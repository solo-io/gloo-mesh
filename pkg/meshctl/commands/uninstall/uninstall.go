package uninstall

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/codegen/helm"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/gloomesh"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context, globalFlags *utils.GlobalFlags) *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall Gloo Mesh from the referenced cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.verbose = globalFlags.Verbose
			return uninstall(ctx, opts)
		},
	}
	opts.addToFlags(cmd.Flags())
	cmd.SilenceUsage = true
	return cmd
}

type options struct {
	kubeconfig  string
	kubecontext string
	namespace   string
	releaseName string
	verbose     bool
	dryRun      bool
}

func (o *options) addToFlags(flags *pflag.FlagSet) {
	utils.AddManagementKubeconfigFlags(&o.kubeconfig, &o.kubecontext, flags)
	flags.BoolVarP(&o.dryRun, "dry-run", "d", false, "Output installation manifest")
	flags.StringVar(&o.namespace, "namespace", defaults.DefaultPodNamespace, "namespace in which to install Gloo Mesh")
	flags.StringVar(&o.releaseName, "release-name", helm.Chart.Data.Name, "Helm release name")
}

func uninstall(ctx context.Context, opts *options) error {
	if err := uninstallGlooMesh(ctx, opts); err != nil {
		return eris.Wrap(err, "uninstalling gloo-mesh")
	}
	return nil
}

func uninstallGlooMesh(ctx context.Context, opts *options) error {
	return gloomesh.Uninstaller{
		KubeConfig:  opts.kubeconfig,
		KubeContext: opts.kubecontext,
		Namespace:   opts.namespace,
		ReleaseName: opts.releaseName,
		Verbose:     opts.verbose,
		DryRun:      opts.dryRun,
	}.UninstallGlooMesh(
		ctx,
	)
}
