package controllers

import (
	apps_controllers "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/apps/v1/controller"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/multicluster/manager"
)

func NewDeploymentEventWatcherFactory() DeploymentEventWatcherFactory {
	return &defaultDeploymentEventWatcherFactory{}
}

type defaultDeploymentEventWatcherFactory struct{}

func (d *defaultDeploymentEventWatcherFactory) Build(mgr mc_manager.AsyncManager, clusterName string) apps_controllers.DeploymentEventWatcher {
	// just directly return the generated autopilot implementation
	return apps_controllers.NewDeploymentEventWatcher(clusterName, mgr.Manager())
}
