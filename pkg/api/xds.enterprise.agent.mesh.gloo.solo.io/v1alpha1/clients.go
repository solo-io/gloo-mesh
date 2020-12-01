// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha1

import (
	"context"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the xds.enterprise.agent.mesh.gloo.solo.io/v1alpha1 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the xds.enterprise.agent.mesh.gloo.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the xds.enterprise.agent.mesh.gloo.solo.io/v1alpha1/v1alpha1 APIs
	XdsConfigs() XdsConfigClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the xds.enterprise.agent.mesh.gloo.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) XdsConfigs() XdsConfigClient {
	return NewXdsConfigClient(c.client)
}

// Reader knows how to read and list XdsConfigs.
type XdsConfigReader interface {
	// Get retrieves a XdsConfig for the given object key
	GetXdsConfig(ctx context.Context, key client.ObjectKey) (*XdsConfig, error)

	// List retrieves list of XdsConfigs for a given namespace and list options.
	ListXdsConfig(ctx context.Context, opts ...client.ListOption) (*XdsConfigList, error)
}

// XdsConfigTransitionFunction instructs the XdsConfigWriter how to transition between an existing
// XdsConfig object and a desired on an Upsert
type XdsConfigTransitionFunction func(existing, desired *XdsConfig) error

// Writer knows how to create, delete, and update XdsConfigs.
type XdsConfigWriter interface {
	// Create saves the XdsConfig object.
	CreateXdsConfig(ctx context.Context, obj *XdsConfig, opts ...client.CreateOption) error

	// Delete deletes the XdsConfig object.
	DeleteXdsConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given XdsConfig object.
	UpdateXdsConfig(ctx context.Context, obj *XdsConfig, opts ...client.UpdateOption) error

	// Patch patches the given XdsConfig object.
	PatchXdsConfig(ctx context.Context, obj *XdsConfig, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all XdsConfig objects matching the given options.
	DeleteAllOfXdsConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the XdsConfig object.
	UpsertXdsConfig(ctx context.Context, obj *XdsConfig, transitionFuncs ...XdsConfigTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a XdsConfig object.
type XdsConfigStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given XdsConfig object.
	UpdateXdsConfigStatus(ctx context.Context, obj *XdsConfig, opts ...client.UpdateOption) error

	// Patch patches the given XdsConfig object's subresource.
	PatchXdsConfigStatus(ctx context.Context, obj *XdsConfig, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on XdsConfigs.
type XdsConfigClient interface {
	XdsConfigReader
	XdsConfigWriter
	XdsConfigStatusWriter
}

type xdsConfigClient struct {
	client client.Client
}

func NewXdsConfigClient(client client.Client) *xdsConfigClient {
	return &xdsConfigClient{client: client}
}

func (c *xdsConfigClient) GetXdsConfig(ctx context.Context, key client.ObjectKey) (*XdsConfig, error) {
	obj := &XdsConfig{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *xdsConfigClient) ListXdsConfig(ctx context.Context, opts ...client.ListOption) (*XdsConfigList, error) {
	list := &XdsConfigList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *xdsConfigClient) CreateXdsConfig(ctx context.Context, obj *XdsConfig, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *xdsConfigClient) DeleteXdsConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &XdsConfig{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *xdsConfigClient) UpdateXdsConfig(ctx context.Context, obj *XdsConfig, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *xdsConfigClient) PatchXdsConfig(ctx context.Context, obj *XdsConfig, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *xdsConfigClient) DeleteAllOfXdsConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &XdsConfig{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *xdsConfigClient) UpsertXdsConfig(ctx context.Context, obj *XdsConfig, transitionFuncs ...XdsConfigTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*XdsConfig), desired.(*XdsConfig)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *xdsConfigClient) UpdateXdsConfigStatus(ctx context.Context, obj *XdsConfig, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *xdsConfigClient) PatchXdsConfigStatus(ctx context.Context, obj *XdsConfig, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides XdsConfigClients for multiple clusters.
type MulticlusterXdsConfigClient interface {
	// Cluster returns a XdsConfigClient for the given cluster
	Cluster(cluster string) (XdsConfigClient, error)
}

type multiclusterXdsConfigClient struct {
	client multicluster.Client
}

func NewMulticlusterXdsConfigClient(client multicluster.Client) MulticlusterXdsConfigClient {
	return &multiclusterXdsConfigClient{client: client}
}

func (m *multiclusterXdsConfigClient) Cluster(cluster string) (XdsConfigClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewXdsConfigClient(client), nil
}
