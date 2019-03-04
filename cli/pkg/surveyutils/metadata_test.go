package surveyutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/cliutil/testutil"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/cli/pkg/helpers"
	. "github.com/solo-io/supergloo/cli/pkg/surveyutils"
)

var _ = Describe("Metadata", func() {
	It("should create the expected install ", func() {
		name, namespace := "hi", helpers.MustGetNamespaces()[1]
		testutil.ExpectInteractive(func(c *testutil.Console) {
			c.ExpectString("name for the resource: ")
			c.SendLine(name)
			c.ExpectString("namespace for the resource: ")
			c.PressDown()
			c.SendLine("")
			c.SendLine("")
			c.ExpectEOF()
		}, func() {
			var meta core.Metadata
			err := SurveyMetadata("test", &meta)
			Expect(err).NotTo(HaveOccurred())
			Expect(meta.Name).To(Equal(name))
			Expect(meta.Namespace).To(Equal(namespace))
		})
	})
})
