package aws_test

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/aws_creds"
	mock_aws_creds "github.com/solo-io/service-mesh-hub/cli/pkg/common/aws_creds/mocks"
	aws2 "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/compute-target/aws"
	mock_rest_api "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/compute-target/aws/mocks"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("CredsHandler", func() {
	var (
		ctrl                *gomock.Controller
		ctx                 context.Context
		mockSecretConverter *mock_aws_creds.MockSecretAwsCredsConverter
		mockReconciler      *mock_rest_api.MockRestAPIDiscoveryReconciler
		awsCredsHandler     aws2.AwsCredsHandler
		secret              *k8s_core_types.Secret
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockSecretConverter = mock_aws_creds.NewMockSecretAwsCredsConverter(ctrl)
		mockReconciler = mock_rest_api.NewMockRestAPIDiscoveryReconciler(ctrl)
		awsCredsHandler = aws2.NewAwsAPIHandler(
			mockSecretConverter,
			[]aws2.RestAPIDiscoveryReconciler{mockReconciler},
		)
		secret = &k8s_core_types.Secret{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      "secret-name",
				Namespace: "service-mesh-hub",
			},
			Type: aws_creds.AWSSecretType,
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should ignore non-AWS secrets when compute target added", func() {
		err := awsCredsHandler.ComputeTargetAdded(ctx, &k8s_core_types.Secret{Type: k8s_core_types.SecretTypeOpaque})
		Expect(err).ToNot(HaveOccurred())
	})

	It("should ignore non-AWS secrets when compute target removed", func() {
		err := awsCredsHandler.ComputeTargetRemoved(ctx, &k8s_core_types.Secret{Type: k8s_core_types.SecretTypeOpaque})
		Expect(err).ToNot(HaveOccurred())
	})

	It("should handle new API registration", func() {
		creds := &credentials.Credentials{}
		mockSecretConverter.EXPECT().SecretToCreds(secret).Return(creds, nil)
		mockReconciler.EXPECT().Reconcile(gomock.Any(), creds, aws2.Region).Return(nil)
		err := awsCredsHandler.ComputeTargetAdded(ctx, secret)
		Expect(err).ToNot(HaveOccurred())
	})

	It("should handle new API deregistration", func() {
		// first register the API for cancelFunc map entry
		creds := &credentials.Credentials{}
		mockSecretConverter.EXPECT().SecretToCreds(secret).Return(creds, nil)
		mockReconciler.EXPECT().Reconcile(gomock.Any(), creds, aws2.Region).Return(nil)
		err := awsCredsHandler.ComputeTargetAdded(ctx, secret)
		Expect(err).ToNot(HaveOccurred())

		err = awsCredsHandler.ComputeTargetRemoved(ctx, secret)
		Expect(err).ToNot(HaveOccurred())
	})
})
