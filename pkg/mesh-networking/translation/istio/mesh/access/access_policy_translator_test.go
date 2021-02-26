package access_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/istio/security.istio.io/v1beta1/sets"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	networkingv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/mesh/access"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	securityv1beta1spec "istio.io/api/security/v1beta1"
	"istio.io/api/type/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("AccessPolicyTranslator", func() {
	var (
		translator access.Translator
	)

	BeforeEach(func() {
		translator = access.NewTranslator(context.Background())
	})

	It("should translate an AuthorizationPolicy for the ingress gateway and in the installation namespace", func() {
		mesh := &discoveryv1.Mesh{
			Spec: discoveryv1.MeshSpec{
				Type: &discoveryv1.MeshSpec_Istio_{
					Istio: &discoveryv1.MeshSpec_Istio{
						Installation: &discoveryv1.MeshSpec_MeshInstallation{
							Namespace: "istio-system",
							Cluster:   "cluster-name",
						},
						IngressGateways: []*discoveryv1.MeshSpec_Istio_IngressGatewayInfo{
							{
								WorkloadLabels: map[string]string{
									"istio": "ingressgateway",
								},
								ExternalAddress: "1.1.1.1",
							},
							{
								WorkloadLabels: map[string]string{
									"istio": "ingressgateway2",
								},
								ExternalAddress: "2.2.2.2",
							},
						},
					},
				},
			},
			Status: discoveryv1.MeshStatus{
				AppliedVirtualMesh: &discoveryv1.MeshStatus_AppliedVirtualMesh{
					Ref: &v1.ObjectRef{
						Name:      "virtual-mesh",
						Namespace: "gloo-mesh",
					},
					Spec: &networkingv1.VirtualMeshSpec{
						GlobalAccessPolicy: networkingv1.VirtualMeshSpec_ENABLED,
					},
				},
			},
		}
		expectedAuthPolicies := v1beta1sets.NewAuthorizationPolicySet(
			&securityv1beta1.AuthorizationPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:        access.IngressGatewayAuthPolicyName + "-1-1-1-1",
					Namespace:   "istio-system",
					ClusterName: "cluster-name",
					Labels:      metautils.TranslatedObjectLabels(),
					Annotations: map[string]string{
						metautils.ParentLabelkey: `{"networking.mesh.gloo.solo.io/v1, Kind=VirtualMesh":[{"name":"virtual-mesh","namespace":"gloo-mesh"}]}`,
					},
				},
				Spec: securityv1beta1spec.AuthorizationPolicy{
					Action: securityv1beta1spec.AuthorizationPolicy_ALLOW,
					// A single empty rule allows all traffic.
					// Reference: https://istio.io/docs/reference/config/security/authorization-policy/#AuthorizationPolicy
					Rules: []*securityv1beta1spec.Rule{{}},
					Selector: &v1beta1.WorkloadSelector{
						MatchLabels: map[string]string{
							"istio": "ingressgateway",
						},
					},
				},
			},
			&securityv1beta1.AuthorizationPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:        access.IngressGatewayAuthPolicyName + "-2-2-2-2",
					Namespace:   "istio-system",
					ClusterName: "cluster-name",
					Labels:      metautils.TranslatedObjectLabels(),
					Annotations: map[string]string{
						metautils.ParentLabelkey: `{"networking.mesh.gloo.solo.io/v1, Kind=VirtualMesh":[{"name":"virtual-mesh","namespace":"gloo-mesh"}]}`,
					},
				},
				Spec: securityv1beta1spec.AuthorizationPolicy{
					Action: securityv1beta1spec.AuthorizationPolicy_ALLOW,
					// A single empty rule allows all traffic.
					// Reference: https://istio.io/docs/reference/config/security/authorization-policy/#AuthorizationPolicy
					Rules: []*securityv1beta1spec.Rule{{}},
					Selector: &v1beta1.WorkloadSelector{
						MatchLabels: map[string]string{
							"istio": "ingressgateway2",
						},
					},
				},
			},
			&securityv1beta1.AuthorizationPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:        access.GlobalAccessControlAuthPolicyName,
					Namespace:   "istio-system",
					ClusterName: "cluster-name",
					Labels:      metautils.TranslatedObjectLabels(),
					Annotations: map[string]string{
						metautils.ParentLabelkey: `{"networking.mesh.gloo.solo.io/v1, Kind=VirtualMesh":[{"name":"virtual-mesh","namespace":"gloo-mesh"}]}`,
					},
				},
				Spec: securityv1beta1spec.AuthorizationPolicy{},
			},
		)
		outputs := istio.NewBuilder(context.TODO(), "")
		translator.Translate(mesh, mesh.Status.AppliedVirtualMesh, outputs)
		Expect(outputs.GetAuthorizationPolicies()).To(Equal(expectedAuthPolicies))
	})

	It("should not translate any AuthorizationPolicies", func() {
		mesh := &discoveryv1.Mesh{
			Spec: discoveryv1.MeshSpec{
				Type: &discoveryv1.MeshSpec_Istio_{
					Istio: &discoveryv1.MeshSpec_Istio{
						Installation: &discoveryv1.MeshSpec_MeshInstallation{
							Namespace: "istio-system",
						},
					},
				},
			},
			Status: discoveryv1.MeshStatus{
				AppliedVirtualMesh: &discoveryv1.MeshStatus_AppliedVirtualMesh{
					Spec: &networkingv1.VirtualMeshSpec{
						GlobalAccessPolicy: networkingv1.VirtualMeshSpec_DISABLED,
					},
				},
			},
		}
		outputs := istio.NewBuilder(context.TODO(), "")
		translator.Translate(mesh, mesh.Status.AppliedVirtualMesh, outputs)
		Expect(outputs.GetAuthorizationPolicies().Length()).To(Equal(0))
	})
})
