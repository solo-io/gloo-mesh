package cert_manager_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/solo-io/go-utils/contextutils"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_networking_types "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	networking_snapshot "github.com/solo-io/service-mesh-hub/pkg/networking-snapshot"
	cert_manager "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/security/cert-manager"
	mock_cert_manager "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/security/cert-manager/mocks"
	test_logging "github.com/solo-io/service-mesh-hub/test/logging"
	mock_zephyr_networking "github.com/solo-io/service-mesh-hub/test/mocks/clients/networking.zephyr.solo.io/v1alpha1"
	"go.uber.org/zap/zapcore"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("snapshot listener", func() {
	var (
		ctrl                *gomock.Controller
		ctx                 context.Context
		testLogger          *test_logging.TestLogger
		csrProcessor        *mock_cert_manager.MockVirtualMeshCertificateManager
		virtualMeshClient   *mock_zephyr_networking.MockVirtualMeshClient
		csrSnapshotListener cert_manager.VMCSRSnapshotListener
	)

	BeforeEach(func() {
		testLogger = test_logging.NewTestLogger()
		ctx = contextutils.WithExistingLogger(context.TODO(), testLogger.Logger())
		ctrl = gomock.NewController(GinkgoT())
		csrProcessor = mock_cert_manager.NewMockVirtualMeshCertificateManager(ctrl)
		virtualMeshClient = mock_zephyr_networking.NewMockVirtualMeshClient(ctrl)
		csrSnapshotListener = cert_manager.NewVMCSRSnapshotListener(csrProcessor, virtualMeshClient)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will do nothing if there are no updated virtual meshes", func() {
		snap := &networking_snapshot.MeshNetworkingSnapshot{}
		csrSnapshotListener.Sync(ctx, snap)
		testLogger.EXPECT().
			LastEntry().
			Level(zapcore.DebugLevel).
			HaveMessage(cert_manager.NoVirtualMeshesChangedMessage)
	})

	It("will process all create events in order", func() {
		vm1 := &zephyr_networking.VirtualMesh{
			TypeMeta:   k8s_meta_types.TypeMeta{},
			ObjectMeta: k8s_meta_types.ObjectMeta{},
			Spec:       zephyr_networking_types.VirtualMeshSpec{},
			Status: zephyr_networking_types.VirtualMeshStatus{
				CertificateStatus: &zephyr_core_types.Status{
					State: zephyr_core_types.Status_ACCEPTED,
				},
			},
		}
		vm2 := &zephyr_networking.VirtualMesh{
			TypeMeta:   k8s_meta_types.TypeMeta{},
			ObjectMeta: k8s_meta_types.ObjectMeta{},
			Spec:       zephyr_networking_types.VirtualMeshSpec{},
			Status: zephyr_networking_types.VirtualMeshStatus{
				CertificateStatus: &zephyr_core_types.Status{
					State: zephyr_core_types.Status_INVALID,
				},
			},
		}
		snap := &networking_snapshot.MeshNetworkingSnapshot{
			VirtualMeshes: []*zephyr_networking.VirtualMesh{vm1, vm2},
		}
		csrProcessor.EXPECT().InitializeCertificateForVirtualMesh(ctx, vm1).Return(vm1.Status)
		csrProcessor.EXPECT().InitializeCertificateForVirtualMesh(ctx, vm2).Return(vm2.Status)

		virtualMeshClient.EXPECT().UpdateVirtualMeshStatus(ctx, vm1).Return(nil)
		virtualMeshClient.EXPECT().UpdateVirtualMeshStatus(ctx, vm2).Return(nil)
		csrSnapshotListener.Sync(ctx, snap)
	})
})
