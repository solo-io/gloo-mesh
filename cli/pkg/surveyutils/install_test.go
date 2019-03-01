package surveyutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/cliutil/testutil"
	"github.com/solo-io/supergloo/cli/pkg/helpers"
	"github.com/solo-io/supergloo/cli/pkg/options"
	. "github.com/solo-io/supergloo/cli/pkg/surveyutils"
	"github.com/solo-io/supergloo/pkg/install/istio"
)

var _ = Describe("Metadata", func() {
	It("should create the expected install ", func() {
		namespace := helpers.MustGetNamespaces()[1]

		testutil.ExpectInteractive(func(c *testutil.Console) {
			c.ExpectString("which namespace to install to? ")
			c.PressDown()
			c.SendLine("")
			c.ExpectString("which version of Istio to install? ")
			c.PressDown()
			c.SendLine("")
			c.ExpectString("enable mtls? ")
			c.SendLine("y")
			c.ExpectString("enable auto-injection? ")
			c.SendLine("y")
			c.ExpectString("add grafana to the install? ")
			c.SendLine("y")
			c.ExpectString("add prometheus to the install? ")
			c.SendLine("y")
			c.ExpectString("add jaeger to the install? ")
			c.SendLine("y")
			c.ExpectEOF()
		}, func() {
			var in options.InputInstall
			err := SurveyIstioInstall(&in)
			Expect(err).NotTo(HaveOccurred())
			Expect(in.IstioInstall.InstallationNamespace).To(Equal(namespace))
			Expect(in.IstioInstall.IstioVersion).To(Equal(istio.IstioVersion105))
			Expect(in.IstioInstall.EnableMtls).To(Equal(true))
			Expect(in.IstioInstall.EnableAutoInject).To(Equal(true))
			Expect(in.IstioInstall.InstallGrafana).To(Equal(true))
			Expect(in.IstioInstall.InstallPrometheus).To(Equal(true))
			Expect(in.IstioInstall.InstallJaeger).To(Equal(true))
		})
	})
})
