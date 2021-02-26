package tests

import (
	"fmt"
	"time"

	settingsv1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/test/extensions"
	"github.com/solo-io/gloo-mesh/test/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func NetworkingExtensionsTest() {
	var (
		err               error
		manifest          utils.Manifest
		glooMeshNamespace = defaults.GetPodNamespace()
	)

	AfterEach(func() {
		manifest, err = utils.NewManifest("default-settings.yaml")
		Expect(err).NotTo(HaveOccurred())
		// update settings to remove our extensions server
		err = manifest.AppendResources(&settingsv1.Settings{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Settings",
				APIVersion: settingsv1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: glooMeshNamespace,
				Name:      "settings", // the default/expected name
			},
		})
		Expect(err).NotTo(HaveOccurred())
		err = manifest.KubeApply(glooMeshNamespace)
		Expect(err).NotTo(HaveOccurred())
	})

	It("enables communication across clusters using global dns names", func() {
		manifest, err = utils.NewManifest("extension-settings.yaml")
		Expect(err).NotTo(HaveOccurred())

		By("with extensions enabled, additional configs can be added to GlooMesh outputs", func() {

			helloMsg := "hello from a 3rd party"

			srv := extensions.NewTestExtensionsServer()

			// run extensions server
			go func() {
				defer GinkgoRecover()
				err := srv.Run()
				Expect(err).NotTo(HaveOccurred())
			}()
			// run hello server
			go func() {
				defer GinkgoRecover()
				err := extensions.RunHelloServer(helloMsg)
				Expect(err).NotTo(HaveOccurred())
			}()

			// update settings to connect our extensions server
			err = manifest.AppendResources(&settingsv1.Settings{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Settings",
					APIVersion: settingsv1.SchemeGroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace: glooMeshNamespace,
					Name:      "settings", // the default/expected name
				},
				Spec: settingsv1.SettingsSpec{
					NetworkingExtensionServers: []*settingsv1.GrpcServer{{
						// use the machine's docker host address
						Address:                    fmt.Sprintf("%v:%v", extensions.DockerHostAddress, extensions.ExtensionsServerPort),
						Insecure:                   true,
						ReconnectOnNetworkFailures: true,
					}},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			err = manifest.KubeApply(glooMeshNamespace)
			Expect(err).NotTo(HaveOccurred())

			// ensure the server eventually connects to us
			Eventually(srv.HasConnected, time.Minute*2).Should(BeTrue())

			// check we can eventually hit the echo server via the gateway.
			// This request verifies that Envoy has config provided by Service Entries from our test extensions server.
			Eventually(CurlHelloServer, "30s", "1s").Should(ContainSubstring(helloMsg))
		})
	})
}
