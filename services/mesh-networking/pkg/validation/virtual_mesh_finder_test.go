package vm_validation_test

import (
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	. "github.com/solo-io/go-utils/testutils"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_networking_types "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	vm_validation "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/validation"
	mock_zephyr_discovery "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("mesh ref finder", func() {
	var (
		ctrl          *gomock.Controller
		ctx           context.Context
		meshClient    *mock_zephyr_discovery.MockMeshClient
		meshRefFinder vm_validation.VirtualMeshFinder

		testErr = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		meshClient = mock_zephyr_discovery.NewMockMeshClient(ctrl)
		meshRefFinder = vm_validation.NewVirtualMeshFinder(meshClient)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will fail if initial mesh list fails", func() {
		meshClient.EXPECT().
			ListMesh(ctx).
			Return(nil, testErr)

		_, err := meshRefFinder.GetMeshesForVirtualMesh(ctx, &zephyr_networking.VirtualMesh{})
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will reutrn nil, nil with no refs as input", func() {
		meshClient.EXPECT().
			ListMesh(ctx).
			Return(nil, nil)
		list, err := meshRefFinder.GetMeshesForVirtualMesh(ctx, &zephyr_networking.VirtualMesh{})
		Expect(err).NotTo(HaveOccurred())
		Expect(list).To(BeNil())
	})

	It("will return an error containing all invalid refs", func() {
		meshList := &zephyr_discovery.MeshList{}
		refs := []*zephyr_core_types.ResourceRef{
			{
				Name:      "name1",
				Namespace: "namespace1",
			},
			{
				Name:      "name2",
				Namespace: "namespace2",
			},
		}
		vm := &zephyr_networking.VirtualMesh{
			Spec: zephyr_networking_types.VirtualMeshSpec{
				Meshes: refs,
			},
		}
		meshClient.EXPECT().
			ListMesh(ctx).
			Return(meshList, nil)
		_, err := meshRefFinder.GetMeshesForVirtualMesh(ctx, vm)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(vm_validation.InvalidMeshRefsError([]string{
			fmt.Sprintf("%s.%s", refs[0].GetName(), refs[0].GetNamespace()),
			fmt.Sprintf("%s.%s", refs[1].GetName(), refs[1].GetNamespace()),
		})))
	})

	It("will return an error containing all invalid refs", func() {
		meshList := &zephyr_discovery.MeshList{
			Items: []zephyr_discovery.Mesh{
				{
					ObjectMeta: k8s_meta_types.ObjectMeta{
						Name:      "name1",
						Namespace: "namespace1",
					},
				},
				{
					ObjectMeta: k8s_meta_types.ObjectMeta{
						Name:      "name2",
						Namespace: "namespace2",
					},
				},
			},
		}
		refs := []*zephyr_core_types.ResourceRef{
			{
				Name:      "name1",
				Namespace: "namespace1",
			},
			{
				Name:      "name2",
				Namespace: "namespace2",
			},
		}
		vm := &zephyr_networking.VirtualMesh{
			Spec: zephyr_networking_types.VirtualMeshSpec{
				Meshes: refs,
			},
		}
		meshClient.EXPECT().
			ListMesh(ctx).
			Return(meshList, nil)
		list, err := meshRefFinder.GetMeshesForVirtualMesh(ctx, vm)
		Expect(err).NotTo(HaveOccurred())
		Expect(list).To(ConsistOf(&meshList.Items[0], &meshList.Items[1]))
	})
})
