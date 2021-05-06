package reconciliation

import (
	"context"
	"time"

	skinput "github.com/solo-io/skv2/contrib/pkg/input"

	"github.com/rotisserie/eris"
	corev1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/issuer/input"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	v1sets "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/certificates/common/secrets"
	"github.com/solo-io/gloo-mesh/pkg/certificates/issuer/utils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// function which defines how the cert issuer reconciler should be registered with internal components.
type RegisterReconcilerFunc func(
	ctx context.Context,
	reconcile skinput.MultiClusterReconcileFunc,
	reconcileInterval time.Duration,
) error

// function which defines how the cert issuer should update the statuses of objects in its input snapshot
type SyncStatusFunc func(ctx context.Context, snapshot input.Snapshot) error

type certIssuerReconciler struct {
	ctx               context.Context
	builder           input.Builder
	syncInputStatuses SyncStatusFunc
	masterSecrets     corev1.SecretClient
}

func Start(
	ctx context.Context,
	registerReconciler RegisterReconcilerFunc,
	builder input.Builder,
	syncInputStatuses SyncStatusFunc,
	masterClient client.Client,
) error {
	r := &certIssuerReconciler{
		ctx:               ctx,
		builder:           builder,
		syncInputStatuses: syncInputStatuses,
		masterSecrets:     corev1.NewSecretClient(masterClient),
	}

	return registerReconciler(ctx, r.reconcile, time.Second/2)
}

// reconcile global state
func (r *certIssuerReconciler) reconcile(_ ezkube.ClusterResourceId) (bool, error) {
	inputSnap, err := r.builder.BuildSnapshot(r.ctx, "cert-issuer", input.BuildOptions{})
	if err != nil {
		// failed to read from cache; should never happen
		return false, err
	}

	for _, certificateRequest := range inputSnap.CertificateRequests().List() {
		if err := r.reconcileCertificateRequest(certificateRequest, inputSnap.IssuedCertificates()); err != nil {
			contextutils.LoggerFrom(r.ctx).Warnf("certificate request could not be processed: %v", err)
			certificateRequest.Status.Error = err.Error()
			certificateRequest.Status.State = v1.CertificateRequestStatus_FAILED
		}
	}

	return false, r.syncInputStatuses(r.ctx, inputSnap)
}

func (r *certIssuerReconciler) reconcileCertificateRequest(certificateRequest *v1.CertificateRequest, issuedCertificates v1sets.IssuedCertificateSet) error {
	// if observed generation is out of sync, treat the issued certificate as Pending (spec has been modified)
	if certificateRequest.Status.ObservedGeneration != certificateRequest.Generation {
		certificateRequest.Status.State = v1.CertificateRequestStatus_PENDING
	}

	// reset & update status
	certificateRequest.Status.ObservedGeneration = certificateRequest.Generation
	certificateRequest.Status.Error = ""

	switch certificateRequest.Status.State {
	case v1.CertificateRequestStatus_FINISHED:
		if len(certificateRequest.Status.SignedCertificate) > 0 {
			contextutils.LoggerFrom(r.ctx).Debugf("skipping cert request %v which has already been fulfilled", sets.Key(certificateRequest))
			return nil
		}
		// else treat as pending
		fallthrough
	case v1.CertificateRequestStatus_FAILED:
		// restart the workflow from PENDING
		fallthrough
	case v1.CertificateRequestStatus_PENDING:
		//
	default:
		return eris.Errorf("unknown certificate request state: %v", certificateRequest.Status.State)
	}

	issuedCertificate, err := issuedCertificates.Find(certificateRequest)
	if err != nil {
		return eris.Wrapf(err, "failed to find issued certificate matching certificate request")
	}

	signingCertificateSecret, err := r.masterSecrets.GetSecret(r.ctx, ezkube.MakeClientObjectKey(issuedCertificate.Spec.SigningCertificateSecret))
	if err != nil {
		return eris.Wrapf(err, "failed to find issuer's signing certificate matching issued request %v", sets.Key(issuedCertificate))
	}

	signingCA := secrets.RootCADataFromSecretData(signingCertificateSecret.Data)

	// generate the issued cert PEM encoded bytes
	signedCert, err := utils.GenCertForCSR(
		issuedCertificate.Spec.Hosts,
		certificateRequest.Spec.CertificateSigningRequest,
		signingCA.RootCert,
		signingCA.PrivateKey,
	)
	if err != nil {
		return eris.Wrapf(err, "failed to generate signed cert for certificate request %v", sets.Key(certificateRequest))
	}

	certificateRequest.Status = v1.CertificateRequestStatus{
		ObservedGeneration: certificateRequest.Generation,
		State:              v1.CertificateRequestStatus_FINISHED,
		SignedCertificate:  signedCert,
		SigningRootCa:      signingCA.RootCert,
	}

	return nil
}
