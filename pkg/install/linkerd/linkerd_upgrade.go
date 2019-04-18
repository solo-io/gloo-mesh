package linkerd

import (
	"fmt"
	"time"

	pb "github.com/linkerd/linkerd2/controller/gen/config"
	"github.com/linkerd/linkerd2/pkg/config"
	"github.com/linkerd/linkerd2/pkg/k8s"
	"github.com/linkerd/linkerd2/pkg/tls"
	"github.com/linkerd/linkerd2/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	okMessage   = "You're on your way to upgrading Linkerd!\nVisit this URL for further instructions: https://linkerd.io/upgrade/#nextsteps\n"
	failMessage = "For troubleshooting help, visit: https://linkerd.io/upgrade/#troubleshooting\n"
)

type (
	upgradeOptions struct{ *installOptions }
)

func newUpgradeOptions(o *installOptions) *upgradeOptions {
	return &upgradeOptions{o}
}

func linkerdAlreadyInstalled(installNamesapce string, k kubernetes.Interface) bool {
	_, err := fetchConfigs(installNamesapce, k)
	return err == nil
}

func (options *upgradeOptions) validateAndBuild(installNamesapce string, k kubernetes.Interface) (*installValues, *pb.All, error) {
	if err := options.validate(); err != nil {
		return nil, nil, err
	}

	// We fetch the configs directly from kubernetes because we need to be able
	// to upgrade/reinstall the control plane when the API is not available; and
	// this also serves as a passive check that we have privileges to access this
	// control plane.
	configs, err := fetchConfigs(installNamesapce, k)
	if err != nil {
		return nil, nil, fmt.Errorf("could not fetch configs from kubernetes: %s", err)
	}

	// If the install config needs to be repaired--either because it did not
	// exist or because it is missing expected fields, repair it.
	repairInstall(func() string {
		return "" //ignore uuid
	}, configs.Install)

	// Update the configs from the synthesized options.
	options.overrideConfigs(configs, map[string]string{})
	if options.proxyAutoInject {
		configs.GetGlobal().AutoInjectContext = &pb.AutoInjectContext{}
	}
	configs.GetInstall().Flags = options.recordedFlags

	var identity *installIdentityValues
	idctx := configs.GetGlobal().GetIdentityContext()
	if idctx.GetTrustDomain() == "" || idctx.GetTrustAnchorsPem() == "" {
		// If there wasn't an idctx, or if it doesn't specify the required fields, we
		// must be upgrading from a version that didn't support identity, so generate it anew...
		identity, err = options.identityOptions.genValues(installNamesapce)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to generate issuer credentials: %s", err)
		}
		configs.GetGlobal().IdentityContext = identity.toIdentityContext()
	} else {
		identity, err = fetchIdentityValues(installNamesapce, k, options.controllerReplicas, idctx)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to fetch the existing issuer credentials from Kubernetes: %s", err)
		}
	}

	// Values have to be generated after any missing identity is generated,
	// otherwise it will be missing from the generated configmap.
	values, err := options.buildValuesWithoutIdentity(configs)
	if err != nil {
		return nil, nil, fmt.Errorf("could not build install configuration: %s", err)
	}
	values.Identity = identity

	return values, configs, nil
}

func repairInstall(generateUUID func() string, install *pb.Install) {
	if install == nil {
		install = &pb.Install{}
	}

	if install.GetUuid() == "" {
		install.Uuid = generateUUID()
	}

	// ALWAYS update the CLI version to the most recent.
	install.CliVersion = version.Version

	// Install flags are updated separately.
}

// fetchConfigs checks the kubernetes API to fetch an existing
// linkerd configuration.
//
// This bypasses the public API so that upgrades can proceed when the API pod is
// not available.
func fetchConfigs(namespace string, k kubernetes.Interface) (*pb.All, error) {
	configMap, err := k.CoreV1().
		ConfigMaps(namespace).
		Get(k8s.ConfigConfigMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return config.FromConfigMap(configMap.Data)
}

// fetchIdentityValue checks the kubernetes API to fetch an existing
// linkerd identity configuration.
//
// This bypasses the public API so that we can access secrets and validate
// permissions.
func fetchIdentityValues(namespace string, k kubernetes.Interface, replicas uint, idctx *pb.IdentityContext) (*installIdentityValues, error) {
	if idctx == nil {
		return nil, nil
	}

	keyPEM, crtPEM, expiry, err := fetchIssuer(namespace, k, idctx.GetTrustAnchorsPem())
	if err != nil {
		return nil, err
	}

	return &installIdentityValues{
		Replicas:        replicas,
		TrustDomain:     idctx.GetTrustDomain(),
		TrustAnchorsPEM: idctx.GetTrustAnchorsPem(),
		Issuer: &issuerValues{
			ClockSkewAllowance:  idctx.GetClockSkewAllowance().String(),
			IssuanceLifetime:    idctx.GetIssuanceLifetime().String(),
			CrtExpiryAnnotation: k8s.IdentityIssuerExpiryAnnotation,

			KeyPEM:    keyPEM,
			CrtPEM:    crtPEM,
			CrtExpiry: expiry,
		},
	}, nil
}

func fetchIssuer(namespace string, k kubernetes.Interface, trustPEM string) (string, string, time.Time, error) {
	roots, err := tls.DecodePEMCertPool(trustPEM)
	if err != nil {
		return "", "", time.Time{}, err
	}

	secret, err := k.CoreV1().
		Secrets(namespace).
		Get(k8s.IdentityIssuerSecretName, metav1.GetOptions{})
	if err != nil {
		return "", "", time.Time{}, err
	}

	keyPEM := string(secret.Data[k8s.IdentityIssuerKeyName])
	key, err := tls.DecodePEMKey(keyPEM)
	if err != nil {
		return "", "", time.Time{}, err
	}

	crtPEM := string(secret.Data[k8s.IdentityIssuerCrtName])
	crt, err := tls.DecodePEMCrt(crtPEM)
	if err != nil {
		return "", "", time.Time{}, err
	}

	cred := &tls.Cred{PrivateKey: key, Crt: *crt}
	if err = cred.Verify(roots, ""); err != nil {
		return "", "", time.Time{}, fmt.Errorf("invalid issuer credentials: %s", err)
	}

	return keyPEM, crtPEM, crt.Certificate.NotAfter, nil
}
