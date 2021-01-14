package destinationrule

import (
	"context"
	"reflect"

	settingsv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2"

	v1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	discoveryv1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/decorators/tls"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/decorators/trafficshift"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/traffictarget/utils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/selectorutils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/rotisserie/eris"
	discoveryv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/decorators"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/fieldutils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/hostutils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/skv2/pkg/equalityutils"
	"github.com/solo-io/skv2/pkg/ezkube"
	networkingv1alpha3spec "istio.io/api/networking/v1alpha3"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
)

//go:generate mockgen -source ./destination_rule_translator.go -destination mocks/destination_rule_translator.go

// the DestinationRule translator translates a TrafficTarget into a DestinationRule.
type Translator interface {
	// Translate translates the appropriate DestinationRule for the given TrafficTarget.
	// returns nil if no DestinationRule is required for the TrafficTarget (i.e. if no DestinationRule features are required, such as subsets).
	//
	// Errors caused by invalid user config will be reported using the Reporter.
	//
	// Note that the input snapshot TrafficTargetSet contains the given TrafficTarget.
	Translate(
		ctx context.Context,
		in input.LocalSnapshot,
		trafficTarget *discoveryv1alpha2.TrafficTarget,
		sourceMeshInstallation *discoveryv1alpha2.MeshSpec_MeshInstallation,
		reporter reporting.Reporter,
	) *networkingv1alpha3.DestinationRule
}

type translator struct {
	settings             *settingsv1alpha2.Settings
	userDestinationRules v1alpha3sets.DestinationRuleSet
	clusterDomains       hostutils.ClusterDomainRegistry
	decoratorFactory     decorators.Factory
	trafficTargets       discoveryv1alpha2sets.TrafficTargetSet
	failoverServices     v1alpha2sets.FailoverServiceSet
}

func NewTranslator(
	settings *settingsv1alpha2.Settings,
	userDestinationRules v1alpha3sets.DestinationRuleSet,
	clusterDomains hostutils.ClusterDomainRegistry,
	decoratorFactory decorators.Factory,
	trafficTargets discoveryv1alpha2sets.TrafficTargetSet,
	failoverServices v1alpha2sets.FailoverServiceSet,
) Translator {
	return &translator{
		settings:             settings,
		userDestinationRules: userDestinationRules,
		clusterDomains:       clusterDomains,
		decoratorFactory:     decoratorFactory,
		trafficTargets:       trafficTargets,
		failoverServices:     failoverServices,
	}
}

// translate the appropriate DestinationRule for the given TrafficTarget.
// returns nil if no DestinationRule is required for the TrafficTarget (i.e. if no DestinationRule features are required, such as subsets).
// The input snapshot TrafficTargetSet contains n the
func (t *translator) Translate(
	ctx context.Context,
	in input.LocalSnapshot,
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	sourceMeshInstallation *discoveryv1alpha2.MeshSpec_MeshInstallation,
	reporter reporting.Reporter,
) *networkingv1alpha3.DestinationRule {
	kubeService := trafficTarget.Spec.GetKubeService()

	if kubeService == nil {
		// TODO(ilackarms): non kube services currently unsupported
		return nil
	}

	sourceClusterName := kubeService.Ref.ClusterName
	if sourceMeshInstallation != nil {
		sourceClusterName = sourceMeshInstallation.Cluster
	}

	destinationRule, err := t.initializeDestinationRule(trafficTarget, t.settings.Spec.Mtls, sourceMeshInstallation)
	if err != nil {
		contextutils.LoggerFrom(ctx).Error(err)
		return nil
	}

	// register the owners of the destinationrule fields
	destinationRuleFields := fieldutils.NewOwnershipRegistry()
	drDecorators := t.decoratorFactory.MakeDecorators(decorators.Parameters{
		ClusterDomains: t.clusterDomains,
		Snapshot:       in,
	})

	// Apply decorators which map a single applicable TrafficPolicy to a field on the DestinationRule.
	for _, policy := range trafficTarget.Status.AppliedTrafficPolicies {

		// Don't translate the trafficPolicy if the sourceClusterName is not selected by the SourceSelectors
		if !selectorutils.WorkloadSelectorContainsCluster(policy.Spec.SourceSelector, sourceClusterName) {
			continue
		}

		registerField := registerFieldFunc(destinationRuleFields, destinationRule, policy.Ref)
		for _, decorator := range drDecorators {

			if destinationRuleDecorator, ok := decorator.(decorators.TrafficPolicyDestinationRuleDecorator); ok {
				if err := destinationRuleDecorator.ApplyTrafficPolicyToDestinationRule(
					policy,
					trafficTarget,
					&destinationRule.Spec,
					registerField,
				); err != nil {
					reporter.ReportTrafficPolicyToTrafficTarget(trafficTarget, policy.Ref, eris.Wrapf(err, "%v", decorator.DecoratorName()))
				}
			}
		}
	}

	// TODO need a more robust implementation of determining whether a DestinationRule has any effect
	if len(destinationRule.Spec.Subsets) == 0 &&
		destinationRule.Spec.GetTrafficPolicy().GetTls().GetMode() == networkingv1alpha3spec.ClientTLSSettings_DISABLE &&
		destinationRule.Spec.GetTrafficPolicy().GetOutlierDetection() == nil {
		// no need to create this DestinationRule as it has no effect
		return nil
	}

	if t.userDestinationRules == nil {
		return destinationRule
	}

	// detect and report error on intersecting config if enabled in settings
	if errs := conflictsWithUserDestinationRule(
		t.userDestinationRules,
		destinationRule,
	); len(errs) > 0 {
		for _, err := range errs {
			for _, policy := range trafficTarget.Status.AppliedTrafficPolicies {
				reporter.ReportTrafficPolicyToTrafficTarget(trafficTarget, policy.Ref, err)
			}
		}
		return nil
	}

	return destinationRule
}

// construct the callback for registering fields in the virtual service
func registerFieldFunc(
	destinationRuleFields fieldutils.FieldOwnershipRegistry,
	destinationRule *networkingv1alpha3.DestinationRule,
	policy ezkube.ResourceId,
) decorators.RegisterField {
	return func(fieldPtr, val interface{}) error {
		fieldVal := reflect.ValueOf(fieldPtr).Elem().Interface()

		if equalityutils.DeepEqual(fieldVal, val) {
			return nil
		}
		if err := destinationRuleFields.RegisterFieldOwnership(
			destinationRule,
			fieldPtr,
			[]ezkube.ResourceId{policy},
			&v1alpha2.TrafficPolicy{},
			0, //TODO(ilackarms): priority
		); err != nil {
			return err
		}
		return nil
	}
}

func (t *translator) initializeDestinationRule(
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	mtlsDefault *v1alpha2.TrafficPolicySpec_MTLS,
	sourceMeshInstallation *discoveryv1alpha2.MeshSpec_MeshInstallation,
) (*networkingv1alpha3.DestinationRule, error) {
	var meta metav1.ObjectMeta
	if sourceMeshInstallation != nil {
		meta = metautils.FederatedObjectMeta(
			trafficTarget.Spec.GetKubeService().Ref,
			sourceMeshInstallation,
			trafficTarget.Annotations,
		)
	} else {
		meta = metautils.TranslatedObjectMeta(
			trafficTarget.Spec.GetKubeService().Ref,
			trafficTarget.Annotations,
		)
	}
	hostname := t.clusterDomains.GetDestinationFQDN(meta.ClusterName, trafficTarget.Spec.GetKubeService().Ref)

	destinationRule := &networkingv1alpha3.DestinationRule{
		ObjectMeta: meta,
		Spec: networkingv1alpha3spec.DestinationRule{
			Host:          hostname,
			TrafficPolicy: &networkingv1alpha3spec.TrafficPolicy{},
			Subsets: trafficshift.MakeDestinationRuleSubsetsForTrafficTarget(
				trafficTarget,
				t.trafficTargets,
				t.failoverServices,
				sourceMeshInstallation.GetCluster(),
			),
		},
	}

	// Initialize Istio TLS mode with default declared in Settings
	istioTlsMode, err := tls.MapIstioTlsMode(mtlsDefault.GetIstio().GetTlsMode())
	if err != nil {
		return nil, err
	}
	destinationRule.Spec.TrafficPolicy.Tls = &networkingv1alpha3spec.ClientTLSSettings{
		Mode: istioTlsMode,
	}

	return destinationRule, nil
}

// Return errors for each user-supplied VirtualService that applies to the same hostname as the translated VirtualService
func conflictsWithUserDestinationRule(
	userDestinationRules v1alpha3sets.DestinationRuleSet,
	translatedDestinationRule *networkingv1alpha3.DestinationRule,
) []error {
	// For each user DR, check whether any hosts match any hosts from translated DR
	var errs []error

	// destination rules from RemoteSnapshot only contain non-translated objects
	userDestinationRules.List(func(dr *networkingv1alpha3.DestinationRule) bool {
		// check if common hostnames exist
		commonHostname := utils.CommonHostnames([]string{dr.Spec.Host}, []string{translatedDestinationRule.Spec.Host})
		if len(commonHostname) > 0 {
			errs = append(
				errs,
				eris.Errorf("Unable to translate AppliedTrafficPolicies to DestinationRule, applies to host %s that is already configured by the existing DestinationRule %s", commonHostname[0], sets.Key(dr)),
			)
		}
		return false
	})

	return errs
}
