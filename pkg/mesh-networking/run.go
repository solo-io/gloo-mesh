package mesh_networking

import (
	"context"
	"time"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/controller"
	mc_manager "github.com/solo-io/service-mesh-hub/pkg/common/compute-target/k8s"
	container_runtime "github.com/solo-io/service-mesh-hub/pkg/common/container-runtime"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/wire"
	"github.com/solo-io/skv2/pkg/reconcile"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func Run(ctx context.Context) {
	ctx = container_runtime.CreateRootContext(ctx, "mesh-networking")
	logger := contextutils.LoggerFrom(ctx)

	// build all the objects needed for multicluster operations
	meshNetworkingContext, err := wire.InitializeMeshNetworking(
		ctx,
	)
	if err != nil {
		logger.Fatalw("error initializing mesh networking clients", zap.Error(err))
	}

	// block until we die; RIP
	err = mc_manager.SetupAndStartLocalManager(
		meshNetworkingContext.MultiClusterDeps,
		[]mc_manager.AsyncManagerStartOptionsFunc{
			mc_manager.AddAllV1Alpha1ToScheme,
			mc_manager.AddAllIstioToScheme,
			mc_manager.AddAllLinkerdToScheme,
			startComponents(meshNetworkingContext),
		},
		[]mc_manager.NamedAsyncManagerHandler{{
			Name:                "mesh-networking-multicluster-controller",
			AsyncManagerHandler: meshNetworkingContext.MeshNetworkingClusterHandler,
		}},
	)

	if err != nil {
		logger.Fatalw("the local manager instance failed to start up or died with an error", zap.Error(err))
	}
}

// Controller-runtime Watches require the manager to be started first, otherwise it will block indefinitely
// Thus we initialize all components (and their associated watches) as an AsyncManagerStartOptionsFunc.
func startComponents(meshNetworkingContext wire.MeshNetworkingContext) func(context.Context, manager.Manager) error {
	return func(ctx context.Context, m manager.Manager) error {
		logger := contextutils.LoggerFrom(ctx)
		var err error
		if err = meshNetworkingContext.MeshNetworkingSnapshotContext.StartListening(
			contextutils.WithLogger(ctx, "mesh_networking_snapshot_listener"),
		); err != nil {
			logger.Fatalw("error initializing mesh networking snapshot listener", zap.Error(err))
		}

		go startTrafficPolicyReconciler(ctx, meshNetworkingContext)

		err = meshNetworkingContext.AccessControlPolicyTranslator.Start(
			contextutils.WithLogger(ctx, "access_control_policy_translator"),
		)
		if err != nil {
			logger.Fatalw("error initializing AccessControlPolicyTranslator", zap.Error(err))
		}

		err = meshNetworkingContext.GlobalAccessPolicyEnforcer.Start(
			contextutils.WithLogger(ctx, "global_access_control_policy_enforcer"),
		)
		if err != nil {
			logger.Fatalw("error initializing GlobalAccessControlPolicyEnforcer", zap.Error(err))
		}

		err = meshNetworkingContext.FederationResolver.Start(
			contextutils.WithLogger(ctx, "federation_resolver"),
		)
		if err != nil {
			logger.Fatalw("error initializing FederationResolver", zap.Error(err))
		}

		failoverServiceReconcileLoop := controller.NewFailoverServiceReconcileLoop("failover-service", m, reconcile.Options{})
		err = failoverServiceReconcileLoop.RunFailoverServiceReconciler(ctx, meshNetworkingContext.FailoverServiceReconciler)
		if err != nil {
			logger.Fatalw("error initializing FailoverServiceReconcileLoop", zap.Error(err))
		}

		return nil
	}
}

// This runs reconcile ever second. since it only writes things that have changed, and reads from cache
// it should generate load on the cluster. we plane to change this in the future for better responsiveness
func startTrafficPolicyReconciler(ctx context.Context, meshNetworkingContext wire.MeshNetworkingContext) {
	for {
		meshNetworkingContext.TrafficPolicyReconciler.Reconcile(ctx)
		select {
		case <-time.After(time.Second):
			continue
		case <-ctx.Done():
			return
		}
	}
}
