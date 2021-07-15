// Code generated by skv2. DO NOT EDIT.

package v1

import (
	istio_enterprise_mesh_gloo_solo_io_v1 "github.com/solo-io/gloo-mesh/pkg/api/istio.enterprise.mesh.gloo.solo.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for IstioInstallationClient from Clientset
func IstioInstallationClientFromClientsetProvider(clients istio_enterprise_mesh_gloo_solo_io_v1.Clientset) istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallationClient {
	return clients.IstioInstallations()
}

// Provider for IstioInstallation Client from Client
func IstioInstallationClientProvider(client client.Client) istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallationClient {
	return istio_enterprise_mesh_gloo_solo_io_v1.NewIstioInstallationClient(client)
}

type IstioInstallationClientFactory func(client client.Client) istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallationClient

func IstioInstallationClientFactoryProvider() IstioInstallationClientFactory {
	return IstioInstallationClientProvider
}

type IstioInstallationClientFromConfigFactory func(cfg *rest.Config) (istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallationClient, error)

func IstioInstallationClientFromConfigFactoryProvider() IstioInstallationClientFromConfigFactory {
	return func(cfg *rest.Config) (istio_enterprise_mesh_gloo_solo_io_v1.IstioInstallationClient, error) {
		clients, err := istio_enterprise_mesh_gloo_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.IstioInstallations(), nil
	}
}