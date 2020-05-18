package traffic_policy_validation

import (
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go

// do data validation on a Traffic Policy; e.g., ensure that its retry attempts are non-negative, that it references real services, etc.
type Validator interface {
	// always returns a non-nil Status that should be written to the cluster
	// if validation failed, the concrete validation error that occurred will be returned with that non-nil status so it can be logged
	ValidateTrafficPolicy(trafficPolicy *zephyr_networking.TrafficPolicy, allMeshServices []*zephyr_discovery.MeshService) (*zephyr_core_types.Status, error)
}
