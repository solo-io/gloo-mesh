package istio

import (
	"context"

	"github.com/solo-io/supergloo/pkg/install/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/test/inputs"
)

type mockIstioInstaller struct {
	enabledInstalls, disabledInstalls v1.InstallList
	errorOnInstall                    bool
}

func (i *mockIstioInstaller) EnsureIstioInstall(ctx context.Context, install *v1.Install, list v1.MeshList) error {
	if i.errorOnInstall {
		return errors.Errorf("i was told to error")
	}
	if install.Disabled {
		i.disabledInstalls = append(i.disabledInstalls, install)
		return nil
	}
	i.enabledInstalls = append(i.enabledInstalls, install)
	return nil
}

type failingMeshClient struct {
	errorOnWrite, errorOnRead, errorOnDelete bool
}

func (c *failingMeshClient) BaseClient() clients.ResourceClient {
	panic("implement me")
}

func (c *failingMeshClient) Register() error {
	panic("implement me")
}

func (c *failingMeshClient) Read(namespace, name string, opts clients.ReadOpts) (*v1.Mesh, error) {
	if c.errorOnRead {
		return nil, errors.Errorf("i was told to error")
	}
	return &v1.Mesh{Metadata: core.Metadata{Name: name, Namespace: namespace}}, nil
}

func (c *failingMeshClient) Write(resource *v1.Mesh, opts clients.WriteOpts) (*v1.Mesh, error) {
	if c.errorOnWrite {
		return nil, errors.Errorf("i was told to error")
	}
	return resource, nil
}

func (c *failingMeshClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	if c.errorOnDelete {
		return errors.Errorf("i was told to error")
	}
	return nil
}

func (c *failingMeshClient) List(namespace string, opts clients.ListOpts) (v1.MeshList, error) {
	panic("implement me")
}

func (c *failingMeshClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1.MeshList, <-chan error, error) {
	panic("implement me")
}

var _ = Describe("Syncer", func() {
	var (
		installer     *mockIstioInstaller
		meshClient    v1.MeshClient
		installClient v1.InstallClient
		report        reporter.Reporter
	)
	Context("happy paths", func() {

		BeforeEach(func() {
			installer = &mockIstioInstaller{}
			meshClient, _ = v1.NewMeshClient(&factory.MemoryResourceClientFactory{
				Cache: memory.NewInMemoryResourceCache(),
			})
			installClient, _ = v1.NewInstallClient(&factory.MemoryResourceClientFactory{
				Cache: memory.NewInMemoryResourceCache(),
			})
			report = reporter.NewReporter("test", installClient.BaseClient())
		})
		Context("multiple active installs", func() {
			It("it reports an error on them, does not call the installer", func() {
				installList := v1.InstallList{
					inputs.IstioInstall("a", "b", "c", "versiondoesntmatter", false),
					inputs.IstioInstall("b", "b", "c", "versiondoesntmatter", false),
				}
				snap := &v1.InstallSnapshot{Installs: map[string]v1.InstallList{"": installList}}
				installeSyncer := newTestInstallSyncer(installer, meshClient, report)
				err := installeSyncer.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())

				Expect(installer.disabledInstalls).To(HaveLen(0))
				Expect(installer.enabledInstalls).To(HaveLen(0))

				i1, err := installClient.Read("b", "a", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred())
				Expect(i1.Status.State).To(Equal(core.Status_Rejected))
				Expect(i1.Status.Reason).To(ContainSubstring("multiple active istio installations are not currently supported"))

				i2, err := installClient.Read("b", "b", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred())
				Expect(i2.Status.State).To(Equal(core.Status_Rejected))
				Expect(i2.Status.Reason).To(ContainSubstring("multiple active istio installations are not currently supported"))
			})
		})
		Context("one active install, one inactive install with a previous install", func() {
			It("it reports success, calls installer, writes the created mesh", func() {
				installList := v1.InstallList{
					inputs.IstioInstall("a", "b", "c", "versiondoesntmatter", true),
					inputs.IstioInstall("b", "b", "c", "versiondoesntmatter", false),
				}
				install := installList[0]
				Expect(install.InstallType).To(BeAssignableToTypeOf(&v1.Install_Mesh{}))
				snap := &v1.InstallSnapshot{Installs: map[string]v1.InstallList{"": installList}}
				installSyncer := newTestInstallSyncer(installer, meshClient, report)
				err := installSyncer.Sync(context.TODO(), snap)
				Expect(err).NotTo(HaveOccurred())

				Expect(installer.disabledInstalls).To(HaveLen(1))
				Expect(installer.enabledInstalls).To(HaveLen(1))

				i1, err := installClient.Read("b", "a", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred())
				Expect(i1.Status.State).To(Equal(core.Status_Accepted))

				i2, err := installClient.Read("b", "b", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred())
				Expect(i2.Status.State).To(Equal(core.Status_Accepted))
			})
		})
	})
	Context("when install fails", func() {
		BeforeEach(func() {
			installer = &mockIstioInstaller{errorOnInstall: true}
			meshClient, _ = v1.NewMeshClient(&factory.MemoryResourceClientFactory{
				Cache: memory.NewInMemoryResourceCache(),
			})
			installClient, _ = v1.NewInstallClient(&factory.MemoryResourceClientFactory{
				Cache: memory.NewInMemoryResourceCache(),
			})
			report = reporter.NewReporter("test", installClient.BaseClient())
		})
		It("it marks the install as rejected", func() {
			installList := v1.InstallList{
				inputs.IstioInstall("a", "b", "c", "versiondoesntmatter", true),
				inputs.IstioInstall("b", "b", "c", "versiondoesntmatter", false),
			}

			snap := &v1.InstallSnapshot{Installs: map[string]v1.InstallList{"": installList}}
			installeSyncer := newTestInstallSyncer(installer, meshClient, report)
			err := installeSyncer.Sync(context.TODO(), snap)
			Expect(err).NotTo(HaveOccurred())

			i1, err := installClient.Read("b", "a", clients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())
			Expect(i1.Status.State).To(Equal(core.Status_Rejected))

			i2, err := installClient.Read("b", "b", clients.ReadOpts{})
			Expect(err).NotTo(HaveOccurred())
			Expect(i2.Status.State).To(Equal(core.Status_Rejected))
			Expect(i2.Status.Reason).To(ContainSubstring("install failed"))
		})
	})
})

func newTestInstallSyncer(istioInstaller Installer, meshClient v1.MeshClient, reporter reporter.Reporter) v1.InstallSyncer {
	return common.NewMeshInstallSyncer("istio", meshClient, reporter, isIstioInstall, istioInstaller.EnsureIstioInstall)
}
