package zephyr_discovery_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_core "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/discovery"
	mock_controller_runtime "github.com/solo-io/service-mesh-hub/test/mocks/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("MeshWorkloadClient", func() {
	var (
		meshWorkloadClient zephyr_core.MeshWorkloadClient
		mockKubeClient     *mock_controller_runtime.MockClient
		ctx                context.Context
		ctrl               *gomock.Controller
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockKubeClient = mock_controller_runtime.NewMockClient(ctrl)
		meshWorkloadClient = zephyr_core.NewMeshWorkloadClient(mockKubeClient)
		ctx = context.TODO()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should get mesh workload", func() {
		objectKey := client.ObjectKey{Namespace: "foo", Name: "bar"}
		mockKubeClient.EXPECT().Get(ctx, objectKey, &v1alpha1.MeshWorkload{}).Return(nil)
		_, err := meshWorkloadClient.Get(ctx, objectKey)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return error", func() {
		objectKey := client.ObjectKey{Namespace: "foo", Name: "bar"}
		mockKubeClient.EXPECT().Get(ctx, objectKey, &v1alpha1.MeshWorkload{}).Return(errors.New(""))
		_, err := meshWorkloadClient.Get(ctx, objectKey)
		Expect(err).To(HaveOccurred())
	})
})
