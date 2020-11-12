package install

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/codegen/helm"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/common/version"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/gloomesh"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/registration"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/utils"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context) *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install Gloo Mesh",
		RunE: func(cmd *cobra.Command, args []string) error {
			return install(ctx, opts)
		},
	}
	opts.addToFlags(cmd.Flags())
	cmd.SilenceUsage = true
	return cmd
}

type options struct {
	kubeCfgPath     string
	kubeContext     string
	namespace       string
	chartPath       string
	chartValuesFile string
	releaseName     string
	version         string
	verbose         bool
	dryRun          bool
	registrationOptions
}

type registrationOptions struct {
	register            bool
	clusterName         string
	apiServerAddress    string
	clusterDomain       string
	certAgentChartPath  string
	certAgentValuesPath string
}

func (o *options) addToFlags(flags *pflag.FlagSet) {
	utils.AddManagementKubeconfigFlags(&o.kubeCfgPath, &o.kubeContext, flags)
	flags.BoolVarP(&o.dryRun, "dry-run", "d", false, "Output installation manifest")
	flags.StringVar(&o.namespace, "namespace", defaults.DefaultPodNamespace, "namespace in which to install Gloo Mesh")
	flags.StringVar(&o.chartPath, "chart-file", "", "Path to a local Helm chart for installing Gloo Mesh. If unset, this command will install Gloo Mesh from the publicly released Helm chart.")
	flags.StringVarP(&o.chartValuesFile, "chart-values-file", "", "", "File containing value overrides for the Gloo Mesh Helm chart")
	flags.StringVar(&o.releaseName, "release-name", helm.Chart.Data.Name, "Helm release name")
	flags.StringVar(&o.version, "version", "", "Version to install, defaults to latest if omitted")

	flags.BoolVarP(&o.register, "register", "r", false, "Register the cluster running Gloo Mesh")
	flags.StringVar(&o.clusterName, "cluster-name", "mgmt-cluster",
		"Name with which to register the cluster running Gloo Mesh, only applies if --register is also set")
	flags.StringVar(&o.apiServerAddress, "api-server-address", "", "Swap out the address of the remote cluster's k8s API server for the value of this flag. Set this flag when the address of the cluster domain used by the Gloo Mesh is different than that specified in the local kubeconfig.")
	flags.StringVar(&o.clusterDomain, "cluster-domain", "", "The Cluster Domain used by the Kubernetes DNS Service in the registered cluster. Defaults to 'cluster.local'. Read more: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/")
	flags.StringVar(&o.certAgentChartPath, "cert-agent-chart-file", "", "Path to a local Helm chart for installing the Certificate Agent. If unset, this command will install the Certificate Agent from the publicly released Helm chart.")
	flags.StringVar(&o.certAgentValuesPath, "cert-agent-chart-values", "", "Path to a Helm values.yaml file for customizing the installation of the Certificate Agent. If unset, this command will install the Certificate Agent with default Helm values.")
	flags.BoolVarP(&o.verbose, "verbose", "v", false, "Enable verbose output")
}

func install(ctx context.Context, opts *options) error {
	// User-specified chartPath takes precedence over specified version.
	gloomeshChartUri := opts.chartPath
	gloomeshVersion := opts.version
	if opts.version == "" {
		gloomeshVersion = version.Version
	}
	if gloomeshChartUri == "" {
		gloomeshChartUri = fmt.Sprintf(gloomesh.GlooMeshChartUriTemplate, gloomeshVersion)
	}

	err := gloomesh.Installer{
		HelmChartPath:  gloomeshChartUri,
		HelmValuesPath: opts.chartValuesFile,
		KubeConfig:     opts.kubeCfgPath,
		KubeContext:    opts.kubeContext,
		Namespace:      opts.namespace,
		ReleaseName:    opts.releaseName,
		Verbose:        opts.verbose,
		DryRun:         opts.dryRun,
	}.InstallGlooMesh(
		ctx,
	)

	if err != nil {
		return eris.Wrap(err, "installing gloo-mesh")
	}

	if opts.register && !opts.dryRun {
		registrantOpts := &registration.RegistrantOptions{
			KubeConfigPath: opts.kubeCfgPath,
			MgmtContext:    opts.kubeContext,
			RemoteContext:  opts.kubeContext,
			Registration: register.RegistrationOptions{
				ClusterName:      opts.clusterName,
				RemoteCtx:        opts.kubeContext,
				Namespace:        opts.namespace,
				RemoteNamespace:  opts.namespace,
				APIServerAddress: opts.apiServerAddress,
				ClusterDomain:    opts.clusterDomain,
			},
			CertAgent: registration.AgentInstallOptions{
				ChartPath:   opts.certAgentChartPath,
				ChartValues: opts.certAgentValuesPath,
			},
			Verbose: opts.verbose,
		}
		registrant, err := registration.NewRegistrant(registrantOpts)
		if err != nil {
			return eris.Wrap(err, "initializing registrant")
		}
		if err := registrant.RegisterCluster(ctx); err != nil {
			return eris.Wrap(err, "registering management-plane cluster")
		}
	}
	return nil
}
