package e2e_test

import (
	"context"
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/supergloo/cli/pkg/helpers"
	"github.com/solo-io/supergloo/pkg/setup"
	"github.com/solo-io/supergloo/test/testutils"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var (
	kube                                kubernetes.Interface
	rootCtx                             context.Context
	cancel                              func()
	basicNamespace, namespaceWithInject string
)

var _ = BeforeSuite(func() {
	basicNamespace, namespaceWithInject = "basic-namespace", "namespace-with-inject"
	kube = helpers.MustKubeClient()
	_, err := kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: basicNamespace,
		},
	})
	Expect(err).NotTo(HaveOccurred())

	_, err = kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   namespaceWithInject,
			Labels: map[string]string{"istio-injection": "enabled"},
		},
	})
	Expect(err).NotTo(HaveOccurred())

	rootCtx, cancel = context.WithCancel(context.TODO())
	// create sg ns
	_, err = kube.CoreV1().Namespaces().Create(&kubev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: "supergloo-system"},
	})
	Expect(err).NotTo(HaveOccurred())

	// start supergloo
	go func() {
		defer GinkgoRecover()
		err := setup.Main(rootCtx, func(e error) {
			defer GinkgoRecover()
			return
			Expect(e).NotTo(HaveOccurred())
		})
		Expect(err).NotTo(HaveOccurred())
	}()
})

var _ = AfterSuite(func() {
	if cancel != nil {
		defer cancel()
	}
	kube.CoreV1().Namespaces().Delete("supergloo-system", nil)
	kube.CoreV1().Namespaces().Delete("istio-system", nil)
	kube.CoreV1().Namespaces().Delete(basicNamespace, nil)
	kube.CoreV1().Namespaces().Delete(namespaceWithInject, nil)
	testutils.TeardownIstio(kube)
	log.Printf("done!")
})
