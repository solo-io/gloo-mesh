package reconcilers

import (
	apps_v1_controller "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/controller"
	core_v1_controller "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	smh_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller"
	"github.com/solo-io/skv2/pkg/reconcile"
	apps_v1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type discoveryReconcilers struct{}

type DiscoveryReconcilers interface {
	smh_discovery_controller.MeshWorkloadReconciler
	smh_discovery_controller.MeshReconciler

	apps_v1_controller.MulticlusterDeploymentReconciler
	core_v1_controller.MulticlusterPodReconciler
	core_v1_controller.MulticlusterServiceReconciler
}

func (d *discoveryReconcilers) ReconcileMeshWorkload(obj *smh_discovery.MeshWorkload) (reconcile.Result, error) {
	panic("implement me")
}

func (d *discoveryReconcilers) ReconcileMesh(obj *smh_discovery.Mesh) (reconcile.Result, error) {
	panic("implement me")
}

func (d *discoveryReconcilers) ReconcileDeployment(clusterName string, obj *apps_v1.Deployment) (reconcile.Result, error) {
	panic("implement me")
}

func (d *discoveryReconcilers) ReconcilePod(clusterName string, obj *v1.Pod) (reconcile.Result, error) {
	panic("implement me")
}

func (d *discoveryReconcilers) ReconcileService(clusterName string, obj *v1.Service) (reconcile.Result, error) {
	panic("implement me")
}
