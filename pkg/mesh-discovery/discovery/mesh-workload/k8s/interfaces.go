package k8s

import (
	"context"

	smh_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	k8s_apps_types "k8s.io/api/apps/v1"
	k8s_core_types "k8s.io/api/core/v1"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_mesh_interfaces.go -package mock_mesh_workload

type MeshWorkloadScanners map[smh_core_types.MeshType]MeshWorkloadScanner

type MeshWorkloadDiscovery interface {
	// Ensure that the existing MeshWorkloads match the set of discovered MeshWorkloads,
	// creating, updating, or deleting MeshWorkloads as necessary.
	// TODO: replace client writes with an output snapshot
	DiscoverMeshWorkloads(ctx context.Context, clusterName string) error
}

// get a resource's controller- i.e., in the case of a pod, get its deployment
type OwnerFetcher interface {
	GetDeployment(ctx context.Context, pod *k8s_core_types.Pod) (*k8s_apps_types.Deployment, error)
}

// Scan a pod to see if it represents a mesh workload and if so return a computed MeshWorkload.
type MeshWorkloadScanner interface {
	ScanPod(ctx context.Context, pod *k8s_core_types.Pod, clusterName string) (*smh_discovery.MeshWorkload, error)
}

// Factory for cluster-scoped MeshWorkloadScanner
type MeshWorkloadScannerFactory func(clusterName string) (MeshWorkloadScanner, error)
