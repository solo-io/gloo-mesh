// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1   string
		namespace2   string
		name1, name2 = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg          *rest.Config
		emitter      RegistrationEmitter
		meshClient   MeshClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace2)
		Expect(err).NotTo(HaveOccurred())
		// Mesh Constructor
		meshClientFactory := &factory.KubeResourceClientFactory{
			Crd:         MeshCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(context.TODO()),
		}
		meshClient, err = NewMeshClient(meshClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewRegistrationEmitter(meshClient)
	})
	AfterEach(func() {
		setup.TeardownKube(namespace1)
		setup.TeardownKube(namespace2)
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *RegistrationSnapshot

		/*
			Mesh
		*/

		assertSnapshotMeshes := func(expectMeshes MeshList, unexpectMeshes MeshList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMeshes {
						if _, err := snap.Meshes.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMeshes {
						if _, err := snap.Meshes.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := meshClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := meshClient.List(namespace2, clients.ListOpts{})
					combined := MeshesByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		mesh1a, err := meshClient.Write(NewMesh(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh1b, err := meshClient.Write(NewMesh(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, nil)
		mesh2a, err := meshClient.Write(NewMesh(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh2b, err := meshClient.Write(NewMesh(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b, mesh2a, mesh2b}, nil)

		err = meshClient.Delete(mesh2a.Metadata.Namespace, mesh2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh2b.Metadata.Namespace, mesh2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, MeshList{mesh2a, mesh2b})

		err = meshClient.Delete(mesh1a.Metadata.Namespace, mesh1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh1b.Metadata.Namespace, mesh1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(nil, MeshList{mesh1a, mesh1b, mesh2a, mesh2b})
	})
	It("tracks snapshots on changes to any resource using AllNamespace", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{""}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *RegistrationSnapshot

		/*
			Mesh
		*/

		assertSnapshotMeshes := func(expectMeshes MeshList, unexpectMeshes MeshList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectMeshes {
						if _, err := snap.Meshes.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectMeshes {
						if _, err := snap.Meshes.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := meshClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := meshClient.List(namespace2, clients.ListOpts{})
					combined := MeshesByNamespace{
						namespace1: nsList1,
						namespace2: nsList2,
					}
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		mesh1a, err := meshClient.Write(NewMesh(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh1b, err := meshClient.Write(NewMesh(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, nil)
		mesh2a, err := meshClient.Write(NewMesh(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		mesh2b, err := meshClient.Write(NewMesh(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b, mesh2a, mesh2b}, nil)

		err = meshClient.Delete(mesh2a.Metadata.Namespace, mesh2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh2b.Metadata.Namespace, mesh2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(MeshList{mesh1a, mesh1b}, MeshList{mesh2a, mesh2b})

		err = meshClient.Delete(mesh1a.Metadata.Namespace, mesh1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = meshClient.Delete(mesh1b.Metadata.Namespace, mesh1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotMeshes(nil, MeshList{mesh1a, mesh1b, mesh2a, mesh2b})
	})
})
