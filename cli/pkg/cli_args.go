package cli

import (
	"context"

	usageclient "github.com/solo-io/reporting-client/pkg/client"
	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/usage"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check"
	clusterroot "github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/demo"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/explore"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/install"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/istio"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/upgrade"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/version"
	"github.com/spf13/cobra"
)

// build an instance of the meshctl implementation
func BuildCli(
	ctx context.Context,
	opts *options.Options,
	usageReporter usageclient.Client,
	clusterCmd clusterroot.ClusterCommand,
	versionCmd version.VersionCommand,
	istioCmd istio.IstioCommand,
	upgradeCmd upgrade.UpgradeCommand,
	installCmd install.InstallCommand,
	uninstallCmd uninstall.UninstallCommand,
	checkCommand check.CheckCommand,
	exploreCommand explore.ExploreCommand,
	demoCommand demo.DemoCommand,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   cliconstants.RootCommand.Use,
		Short: cliconstants.RootCommand.Short,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			usageReporter.StartReportingUsage(ctx, usage.UsageReportingInterval)
			return nil
		},
	}
	options.AddRootFlags(cmd, opts)
	cmd.AddCommand(
		clusterCmd,
		versionCmd,
		installCmd,
		upgradeCmd,
		istioCmd,
		uninstallCmd,
		checkCommand,
		exploreCommand,
		demoCommand,
	)
	return cmd
}
