package selector

import (
	"context"

	core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	discovery_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
)

//go:generate mockgen -source ./interfaces.go -destination mocks/mock_interfaces.go

// Find Service Mesh Hub resources that correspond to k8s-native resources
type ResourceSelector interface {
	// fetch all MeshServices that match the given selector
	GetMeshServicesByServiceSelector(
		ctx context.Context,
		selector *core_types.ServiceSelector,
	) ([]*discovery_v1alpha1.MeshService, error)

	// get the workloads that the given IdentitySelector applies to
	GetMeshWorkloadsByIdentitySelector(
		ctx context.Context,
		identitySelector *core_types.IdentitySelector,
	) ([]*discovery_v1alpha1.MeshWorkload, error)

	// get the workloads that the given WorkloadSelector applies to
	GetMeshWorkloadsByWorkloadSelector(
		ctx context.Context,
		workloadSelector *core_types.WorkloadSelector,
	) ([]*discovery_v1alpha1.MeshWorkload, error)

	// fetch the MeshService backing a k8s Service by the Service's name, namespace, cluster name
	// return error if no MeshService found, or multiple
	GetMeshServiceByRefSelector(
		ctx context.Context,
		kubeServiceName string,
		kubeServiceNamespace string,
		kubeServiceCluster string,
	) (*discovery_v1alpha1.MeshService, error)

	// get the Mesh Workload corresponding to the indicated pod controller (eg deployment)
	GetMeshWorkloadByRefSelector(
		ctx context.Context,
		podControllerName string,
		podControllerNamespace string,
		podControllerCluster string,
	) (*discovery_v1alpha1.MeshWorkload, error)
}
