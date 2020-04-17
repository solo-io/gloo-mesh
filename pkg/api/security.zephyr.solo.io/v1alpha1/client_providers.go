package v1alpha1

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provider for the security.zephyr.solo.io/v1alpha1 Clientset from config
func ClientsetFromConfigProvider(cfg *rest.Config) (Clientset, error) {
	return NewClientsetFromConfig(cfg)
}

// Provider for the security.zephyr.solo.io/v1alpha1 Clientset from client
func ClientsProvider(client client.Client) Clientset {
	return NewClientset(client)
}

// Provider for VirtualMeshCertificateSigningRequestClient from Clientset
func VirtualMeshCertificateSigningRequestClientFromClientsetProvider(clients Clientset) VirtualMeshCertificateSigningRequestClient {
	return clients.VirtualMeshCertificateSigningRequests()
}

// Provider for VirtualMeshCertificateSigningRequestClient from Client
func VirtualMeshCertificateSigningRequestClientProvider(client client.Client) VirtualMeshCertificateSigningRequestClient {
	return NewVirtualMeshCertificateSigningRequestClient(client)
}

type VirtualMeshCertificateSigningRequestClientFactory func(client client.Client) VirtualMeshCertificateSigningRequestClient

func VirtualMeshCertificateSigningRequestClientFactoryProvider() VirtualMeshCertificateSigningRequestClientFactory {
	return VirtualMeshCertificateSigningRequestClientProvider
}

type VirtualMeshCertificateSigningRequestClientFromConfigFactory func(cfg *rest.Config) (VirtualMeshCertificateSigningRequestClient, error)

func VirtualMeshCertificateSigningRequestClientFromConfigFactoryProvider() VirtualMeshCertificateSigningRequestClientFromConfigFactory {
	return func(cfg *rest.Config) (VirtualMeshCertificateSigningRequestClient, error) {
		clients, err := NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualMeshCertificateSigningRequests(), nil
	}
}
