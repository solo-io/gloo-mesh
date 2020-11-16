package commands

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/dashboard"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/debug"

	"github.com/sirupsen/logrus"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/check"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/cluster"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/demo"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/describe"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/install"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/mesh"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/uninstall"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/version"

	"github.com/spf13/cobra"

	// required import to enable kube client-go auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func RootCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "meshctl [command]",
		Short: "The Command Line Interface for managing Gloo Mesh.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logrus.SetLevel(logrus.DebugLevel)
		},
	}

	cmd.AddCommand(
		cluster.Command(ctx),
		demo.Command(ctx),
		debug.Command(ctx),
		describe.Command(ctx),
		mesh.Command(ctx),
		install.Command(ctx),
		uninstall.Command(ctx),
		check.Command(ctx),
		dashboard.Command(ctx),
		version.Command(ctx),
	)

	return cmd
}
