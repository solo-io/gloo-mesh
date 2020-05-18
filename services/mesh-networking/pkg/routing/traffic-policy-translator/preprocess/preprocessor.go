package preprocess

import (
	"context"

	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/selector"
)

type trafficPolicyPreprocessor struct {
	resourceSelector selector.ResourceSelector
	merger           TrafficPolicyMerger
	validator        TrafficPolicyValidator
}

func NewTrafficPolicyPreprocessor(
	resourceSelector selector.ResourceSelector,
	merger TrafficPolicyMerger,
	validator TrafficPolicyValidator,
) TrafficPolicyPreprocessor {
	return &trafficPolicyPreprocessor{
		resourceSelector: resourceSelector,
		merger:           merger,
		validator:        validator,
	}
}

/*
	Given a TrafficPolicy, do the following:
	1. Fetch all destination MeshServices
	2. For each MeshService, do the following:
		a. Fetch existing TrafficPolicies that apply to it
		b. Sort the TrafficPolicies by creation time ascending
		c. Merge the TrafficPolicies
			- If conflict encountered on any TrafficPolicy, return conflict error
*/
func (d *trafficPolicyPreprocessor) PreprocessTrafficPolicy(
	ctx context.Context,
	trafficPolicy *zephyr_networking.TrafficPolicy,
) (map[selector.MeshServiceId][]*zephyr_networking.TrafficPolicy, error) {
	err := d.validator.Validate(ctx, trafficPolicy)
	if err != nil {
		return nil, err
	}
	meshServices, err := d.resourceSelector.GetAllMeshServicesByServiceSelector(
		ctx,
		trafficPolicy.Spec.GetDestinationSelector(),
	)
	if err != nil {
		return nil, err
	}
	return d.merger.MergeTrafficPoliciesForMeshServices(ctx, meshServices)
}

/*
	Given a MeshService, do the following:
	1. Fetch existing TrafficPolicies that apply to it
	2. Sort the TrafficPolicies by creation time ascending
	3. Merge the TrafficPolicies
		- If conflict encountered on any TrafficPolicy, do not apply any of its rules and update its status to CONFLICT
*/
func (d *trafficPolicyPreprocessor) PreprocessTrafficPoliciesForMeshService(
	ctx context.Context,
	meshService *zephyr_discovery.MeshService,
) (map[selector.MeshServiceId][]*zephyr_networking.TrafficPolicy, error) {
	return d.merger.MergeTrafficPoliciesForMeshServices(ctx, []*zephyr_discovery.MeshService{meshService})
}
