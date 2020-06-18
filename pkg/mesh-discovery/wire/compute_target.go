package wire

import (
	"github.com/google/wire"
	k8s_apps_providers "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/providers"
	k8s_core_providers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/providers"
	smh_discovery_providers "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/providers"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/aws_creds"
	appmesh2 "github.com/solo-io/service-mesh-hub/pkg/common/aws/clients"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/cloud"
	"github.com/solo-io/service-mesh-hub/pkg/common/aws/matcher"
	aws_utils "github.com/solo-io/service-mesh-hub/pkg/common/aws/parser"
	cluster_registration "github.com/solo-io/service-mesh-hub/pkg/common/cluster-registration"
	compute_target "github.com/solo-io/service-mesh-hub/pkg/common/compute-target"
	mc_manager "github.com/solo-io/service-mesh-hub/pkg/common/compute-target/k8s"
	"github.com/solo-io/service-mesh-hub/pkg/common/container-runtime/docker"
	"github.com/solo-io/service-mesh-hub/pkg/common/container-runtime/version"
	"github.com/solo-io/service-mesh-hub/pkg/common/csr/installation"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/auth"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/helm"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/kubeconfig"
	compute_target_aws "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/compute-target/aws"
	eks_client "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/compute-target/aws/clients/eks"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/k8s-cluster/rest/eks"
	meshworkload_appmesh "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh-workload/k8s/appmesh"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/discovery/mesh/rest/appmesh"
	"k8s.io/client-go/rest"
)

var AwsSet = wire.NewSet(
	compute_target_aws.NewAwsAPIHandler,
	aws_creds.DefaultSecretAwsCredsConverter,
	aws_utils.NewArnParser,
	aws_utils.NewAppMeshScanner,
	aws_utils.NewAwsAccountIdFetcher,
	meshworkload_appmesh.AppMeshWorkloadScannerFactoryProvider,
	smh_discovery_providers.KubernetesClusterClientProvider,
	eks_client.EksConfigBuilderFactoryProvider,
	appmesh2.AppmeshClientFactoryProvider,
	AwsDiscoveryReconcilersProvider,
	appmesh.NewAppMeshDiscoveryReconciler,
	eks.NewEksDiscoveryReconciler,
	matcher.NewAppmeshMatcher,
	cloud.NewAwsCloudStore,
)

var ClusterRegistrationSet = wire.NewSet(
	helm.HelmClientForMemoryConfigFactoryProvider,
	helm.HelmClientForFileConfigFactoryProvider,
	k8s_core_providers.SecretClientFromConfigFactoryProvider,
	k8s_core_providers.NamespaceClientFromConfigFactoryProvider,
	smh_discovery_providers.KubernetesClusterClientFromConfigFactoryProvider,
	k8s_apps_providers.DeploymentClientFromConfigFactoryProvider,
	k8s_core_providers.ServiceAccountClientFromConfigFactoryProvider,
	auth.RbacClientFactoryProvider,
	auth.ClusterAuthorizationFactoryProvider,
	installation.NewCsrAgentInstallerFactory,
	DeployedVersionFinderProvider,
)

func AwsDiscoveryReconcilersProvider(
	appMeshReconciler compute_target_aws.AppMeshDiscoveryReconciler,
	eksReconciler compute_target_aws.EksDiscoveryReconciler,
) []compute_target_aws.RestAPIDiscoveryReconciler {
	return []compute_target_aws.RestAPIDiscoveryReconciler{appMeshReconciler, eksReconciler}
}

func ComputeTargetCredentialsHandlersProvider(
	asyncManagerController *mc_manager.AsyncManagerController,
	awsCredsHandler compute_target_aws.AwsCredsHandler,
) []compute_target.ComputeTargetCredentialsHandler {
	return []compute_target.ComputeTargetCredentialsHandler{
		asyncManagerController,
		awsCredsHandler,
	}
}

func DeployedVersionFinderProvider(
	masterCfg *rest.Config,
	deploymentClientFromConfigFactory k8s_apps_providers.DeploymentClientFromConfigFactory,
	imageNameParser docker.ImageNameParser,
) (version.DeployedVersionFinder, error) {
	deploymentClient, err := deploymentClientFromConfigFactory(masterCfg)
	if err != nil {
		return nil, err
	}
	return version.NewDeployedVersionFinder(deploymentClient, imageNameParser), nil
}

func ClusterRegistrationClientProvider(
	masterCfg *rest.Config,
	secretClientFactory k8s_core_providers.SecretClientFromConfigFactory,
	kubeClusterClient smh_discovery_providers.KubernetesClusterClientFromConfigFactory,
	namespaceClientFactory k8s_core_providers.NamespaceClientFromConfigFactory,
	kubeConverter kubeconfig.Converter,
	csrAgentInstallerFactory installation.CsrAgentInstallerFactory,
) (cluster_registration.ClusterRegistrationClient, error) {
	masterSecretClient, err := secretClientFactory(masterCfg)
	if err != nil {
		return nil, err
	}
	masterKubeClusterClient, err := kubeClusterClient(masterCfg)
	if err != nil {
		return nil, err
	}
	return cluster_registration.NewClusterRegistrationClient(
		masterSecretClient,
		masterKubeClusterClient,
		namespaceClientFactory,
		kubeConverter,
		csrAgentInstallerFactory,
		auth.DefaultClusterAuthClientFromConfig,
	), nil
}
