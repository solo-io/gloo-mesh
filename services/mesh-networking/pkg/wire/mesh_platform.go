package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/aws_creds"
	mesh_platform "github.com/solo-io/service-mesh-hub/services/common/mesh-platform"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/mesh-platform/k8s"
	"github.com/solo-io/service-mesh-hub/services/common/mesh-platform/rest"
	v1 "k8s.io/api/core/v1"
)

func MeshPlatformCredentialsHandlersProvider(
	asyncManagerController *mc_manager.AsyncManagerController,
	awsCredsHandler rest.AwsCredsHandler,
) []mesh_platform.MeshPlatformCredentialsHandler {
	return []mesh_platform.MeshPlatformCredentialsHandler{
		asyncManagerController,
		awsCredsHandler,
	}
}

var AwsSet = wire.NewSet(
	NewNetworkingAwsCredsHandler,
	aws_creds.DefaultSecretAwsCredsConverter,
)

// Temporary stub until AppMesh translation is implemented
type networkingAwsCredsHandler struct {
}

func NewNetworkingAwsCredsHandler() rest.AwsCredsHandler {
	return &networkingAwsCredsHandler{}
}

func (n *networkingAwsCredsHandler) MeshPlatformAdded(ctx context.Context, secret *v1.Secret) error {
	return nil
}

func (n *networkingAwsCredsHandler) MeshPlatformRemoved(ctx context.Context, secret *v1.Secret) error {
	return nil
}
