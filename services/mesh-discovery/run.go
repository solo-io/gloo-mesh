package mesh_discovery

import (
	"context"

	"github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh-workload/istio"
	"github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh-workload/linkerd"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/service-mesh-hub/services/common/multicluster"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/multicluster/manager"
	"github.com/solo-io/service-mesh-hub/services/internal/config"
	"github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh"
	mesh_workload "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh-workload"
	md_multicluster "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/multicluster"
	"github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/wire"
	"go.uber.org/zap"
)

func Run(rootCtx context.Context) {
	ctx := config.CreateRootContext(rootCtx, "mesh-discovery")

	logger := contextutils.LoggerFrom(ctx)

	// build all the objects needed for multicluster operations
	discoveryContext, err := wire.InitializeDiscovery(ctx)
	if err != nil {
		logger.Fatalw("error initializing discovery clients", zap.Error(err))
	}

	localManager := discoveryContext.MultiClusterDeps.LocalManager

	// this is our main entrypoint for mesh-discovery
	// when it detects a new cluster get registered, it builds a deployment controller
	// with the controller factory, and attaches the given mesh finders to it
	deploymentHandler, err := md_multicluster.NewDiscoveryClusterHandler(
		localManager,
		[]mesh.MeshScanner{
			discoveryContext.MeshDiscovery.IstioMeshScanner,
			discoveryContext.MeshDiscovery.ConsulConnectMeshScanner,
			discoveryContext.MeshDiscovery.LinkerdMeshScanner,
		},
		[]mesh_workload.MeshWorkloadScannerFactory{
			istio.NewIstioMeshWorkloadScanner,
			linkerd.NewLinkerdMeshWorkloadScanner,
		},
		discoveryContext,
	)

	if err != nil {
		logger.Fatalw("error initializing discovery cluster handler", zap.Error(err))
	}

	// block until we die; RIP
	err = multicluster.SetupAndStartLocalManager(
		discoveryContext.MultiClusterDeps,

		// need to be sure to register the v1alpha1 CRDs with the controller runtime
		[]mc_manager.AsyncManagerStartOptionsFunc{multicluster.AddAllV1Alpha1ToScheme},

		[]multicluster.NamedAsyncManagerHandler{{
			Name:                "discovery-controller",
			AsyncManagerHandler: deploymentHandler,
		}},
	)

	if err != nil {
		logger.Fatalw("the local manager instance failed to start up or died with an error", zap.Error(err))
	}
}
