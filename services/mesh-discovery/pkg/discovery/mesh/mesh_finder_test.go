package mesh_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-multierror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	mp_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/clients"
	mesh_discovery "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh"
	mock_discovery "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh/mocks"
	mock_core "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	mock_k8s_apps_clients "github.com/solo-io/service-mesh-hub/test/mocks/clients/kubernetes/apps/v1"
	mock_controller_runtime "github.com/solo-io/service-mesh-hub/test/mocks/controller-runtime"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BuildDeployment(objMeta metav1.ObjectMeta) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: objMeta,
	}
}

func BuildMesh(objMeta metav1.ObjectMeta) *mp_v1alpha1.Mesh {
	return &mp_v1alpha1.Mesh{
		ObjectMeta: objMeta,
	}
}

var _ = Describe("Mesh Finder", func() {
	var (
		ctrl            *gomock.Controller
		ctx             = context.TODO()
		clusterName     = "cluster-name"
		remoteNamespace = "remote-namespace"
		clusterClient   *mock_controller_runtime.MockClient
		testErr         = eris.New("test-err")
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		clusterClient = mock_controller_runtime.NewMockClient(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Create Event", func() {
		It("can discover a mesh", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})
			meshObjectMeta := metav1.ObjectMeta{Name: "test-mesh", Namespace: remoteNamespace}
			mesh := BuildMesh(meshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(mesh, nil)

			localMeshClient.
				EXPECT().
				UpsertMeshSpec(ctx, mesh).
				Return(nil)

			err := eventHandler.CreateDeployment(deployment)
			Expect(err).NotTo(HaveOccurred())
		})

		It("can go on to discover a mesh if one of the other finders errors out", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			brokenMeshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})
			meshObjectMeta := metav1.ObjectMeta{Name: "test-mesh", Namespace: remoteNamespace}
			mesh := BuildMesh(meshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{brokenMeshFinder, meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(mesh, nil)

			localMeshClient.
				EXPECT().
				UpsertMeshSpec(ctx, mesh).
				Return(nil)

			brokenMeshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(nil, testErr)

			err := eventHandler.CreateDeployment(deployment)
			Expect(err).NotTo(HaveOccurred())

		})

		It("responds with an error if no mesh was found and the finders reported an error", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(nil, testErr)

			err := eventHandler.CreateDeployment(deployment)
			multierr, ok := err.(*multierror.Error)
			Expect(ok).To(BeTrue())
			Expect(multierr.Errors).To(HaveLen(1))
			Expect(multierr.Errors[0]).To(testutils.HaveInErrorChain(testErr))

		})

		It("doesn't do anything if no mesh was discovered", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(nil, nil)

			err := eventHandler.CreateDeployment(deployment)
			Expect(err).NotTo(HaveOccurred())

		})

		It("performs an upsert if we discovered a mesh that we discovered previously", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})
			meshObjectMeta := metav1.ObjectMeta{Name: "test-mesh", Namespace: remoteNamespace}
			mesh := BuildMesh(meshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(mesh, nil)

			localMeshClient.
				EXPECT().
				UpsertMeshSpec(ctx, mesh).
				Return(nil)

			err := eventHandler.CreateDeployment(deployment)
			Expect(err).NotTo(HaveOccurred())

		})

		It("returns error from Upsert if upsert fails", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})
			meshObjectMeta := metav1.ObjectMeta{Name: "test-mesh", Namespace: remoteNamespace}
			mesh := BuildMesh(meshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(mesh, nil)

			localMeshClient.
				EXPECT().
				UpsertMeshSpec(ctx, mesh).
				Return(testErr)

			err := eventHandler.CreateDeployment(deployment)
			Expect(err).To(Equal(testErr))

		})
	})

	Context("Update Event", func() {
		It("can discover a mesh", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			newDeployment := BuildDeployment(metav1.ObjectMeta{Name: "new-deployment", Namespace: remoteNamespace})
			meshObjectMeta := metav1.ObjectMeta{Name: "test-mesh", Namespace: remoteNamespace}
			mesh := BuildMesh(meshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, newDeployment, clusterClient).
				Return(mesh, nil)

			localMeshClient.
				EXPECT().
				UpsertMeshSpec(ctx, mesh).
				Return(nil)

			err := eventHandler.UpdateDeployment(nil, newDeployment)
			Expect(err).NotTo(HaveOccurred())

		})

		It("doesn't do anything if no mesh was discovered", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			newDeployment := BuildDeployment(metav1.ObjectMeta{Name: "new-deployment", Namespace: remoteNamespace})

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, newDeployment, clusterClient).
				Return(nil, nil)

			err := eventHandler.UpdateDeployment(nil, newDeployment)
			Expect(err).NotTo(HaveOccurred())

		})

		It("writes a new CR if an update event changes the mesh type", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			newDeployment := BuildDeployment(metav1.ObjectMeta{Name: "new-deployment", Namespace: remoteNamespace})
			newMeshObjectMeta := metav1.ObjectMeta{Name: "new-test-mesh", Namespace: remoteNamespace}
			newMesh := BuildMesh(newMeshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)
			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, newDeployment, clusterClient).
				Return(newMesh, nil)
			localMeshClient.
				EXPECT().
				UpsertMeshSpec(ctx, newMesh).
				Return(nil)
			err := eventHandler.UpdateDeployment(nil, newDeployment)
			Expect(err).NotTo(HaveOccurred())

		})
	})

	Context("Delete Event", func() {
		It("can delete a mesh when appropriate", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})
			meshObjectMeta := metav1.ObjectMeta{Name: "test-mesh", Namespace: remoteNamespace}
			mesh := BuildMesh(meshObjectMeta)

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(mesh, nil)

			localMeshClient.
				EXPECT().
				DeleteMesh(ctx, clients.ObjectMetaToObjectKey(mesh.ObjectMeta)).
				Return(nil)

			err := eventHandler.DeleteDeployment(deployment)
			Expect(err).NotTo(HaveOccurred())
		})

		It("does not delete a mesh if the deployment is not a control plane", func() {
			meshFinder := mock_discovery.NewMockMeshScanner(ctrl)
			localMeshClient := mock_core.NewMockMeshClient(ctrl)
			deploymentClient := mock_k8s_apps_clients.NewMockDeploymentClient(ctrl)
			deployment := BuildDeployment(metav1.ObjectMeta{Name: "test-deployment", Namespace: remoteNamespace})

			eventHandler := mesh_discovery.NewMeshFinder(
				ctx,
				clusterName,
				[]mesh_discovery.MeshScanner{meshFinder},
				localMeshClient,
				clusterClient,
				deploymentClient,
			)

			meshFinder.
				EXPECT().
				ScanDeployment(ctx, clusterName, deployment, clusterClient).
				Return(nil, nil)

			err := eventHandler.DeleteDeployment(deployment)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
