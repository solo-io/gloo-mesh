package istio_multicluster

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/demo/common/cleanup"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/demo/common/initialize"
	"github.com/spf13/cobra"
)

const (
	mgmtCluster   = "mgmt-cluster"
	remoteCluster = "remote-cluster"
)

func Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "istio-multicluster",
		Short: "Demo Gloo Mesh functionality with two Istio control planes deployed on separate k8s clusters",
	}

	cmd.AddCommand(
		initialize.IstioCommand(ctx, mgmtCluster, remoteCluster),
		cleanup.Command(ctx, mgmtCluster, remoteCluster),
	)

	return cmd
}
