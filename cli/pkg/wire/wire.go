// +build wireinject

package wire

import (
	"context"
	"io"

	"github.com/google/wire"
	"github.com/solo-io/go-utils/installutils/helminstall"
	usageclient "github.com/solo-io/reporting-client/pkg/client"
	cli "github.com/solo-io/service-mesh-hub/cli/pkg"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	common_config "github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/interactive"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/kube"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/resource_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/usage"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/healthcheck"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/status"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/deregister"
	register "github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/register/csr"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/demo"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/describe"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/describe/description"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/install"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh/install/istio/operator"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/config_lookup"
	crd_uninstall "github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/crd"
	helm_uninstall "github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/helm"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/upgrade"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/version"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/version/server"
	discovery_versioned "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/clientset/versioned"
	networking_versioned "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/clientset/versioned"
	security_versioned "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1/clientset/versioned"
	"github.com/solo-io/service-mesh-hub/pkg/auth"
	apiext2 "github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/apiext"
	kubernetes_apps "github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/apps"
	kubernetes_core "github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/core"
	kubernetes_discovery "github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/discovery"
	discovery_core "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/discovery"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/networking"
	zephyr_security "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/security"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	"github.com/solo-io/service-mesh-hub/pkg/kubeconfig"
	"github.com/solo-io/service-mesh-hub/pkg/selector"
	version2 "github.com/solo-io/service-mesh-hub/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func DefaultKubeClientsFactory(masterConfig *rest.Config, writeNamespace string) (clients *common.KubeClients, err error) {
	wire.Build(
		kubernetes.NewForConfig,
		wire.Bind(new(kubernetes.Interface), new(*kubernetes.Clientset)),
		discovery_versioned.NewForConfig,
		wire.Bind(new(discovery_versioned.Interface), new(*discovery_versioned.Clientset)),
		networking_versioned.NewForConfig,
		wire.Bind(new(networking_versioned.Interface), new(*networking_versioned.Clientset)),
		security_versioned.NewForConfig,
		wire.Bind(new(security_versioned.Interface), new(*security_versioned.Clientset)),
		kubernetes_core.NewGeneratedServiceAccountClient,
		discovery_core.NewGeneratedKubernetesClusterClient,
		kubernetes_core.NewGeneratedSecretsClient,
		kubernetes_core.NewGeneratedNamespaceClient,
		kubernetes_discovery.NewGeneratedServerVersionClient,
		kubernetes_core.NewGeneratedPodClient,
		discovery_core.NewGeneratedMeshClient,
		discovery_core.NewGeneratedMeshServiceClient,
		zephyr_networking.NewGeneratedVirtualMeshClient,
		zephyr_security.NewGeneratedVirtualMeshCSRClient,
		kubernetes_apps.NewGeneratedDeploymentClient,
		zephyr_networking.NewGeneratedAccessControlPolicyClient,
		zephyr_networking.NewGeneratedTrafficPolicyClient,
		discovery_core.NewGeneratedMeshWorkloadClient,
		apiext2.NewGeneratedCrdClientFactory,
		auth.NewRemoteAuthorityConfigCreator,
		auth.RbacClientProvider,
		auth.NewRemoteAuthorityManager,
		auth.NewClusterAuthorization,
		docker.NewImageNameParser,
		version2.NewDeployedVersionFinder,
		helminstall.DefaultHelmClient,
		install.HelmInstallerProvider,
		healthcheck.ClientsProvider,
		crd_uninstall.NewCrdRemover,
		kubeconfig.SecretToConfigConverterProvider,
		common.UninstallClientsProvider,
		common_config.NewInMemoryRESTClientGetterFactory,
		helm_uninstall.NewUninstallerFactory,
		config_lookup.NewKubeConfigLookup,
		deregister.NewClusterDeregistrationClient,
		common.KubeClientsProvider,
		description.NewResourceDescriber,
		selector.NewResourceSelector,
		kubernetes_apps.GeneratedDeploymentClientFactoryProvider,
	)
	return nil, nil
}

func DefaultClientsFactory(opts *options.Options) (*common.Clients, error) {
	wire.Build(
		upgrade.UpgraderClientSet,
		docker.NewImageNameParser,
		common_config.DefaultKubeLoaderProvider,
		common_config.NewMasterKubeConfigVerifier,
		server.DefaultServerVersionClientProvider,
		kube.NewUnstructuredKubeClientFactory,
		server.NewDeploymentClient,
		operator.NewInstallerManifestBuilder,
		common.IstioClientsProvider,
		register.NewCsrAgentInstallerFactory,
		common.ClusterRegistrationClientsProvider,
		operator.NewOperatorManagerFactory,
		status.StatusClientFactoryProvider,
		healthcheck.DefaultHealthChecksProvider,
		common.ClientsProvider,
	)
	return nil, nil
}

func InitializeCLI(ctx context.Context, out io.Writer, in io.Reader) *cobra.Command {
	wire.Build(
		docker.NewImageNameParser,
		common.NewDefaultFileReader,
		common_config.DefaultKubeLoaderProvider,
		options.NewOptionsProvider,
		DefaultKubeClientsFactoryProvider,
		DefaultClientsFactoryProvider,
		table_printing.TablePrintingSet,
		resource_printing.NewResourcePrinter,
		common.PrintersProvider,
		usage.DefaultUsageReporterProvider,
		interactive.NewSurveyInteractivePrompt,
		demo.NewCommandLineRunner,
		demo.DemoSet,
		upgrade.UpgradeSet,
		cluster.ClusterSet,
		version.VersionSet,
		mesh.MeshProviderSet,
		install.InstallSet,
		uninstall.UninstallSet,
		check.CheckSet,
		describe.DescribeSet,
		create.CreateSet,
		get.GetSet,
		cli.BuildCli,
	)
	return nil
}

func InitializeCLIWithMocks(
	ctx context.Context,
	out io.Writer,
	in io.Reader,
	usageClient usageclient.Client,
	kubeClientsFactory common.KubeClientsFactory,
	clientsFactory common.ClientsFactory,
	kubeLoader common_config.KubeLoader,
	imageNameParser docker.ImageNameParser,
	fileReader common.FileReader,
	secretToConfigConverter kubeconfig.SecretToConfigConverter,
	printers common.Printers,
	runnner demo.CommandLineRunner,
	interactivePrompt interactive.InteractivePrompt,
) *cobra.Command {
	wire.Build(
		options.NewOptionsProvider,
		demo.DemoSet,
		cluster.ClusterSet,
		version.VersionSet,
		mesh.MeshProviderSet,
		install.InstallSet,
		upgrade.UpgradeSet,
		uninstall.UninstallSet,
		check.CheckSet,
		get.GetSet,
		describe.DescribeSet,
		create.CreateSet,
		cli.BuildCli,
	)
	return nil
}
