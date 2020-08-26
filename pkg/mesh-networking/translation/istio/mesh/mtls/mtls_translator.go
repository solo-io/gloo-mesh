package mtls

import (
	"context"
	"fmt"
	"time"

	discoveryv1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/output/local"
	"github.com/solo-io/skv2/pkg/ezkube"

	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"

	"github.com/solo-io/service-mesh-hub/pkg/common/defaults"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/output/istio"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/certificates/common/secrets"
	"istio.io/istio/pkg/spiffe"
	"istio.io/istio/security/pkg/pki/util"
	corev1 "k8s.io/api/core/v1"

	"github.com/solo-io/go-utils/contextutils"
	certificatesv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

//go:generate mockgen -source ./mtls_translator.go -destination mocks/mtls_translator.go

const (
	defaultIstioOrg              = "Istio"
	defaultCitadelServiceAccount = "istio-citadel"
	defaultTrustDomain           = "cluster.local" // The default SPIFFE URL value for trust domain
	defaultIstioNamespace        = "istio-system"
	// name of the istio root CA secret
	// https://istio.io/latest/docs/tasks/security/cert-management/plugin-ca-cert/
	istioCaSecretName = "cacerts"
)

var (
	signingCertSecretType = corev1.SecretType(fmt.Sprintf("%s/generated_signing_cert", certificatesv1alpha2.SchemeGroupVersion.Group))
)

// used by networking reconciler to filter ignored secrets
func IsSigningCert(secret *corev1.Secret) bool {
	return secret.Type == signingCertSecretType
}

// the VirtualService translator translates a Mesh into a VirtualService.
type Translator interface {
	// Translate translates the appropriate VirtualService and DestinationRule for the given Mesh.
	// returns nil if no VirtualService or DestinationRule is required for the Mesh (i.e. if no VirtualService/DestinationRule features are required, such as subsets).
	// Output resources will be added to the istio.Builder
	// Errors caused by invalid user config will be reported using the Reporter.
	Translate(
		mesh *discoveryv1alpha2.Mesh,
		virtualMesh *discoveryv1alpha2.MeshStatus_AppliedVirtualMesh,
		istioOutputs istio.Builder,
		localOutputs local.Builder,
		reporter reporting.Reporter,
	)
}

type translator struct {
	ctx       context.Context
	secrets   corev1sets.SecretSet
	workloads discoveryv1alpha2sets.WorkloadSet
}

func NewTranslator(ctx context.Context, secrets corev1sets.SecretSet, workloads discoveryv1alpha2sets.WorkloadSet) Translator {
	return &translator{
		ctx:       ctx,
		secrets:   secrets,
		workloads: workloads,
	}
}

// translate the appropriate resources for the given Mesh.
func (t *translator) Translate(
	mesh *discoveryv1alpha2.Mesh,
	virtualMesh *discoveryv1alpha2.MeshStatus_AppliedVirtualMesh,
	istioOutputs istio.Builder,
	localOutputs local.Builder,
	reporter reporting.Reporter,
) {
	istioMesh := mesh.Spec.GetIstio()
	if istioMesh == nil {
		contextutils.LoggerFrom(t.ctx).Debugf("ignoring non istio mesh %v %T", sets.Key(mesh), mesh.Spec.MeshType)
		return
	}

	if virtualMesh == nil || virtualMesh.Spec.MtlsConfig == nil {
		contextutils.LoggerFrom(t.ctx).Debugf("no translation for virtual mesh %v which has no mTLS configuration", sets.Key(mesh))
		return
	}
	mtlsConfig := virtualMesh.Spec.MtlsConfig

	// TODO(ilackarms): currently we assume a shared trust model
	// we'll want to expand this to support limited trust in the future
	sharedTrust := mtlsConfig.GetShared()
	rootCA := sharedTrust.GetRootCertificateAuthority()
	agentInfo := mesh.Spec.AgentInfo
	if agentInfo == nil {
		contextutils.LoggerFrom(t.ctx).Debugf("cannot configure root certificates for mesh %v which has no cert-agent", sets.Key(mesh))
		return
	}

	var rootCaSecret *v1.ObjectRef
	switch caType := rootCA.CaSource.(type) {
	case *v1alpha2.VirtualMeshSpec_RootCertificateAuthority_Generated:
		generatedSecretName := virtualMesh.Ref.Name + "." + virtualMesh.Ref.Namespace
		// write the signing secret to the smh namespace
		generatedSecretNamespace := defaults.GetPodNamespace()
		// use the existing secret if it exists
		rootCaSecret = &v1.ObjectRef{
			Name:      generatedSecretName,
			Namespace: generatedSecretNamespace,
		}
		selfSignedCertSecret, err := t.secrets.Find(rootCaSecret)
		if err != nil {
			selfSignedCert, err := generateSelfSignedCert(caType.Generated)
			if err != nil {
				// should never happen
				reporter.ReportVirtualMeshToMesh(mesh, virtualMesh.Ref, err)
				return
			}
			// the self signed cert goes to the master/local cluster
			selfSignedCertSecret = &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: generatedSecretName,
					// write to the agent namespace
					Namespace: generatedSecretNamespace,
					// ensure the secret is written to the maser/local cluster
					ClusterName: "",
					Labels:      metautils.TranslatedObjectLabels(),
				},
				Data: selfSignedCert.ToSecretData(),
				Type: signingCertSecretType,
			}
		}
		localOutputs.AddSecrets(selfSignedCertSecret)
	case *v1alpha2.VirtualMeshSpec_RootCertificateAuthority_Secret:
		rootCaSecret = caType.Secret
	}

	trustDomain := istioMesh.GetCitadelInfo().GetTrustDomain()
	if trustDomain == "" {
		trustDomain = defaultTrustDomain
	}
	citadelServiceAccount := istioMesh.GetCitadelInfo().GetCitadelServiceAccount()
	if citadelServiceAccount == "" {
		citadelServiceAccount = defaultCitadelServiceAccount
	}
	istioNamespace := istioMesh.GetInstallation().GetNamespace()
	if istioNamespace == "" {
		istioNamespace = defaultIstioNamespace
	}

	// the default location of the istio CA Certs secret
	// the certificate workflow will produce a cert with this ref
	istioCaCerts := &v1.ObjectRef{
		Name:      istioCaSecretName,
		Namespace: istioNamespace,
	}

	// get the pods that need to be bounced for this mesh
	podsToBounce := getPodsToBounce(mesh, t.workloads, mtlsConfig.AutoRestartPods)

	// issue a certificate to the mesh agent
	issuedCertificate := &certificatesv1alpha2.IssuedCertificate{
		ObjectMeta: metav1.ObjectMeta{
			Name: mesh.Name,
			// write to the agent namespace
			Namespace: agentInfo.AgentNamespace,
			// write to the mesh cluster
			ClusterName: istioMesh.GetInstallation().GetCluster(),
			Labels:      metautils.TranslatedObjectLabels(),
		},
		Spec: certificatesv1alpha2.IssuedCertificateSpec{
			Hosts:                    []string{buildSpiffeURI(trustDomain, istioNamespace, citadelServiceAccount)},
			Org:                      defaultIstioOrg,
			SigningCertificateSecret: rootCaSecret,
			IssuedCertificateSecret:  istioCaCerts,
			PodsToBounce:             podsToBounce,
		},
	}
	istioOutputs.AddIssuedCertificates(issuedCertificate)
}

const (
	defaultRootCertTTLDays     = 365
	defaultRootCertTTLDuration = defaultRootCertTTLDays * 24 * time.Hour
	defaultRootCertRsaKeySize  = 4096
	defaultOrgName             = "service-mesh-hub"
)

func generateSelfSignedCert(
	builtinCA *v1alpha2.VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert,
) (*secrets.RootCAData, error) {
	org := defaultOrgName
	if builtinCA.GetOrgName() != "" {
		org = builtinCA.GetOrgName()
	}
	ttl := defaultRootCertTTLDuration
	if builtinCA.GetTtlDays() > 0 {
		ttl = time.Duration(builtinCA.GetTtlDays()) * 24 * time.Hour
	}
	rsaKeySize := defaultRootCertRsaKeySize
	if builtinCA.GetRsaKeySizeBytes() > 0 {
		rsaKeySize = int(builtinCA.GetRsaKeySizeBytes())
	}
	options := util.CertOptions{
		Org:          org,
		IsCA:         true,
		IsSelfSigned: true,
		TTL:          ttl,
		RSAKeySize:   rsaKeySize,
		PKCS8Key:     false, // currently only supporting PKCS1
	}
	cert, key, err := util.GenCertKeyFromOptions(options)
	if err != nil {
		return nil, err
	}
	rootCaData := &secrets.RootCAData{
		PrivateKey: key,
		RootCert:   cert,
	}
	return rootCaData, nil
}

func buildSpiffeURI(trustDomain, namespace, serviceAccount string) string {
	return fmt.Sprintf("%s%s/ns/%s/sa/%s", spiffe.URIPrefix, trustDomain, namespace, serviceAccount)
}

// get selectors for all the pods in a mesh; they need to be bounced (including the mesh control plane itself)
func getPodsToBounce(mesh *discoveryv1alpha2.Mesh, allWorkloads discoveryv1alpha2sets.WorkloadSet, autoRestartPods bool) []*certificatesv1alpha2.IssuedCertificateSpec_PodSelector {
	// if autoRestartPods is false, we rely on the user to manually restart their pods
	if !autoRestartPods {
		return nil
	}
	istioMesh := mesh.Spec.GetIstio()
	istioInstall := istioMesh.GetInstallation()

	// bounce the control plane pod
	podsToBounce := []*certificatesv1alpha2.IssuedCertificateSpec_PodSelector{
		{
			Namespace: istioInstall.Namespace,
			Labels:    istioInstall.PodLabels,
		},
	}

	// bounce the ingress gateway pods
	for _, gateway := range istioMesh.IngressGateways {
		podsToBounce = append(podsToBounce, &certificatesv1alpha2.IssuedCertificateSpec_PodSelector{
			Namespace: istioInstall.Namespace,
			Labels:    gateway.WorkloadLabels,
		})
	}

	// collect selectors from workloads matching this mesh
	allWorkloads.List(func(workload *discoveryv1alpha2.Workload) bool {
		kubeWorkload := workload.Spec.GetKubernetes()

		if kubeWorkload != nil && ezkube.RefsMatch(workload.Spec.Mesh, mesh) {
			podsToBounce = append(podsToBounce, &certificatesv1alpha2.IssuedCertificateSpec_PodSelector{
				Namespace: kubeWorkload.Controller.GetNamespace(),
				Labels:    kubeWorkload.PodLabels,
			})
		}

		return false
	})

	return podsToBounce
}
