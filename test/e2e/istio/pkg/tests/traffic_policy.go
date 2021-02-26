package tests

import (
	"context"
	"time"

	networkingv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	"github.com/solo-io/gloo-mesh/test/e2e"
	"github.com/solo-io/gloo-mesh/test/utils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	istionetworkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/solo-io/gloo-mesh/test/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TrafficPolicyTest() {
	var (
		err      error
		manifest utils.Manifest
		ctx      = context.Background()
	)

	BeforeEach(func() {
		manifest, err = utils.NewManifest("bookinfo-policies.yaml")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		manifest.Cleanup(BookinfoNamespace)
	})

	It("applies traffic shift policies to local subsets", func() {

		By("initially curling reviews should return both reviews-v1 and reviews-v2", func() {
			Eventually(CurlReviews, "1m", "1s").Should(ContainSubstring(`"color": "black"`))
			Eventually(CurlReviews, "1m", "1s").ShouldNot(ContainSubstring(`"color": "black"`))
		})

		By("creating a TrafficPolicy with traffic shift to reviews-v2 should consistently shift traffic", func() {
			trafficShiftReviewsV2 := data.LocalTrafficShiftPolicy("bookinfo-policy", BookinfoNamespace, &v1.ClusterObjectRef{
				Name:        "reviews",
				Namespace:   BookinfoNamespace,
				ClusterName: MgmtClusterName,
			}, map[string]string{"version": "v2"}, 9080)

			err = manifest.AppendResources(trafficShiftReviewsV2)
			Expect(err).NotTo(HaveOccurred())
			err = manifest.KubeApply(BookinfoNamespace)
			Expect(err).NotTo(HaveOccurred())

			// ensure status is updated
			utils.AssertTrafficPolicyStatuses(ctx, e2e.GetEnv().Management.TrafficPolicyClient, BookinfoNamespace)

			// insert a sleep because we can't effectively use an Eventually here to ensure config has propagated to envoy
			time.Sleep(time.Second * 5)

			// check we can eventually (consistently) hit the v2 subset
			Eventually(CurlReviews, "30s", "0.1s").Should(ContainSubstring(`"color": "black"`))
			Consistently(CurlReviews, "10s", "0.1s").Should(ContainSubstring(`"color": "black"`))
		})

		By("delete TrafficPolicy should remove traffic shift", func() {
			err = manifest.KubeDelete(BookinfoNamespace)
			Expect(err).NotTo(HaveOccurred())

			Eventually(CurlReviews, "1m", "1s").Should(ContainSubstring(`"color": "black"`))
			Eventually(CurlReviews, "1m", "1s").ShouldNot(ContainSubstring(`"color": "black"`))
		})
	})

	It("disables mTLS for traffic target", func() {
		var getReviewsDestinationRule = func() (*istionetworkingv1alpha3.DestinationRule, error) {
			env := e2e.GetEnv()
			destRuleClient := env.Management.DestinationRuleClient
			meta := metautils.TranslatedObjectMeta(
				&v1.ClusterObjectRef{
					Name:        "reviews",
					Namespace:   BookinfoNamespace,
					ClusterName: MgmtClusterName,
				},
				nil,
			)
			return destRuleClient.GetDestinationRule(ctx, client.ObjectKey{Name: meta.Name, Namespace: meta.Namespace})
		}

		By("initially ensure that DestinationRule exists for mgmt reviews traffic target", func() {
			Eventually(func() *istionetworkingv1alpha3.DestinationRule {
				destRule, err := getReviewsDestinationRule()
				if err != nil {
					return nil
				}
				return destRule
			}, "30s", "1s").ShouldNot(BeNil())
		})

		By("creating TrafficPolicy that overrides default mTLS settings for reviews traffic target", func() {
			tp := &networkingv1.TrafficPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "TrafficPolicy",
					APIVersion: networkingv1.SchemeGroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:        "mtls-disable",
					Namespace:   BookinfoNamespace,
					ClusterName: MgmtClusterName,
				},
				Spec: networkingv1.TrafficPolicySpec{
					Policy: &networkingv1.TrafficPolicySpec_Policy{
						Mtls: &networkingv1.TrafficPolicySpec_Policy_MTLS{
							Istio: &networkingv1.TrafficPolicySpec_Policy_MTLS_Istio{
								TlsMode: networkingv1.TrafficPolicySpec_Policy_MTLS_Istio_DISABLE,
							},
						},
					},
				},
			}
			err = manifest.AppendResources(tp)
			Expect(err).NotTo(HaveOccurred())
			err = manifest.KubeApply(BookinfoNamespace)
			Expect(err).NotTo(HaveOccurred())

			// ensure status is updated
			utils.AssertTrafficPolicyStatuses(ctx, e2e.GetEnv().Management.TrafficPolicyClient, BookinfoNamespace)

			// Check that DestinationRule for reviews no longer exists
			Eventually(func() bool {
				_, err := getReviewsDestinationRule()
				return errors.IsNotFound(err)
			}, "30s", "1s").Should(BeTrue())
		})

		By("first ensure that DestinationRule for mgmt reviews traffic target exists", func() {
			err = manifest.KubeDelete(BookinfoNamespace)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() *istionetworkingv1alpha3.DestinationRule {
				destRule, _ := getReviewsDestinationRule()
				return destRule
			}, "30s", "1s").ShouldNot(BeNil())
		})
	})
}
