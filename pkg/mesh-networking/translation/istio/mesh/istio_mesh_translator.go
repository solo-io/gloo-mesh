package mesh

import (
	"context"

	certificatesv1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2/sets"

	istiov1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"
	"github.com/solo-io/go-utils/contextutils"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/mesh/access"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/mesh/failoverservice"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/mesh/federation"
	"github.com/solo-io/skv2/contrib/pkg/sets"
)

//go:generate mockgen -source ./istio_mesh_translator.go -destination mocks/istio_mesh_translator.go

// outputs of translating a single Mesh
type Outputs struct {
	Gateways              istiov1alpha3sets.GatewaySet
	EnvoyFilters          istiov1alpha3sets.EnvoyFilterSet
	DestinationRules      istiov1alpha3sets.DestinationRuleSet
	ServiceEntries        istiov1alpha3sets.ServiceEntrySet
	AuthorizationPolicies v1beta1sets.AuthorizationPolicySet
	IssuedCertificates    certificatesv1alpha2sets.IssuedCertificateSet
}

// the VirtualService translator translates a Mesh into a VirtualService.
type Translator interface {
	// Translate translates the appropriate resources to apply the VirtualMesh to the given Mesh.
	//
	// Errors caused by invalid user config will be reported using the Reporter.
	Translate(
		in input.Snapshot,
		istioMesh *discoveryv1alpha2.Mesh,
		reporter reporting.Reporter,
	) Outputs
}

type translator struct {
	ctx                       context.Context
	federationTranslator      federation.Translator
	accessTranslator          access.Translator
	failoverServiceTranslator failoverservice.Translator
}

func NewTranslator(
	ctx context.Context,
	federationTranslator federation.Translator,
	accessTranslator access.Translator,
	failoverServiceTranslator failoverservice.Translator,
) Translator {
	return &translator{
		ctx:                       ctx,
		federationTranslator:      federationTranslator,
		accessTranslator:          accessTranslator,
		failoverServiceTranslator: failoverServiceTranslator,
	}
}

// translate the appropriate resources for the given Mesh.
func (t *translator) Translate(
	in input.Snapshot,
	mesh *discoveryv1alpha2.Mesh,
	reporter reporting.Reporter,
) Outputs {
	istioMesh := mesh.Spec.GetIstio()
	if istioMesh == nil {
		contextutils.LoggerFrom(t.ctx).Debugf("ignoring non istio mesh %v %T", sets.Key(mesh), mesh.Spec.MeshType)
		return Outputs{}
	}

	gateways := istiov1alpha3sets.NewGatewaySet()
	envoyFilters := istiov1alpha3sets.NewEnvoyFilterSet()
	destinationRules := istiov1alpha3sets.NewDestinationRuleSet()
	serviceEntries := istiov1alpha3sets.NewServiceEntrySet()
	authPolicies := v1beta1sets.NewAuthorizationPolicySet()
	issuedCertificates := certificatesv1alpha2sets.NewIssuedCertificateSet()

	for _, vMesh := range mesh.Status.AppliedVirtualMeshes {

		federationOutputs := t.federationTranslator.Translate(in, mesh, vMesh, reporter)
		accessAuthPolicies := t.accessTranslator.Translate(mesh, vMesh)

		if federationOutputs.Gateway != nil {
			gateways.Insert(federationOutputs.Gateway)
		}
		if federationOutputs.EnvoyFilter != nil {
			envoyFilters.Insert(federationOutputs.EnvoyFilter)
		}
		destinationRules = destinationRules.Union(federationOutputs.DestinationRules)
		serviceEntries = serviceEntries.Union(federationOutputs.ServiceEntries)
		authPolicies = authPolicies.Union(accessAuthPolicies)
	}

	for _, failoverService := range mesh.Status.AppliedFailoverServices {
		failoverServiceOutputs := t.failoverServiceTranslator.Translate(in, mesh, failoverService, reporter)

		envoyFilters = envoyFilters.Union(failoverServiceOutputs.EnvoyFilters)
		serviceEntries = serviceEntries.Union(failoverServiceOutputs.ServiceEntries)
	}

	return Outputs{
		Gateways:              gateways,
		EnvoyFilters:          envoyFilters,
		DestinationRules:      destinationRules,
		ServiceEntries:        serviceEntries,
		AuthorizationPolicies: authPolicies,
		IssuedCertificates:    issuedCertificates,
	}
}
