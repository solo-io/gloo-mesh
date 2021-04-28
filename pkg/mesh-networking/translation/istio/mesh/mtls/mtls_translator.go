package mtls

import (
	"context"
	"fmt"
	"time"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	"github.com/solo-io/gloo-mesh/pkg/common/version"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"

	discoveryv1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/local"
	"github.com/solo-io/skv2/pkg/ezkube"

	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"

	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/certificates/common/secrets"
	"istio.io/istio/pkg/spiffe"
	"istio.io/istio/security/pkg/pki/util"
	corev1 "k8s.io/api/core/v1"

	certificatesv1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
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
	// name of the istio root CA configmap distributed to all namespaces
	// copied from https://github.com/istio/istio/blob/88a2bfb/pilot/pkg/serviceregistry/kube/controller/namespacecontroller.go#L39
	// not imported due to issues with dependeny imports
	istioCaConfigMapName = "istio-ca-root-cert"
)

var (
	signingCertSecretType = corev1.SecretType(fmt.Sprintf("%s/generated_signing_cert", certificatesv1.SchemeGroupVersion.Group))

	// used when the user provides a nil root cert
	defaultSelfSignedRootCa = &v1.VirtualMeshSpec_RootCertificateAuthority{
		CaSource: &v1.VirtualMeshSpec_RootCertificateAuthority_Generated{
			Generated: &v1.VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert{
				TtlDays:         defaultRootCertTTLDays,
				RsaKeySizeBytes: defaultRootCertRsaKeySize,
				OrgName:         defaultOrgName,
			},
		},
	}
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
		mesh *discoveryv1.Mesh,
		virtualMesh *discoveryv1.MeshStatus_AppliedVirtualMesh,
		istioOutputs istio.Builder,
		localOutputs local.Builder,
		reporter reporting.Reporter,
	)
}

type translator struct {
	ctx       context.Context
	secrets   corev1sets.SecretSet
	workloads discoveryv1sets.WorkloadSet
}

func NewTranslator(ctx context.Context, secrets corev1sets.SecretSet, workloads discoveryv1sets.WorkloadSet) Translator {
	return &translator{
		ctx:       ctx,
		secrets:   secrets,
		workloads: workloads,
	}
}

// translate the appropriate resources for the given Mesh.
func (t *translator) Translate(
	mesh *discoveryv1.Mesh,
	virtualMesh *discoveryv1.MeshStatus_AppliedVirtualMesh,
	istioOutputs istio.Builder,
	localOutputs local.Builder,
	reporter reporting.Reporter,
) {
	istioMesh := mesh.Spec.GetIstio()
	if istioMesh == nil {
		contextutils.LoggerFrom(t.ctx).Debugf("ignoring non istio mesh %v %T", sets.Key(mesh), mesh.Spec.Type)
		return
	}

	if err := t.updateMtlsOutputs(mesh, virtualMesh, istioOutputs, localOutputs); err != nil {
		reporter.ReportVirtualMeshToMesh(mesh, virtualMesh.Ref, err)
	}
}

func (t *translator) updateMtlsOutputs(
	mesh *discoveryv1.Mesh,
	virtualMesh *discoveryv1.MeshStatus_AppliedVirtualMesh,
	istioOutputs istio.Builder,
	localOutputs local.Builder,
) error {
	mtlsConfig := virtualMesh.Spec.MtlsConfig
	if mtlsConfig == nil {
		// nothing to do
		contextutils.LoggerFrom(t.ctx).Debugf("no translation for VirtualMesh %v which has no mTLS configuration", sets.Key(mesh))
		return nil
	}

	if mtlsConfig.TrustModel == nil {
		return eris.Errorf("must specify trust model to use for issuing certificates")
	}

	switch trustModel := mtlsConfig.TrustModel.(type) {
	case *v1.VirtualMeshSpec_MTLSConfig_Shared:
		return t.configureSharedTrust(
			mesh,
			trustModel.Shared,
			virtualMesh.Ref,
			istioOutputs,
			localOutputs,
			mtlsConfig.AutoRestartPods,
		)
	case *v1.VirtualMeshSpec_MTLSConfig_Limited:
		return eris.Errorf("limited trust not supported in version %v of Gloo Mesh", version.Version)
	}

	return nil
}

// will create the secret if it is self-signed,
// otherwise will return the user-provided secret ref in the mtls config
func (t *translator) configureSharedTrust(
	mesh *discoveryv1.Mesh,
	sharedTrust *v1.VirtualMeshSpec_MTLSConfig_SharedTrust,
	virtualMeshRef *skv2corev1.ObjectRef,
	istioOutputs istio.Builder,
	localOutputs local.Builder,
	autoRestartPods bool,
) error {
	rootCA := sharedTrust.GetRootCertificateAuthority()

	rootCaSecret, err := t.getOrCreateRootCaSecret(
		rootCA,
		virtualMeshRef,
		localOutputs,
	)
	if err != nil {
		return err
	}

	agentInfo := mesh.Spec.AgentInfo
	if agentInfo == nil {
		contextutils.LoggerFrom(t.ctx).Debugf("cannot configure root certificates for mesh %v which has no cert-agent", sets.Key(mesh))
		return nil
	}

	issuedCertificate, podBounceDirective := t.constructIssuedCertificate(
		mesh,
		rootCaSecret,
		agentInfo.AgentNamespace,
		autoRestartPods,
	)

	// Append the VirtualMesh as a parent to each output resource
	metautils.AppendParent(t.ctx, issuedCertificate, virtualMeshRef, v1.VirtualMesh{}.GVK())
	metautils.AppendParent(t.ctx, podBounceDirective, virtualMeshRef, v1.VirtualMesh{}.GVK())

	istioOutputs.AddIssuedCertificates(issuedCertificate)
	istioOutputs.AddPodBounceDirectives(podBounceDirective)
	return nil
}

// will create the secret if it is self-signed,
// otherwise will return the user-provided secret ref in the mtls config
func (t *translator) getOrCreateRootCaSecret(
	rootCA *v1.VirtualMeshSpec_RootCertificateAuthority,
	virtualMeshRef *skv2corev1.ObjectRef,
	localOutputs local.Builder,
) (*skv2corev1.ObjectRef, error) {
	if rootCA == nil || rootCA.CaSource == nil {
		rootCA = defaultSelfSignedRootCa
	}

	var rootCaSecret *skv2corev1.ObjectRef
	switch caType := rootCA.CaSource.(type) {
	case *v1.VirtualMeshSpec_RootCertificateAuthority_Generated:
		generatedSecretName := virtualMeshRef.Name + "." + virtualMeshRef.Namespace
		// write the signing secret to the gloomesh namespace
		generatedSecretNamespace := defaults.GetPodNamespace()
		// use the existing secret if it exists
		rootCaSecret = &skv2corev1.ObjectRef{
			Name:      generatedSecretName,
			Namespace: generatedSecretNamespace,
		}
		selfSignedCertSecret, err := t.secrets.Find(rootCaSecret)
		if err != nil {
			selfSignedCert, err := generateSelfSignedCert(caType.Generated)
			if err != nil {
				// should never happen
				return nil, err
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

		// Append the VirtualMesh as a parent to the output secret
		metautils.AppendParent(t.ctx, selfSignedCertSecret, virtualMeshRef, v1.VirtualMesh{}.GVK())

		localOutputs.AddSecrets(selfSignedCertSecret)
	case *v1.VirtualMeshSpec_RootCertificateAuthority_Secret:
		rootCaSecret = caType.Secret
	}

	return rootCaSecret, nil
}

func (t *translator) constructIssuedCertificate(
	mesh *discoveryv1.Mesh,
	rootCaSecret *skv2corev1.ObjectRef,
	agentNamespace string,
	autoRestartPods bool,
) (*certificatesv1.IssuedCertificate, *certificatesv1.PodBounceDirective) {
	istioMesh := mesh.Spec.GetIstio()

	trustDomain := istioMesh.GetTrustDomain()
	if trustDomain == "" {
		trustDomain = defaultTrustDomain
	}
	istiodServiceAccount := istioMesh.GetIstiodServiceAccount()
	if istiodServiceAccount == "" {
		istiodServiceAccount = defaultCitadelServiceAccount
	}
	istioNamespace := istioMesh.GetInstallation().GetNamespace()
	if istioNamespace == "" {
		istioNamespace = defaultIstioNamespace
	}

	// the default location of the istio CA Certs secret
	// the certificate workflow will produce a cert with this ref
	istioCaCerts := &skv2corev1.ObjectRef{
		Name:      istioCaSecretName,
		Namespace: istioNamespace,
	}

	clusterName := istioMesh.GetInstallation().GetCluster()
	issuedCertificateMeta := metav1.ObjectMeta{
		Name: mesh.Name,
		// write to the agent namespace
		Namespace: agentNamespace,
		// write to the mesh cluster
		ClusterName: clusterName,
		Labels:      metautils.TranslatedObjectLabels(),
	}

	// get the pods that need to be bounced for this mesh
	podsToBounce := getPodsToBounce(mesh, t.workloads, autoRestartPods)
	var (
		podBounceDirective *certificatesv1.PodBounceDirective
		podBounceRef       *skv2corev1.ObjectRef
	)
	if len(podsToBounce) > 0 {
		podBounceDirective = &certificatesv1.PodBounceDirective{
			ObjectMeta: issuedCertificateMeta,
			Spec: certificatesv1.PodBounceDirectiveSpec{
				PodsToBounce: podsToBounce,
			},
		}
		podBounceRef = ezkube.MakeObjectRef(podBounceDirective)
	}

	// issue a certificate to the mesh agent
	return &certificatesv1.IssuedCertificate{
		ObjectMeta: issuedCertificateMeta,
		Spec: certificatesv1.IssuedCertificateSpec{
			Hosts:                    []string{buildSpiffeURI(trustDomain, istioNamespace, istiodServiceAccount)},
			Org:                      defaultIstioOrg,
			SigningCertificateSecret: rootCaSecret,
			IssuedCertificateSecret:  istioCaCerts,
			PodBounceDirective:       podBounceRef,
		},
	}, podBounceDirective
}

const (
	defaultRootCertTTLDays     = 365
	defaultRootCertTTLDuration = defaultRootCertTTLDays * 24 * time.Hour
	defaultRootCertRsaKeySize  = 4096
	defaultOrgName             = "gloo-mesh"
)

func generateSelfSignedCert(
	builtinCA *v1.VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert,
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
func getPodsToBounce(mesh *discoveryv1.Mesh, allWorkloads discoveryv1sets.WorkloadSet, autoRestartPods bool) []*certificatesv1.PodBounceDirectiveSpec_PodSelector {
	// if autoRestartPods is false, we rely on the user to manually restart their pods
	if !autoRestartPods {
		return nil
	}
	istioMesh := mesh.Spec.GetIstio()
	istioInstall := istioMesh.GetInstallation()

	// bounce the control plane pod first
	// order matters
	podsToBounce := []*certificatesv1.PodBounceDirectiveSpec_PodSelector{
		{
			Namespace: istioInstall.Namespace,
			Labels:    istioInstall.PodLabels,
			// ensure at least one replica of istiod is ready before restarting the other pods
			WaitForReplicas: 1,
		},
	}

	// bounce the ingress gateway pods
	for _, gateway := range istioMesh.IngressGateways {
		podsToBounce = append(podsToBounce, &certificatesv1.PodBounceDirectiveSpec_PodSelector{
			Namespace: istioInstall.Namespace,
			Labels:    gateway.WorkloadLabels,
			RootCertSync: &certificatesv1.PodBounceDirectiveSpec_PodSelector_RootCertSync{
				SecretRef: &skv2corev1.ObjectRef{
					Name:      istioCaSecretName,
					Namespace: istioInstall.Namespace,
				},
				SecretKey: secrets.RootCertID,
				ConfigMapRef: &skv2corev1.ObjectRef{
					Name:      istioCaConfigMapName,
					Namespace: istioInstall.Namespace,
				},
				ConfigMapKey: secrets.RootCertID,
			},
		})
	}

	// collect selectors from workloads matching this mesh
	allWorkloads.List(func(workload *discoveryv1.Workload) bool {
		kubeWorkload := workload.Spec.GetKubernetes()

		if kubeWorkload != nil && ezkube.RefsMatch(workload.Spec.Mesh, mesh) {
			podsToBounce = append(podsToBounce, &certificatesv1.PodBounceDirectiveSpec_PodSelector{
				Namespace: kubeWorkload.Controller.GetNamespace(),
				Labels:    kubeWorkload.PodLabels,
				RootCertSync: &certificatesv1.PodBounceDirectiveSpec_PodSelector_RootCertSync{
					SecretRef: &skv2corev1.ObjectRef{
						Name:      istioCaSecretName,
						Namespace: istioInstall.Namespace,
					},
					SecretKey: secrets.RootCertID,
					ConfigMapRef: &skv2corev1.ObjectRef{
						Name:      istioCaConfigMapName,
						Namespace: kubeWorkload.Controller.GetNamespace(),
					},
					ConfigMapKey: secrets.RootCertID,
				},
			})
		}

		return false
	})

	return podsToBounce
}
