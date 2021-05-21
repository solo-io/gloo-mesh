package reconciliation

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/agent/input"
	mock_certagent "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/agent/output/certagent/mocks"
	certificatesv1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("CertAgentReconciler", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		mockTranslator *mock_translation.MockTranslator
		mockPodBouncer *mock_podbouncer.MockPodBouncer
		mockOutput     *mock_certagent.MockBuilder
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.Background(), GinkgoT())

		mockTranslator = mock_translation.NewMockTranslator(ctrl)
		mockPodBouncer = mock_podbouncer.NewMockPodBouncer(ctrl)
		mockOutput = mock_certagent.NewMockBuilder(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("Will search for secret/readd it if Status==FINISHED", func() {

		// reconciler := reconciliation.NewCertAgentReconciler(ctx, mockPodBouncer, mockTranslator)
		reconciler := &certAgentReconciler{
			ctx:        ctx,
			podBouncer: mockPodBouncer,
			translator: mockTranslator,
		}
		issuedCert := &certificatesv1.IssuedCertificate{
			ObjectMeta: metav1.ObjectMeta{
				Generation: 2,
			},
			Spec: certificatesv1.IssuedCertificateSpec{
				IssuedCertificateSecret: &skv2corev1.ObjectRef{
					Name:      "hello",
					Namespace: "world",
				},
			},
			Status: certificatesv1.IssuedCertificateStatus{
				State:              certificatesv1.IssuedCertificateStatus_FINISHED,
				ObservedGeneration: 2,
			},
		}

		writtenSecret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      issuedCert.Spec.GetIssuedCertificateSecret().GetName(),
				Namespace: issuedCert.Spec.GetIssuedCertificateSecret().GetNamespace(),
			},
		}

		inputSnap := input.NewInputSnapshotManualBuilder("hello").
			AddSecrets([]*corev1.Secret{writtenSecret}).
			Build()

		mockTranslator.EXPECT().ShouldProcess(gomock.Any(), issuedCert).Return(true)
		mockOutput.EXPECT().AddSecrets(writtenSecret)

		err := reconciler.reconcileIssuedCertificate(issuedCert, inputSnap, mockOutput)
	Context("IssuedCertificatePending", func() {
			issuedCert *certificatesv1.IssuedCertificate
			csr        *certificatesv1.CertificateRequest
			csrBytes   []byte
					State:              certificatesv1.IssuedCertificateStatus_FINISHED,
				},
			}

			csrBytes = []byte("I'm a CSR")

			csr = &certificatesv1.CertificateRequest{
				ObjectMeta: metav1.ObjectMeta{
					Name:      issuedCert.GetName(),
					Namespace: issuedCert.GetNamespace(),
					Labels: map[string]string{
						"agent.certificates.mesh.gloo.solo.io": "gloo-mesh",
					},
				},
				Spec: certificatesv1.CertificateRequestSpec{
					CertificateSigningRequest: csrBytes,
				},
			}

		})

		It("Create CSR if translator Pending func returns properly", func() {
			reconciler := &certAgentReconciler{
				ctx:        ctx,
				podBouncer: mockPodBouncer,
				translator: mockTranslator,
			}

			inputSnap := input.NewInputSnapshotManualBuilder("hello").
				Build()

			mockTranslator.EXPECT().
				IssuedCertiticatePending(gomock.Any(), issuedCert, inputSnap, mockOutput).
				Return(csrBytes, nil)

			mockOutput.EXPECT().AddCertificateRequests(csr)

			mockTranslator.EXPECT().
				ShouldProcess(ctx, issuedCert).
				Return(true)

			reconciler := &certAgentReconciler{
				ctx:        ctx,
				podBouncer: mockPodBouncer,
				translator: mockTranslator,
			}

			inputSnap := input.NewInputSnapshotManualBuilder("hello").
				Build()

			mockTranslator.EXPECT().
				ShouldProcess(ctx, issuedCert).
				Return(false)

			err := reconciler.reconcileIssuedCertificate(issuedCert, inputSnap, mockOutput)
			Expect(err).NotTo(HaveOccurred())
			Expect(issuedCert.Status.State).To(Equal(certificatesv1.IssuedCertificateStatus_FINISHED))
		})

	})

	Context("IssuedCertificateRequested", func() {
		var (
			issuedCert *certificatesv1.IssuedCertificate
			csr        *certificatesv1.CertificateRequest
			csrBytes   []byte
		)
		BeforeEach(func() {
			issuedCert = &certificatesv1.IssuedCertificate{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "hello",
					Namespace:  "world",
					Generation: 2,
				},
				Spec: certificatesv1.IssuedCertificateSpec{},
				Status: certificatesv1.IssuedCertificateStatus{
					State:              certificatesv1.IssuedCertificateStatus_REQUESTED,
					ObservedGeneration: 2,
				},
			}

			csrBytes = []byte("I'm a CSR")

			csr = &certificatesv1.CertificateRequest{
				ObjectMeta: metav1.ObjectMeta{
					Name:      issuedCert.GetName(),
					Namespace: issuedCert.GetNamespace(),
					Labels: map[string]string{
						"agent.certificates.mesh.gloo.solo.io": "gloo-mesh",
					},
				},
				Spec: certificatesv1.CertificateRequestSpec{
					CertificateSigningRequest: csrBytes,
				},
			}
		})

		It("Find CSR and pass into translator when cert is requested", func() {

			reconciler := &certAgentReconciler{
				ctx:        ctx,
				podBouncer: mockPodBouncer,
				translator: mockTranslator,
			}

			inputSnap := input.NewInputSnapshotManualBuilder("hello").
				AddCertificateRequests([]*certificatesv1.CertificateRequest{csr}).
				Build()

			mockTranslator.EXPECT().
				ShouldProcess(gomock.Any(), issuedCert).
				Return(true)

			mockTranslator.EXPECT().
				IssuedCertificateRequested(gomock.Any(), issuedCert, csr, inputSnap, mockOutput).
				Return(nil)

			err := reconciler.reconcileIssuedCertificate(issuedCert, inputSnap, mockOutput)
			Expect(err).NotTo(HaveOccurred())
			Expect(issuedCert.Status.State).To(Equal(certificatesv1.IssuedCertificateStatus_ISSUED))
		})

		It("Will not update status when translator.ShouldProcess == false", func() {

			reconciler := &certAgentReconciler{
				ctx:        ctx,
				podBouncer: mockPodBouncer,
				translator: mockTranslator,
			}

			inputSnap := input.NewInputSnapshotManualBuilder("hello").
				AddCertificateRequests([]*certificatesv1.CertificateRequest{csr}).
				Build()

			mockTranslator.EXPECT().
				ShouldProcess(gomock.Any(), issuedCert).
				Return(false)

			err := reconciler.reconcileIssuedCertificate(issuedCert, inputSnap, mockOutput)
			Expect(err).NotTo(HaveOccurred())
			Expect(issuedCert.Status.State).To(Equal(certificatesv1.IssuedCertificateStatus_REQUESTED))
		})
	})

	Context("IssuedCertificateIssued", func() {
		var (
			issuedCert *certificatesv1.IssuedCertificate
			pbd        *certificatesv1.PodBounceDirective
		)

		BeforeEach(func() {
			pbd = &certificatesv1.PodBounceDirective{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "hello",
					Namespace: "world",
				},
				Spec: certificatesv1.PodBounceDirectiveSpec{},
			}

			issuedCert = &certificatesv1.IssuedCertificate{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "hello",
					Namespace:  "world",
					Generation: 2,
				},
				Spec: certificatesv1.IssuedCertificateSpec{
					PodBounceDirective: ezkube.MakeObjectRef(pbd),
				},
				Status: certificatesv1.IssuedCertificateStatus{
					State:              certificatesv1.IssuedCertificateStatus_ISSUED,
					ObservedGeneration: 2,
				},
			}
		})

		It("Will delete pods when cert has been issued", func() {

			reconciler := &certAgentReconciler{
				ctx:        ctx,
				podBouncer: mockPodBouncer,
				translator: mockTranslator,
			}

			pods := v1sets.NewPodSet(&corev1.Pod{})
			configMaps := v1sets.NewConfigMapSet(&corev1.ConfigMap{})
			secrets := v1sets.NewSecretSet(&corev1.Secret{})

			inputSnap := input.NewInputSnapshotManualBuilder("hello").
				AddPodBounceDirectives([]*certificatesv1.PodBounceDirective{pbd}).
				AddPods(pods.List()).
				AddSecrets(secrets.List()).
				AddConfigMaps(configMaps.List()).
				Build()

			mockTranslator.EXPECT().
				ShouldProcess(gomock.Any(), issuedCert).
				Return(true)

			mockTranslator.EXPECT().
				IssuedCertificateIssued(gomock.Any(), issuedCert, inputSnap, mockOutput).
				Return(nil)

			mockPodBouncer.EXPECT().
				BouncePods(gomock.Any(), pbd, pods, configMaps, secrets).
				Return(false, nil)

			err := reconciler.reconcileIssuedCertificate(issuedCert, inputSnap, mockOutput)
			Expect(err).NotTo(HaveOccurred())
			Expect(issuedCert.Status.State).To(Equal(certificatesv1.IssuedCertificateStatus_FINISHED))
		})

		It("Will not delete pods when translator.ShouldProcess==false", func() {

			reconciler := &certAgentReconciler{
				ctx:        ctx,
				podBouncer: mockPodBouncer,
				translator: mockTranslator,
			}

			pods := v1sets.NewPodSet(&corev1.Pod{})
			configMaps := v1sets.NewConfigMapSet(&corev1.ConfigMap{})
			secrets := v1sets.NewSecretSet(&corev1.Secret{})

			inputSnap := input.NewInputSnapshotManualBuilder("hello").
				AddPodBounceDirectives([]*certificatesv1.PodBounceDirective{pbd}).
				AddPods(pods.List()).
				AddSecrets(secrets.List()).
				AddConfigMaps(configMaps.List()).
				Build()

			mockTranslator.EXPECT().
				ShouldProcess(gomock.Any(), issuedCert).
				Return(false)

			err := reconciler.reconcileIssuedCertificate(issuedCert, inputSnap, mockOutput)
			Expect(err).NotTo(HaveOccurred())
			Expect(issuedCert.Status.State).To(Equal(certificatesv1.IssuedCertificateStatus_ISSUED))
		})
	})
})
