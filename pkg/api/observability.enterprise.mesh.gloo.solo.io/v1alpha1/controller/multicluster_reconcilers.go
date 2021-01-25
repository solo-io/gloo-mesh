// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	observability_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/observability.enterprise.mesh.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the AccessLogRecord Resource across clusters.
// implemented by the user
type MulticlusterAccessLogRecordReconciler interface {
	ReconcileAccessLogRecord(clusterName string, obj *observability_enterprise_mesh_gloo_solo_io_v1alpha1.AccessLogRecord) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessLogRecord Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAccessLogRecordDeletionReconciler interface {
	ReconcileAccessLogRecordDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterAccessLogRecordReconcilerFuncs struct {
	OnReconcileAccessLogRecord         func(clusterName string, obj *observability_enterprise_mesh_gloo_solo_io_v1alpha1.AccessLogRecord) (reconcile.Result, error)
	OnReconcileAccessLogRecordDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterAccessLogRecordReconcilerFuncs) ReconcileAccessLogRecord(clusterName string, obj *observability_enterprise_mesh_gloo_solo_io_v1alpha1.AccessLogRecord) (reconcile.Result, error) {
	if f.OnReconcileAccessLogRecord == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessLogRecord(clusterName, obj)
}

func (f *MulticlusterAccessLogRecordReconcilerFuncs) ReconcileAccessLogRecordDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileAccessLogRecordDeletion == nil {
		return nil
	}
	return f.OnReconcileAccessLogRecordDeletion(clusterName, req)
}

type MulticlusterAccessLogRecordReconcileLoop interface {
	// AddMulticlusterAccessLogRecordReconciler adds a MulticlusterAccessLogRecordReconciler to the MulticlusterAccessLogRecordReconcileLoop.
	AddMulticlusterAccessLogRecordReconciler(ctx context.Context, rec MulticlusterAccessLogRecordReconciler, predicates ...predicate.Predicate)
}

type multiclusterAccessLogRecordReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAccessLogRecordReconcileLoop) AddMulticlusterAccessLogRecordReconciler(ctx context.Context, rec MulticlusterAccessLogRecordReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAccessLogRecordMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAccessLogRecordReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterAccessLogRecordReconcileLoop {
	return &multiclusterAccessLogRecordReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &observability_enterprise_mesh_gloo_solo_io_v1alpha1.AccessLogRecord{}, options)}
}

type genericAccessLogRecordMulticlusterReconciler struct {
	reconciler MulticlusterAccessLogRecordReconciler
}

func (g genericAccessLogRecordMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAccessLogRecordDeletionReconciler); ok {
		return deletionReconciler.ReconcileAccessLogRecordDeletion(cluster, req)
	}
	return nil
}

func (g genericAccessLogRecordMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*observability_enterprise_mesh_gloo_solo_io_v1alpha1.AccessLogRecord)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessLogRecord handler received event for %T", object)
	}
	return g.reconciler.ReconcileAccessLogRecord(cluster, obj)
}
