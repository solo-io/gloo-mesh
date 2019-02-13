package helm_test

import (
	"context"

	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/testutils"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/supergloo/pkg2/install/helm"
)

var istioCrd = apiextensions.CustomResourceDefinition{}

var _ = Describe("HelmChartInstaller", func() {
	var ns string
	BeforeEach(func() {
		ns = "test" + testutils.RandString(5)
	})
	AfterEach(func() {
		testutils.TeardownKube(ns)
	})
	Context("create manifest", func() {
		It("creates resources from a helm chart", func() {
			values := `
mixer:
  enabled: true #should install mixer

`
			manifests, err := RenderManifests(
				context.TODO(),
				"https://s3.amazonaws.com/supergloo.solo.io/istio-1.0.3.tgz",
				values,
				"yella",
				ns,
				"",
				true,
			)
			defer DeleteFromManifests(context.TODO(), ns, manifests)
			Expect(err).NotTo(HaveOccurred())
			err = CreateFromManifests(context.TODO(), ns, manifests)
			Expect(err).NotTo(HaveOccurred())

			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())

			kubeClient, err := kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())

			// yes mixer
			_, err = kubeClient.AppsV1().Deployments(ns).Get("istio-policy", v1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			_, err = kubeClient.AppsV1().Deployments(ns).Get("istio-telemetry", v1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())

			err = DeleteFromManifests(context.TODO(), ns, manifests)
			Expect(err).NotTo(HaveOccurred())
		})
		It("handles value overrides correctly", func() {
			values := `
mixer:
  enabled: false #should not install mixer

`
			manifests, err := RenderManifests(
				context.TODO(),
				"https://s3.amazonaws.com/supergloo.solo.io/istio-1.0.3.tgz",
				values,
				"yella",
				ns,
				"",
				true,
			)
			Expect(err).NotTo(HaveOccurred())
			defer DeleteFromManifests(context.TODO(), ns, manifests)

			// no security crds
			for _, man := range manifests {
				Expect(man.Content).NotTo(ContainSubstring("policies.authentication.istio.io"))
			}

			err = CreateFromManifests(context.TODO(), ns, manifests)
			Expect(err).NotTo(HaveOccurred())

			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())

			kubeClient, err := kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())

			// no mixer
			_, err = kubeClient.AppsV1().Deployments(ns).Get("mixer", v1.GetOptions{})
			_, err = kubeClient.AppsV1().Deployments(ns).Get("istio-policy", v1.GetOptions{})
			Expect(err).To(HaveOccurred())
			_, err = kubeClient.AppsV1().Deployments(ns).Get("istio-telemetry", v1.GetOptions{})
			Expect(err).To(HaveOccurred())

			err = DeleteFromManifests(context.TODO(), ns, manifests)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
