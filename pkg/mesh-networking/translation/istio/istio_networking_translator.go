package istio

import (
	"context"
	"fmt"

	certificatesv1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2/sets"

	v1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/reporting"
	"github.com/solo-io/skv2/contrib/pkg/sets"
)

// output types of Istio translation
type Outputs struct {
	IssuedCertificates    certificatesv1alpha2sets.IssuedCertificateSet
	DestinationRules      v1alpha3sets.DestinationRuleSet
	EnvoyFilters          v1alpha3sets.EnvoyFilterSet
	Gateways              v1alpha3sets.GatewaySet
	ServiceEntries        v1alpha3sets.ServiceEntrySet
	VirtualServices       v1alpha3sets.VirtualServiceSet
	AuthorizationPolicies v1beta1sets.AuthorizationPolicySet
}

// the istio translator translates an input networking snapshot to an output snapshot of Istio resources
type Translator interface {
	Translate(
		ctx context.Context,
		in input.Snapshot,
		reporter reporting.Reporter,
	) Outputs
}

type istioTranslator struct {
	totalTranslates int // TODO(ilackarms): metric
	dependencies    dependencyFactory
}

func NewIstioTranslator() Translator {
	return &istioTranslator{
		dependencies: dependencyFactoryImpl{},
	}
}

func (t *istioTranslator) Translate(
	ctx context.Context,
	in input.Snapshot,
	reporter reporting.Reporter,
) Outputs {
	ctx = contextutils.WithLogger(ctx, fmt.Sprintf("istio-translator-%v", t.totalTranslates))

	meshServiceTranslator := t.dependencies.makeMeshServiceTranslator(in.KubernetesClusters())

	destinationRules := v1alpha3sets.NewDestinationRuleSet()
	virtualServices := v1alpha3sets.NewVirtualServiceSet()
	authorizationPolicies := v1beta1sets.NewAuthorizationPolicySet()

	for _, meshService := range in.MeshServices().List() {
		meshService := meshService // pike

		serviceOutputs := meshServiceTranslator.Translate(in, meshService, reporter)

		destinationRule := serviceOutputs.DestinationRule
		if destinationRule != nil {
			destinationRules.Insert(destinationRule)
			contextutils.LoggerFrom(ctx).Debugf("translated destination rule %v", sets.Key(destinationRule))
		}

		virtualService := serviceOutputs.VirtualService
		if virtualService != nil {
			contextutils.LoggerFrom(ctx).Debugf("translated virtual service %v", sets.Key(virtualService))
			virtualServices.Insert(virtualService)
		}

		authorizationPolicy := serviceOutputs.AuthorizationPolicy
		if authorizationPolicy != nil {
			contextutils.LoggerFrom(ctx).Debugf("translated authorization policy %v", sets.Key(authorizationPolicy))
			authorizationPolicies.Insert(authorizationPolicy)
		}
	}

	envoyFilters := v1alpha3sets.NewEnvoyFilterSet()
	gateways := v1alpha3sets.NewGatewaySet()
	serviceEntries := v1alpha3sets.NewServiceEntrySet()
	issuedCertificates := certificatesv1alpha2sets.NewIssuedCertificateSet()

	meshTranslator := t.dependencies.makeMeshTranslator(ctx, in.KubernetesClusters())
	for _, mesh := range in.Meshes().List() {
		meshOutputs := meshTranslator.Translate(in, mesh, reporter)

		gateways = gateways.Union(meshOutputs.Gateways)
		serviceEntries = serviceEntries.Union(meshOutputs.ServiceEntries)
		envoyFilters = envoyFilters.Union(meshOutputs.EnvoyFilters)
		destinationRules = destinationRules.Union(meshOutputs.DestinationRules)
		issuedCertificates = issuedCertificates.Union(meshOutputs.IssuedCertificates)
	}

	t.totalTranslates++

	return Outputs{
		IssuedCertificates:    issuedCertificates,
		DestinationRules:      destinationRules,
		EnvoyFilters:          envoyFilters,
		Gateways:              gateways,
		ServiceEntries:        serviceEntries,
		VirtualServices:       virtualServices,
		AuthorizationPolicies: authorizationPolicies,
	}
}
