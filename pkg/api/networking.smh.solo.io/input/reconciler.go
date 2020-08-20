// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconciler.go -destination mocks/reconciler.go

// The Input Reconciler calls a simple func() error whenever a
// storage event is received for any of:
// * TrafficTargets
// * MeshWorkloads
// * Meshes
// * TrafficPolicies
// * AccessPolicies
// * VirtualMeshes
// * FailoverServices
// * Secrets
// * KubernetesClusters
// for a given cluster or set of clusters.
//
// Input Reconcilers can be be constructed from either a single Manager (watch events in a single cluster)
// or a ClusterWatcher (watch events in multiple clusters).
package input

import (
	"context"
	"time"

	"github.com/solo-io/skv2/contrib/pkg/input"
	sk_core_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	discovery_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	discovery_smh_solo_io_v1alpha2_controllers "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/controller"

	networking_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	networking_smh_solo_io_v1alpha2_controllers "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/controller"

	v1_controllers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	v1 "k8s.io/api/core/v1"

	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	multicluster_solo_io_v1alpha1_controllers "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/controller"
)

// the multiClusterReconciler reconciles events for input resources across clusters
type multiClusterReconciler interface {
	discovery_smh_solo_io_v1alpha2_controllers.MulticlusterTrafficTargetReconciler
	discovery_smh_solo_io_v1alpha2_controllers.MulticlusterMeshWorkloadReconciler
	discovery_smh_solo_io_v1alpha2_controllers.MulticlusterMeshReconciler

	networking_smh_solo_io_v1alpha2_controllers.MulticlusterTrafficPolicyReconciler
	networking_smh_solo_io_v1alpha2_controllers.MulticlusterAccessPolicyReconciler
	networking_smh_solo_io_v1alpha2_controllers.MulticlusterVirtualMeshReconciler
	networking_smh_solo_io_v1alpha2_controllers.MulticlusterFailoverServiceReconciler

	v1_controllers.MulticlusterSecretReconciler

	multicluster_solo_io_v1alpha1_controllers.MulticlusterKubernetesClusterReconciler
}

var _ multiClusterReconciler = &multiClusterReconcilerImpl{}

type multiClusterReconcilerImpl struct {
	base input.MultiClusterReconciler
}

// register the reconcile func with the cluster watcher
// the reconcileInterval, if greater than 0, will limit the number of reconciles
// to one per interval.
func RegisterMultiClusterReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	reconcileFunc input.MultiClusterReconcileFunc,
	reconcileInterval time.Duration,
	predicates ...predicate.Predicate,
) {

	base := input.NewMultiClusterReconcilerImpl(
		ctx,
		reconcileFunc,
		reconcileInterval,
	)

	r := &multiClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	discovery_smh_solo_io_v1alpha2_controllers.NewMulticlusterTrafficTargetReconcileLoop("TrafficTarget", clusters).AddMulticlusterTrafficTargetReconciler(ctx, r, predicates...)
	discovery_smh_solo_io_v1alpha2_controllers.NewMulticlusterMeshWorkloadReconcileLoop("MeshWorkload", clusters).AddMulticlusterMeshWorkloadReconciler(ctx, r, predicates...)
	discovery_smh_solo_io_v1alpha2_controllers.NewMulticlusterMeshReconcileLoop("Mesh", clusters).AddMulticlusterMeshReconciler(ctx, r, predicates...)

	networking_smh_solo_io_v1alpha2_controllers.NewMulticlusterTrafficPolicyReconcileLoop("TrafficPolicy", clusters).AddMulticlusterTrafficPolicyReconciler(ctx, r, predicates...)
	networking_smh_solo_io_v1alpha2_controllers.NewMulticlusterAccessPolicyReconcileLoop("AccessPolicy", clusters).AddMulticlusterAccessPolicyReconciler(ctx, r, predicates...)
	networking_smh_solo_io_v1alpha2_controllers.NewMulticlusterVirtualMeshReconcileLoop("VirtualMesh", clusters).AddMulticlusterVirtualMeshReconciler(ctx, r, predicates...)
	networking_smh_solo_io_v1alpha2_controllers.NewMulticlusterFailoverServiceReconcileLoop("FailoverService", clusters).AddMulticlusterFailoverServiceReconciler(ctx, r, predicates...)

	v1_controllers.NewMulticlusterSecretReconcileLoop("Secret", clusters).AddMulticlusterSecretReconciler(ctx, r, predicates...)

	multicluster_solo_io_v1alpha1_controllers.NewMulticlusterKubernetesClusterReconcileLoop("KubernetesCluster", clusters).AddMulticlusterKubernetesClusterReconciler(ctx, r, predicates...)

}

func (r *multiClusterReconcilerImpl) ReconcileTrafficTarget(clusterName string, obj *discovery_smh_solo_io_v1alpha2.TrafficTarget) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileTrafficTargetDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileMeshWorkload(clusterName string, obj *discovery_smh_solo_io_v1alpha2.MeshWorkload) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshWorkloadDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileMesh(clusterName string, obj *discovery_smh_solo_io_v1alpha2.Mesh) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileMeshDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileTrafficPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha2.TrafficPolicy) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileTrafficPolicyDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileAccessPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha2.AccessPolicy) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileAccessPolicyDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileVirtualMesh(clusterName string, obj *networking_smh_solo_io_v1alpha2.VirtualMesh) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileVirtualMeshDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileFailoverService(clusterName string, obj *networking_smh_solo_io_v1alpha2.FailoverService) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileFailoverServiceDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileSecret(clusterName string, obj *v1.Secret) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileSecretDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

func (r *multiClusterReconcilerImpl) ReconcileKubernetesCluster(clusterName string, obj *multicluster_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileKubernetesClusterDeletion(clusterName string, obj reconcile.Request) error {
	ref := &sk_core_v1.ClusterObjectRef{
		Name:        obj.Name,
		Namespace:   obj.Namespace,
		ClusterName: clusterName,
	}
	_, err := r.base.ReconcileClusterGeneric(ref)
	return err
}

// the singleClusterReconciler reconciles events for input resources across clusters
type singleClusterReconciler interface {
	discovery_smh_solo_io_v1alpha2_controllers.TrafficTargetReconciler
	discovery_smh_solo_io_v1alpha2_controllers.MeshWorkloadReconciler
	discovery_smh_solo_io_v1alpha2_controllers.MeshReconciler

	networking_smh_solo_io_v1alpha2_controllers.TrafficPolicyReconciler
	networking_smh_solo_io_v1alpha2_controllers.AccessPolicyReconciler
	networking_smh_solo_io_v1alpha2_controllers.VirtualMeshReconciler
	networking_smh_solo_io_v1alpha2_controllers.FailoverServiceReconciler

	v1_controllers.SecretReconciler

	multicluster_solo_io_v1alpha1_controllers.KubernetesClusterReconciler
}

var _ singleClusterReconciler = &singleClusterReconcilerImpl{}

type singleClusterReconcilerImpl struct {
	base input.SingleClusterReconciler
}

// register the reconcile func with the manager
// the reconcileInterval, if greater than 0, will limit the number of reconciles
// to one per interval.
func RegisterSingleClusterReconciler(
	ctx context.Context,
	mgr manager.Manager,
	reconcileFunc input.SingleClusterReconcileFunc,
	reconcileInterval time.Duration,
	predicates ...predicate.Predicate,
) error {

	base := input.NewSingleClusterReconciler(
		ctx,
		reconcileFunc,
		reconcileInterval,
	)

	r := &singleClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	if err := discovery_smh_solo_io_v1alpha2_controllers.NewTrafficTargetReconcileLoop("TrafficTarget", mgr, reconcile.Options{}).RunTrafficTargetReconciler(ctx, r, predicates...); err != nil {
		return err
	}
	if err := discovery_smh_solo_io_v1alpha2_controllers.NewMeshWorkloadReconcileLoop("MeshWorkload", mgr, reconcile.Options{}).RunMeshWorkloadReconciler(ctx, r, predicates...); err != nil {
		return err
	}
	if err := discovery_smh_solo_io_v1alpha2_controllers.NewMeshReconcileLoop("Mesh", mgr, reconcile.Options{}).RunMeshReconciler(ctx, r, predicates...); err != nil {
		return err
	}

	if err := networking_smh_solo_io_v1alpha2_controllers.NewTrafficPolicyReconcileLoop("TrafficPolicy", mgr, reconcile.Options{}).RunTrafficPolicyReconciler(ctx, r, predicates...); err != nil {
		return err
	}
	if err := networking_smh_solo_io_v1alpha2_controllers.NewAccessPolicyReconcileLoop("AccessPolicy", mgr, reconcile.Options{}).RunAccessPolicyReconciler(ctx, r, predicates...); err != nil {
		return err
	}
	if err := networking_smh_solo_io_v1alpha2_controllers.NewVirtualMeshReconcileLoop("VirtualMesh", mgr, reconcile.Options{}).RunVirtualMeshReconciler(ctx, r, predicates...); err != nil {
		return err
	}
	if err := networking_smh_solo_io_v1alpha2_controllers.NewFailoverServiceReconcileLoop("FailoverService", mgr, reconcile.Options{}).RunFailoverServiceReconciler(ctx, r, predicates...); err != nil {
		return err
	}

	if err := v1_controllers.NewSecretReconcileLoop("Secret", mgr, reconcile.Options{}).RunSecretReconciler(ctx, r, predicates...); err != nil {
		return err
	}

	if err := multicluster_solo_io_v1alpha1_controllers.NewKubernetesClusterReconcileLoop("KubernetesCluster", mgr, reconcile.Options{}).RunKubernetesClusterReconciler(ctx, r, predicates...); err != nil {
		return err
	}

	return nil
}

func (r *singleClusterReconcilerImpl) ReconcileTrafficTarget(obj *discovery_smh_solo_io_v1alpha2.TrafficTarget) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileTrafficTargetDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileMeshWorkload(obj *discovery_smh_solo_io_v1alpha2.MeshWorkload) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileMeshWorkloadDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileMesh(obj *discovery_smh_solo_io_v1alpha2.Mesh) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileMeshDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileTrafficPolicy(obj *networking_smh_solo_io_v1alpha2.TrafficPolicy) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileTrafficPolicyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileAccessPolicy(obj *networking_smh_solo_io_v1alpha2.AccessPolicy) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileAccessPolicyDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileVirtualMesh(obj *networking_smh_solo_io_v1alpha2.VirtualMesh) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileVirtualMeshDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileFailoverService(obj *networking_smh_solo_io_v1alpha2.FailoverService) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileFailoverServiceDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileSecret(obj *v1.Secret) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileSecretDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}

func (r *singleClusterReconcilerImpl) ReconcileKubernetesCluster(obj *multicluster_solo_io_v1alpha1.KubernetesCluster) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileKubernetesClusterDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}
