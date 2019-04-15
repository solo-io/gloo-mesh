package linkerd_test

import (
	"fmt"
	"time"

	skerrors "github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/supergloo/cli/pkg/helpers/clients"
	sgtestutils "github.com/solo-io/supergloo/test/testutils"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	skclients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/cli/test/utils"
	sgutils "github.com/solo-io/supergloo/test/e2e/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const superglooNamespace = "supergloo-system"

var _ = Describe("linkerd e2e", func() {
	meshName := "my-linkerd"

	It("it installs linkerd", func() {
		testInstallLinkerd(meshName)
	})
	It("it uninstalls linkerd", func() {
		testUninstallLinkerd(meshName)
	})
})

/*
   tests
*/
func testInstallLinkerd(meshName string) {
	err := utils.Supergloo(fmt.Sprintf("install linkerd --name=%v --mtls=true --auto-inject=true", meshName))
	Expect(err).NotTo(HaveOccurred())

	installClient := clients.MustInstallClient()

	Eventually(func() (core.Status_State, error) {
		i, err := installClient.Read(superglooNamespace, meshName, skclients.ReadOpts{})
		if err != nil {
			return 0, err
		}
		Expect(i.Status.Reason).To(Equal(""))
		return i.Status.State, nil
	}, time.Minute*2).Should(Equal(core.Status_Accepted))

	Eventually(func() error {
		_, err := kube.CoreV1().Services(linkerdNamesapce).Get("linkerd-controller-api", metav1.GetOptions{})
		return err
	}).ShouldNot(HaveOccurred())

	meshClient := clients.MustMeshClient()
	Eventually(func() error {
		_, err := meshClient.Read(superglooNamespace, meshName, skclients.ReadOpts{})
		return err
	}).ShouldNot(HaveOccurred())

	err = sgtestutils.WaitUntilPodsRunning(time.Minute*2, linkerdNamesapce,
		"linkerd-controller",
		"linkerd-web",
		"linkerd-prometheus",
		"linkerd-grafana",
		"linkerd-ca",
		"linkerd-proxy-injector",
	)
	Expect(err).NotTo(HaveOccurred())

	err = sgutils.DeployTestRunner(basicNamespace)
	Expect(err).NotTo(HaveOccurred())

	// the sidecar injector might take some time to become available
	Eventually(func() error {
		return sgutils.DeployTestRunner(namespaceWithInject)
	}, time.Minute*1).ShouldNot(HaveOccurred())

	err = sgutils.DeployBookInfoIstio(namespaceWithInject)
	Expect(err).NotTo(HaveOccurred())

	err = sgtestutils.WaitUntilPodsRunning(time.Minute*4, basicNamespace,
		"testrunner",
	)
	Expect(err).NotTo(HaveOccurred())

	err = sgtestutils.WaitUntilPodsRunning(time.Minute*2, namespaceWithInject,
		"testrunner",
		"reviews-v1",
		"reviews-v2",
		"reviews-v3",
	)
	Expect(err).NotTo(HaveOccurred())
}

func testUninstallLinkerd(meshName string) {
	// test uninstall works
	err := utils.Supergloo("uninstall --name=" + meshName)
	Expect(err).NotTo(HaveOccurred())

	err = nil
	Eventually(func() error {
		_, err = kube.CoreV1().Services(linkerdNamesapce).Get("linkerd-controller-api", metav1.GetOptions{})
		return err
	}, time.Minute*2).Should(HaveOccurred())
	Expect(kubeerrs.IsNotFound(err)).To(BeTrue())

	err = nil
	Eventually(func() bool {
		_, err = clients.MustMeshClient().Read(superglooNamespace, meshName, skclients.ReadOpts{})
		if err == nil {
			return false
		}
		return skerrors.IsNotExist(err)
	}, time.Minute*2).Should(BeTrue())
}
