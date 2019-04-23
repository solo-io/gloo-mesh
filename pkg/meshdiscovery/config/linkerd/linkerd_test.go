package linkerd

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/supergloo/cli/pkg/helpers/clients"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/clientset"
)

var _ = Describe("istio discovery config", func() {

	var (
		cs  *clientset.Clientset
		ctx context.Context
	)

	BeforeEach(func() {
		var err error
		ctx = context.TODO()
		cs, err = clientset.ClientsetFromContext(ctx)
		Expect(err).NotTo(HaveOccurred())
		clients.UseMemoryClients()
	})

	Context("plugin creation", func() {
		It("can be initialized without an error", func() {
			_, err := NewLinkerdConfigDiscoveryRunner(ctx, cs)
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Context("full mesh", func() {

		var (
			mesh    *v1.Mesh
			install *v1.Install
		)
		BeforeEach(func() {
			mesh = &v1.Mesh{
				MeshType: &v1.Mesh_Linkerd{
					Linkerd: &v1.LinkerdMesh{
						InstallationNamespace: "world",
					},
				},
				MtlsConfig: &v1.MtlsConfig{},
				DiscoveryMetadata: &v1.DiscoveryMetadata{
					InstallationNamespace: "world",
				},
			}
			install = &v1.Install{
				InstallationNamespace: "world",
				InstallType: &v1.Install_Mesh{
					Mesh: &v1.MeshInstall{
						MeshInstallType: &v1.MeshInstall_Linkerd{
							Linkerd: &v1.LinkerdInstall{
								Version:          "2.2.1",
								EnableMtls:       true,
								EnableAutoInject: true,
							},
						},
					},
				},
			}
		})

		It("can organize meshes", func() {
			fullMeshes := organizeMeshes(
				v1.MeshList{mesh},
				v1.InstallList{install},
				nil,
				nil,
			)
			Expect(fullMeshes).To(HaveLen(1))
			Expect(fullMeshes[0].Mesh).To(BeEquivalentTo(mesh))
			Expect(fullMeshes[0].Install).To(BeEquivalentTo(install))
		})
		It("Can merge properly with no install or mesh policy", func() {
			fm := &meshResources{
				Mesh: mesh,
			}
			Expect(fm.merge()).To(BeEquivalentTo(fm.Mesh))
		})
		It("can merge properly with install", func() {
			fm := &meshResources{
				Mesh:    mesh,
				Install: install,
			}
			merge := fm.merge()
			Expect(merge.DiscoveryMetadata.MtlsConfig).To(BeEquivalentTo(&v1.MtlsConfig{
				MtlsEnabled: true,
			}))
			Expect(merge.DiscoveryMetadata).To(BeEquivalentTo(&v1.DiscoveryMetadata{
				MtlsConfig: &v1.MtlsConfig{
					MtlsEnabled: true,
				},
				InstallationNamespace:  "world",
				MeshVersion:            "2.2.1",
				InjectedNamespaceLabel: injectionAnnotation,
				EnableAutoInject:       true,
			}))
		})

	})
})
