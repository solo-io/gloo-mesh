package k8s

import (
	"context"
	"fmt"
	"strings"

	smh_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	smh_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller"
	smh_discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/types"
	k8s_core "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	k8s_core_controller "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	container_runtime "github.com/solo-io/service-mesh-hub/pkg/container-runtime"
	"github.com/solo-io/service-mesh-hub/pkg/kube"
	"github.com/solo-io/service-mesh-hub/pkg/kube/metadata"
	"github.com/solo-io/service-mesh-hub/pkg/kube/selection"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	DiscoveryLabels = func(meshType smh_core_types.MeshType, cluster, kubeServiceName, kubeServiceNamespace string) map[string]string {
		return map[string]string{
			kube.DISCOVERED_BY:          kube.MESH_WORKLOAD_DISCOVERY,
			kube.MESH_TYPE:              strings.ToLower(meshType.String()),
			kube.KUBE_SERVICE_NAME:      kubeServiceName,
			kube.KUBE_SERVICE_NAMESPACE: kubeServiceNamespace,
			kube.COMPUTE_TARGET:         cluster,
		}
	}

	skippedLabels = sets.NewString(
		"pod-template-hash",
		"service.istio.io/canonical-revision",
	)
)

func NewMeshServiceFinder(
	ctx context.Context,
	clusterName, writeNamespace string,
	serviceClient k8s_core.ServiceClient,
	meshServiceClient smh_discovery.MeshServiceClient,
	meshWorkloadClient smh_discovery.MeshWorkloadClient,
	meshClient smh_discovery.MeshClient,
) MeshServiceFinder {
	return &meshServiceFinder{
		ctx:                ctx,
		writeNamespace:     writeNamespace,
		clusterName:        clusterName,
		serviceClient:      serviceClient,
		meshServiceClient:  meshServiceClient,
		meshWorkloadClient: meshWorkloadClient,
		meshClient:         meshClient,
	}
}

func (m *meshServiceFinder) StartDiscovery(
	serviceEventWatcher k8s_core_controller.ServiceEventWatcher,
	meshWorkloadEventWatcher smh_discovery_controller.MeshWorkloadEventWatcher,
) error {
	err := m.reconcileMeshServices()
	if err != nil {
		return err
	}

	err = serviceEventWatcher.AddEventHandler(m.ctx, &k8s_core_controller.ServiceEventHandlerFuncs{
		OnCreate: func(obj *k8s_core_types.Service) error {
			logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.CreateEvent, obj)
			err := m.reconcileMeshServices()
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnUpdate: func(_, new *k8s_core_types.Service) error {
			logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.UpdateEvent, new)
			err := m.reconcileMeshServices()
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnDelete: func(obj *k8s_core_types.Service) error {
			logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.DeleteEvent, obj)
			err := m.reconcileMeshServices()
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
	})
	if err != nil {
		return err
	}
	return meshWorkloadEventWatcher.AddEventHandler(m.ctx, &smh_discovery_controller.MeshWorkloadEventHandlerFuncs{
		OnCreate: func(obj *smh_discovery.MeshWorkload) error {
			logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.CreateEvent, obj)
			err := m.reconcileMeshServices()
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnUpdate: func(_, new *smh_discovery.MeshWorkload) error {
			logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.UpdateEvent, new)
			err := m.reconcileMeshServices()
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnDelete: func(obj *smh_discovery.MeshWorkload) error {
			logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.DeleteEvent, obj)
			err := m.reconcileMeshServices()
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
	})
}

type meshServiceFinder struct {
	ctx                context.Context
	writeNamespace     string
	clusterName        string
	serviceClient      k8s_core.ServiceClient
	meshServiceClient  smh_discovery.MeshServiceClient
	meshWorkloadClient smh_discovery.MeshWorkloadClient
	meshClient         smh_discovery.MeshClient
}

func (m *meshServiceFinder) reconcileMeshServices() error {
	existingMeshServicesByName, existingMeshServiceNames, err := m.getExistingMeshServices()
	if err != nil {
		return err
	}
	discoveredMeshServices, discoveredMeshServiceNames, err := m.discoverMeshServices()
	if err != nil {
		return err
	}
	// For each service that is discovered, create if it doesn't exist or update it if the spec has changed.
	for _, discoveredMeshService := range discoveredMeshServices {
		existingMeshService, ok := existingMeshServicesByName[discoveredMeshService.GetName()]
		if !ok || !existingMeshService.Spec.Equal(discoveredMeshService.Spec) {
			err = m.meshServiceClient.UpsertMeshServiceSpec(m.ctx, discoveredMeshService)
			if err != nil {
				return err
			}
		}
	}
	// Delete MeshServices that no longer exist.
	for _, existingMeshServiceName := range existingMeshServiceNames.Difference(discoveredMeshServiceNames).List() {
		existingMeshService, ok := existingMeshServicesByName[existingMeshServiceName]
		if !ok {
			continue
		}
		err = m.meshServiceClient.DeleteMeshService(m.ctx, selection.ObjectMetaToObjectKey(existingMeshService.ObjectMeta))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *meshServiceFinder) getExistingMeshServices() (map[string]*smh_discovery.MeshService, sets.String, error) {
	meshServiceNames := sets.NewString()
	existingMeshServices, err := m.meshServiceClient.ListMeshService(m.ctx, client.MatchingLabels{
		kube.COMPUTE_TARGET: m.clusterName,
	})
	if err != nil {
		return nil, nil, err
	}
	meshServicesByName := make(map[string]*smh_discovery.MeshService)
	for _, existingMeshService := range existingMeshServices.Items {
		existingMeshService := existingMeshService
		meshServicesByName[existingMeshService.GetName()] = &existingMeshService
		meshServiceNames.Insert(existingMeshService.GetName())
	}
	return meshServicesByName, meshServiceNames, nil
}

func (m *meshServiceFinder) discoverMeshServices() ([]*smh_discovery.MeshService, sets.String, error) {
	discoveredMeshServiceNames := sets.NewString()
	services, err := m.serviceClient.ListService(m.ctx)
	if err != nil {
		return nil, nil, err
	}
	var discoveredMeshServices []*smh_discovery.MeshService
	for _, kubeService := range services.Items {
		kubeService := kubeService
		mesh, backingWorkloads, err := m.findMeshAndWorkloadsForService(&kubeService)
		if err != nil {
			return nil, nil, err
		}
		if len(backingWorkloads) == 0 {
			continue
		}
		discoveredMeshService, err := m.buildMeshService(&kubeService, mesh, m.findSubsets(backingWorkloads), m.clusterName)
		if err != nil {
			return nil, nil, err
		}
		discoveredMeshServices = append(discoveredMeshServices, discoveredMeshService)
		discoveredMeshServiceNames.Insert(discoveredMeshService.GetName())
	}
	return discoveredMeshServices, discoveredMeshServiceNames, nil
}

func (m *meshServiceFinder) findMeshAndWorkloadsForService(
	service *k8s_core_types.Service,
) (*smh_discovery.Mesh, []*smh_discovery.MeshWorkload, error) {
	// early optimization- bail out early if we know that this service can't select anything
	// otherwise we'll have to check all the mesh workloads
	if len(service.Spec.Selector) == 0 {
		return nil, nil, nil
	}
	meshWorkloads, err := m.meshWorkloadClient.ListMeshWorkload(m.ctx, client.MatchingLabels{
		kube.COMPUTE_TARGET: m.clusterName,
	})
	if err != nil {
		return nil, nil, err
	}
	var backingWorkloads []*smh_discovery.MeshWorkload
	var mesh *smh_discovery.Mesh
	for _, meshWorkloadIter := range meshWorkloads.Items {
		meshWorkload := meshWorkloadIter
		meshForWorkload, err := m.meshClient.GetMesh(m.ctx, selection.ResourceRefToObjectKey(meshWorkload.Spec.GetMesh()))
		if err != nil {
			return nil, nil, err
		}
		if m.isServiceBackedByWorkload(service, &meshWorkload) {
			mesh = meshForWorkload
			backingWorkloads = append(backingWorkloads, &meshWorkload)
		}
	}
	return mesh, backingWorkloads, nil
}

// expects a list of just the workloads that back the service you're finding subsets for
func (m *meshServiceFinder) findSubsets(backingWorkloads []*smh_discovery.MeshWorkload) map[string]*smh_discovery_types.MeshServiceSpec_Subset {

	uniqueLabels := make(map[string]sets.String)
	for _, backingWorkload := range backingWorkloads {
		for key, val := range backingWorkload.Spec.GetKubeController().GetLabels() {
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
	subsets := make(map[string]*smh_discovery_types.MeshServiceSpec_Subset)
	for k, v := range uniqueLabels {
		if v.Len() > 1 {
			subsets[k] = &smh_discovery_types.MeshServiceSpec_Subset{Values: v.List()}
		}
	}
	return subsets
}

func (m *meshServiceFinder) isServiceBackedByWorkload(
	service *k8s_core_types.Service,
	meshWorkload *smh_discovery.MeshWorkload,
) bool {
	workloadCluster := meshWorkload.Labels[kube.COMPUTE_TARGET]

	// If the meshworkload is not on the same cluster as the service, it can be skipped safely
	// The event handler accepts events from MeshWorkloads which may "match" the incoming service
	// but be on a different cluster, so it is important to check that here.
	if workloadCluster != m.clusterName {
		return false
	}

	// if either the service has no selector labels or the mesh workload's corresponding pod has no labels,
	// then this service cannot be backed by this mesh workload
	// the library call below returns true for either case, so we explicitly check for it here
	if len(service.Spec.Selector) == 0 || len(meshWorkload.Spec.GetKubeController().GetLabels()) == 0 {
		return false
	}

	// If service not in same namespace as workload, continue
	if service.GetNamespace() != meshWorkload.Spec.GetKubeController().GetKubeControllerRef().GetNamespace() {
		return false
	}

	return labels.AreLabelsInWhiteList(service.Spec.Selector, meshWorkload.Spec.GetKubeController().GetLabels())
}

func (m *meshServiceFinder) buildMeshService(
	service *k8s_core_types.Service,
	mesh *smh_discovery.Mesh,
	subsets map[string]*smh_discovery_types.MeshServiceSpec_Subset,
	clusterName string,
) (*smh_discovery.MeshService, error) {
	meshType, err := metadata.MeshToMeshType(mesh)
	if err != nil {
		return nil, err
	}

	return &smh_discovery.MeshService{
		ObjectMeta: k8s_meta_types.ObjectMeta{
			Name:      m.buildMeshServiceName(service, clusterName),
			Namespace: m.writeNamespace,
			Labels:    DiscoveryLabels(meshType, clusterName, service.GetName(), service.GetNamespace()),
		},
		Spec: smh_discovery_types.MeshServiceSpec{
			KubeService: &smh_discovery_types.MeshServiceSpec_KubeService{
				Ref: &smh_core_types.ResourceRef{
					Name:      service.GetName(),
					Namespace: service.GetNamespace(),
					Cluster:   clusterName,
				},
				WorkloadSelectorLabels: service.Spec.Selector,
				Labels:                 service.GetLabels(),
				Ports:                  m.convertPorts(service),
			},
			Mesh:    selection.ObjectMetaToResourceRef(mesh.ObjectMeta),
			Subsets: subsets,
		},
	}, nil
}

func (m *meshServiceFinder) convertPorts(service *k8s_core_types.Service) (ports []*smh_discovery_types.MeshServiceSpec_KubeService_KubeServicePort) {
	for _, kubePort := range service.Spec.Ports {
		ports = append(ports, &smh_discovery_types.MeshServiceSpec_KubeService_KubeServicePort{
			Port:     uint32(kubePort.Port),
			Name:     kubePort.Name,
			Protocol: string(kubePort.Protocol),
		})
	}
	return ports
}

func (m *meshServiceFinder) buildMeshServiceName(service *k8s_core_types.Service, clusterName string) string {
	return fmt.Sprintf("%s-%s-%s", service.GetName(), service.GetNamespace(), clusterName)
}
