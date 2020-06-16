// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	security_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for VirtualMeshCertificateSigningRequestClient from Clientset
func VirtualMeshCertificateSigningRequestClientFromClientsetProvider(clients security_smh_solo_io_v1alpha1.Clientset) security_smh_solo_io_v1alpha1.VirtualMeshCertificateSigningRequestClient {
	return clients.VirtualMeshCertificateSigningRequests()
}

// Provider for VirtualMeshCertificateSigningRequest Client from Client
func VirtualMeshCertificateSigningRequestClientProvider(client client.Client) security_smh_solo_io_v1alpha1.VirtualMeshCertificateSigningRequestClient {
	return security_smh_solo_io_v1alpha1.NewVirtualMeshCertificateSigningRequestClient(client)
}

type VirtualMeshCertificateSigningRequestClientFactory func(client client.Client) security_smh_solo_io_v1alpha1.VirtualMeshCertificateSigningRequestClient

func VirtualMeshCertificateSigningRequestClientFactoryProvider() VirtualMeshCertificateSigningRequestClientFactory {
	return VirtualMeshCertificateSigningRequestClientProvider
}

type VirtualMeshCertificateSigningRequestClientFromConfigFactory func(cfg *rest.Config) (security_smh_solo_io_v1alpha1.VirtualMeshCertificateSigningRequestClient, error)

func VirtualMeshCertificateSigningRequestClientFromConfigFactoryProvider() VirtualMeshCertificateSigningRequestClientFromConfigFactory {
	return func(cfg *rest.Config) (security_smh_solo_io_v1alpha1.VirtualMeshCertificateSigningRequestClient, error) {
		clients, err := security_smh_solo_io_v1alpha1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualMeshCertificateSigningRequests(), nil
	}
}
