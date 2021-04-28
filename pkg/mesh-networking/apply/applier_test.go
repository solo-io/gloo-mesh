package apply_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	commonv1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	networkingv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/solo-io/gloo-mesh/pkg/mesh-networking/apply"
)

var _ = Describe("Applier", func() {
	Context("applied traffic policies", func() {
		var (
			destination = &discoveryv1.Destination{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ms1",
					Namespace: "ns",
				},
				Spec: discoveryv1.DestinationSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh1",
						Namespace: "ns",
					},
					Type: &discoveryv1.DestinationSpec_KubeService_{
						KubeService: &discoveryv1.DestinationSpec_KubeService{
							Ref: &skv2corev1.ClusterObjectRef{
								Name:        "svc-name",
								Namespace:   "svc-namespace",
								ClusterName: "svc-cluster",
							},
						},
					},
				},
			}
			workload = &discoveryv1.Workload{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "wkld1",
					Namespace: "ns",
				},
				Spec: discoveryv1.WorkloadSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh1",
						Namespace: "ns",
					},
				},
			}
			mesh = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh1",
					Namespace: "ns",
				},
			}
			trafficPolicy1 = &networkingv1.TrafficPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tp1",
					Namespace: "ns",
				},
				Spec: networkingv1.TrafficPolicySpec{
					Policy: &networkingv1.TrafficPolicySpec_Policy{
						// fill an arbitrary part of the spec
						Mirror: &networkingv1.TrafficPolicySpec_Policy_Mirror{},
					},
				},
			}
			trafficPolicy2 = &networkingv1.TrafficPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tp2",
					Namespace: "ns",
				},
				Spec: networkingv1.TrafficPolicySpec{
					Policy: &networkingv1.TrafficPolicySpec_Policy{
						// fill an arbitrary part of the spec
						FaultInjection: &networkingv1.TrafficPolicySpec_Policy_FaultInjection{},
					},
				},
			}

			snap = input.NewInputLocalSnapshotManualBuilder("").
				AddDestinations(discoveryv1.DestinationSlice{destination}).
				AddTrafficPolicies(networkingv1.TrafficPolicySlice{trafficPolicy1, trafficPolicy2}).
				AddWorkloads(discoveryv1.WorkloadSlice{workload}).
				AddMeshes(discoveryv1.MeshSlice{mesh}).
				Build()
		)

		BeforeEach(func() {
			translator := testIstioTranslator{callReporter: func(reporter reporting.Reporter) {
				// no report = accept
			}}
			applier := NewApplier(translator)
			applier.Apply(context.TODO(), snap, nil)
		})
		It("updates status on input traffic policies", func() {
			Expect(trafficPolicy1.Status.Destinations).To(HaveKey(sets.Key(destination)))
			Expect(trafficPolicy1.Status.Destinations[sets.Key(destination)]).To(Equal(&networkingv1.ApprovalStatus{
				AcceptanceOrder: 0,
				State:           commonv1.ApprovalState_ACCEPTED,
			}))
			Expect(trafficPolicy1.Status.Workloads).To(HaveLen(1))
			Expect(trafficPolicy1.Status.Workloads[0]).To(Equal(sets.Key(workload)))
			Expect(trafficPolicy2.Status.Destinations).To(HaveKey(sets.Key(destination)))
			Expect(trafficPolicy2.Status.Destinations[sets.Key(destination)]).To(Equal(&networkingv1.ApprovalStatus{
				AcceptanceOrder: 1,
				State:           commonv1.ApprovalState_ACCEPTED,
			}))
			Expect(trafficPolicy2.Status.Workloads).To(HaveLen(1))
			Expect(trafficPolicy2.Status.Workloads[0]).To(Equal(sets.Key(workload)))

		})
		It("updates status on input Destination policies", func() {
			Expect(destination.Status.AppliedTrafficPolicies).To(HaveLen(2))
			Expect(destination.Status.AppliedTrafficPolicies[0].Ref).To(Equal(ezkube.MakeObjectRef(trafficPolicy1)))
			Expect(destination.Status.AppliedTrafficPolicies[0].Spec).To(Equal(&trafficPolicy1.Spec))
			Expect(destination.Status.AppliedTrafficPolicies[1].Ref).To(Equal(ezkube.MakeObjectRef(trafficPolicy2)))
			Expect(destination.Status.AppliedTrafficPolicies[1].Spec).To(Equal(&trafficPolicy2.Spec))
			Expect(destination.Status.LocalFqdn).To(Equal("svc-name.svc-namespace.svc.cluster.local"))
		})
	})
	Context("invalid traffic policies", func() {
		var (
			destination = &discoveryv1.Destination{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ms1",
					Namespace: "ns",
				},
			}
			trafficPolicy = &networkingv1.TrafficPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tp1",
					Namespace: "ns",
				},
			}

			snap = input.NewInputLocalSnapshotManualBuilder("").
				AddDestinations(discoveryv1.DestinationSlice{destination}).
				AddTrafficPolicies(networkingv1.TrafficPolicySlice{trafficPolicy}).
				Build()
		)

		BeforeEach(func() {
			translator := testIstioTranslator{callReporter: func(reporter reporting.Reporter) {
				// report = reject
				reporter.ReportTrafficPolicyToDestination(destination, trafficPolicy, errors.New("did an oopsie"))
			}}
			applier := NewApplier(translator)
			applier.Apply(context.TODO(), snap, nil)
		})
		It("updates status on input traffic policies", func() {
			Expect(trafficPolicy.Status.Destinations).To(HaveKey(sets.Key(destination)))
			Expect(trafficPolicy.Status.Destinations[sets.Key(destination)]).To(Equal(&networkingv1.ApprovalStatus{
				AcceptanceOrder: 0,
				State:           commonv1.ApprovalState_INVALID,
				Errors:          []string{"did an oopsie"},
			}))
		})
		It("does not add the policy to the Destination status", func() {
			Expect(destination.Status.AppliedTrafficPolicies).To(HaveLen(0))
		})
	})

	Context("setting workloads status", func() {
		var (
			destination = &discoveryv1.Destination{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ms1",
					Namespace: "ns",
				},
				Spec: discoveryv1.DestinationSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh1",
						Namespace: "ns",
					},
				},
			}
			workload1 = &discoveryv1.Workload{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "wkld1",
					Namespace: "ns",
				},
				Spec: discoveryv1.WorkloadSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh1",
						Namespace: "ns",
					},
				},
			}
			workload2 = &discoveryv1.Workload{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "wkld2",
					Namespace: "ns",
				},
				Spec: discoveryv1.WorkloadSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh2",
						Namespace: "ns",
					},
				},
			}
			mesh1 = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh1",
					Namespace: "ns",
				},
			}
			mesh2 = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh2",
					Namespace: "ns",
				},
			}
			virtualMesh = &networkingv1.VirtualMesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "vmesh1",
					Namespace: "ns",
				},
				Spec: networkingv1.VirtualMeshSpec{
					Meshes: []*skv2corev1.ObjectRef{
						{Name: "mesh1", Namespace: "ns"},
						{Name: "mesh2", Namespace: "ns"},
					},
				},
			}
			trafficPolicy = &networkingv1.TrafficPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tp1",
					Namespace: "ns",
				},
			}
			accessPolicy = &networkingv1.AccessPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ap1",
					Namespace: "ns",
				},
			}
		)

		It("sets policy workloads using mesh", func() {
			snap := input.NewInputLocalSnapshotManualBuilder("").
				AddTrafficPolicies(networkingv1.TrafficPolicySlice{trafficPolicy}).
				AddAccessPolicies(networkingv1.AccessPolicySlice{accessPolicy}).
				AddDestinations(discoveryv1.DestinationSlice{destination}).
				AddWorkloads(discoveryv1.WorkloadSlice{workload1, workload2}).
				AddMeshes(discoveryv1.MeshSlice{mesh1, mesh2}).
				Build()
			translator := testIstioTranslator{callReporter: func(reporter reporting.Reporter) {
				// no report = accept
			}}
			applier := NewApplier(translator)
			applier.Apply(context.TODO(), snap, nil)

			// destination and workload1 are both in mesh1
			Expect(trafficPolicy.Status.Workloads).To(HaveLen(1))
			Expect(trafficPolicy.Status.Workloads[0]).To(Equal(sets.Key(workload1)))
			Expect(accessPolicy.Status.Workloads).To(HaveLen(1))
			Expect(accessPolicy.Status.Workloads[0]).To(Equal(sets.Key(workload1)))
		})
		It("sets policy workloads using VirtualMesh", func() {
			snap := input.NewInputLocalSnapshotManualBuilder("").
				AddTrafficPolicies(networkingv1.TrafficPolicySlice{trafficPolicy}).
				AddAccessPolicies(networkingv1.AccessPolicySlice{accessPolicy}).
				AddDestinations(discoveryv1.DestinationSlice{destination}).
				AddWorkloads(discoveryv1.WorkloadSlice{workload1, workload2}).
				AddMeshes(discoveryv1.MeshSlice{mesh1, mesh2}).
				AddVirtualMeshes(networkingv1.VirtualMeshSlice{virtualMesh}).
				Build()
			translator := testIstioTranslator{callReporter: func(reporter reporting.Reporter) {
				// no report = accept
			}}
			applier := NewApplier(translator)
			applier.Apply(context.TODO(), snap, nil)

			// destination is in mesh1, workload1 is in mesh1, and workload2 is in mesh2.
			// since mesh1 and mesh2 are in the same VirtualMesh, both workloads are returned
			Expect(trafficPolicy.Status.Workloads).To(HaveLen(2))
			Expect(trafficPolicy.Status.Workloads[0]).To(Equal(sets.Key(workload1)))
			Expect(trafficPolicy.Status.Workloads[1]).To(Equal(sets.Key(workload2)))
			Expect(accessPolicy.Status.Workloads).To(HaveLen(2))
			Expect(accessPolicy.Status.Workloads[0]).To(Equal(sets.Key(workload1)))
			Expect(accessPolicy.Status.Workloads[1]).To(Equal(sets.Key(workload2)))
		})
		It("sets no policy workloads when there is no matching mesh", func() {
			workload1.Spec.Mesh.Name = "mesh2"
			snap := input.NewInputLocalSnapshotManualBuilder("").
				AddTrafficPolicies(networkingv1.TrafficPolicySlice{trafficPolicy}).
				AddAccessPolicies(networkingv1.AccessPolicySlice{accessPolicy}).
				AddDestinations(discoveryv1.DestinationSlice{destination}).
				AddWorkloads(discoveryv1.WorkloadSlice{workload1, workload2}).
				AddMeshes(discoveryv1.MeshSlice{mesh1, mesh2}).
				Build()
			translator := testIstioTranslator{callReporter: func(reporter reporting.Reporter) {
				// no report = accept
			}}
			applier := NewApplier(translator)
			applier.Apply(context.TODO(), snap, nil)

			// destination is in mesh1, but both workloads are in mesh2
			Expect(trafficPolicy.Status.Workloads).To(BeNil())
			Expect(accessPolicy.Status.Workloads).To(BeNil())
		})
	})

	Context("applied federation", func() {
		var (
			applier Applier

			destination = &discoveryv1.Destination{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ms1",
					Namespace: "ns",
				},
				Spec: discoveryv1.DestinationSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh1",
						Namespace: "ns",
					},
					Type: &discoveryv1.DestinationSpec_KubeService_{
						KubeService: &discoveryv1.DestinationSpec_KubeService{
							Ref: &skv2corev1.ClusterObjectRef{
								Name:        "svc-name",
								Namespace:   "svc-namespace",
								ClusterName: "svc-cluster",
							},
						},
					},
				},
			}

			mesh1 = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh1",
					Namespace: "ns",
				},
			}
			mesh2 = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh2",
					Namespace: "ns",
				},
			}
			mesh3 = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh3",
					Namespace: "ns",
				},
			}
			mesh4 = &discoveryv1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "mesh4",
					Namespace: "ns",
				},
			}

			snap = input.NewInputLocalSnapshotManualBuilder("").
				AddDestinations(discoveryv1.DestinationSlice{destination}).
				AddMeshes(discoveryv1.MeshSlice{mesh1, mesh2, mesh3, mesh4})
		)

		BeforeEach(func() {
			translator := testIstioTranslator{callReporter: func(reporter reporting.Reporter) {
				// no report = accept
			}}
			applier = NewApplier(translator)
		})

		It("applies VirtualMesh with permissive federation", func() {
			permissiveVirtualMesh := &networkingv1.VirtualMesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "vm1",
					Namespace: "ns",
				},
				Spec: networkingv1.VirtualMeshSpec{
					Meshes: []*skv2corev1.ObjectRef{
						ezkube.MakeObjectRef(mesh1),
						ezkube.MakeObjectRef(mesh2),
						ezkube.MakeObjectRef(mesh3),
						ezkube.MakeObjectRef(mesh4),
					},
					Federation: &networkingv1.VirtualMeshSpec_Federation{
						Mode: &networkingv1.VirtualMeshSpec_Federation_Permissive{},
					},
				},
			}

			snap.AddVirtualMeshes([]*networkingv1.VirtualMesh{permissiveVirtualMesh})

			expectedAppliedFederation := &discoveryv1.DestinationStatus_AppliedFederation{
				FederatedHostname: "svc-name.svc-namespace.svc.svc-cluster.global",
				FederatedToMeshes: []*skv2corev1.ObjectRef{
					ezkube.MakeObjectRef(mesh2),
					ezkube.MakeObjectRef(mesh3),
					ezkube.MakeObjectRef(mesh4),
				},
				VirtualMeshRef: ezkube.MakeObjectRef(permissiveVirtualMesh),
			}

			applier.Apply(context.TODO(), snap.Build(), nil)

			Expect(destination.Status.AppliedFederation).To(Equal(expectedAppliedFederation))
		})

		It("restrictive federation with empty selectors should have permissive semantics", func() {
			restrictiveVirtualMesh := &networkingv1.VirtualMesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "vm1",
					Namespace: "ns",
				},
				Spec: networkingv1.VirtualMeshSpec{
					Meshes: []*skv2corev1.ObjectRef{
						ezkube.MakeObjectRef(mesh1),
						ezkube.MakeObjectRef(mesh2),
						ezkube.MakeObjectRef(mesh3),
						ezkube.MakeObjectRef(mesh4),
					},
					Federation: &networkingv1.VirtualMeshSpec_Federation{
						Mode: &networkingv1.VirtualMeshSpec_Federation_Restrictive{},
					},
				},
			}

			snap.AddVirtualMeshes([]*networkingv1.VirtualMesh{restrictiveVirtualMesh})

			expectedAppliedFederation := &discoveryv1.DestinationStatus_AppliedFederation{
				FederatedHostname: "svc-name.svc-namespace.svc.svc-cluster.global",
				FederatedToMeshes: []*skv2corev1.ObjectRef{
					ezkube.MakeObjectRef(mesh2),
					ezkube.MakeObjectRef(mesh3),
					ezkube.MakeObjectRef(mesh4),
				},
				VirtualMeshRef: ezkube.MakeObjectRef(restrictiveVirtualMesh),
			}

			applier.Apply(context.TODO(), snap.Build(), nil)

			Expect(destination.Status.AppliedFederation).To(Equal(expectedAppliedFederation))
		})

		It("restrictive federation with defined selectors should selectively federate Destinations", func() {
			destination2 := &discoveryv1.Destination{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "d2",
					Namespace: "ns",
				},
				Spec: discoveryv1.DestinationSpec{
					Mesh: &skv2corev1.ObjectRef{
						Name:      "mesh1",
						Namespace: "ns",
					},
					Type: &discoveryv1.DestinationSpec_KubeService_{
						KubeService: &discoveryv1.DestinationSpec_KubeService{
							Ref: &skv2corev1.ClusterObjectRef{
								Name:        "svc-name2",
								Namespace:   "svc-namespace",
								ClusterName: "svc-cluster",
							},
						},
					},
				},
			}

			restrictiveVirtualMesh := &networkingv1.VirtualMesh{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "vm1",
					Namespace: "ns",
				},
				Spec: networkingv1.VirtualMeshSpec{
					Meshes: []*skv2corev1.ObjectRef{
						ezkube.MakeObjectRef(mesh1),
						ezkube.MakeObjectRef(mesh2),
						ezkube.MakeObjectRef(mesh3),
						ezkube.MakeObjectRef(mesh4),
					},
					Federation: &networkingv1.VirtualMeshSpec_Federation{
						Mode: &networkingv1.VirtualMeshSpec_Federation_Restrictive{
							Restrictive: &networkingv1.VirtualMeshSpec_Federation_RestrictiveFederation{
								FederationSelectors: []*networkingv1.VirtualMeshSpec_Federation_RestrictiveFederation_FederationSelector{
									{
										DestinationSelectors: []*commonv1.DestinationSelector{
											{
												KubeServiceRefs: &commonv1.DestinationSelector_KubeServiceRefs{
													Services: []*skv2corev1.ClusterObjectRef{
														{
															Name:        destination.Spec.GetKubeService().GetRef().GetName(),
															Namespace:   destination.Spec.GetKubeService().GetRef().GetNamespace(),
															ClusterName: destination.Spec.GetKubeService().GetRef().GetClusterName(),
														},
													},
												},
											},
										},
										Meshes: []*skv2corev1.ObjectRef{
											ezkube.MakeObjectRef(mesh2),
											ezkube.MakeObjectRef(mesh4),
										},
									},
									{
										DestinationSelectors: []*commonv1.DestinationSelector{
											{
												KubeServiceMatcher: &commonv1.DestinationSelector_KubeServiceMatcher{
													Namespaces: []string{destination.Spec.GetKubeService().GetRef().GetNamespace()},
													Clusters:   []string{destination.Spec.GetKubeService().GetRef().GetClusterName()},
												},
											},
										},
										// multiple references to the same mesh across different federation selectors should be deduplicated
										Meshes: []*skv2corev1.ObjectRef{
											ezkube.MakeObjectRef(mesh2),
											ezkube.MakeObjectRef(mesh3),
										},
									},
								},
							},
						},
					},
				},
			}

			snap.AddDestinations([]*discoveryv1.Destination{destination2})
			snap.AddVirtualMeshes([]*networkingv1.VirtualMesh{restrictiveVirtualMesh})

			expectedAppliedFederation1 := &discoveryv1.DestinationStatus_AppliedFederation{
				FederatedHostname: "svc-name.svc-namespace.svc.svc-cluster.global",
				FederatedToMeshes: []*skv2corev1.ObjectRef{
					ezkube.MakeObjectRef(mesh2),
					ezkube.MakeObjectRef(mesh3),
					ezkube.MakeObjectRef(mesh4),
				},
				VirtualMeshRef: ezkube.MakeObjectRef(restrictiveVirtualMesh),
			}

			expectedAppliedFederation2 := &discoveryv1.DestinationStatus_AppliedFederation{
				FederatedHostname: "svc-name2.svc-namespace.svc.svc-cluster.global",
				FederatedToMeshes: []*skv2corev1.ObjectRef{
					ezkube.MakeObjectRef(mesh2),
					ezkube.MakeObjectRef(mesh3),
				},
				VirtualMeshRef: ezkube.MakeObjectRef(restrictiveVirtualMesh),
			}

			applier.Apply(context.TODO(), snap.Build(), nil)

			Expect(destination.Status.AppliedFederation).To(Equal(expectedAppliedFederation1))
			Expect(destination2.Status.AppliedFederation).To(Equal(expectedAppliedFederation2))
		})
	})
})

// NOTE(ilackarms): we implement a test translator here instead of using a mock because
// we need to call methods on the reporter which is passed as an argument to the translator
type testIstioTranslator struct {
	callReporter func(reporter reporting.Reporter)
}

func (t testIstioTranslator) Translate(
	ctx context.Context,
	in input.LocalSnapshot,
	existingIstioResources input.RemoteSnapshot,
	reporter reporting.Reporter,
) (*translation.Outputs, error) {
	t.callReporter(reporter)
	return &translation.Outputs{}, nil
}
