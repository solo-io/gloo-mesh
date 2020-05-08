package get_service_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	mock_table_printing "github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing/mocks"
	cli_test "github.com/solo-io/service-mesh-hub/cli/pkg/test"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	mock_kubeconfig "github.com/solo-io/service-mesh-hub/pkg/kubeconfig/mocks"
	mock_zephyr_discovery "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Get Service Cmd", func() {
	var (
		ctrl               *gomock.Controller
		ctx                context.Context
		meshctl            *cli_test.MockMeshctl
		mockKubeLoader     *mock_kubeconfig.MockKubeLoader
		mockServicePrinter *mock_table_printing.MockMeshServicePrinter
		mockServiceClient  *mock_zephyr_discovery.MockMeshServiceClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockKubeLoader = mock_kubeconfig.NewMockKubeLoader(ctrl)
		mockServicePrinter = mock_table_printing.NewMockMeshServicePrinter(ctrl)
		mockServiceClient = mock_zephyr_discovery.NewMockMeshServiceClient(ctrl)
		meshctl = &cli_test.MockMeshctl{
			MockController: ctrl,
			Ctx:            ctx,
			Clients:        common.Clients{},
			KubeClients: common.KubeClients{
				MeshServiceClient: mockServiceClient,
			},
			KubeLoader: mockKubeLoader,
			Printers: common.Printers{
				MeshServicePrinter: mockServicePrinter,
			},
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will call the mesh service printer with the proper data", func() {

		meshServices := []*zephyr_discovery.MeshService{
			{
				ObjectMeta: v1.ObjectMeta{
					Name: "mesh-1",
				},
			},
			{
				ObjectMeta: v1.ObjectMeta{
					Name: "mesh-2",
				},
			},
		}
		mockKubeLoader.EXPECT().
			GetRestConfigForContext("", "").
			Return(nil, nil)
		mockServiceClient.EXPECT().
			ListMeshService(ctx).
			Return(&zephyr_discovery.MeshServiceList{
				Items: []zephyr_discovery.MeshService{*meshServices[0], *meshServices[1]},
			}, nil)
		mockServicePrinter.EXPECT().
			Print(gomock.Any(), meshServices).
			Return(nil)
		_, err := meshctl.Invoke("get services")
		Expect(err).NotTo(HaveOccurred())
	})
})
