package mesh

import (
	"context"

	istiov1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	"github.com/solo-io/go-utils/contextutils"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/snapshot/input"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/mesh/federation"
	"github.com/solo-io/skv2/contrib/pkg/sets"
)

// outputs of translating a single Mesh
type Outputs struct {
	Gateways         istiov1alpha3sets.GatewaySet
	EnvoyFilters     istiov1alpha3sets.EnvoyFilterSet
	DestinationRules istiov1alpha3sets.DestinationRuleSet
	ServiceEntries   istiov1alpha3sets.ServiceEntrySet
}

// the VirtualService translator translates a Mesh into a VirtualService.
type Translator interface {
	// Translate translates the appropriate resources to apply the VirtualMesh to the given Mesh.
	//
	// Errors caused by invalid user config will be reported using the Reporter.
	Translate(
		in input.Snapshot,
		mesh *discoveryv1alpha2.Mesh,
		reporter reporting.Reporter,
	) Outputs
}

type translator struct {
	ctx                  context.Context
	federationTranslator federation.Translator
}

func NewTranslator(ctx context.Context, federationTranslator federation.Translator) Translator {
	return &translator{ctx: ctx, federationTranslator: federationTranslator}
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

	for _, vMesh := range mesh.Status.AppliedVirtualMeshes {
		federationOutputs := t.federationTranslator.Translate(in, mesh, vMesh, reporter)

		if federationOutputs.Gateway != nil {
			gateways.Insert(federationOutputs.Gateway)
		}
		if federationOutputs.EnvoyFilter != nil {
			envoyFilters.Insert(federationOutputs.EnvoyFilter)
		}
		destinationRules = destinationRules.Union(federationOutputs.DestinationRules)
		serviceEntries = serviceEntries.Union(federationOutputs.ServiceEntries)
	}

	return Outputs{
		Gateways:         gateways,
		EnvoyFilters:     envoyFilters,
		DestinationRules: destinationRules,
		ServiceEntries:   serviceEntries,
	}
}
