// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	networking_zephyr_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the TrafficPolicy Resource.
// implemented by the user
type TrafficPolicyReconciler interface {
	ReconcileTrafficPolicy(obj *networking_zephyr_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the TrafficPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type TrafficPolicyDeletionReconciler interface {
	ReconcileTrafficPolicyDeletion(req reconcile.Request)
}

type TrafficPolicyReconcilerFuncs struct {
	OnReconcileTrafficPolicy         func(obj *networking_zephyr_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error)
	OnReconcileTrafficPolicyDeletion func(req reconcile.Request)
}

func (f *TrafficPolicyReconcilerFuncs) ReconcileTrafficPolicy(obj *networking_zephyr_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	if f.OnReconcileTrafficPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTrafficPolicy(obj)
}

func (f *TrafficPolicyReconcilerFuncs) ReconcileTrafficPolicyDeletion(req reconcile.Request) {
	if f.OnReconcileTrafficPolicyDeletion == nil {
		return
	}
	f.OnReconcileTrafficPolicyDeletion(req)
}

// Reconcile and finalize the TrafficPolicy Resource
// implemented by the user
type TrafficPolicyFinalizer interface {
	TrafficPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	TrafficPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeTrafficPolicy(obj *networking_zephyr_solo_io_v1alpha1.TrafficPolicy) error
}

type TrafficPolicyReconcileLoop interface {
	RunTrafficPolicyReconciler(ctx context.Context, rec TrafficPolicyReconciler, predicates ...predicate.Predicate) error
}

type trafficPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewTrafficPolicyReconcileLoop(name string, mgr manager.Manager) TrafficPolicyReconcileLoop {
	return &trafficPolicyReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &networking_zephyr_solo_io_v1alpha1.TrafficPolicy{}),
	}
}

func (c *trafficPolicyReconcileLoop) RunTrafficPolicyReconciler(ctx context.Context, reconciler TrafficPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericTrafficPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(TrafficPolicyFinalizer); ok {
		reconcilerWrapper = genericTrafficPolicyFinalizer{
			genericTrafficPolicyReconciler: genericReconciler,
			finalizingReconciler:           finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericTrafficPolicyHandler implements a generic reconcile.Reconciler
type genericTrafficPolicyReconciler struct {
	reconciler TrafficPolicyReconciler
}

func (r genericTrafficPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_zephyr_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileTrafficPolicy(obj)
}

func (r genericTrafficPolicyReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(TrafficPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileTrafficPolicyDeletion(request)
	}
}

// genericTrafficPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericTrafficPolicyFinalizer struct {
	genericTrafficPolicyReconciler
	finalizingReconciler TrafficPolicyFinalizer
}

func (r genericTrafficPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.TrafficPolicyFinalizerName()
}

func (r genericTrafficPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_zephyr_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeTrafficPolicy(obj)
}

// Reconcile Upsert events for the AccessControlPolicy Resource.
// implemented by the user
type AccessControlPolicyReconciler interface {
	ReconcileAccessControlPolicy(obj *networking_zephyr_solo_io_v1alpha1.AccessControlPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessControlPolicy Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type AccessControlPolicyDeletionReconciler interface {
	ReconcileAccessControlPolicyDeletion(req reconcile.Request)
}

type AccessControlPolicyReconcilerFuncs struct {
	OnReconcileAccessControlPolicy         func(obj *networking_zephyr_solo_io_v1alpha1.AccessControlPolicy) (reconcile.Result, error)
	OnReconcileAccessControlPolicyDeletion func(req reconcile.Request)
}

func (f *AccessControlPolicyReconcilerFuncs) ReconcileAccessControlPolicy(obj *networking_zephyr_solo_io_v1alpha1.AccessControlPolicy) (reconcile.Result, error) {
	if f.OnReconcileAccessControlPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessControlPolicy(obj)
}

func (f *AccessControlPolicyReconcilerFuncs) ReconcileAccessControlPolicyDeletion(req reconcile.Request) {
	if f.OnReconcileAccessControlPolicyDeletion == nil {
		return
	}
	f.OnReconcileAccessControlPolicyDeletion(req)
}

// Reconcile and finalize the AccessControlPolicy Resource
// implemented by the user
type AccessControlPolicyFinalizer interface {
	AccessControlPolicyReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	AccessControlPolicyFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeAccessControlPolicy(obj *networking_zephyr_solo_io_v1alpha1.AccessControlPolicy) error
}

type AccessControlPolicyReconcileLoop interface {
	RunAccessControlPolicyReconciler(ctx context.Context, rec AccessControlPolicyReconciler, predicates ...predicate.Predicate) error
}

type accessControlPolicyReconcileLoop struct {
	loop reconcile.Loop
}

func NewAccessControlPolicyReconcileLoop(name string, mgr manager.Manager) AccessControlPolicyReconcileLoop {
	return &accessControlPolicyReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &networking_zephyr_solo_io_v1alpha1.AccessControlPolicy{}),
	}
}

func (c *accessControlPolicyReconcileLoop) RunAccessControlPolicyReconciler(ctx context.Context, reconciler AccessControlPolicyReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericAccessControlPolicyReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(AccessControlPolicyFinalizer); ok {
		reconcilerWrapper = genericAccessControlPolicyFinalizer{
			genericAccessControlPolicyReconciler: genericReconciler,
			finalizingReconciler:                 finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericAccessControlPolicyHandler implements a generic reconcile.Reconciler
type genericAccessControlPolicyReconciler struct {
	reconciler AccessControlPolicyReconciler
}

func (r genericAccessControlPolicyReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_zephyr_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessControlPolicy handler received event for %T", object)
	}
	return r.reconciler.ReconcileAccessControlPolicy(obj)
}

func (r genericAccessControlPolicyReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(AccessControlPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileAccessControlPolicyDeletion(request)
	}
}

// genericAccessControlPolicyFinalizer implements a generic reconcile.FinalizingReconciler
type genericAccessControlPolicyFinalizer struct {
	genericAccessControlPolicyReconciler
	finalizingReconciler AccessControlPolicyFinalizer
}

func (r genericAccessControlPolicyFinalizer) FinalizerName() string {
	return r.finalizingReconciler.AccessControlPolicyFinalizerName()
}

func (r genericAccessControlPolicyFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_zephyr_solo_io_v1alpha1.AccessControlPolicy)
	if !ok {
		return errors.Errorf("internal error: AccessControlPolicy handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeAccessControlPolicy(obj)
}

// Reconcile Upsert events for the VirtualMesh Resource.
// implemented by the user
type VirtualMeshReconciler interface {
	ReconcileVirtualMesh(obj *networking_zephyr_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualMesh Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type VirtualMeshDeletionReconciler interface {
	ReconcileVirtualMeshDeletion(req reconcile.Request)
}

type VirtualMeshReconcilerFuncs struct {
	OnReconcileVirtualMesh         func(obj *networking_zephyr_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error)
	OnReconcileVirtualMeshDeletion func(req reconcile.Request)
}

func (f *VirtualMeshReconcilerFuncs) ReconcileVirtualMesh(obj *networking_zephyr_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error) {
	if f.OnReconcileVirtualMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualMesh(obj)
}

func (f *VirtualMeshReconcilerFuncs) ReconcileVirtualMeshDeletion(req reconcile.Request) {
	if f.OnReconcileVirtualMeshDeletion == nil {
		return
	}
	f.OnReconcileVirtualMeshDeletion(req)
}

// Reconcile and finalize the VirtualMesh Resource
// implemented by the user
type VirtualMeshFinalizer interface {
	VirtualMeshReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	VirtualMeshFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeVirtualMesh(obj *networking_zephyr_solo_io_v1alpha1.VirtualMesh) error
}

type VirtualMeshReconcileLoop interface {
	RunVirtualMeshReconciler(ctx context.Context, rec VirtualMeshReconciler, predicates ...predicate.Predicate) error
}

type virtualMeshReconcileLoop struct {
	loop reconcile.Loop
}

func NewVirtualMeshReconcileLoop(name string, mgr manager.Manager) VirtualMeshReconcileLoop {
	return &virtualMeshReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &networking_zephyr_solo_io_v1alpha1.VirtualMesh{}),
	}
}

func (c *virtualMeshReconcileLoop) RunVirtualMeshReconciler(ctx context.Context, reconciler VirtualMeshReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericVirtualMeshReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(VirtualMeshFinalizer); ok {
		reconcilerWrapper = genericVirtualMeshFinalizer{
			genericVirtualMeshReconciler: genericReconciler,
			finalizingReconciler:         finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(ctx, reconcilerWrapper, predicates...)
}

// genericVirtualMeshHandler implements a generic reconcile.Reconciler
type genericVirtualMeshReconciler struct {
	reconciler VirtualMeshReconciler
}

func (r genericVirtualMeshReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_zephyr_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return r.reconciler.ReconcileVirtualMesh(obj)
}

func (r genericVirtualMeshReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(VirtualMeshDeletionReconciler); ok {
		deletionReconciler.ReconcileVirtualMeshDeletion(request)
	}
}

// genericVirtualMeshFinalizer implements a generic reconcile.FinalizingReconciler
type genericVirtualMeshFinalizer struct {
	genericVirtualMeshReconciler
	finalizingReconciler VirtualMeshFinalizer
}

func (r genericVirtualMeshFinalizer) FinalizerName() string {
	return r.finalizingReconciler.VirtualMeshFinalizerName()
}

func (r genericVirtualMeshFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*networking_zephyr_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeVirtualMesh(obj)
}
