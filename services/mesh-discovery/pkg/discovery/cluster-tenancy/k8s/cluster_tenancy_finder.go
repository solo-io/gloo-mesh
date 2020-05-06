package k8s_tenancy

import (
	"context"

	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	k8s_core "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	k8s_core_controller "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	"github.com/solo-io/service-mesh-hub/pkg/clients"
	"github.com/solo-io/service-mesh-hub/pkg/logging"
	k8s_core_types "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type clusterTenancyFinder struct {
	clusterName       string
	tenancyRegistrars []ClusterTenancyRegistrar
	podClient         k8s_core.PodClient
	localMeshClient   zephyr_discovery.MeshClient
}

func NewClusterTenancyFinder(
	clusterName string,
	tenancyRegistrars []ClusterTenancyRegistrar,
	podClient k8s_core.PodClient,
	localMeshClient zephyr_discovery.MeshClient,
) ClusterTenancyRegistrarLoop {
	return &clusterTenancyFinder{
		clusterName:       clusterName,
		tenancyRegistrars: tenancyRegistrars,
		podClient:         podClient,
		localMeshClient:   localMeshClient,
	}
}

func (c *clusterTenancyFinder) StartRegistration(
	ctx context.Context,
	podEventWatcher k8s_core_controller.PodEventWatcher,
	meshEventWatcher zephyr_discovery_controller.MeshEventWatcher,
) (err error) {
	if err = podEventWatcher.AddEventHandler(ctx, &k8s_core_controller.PodEventHandlerFuncs{
		OnCreate: func(pod *k8s_core_types.Pod) error {
			logger := logging.BuildEventLogger(ctx, logging.CreateEvent, pod)
			logger.Debugf("Handling for %s.%s", pod.GetName(), pod.GetNamespace())
			err := c.reconcile(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnUpdate: func(_, pod *k8s_core_types.Pod) error {
			logger := logging.BuildEventLogger(ctx, logging.UpdateEvent, pod)
			logger.Debugf("Handling for %s.%s", pod.GetName(), pod.GetNamespace())
			err := c.reconcile(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnDelete: func(pod *k8s_core_types.Pod) error {
			logger := logging.BuildEventLogger(ctx, logging.DeleteEvent, pod)
			logger.Debugf("Handling for %s.%s", pod.GetName(), pod.GetNamespace())
			err := c.reconcile(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
	}); err != nil {
		return err
	}
	return meshEventWatcher.AddEventHandler(ctx, &zephyr_discovery_controller.MeshEventHandlerFuncs{
		OnCreate: func(mesh *zephyr_discovery.Mesh) error {
			logger := logging.BuildEventLogger(ctx, logging.CreateEvent, mesh)
			logger.Debugf("Handling for %s.%s", mesh.GetName(), mesh.GetNamespace())
			err := c.reconcile(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
		OnUpdate: func(_, mesh *zephyr_discovery.Mesh) error {
			logger := logging.BuildEventLogger(ctx, logging.UpdateEvent, mesh)
			logger.Debugf("Handling for %s.%s", mesh.GetName(), mesh.GetNamespace())
			err := c.reconcile(ctx)
			if err != nil {
				logger.Errorf("%+v", err)
			}
			return nil
		},
	})
}

/*
1. List all pods and aggregate all discovered Meshes for this cluster.
2. List all meshes and for each:
   If cluster hosts mesh and cluster not registered on mesh:
      update mesh with cluster
   Else if cluster does not host mesh and cluster registered on mesh:
      update Mesh by removing cluster
*/
func (c *clusterTenancyFinder) reconcile(ctx context.Context) error {
	podList, err := c.podClient.ListPod(ctx)
	if err != nil {
		return err
	}
	meshesOnCluster := make(map[ClusterTenancyRegistrar][]client.ObjectKey)
	for _, pod := range podList.Items {
		pod := pod
		registrar, meshObjKey, err := c.meshFromPod(ctx, &pod)
		if err != nil {
			return err
		}
		if meshObjKey != nil {
			meshObjKeys, ok := meshesOnCluster[registrar]
			if !ok {
				meshesOnCluster[registrar] = []client.ObjectKey{*meshObjKey}
			} else if !containsMesh(*meshObjKey, meshObjKeys) {
				meshesOnCluster[registrar] = append(meshObjKeys, *meshObjKey)
			}
		}
	}
	meshList, err := c.localMeshClient.ListMesh(ctx)
	if err != nil {
		return err
	}
	for _, mesh := range meshList.Items {
		mesh := mesh
		err := c.reconcileClusterTenancyForMesh(ctx, meshesOnCluster, &mesh)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *clusterTenancyFinder) meshFromPod(
	ctx context.Context,
	pod *k8s_core_types.Pod,
) (ClusterTenancyRegistrar, *client.ObjectKey, error) {
	for _, registrar := range c.tenancyRegistrars {
		mesh, err := registrar.MeshFromSidecar(ctx, pod)
		if err != nil {
			return nil, nil, err
		}
		if mesh != nil {
			key := clients.ObjectMetaToObjectKey(mesh.ObjectMeta)
			return registrar, &key, nil
		}
	}
	return nil, nil, nil
}

func (c *clusterTenancyFinder) reconcileClusterTenancyForMesh(
	ctx context.Context,
	meshesOnClusterByRegistrar map[ClusterTenancyRegistrar][]client.ObjectKey,
	mesh *zephyr_discovery.Mesh,
) error {
	meshObjKey := clients.ObjectMetaToObjectKey(mesh.ObjectMeta)
	for registrar, meshesOnCluster := range meshesOnClusterByRegistrar {
		clusterHostsMesh := registrar.ClusterHostsMesh(c.clusterName, mesh)
		if clusterHostsMesh && !containsMesh(meshObjKey, meshesOnCluster) {
			return registrar.DeregisterMesh(ctx, c.clusterName, mesh)
		} else if !clusterHostsMesh && containsMesh(meshObjKey, meshesOnCluster) {
			return registrar.RegisterMesh(ctx, c.clusterName, mesh)
		}
	}
	return nil
}

func containsMesh(mesh client.ObjectKey, meshObjKeys []client.ObjectKey) bool {
	for _, meshObjKey := range meshObjKeys {
		if mesh == meshObjKey {
			return true
		}
	}
	return false
}
