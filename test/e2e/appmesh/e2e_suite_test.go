package appmesh_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/solo-io/supergloo/pkg/constants"

	"github.com/avast/retry-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/testutils/clusterlock"
	sgutils "github.com/solo-io/supergloo/cli/test/utils"
	mdsetup "github.com/solo-io/supergloo/pkg/meshdiscovery/setup"
	"github.com/solo-io/supergloo/pkg/setup"
	"github.com/solo-io/supergloo/pkg/version"
	"github.com/solo-io/supergloo/test/e2e/utils"
	"github.com/solo-io/supergloo/test/testutils"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AWS App Mesh e2e Suite")
}

const superglooNamespace = "supergloo-system"

var (
	kube                                kubernetes.Interface
	lock                                *clusterlock.TestClusterLocker
	rootCtx                             context.Context
	cancel                              func()
	basicNamespace, namespaceWithInject string
	chartUrl                            string
)

var _ = BeforeSuite(func() {
	var err error
	kube = testutils.MustKubeClient()

	// Get build information
	buildVersion, helmChartUrl, imageRepoPrefix, err := utils.GetBuildInformation()
	Expect(err).NotTo(HaveOccurred())

	// Set the supergloo version (will be equal to the BUILD_ID env)
	version.Version = buildVersion
	chartUrl = helmChartUrl

	// Acquire cluster lock
	lock, err = clusterlock.NewTestClusterLocker(kube, clusterlock.Options{
		IdPrefix: os.ExpandEnv("superglooe2e-{$BUILD_ID}-"),
	})
	Expect(err).NotTo(HaveOccurred())
	Expect(lock.AcquireLock(retry.OnRetry(func(n uint, err error) {
		log.Printf("waiting to acquire lock with err: %v", err)
	}))).NotTo(HaveOccurred())

	// If present, delete all namespaces used in this test
	teardown()

	// Create namespaces
	basicNamespace, namespaceWithInject = "basic-namespace", "namespace-with-inject"
	_, err = kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: basicNamespace,
		},
	})
	Expect(err).NotTo(HaveOccurred())

	_, err = kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   namespaceWithInject,
			Labels: map[string]string{"app-mesh-injection": "enabled"},
		},
	})
	Expect(err).NotTo(HaveOccurred())

	_, err = kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: superglooNamespace},
	})
	Expect(err).NotTo(HaveOccurred())

	Expect(os.Setenv(constants.PodNamespaceEnvName, superglooNamespace)).NotTo(HaveOccurred())
	image := fmt.Sprintf("%s/%s:%s", imageRepoPrefix, constants.SidecarInjectorImageName, buildVersion)
	Expect(os.Setenv(constants.SidecarInjectorImageNameEnvName, image)).NotTo(HaveOccurred())
	Expect(os.Setenv(constants.SidecarInjectorImagePullPolicyEnvName, "Always")).NotTo(HaveOccurred())

	// start supergloo (requires setting the envs if running locally)
	rootCtx, cancel = context.WithCancel(context.TODO())
	go func() {
		defer GinkgoRecover()
		err := setup.Main(rootCtx, func(e error) {
			defer GinkgoRecover()
			Expect(e).NotTo(HaveOccurred())
			return
		})
		Expect(err).NotTo(HaveOccurred())
	}()

	// start mesh discovery
	go func() {
		defer GinkgoRecover()
		err := mdsetup.Main(rootCtx, func(e error) {
			defer GinkgoRecover()
			return
			Expect(e).NotTo(HaveOccurred())
		}, nil)
		Expect(err).NotTo(HaveOccurred())
	}()

	// Install supergloo using the helm chart specific to this test run
	superglooErr := sgutils.Supergloo(fmt.Sprintf("init -f %s", chartUrl))
	Expect(superglooErr).NotTo(HaveOccurred())

	// TODO (ilackarms): add a flag to switch between starting supergloo locally and deploying via cli
	testutils.DeleteSuperglooPods(kube, superglooNamespace)

})

var _ = AfterSuite(func() {
	defer lock.ReleaseLock()
	teardown()
})

func teardown() {
	if cancel != nil {
		cancel()
	}
	testutils.TeardownSuperGloo(testutils.MustKubeClient())
	kube.CoreV1().Namespaces().Delete(superglooNamespace, nil)
	kube.CoreV1().Namespaces().Delete(basicNamespace, nil)
	kube.CoreV1().Namespaces().Delete(namespaceWithInject, nil)

	testutils.WaitForNamespaceTeardown(superglooNamespace)
	testutils.WaitForNamespaceTeardown(basicNamespace)
	testutils.WaitForNamespaceTeardown(namespaceWithInject)
	log.Printf("done!")
}
