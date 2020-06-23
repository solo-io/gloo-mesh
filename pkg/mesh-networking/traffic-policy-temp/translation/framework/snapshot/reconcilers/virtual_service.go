package reconcilers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	istio_networking_clients "github.com/solo-io/external-apis/pkg/api/istio/networking.istio.io/v1alpha3"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/selection"
	istio_networking "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VirtualServiceReconciler interface {
	Reconcile(ctx context.Context, desiredGlobalState []*istio_networking.VirtualService) error
}

// a reconciler can either be whole-cluster scoped or scoped to a namespace. In addition labels can be used.
// either ScopedToWholeCluster() must be set or one of the following two methods must be called with non-zero values
// that's just to force it to be obvious when we're going to be reconciling EVERYTHING across a whole cluster
type VirtualServiceReconcilerBuilder interface {
	ScopedToWholeCluster() VirtualServiceReconcilerBuilder
	ScopedToNamespace(namespace string) VirtualServiceReconcilerBuilder
	ScopedToLabels(labels map[string]string) VirtualServiceReconcilerBuilder
	WithClient(client client.Client) VirtualServiceReconcilerBuilder

	Build() (VirtualServiceReconciler, error)
}

type virtualServiceReconcilerBuilder struct {
	clusterScoped bool
	namespace     string
	labels        map[string]string
	client        istio_networking_clients.VirtualServiceClient
}

func NewVirtualServiceReconcilerBuilder() VirtualServiceReconcilerBuilder {
	return &virtualServiceReconcilerBuilder{}
}

func (v *virtualServiceReconcilerBuilder) ScopedToWholeCluster() VirtualServiceReconcilerBuilder {
	v.clusterScoped = true
	return v
}

func (v *virtualServiceReconcilerBuilder) ScopedToNamespace(namespace string) VirtualServiceReconcilerBuilder {
	v.namespace = namespace
	return v
}

func (v *virtualServiceReconcilerBuilder) ScopedToLabels(labels map[string]string) VirtualServiceReconcilerBuilder {
	v.labels = labels
	return v
}

func (v *virtualServiceReconcilerBuilder) WithClient(client client.Client) VirtualServiceReconcilerBuilder {
	v.client = istio_networking_clients.NewVirtualServiceClient(client)
	return v
}

func (v *virtualServiceReconcilerBuilder) Build() (VirtualServiceReconciler, error) {
	if v.client == nil {
		return nil, eris.New("Must provide a client")
	}

	if !v.clusterScoped && v.namespace == "" && len(v.labels) == 0 {
		return nil, eris.New("Must either configure this reconciler to be cluster-scoped or explicitly scope it to a namespace/label")
	}

	if v.clusterScoped && v.namespace != "" {
		return nil, eris.New("Cannot be cluster-scoped and scoped to a namespace")
	}

	return &virtualServiceReconciler{
		namespace:            v.namespace,
		labels:               v.labels,
		virtualServiceClient: v.client,
	}, nil
}

func NewVirtualServiceReconciler(namespace string, labels map[string]string, client istio_networking_clients.VirtualServiceClient) VirtualServiceReconciler {
	return &virtualServiceReconciler{
		namespace:            namespace,
		labels:               labels,
		virtualServiceClient: client,
	}
}

type virtualServiceReconciler struct {
	virtualServiceClient istio_networking_clients.VirtualServiceClient

	namespace string
	labels    map[string]string
}

func (v *virtualServiceReconciler) Reconcile(ctx context.Context, desiredGlobalState []*istio_networking.VirtualService) error {
	virtualServiceList, err := v.virtualServiceClient.ListVirtualService(
		ctx,

		// if this reconciler has been scoped to the whole cluster, these two values will be their respective zero-values and this will list all the objects on the cluster
		client.InNamespace(v.namespace),
		client.MatchingLabels(v.labels),
	)
	if err != nil {
		return err
	}

	nameNamespaceToDesiredState := map[string]*istio_networking.VirtualService{}

	for _, desiredVsIter := range desiredGlobalState {
		desiredVs := desiredVsIter
		nameNamespaceToDesiredState[selection.ToUniqueSingleClusterString(desiredVs.ObjectMeta)] = desiredVs
	}

	var multierr error
	// update and delete existing VS's
	for _, existingVirtualService := range virtualServiceList.Items {
		key := selection.ToUniqueSingleClusterString(existingVirtualService.ObjectMeta)
		desiredState, shouldBeAlive := nameNamespaceToDesiredState[key]
		delete(nameNamespaceToDesiredState, key)

		if !shouldBeAlive {
			err = v.virtualServiceClient.DeleteVirtualService(ctx, selection.ObjectMetaToObjectKey(existingVirtualService.ObjectMeta))
			if err != nil {
			}
		} else if !proto.Equal(&existingVirtualService.Spec, &desiredState.Spec) {
			// make sure we use the same resource version for updates
			desiredState.ObjectMeta.ResourceVersion = existingVirtualService.ObjectMeta.ResourceVersion
			v.addLabels(desiredState)
			err = v.virtualServiceClient.UpdateVirtualService(ctx, desiredState)
			if err != nil {
				multierr = multierror.Append(multierr, err)
			}
		}
	}

	// create new VS's of what's left in the map
	for _, desiredVirtualService := range nameNamespaceToDesiredState { // add our labels:
		v.addLabels(desiredVirtualService)
		err := v.virtualServiceClient.CreateVirtualService(ctx, desiredVirtualService)
		if err != nil {
			multierr = multierror.Append(multierr, err)
		}
	}

	return multierr
}

func (v *virtualServiceReconciler) addLabels(desiredVirtualService *istio_networking.VirtualService) {
	if desiredVirtualService.Labels == nil && len(v.labels) != 0 {
		desiredVirtualService.Labels = make(map[string]string)
	}
	for k, v := range v.labels {
		desiredVirtualService.Labels[k] = v
	}
}
