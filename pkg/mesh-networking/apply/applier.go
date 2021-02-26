package apply

import (
	"context"
	"sort"

	commonv1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discoveryv1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	networkingv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/apply/configtarget"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/hostutils"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/selectorutils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	utilsets "k8s.io/apimachinery/pkg/util/sets"
)

// the Applier validates user-applied configuration
// and produces a snapshot that is ready for translation (i.e. with accepted policies applied to all the Status of all targeted Destinations)
// Note that the Applier also updates the statuses of objects contained in the input Snapshot.
// The Input Snapshot's SyncStatuses method should usually be called after running the Applier.
type Applier interface {
	Apply(ctx context.Context, input input.LocalSnapshot, userSupplied input.RemoteSnapshot)
}

type applier struct {
	// the applier runs the networking translator in order to detect & report translation errors
	translator translation.Translator
}

func NewApplier(
	translator translation.Translator,
) Applier {
	return &applier{
		translator: translator,
	}
}

func (v *applier) Apply(ctx context.Context, input input.LocalSnapshot, userSupplied input.RemoteSnapshot) {
	ctx = contextutils.WithLogger(ctx, "validation")
	reporter := newApplyReporter()

	initializePolicyStatuses(input)

	setDiscoveryStatusMetadata(input)

	validateConfigTargetReferences(input)

	applyPoliciesToConfigTargets(input)

	// perform a dry run of translation to find any errors
	_, err := v.translator.Translate(ctx, input, userSupplied, reporter)
	if err != nil {
		// should never happen
		contextutils.LoggerFrom(ctx).DPanicf("internal error: failed to run translator: %v", err)
	}

	reportTranslationErrors(ctx, reporter, input)
}

// Optimistically initialize policy statuses to accepted, which may be set to invalid or failed pending subsequent validation.
func initializePolicyStatuses(input input.LocalSnapshot) {
	trafficPolicies := input.TrafficPolicies().List()
	accessPolicies := input.AccessPolicies().List()
	virtualMeshes := input.VirtualMeshes().List()

	// initialize TrafficPolicy statuses
	for _, trafficPolicy := range trafficPolicies {
		trafficPolicy.Status = networkingv1.TrafficPolicyStatus{
			State:              commonv1.ApprovalState_ACCEPTED,
			ObservedGeneration: trafficPolicy.Generation,
			Destinations:       map[string]*networkingv1.ApprovalStatus{},
		}
	}

	// initialize AccessPolicy statuses
	for _, accessPolicy := range accessPolicies {
		accessPolicy.Status = networkingv1.AccessPolicyStatus{
			State:              commonv1.ApprovalState_ACCEPTED,
			ObservedGeneration: accessPolicy.Generation,
			Destinations:       map[string]*networkingv1.ApprovalStatus{},
		}
	}

	// By this point, VirtualMeshes have already undergone pre-translation validation.
	for _, virtualMesh := range virtualMeshes {
		virtualMesh.Status = networkingv1.VirtualMeshStatus{
			State:              commonv1.ApprovalState_ACCEPTED,
			ObservedGeneration: virtualMesh.Generation,
			Meshes:             map[string]*networkingv1.ApprovalStatus{},
		}
	}
}

// Append status metadata to relevant discovery resources.
func setDiscoveryStatusMetadata(input input.LocalSnapshot) {
	clusterDomains := hostutils.NewClusterDomainRegistry(input.KubernetesClusters(), input.Destinations())
	for _, destination := range input.Destinations().List() {
		if destination.Spec.GetKubeService() != nil {
			ref := destination.Spec.GetKubeService().GetRef()
			destination.Status.LocalFqdn = clusterDomains.GetLocalFQDN(ref)
		}
	}
}

// Validate that configuration target references.
func validateConfigTargetReferences(input input.LocalSnapshot) {
	configTargetValidator := configtarget.NewConfigTargetValidator(input.Meshes(), input.Destinations())
	configTargetValidator.ValidateAccessPolicies(input.AccessPolicies().List())
	configTargetValidator.ValidateTrafficPolicies(input.TrafficPolicies().List())
	configTargetValidator.ValidateVirtualMeshes(input.VirtualMeshes().List())
}

// Apply networking configuration policies to relevant discovery entities.
func applyPoliciesToConfigTargets(input input.LocalSnapshot) {
	for _, destination := range input.Destinations().List() {
		destination.Status.AppliedTrafficPolicies = getAppliedTrafficPolicies(input.TrafficPolicies().List(), destination)
		destination.Status.AppliedAccessPolicies = getAppliedAccessPolicies(input.AccessPolicies().List(), destination)
	}

	for _, mesh := range input.Meshes().List() {
		mesh.Status.AppliedVirtualMesh = getAppliedVirtualMesh(input.VirtualMeshes().List(), mesh)
	}
}

// For all discovery entities, update status with any translation errors.
// Also update observed generation to indicate that it's been processed.
func reportTranslationErrors(ctx context.Context, reporter *applyReporter, input input.LocalSnapshot) {
	for _, workload := range input.Workloads().List() {
		// TODO: validate config applied to workloads when introduced
		workload.Status.ObservedGeneration = workload.Generation
	}

	for _, destination := range input.Destinations().List() {
		destination.Status.ObservedGeneration = destination.Generation
		destination.Status.AppliedTrafficPolicies = validateAndReturnApprovedTrafficPolicies(ctx, input, reporter, destination)
		destination.Status.AppliedAccessPolicies = validateAndReturnApprovedAccessPolicies(ctx, input, reporter, destination)
	}

	for _, mesh := range input.Meshes().List() {
		mesh.Status.ObservedGeneration = mesh.Generation
		mesh.Status.AppliedVirtualMesh = validateAndReturnVirtualMesh(ctx, input, reporter, mesh)
	}

	setWorkloadsForTrafficPolicies(ctx, input.TrafficPolicies().List(), input.Workloads().List(), input.Destinations(), input.Meshes())
	setWorkloadsForAccessPolicies(ctx, input.AccessPolicies().List(), input.Workloads().List(), input.Destinations(), input.Meshes())
}

// A workload is associated with a TrafficPolicy if the workload matches the policy's workload selector
// AND the workload is in the same mesh or VirtualMesh as any of the policy's selected Destinations
func setWorkloadsForTrafficPolicies(
	ctx context.Context,
	trafficPolicies networkingv1.TrafficPolicySlice,
	workloads discoveryv1.WorkloadSlice,
	destinations discoveryv1sets.DestinationSet,
	meshes discoveryv1sets.MeshSet) {

	// create a map of mesh to VirtualMesh for lookup
	meshToVirtualMesh := makeMeshToVirtualMeshMap(meshes.List())

	for _, trafficPolicy := range trafficPolicies {
		// get the selected Destinations on the policy
		matchingDestinations := destinations.List(func(destination *discoveryv1.Destination) bool {
			return trafficPolicy.Status.GetDestinations()[sets.Key(destination.GetObjectMeta())] == nil
		})
		// get all the mesh and VirtualMesh refs from those Destinations
		matchingMeshes, matchingVirtualMeshes := getMeshesFromDestinations(ctx, matchingDestinations, meshes)

		var matchingWorkloads []string
		// TODO(awang) optimize if the returned workloads list gets too large
		//if len(trafficPolicy.Spec.GetSourceSelector()) == 0 {
		//	trafficPolicy.Status.Workloads = []string{"*"}
		//	return
		//}
		for _, workload := range workloads {
			if selectorutils.SelectorMatchesWorkload(trafficPolicy.Spec.GetSourceSelector(), workload) &&
				meshMatches(workload.Spec.GetMesh(), matchingMeshes, matchingVirtualMeshes, meshToVirtualMesh) {
				matchingWorkloads = append(matchingWorkloads, sets.Key(workload))
			}
		}
		trafficPolicy.Status.Workloads = matchingWorkloads
	}
}

// A workload is associated with an AccessPolicy if the workload matches the policy's identity selector
// AND the workload is in the same mesh or VirtualMesh as any of the policy's selected Destinations
func setWorkloadsForAccessPolicies(
	ctx context.Context,
	accessPolicies networkingv1.AccessPolicySlice,
	workloads discoveryv1.WorkloadSlice,
	destinations discoveryv1sets.DestinationSet,
	meshes discoveryv1sets.MeshSet) {

	// create a map of mesh to VirtualMesh for lookup
	meshToVirtualMesh := makeMeshToVirtualMeshMap(meshes.List())

	for _, accessPolicy := range accessPolicies {
		// get the selected Destinations on the policy
		matchingDestinations := destinations.List(func(destination *discoveryv1.Destination) bool {
			return accessPolicy.Status.GetDestinations()[sets.Key(destination.GetObjectMeta())] == nil
		})
		// get all the mesh and VirtualMesh refs from those Destinations
		matchingMeshes, matchingVirtualMeshes := getMeshesFromDestinations(ctx, matchingDestinations, meshes)

		var matchingWorkloads []string
		// TODO(awang) optimize if the returned workloads list gets too large
		for _, workload := range workloads {
			if selectorutils.IdentityMatchesWorkload(accessPolicy.Spec.GetSourceSelector(), workload) &&
				meshMatches(workload.Spec.GetMesh(), matchingMeshes, matchingVirtualMeshes, meshToVirtualMesh) {
				matchingWorkloads = append(matchingWorkloads, sets.Key(workload))
			}
		}
		accessPolicy.Status.Workloads = matchingWorkloads
	}
}

// this function both validates the status of TrafficPolicies (sets error or accepted state)
// as well as returns a list of accepted traffic policies for the Destination status
func validateAndReturnApprovedTrafficPolicies(ctx context.Context, input input.LocalSnapshot, reporter *applyReporter, destination *discoveryv1.Destination) []*discoveryv1.DestinationStatus_AppliedTrafficPolicy {
	var validatedTrafficPolicies []*discoveryv1.DestinationStatus_AppliedTrafficPolicy

	// track accepted index
	var acceptedIndex uint32
	for _, appliedTrafficPolicy := range destination.Status.AppliedTrafficPolicies {
		errsForTrafficPolicy := reporter.getTrafficPolicyErrors(destination, appliedTrafficPolicy.Ref)

		trafficPolicy, err := input.TrafficPolicies().Find(appliedTrafficPolicy.Ref)
		if err != nil {
			// should never happen
			contextutils.LoggerFrom(ctx).Errorf("internal error: failed to look up applied TrafficPolicy %v: %v", appliedTrafficPolicy.Ref, err)
			continue
		}

		if trafficPolicy.Status.Destinations == nil {
			trafficPolicy.Status.Destinations = map[string]*networkingv1.ApprovalStatus{}
		}

		if len(errsForTrafficPolicy) == 0 {
			trafficPolicy.Status.Destinations[sets.Key(destination)] = &networkingv1.ApprovalStatus{
				AcceptanceOrder: acceptedIndex,
				State:           commonv1.ApprovalState_ACCEPTED,
			}
			validatedTrafficPolicies = append(validatedTrafficPolicies, appliedTrafficPolicy)
			acceptedIndex++
		} else {
			var errMsgs []string
			for _, tpErr := range errsForTrafficPolicy {
				errMsgs = append(errMsgs, tpErr.Error())
			}
			trafficPolicy.Status.Destinations[sets.Key(destination)] = &networkingv1.ApprovalStatus{
				State:  commonv1.ApprovalState_INVALID,
				Errors: errMsgs,
			}
			trafficPolicy.Status.State = commonv1.ApprovalState_INVALID
		}
	}

	return validatedTrafficPolicies
}

// this function both validates the status of AccessPolicies (sets error or accepted state)
// as well as returns a list of accepted AccessPolicies for the Destination status
func validateAndReturnApprovedAccessPolicies(
	ctx context.Context,
	input input.LocalSnapshot,
	reporter *applyReporter,
	destination *discoveryv1.Destination,
) []*discoveryv1.DestinationStatus_AppliedAccessPolicy {
	var validatedAccessPolicies []*discoveryv1.DestinationStatus_AppliedAccessPolicy

	// track accepted index
	var acceptedIndex uint32
	for _, appliedAccessPolicy := range destination.Status.AppliedAccessPolicies {
		errsForAccessPolicy := reporter.getAccessPolicyErrors(destination, appliedAccessPolicy.Ref)

		accessPolicy, err := input.AccessPolicies().Find(appliedAccessPolicy.Ref)
		if err != nil {
			// should never happen
			contextutils.LoggerFrom(ctx).Errorf("internal error: failed to look up applied AccessPolicy %v: %v", appliedAccessPolicy.Ref, err)
			continue
		}

		if len(errsForAccessPolicy) == 0 {
			accessPolicy.Status.Destinations[sets.Key(destination)] = &networkingv1.ApprovalStatus{
				AcceptanceOrder: acceptedIndex,
				State:           commonv1.ApprovalState_ACCEPTED,
			}
			validatedAccessPolicies = append(validatedAccessPolicies, appliedAccessPolicy)
			acceptedIndex++
		} else {
			var errMsgs []string
			for _, apErr := range errsForAccessPolicy {
				errMsgs = append(errMsgs, apErr.Error())
			}
			accessPolicy.Status.Destinations[sets.Key(destination)] = &networkingv1.ApprovalStatus{
				State:  commonv1.ApprovalState_INVALID,
				Errors: errMsgs,
			}
			accessPolicy.Status.State = commonv1.ApprovalState_INVALID
		}
	}

	return validatedAccessPolicies
}

func validateAndReturnVirtualMesh(
	ctx context.Context,
	input input.LocalSnapshot,
	reporter *applyReporter,
	mesh *discoveryv1.Mesh,
) *discoveryv1.MeshStatus_AppliedVirtualMesh {
	appliedVirtualMesh := mesh.Status.AppliedVirtualMesh
	if appliedVirtualMesh == nil {
		return nil
	}
	errsForVirtualMesh := reporter.getVirtualMeshErrors(mesh, appliedVirtualMesh.Ref)

	virtualMesh, err := input.VirtualMeshes().Find(appliedVirtualMesh.Ref)
	if err != nil {
		// should never happen
		contextutils.LoggerFrom(ctx).Errorf("internal error: failed to look up applied VirtualMesh %v: %v", appliedVirtualMesh.Ref, err)
		return nil
	}

	if len(errsForVirtualMesh) == 0 {
		virtualMesh.Status.Meshes[sets.Key(mesh)] = &networkingv1.ApprovalStatus{
			State: commonv1.ApprovalState_ACCEPTED,
		}
		return appliedVirtualMesh
	} else {
		var errMsgs []string
		for _, fsErr := range errsForVirtualMesh {
			errMsgs = append(errMsgs, fsErr.Error())
		}
		virtualMesh.Status.Meshes[sets.Key(mesh)] = &networkingv1.ApprovalStatus{
			State:  commonv1.ApprovalState_INVALID,
			Errors: errMsgs,
		}
		virtualMesh.Status.State = commonv1.ApprovalState_INVALID
		return nil
	}
}

// the applyReporter validates individual policies and reports any encountered errors
type applyReporter struct {
	// NOTE(ilackarms): map access should be synchronous (called in a single context),
	// so locking should not be necessary.
	unappliedTrafficPolicies  map[*discoveryv1.Destination]map[string][]error
	unappliedAccessPolicies   map[*discoveryv1.Destination]map[string][]error
	unappliedFailoverServices map[*discoveryv1.Mesh]map[string][]error
	unappliedVirtualMeshes    map[*discoveryv1.Mesh]map[string][]error
	invalidFailoverServices   map[string][]error
}

func newApplyReporter() *applyReporter {
	return &applyReporter{
		unappliedTrafficPolicies:  map[*discoveryv1.Destination]map[string][]error{},
		unappliedAccessPolicies:   map[*discoveryv1.Destination]map[string][]error{},
		unappliedFailoverServices: map[*discoveryv1.Mesh]map[string][]error{},
		unappliedVirtualMeshes:    map[*discoveryv1.Mesh]map[string][]error{},
		invalidFailoverServices:   map[string][]error{},
	}
}

// mark the policy with an error; will be used to filter the policy out of
// the accepted status later
func (v *applyReporter) ReportTrafficPolicyToDestination(destination *discoveryv1.Destination, trafficPolicy ezkube.ResourceId, err error) {
	invalidTrafficPoliciesForDestination := v.unappliedTrafficPolicies[destination]
	if invalidTrafficPoliciesForDestination == nil {
		invalidTrafficPoliciesForDestination = map[string][]error{}
	}
	key := sets.Key(trafficPolicy)
	errs := invalidTrafficPoliciesForDestination[key]
	errs = append(errs, err)
	invalidTrafficPoliciesForDestination[key] = errs
	v.unappliedTrafficPolicies[destination] = invalidTrafficPoliciesForDestination
}

func (v *applyReporter) ReportAccessPolicyToDestination(destination *discoveryv1.Destination, accessPolicy ezkube.ResourceId, err error) {
	invalidAccessPoliciesForDestination := v.unappliedAccessPolicies[destination]
	if invalidAccessPoliciesForDestination == nil {
		invalidAccessPoliciesForDestination = map[string][]error{}
	}
	key := sets.Key(accessPolicy)
	errs := invalidAccessPoliciesForDestination[key]
	errs = append(errs, err)
	invalidAccessPoliciesForDestination[key] = errs
	v.unappliedAccessPolicies[destination] = invalidAccessPoliciesForDestination
}

func (v *applyReporter) ReportVirtualMeshToMesh(mesh *discoveryv1.Mesh, virtualMesh ezkube.ResourceId, err error) {
	invalidVirtualMeshesForMesh := v.unappliedVirtualMeshes[mesh]
	if invalidVirtualMeshesForMesh == nil {
		invalidVirtualMeshesForMesh = map[string][]error{}
	}
	key := sets.Key(virtualMesh)
	errs := invalidVirtualMeshesForMesh[key]
	errs = append(errs, err)
	invalidVirtualMeshesForMesh[key] = errs
	v.unappliedVirtualMeshes[mesh] = invalidVirtualMeshesForMesh
}

func (v *applyReporter) ReportFailoverServiceToMesh(mesh *discoveryv1.Mesh, failoverService ezkube.ResourceId, err error) {
	invalidFailoverServicesForMesh := v.unappliedFailoverServices[mesh]
	if invalidFailoverServicesForMesh == nil {
		invalidFailoverServicesForMesh = map[string][]error{}
	}
	key := sets.Key(failoverService)
	errs := invalidFailoverServicesForMesh[key]
	errs = append(errs, err)
	invalidFailoverServicesForMesh[key] = errs
	v.unappliedFailoverServices[mesh] = invalidFailoverServicesForMesh
}

func (v *applyReporter) ReportFailoverService(failoverService ezkube.ResourceId, newErrs []error) {
	key := sets.Key(failoverService)
	errs := v.invalidFailoverServices[key]
	if errs == nil {
		errs = []error{}
	}
	errs = append(errs, newErrs...)
	v.invalidFailoverServices[key] = errs
}

func (v *applyReporter) getTrafficPolicyErrors(destination *discoveryv1.Destination, trafficPolicy ezkube.ResourceId) []error {
	invalidTrafficPoliciesForDestination, ok := v.unappliedTrafficPolicies[destination]
	if !ok {
		return nil
	}
	tpErrors, ok := invalidTrafficPoliciesForDestination[sets.Key(trafficPolicy)]
	if !ok {
		return nil
	}
	return tpErrors
}

func (v *applyReporter) getAccessPolicyErrors(destination *discoveryv1.Destination, accessPolicy ezkube.ResourceId) []error {
	invalidAccessPoliciesForDestination, ok := v.unappliedAccessPolicies[destination]
	if !ok {
		return nil
	}
	apErrors, ok := invalidAccessPoliciesForDestination[sets.Key(accessPolicy)]
	if !ok {
		return nil
	}
	return apErrors
}

func (v *applyReporter) getFailoverServiceErrors(mesh *discoveryv1.Mesh, failoverService ezkube.ResourceId) []error {
	var errs []error
	// Mesh-dependent errors
	invalidAccessPoliciesForDestination, ok := v.unappliedFailoverServices[mesh]
	if ok {
		fsErrors, ok := invalidAccessPoliciesForDestination[sets.Key(failoverService)]
		if ok {
			errs = append(errs, fsErrors...)
		}
	}

	// Mesh-independent errors
	fsErrs := v.invalidFailoverServices[sets.Key(failoverService)]
	if fsErrs != nil {
		errs = append(errs, fsErrs...)
	}
	return errs
}

func (v *applyReporter) getVirtualMeshErrors(mesh *discoveryv1.Mesh, virtualMesh ezkube.ResourceId) []error {
	var errs []error
	// Mesh-dependent errors
	invalidAccessPoliciesForDestination, ok := v.unappliedVirtualMeshes[mesh]
	if ok {
		fsErrors, ok := invalidAccessPoliciesForDestination[sets.Key(virtualMesh)]
		if ok {
			errs = append(errs, fsErrors...)
		}
	}

	return errs
}

func getAppliedTrafficPolicies(
	trafficPolicies networkingv1.TrafficPolicySlice,
	destination *discoveryv1.Destination,
) []*discoveryv1.DestinationStatus_AppliedTrafficPolicy {
	var matchingTrafficPolicies networkingv1.TrafficPolicySlice
	for _, policy := range trafficPolicies {
		if policy.Status.State != commonv1.ApprovalState_ACCEPTED {
			continue
		}
		if selectorutils.SelectorMatchesService(policy.Spec.DestinationSelector, destination) {
			matchingTrafficPolicies = append(matchingTrafficPolicies, policy)
		}
	}

	sortTrafficPoliciesByAcceptedDate(destination, matchingTrafficPolicies)

	var appliedPolicies []*discoveryv1.DestinationStatus_AppliedTrafficPolicy
	for _, policy := range matchingTrafficPolicies {
		policy := policy // pike
		appliedPolicies = append(appliedPolicies, &discoveryv1.DestinationStatus_AppliedTrafficPolicy{
			Ref:                ezkube.MakeObjectRef(policy),
			Spec:               &policy.Spec,
			ObservedGeneration: policy.Generation,
		})
	}
	return appliedPolicies
}

// sort the set of traffic policies in the order in which they were accepted.
// Traffic policies which were accepted first and have not changed (i.e. their observedGeneration is up-to-date) take precedence.
// Next are policies that were previously accepted but whose observedGeneration is out of date. This permits policies which were modified but formerly correct to maintain
// their acceptance status ahead of policies which were unomdified and previously rejected.
// Next will be the policies which have been modified and rejected.
// Finally, policies which are rejected and modified
func sortTrafficPoliciesByAcceptedDate(destination *discoveryv1.Destination, trafficPolicies networkingv1.TrafficPolicySlice) {
	isUpToDate := func(tp *networkingv1.TrafficPolicy) bool {
		return tp.Status.ObservedGeneration == tp.Generation
	}

	sort.SliceStable(trafficPolicies, func(i, j int) bool {
		tp1, tp2 := trafficPolicies[i], trafficPolicies[j]

		status1 := tp1.Status.Destinations[sets.Key(destination)]
		status2 := tp2.Status.Destinations[sets.Key(destination)]

		if status2 == nil {
			// if status is not set, the TrafficPolicy is "pending" for this Destination
			// and should get sorted after an accepted status.
			return status1 != nil
		} else if status1 == nil {
			return false
		}

		switch {
		case status1.State == commonv1.ApprovalState_ACCEPTED:
			if status2.State != commonv1.ApprovalState_ACCEPTED {
				// accepted comes before non accepted
				return true
			}

			if tp1UpToDate := isUpToDate(tp1); tp1UpToDate != isUpToDate(tp2) {
				// up to date is validated before modified
				return tp1UpToDate
			}

			// sort by the previous acceptance order
			return status1.AcceptanceOrder < status2.AcceptanceOrder
		case status2.State == commonv1.ApprovalState_ACCEPTED:
			// accepted comes before non accepted
			return false
		default:
			// neither policy has been accepted, we can simply sort by unique key
			return sets.Key(tp1) < sets.Key(tp2)
		}
	})
}

// Fetch all AccessPolicies applicable to the given Destination.
// Sorting is not needed because the additive semantics of AccessPolicies does not allow for conflicts.
func getAppliedAccessPolicies(
	accessPolicies networkingv1.AccessPolicySlice,
	destination *discoveryv1.Destination,
) []*discoveryv1.DestinationStatus_AppliedAccessPolicy {
	var appliedPolicies []*discoveryv1.DestinationStatus_AppliedAccessPolicy
	for _, policy := range accessPolicies {
		policy := policy // pike
		if policy.Status.State != commonv1.ApprovalState_ACCEPTED {
			continue
		}
		if !selectorutils.SelectorMatchesService(policy.Spec.DestinationSelector, destination) {
			continue
		}
		appliedPolicies = append(appliedPolicies, &discoveryv1.DestinationStatus_AppliedAccessPolicy{
			Ref:                ezkube.MakeObjectRef(policy),
			Spec:               &policy.Spec,
			ObservedGeneration: policy.Generation,
		})
	}

	return appliedPolicies
}

func getAppliedVirtualMesh(
	virtualMeshes networkingv1.VirtualMeshSlice,
	mesh *discoveryv1.Mesh,
) *discoveryv1.MeshStatus_AppliedVirtualMesh {
	for _, vMesh := range virtualMeshes {
		vMesh := vMesh // pike
		if vMesh.Status.State != commonv1.ApprovalState_ACCEPTED {
			continue
		}
		for _, meshRef := range vMesh.Spec.Meshes {
			if ezkube.RefsMatch(mesh, meshRef) {
				return &discoveryv1.MeshStatus_AppliedVirtualMesh{
					Ref:                ezkube.MakeObjectRef(vMesh),
					Spec:               &vMesh.Spec,
					ObservedGeneration: vMesh.Generation,
				}
			}
		}
	}
	return nil
}

// Get all the meshes and corresponding VirtualMeshes of the given Destinations.
// Results are returned as maps keyed by mesh ObjectRef keys and VirtualMesh ObjectRef keys
func getMeshesFromDestinations(ctx context.Context, destinations []*discoveryv1.Destination,
	allMeshes discoveryv1sets.MeshSet) (utilsets.String, utilsets.String) {

	meshes := utilsets.NewString()
	virtualMeshes := utilsets.NewString()
	for _, destination := range destinations {
		meshRef := destination.Spec.GetMesh()
		if meshRef == nil {
			continue
		}
		meshKey := sets.Key(meshRef)
		if !meshes.Has(meshKey) {
			meshes.Insert(meshKey)

			// get the full mesh object to get the VirtualMesh
			mesh, err := allMeshes.Find(meshRef)
			if err != nil {
				// should never happen
				contextutils.LoggerFrom(ctx).Errorf("internal error: failed to look up mesh %v: %v", meshRef, err)
				continue
			}
			if virtualMeshRef := mesh.Status.GetAppliedVirtualMesh().GetRef(); virtualMeshRef != nil {
				virtualMeshes.Insert(sets.Key(virtualMeshRef))
			}
		}
	}
	return meshes, virtualMeshes
}

// Map each mesh ref to its VirtualMesh ref (if any).
// The keys in the returned map are mesh ref keys, and the values are VirtualMesh ref keys.
func makeMeshToVirtualMeshMap(meshes discoveryv1.MeshSlice) map[string]string {
	meshToVirtualMesh := make(map[string]string)
	for _, mesh := range meshes {
		if virtualMeshRef := mesh.Status.GetAppliedVirtualMesh().GetRef(); virtualMeshRef != nil {
			meshToVirtualMesh[sets.Key(mesh)] = sets.Key(virtualMeshRef)
		}
	}
	return meshToVirtualMesh
}

// Returns true if the given mesh either matches one of the given matchingMeshes, or it is in a VirtualMesh that
// matches one of the given matchingVirtualMeshes
func meshMatches(meshRef *v1.ObjectRef, matchingMeshes utilsets.String, matchingVirtualMeshes utilsets.String,
	meshToVirtualMesh map[string]string) bool {
	meshKey := sets.Key(meshRef)
	if matchingMeshes.Has(meshKey) {
		return true
	}
	if virtualMeshRefKey, ok := meshToVirtualMesh[meshKey]; ok {
		return matchingVirtualMeshes.Has(virtualMeshRefKey)
	}
	return false
}
