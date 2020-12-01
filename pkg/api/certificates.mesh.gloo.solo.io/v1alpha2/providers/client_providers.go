// Code generated by skv2. DO NOT EDIT.

package v1alpha2

import (
	certificates_mesh_gloo_solo_io_v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for IssuedCertificateClient from Clientset
func IssuedCertificateClientFromClientsetProvider(clients certificates_mesh_gloo_solo_io_v1alpha2.Clientset) certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificateClient {
	return clients.IssuedCertificates()
}

// Provider for IssuedCertificate Client from Client
func IssuedCertificateClientProvider(client client.Client) certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificateClient {
	return certificates_mesh_gloo_solo_io_v1alpha2.NewIssuedCertificateClient(client)
}

type IssuedCertificateClientFactory func(client client.Client) certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificateClient

func IssuedCertificateClientFactoryProvider() IssuedCertificateClientFactory {
	return IssuedCertificateClientProvider
}

type IssuedCertificateClientFromConfigFactory func(cfg *rest.Config) (certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificateClient, error)

func IssuedCertificateClientFromConfigFactoryProvider() IssuedCertificateClientFromConfigFactory {
	return func(cfg *rest.Config) (certificates_mesh_gloo_solo_io_v1alpha2.IssuedCertificateClient, error) {
		clients, err := certificates_mesh_gloo_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.IssuedCertificates(), nil
	}
}

// Provider for CertificateRequestClient from Clientset
func CertificateRequestClientFromClientsetProvider(clients certificates_mesh_gloo_solo_io_v1alpha2.Clientset) certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequestClient {
	return clients.CertificateRequests()
}

// Provider for CertificateRequest Client from Client
func CertificateRequestClientProvider(client client.Client) certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequestClient {
	return certificates_mesh_gloo_solo_io_v1alpha2.NewCertificateRequestClient(client)
}

type CertificateRequestClientFactory func(client client.Client) certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequestClient

func CertificateRequestClientFactoryProvider() CertificateRequestClientFactory {
	return CertificateRequestClientProvider
}

type CertificateRequestClientFromConfigFactory func(cfg *rest.Config) (certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequestClient, error)

func CertificateRequestClientFromConfigFactoryProvider() CertificateRequestClientFromConfigFactory {
	return func(cfg *rest.Config) (certificates_mesh_gloo_solo_io_v1alpha2.CertificateRequestClient, error) {
		clients, err := certificates_mesh_gloo_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.CertificateRequests(), nil
	}
}

// Provider for PodBounceDirectiveClient from Clientset
func PodBounceDirectiveClientFromClientsetProvider(clients certificates_mesh_gloo_solo_io_v1alpha2.Clientset) certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirectiveClient {
	return clients.PodBounceDirectives()
}

// Provider for PodBounceDirective Client from Client
func PodBounceDirectiveClientProvider(client client.Client) certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirectiveClient {
	return certificates_mesh_gloo_solo_io_v1alpha2.NewPodBounceDirectiveClient(client)
}

type PodBounceDirectiveClientFactory func(client client.Client) certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirectiveClient

func PodBounceDirectiveClientFactoryProvider() PodBounceDirectiveClientFactory {
	return PodBounceDirectiveClientProvider
}

type PodBounceDirectiveClientFromConfigFactory func(cfg *rest.Config) (certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirectiveClient, error)

func PodBounceDirectiveClientFromConfigFactoryProvider() PodBounceDirectiveClientFromConfigFactory {
	return func(cfg *rest.Config) (certificates_mesh_gloo_solo_io_v1alpha2.PodBounceDirectiveClient, error) {
		clients, err := certificates_mesh_gloo_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.PodBounceDirectives(), nil
	}
}
