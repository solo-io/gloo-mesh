package istio_multicluster

import (
	"context"

	"github.com/solo-io/service-mesh-hub/pkg/meshctl/commands/demo/common/cleanup"
	"github.com/solo-io/service-mesh-hub/pkg/meshctl/commands/demo/common/initialize"
	"github.com/spf13/cobra"
)

const (
	mgmtCluster   = "mgmt-cluster"
	remoteCluster = "remote-cluster"
)

func Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "istio-multicluster",
		Short: "Demo Service Mesh Hub functionality with two Istio control planes deployed on separate k8s clusters",
		Long: `
Demo Service Mesh Hub functionality with two Istio control planes deployed on separate k8s clusters.

Requires kubectl >= v1.18.8, kind >= v0.8.1, istioctl <= v1.7.0, and docker.
We recommend allocating at least 8GB of RAM for Docker.
`,
	}

	cmd.AddCommand(
		initialize.IstioCommand(ctx, mgmtCluster, remoteCluster),
		cleanup.Command(ctx, mgmtCluster, remoteCluster),
	)

	return cmd
}
