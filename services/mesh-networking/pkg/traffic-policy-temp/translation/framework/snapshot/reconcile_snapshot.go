package snapshot

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/multicluster"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/traffic-policy-temp/translation/framework/snapshot/reconcilers"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type snapshotReconciler struct {
	dynamicClientGetter multicluster.DynamicClientGetter

	virtualServiceReconciler  reconcilers.VirtualServiceReconcilerBuilder
	destinationRuleReconciler reconcilers.DestinationRuleReconcilerBuilder
}

func (r *snapshotReconciler) ReconcileAllSnapshots(ctx context.Context, clusterNameToSnapshot ClusterNameToSnapshot) error {
	var multierr error
	for cluster, snapshot := range clusterNameToSnapshot {
		err := r.reconcileCluster(ctx, cluster, snapshot)
		if err != nil {
			multierr = multierror.Append(multierr, err)
			continue
		}
	}
	return multierr
}

func (r *snapshotReconciler) reconcileCluster(ctx context.Context, cluster types.NamespacedName, snapshot *TranslatedSnapshot) error {
	client, err := r.dynamicClientGetter.GetClientForCluster(ctx, cluster.Name)
	if err != nil {
		return err
	}
	var multierr error
	if snapshot.Istio != nil {
		if err := r.reconcileIstio(ctx, client, snapshot.Istio); err != nil {
			multierr = multierror.Append(multierr, err)
		}
	}
	return multierr
}

func (r *snapshotReconciler) reconcileIstio(ctx context.Context, client client.Client, snapshot *IstioSnapshot) error {
	virtualServiceReconciler, err := r.virtualServiceReconciler.WithClient(client).ScopedToWholeCluster().Build()
	if err != nil {
		return err
	}
	err = virtualServiceReconciler.Reconcile(ctx, snapshot.VirtualServices)
	if err != nil {
		return err
	}
	destinationRuleReconciler, err := r.destinationRuleReconciler.WithClient(client).ScopedToWholeCluster().Build()
	if err != nil {
		return err
	}
	err = destinationRuleReconciler.Reconcile(ctx, snapshot.DestinationRules)
	if err != nil {
		return err
	}
	return nil
}
