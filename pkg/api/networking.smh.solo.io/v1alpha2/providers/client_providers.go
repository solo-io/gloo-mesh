// Code generated by skv2. DO NOT EDIT.

package v1alpha2

import (
	networking_smh_solo_io_v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for TrafficPolicyClient from Clientset
func TrafficPolicyClientFromClientsetProvider(clients networking_smh_solo_io_v1alpha2.Clientset) networking_smh_solo_io_v1alpha2.TrafficPolicyClient {
	return clients.TrafficPolicies()
}

// Provider for TrafficPolicy Client from Client
func TrafficPolicyClientProvider(client client.Client) networking_smh_solo_io_v1alpha2.TrafficPolicyClient {
	return networking_smh_solo_io_v1alpha2.NewTrafficPolicyClient(client)
}

type TrafficPolicyClientFactory func(client client.Client) networking_smh_solo_io_v1alpha2.TrafficPolicyClient

func TrafficPolicyClientFactoryProvider() TrafficPolicyClientFactory {
	return TrafficPolicyClientProvider
}

type TrafficPolicyClientFromConfigFactory func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.TrafficPolicyClient, error)

func TrafficPolicyClientFromConfigFactoryProvider() TrafficPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.TrafficPolicyClient, error) {
		clients, err := networking_smh_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.TrafficPolicies(), nil
	}
}

// Provider for AccessPolicyClient from Clientset
func AccessPolicyClientFromClientsetProvider(clients networking_smh_solo_io_v1alpha2.Clientset) networking_smh_solo_io_v1alpha2.AccessPolicyClient {
	return clients.AccessPolicies()
}

// Provider for AccessPolicy Client from Client
func AccessPolicyClientProvider(client client.Client) networking_smh_solo_io_v1alpha2.AccessPolicyClient {
	return networking_smh_solo_io_v1alpha2.NewAccessPolicyClient(client)
}

type AccessPolicyClientFactory func(client client.Client) networking_smh_solo_io_v1alpha2.AccessPolicyClient

func AccessPolicyClientFactoryProvider() AccessPolicyClientFactory {
	return AccessPolicyClientProvider
}

type AccessPolicyClientFromConfigFactory func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.AccessPolicyClient, error)

func AccessPolicyClientFromConfigFactoryProvider() AccessPolicyClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.AccessPolicyClient, error) {
		clients, err := networking_smh_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.AccessPolicies(), nil
	}
}

// Provider for VirtualMeshClient from Clientset
func VirtualMeshClientFromClientsetProvider(clients networking_smh_solo_io_v1alpha2.Clientset) networking_smh_solo_io_v1alpha2.VirtualMeshClient {
	return clients.VirtualMeshes()
}

// Provider for VirtualMesh Client from Client
func VirtualMeshClientProvider(client client.Client) networking_smh_solo_io_v1alpha2.VirtualMeshClient {
	return networking_smh_solo_io_v1alpha2.NewVirtualMeshClient(client)
}

type VirtualMeshClientFactory func(client client.Client) networking_smh_solo_io_v1alpha2.VirtualMeshClient

func VirtualMeshClientFactoryProvider() VirtualMeshClientFactory {
	return VirtualMeshClientProvider
}

type VirtualMeshClientFromConfigFactory func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.VirtualMeshClient, error)

func VirtualMeshClientFromConfigFactoryProvider() VirtualMeshClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.VirtualMeshClient, error) {
		clients, err := networking_smh_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualMeshes(), nil
	}
}

// Provider for FailoverServiceClient from Clientset
func FailoverServiceClientFromClientsetProvider(clients networking_smh_solo_io_v1alpha2.Clientset) networking_smh_solo_io_v1alpha2.FailoverServiceClient {
	return clients.FailoverServices()
}

// Provider for FailoverService Client from Client
func FailoverServiceClientProvider(client client.Client) networking_smh_solo_io_v1alpha2.FailoverServiceClient {
	return networking_smh_solo_io_v1alpha2.NewFailoverServiceClient(client)
}

type FailoverServiceClientFactory func(client client.Client) networking_smh_solo_io_v1alpha2.FailoverServiceClient

func FailoverServiceClientFactoryProvider() FailoverServiceClientFactory {
	return FailoverServiceClientProvider
}

type FailoverServiceClientFromConfigFactory func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.FailoverServiceClient, error)

func FailoverServiceClientFromConfigFactoryProvider() FailoverServiceClientFromConfigFactory {
	return func(cfg *rest.Config) (networking_smh_solo_io_v1alpha2.FailoverServiceClient, error) {
		clients, err := networking_smh_solo_io_v1alpha2.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.FailoverServices(), nil
	}
}
