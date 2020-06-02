package k8s

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	k8s_core "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	k8s_core_controller "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	container_runtime "github.com/solo-io/service-mesh-hub/pkg/container-runtime"
	"github.com/solo-io/service-mesh-hub/pkg/kube"
	"github.com/solo-io/service-mesh-hub/pkg/kube/metadata"
	"github.com/solo-io/service-mesh-hub/pkg/kube/selection"
	k8s_tenancy "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/cluster-tenancy/k8s"
	k8s_core_types "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewMeshWorkloadFinder(
	ctx context.Context,
	clusterName string,
	localMeshClient zephyr_discovery.MeshClient,
	localMeshWorkloadClient zephyr_discovery.MeshWorkloadClient,
	meshWorkloadScannerImplementations MeshWorkloadScannerImplementations,
	podClient k8s_core.PodClient,
) MeshWorkloadFinder {

	return &meshWorkloadFinder{
		ctx:                                ctx,
		clusterName:                        clusterName,
		meshWorkloadScannerImplementations: meshWorkloadScannerImplementations,
		localMeshClient:                    localMeshClient,
		localMeshWorkloadClient:            localMeshWorkloadClient,
		podClient:                          podClient,
	}
}

type meshWorkloadFinder struct {
	clusterName                        string
	ctx                                context.Context
	meshWorkloadScannerImplementations MeshWorkloadScannerImplementations
	localMeshClient                    zephyr_discovery.MeshClient
	localMeshWorkloadClient            zephyr_discovery.MeshWorkloadClient
	podClient                          k8s_core.PodClient
}

func (m *meshWorkloadFinder) StartDiscovery(
	podEventWatcher k8s_core_controller.PodEventWatcher,
	meshEventWatcher zephyr_discovery_controller.MeshEventWatcher,
) error {
	// ensure the existing state in the cluster is accurate before starting to handle events
	err := m.reconcile()
	if err != nil {
		return err
	}

	err = podEventWatcher.AddEventHandler(
		m.ctx,
		&k8s_core_controller.PodEventHandlerFuncs{
			OnCreate: func(obj *k8s_core_types.Pod) error {
				logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.CreateEvent, obj)
				logger.Debugf("Handling create for pod %s.%s", obj.GetName(), obj.GetNamespace())
				err = m.reconcile()
				if err != nil {
					logger.Errorf("%+v", err)
				}
				return err
			},
			OnUpdate: func(old, new *k8s_core_types.Pod) error {
				logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.UpdateEvent, new)
				logger.Debugf("Handling update for pod %s.%s", new.GetName(), new.GetNamespace())
				err = m.reconcile()
				if err != nil {
					logger.Errorf("%+v", err)
				}
				return err
			},
			OnDelete: func(obj *k8s_core_types.Pod) error {
				logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.DeleteEvent, obj)
				logger.Debugf("Handling delete for pod %s.%s", obj.GetName(), obj.GetNamespace())
				err = m.reconcile()
				if err != nil {
					logger.Errorf("%+v", err)
				}
				return err
			},
		},
	)
	if err != nil {
		return err
	}
	err = meshEventWatcher.AddEventHandler(
		m.ctx,
		&zephyr_discovery_controller.MeshEventHandlerFuncs{
			OnCreate: func(obj *zephyr_discovery.Mesh) error {
				logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.CreateEvent, obj)
				logger.Debugf("mesh create event for %s.%s", obj.GetName(), obj.GetNamespace())
				err = m.reconcile()
				if err != nil {
					logger.Errorf("%+v", err)
				}
				return err
			},
			OnUpdate: func(old, new *zephyr_discovery.Mesh) error {
				logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.UpdateEvent, new)
				logger.Debugf("mesh update event for %s.%s", new.GetName(), new.GetNamespace())
				err = m.reconcile()
				if err != nil {
					logger.Errorf("%+v", err)
				}
				return err
			},
			OnDelete: func(obj *zephyr_discovery.Mesh) error {
				logger := container_runtime.BuildEventLogger(m.ctx, container_runtime.DeleteEvent, obj)
				logger.Debugf("mesh delete event for %s.%s", obj.GetName(), obj.GetNamespace())
				err = m.reconcile()
				if err != nil {
					logger.Errorf("%+v", err)
				}
				return err
			},
		},
	)
	return err
}

func (m *meshWorkloadFinder) reconcile() error {
	discoveredMeshTypes, err := m.getDiscoveredMeshTypes(m.ctx)
	if err != nil {
		return err
	}
	existingWorkloadsByName, existingWorkloadNames, err := m.getExistingWorkloads()
	if err != nil {
		return err
	}
	discoveredWorkloads, discoveredWorkloadNames, err := m.discoverAllWorkloads(discoveredMeshTypes)
	if err != nil {
		return err
	}
	// For each workload that is discovered, create if it doesn't exist or update it if the spec has changed.
	for _, discoveredWorkload := range discoveredWorkloads {
		existingWorkload, ok := existingWorkloadsByName[discoveredWorkload.GetName()]
		if !ok || !existingWorkload.Spec.Equal(discoveredWorkload.Spec) {
			err = m.localMeshWorkloadClient.UpsertMeshWorkloadSpec(m.ctx, discoveredWorkload)
			if err != nil {
				return err
			}
		}
	}
	// Delete MeshWorkloads that no longer exist.
	for _, existingWorkloadName := range existingWorkloadNames.Difference(discoveredWorkloadNames).List() {
		existingWorkload, ok := existingWorkloadsByName[existingWorkloadName]
		if !ok {
			continue
		}
		err = m.localMeshWorkloadClient.DeleteMeshWorkload(m.ctx, selection.ObjectMetaToObjectKey(existingWorkload.ObjectMeta))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *meshWorkloadFinder) getExistingWorkloads() (map[string]*zephyr_discovery.MeshWorkload, sets.String, error) {
	inThisCluster := client.MatchingLabels{
		kube.COMPUTE_TARGET: m.clusterName,
	}
	meshWorkloadList, err := m.localMeshWorkloadClient.ListMeshWorkload(m.ctx, inThisCluster)
	if err != nil {
		return nil, nil, err
	}
	workloadsByName := make(map[string]*zephyr_discovery.MeshWorkload)
	workloadNames := sets.NewString()
	for _, meshWorkload := range meshWorkloadList.Items {
		meshWorkload := meshWorkload
		workloadsByName[meshWorkload.GetName()] = &meshWorkload
		workloadNames.Insert(meshWorkload.GetName())
	}
	return workloadsByName, workloadNames, nil
}

func (m *meshWorkloadFinder) discoverAllWorkloads(discoveredMeshTypes sets.Int32) ([]*zephyr_discovery.MeshWorkload, sets.String, error) {
	podList, err := m.podClient.ListPod(m.ctx)
	if err != nil {
		return nil, nil, err
	}
	var meshWorkloads []*zephyr_discovery.MeshWorkload
	meshWorkloadNames := sets.NewString()
	for _, pod := range podList.Items {
		pod := pod
		discoveredWorkload, err := m.discoverMeshWorkload(&pod, discoveredMeshTypes)
		if err != nil {
			contextutils.LoggerFrom(m.ctx).Warnf("Error scanning pod %s.%s: %+v", pod.GetName(), pod.GetNamespace(), err)
			continue
		} else if discoveredWorkload == nil {
			continue
		}
		meshWorkloads = append(meshWorkloads, discoveredWorkload)
		meshWorkloadNames.Insert(discoveredWorkload.GetName())
	}
	return meshWorkloads, meshWorkloadNames, nil
}

func (m *meshWorkloadFinder) getDiscoveredMeshTypes(ctx context.Context) (sets.Int32, error) {
	meshList, err := m.localMeshClient.ListMesh(ctx)
	if err != nil {
		return nil, err
	}
	discoveredMeshTypes := sets.Int32{}
	for _, mesh := range meshList.Items {
		mesh := mesh
		// ensure we are only watching for meshes discovered on this same cluster
		// otherwise we can hit a race where:
		//  - Istio is discovered on cluster A
		//  - that mesh is recorded here
		//  - we start discovering workloads on cluster B using the Istio mesh workload discovery
		//  - but we haven't yet discovered Istio on this cluster
		if !k8s_tenancy.ClusterHostsMesh(m.clusterName, &mesh) {
			continue
		}
		meshType, err := metadata.MeshToMeshType(&mesh)
		if err != nil {
			return nil, err
		}
		discoveredMeshTypes.Insert(int32(meshType))
	}
	return discoveredMeshTypes, nil
}

func (m *meshWorkloadFinder) discoverMeshWorkload(pod *k8s_core_types.Pod, discoveredMeshTypes sets.Int32) (*zephyr_discovery.MeshWorkload, error) {
	logger := contextutils.LoggerFrom(m.ctx)
	var discoveredMeshWorkload *zephyr_discovery.MeshWorkload
	var err error

	for _, discoveredMeshType := range discoveredMeshTypes.List() {
		meshWorkloadScanner, ok := m.meshWorkloadScannerImplementations[zephyr_core_types.MeshType(discoveredMeshType)]
		if !ok {
			logger.Warnf("No MeshWorkloadScanner found for mesh type: %s", zephyr_core_types.MeshType(discoveredMeshType).String())
			continue
		}
		discoveredMeshWorkload, err = meshWorkloadScanner.ScanPod(m.ctx, pod, m.clusterName)
		if err != nil {
			return nil, err
		}
		if discoveredMeshWorkload != nil {
			break
		}
	}
	// the mesh workload needs to have our standard discovery labels attached to it, like cluster name, etc
	if discoveredMeshWorkload != nil {
		m.attachGeneralDiscoveryLabels(discoveredMeshWorkload)
	}
	return discoveredMeshWorkload, nil
}

func (m *meshWorkloadFinder) attachGeneralDiscoveryLabels(meshWorkload *zephyr_discovery.MeshWorkload) {
	if meshWorkload.Labels == nil {
		meshWorkload.Labels = map[string]string{}
	}
	meshWorkload.Labels[kube.DISCOVERED_BY] = kube.MESH_WORKLOAD_DISCOVERY
	meshWorkload.Labels[kube.COMPUTE_TARGET] = m.clusterName
	meshWorkload.Labels[kube.KUBE_CONTROLLER_NAME] = meshWorkload.Spec.GetKubeController().GetKubeControllerRef().GetName()
	meshWorkload.Labels[kube.KUBE_CONTROLLER_NAMESPACE] = meshWorkload.Spec.GetKubeController().GetKubeControllerRef().GetNamespace()
}
