package destination_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	smiaccessv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	smispecsv1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	smisplitv1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	mock_output "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/smi/mocks"
	mock_reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting/mocks"
	. "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/smi/destination"
	mock_access "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/smi/destination/access/mocks"
	mock_split "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/smi/destination/split/mocks"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

var _ = Describe("SmiDestinationTranslator", func() {
	var (
		ctx                      context.Context
		ctrl                     *gomock.Controller
		mockOutputs              *mock_output.MockBuilder
		mockReporter             *mock_reporting.MockReporter
		mockSplitTranslator      *mock_split.MockTranslator
		mockAccessTranslator     *mock_access.MockTranslator
		smiDestinationTranslator Translator
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.Background(), GinkgoT())
		mockOutputs = mock_output.NewMockBuilder(ctrl)
		mockReporter = mock_reporting.NewMockReporter(ctrl)
		mockSplitTranslator = mock_split.NewMockTranslator(ctrl)
		mockAccessTranslator = mock_access.NewMockTranslator(ctrl)
		smiDestinationTranslator = NewTranslator(mockSplitTranslator, mockAccessTranslator)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should translate when an smi Destination", func() {
		destination := &v1.Destination{
			Spec: v1.DestinationSpec{
				Mesh: &skv2corev1.ObjectRef{
					Name:      "hello",
					Namespace: "world",
				},
			},
		}
		in := input.NewInputLocalSnapshotManualBuilder("").Build()

		ts := &smisplitv1alpha2.TrafficSplit{}

		mockSplitTranslator.
			EXPECT().
			Translate(gomock.AssignableToTypeOf(ctx), in, destination, mockReporter).
			Return(ts)

		mockOutputs.
			EXPECT().
			AddTrafficSplits(ts)

		tt := &smiaccessv1alpha2.TrafficTarget{}
		hrg := &smispecsv1alpha3.HTTPRouteGroup{}
		mockAccessTranslator.
			EXPECT().
			Translate(gomock.AssignableToTypeOf(ctx), in, destination, mockReporter).
			Return([]*smiaccessv1alpha2.TrafficTarget{tt}, []*smispecsv1alpha3.HTTPRouteGroup{hrg})

		mockOutputs.
			EXPECT().
			AddTrafficTargets(tt)

		mockOutputs.
			EXPECT().
			AddHTTPRouteGroups(hrg)

		smiDestinationTranslator.Translate(ctx, in, destination, mockOutputs, mockReporter)
	})
})
