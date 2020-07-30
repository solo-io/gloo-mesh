// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./reconciler.go -destination mocks/reconciler.go

// The Input Reconciler calls a simple func() error whenever a
// storage event is received for any of:
// * CertificateRequests
// for a given cluster or set of clusters.
//
// Input Reconcilers can be be constructed from either a single Manager (watch events in a single cluster)
// or a ClusterWatcher (watch events in multiple clusters).
package input

import (
	"context"

	"github.com/solo-io/skv2/contrib/pkg/input"
	sk_core_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/multicluster"
	"github.com/solo-io/skv2/pkg/reconcile"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	certificates_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2"
	certificates_smh_solo_io_v1alpha2_controllers "github.com/solo-io/service-mesh-hub/pkg/api/certificates.smh.solo.io/v1alpha2/controller"
)

// the multiClusterReconciler reconciles events for input resources across clusters
type multiClusterReconciler interface {
	certificates_smh_solo_io_v1alpha2_controllers.MulticlusterCertificateRequestReconciler
}

var _ multiClusterReconciler = &multiClusterReconcilerImpl{}

type multiClusterReconcilerImpl struct {
	base input.MultiClusterReconciler
}

// register the reconcile func with the cluster watcher
func RegisterMultiClusterReconciler(
	ctx context.Context,
	clusters multicluster.ClusterWatcher,
	reconcileFunc input.MultiClusterReconcileFunc,
) {

	base := input.NewMultiClusterReconcilerImpl(
		ctx,
		reconcileFunc,
	)

	r := &multiClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	certificates_smh_solo_io_v1alpha2_controllers.NewMulticlusterCertificateRequestReconcileLoop("CertificateRequest", clusters).AddMulticlusterCertificateRequestReconciler(ctx, r)

}

func (r *multiClusterReconcilerImpl) ReconcileCertificateRequest(clusterName string, obj *certificates_smh_solo_io_v1alpha2.CertificateRequest) (reconcile.Result, error) {
	obj.ClusterName = clusterName
	return r.base.ReconcileClusterGeneric(obj)
}

func (r *multiClusterReconcilerImpl) ReconcileCertificateRequestDeletion(clusterName string, obj reconcile.Request) error {
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
	certificates_smh_solo_io_v1alpha2_controllers.CertificateRequestReconciler
}

var _ singleClusterReconciler = &singleClusterReconcilerImpl{}

type singleClusterReconcilerImpl struct {
	base input.SingleClusterReconciler
}

// register the reconcile func with the manager
func RegisterSingleClusterReconciler(
	ctx context.Context,
	mgr manager.Manager,
	reconcileFunc input.SingleClusterReconcileFunc,
) error {

	base := input.NewSingleClusterReconciler(
		ctx,
		reconcileFunc,
	)

	r := &singleClusterReconcilerImpl{
		base: base,
	}

	// initialize reconcile loops

	if err := certificates_smh_solo_io_v1alpha2_controllers.NewCertificateRequestReconcileLoop("CertificateRequest", mgr, reconcile.Options{}).RunCertificateRequestReconciler(ctx, r); err != nil {
		return err
	}

	return nil
}

func (r *singleClusterReconcilerImpl) ReconcileCertificateRequest(obj *certificates_smh_solo_io_v1alpha2.CertificateRequest) (reconcile.Result, error) {
	return r.base.ReconcileGeneric(obj)
}

func (r *singleClusterReconcilerImpl) ReconcileCertificateRequestDeletion(obj reconcile.Request) error {
	ref := &sk_core_v1.ObjectRef{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}
	_, err := r.base.ReconcileGeneric(ref)
	return err
}
