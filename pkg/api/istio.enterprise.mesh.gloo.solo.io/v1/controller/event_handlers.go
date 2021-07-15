// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./event_handlers.go -destination mocks/event_handlers.go

// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	istio_enterprise_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/istio.enterprise.mesh.gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the IstioInstallation Resource
// DEPRECATED: Prefer reconciler pattern.
type IstioInstallationEventHandler interface {
	CreateIstioInstallation(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
	UpdateIstioInstallation(old, new *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
	DeleteIstioInstallation(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
	GenericIstioInstallation(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
}

type IstioInstallationEventHandlerFuncs struct {
	OnCreate  func(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
	OnUpdate  func(old, new *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
	OnDelete  func(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
	OnGeneric func(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error
}

func (f *IstioInstallationEventHandlerFuncs) CreateIstioInstallation(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *IstioInstallationEventHandlerFuncs) DeleteIstioInstallation(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *IstioInstallationEventHandlerFuncs) UpdateIstioInstallation(objOld, objNew *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *IstioInstallationEventHandlerFuncs) GenericIstioInstallation(obj *istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type IstioInstallationEventWatcher interface {
	AddEventHandler(ctx context.Context, h IstioInstallationEventHandler, predicates ...predicate.Predicate) error
}

type istioInstallationEventWatcher struct {
	watcher events.EventWatcher
}

func NewIstioInstallationEventWatcher(name string, mgr manager.Manager) IstioInstallationEventWatcher {
	return &istioInstallationEventWatcher{
		watcher: events.NewWatcher(name, mgr, &istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation{}),
	}
}

func (c *istioInstallationEventWatcher) AddEventHandler(ctx context.Context, h IstioInstallationEventHandler, predicates ...predicate.Predicate) error {
	handler := genericIstioInstallationHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericIstioInstallationHandler implements a generic events.EventHandler
type genericIstioInstallationHandler struct {
	handler IstioInstallationEventHandler
}

func (h genericIstioInstallationHandler) Create(object client.Object) error {
	obj, ok := object.(*istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation)
	if !ok {
		return errors.Errorf("internal error: IstioInstallation handler received event for %T", object)
	}
	return h.handler.CreateIstioInstallation(obj)
}

func (h genericIstioInstallationHandler) Delete(object client.Object) error {
	obj, ok := object.(*istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation)
	if !ok {
		return errors.Errorf("internal error: IstioInstallation handler received event for %T", object)
	}
	return h.handler.DeleteIstioInstallation(obj)
}

func (h genericIstioInstallationHandler) Update(old, new client.Object) error {
	objOld, ok := old.(*istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation)
	if !ok {
		return errors.Errorf("internal error: IstioInstallation handler received event for %T", old)
	}
	objNew, ok := new.(*istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation)
	if !ok {
		return errors.Errorf("internal error: IstioInstallation handler received event for %T", new)
	}
	return h.handler.UpdateIstioInstallation(objOld, objNew)
}

func (h genericIstioInstallationHandler) Generic(object client.Object) error {
	obj, ok := object.(*istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallation)
	if !ok {
		return errors.Errorf("internal error: IstioInstallation handler received event for %T", object)
	}
	return h.handler.GenericIstioInstallation(obj)
}