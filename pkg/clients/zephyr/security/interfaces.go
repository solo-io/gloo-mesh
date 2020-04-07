package zephyr_security

import (
	"context"

	"github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -destination ./mocks/mock_interfaces.go -source ./interfaces.go

type VirtualMeshCSRClient interface {
	Create(ctx context.Context, csr *v1alpha1.VirtualMeshCertificateSigningRequest, opts ...client.CreateOption) error
	Update(ctx context.Context, csr *v1alpha1.VirtualMeshCertificateSigningRequest, opts ...client.UpdateOption) error
	UpdateStatus(ctx context.Context, csr *v1alpha1.VirtualMeshCertificateSigningRequest, opts ...client.UpdateOption) error
	Get(ctx context.Context, name, namespace string) (*v1alpha1.VirtualMeshCertificateSigningRequest, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.VirtualMeshCertificateSigningRequestList, error)
	Delete(ctx context.Context, csr *v1alpha1.VirtualMeshCertificateSigningRequest, opts ...client.DeleteOption) error
}
