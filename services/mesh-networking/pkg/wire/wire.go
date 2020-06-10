// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	kubernetes_apps "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/apps/v1"
	kubernetes_core "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	smh_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/csr/certgen"
	"github.com/solo-io/service-mesh-hub/pkg/filesystem/files"
	"github.com/solo-io/service-mesh-hub/pkg/kube/kubeconfig"
	multicluster_wire "github.com/solo-io/service-mesh-hub/services/common/compute-target/wire"
	csr_generator "github.com/solo-io/service-mesh-hub/services/csr-agent/pkg/csr-generator"
	networking_multicluster "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/compute-target"
	controller_factories "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/compute-target/controllers"
	cert_manager "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/security/cert-manager"
	cert_signer "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/security/cert-signer"
	vm_validation "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/validation"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func MeshServiceReaderProvider(client client.Client) smh_discovery.MeshServiceReader {
	return smh_discovery.MeshServiceClientProvider(client)
}
func MeshWorkloadReaderProvider(client client.Client) smh_discovery.MeshWorkloadReader {
	return smh_discovery.MeshWorkloadClientProvider(client)
}

func InitializeMeshNetworking(ctx context.Context) (MeshNetworkingContext, error) {
	wire.Build(
		kubernetes_core.SecretClientProvider,
		kubernetes_core.ConfigMapClientProvider,
		kubernetes_core.PodClientFactoryProvider,
		kubernetes_core.NodeClientFactoryProvider,
		kubernetes_apps.DeploymentClientFactoryProvider,
		smh_discovery.MeshClientProvider,
		MeshServiceReaderProvider,
		smh_discovery.MeshServiceClientProvider,
		smh_discovery.MeshWorkloadClientProvider,
		MeshWorkloadReaderProvider,
		smh_networking.VirtualMeshClientProvider,
		smh_networking.TrafficPolicyClientProvider,
		smh_networking.AccessControlPolicyClientProvider,
		csr_generator.NewVirtualMeshCSRDataSourceFactory,
		vm_validation.NewVirtualMeshFinder,
		cert_signer.NewVirtualMeshCertClient,
		multicluster_wire.MulticlusterProviderSet,
		multicluster_wire.DynamicClientGetterProvider,
		certgen.NewSigner,
		certgen.NewRootCertGenerator,
		LocalMeshWorkloadEventWatcherProvider,
		LocalMeshServiceEventWatcherProvider,
		ClientFactoryProviderSet,
		ControllerFactoryProviderSet,
		TrafficPolicyProviderSet,
		AccessControlPolicySet,
		FederationProviderSet,
		networking_multicluster.NewMeshNetworkingClusterHandler,
		controller_factories.NewLocalVirtualMeshEventWatcher,
		vm_validation.NewVirtualMeshValidator,
		cert_manager.VMCSRSnapshotListenerSet,
		MeshNetworkingSnapshotContextProvider,
		MeshNetworkingContextProvider,
		AwsSet,
		ComputeTargetCredentialsHandlersProvider,
		kubeconfig.NewConverter,
		files.NewDefaultFileReader,
	)

	return MeshNetworkingContext{}, nil
}
