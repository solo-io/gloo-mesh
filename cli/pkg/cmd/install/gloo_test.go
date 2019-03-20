package install_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/supergloo/cli/pkg/helpers"
	"github.com/solo-io/supergloo/cli/test/utils"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/test/inputs"
)

var _ = Describe("Install", func() {

	BeforeEach(func() {
		helpers.UseMemoryClients()
	})

	getInstall := func(name string) *v1.Install {
		in, err := helpers.MustInstallClient().Read("supergloo-system", name, clients.ReadOpts{})
		ExpectWithOffset(1, err).NotTo(HaveOccurred())
		return in
	}

	Describe("non-interactive", func() {
		It("should create the expected install ", func() {
			installAndVerifyGloo := func(
				name,
				namespace,
				version string) {

				err := utils.Supergloo("install gloo " +
					fmt.Sprintf("--name=%v ", name) +
					fmt.Sprintf("--installation-namespace %s ", namespace) +
					fmt.Sprintf("--version=%v ", version))
				if version == "badver" {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("is not a supported gloo version"))
					return
				}

				Expect(err).NotTo(HaveOccurred())
				install := getInstall(name)
				glooIngress := MustGlooInstallType(install)
				if version != "latest" {
					Expect(glooIngress.Gloo.GlooVersion).To(Equal(strings.TrimPrefix(version, "v")))
				}
			}

			installAndVerifyGloo("a1a", "ns", "latest")
			installAndVerifyGloo("b1a", "ns", "v0.10.5")
			installAndVerifyGloo("c1a", "ns", "badver")
		})
		It("should enable an existing + disabled install", func() {
			name := "input"
			namespace := "ns"
			inst := inputs.IstioInstall(name, namespace, "any", "1.0.5", true)
			ic := helpers.MustInstallClient()
			_, err := ic.Write(inst, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			err = utils.Supergloo("install gloo " +
				fmt.Sprintf("--name=%v ", name) +
				fmt.Sprintf("--installation-namespace %s ", namespace) +
				fmt.Sprintf("--namespace=%v ", namespace))
			Expect(err).NotTo(HaveOccurred())

			updatedInstall, err := ic.Read(namespace, name, clients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())
			Expect(updatedInstall.Disabled).To(BeFalse())

		})
		It("should error enable on existing enabled install", func() {
			name := "input"
			namespace := "ns"
			inst := inputs.IstioInstall(name, namespace, "any", "1.0.5", false)
			ic := helpers.MustInstallClient()
			_, err := ic.Write(inst, clients.WriteOpts{})
			Expect(err).NotTo(HaveOccurred())

			err = utils.Supergloo("install gloo " +
				fmt.Sprintf("--name=%v ", name) +
				fmt.Sprintf("--installation-namespace %s ", namespace) +
				fmt.Sprintf("--namespace=%v ", namespace))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("already installed and enabled"))
		})
	})
})

func MustGlooInstallType(install *v1.Install) *v1.MeshIngressInstall_Gloo {
	Expect(install.InstallType).To(BeAssignableToTypeOf(&v1.Install_Ingress{}))
	ingress := install.InstallType.(*v1.Install_Ingress)
	Expect(ingress.Ingress.InstallType).To(BeAssignableToTypeOf(&v1.MeshIngressInstall_Gloo{}))
	glooIngress := ingress.Ingress.InstallType.(*v1.MeshIngressInstall_Gloo)
	return glooIngress
}
