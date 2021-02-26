package detector

import (
	"context"
	"strings"

	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discoveryv1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/utils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/utils/localityutils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/utils/workloadutils"
	"github.com/solo-io/go-utils/contextutils"
	sets2 "github.com/solo-io/skv2/contrib/pkg/sets"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/pointer"
)

const (
	// TODO: allow for specifying specific meshes.
	// Currently this annotation assumes that there is only one mesh per cluster, and therefore the corresponding
	// Destination will be associated with that mesh.
	DiscoveryMeshAnnotation = "discovery.mesh.gloo.solo.io/enabled"
)

var (
	skippedLabels = sets.NewString(
		"pod-template-hash",
		"service.istio.io/canonical-revision",
	)
)

// the DestinationDetector detects Destinations from services
// whose backing pods are injected with a Mesh sidecar.
// If no Mesh is detected, nil is returned
type DestinationDetector interface {
	DetectDestination(
		ctx context.Context,
		service *corev1.Service,
		pods corev1sets.PodSet,
		nodes corev1sets.NodeSet,
		workloads discoveryv1sets.WorkloadSet,
		meshes discoveryv1sets.MeshSet,
		endpoints corev1sets.EndpointsSet,
	) *v1.Destination
}

type destinationDetector struct{}

func NewDestinationDetector() DestinationDetector {
	return &destinationDetector{}
}

func (t *destinationDetector) DetectDestination(
	ctx context.Context,
	service *corev1.Service,
	pods corev1sets.PodSet,
	nodes corev1sets.NodeSet,
	workloads discoveryv1sets.WorkloadSet,
	meshes discoveryv1sets.MeshSet,
	endpoints corev1sets.EndpointsSet,
) *v1.Destination {

	kubeService := &v1.DestinationSpec_KubeService{
		Ref:                    ezkube.MakeClusterObjectRef(service),
		WorkloadSelectorLabels: service.Spec.Selector,
		Labels:                 service.Labels,
		Ports:                  convertPorts(service),
	}

	// add locality to the destination
	region, err := localityutils.GetServiceRegion(service, pods, nodes)
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnw("could not get region for destination", zap.Error(err))
	}
	kubeService.Region = region

	destination := &v1.Destination{
		ObjectMeta: utils.DiscoveredObjectMeta(service),
		Spec: v1.DestinationSpec{
			Type: &v1.DestinationSpec_KubeService_{
				KubeService: kubeService,
			},
		},
	}

	// If the service is not associated with a mesh, do not create a destination
	validDestination := addMeshForKubeService(
		ctx,
		destination,
		service,
		workloads,
		meshes,
		endpoints,
		nodes,
	)
	if !validDestination {
		return nil
	}

	return destination
}

func addMeshForKubeService(
	ctx context.Context,
	tt *v1.Destination,
	service *corev1.Service,
	meshWorkloads discoveryv1sets.WorkloadSet,
	meshes discoveryv1sets.MeshSet,
	endpoints corev1sets.EndpointsSet,
	nodes corev1sets.NodeSet,
) bool {

	var validMesh *skv2corev1.ObjectRef

	// TODO: support subsets from services which have been discovered via the annotation
	discoveryEnabled, ok := service.Annotations[DiscoveryMeshAnnotation]
	if ok && discoveryEnabled == "true" {

		// Search for mesh which exists on the same cluster as the annotated service
		for _, mesh := range meshes.List() {
			mesh := mesh
			switch typedMesh := mesh.Spec.GetType().(type) {
			case *v1.MeshSpec_Osm:
				if typedMesh.Osm.GetInstallation().GetCluster() == service.GetClusterName() {
					validMesh = ezkube.MakeObjectRef(mesh)
					break
				}
			case *v1.MeshSpec_Istio_:
				if typedMesh.Istio.GetInstallation().GetCluster() == service.GetClusterName() {
					validMesh = ezkube.MakeObjectRef(mesh)
					break
				}
			}
		}

		if validMesh == nil {
			contextutils.LoggerFrom(ctx).Errorf(
				"mesh could not be found for annotated service %s",
				sets2.TypedKey(service),
			)
		}
	}

	if validMesh != nil {
		tt.Spec.Mesh = validMesh
		return true
	}

	// if no mesh was found from the annotation, check the workloads
	backingWorkloads := workloadutils.FindBackingWorkloads(tt.Spec.GetKubeService(), meshWorkloads)
	// If there are no backing workloads, then we cannot find the associated mesh
	if len(backingWorkloads) == 0 {
		return false
	}
	handleWorkloadDiscoveredMesh(ctx, tt, backingWorkloads, endpoints, nodes)
	return true
}

func handleWorkloadDiscoveredMesh(
	ctx context.Context,
	tt *v1.Destination,
	backingWorkloads v1.WorkloadSlice,
	endpoints corev1sets.EndpointsSet,
	nodes corev1sets.NodeSet,
) {

	// all backing workloads should be in the same mesh
	tt.Spec.Mesh = backingWorkloads[0].Spec.Mesh

	// derive subsets from backing workloads
	tt.Spec.GetKubeService().Subsets = findSubsets(backingWorkloads)

	ep, err := endpoints.Find(tt.Spec.GetKubeService().GetRef())
	if err != nil {
		contextutils.LoggerFrom(ctx).Errorf(
			"endpoints could not be found for kube service %s",
			sets2.TypedKey(tt.Spec.GetKubeService().GetRef()),
		)
		return
	}

	// dervive endpoints from kubernetes endpoints, and backing workloads
	findEndpoints(ctx, backingWorkloads, ep, nodes, tt.Spec.GetKubeService())
}

func findEndpoints(
	ctx context.Context,
	backingWorkloads v1.WorkloadSlice,
	endpoint *corev1.Endpoints,
	nodes corev1sets.NodeSet,
	kubeService *v1.DestinationSpec_KubeService,
) {

	for _, epSub := range endpoint.Subsets {
		sub := &v1.DestinationSpec_KubeService_EndpointsSubset{}
		for _, addr := range epSub.Addresses {
			addr := addr
			ep := &v1.DestinationSpec_KubeService_EndpointsSubset_Endpoint{
				IpAddress: addr.IP,
			}

			if addr.NodeName != nil {
				subLocality, err := localityutils.GetSubLocality(kubeService.GetRef().GetClusterName(), *addr.NodeName, nodes)
				if err != nil {
					// Log the error but continue processing. We just won't be able to get a locality for this address
					contextutils.LoggerFrom(ctx).Warnw("could not get locality for address", "error", err)
				} else {
					ep.SubLocality = subLocality
				}
			} else {
				contextutils.LoggerFrom(ctx).Warnw("address does not have a node", "address", addr)
			}

			if addr.TargetRef != nil {
				for _, workload := range backingWorkloads {
					kubeWorkload := workload.Spec.GetKubernetes()
					if kubeWorkload == nil {
						continue
					}
					// Check if TargetRef points to a child of a backing workload to get the labels
					if addr.TargetRef.Namespace == kubeWorkload.GetController().GetNamespace() &&
						strings.HasPrefix(addr.TargetRef.Name, kubeWorkload.GetController().GetName()+"-") {
						ep.Labels = kubeWorkload.GetPodLabels()
						break
					}
				}
			} else {
				contextutils.LoggerFrom(ctx).Debugf(
					"skipping endpoint subset addr (%v) because targetRef is nil",
					addr,
				)
			}

			sub.Endpoints = append(sub.Endpoints, ep)

		}

		for _, port := range epSub.Ports {
			port := port
			svcPort := &v1.DestinationSpec_KubeService_KubeServicePort{
				Port:     uint32(port.Port),
				Name:     port.Name,
				Protocol: string(port.Protocol),
			}
			if port.AppProtocol != nil {
				svcPort.AppProtocol = *port.AppProtocol
			}
			sub.Ports = append(sub.Ports, svcPort)
		}

		// Only add this subset to the list if any IPs matched the workload in question
		if len(sub.GetEndpoints()) == 0 {
			contextutils.LoggerFrom(ctx).Debugf(
				"skipping endpoint address %v because no ip addresses were found",
				epSub,
			)
			continue
		}

		kubeService.EndpointSubsets = append(kubeService.EndpointSubsets, sub)
	}
}

// expects a list of just the workloads that back the service you're finding subsets for
func findSubsets(backingWorkloads v1.WorkloadSlice) map[string]*v1.DestinationSpec_KubeService_Subset {
	uniqueLabels := make(map[string]sets.String)
	for _, backingWorkload := range backingWorkloads {
		for key, val := range backingWorkload.Spec.GetKubernetes().GetPodLabels() {
			// skip known kubernetes values
			if skippedLabels.Has(key) {
				continue
			}
			existing, ok := uniqueLabels[key]
			if !ok {
				uniqueLabels[key] = sets.NewString(val)
			} else {
				existing.Insert(val)
			}
		}
	}
	/*
		Only select the keys with > 1 value
		The subsets worth noting will be sets of labels which share the same key, but have different values, such as:

			version:
				- v1
				- v2
	*/
	subsets := make(map[string]*v1.DestinationSpec_KubeService_Subset)
	for k, v := range uniqueLabels {
		if v.Len() > 1 {
			subsets[k] = &v1.DestinationSpec_KubeService_Subset{Values: v.List()}
		}
	}
	if len(subsets) == 0 {
		// important to return nil instead of empty map for asserting equality
		return nil
	}
	return subsets
}

func convertPorts(service *corev1.Service) (ports []*v1.DestinationSpec_KubeService_KubeServicePort) {
	for _, kubePort := range service.Spec.Ports {
		ports = append(ports, &v1.DestinationSpec_KubeService_KubeServicePort{
			Port:        uint32(kubePort.Port),
			Name:        kubePort.Name,
			Protocol:    string(kubePort.Protocol),
			AppProtocol: pointer.StringPtrDerefOr(kubePort.AppProtocol, ""),
		})
	}
	return ports
}
