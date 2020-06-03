package wire

import (
	"github.com/google/wire"
	"github.com/solo-io/service-mesh-hub/pkg/aws/aws_creds"
	"github.com/solo-io/service-mesh-hub/pkg/aws/clients"
	"github.com/solo-io/service-mesh-hub/pkg/aws/credentials"
	compute_target "github.com/solo-io/service-mesh-hub/services/common/compute-target"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/compute-target/k8s"
	aws2 "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/compute-target/aws"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/compute-target/aws"
)

func ComputeTargetCredentialsHandlersProvider(
	asyncManagerController *mc_manager.AsyncManagerController,
	awsCredsHandler aws2.AwsCredsHandler,
) []compute_target.ComputeTargetCredentialsHandler {
	return []compute_target.ComputeTargetCredentialsHandler{
		asyncManagerController,
		awsCredsHandler,
	}
}

var AwsSet = wire.NewSet(
	credentials.NewCredentialsGetter,
	aws.NewNetworkingAwsCredsHandler,
	aws_creds.DefaultSecretAwsCredsConverter,
	clients.STSClientFactoryProvider,
)
