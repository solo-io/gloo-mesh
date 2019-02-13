package shared_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/prometheus/prometheus/config"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/tests/typed"
	"github.com/solo-io/supergloo/pkg2/api/external/prometheus"
	prometheusv1 "github.com/solo-io/supergloo/pkg2/api/external/prometheus/v1"
	v1 "github.com/solo-io/supergloo/pkg2/api/v1"
	"github.com/solo-io/supergloo/pkg2/translator/istio"
	"github.com/solo-io/supergloo/pkg2/translator/linkerd2"
	. "github.com/solo-io/supergloo/pkg2/translator/shared"
	"github.com/solo-io/supergloo/test/utils"
	"k8s.io/client-go/kubernetes"
)

var _ = Describe("PrometheusSyncer", func() {
	type test struct {
		meshType      string
		scrapeConfigs []*config.ScrapeConfig
	}

	tester := &typed.KubeConfigMapRcTester{}
	var (
		namespace                string
		kube                     kubernetes.Interface
		prometheusConfigName     = "prometheus"
		prometheusDeploymentName = "prometheus"
	)
	BeforeEach(func() {
		namespace = helpers.RandString(6)
		fact := tester.Setup(namespace)
		kube = fact.(*factory.KubeConfigMapClientFactory).Clientset
	})
	AfterEach(func() {
		tester.Teardown(namespace)
	})
	table.DescribeTable("prometheus tests for various meshes",
		func(port int, testCase test) {
			err := utils.DeployPrometheus(namespace, prometheusDeploymentName, prometheusConfigName, uint32(port), kube)
			Expect(err).NotTo(HaveOccurred())
			err = utils.DeployPrometheusConfigmap(namespace, prometheusConfigName, kube)
			Expect(err).NotTo(HaveOccurred())
			prometheusClient, err := prometheusv1.NewPrometheusConfigClient(&factory.KubeConfigMapClientFactory{
				Clientset: kube,
			})
			Expect(err).NotTo(HaveOccurred())
			err = prometheusClient.Register()
			Expect(err).NotTo(HaveOccurred())
			s := &PrometheusSyncer{
				PrometheusClient:     prometheusClient,
				Kube:                 kube,
				DesiredScrapeConfigs: testCase.scrapeConfigs,
				GetConfigMap: func(mesh *v1.Mesh) *core.ResourceRef {
					return &core.ResourceRef{
						Namespace: namespace,
						Name:      prometheusConfigName,
					}
				},
			}
			original := getPrometheusConfig(prometheusClient, namespace, prometheusConfigName)
			for _, sc := range s.DesiredScrapeConfigs {
				Expect(original.ScrapeConfigs).NotTo(ContainElement(sc))
			}

			err = s.Sync(context.TODO(), &v1.TranslatorSnapshot{
				Meshes: map[string]v1.MeshList{
					"ignored-at-this-point": {{
						Observability: &v1.Observability{
							Prometheus: &v1.Prometheus{
								EnableMetrics: true,
								PodLabels: map[string]string{
									"app": "prometheus-server",
								},
							},
						},
					}},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			updated := getPrometheusConfig(prometheusClient, namespace, prometheusConfigName)
			for _, sc := range s.DesiredScrapeConfigs {
				Expect(updated.ScrapeConfigs).To(ContainElement(sc))
			}
		},
		table.Entry("istio", 31000, test{
			scrapeConfigs: istio.IstioScrapeConfigs,
		}),
		table.Entry("linkerd2", 31001, test{
			scrapeConfigs: linkerd2.LinkerdScrapeConfigs,
		}),
	)
})

func getPrometheusConfig(promClient prometheusv1.PrometheusConfigClient, namespace, name string) *prometheus.Config {
	cfg, err := promClient.Read(namespace, name, clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())
	promCfg, err := prometheus.ConfigFromResource(cfg)
	Expect(err).NotTo(HaveOccurred())
	return promCfg
}
