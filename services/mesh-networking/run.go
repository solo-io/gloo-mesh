package mesh_networking

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/service-mesh-hub/services/common/multicluster"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/multicluster/manager"
	"github.com/solo-io/service-mesh-hub/services/internal/config"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/wire"
	"go.uber.org/zap"
)

func Run(ctx context.Context) {
	ctx = config.CreateRootContext(ctx, "mesh-networking")
	logger := contextutils.LoggerFrom(ctx)

	// build all the objects needed for multicluster operations
	meshNetworkingContext, err := wire.InitializeMeshNetworking(
		contextutils.WithLogger(ctx, "access_control_enforcer"),
	)
	if err != nil {
		logger.Fatalw("error initializing mesh networking clients", zap.Error(err))
	}
	if err = meshNetworkingContext.MeshNetworkingSnapshotContext.StartListening(
		contextutils.WithLogger(ctx, "mesh_networking_snapshot_listener"),
	); err != nil {
		logger.Fatalw("error initializing mesh networking snapshot listener", zap.Error(err))
	}
	// start the TrafficPolicyTranslator
	err = meshNetworkingContext.TrafficPolicyTranslator.Start(
		contextutils.WithLogger(ctx, "traffic_policy_translator"),
	)
	if err != nil {
		logger.Fatalw("error initializing TrafficPolicyTranslator", zap.Error(err))
	}

	err = meshNetworkingContext.AccessControlPolicyTranslator.Start(
		contextutils.WithLogger(ctx, "access_control_policy_translator"),
	)
	if err != nil {
		logger.Fatalw("error intitializing AccessControlPolicyTranslator", zap.Error(err))
	}

	err = meshNetworkingContext.GlobalAccessPolicyEnforcer.Start(
		contextutils.WithLogger(ctx, "global_access_control_policy_enforcer"),
	)
	if err != nil {
		logger.Fatalw("error intitializing GlobalAccessControlPolicyEnforcer", zap.Error(err))
	}

	err = meshNetworkingContext.FederationResolver.Start(
		contextutils.WithLogger(ctx, "federation_resolver"),
	)
	if err != nil {
		logger.Fatalw("error intitializing FederationResolver", zap.Error(err))
	}

	// block until we die; RIP
	err = multicluster.SetupAndStartLocalManager(
		meshNetworkingContext.MultiClusterDeps,
		[]mc_manager.AsyncManagerStartOptionsFunc{
			multicluster.AddAllV1Alpha1ToScheme,
			multicluster.AddAllIstioToScheme,
		},
		[]multicluster.NamedAsyncManagerHandler{{
			Name:                "mesh-networking-multicluster-controller",
			AsyncManagerHandler: meshNetworkingContext.MeshNetworkingClusterHandler,
		}},
	)

	if err != nil {
		logger.Fatalw("the local manager instance failed to start up or died with an error", zap.Error(err))
	}
}
