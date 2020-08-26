package federation_test

import (
	"context"

	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/output/istio"
	"github.com/solo-io/service-mesh-hub/test/data"

	"github.com/gogo/protobuf/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	istiov1alpha3sets "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3/sets"
	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	discoveryv1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/input"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/sets"
	"github.com/solo-io/service-mesh-hub/pkg/common/defaults"
	. "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/istio/mesh/federation"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/utils/hostutils"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-networking/translation/utils/metautils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	skv1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	skv1alpha1sets "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	networkingv1alpha3spec "istio.io/api/networking/v1alpha3"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("FederationTranslator", func() {
	ctx := context.TODO()
	clusterDomains := hostutils.NewClusterDomainRegistry(skv1alpha1sets.NewKubernetesClusterSet())

	It("translates federation resources for a virtual mesh", func() {

		namespace := "namespace"
		clusterName := "cluster"

		mesh := &discoveryv1alpha2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "config-namespace",
				Name:      "federated-mesh",
			},
			Spec: discoveryv1alpha2.MeshSpec{
				MeshType: &discoveryv1alpha2.MeshSpec_Istio_{Istio: &discoveryv1alpha2.MeshSpec_Istio{
					Installation: &discoveryv1alpha2.MeshSpec_MeshInstallation{
						Namespace: namespace,
						Cluster:   clusterName,
						Version:   "1.7.0-rc1",
					},
					IngressGateways: []*discoveryv1alpha2.MeshSpec_Istio_IngressGatewayInfo{{
						ExternalAddress:  "mesh-gateway.dns.name",
						ExternalTlsPort:  8181,
						TlsContainerPort: 9191,
						WorkloadLabels:   map[string]string{"gatewaylabels": "righthere"},
					}},
				}},
			},
		}

		clientMesh := &discoveryv1alpha2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "config-namespace",
				Name:      "client-mesh",
			},
			Spec: discoveryv1alpha2.MeshSpec{
				MeshType: &discoveryv1alpha2.MeshSpec_Istio_{Istio: &discoveryv1alpha2.MeshSpec_Istio{
					Installation: &discoveryv1alpha2.MeshSpec_MeshInstallation{
						Namespace: "remote-namespace",
						Cluster:   "remote-cluster",
					},
				}},
			},
		}

		meshRef := ezkube.MakeObjectRef(mesh)
		clientMeshRef := ezkube.MakeObjectRef(clientMesh)

		makeTrafficSplit := func(backingService *v1.ClusterObjectRef, subset map[string]string) *discoveryv1alpha2.TrafficTargetStatus_AppliedTrafficPolicy {
			return &discoveryv1alpha2.TrafficTargetStatus_AppliedTrafficPolicy{Spec: &data.RemoteTrafficShiftPolicy(
				"",
				"",
				backingService,
				clusterName,
				// NOTE(ilackarms): we only care about the subset labels here
				subset,
				0,
			).Spec}
		}

		backingService := &v1.ClusterObjectRef{
			Name:        "some-svc",
			Namespace:   "some-ns",
			ClusterName: clusterName,
		}
		trafficTarget1 := &discoveryv1alpha2.TrafficTarget{
			ObjectMeta: metav1.ObjectMeta{},
			Spec: discoveryv1alpha2.TrafficTargetSpec{
				Type: &discoveryv1alpha2.TrafficTargetSpec_KubeService_{KubeService: &discoveryv1alpha2.TrafficTargetSpec_KubeService{
					Ref: backingService,
					Ports: []*discoveryv1alpha2.TrafficTargetSpec_KubeService_KubeServicePort{
						{
							Port:     1234,
							Name:     "http",
							Protocol: "TCP",
						},
						{
							Port:     5678,
							Name:     "grpc",
							Protocol: "TCP",
						},
					},
				}},
				Mesh: meshRef,
			},
			// include some applied subsets
			Status: discoveryv1alpha2.TrafficTargetStatus{
				AppliedTrafficPolicies: []*discoveryv1alpha2.TrafficTargetStatus_AppliedTrafficPolicy{
					makeTrafficSplit(backingService, map[string]string{"foo": "bar"}),
					makeTrafficSplit(backingService, map[string]string{"foo": "baz"}),
					makeTrafficSplit(backingService, map[string]string{"bar": "qux"}),
				},
			},
		}

		vMesh := &discoveryv1alpha2.MeshStatus_AppliedVirtualMesh{
			Ref: &v1.ObjectRef{
				Name:      "my-virtual-mesh",
				Namespace: "config-namespace",
			},
			Spec: &v1alpha2.VirtualMeshSpec{
				Meshes: []*v1.ObjectRef{
					meshRef,
					clientMeshRef,
				},
			},
		}

		kubeCluster := &skv1alpha1.KubernetesCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name:      clusterName,
				Namespace: defaults.GetPodNamespace(),
			},
		}

		in := input.NewSnapshot(
			"ignored",
			discoveryv1alpha2sets.NewTrafficTargetSet(trafficTarget1), discoveryv1alpha2sets.NewWorkloadSet(), discoveryv1alpha2sets.NewMeshSet(mesh, clientMesh),

			v1alpha2sets.NewTrafficPolicySet(),
			v1alpha2sets.NewAccessPolicySet(),
			v1alpha2sets.NewVirtualMeshSet(),
			v1alpha2sets.NewFailoverServiceSet(),
			corev1sets.NewSecretSet(),
			skv1alpha1sets.NewKubernetesClusterSet(kubeCluster),
		)

		t := NewTranslator(ctx, clusterDomains, in.TrafficTargets())
		outputs := istio.NewBuilder(context.TODO(), "")
		t.Translate(
			in,
			mesh,
			vMesh,
			outputs,
			nil, // no reports expected
		)

		Expect(outputs.GetGateways().Length()).To(Equal(1))
		Expect(outputs.GetGateways().List()[0]).To(Equal(expectedGateway))
		Expect(outputs.GetEnvoyFilters().Length()).To(Equal(1))
		Expect(outputs.GetEnvoyFilters().List()[0]).To(Equal(expectedEnvoyFilter))
		Expect(outputs.GetDestinationRules()).To(Equal(expectedDestinationRules))
		Expect(outputs.GetServiceEntries()).To(Equal(expectedServiceEntries))

	})
})

var expectedGateway = &networkingv1alpha3.Gateway{
	ObjectMeta: metav1.ObjectMeta{
		Name:        "my-virtual-mesh-config-namespace",
		Namespace:   "namespace",
		ClusterName: "cluster",
		Labels:      metautils.TranslatedObjectLabels(),
	},
	Spec: networkingv1alpha3spec.Gateway{
		Servers: []*networkingv1alpha3spec.Server{
			{
				Port: &networkingv1alpha3spec.Port{
					Number:   9191,
					Protocol: "TLS",
					Name:     "tls",
				},
				Hosts: []string{
					"*.global",
				},
				Tls: &networkingv1alpha3spec.ServerTLSSettings{
					Mode: networkingv1alpha3spec.ServerTLSSettings_AUTO_PASSTHROUGH,
				},
			},
		},
		Selector: map[string]string{"gatewaylabels": "righthere"},
	},
}
var expectedEnvoyFilter = &networkingv1alpha3.EnvoyFilter{
	ObjectMeta: metav1.ObjectMeta{
		Name:        "my-virtual-mesh.config-namespace",
		Namespace:   "namespace",
		ClusterName: "cluster",
		Labels:      metautils.TranslatedObjectLabels(),
	},
	Spec: networkingv1alpha3spec.EnvoyFilter{
		WorkloadSelector: &networkingv1alpha3spec.WorkloadSelector{
			Labels: map[string]string{"gatewaylabels": "righthere"},
		},
		ConfigPatches: []*networkingv1alpha3spec.EnvoyFilter_EnvoyConfigObjectPatch{
			{
				ApplyTo: networkingv1alpha3spec.EnvoyFilter_NETWORK_FILTER,
				Match: &networkingv1alpha3spec.EnvoyFilter_EnvoyConfigObjectMatch{
					Context: networkingv1alpha3spec.EnvoyFilter_GATEWAY,
					ObjectTypes: &networkingv1alpha3spec.EnvoyFilter_EnvoyConfigObjectMatch_Listener{
						Listener: &networkingv1alpha3spec.EnvoyFilter_ListenerMatch{
							PortNumber: 9191,
							FilterChain: &networkingv1alpha3spec.EnvoyFilter_ListenerMatch_FilterChainMatch{
								Filter: &networkingv1alpha3spec.EnvoyFilter_ListenerMatch_FilterMatch{
									Name: "envoy.filters.network.sni_cluster",
								},
							},
						},
					},
				},
				Patch: &networkingv1alpha3spec.EnvoyFilter_Patch{
					Operation: 5,
					Value: &types.Struct{
						Fields: map[string]*types.Value{
							"name": {
								Kind: &types.Value_StringValue{
									StringValue: "envoy.filters.network.tcp_cluster_rewrite",
								},
							},
							"typed_config": {
								Kind: &types.Value_StructValue{
									StructValue: &types.Struct{
										Fields: map[string]*types.Value{
											"@type": {
												Kind: &types.Value_StringValue{
													StringValue: "type.googleapis.com/istio.envoy.config.filter.network.tcp_cluster_rewrite.v2alpha1.TcpClusterRewrite",
												},
											},
											"cluster_replacement": {
												Kind: &types.Value_StringValue{
													StringValue: ".cluster.local",
												},
											},
											"cluster_pattern": {
												Kind: &types.Value_StringValue{
													StringValue: "\\.cluster.global$",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
var expectedDestinationRules = istiov1alpha3sets.NewDestinationRuleSet(&networkingv1alpha3.DestinationRule{
	ObjectMeta: metav1.ObjectMeta{
		Name:        "some-svc.some-ns.svc.cluster.global",
		Namespace:   "remote-namespace",
		ClusterName: "remote-cluster",
		Labels:      metautils.TranslatedObjectLabels(),
	},
	Spec: networkingv1alpha3spec.DestinationRule{
		Host: "some-svc.some-ns.svc.cluster.global",
		TrafficPolicy: &networkingv1alpha3spec.TrafficPolicy{
			Tls: &networkingv1alpha3spec.ClientTLSSettings{
				Mode: networkingv1alpha3spec.ClientTLSSettings_ISTIO_MUTUAL,
			},
		},
		Subsets: []*networkingv1alpha3spec.Subset{
			{
				Name:   "foo-bar",
				Labels: map[string]string{"cluster": "cluster"},
			},
			{
				Name:   "foo-baz",
				Labels: map[string]string{"cluster": "cluster"},
			},
			{
				Name:   "bar-qux",
				Labels: map[string]string{"cluster": "cluster"},
			},
		},
	},
})
var expectedServiceEntries = istiov1alpha3sets.NewServiceEntrySet(&networkingv1alpha3.ServiceEntry{
	ObjectMeta: metav1.ObjectMeta{
		Name:        "some-svc.some-ns.svc.cluster.global",
		Namespace:   "remote-namespace",
		ClusterName: "remote-cluster",
		Labels:      metautils.TranslatedObjectLabels(),
	},
	Spec: networkingv1alpha3spec.ServiceEntry{
		Hosts: []string{
			"some-svc.some-ns.svc.cluster.global",
		},
		Addresses: []string{
			"243.21.204.125",
		},
		Ports: []*networkingv1alpha3spec.Port{
			{
				Number:   1234,
				Protocol: "TCP",
				Name:     "http",
			},
			{
				Number:   5678,
				Protocol: "TCP",
				Name:     "grpc",
			},
		},
		Location:   networkingv1alpha3spec.ServiceEntry_MESH_INTERNAL,
		Resolution: networkingv1alpha3spec.ServiceEntry_DNS,
		Endpoints: []*networkingv1alpha3spec.WorkloadEntry{
			{
				Address: "mesh-gateway.dns.name",
				Ports: map[string]uint32{
					"http": 8181,
					"grpc": 8181,
				},
				Labels: map[string]string{"cluster": "cluster"},
			},
		},
	},
})
