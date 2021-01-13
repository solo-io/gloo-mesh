package mirror

import (
	"github.com/rotisserie/eris"
	discoveryv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	discoveryv1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2/sets"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/decorators"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/hostutils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/trafficpolicyutils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/traffictargetutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	networkingv1alpha3spec "istio.io/api/networking/v1alpha3"
)

const (
	decoratorName = "mirror"
)

func init() {
	decorators.Register(decoratorConstructor)
}

func decoratorConstructor(params decorators.Parameters) decorators.Decorator {
	return NewMirrorDecorator(params.ClusterDomains, params.Snapshot.TrafficTargets())
}

// handles setting Mirror on a VirtualService
type mirrorDecorator struct {
	clusterDomains hostutils.ClusterDomainRegistry
	trafficTargets discoveryv1alpha2sets.TrafficTargetSet
}

var _ decorators.TrafficPolicyVirtualServiceDecorator = &mirrorDecorator{}

func NewMirrorDecorator(
	clusterDomains hostutils.ClusterDomainRegistry,
	trafficTargets discoveryv1alpha2sets.TrafficTargetSet,
) *mirrorDecorator {
	return &mirrorDecorator{
		clusterDomains: clusterDomains,
		trafficTargets: trafficTargets,
	}
}

func (d *mirrorDecorator) DecoratorName() string {
	return decoratorName
}

func (d *mirrorDecorator) ApplyTrafficPolicyToVirtualService(
	appliedPolicy *discoveryv1alpha2.TrafficTargetStatus_AppliedTrafficPolicy,
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	sourceMeshInstallation *discoveryv1alpha2.MeshSpec_MeshInstallation,
	output *networkingv1alpha3spec.HTTPRoute,
	registerField decorators.RegisterField,
) error {
	mirror, percentage, err := d.translateMirror(trafficTarget, appliedPolicy.Spec, sourceMeshInstallation.GetCluster())
	if err != nil {
		return err
	}
	if mirror != nil {
		if err := registerField(&output.Mirror, mirror); err != nil {
			return err
		}
		output.Mirror = mirror
		output.MirrorPercentage = percentage
	}
	return nil
}

// If federatedClusterName is non-empty, it indicates translation for a federated VirtualService, so use it as the source cluster name.
func (d *mirrorDecorator) translateMirror(
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	trafficPolicy *v1alpha2.TrafficPolicySpec,
	sourceClusterName string,
) (*networkingv1alpha3spec.Destination, *networkingv1alpha3spec.Percent, error) {
	mirror := trafficPolicy.Mirror
	if mirror == nil {
		return nil, nil, nil
	}
	if mirror.DestinationType == nil {
		return nil, nil, eris.Errorf("must provide mirror destination")
	}

	var translatedMirror *networkingv1alpha3spec.Destination
	switch destinationType := mirror.DestinationType.(type) {
	case *v1alpha2.TrafficPolicySpec_Mirror_KubeService:
		var err error
		translatedMirror, err = d.makeKubeDestinationMirror(
			destinationType,
			mirror.Port,
			trafficTarget,
			sourceClusterName,
		)
		if err != nil {
			return nil, nil, err
		}
	}

	mirrorPercentage := &networkingv1alpha3spec.Percent{
		Value: mirror.GetPercentage(),
	}

	return translatedMirror, mirrorPercentage, nil
}

func (d *mirrorDecorator) makeKubeDestinationMirror(
	destination *v1alpha2.TrafficPolicySpec_Mirror_KubeService,
	port uint32,
	trafficTarget *discoveryv1alpha2.TrafficTarget,
	sourceClusterName string,
) (*networkingv1alpha3spec.Destination, error) {
	destinationRef := destination.KubeService
	mirrorService, err := traffictargetutils.FindTrafficTargetForKubeService(d.trafficTargets.List(), destinationRef)
	if err != nil {
		return nil, eris.Wrapf(err, "invalid mirror destination")
	}
	mirrorKubeService := mirrorService.Spec.GetKubeService()

	// TODO(ilackarms): support other types of TrafficTarget destinations, e.g. via ServiceEntries

	// An empty sourceClusterName indicates translation for VirtualService local to trafficTarget
	if sourceClusterName == "" {
		sourceClusterName = trafficTarget.Spec.GetKubeService().GetRef().GetClusterName()
	}

	destinationHostname := d.clusterDomains.GetDestinationFQDN(
		sourceClusterName,
		destinationRef,
	)

	translatedMirror := &networkingv1alpha3spec.Destination{
		Host: destinationHostname,
	}

	if port != 0 {
		if !trafficpolicyutils.ContainsPort(mirrorKubeService.Ports, port) {
			return nil, eris.Errorf("specified port %d does not exist for mirror destination service %v", port, sets.Key(mirrorKubeService.Ref))
		}
		translatedMirror.Port = &networkingv1alpha3spec.PortSelector{
			Number: port,
		}
	} else {
		// validate that traffic target only has one port
		if numPorts := len(mirrorKubeService.GetPorts()); numPorts > 1 {
			return nil, eris.Errorf("must provide port for mirror destination service %v with multiple ports (%v) defined", sets.Key(mirrorKubeService.GetRef()), numPorts)
		}
	}

	return translatedMirror, nil
}
