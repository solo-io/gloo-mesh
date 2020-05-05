package k8s_tenancy_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	zephyr_discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	k8s_core_controller "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	mock_controllers "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/compute-target/event-watcher-factories/mocks"
	k8s_tenancy "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/cluster-tenancy/k8s"
	mock_k8s_tenancy "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/cluster-tenancy/k8s/mocks"
	mock_core "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	mock_k8s_core_clients "github.com/solo-io/service-mesh-hub/test/mocks/clients/kubernetes/core/v1"
	mock_zephyr_discovery "github.com/solo-io/service-mesh-hub/test/mocks/zephyr/discovery"
	"github.com/solo-io/skv2/pkg/utils"
	k8s_core "k8s.io/api/core/v1"
	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("ClusterTenancyRegistrarLoop", func() {
	var (
		ctrl                  *gomock.Controller
		ctx                   = context.TODO()
		clusterName           = "test-cluster-name"
		mockTenancyRegistrar  *mock_k8s_tenancy.MockClusterTenancyRegistrar
		mockPodClient         *mock_k8s_core_clients.MockPodClient
		mockMeshClient        *mock_core.MockMeshClient
		mockPodEventWatcher   *mock_controllers.MockPodEventWatcher
		mockMeshEventWatcher  *mock_zephyr_discovery.MockMeshEventWatcher
		podEventHandlerFuncs  k8s_core_controller.PodEventHandler
		meshEventHandlerFuncs zephyr_discovery_controller.MeshEventHandler
		clusterTenancyFinder  k8s_tenancy.ClusterTenancyRegistrarLoop
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockTenancyRegistrar = mock_k8s_tenancy.NewMockClusterTenancyRegistrar(ctrl)
		mockPodClient = mock_k8s_core_clients.NewMockPodClient(ctrl)
		mockMeshClient = mock_core.NewMockMeshClient(ctrl)
		mockPodEventWatcher = mock_controllers.NewMockPodEventWatcher(ctrl)
		mockMeshEventWatcher = mock_zephyr_discovery.NewMockMeshEventWatcher(ctrl)
		clusterTenancyFinder = k8s_tenancy.NewClusterTenancyFinder(
			clusterName,
			[]k8s_tenancy.ClusterTenancyRegistrar{mockTenancyRegistrar},
			mockPodClient,
			mockMeshClient,
		)
		mockPodEventWatcher.
			EXPECT().
			AddEventHandler(ctx, gomock.Any()).
			DoAndReturn(func(ctx context.Context, eventHandlerFuncs k8s_core_controller.PodEventHandler) error {
				podEventHandlerFuncs = eventHandlerFuncs
				return nil
			})
		mockMeshEventWatcher.
			EXPECT().
			AddEventHandler(ctx, gomock.Any()).
			DoAndReturn(func(ctx context.Context, eventHandlerFuncs zephyr_discovery_controller.MeshEventHandler) error {
				meshEventHandlerFuncs = eventHandlerFuncs
				return nil
			})
		err := clusterTenancyFinder.StartRegistration(ctx, mockPodEventWatcher, mockMeshEventWatcher)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	var expectRegisterForPodUpsert = func(pod *k8s_core.Pod) {
		mesh := &zephyr_discovery.Mesh{
			Spec: zephyr_discovery_types.MeshSpec{
				Cluster: &types.ResourceRef{
					Name: clusterName,
				},
			},
		}
		mockTenancyRegistrar.EXPECT().MeshFromSidecar(ctx, pod).Return(mesh, nil)
		mockTenancyRegistrar.EXPECT().RegisterMesh(ctx, clusterName, mesh).Return(nil)
	}

	var expectReconcileTenancyForMeshUpsert = func(mesh *zephyr_discovery.Mesh) {
		podList := &k8s_core.PodList{
			Items: []k8s_core.Pod{
				{ObjectMeta: k8s_meta.ObjectMeta{Name: "1"}},
				{ObjectMeta: k8s_meta.ObjectMeta{Name: "2"}},
			},
		}
		mockPodClient.EXPECT().ListPod(ctx).Return(podList, nil)
		for _, pod := range podList.Items {
			pod := pod
			expectRegisterForPodUpsert(&pod)
		}
	}

	var expectDeregisterMesh = func(mesh *zephyr_discovery.Mesh) {
		mockTenancyRegistrar.EXPECT().DeregisterMesh(ctx, clusterName, mesh).Return(nil)
	}

	var expectReconcileTenancyForCluster = func() {
		meshList := &zephyr_discovery.MeshList{
			Items: []zephyr_discovery.Mesh{
				{
					Spec: zephyr_discovery_types.MeshSpec{
						MeshType: &zephyr_discovery_types.MeshSpec_AwsAppMesh_{
							AwsAppMesh: &zephyr_discovery_types.MeshSpec_AwsAppMesh{
								Clusters: []string{"a", "b", clusterName},
							},
						},
					},
				},
				{
					Spec: zephyr_discovery_types.MeshSpec{
						// This mesh should be ignored
						MeshType: &zephyr_discovery_types.MeshSpec_Istio{},
					},
				},
			},
		}
		mockMeshClient.EXPECT().ListMesh(ctx).Return(meshList, nil)
		expectDeregisterMesh(&meshList.Items[0])
		expectReconcileTenancyForMeshUpsert(&meshList.Items[0])
	}

	It("should reconcile tenancy upon Pod create", func() {
		pod := &k8s_core.Pod{}
		expectRegisterForPodUpsert(pod)
		err := podEventHandlerFuncs.CreatePod(pod)
		Expect(err).To(BeNil())
	})

	It("should reconcile tenancy upon Pod update", func() {
		pod := &k8s_core.Pod{}
		expectRegisterForPodUpsert(pod)
		err := podEventHandlerFuncs.UpdatePod(nil, pod)
		Expect(err).To(BeNil())
	})

	It("should reconcile tenancy for pod delete", func() {
		expectReconcileTenancyForCluster()
		err := podEventHandlerFuncs.DeletePod(&k8s_core.Pod{})
		Expect(err).To(BeNil())
	})

	It("should reconcile cluster tenancy for mesh create", func() {
		meshWithCluster := &zephyr_discovery.Mesh{
			Spec: zephyr_discovery_types.MeshSpec{
				MeshType: &zephyr_discovery_types.MeshSpec_AwsAppMesh_{
					AwsAppMesh: &zephyr_discovery_types.MeshSpec_AwsAppMesh{
						Clusters: []string{"a", "b", clusterName},
					},
				},
			},
		}
		meshWithoutCluster := meshWithCluster.DeepCopy()
		meshWithoutCluster.Spec.GetAwsAppMesh().Clusters = utils.RemoveString(meshWithoutCluster.Spec.GetAwsAppMesh().GetClusters(), clusterName)
		expectDeregisterMesh(meshWithCluster)
		expectReconcileTenancyForMeshUpsert(meshWithoutCluster)
		err := meshEventHandlerFuncs.CreateMesh(meshWithCluster)
		Expect(err).To(BeNil())
	})

	It("should reconcile cluster tenancy for mesh create", func() {
		meshWithCluster := &zephyr_discovery.Mesh{
			Spec: zephyr_discovery_types.MeshSpec{
				MeshType: &zephyr_discovery_types.MeshSpec_AwsAppMesh_{
					AwsAppMesh: &zephyr_discovery_types.MeshSpec_AwsAppMesh{
						Clusters: []string{"a", "b", clusterName},
					},
				},
			},
		}
		meshWithoutCluster := meshWithCluster.DeepCopy()
		meshWithoutCluster.Spec.GetAwsAppMesh().Clusters = utils.RemoveString(meshWithoutCluster.Spec.GetAwsAppMesh().GetClusters(), clusterName)
		expectDeregisterMesh(meshWithCluster)
		expectReconcileTenancyForMeshUpsert(meshWithoutCluster)
		err := meshEventHandlerFuncs.UpdateMesh(nil, meshWithCluster)
		Expect(err).To(BeNil())
	})
})
