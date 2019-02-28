package testutils

import (
	"time"

	. "github.com/onsi/gomega"
	kubev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var istioInstalledCrds = []string{
	"meshpolicies.authentication.istio.io",
	"policies.authentication.istio.io",
}

func installedByIstio(crdName string) bool {
	for _, n := range istioInstalledCrds {
		if crdName == n {
			return true
		}
	}
	return false
}

func WaitForIstioTeardown(ns string) {
	EventuallyWithOffset(1, func() []kubev1.Service {
		svcs, err := MustKubeClient().CoreV1().Services(ns).List(v1.ListOptions{})
		if err != nil {
			// namespace is gone
			return []kubev1.Service{}
		}
		return svcs.Items
	}, time.Second*30).Should(BeEmpty())
}
