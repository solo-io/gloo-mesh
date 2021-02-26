package osm

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	mock_output "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/smi/mocks"
	mock_reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting/mocks"
	mock_traffictarget "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/osm/destination/mocks"
	. "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/osm/internal/mocks"
	mock_mesh "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/osm/mesh/mocks"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("SmiNetworkingTranslator", func() {
	var (
		ctrl                      *gomock.Controller
		ctx                       context.Context
		mockReporter              *mock_reporting.MockReporter
		mockOutputs               *mock_output.MockBuilder
		mockDependencyFactory     *MockDependencyFactory
		mockMeshTranslator        *mock_mesh.MockTranslator
		mockDestinationTranslator *mock_traffictarget.MockTranslator
		translator                *osmTranslator
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockReporter = mock_reporting.NewMockReporter(ctrl)
		mockDependencyFactory = NewMockDependencyFactory(ctrl)
		mockOutputs = mock_output.NewMockBuilder(ctrl)
		mockMeshTranslator = mock_mesh.NewMockTranslator(ctrl)
		mockDestinationTranslator = mock_traffictarget.NewMockTranslator(ctrl)
		translator = &osmTranslator{dependencies: mockDependencyFactory}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should translate all meshes and traffictargets", func() {
		in := input.NewInputLocalSnapshotManualBuilder("").
			AddMeshes([]*discoveryv1.Mesh{
				{
					ObjectMeta: metav1.ObjectMeta{},
					Spec:       discoveryv1.MeshSpec{},
					Status:     discoveryv1.MeshStatus{},
				},
			}).
			AddDestinations([]*discoveryv1.Destination{
				{
					ObjectMeta: metav1.ObjectMeta{},
					Spec:       discoveryv1.DestinationSpec{},
					Status:     discoveryv1.DestinationStatus{},
				},
			}).
			Build()

		mockDependencyFactory.
			EXPECT().
			MakeMeshTranslator().
			Return(mockMeshTranslator)

		mockDependencyFactory.
			EXPECT().
			MakeDestinationTranslator().
			Return(mockDestinationTranslator)

		for i := range in.Meshes().List() {
			mockMeshTranslator.
				EXPECT().
				Translate(gomock.Any(), in, in.Meshes().List()[i], mockOutputs, mockReporter)
		}

		for i := range in.Destinations().List() {
			mockDestinationTranslator.
				EXPECT().
				Translate(gomock.Any(), in, in.Destinations().List()[i], mockOutputs, mockReporter)
		}

		translator.Translate(ctx, in, mockOutputs, mockReporter)
	})
})
