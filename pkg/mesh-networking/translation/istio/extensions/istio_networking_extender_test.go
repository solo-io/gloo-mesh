package extensions_test

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/extensions/v1beta1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"

	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/extensions"
	mock_extensions "github.com/solo-io/gloo-mesh/pkg/mesh-networking/extensions/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	mock_istio_extensions "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/extensions/mocks"
	networkingv1alpha3spec "istio.io/api/networking/v1alpha3"
	istionetworkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/extensions"
)

//go:generate mockgen -destination mocks/mock_extensions_client.go -package mock_extensions github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/extensions/v1beta1 NetworkingExtensionsClient,NetworkingExtensions_WatchPushNotificationsClient

var _ = Describe("IstioNetworkingExtender", func() {
	var (
		ctl         *gomock.Controller
		client      *mock_istio_extensions.MockNetworkingExtensionsClient
		mockClients extensions.Clients
		clientset   *mock_extensions.MockClientset
		ctx         = context.TODO()
		exts        IstioExtender
	)
	BeforeEach(func() {
		ctl = gomock.NewController(GinkgoT())
		client = mock_istio_extensions.NewMockNetworkingExtensionsClient(ctl)
		clientset = mock_extensions.NewMockClientset(ctl)
		exts = NewIstioExtender(clientset)
		mockClients = extensions.Clients{client}
	})
	AfterEach(func() {
		ctl.Finish()
	})

	It("applies patches to istio outputs", func() {
		mesh := &discoveryv1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "mesh",
				Namespace: "namespace",
			},
		}

		workload := &discoveryv1.Workload{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "workload",
				Namespace: "namespace",
			},
		}

		destination := &discoveryv1.Destination{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "destination",
				Namespace: "namespace",
			},
		}

		inputs := input.NewInputLocalSnapshotManualBuilder("istio-extender-test").
			AddMeshes([]*discoveryv1.Mesh{mesh}).
			AddWorkloads([]*discoveryv1.Workload{workload}).
			AddDestinations([]*discoveryv1.Destination{destination}).
			Build()

		outputs := istio.NewBuilder(ctx, "test")
		outputs.AddVirtualServices(&istionetworkingv1alpha3.VirtualService{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "bar",
			},
		})
		expectedOutputs := outputs.Clone()
		// modify
		expectedOutputs.AddVirtualServices(&istionetworkingv1alpha3.VirtualService{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "bar",
			},
			Spec: networkingv1alpha3spec.VirtualService{
				Hosts: []string{"added-a-host"},
			},
		})
		// add
		expectedOutputs.AddDestinationRules(&istionetworkingv1alpha3.DestinationRule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "bar",
			},
		})

		clientset.EXPECT().GetClients().Return(mockClients)
		client.EXPECT().GetExtensionPatches(ctx, &v1beta1.ExtensionPatchRequest{
			Inputs:  extensions.InputSnapshotToProto(inputs),
			Outputs: OutputsToProto(outputs),
		}).Return(&v1beta1.ExtensionPatchResponse{
			PatchedOutputs: OutputsToProto(expectedOutputs),
		}, nil)

		// sanity check
		Expect(outputs).NotTo(Equal(expectedOutputs))

		err := exts.PatchOutputs(ctx, inputs, outputs)
		Expect(err).NotTo(HaveOccurred())

		// expect patches to be applied
		Expect(outputs).To(Equal(expectedOutputs))

	})
})
