package istio

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/cli/pkg/helpers/clients"
	"github.com/solo-io/supergloo/pkg/api/custom/clients/kubernetes"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("istio mesh discovery unit tests", func() {
	var (
		istioNamespace     = "istio-system"
		superglooNamespace = "supergloo-system"
	)

	var constructPod = func(container kubev1.Container, namespace string) *v1.Pod {

		pod := &kubev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      "istio-pilot",
			},
			Spec: kubev1.PodSpec{
				Containers: []kubev1.Container{
					container,
				},
			},
		}
		return kubernetes.FromKube(pod)
	}

	BeforeEach(func() {
		clients.UseMemoryClients()
	})
	Context("get version from pod", func() {
		It("errors when no pilot container is found", func() {
			container := kubev1.Container{
				Image: "istio-",
			}
			pod := constructPod(container, istioNamespace)
			_, err := getVersionFromPod(pod)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unable to find pilot container from pod"))
		})
		It("errors when no version is found in image name", func() {
			container := kubev1.Container{
				Image: "istio-pilot",
			}
			pod := constructPod(container, istioNamespace)
			_, err := getVersionFromPod(pod)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unable to find image version for image"))
		})
		It("fails when image is the incorrect format", func() {
			container := kubev1.Container{
				Image: "istio-pilot:10.6",
			}
			pod := constructPod(container, istioNamespace)
			_, err := getVersionFromPod(pod)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unable to find image version for image"))
		})
		It("errors when no version is found in image name", func() {
			container := kubev1.Container{
				Image: "istio-pilot:1.0.6",
			}
			pod := constructPod(container, istioNamespace)
			version, err := getVersionFromPod(pod)
			Expect(err).NotTo(HaveOccurred())
			Expect(version).To(Equal("1.0.6"))
		})
	})
	Context("discovery data", func() {
		It("can properly construct the discovery data", func() {
			container := kubev1.Container{
				Image: "istio-pilot:1.0.6",
			}
			pod := constructPod(container, istioNamespace)
			mesh, err := constructDiscoveredMesh(context.TODO(), pod, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(mesh.Metadata).To(BeEquivalentTo(core.Metadata{
				Labels:    DiscoverySelector,
				Namespace: superglooNamespace,
				Name:      fmt.Sprintf("istio-%s", istioNamespace),
			}))
		})
		It("overwrites discovery data with install info", func() {
			container := kubev1.Container{
				Image: "istio-pilot:1.0.6",
			}
			pod := constructPod(container, istioNamespace)
			helloWorldCert := &core.ResourceRef{
				Namespace: "hello",
				Name:      "world",
			}
			installMeta := core.Metadata{
				Namespace: superglooNamespace,
				Name:      "my-istio",
			}
			installs := v1.InstallList{
				{
					Metadata:              installMeta,
					InstallationNamespace: istioNamespace,
					InstallType: &v1.Install_Mesh{
						Mesh: &v1.MeshInstall{
							MeshInstallType: &v1.MeshInstall_IstioMesh{
								IstioMesh: &v1.IstioInstall{
									CustomRootCert: helloWorldCert,
									EnableMtls:     true,
								},
							},
						},
					},
				},
			}
			mesh, err := constructDiscoveredMesh(context.TODO(), pod, installs)
			Expect(err).NotTo(HaveOccurred())
			Expect(mesh.MtlsConfig).To(BeEquivalentTo(&v1.MtlsConfig{
				RootCertificate: helloWorldCert,
				MtlsEnabled:     true,
			}))
			Expect(mesh.Metadata).To(BeEquivalentTo(core.Metadata{
				Labels:    DiscoverySelector,
				Namespace: installMeta.Namespace,
				Name:      installMeta.Name,
			}))
		})
	})
})
