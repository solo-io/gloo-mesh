package event_watcher_factories

import (
	discovery_controllers "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/mesh-platform/k8s"
)

func NewMeshEventWatcherFactory() MeshEventWatcherFactory {
	return &meshEventWatcherFactory{}
}

type meshEventWatcherFactory struct{}

func (d *meshEventWatcherFactory) Build(mgr mc_manager.AsyncManager, clusterName string) discovery_controllers.MeshEventWatcher {
	return discovery_controllers.NewMeshEventWatcher(clusterName, mgr.Manager())
}
