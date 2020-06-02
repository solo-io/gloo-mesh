package access_policy_enforcer

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	access_control_enforcer "github.com/solo-io/service-mesh-hub/pkg/access-control/enforcer"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_networking_controller "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/controller"
	container_runtime "github.com/solo-io/service-mesh-hub/pkg/container-runtime"
	"github.com/solo-io/service-mesh-hub/pkg/kube/selection"
	"go.uber.org/zap"
)

type enforcerLoop struct {
	virtualMeshEventWatcher zephyr_networking_controller.VirtualMeshEventWatcher
	virtualMeshClient       zephyr_networking.VirtualMeshClient
	meshClient              zephyr_discovery.MeshClient
	meshEnforcers           []access_control_enforcer.AccessPolicyMeshEnforcer
}

func NewEnforcerLoop(
	virtualMeshEventWatcher zephyr_networking_controller.VirtualMeshEventWatcher,
	virtualMeshClient zephyr_networking.VirtualMeshClient,
	meshClient zephyr_discovery.MeshClient,
	meshEnforcers []access_control_enforcer.AccessPolicyMeshEnforcer,
) AccessPolicyEnforcerLoop {
	return &enforcerLoop{
		virtualMeshEventWatcher: virtualMeshEventWatcher,
		virtualMeshClient:       virtualMeshClient,
		meshClient:              meshClient,
		meshEnforcers:           meshEnforcers,
	}
}

func (e *enforcerLoop) Start(ctx context.Context) error {
	return e.virtualMeshEventWatcher.AddEventHandler(ctx, &zephyr_networking_controller.VirtualMeshEventHandlerFuncs{
		OnCreate: func(obj *zephyr_networking.VirtualMesh) error {
			logger := container_runtime.BuildEventLogger(ctx, container_runtime.CreateEvent, obj)
			logger.Debugw("event handler enter",
				zap.Any("spec", obj.Spec),
				zap.Any("status", obj.Status),
			)
			return e.reconcile(ctx, container_runtime.CreateEvent)
		},
		OnUpdate: func(old, new *zephyr_networking.VirtualMesh) error {
			logger := container_runtime.BuildEventLogger(ctx, container_runtime.UpdateEvent, new)
			logger.Debugw("event handler enter",
				zap.Any("old_spec", old.Spec),
				zap.Any("old_status", old.Status),
				zap.Any("new_spec", new.Spec),
				zap.Any("new_status", new.Status),
			)
			return e.reconcile(ctx, container_runtime.UpdateEvent)
		},
		OnDelete: func(obj *zephyr_networking.VirtualMesh) error {
			logger := container_runtime.BuildEventLogger(ctx, container_runtime.DeleteEvent, obj)
			logger.Debugw("event handler enter",
				zap.Any("spec", obj.Spec),
				zap.Any("status", obj.Status),
			)
			return e.reconcile(ctx, container_runtime.DeleteEvent)
		},
		OnGeneric: func(virtualMesh *zephyr_networking.VirtualMesh) error {
			logger := container_runtime.BuildEventLogger(ctx, container_runtime.GenericEvent, virtualMesh)
			logger.Debugf("Ignoring event: %+v", virtualMesh)
			return nil
		},
	})
}

func (e *enforcerLoop) reconcile(ctx context.Context, eventType container_runtime.EventType) error {
	vmList, err := e.virtualMeshClient.ListVirtualMesh(ctx)
	if err != nil {
		return err
	}
	for _, vm := range vmList.Items {
		vm := vm
		logger := container_runtime.BuildEventLogger(ctx, eventType, &vm)
		err = e.enforceGlobalAccessControl(ctx, &vm)
		if err != nil {
			logger.Errorf("Error while handling VirtualMesh event: %+v", err)
		}
		if eventType != container_runtime.DeleteEvent {
			statusErr := e.setStatus(ctx, &vm, err)
			if statusErr != nil {
				logger.Errorf("Error while settings status for VirtualMesh: %+v", statusErr)
				return statusErr
			}
		}
	}
	return err
}

// If enforceMeshDefault, ignore user declared enforce_access_control setting and use mesh-specific default as defined on the
// VirtualMesh API, used for VM deletion to clean up mesh resources.
// If virtualMesh is nil, enforce the mesh specific default.
func (e *enforcerLoop) enforceGlobalAccessControl(
	ctx context.Context,
	virtualMesh *zephyr_networking.VirtualMesh,
) error {
	meshes, err := e.fetchMeshes(ctx, virtualMesh)
	if err != nil {
		return err
	}
	if len(meshes) == 0 {
		return nil
	}
	for _, mesh := range meshes {
		for _, meshEnforcer := range e.meshEnforcers {
			err = meshEnforcer.ReconcileAccessControl(
				contextutils.WithLogger(ctx, meshEnforcer.Name()),
				mesh,
				virtualMesh,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *enforcerLoop) fetchMeshes(
	ctx context.Context,
	virtualMesh *zephyr_networking.VirtualMesh,
) ([]*zephyr_discovery.Mesh, error) {
	var meshes []*zephyr_discovery.Mesh
	for _, meshRef := range virtualMesh.Spec.GetMeshes() {
		mesh, err := e.meshClient.GetMesh(ctx, selection.ResourceRefToObjectKey(meshRef))
		if err != nil {
			return nil, err
		}
		meshes = append(meshes, mesh)
	}
	return meshes, nil
}

func (e *enforcerLoop) setStatus(
	ctx context.Context,
	virtualMesh *zephyr_networking.VirtualMesh,
	err error,
) error {
	if err != nil {
		virtualMesh.Status.AccessControlEnforcementStatus = &zephyr_core_types.Status{
			State:   zephyr_core_types.Status_PROCESSING_ERROR,
			Message: err.Error(),
		}
	} else {
		virtualMesh.Status.AccessControlEnforcementStatus = &zephyr_core_types.Status{
			State: zephyr_core_types.Status_ACCEPTED,
		}
	}
	return e.virtualMeshClient.UpdateVirtualMeshStatus(ctx, virtualMesh)
}
