package snapshot

import (
	"context"
	"sync"
	"time"

	"github.com/solo-io/go-utils/contextutils"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_networking_controller "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/controller"
	"go.uber.org/zap"
)

type MeshNetworkingSnapshot struct {
	MeshServices  []*zephyr_discovery.MeshService
	VirtualMeshes []*zephyr_networking.VirtualMesh
	MeshWorkloads []*zephyr_discovery.MeshWorkload
}

type UpdatedMeshService struct {
	Old *zephyr_discovery.MeshService
	New *zephyr_discovery.MeshService
}

type UpdatedVirtualMesh struct {
	Old *zephyr_networking.VirtualMesh
	New *zephyr_networking.VirtualMesh
}

type UpdatedMeshWorkload struct {
	Old *zephyr_discovery.MeshWorkload
	New *zephyr_discovery.MeshWorkload
}

type UpdatedResources struct {
	MeshServices  []UpdatedMeshService
	VirtualMeshes []UpdatedVirtualMesh
	MeshWorkloads []UpdatedMeshWorkload
}

// an implementation of `MeshNetworkingSnapshotGenerator` that is guaranteed to only ever push
// snapshots that are considered valid by the `MeshNetworkingSnapshotValidator` to its listeners
func NewMeshNetworkingSnapshotGenerator(
	ctx context.Context,
	snapshotValidator MeshNetworkingSnapshotValidator,
	MeshServiceEventWatcher zephyr_discovery_controller.MeshServiceEventWatcher,
	virtualMeshEventWatcher zephyr_networking_controller.VirtualMeshEventWatcher,
	meshWorkloadEventWatcher zephyr_discovery_controller.MeshWorkloadEventWatcher,
) (MeshNetworkingSnapshotGenerator, error) {
	generator := &networkingSnapshotGenerator{
		snapshotValidator: snapshotValidator,
		snapshot:          MeshNetworkingSnapshot{},
	}

	err := MeshServiceEventWatcher.AddEventHandler(ctx, &zephyr_discovery_controller.MeshServiceEventHandlerFuncs{
		OnCreate: func(obj *zephyr_discovery.MeshService) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			updatedMeshServices := append([]*zephyr_discovery.MeshService{}, generator.snapshot.MeshServices...)
			updatedMeshServices = append(updatedMeshServices, obj)

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshServices = updatedMeshServices
			if generator.snapshotValidator.ValidateMeshServiceUpsert(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnUpdate: func(old, new *zephyr_discovery.MeshService) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshServices []*zephyr_discovery.MeshService
			for _, existingMeshService := range generator.snapshot.MeshServices {
				if existingMeshService.GetName() == old.GetName() && existingMeshService.GetNamespace() == old.GetNamespace() {
					updatedMeshServices = append(updatedMeshServices, new)
				} else {
					updatedMeshServices = append(updatedMeshServices, existingMeshService)
				}
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshServices = updatedMeshServices

			if generator.snapshotValidator.ValidateMeshServiceUpsert(ctx, new, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnDelete: func(obj *zephyr_discovery.MeshService) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshServices []*zephyr_discovery.MeshService
			for _, meshService := range generator.snapshot.MeshServices {
				if meshService.GetName() == obj.GetName() && meshService.GetNamespace() == obj.GetNamespace() {
					continue
				}

				updatedMeshServices = append(updatedMeshServices, meshService)
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshServices = updatedMeshServices

			if generator.snapshotValidator.ValidateMeshServiceDelete(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnGeneric: func(obj *zephyr_discovery.MeshService) error {
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	err = virtualMeshEventWatcher.AddEventHandler(ctx, &zephyr_networking_controller.VirtualMeshEventHandlerFuncs{
		OnCreate: func(obj *zephyr_networking.VirtualMesh) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			updatedVirtualMeshes := append([]*zephyr_networking.VirtualMesh{}, generator.snapshot.VirtualMeshes...)
			updatedVirtualMeshes = append(updatedVirtualMeshes, obj)

			updatedSnapshot := generator.snapshot
			updatedSnapshot.VirtualMeshes = updatedVirtualMeshes

			if generator.snapshotValidator.ValidateVirtualMeshUpsert(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnUpdate: func(old, new *zephyr_networking.VirtualMesh) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedVirtualMeshes []*zephyr_networking.VirtualMesh
			for _, existingVirtualMesh := range generator.snapshot.VirtualMeshes {
				if existingVirtualMesh.GetName() == old.GetName() && existingVirtualMesh.GetNamespace() == old.GetNamespace() {
					updatedVirtualMeshes = append(updatedVirtualMeshes, new)
				} else {
					updatedVirtualMeshes = append(updatedVirtualMeshes, existingVirtualMesh)
				}
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.VirtualMeshes = updatedVirtualMeshes

			if generator.snapshotValidator.ValidateVirtualMeshUpsert(ctx, new, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnDelete: func(obj *zephyr_networking.VirtualMesh) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedVirtualMeshes []*zephyr_networking.VirtualMesh
			for _, virtualMesh := range generator.snapshot.VirtualMeshes {
				if virtualMesh.GetName() == obj.GetName() && virtualMesh.GetNamespace() == obj.GetNamespace() {
					continue
				}

				updatedVirtualMeshes = append(updatedVirtualMeshes, virtualMesh)
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.VirtualMeshes = updatedVirtualMeshes

			if generator.snapshotValidator.ValidateVirtualMeshDelete(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnGeneric: func(obj *zephyr_networking.VirtualMesh) error {
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	err = meshWorkloadEventWatcher.AddEventHandler(ctx, &zephyr_discovery_controller.MeshWorkloadEventHandlerFuncs{
		OnCreate: func(obj *zephyr_discovery.MeshWorkload) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			updatedMeshWorkloads := append([]*zephyr_discovery.MeshWorkload{}, generator.snapshot.MeshWorkloads...)
			updatedMeshWorkloads = append(updatedMeshWorkloads, obj)

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshWorkloads = updatedMeshWorkloads

			if generator.snapshotValidator.ValidateMeshWorkloadUpsert(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnUpdate: func(old, new *zephyr_discovery.MeshWorkload) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshWorkloads []*zephyr_discovery.MeshWorkload
			for _, existingMeshWorkload := range generator.snapshot.MeshWorkloads {
				if existingMeshWorkload.GetName() == old.GetName() && existingMeshWorkload.GetNamespace() == old.GetNamespace() {
					updatedMeshWorkloads = append(updatedMeshWorkloads, new)
				} else {
					updatedMeshWorkloads = append(updatedMeshWorkloads, existingMeshWorkload)
				}
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshWorkloads = updatedMeshWorkloads

			if generator.snapshotValidator.ValidateMeshWorkloadUpsert(ctx, new, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnDelete: func(obj *zephyr_discovery.MeshWorkload) error {
			generator.snapshotMutex.Lock()
			defer generator.snapshotMutex.Unlock()

			var updatedMeshWorkloads []*zephyr_discovery.MeshWorkload
			for _, meshWorkload := range generator.snapshot.MeshWorkloads {
				if meshWorkload.GetName() == obj.GetName() && meshWorkload.GetNamespace() == obj.GetNamespace() {
					continue
				}

				updatedMeshWorkloads = append(updatedMeshWorkloads, meshWorkload)
			}

			updatedSnapshot := generator.snapshot
			updatedSnapshot.MeshWorkloads = updatedMeshWorkloads

			if generator.snapshotValidator.ValidateMeshWorkloadDelete(ctx, obj, &updatedSnapshot) {
				generator.isSnapshotPushNeeded = true
				generator.snapshot = updatedSnapshot
			}

			return nil
		},
		OnGeneric: func(obj *zephyr_discovery.MeshWorkload) error {
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	return generator, nil
}

type networkingSnapshotGenerator struct {
	snapshotValidator MeshNetworkingSnapshotValidator

	listeners     []MeshNetworkingSnapshotListener
	listenerMutex sync.Mutex

	// important that snapshot is NOT a reference- we depend on being able to copy it
	// and change fields without mutating the real thing
	// accesses to `isSnapshotPushNeeded` should be gated on the `snapshotMutex`
	snapshot MeshNetworkingSnapshot
	// version of the snapshot being sent, will appear in the logger context values
	version              uint
	isSnapshotPushNeeded bool
	snapshotMutex        sync.Mutex
}

func (f *networkingSnapshotGenerator) RegisterListener(listener MeshNetworkingSnapshotListener) {
	f.listenerMutex.Lock()
	defer f.listenerMutex.Unlock()

	f.listeners = append(f.listeners, listener)
}

func (f *networkingSnapshotGenerator) StartPushingSnapshots(ctx context.Context, snapshotFrequency time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(snapshotFrequency):
			f.snapshotMutex.Lock()
			f.listenerMutex.Lock()

			if f.isSnapshotPushNeeded {
				f.version++
				snapshotContext := contextutils.WithLoggerValues(ctx,
					zap.Uint("snapshot_version", f.version),
					zap.Int("num_mesh_services", len(f.snapshot.MeshServices)),
					zap.Int("num_mesh_workloads", len(f.snapshot.MeshWorkloads)),
					zap.Int("num_virtual_meshs", len(f.snapshot.VirtualMeshes)),
				)
				for _, listener := range f.listeners {
					listener.Sync(snapshotContext, &f.snapshot)
				}

				f.isSnapshotPushNeeded = false
			}

			// important to unlock the mutexes in the same order as they were locked here
			// it's a runtime error to attempt to unlock an already unlocked mutex
			// if the order is changed here, a race condition could cause a repeated unlock
			f.snapshotMutex.Unlock()
			f.listenerMutex.Unlock()
		}
	}
}
