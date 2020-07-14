package healthcheck

import (
	k8s_core "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/healthcheck/internal"
	healthcheck_types "github.com/solo-io/service-mesh-hub/cli/pkg/tree/check/healthcheck/types"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	k8s_discovery "github.com/solo-io/service-mesh-hub/pkg/common/kube/discovery"
)

var (
	KubernetesAPI = healthcheck_types.Category{
		Name:   "Kubernetes API",
		Weight: 10,
	}
	ManagementPlane = healthcheck_types.Category{
		Name:   "Service Mesh Hub Management Plane",
		Weight: 9,
	}
	ServiceFederation = healthcheck_types.Category{
		Name:   "Service Federation",
		Weight: 8,
	}
)

func DefaultHealthChecksProvider() healthcheck_types.HealthCheckSuite {
	return map[healthcheck_types.Category][]healthcheck_types.HealthCheck{
		KubernetesAPI: {
			internal.NewKubeConnectivityCheck(), // make this one first- doesn't make sense to do the others if we can't talk to the api server
			internal.NewK8sServerVersionCheck(),
		},
		ManagementPlane: {
			internal.NewInstallNamespaceExistenceCheck(),
			internal.NewSmhComponentsHealthCheck(),
		},
		ServiceFederation: {
			internal.NewFederationDecisionCheck(),
		},
	}
}

func ClientsProvider(
	namespaceClient k8s_core.NamespaceClient,
	serverVersionClient k8s_discovery.ServerVersionClient,
	podClient k8s_core.PodClient,
	meshServiceClient smh_discovery.MeshServiceClient,
) healthcheck_types.Clients {
	return healthcheck_types.Clients{
		NamespaceClient:     namespaceClient,
		ServerVersionClient: serverVersionClient,
		PodClient:           podClient,
		MeshServiceClient:   meshServiceClient,
	}
}
