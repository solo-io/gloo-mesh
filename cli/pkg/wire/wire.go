// +build wireinject

package wire

import (
	"context"
	"io"

	kubernetes_apiext_providers "github.com/solo-io/external-apis/pkg/api/k8s/apiextensions.k8s.io/v1beta1/providers"
	kubernetes_apps "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	kubernetes_apps_providers "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/providers"
	k8s_core "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	k8s_core_providers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/providers"
	smh_discovery_providers "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/providers"
	smh_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	smh_networking_providers "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/providers"
	smh_security "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"
	smh_security_providers "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/providers"

	"github.com/google/wire"
	usageclient "github.com/solo-io/reporting-client/pkg/client"
	cli "github.com/solo-io/service-mesh-hub/cli/pkg"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	common_config "github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/exec"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/interactive"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/resource_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/usage"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/healthcheck"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/status"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/demo"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/describe"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/describe/description"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/get"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/install"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/config_lookup"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/upgrade"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/version"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/version/server"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	cluster_registration "github.com/solo-io/service-mesh-hub/pkg/common/cluster-registration"
	"github.com/solo-io/service-mesh-hub/pkg/common/container-runtime/docker"
	version2 "github.com/solo-io/service-mesh-hub/pkg/common/container-runtime/version"
	"github.com/solo-io/service-mesh-hub/pkg/common/csr/installation"
	"github.com/solo-io/service-mesh-hub/pkg/common/filesystem/files"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/auth"
	crd_uninstall "github.com/solo-io/service-mesh-hub/pkg/common/kube/crd"
	kubernetes_discovery "github.com/solo-io/service-mesh-hub/pkg/common/kube/discovery"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/helm"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/kubeconfig"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/selection"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/unstructured"
	"github.com/solo-io/service-mesh-hub/pkg/common/mesh-installation/istio/operator"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func MeshServiceReaderProvider(clients smh_discovery.Clientset) smh_discovery.MeshServiceReader {
	return smh_discovery_providers.MeshServiceClientFromClientsetProvider(clients)
}
func MeshWorkloadReaderProvider(clients smh_discovery.Clientset) smh_discovery.MeshWorkloadReader {
	return smh_discovery_providers.MeshWorkloadClientFromClientsetProvider(clients)
}

func DefaultKubeClientsFactory(masterConfig *rest.Config, writeNamespace string) (clients *common.KubeClients, err error) {
	wire.Build(
		kubernetes.NewForConfig,
		wire.Bind(new(kubernetes.Interface), new(*kubernetes.Clientset)),
		kubernetes_discovery.NewGeneratedServerVersionClient,
		k8s_core.NewClientsetFromConfig,
		k8s_core_providers.ServiceAccountClientFromClientsetProvider,
		k8s_core_providers.SecretClientFromClientsetProvider,
		k8s_core_providers.NamespaceClientFromClientsetProvider,
		k8s_core_providers.PodClientFromClientsetProvider,
		k8s_core_providers.SecretClientFactoryProvider,
		k8s_core_providers.ServiceAccountClientFactoryProvider,
		k8s_core_providers.NamespaceClientFromConfigFactoryProvider,
		kubernetes_apps.NewClientsetFromConfig,
		kubernetes_apps_providers.DeploymentClientFromClientsetProvider,
		kubernetes_apps_providers.DeploymentClientFactoryProvider,
		kubernetes_apiext_providers.CustomResourceDefinitionClientFromConfigFactoryProvider,
		files.NewDefaultFileReader,
		auth.NewRemoteAuthorityConfigCreator,
		auth.RbacClientProvider,
		auth.NewRemoteAuthorityManager,
		auth.NewClusterAuthorization,
		docker.NewImageNameParser,
		version2.NewDeployedVersionFinder,
		install.HelmInstallerProvider,
		healthcheck.ClientsProvider,
		crd_uninstall.NewCrdRemover,
		kubeconfig.NewConverter,
		common.UninstallClientsProvider,
		kubeconfig.NewKubeConfigLookup,
		config_lookup.NewDynamicClientGetter,
		cluster_registration.NewClusterDeregistrationClient,
		common.KubeClientsProvider,
		description.NewResourceDescriber,
		selection.NewResourceSelector,
		smh_discovery.NewClientsetFromConfig,
		smh_networking.NewClientsetFromConfig,
		smh_security.NewClientsetFromConfig,
		smh_discovery_providers.KubernetesClusterClientFromClientsetProvider,
		smh_discovery_providers.MeshServiceClientFromClientsetProvider,
		MeshServiceReaderProvider,
		smh_discovery_providers.MeshWorkloadClientFromClientsetProvider,
		MeshWorkloadReaderProvider,
		smh_discovery_providers.MeshClientFromClientsetProvider,
		smh_networking_providers.TrafficPolicyClientFromClientsetProvider,
		smh_networking_providers.AccessControlPolicyClientFromClientsetProvider,
		smh_networking_providers.VirtualMeshClientFromClientsetProvider,
		smh_security_providers.VirtualMeshCertificateSigningRequestClientFromClientsetProvider,
		installation.NewCsrAgentInstallerFactory,
		helm.HelmClientForMemoryConfigFactoryProvider,
		helm.HelmClientForFileConfigFactoryProvider,
		cluster_registration.NewClusterRegistrationClient,
		auth.ClusterAuthClientFromConfigFactoryProvider,
	)
	return nil, nil
}

func DefaultClientsFactory(opts *options.Options) (*common.Clients, error) {
	wire.Build(
		files.NewDefaultFileReader,
		kubeconfig.NewConverter,
		operator.NewInstallerManifestBuilder,
		operator.NewOperatorDaoFactory,
		operator.NewOperatorManagerFactory,
		upgrade.UpgraderClientSet,
		docker.NewImageNameParser,
		kubeconfig.NewKubeLoader,
		common_config.NewMasterKubeConfigVerifier,
		server.DefaultServerVersionClientProvider,
		unstructured.NewUnstructuredKubeClientFactory,
		server.NewDeploymentClient,
		common.IstioClientsProvider,
		status.StatusClientFactoryProvider,
		healthcheck.DefaultHealthChecksProvider,
		common.ClientsProvider,
	)
	return nil, nil
}

func InitializeCLI(ctx context.Context, out io.Writer, in io.Reader) *cobra.Command {
	wire.Build(
		docker.NewImageNameParser,
		files.NewDefaultFileReader,
		kubeconfig.NewKubeLoader,
		options.NewOptionsProvider,
		DefaultKubeClientsFactoryProvider,
		DefaultClientsFactoryProvider,
		table_printing.TablePrintingSet,
		resource_printing.NewResourcePrinter,
		common.PrintersProvider,
		usage.DefaultUsageReporterProvider,
		interactive.NewSurveyInteractivePrompt,
		exec.NewShellRunner,
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
		afero.NewOsFs,
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
	kubeLoader kubeconfig.KubeLoader,
	imageNameParser docker.ImageNameParser,
	fileReader files.FileReader,
	kubeconfigConverter kubeconfig.Converter,
	printers common.Printers,
	runner exec.Runner,
	interactivePrompt interactive.InteractivePrompt,
	fs afero.Fs,
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
