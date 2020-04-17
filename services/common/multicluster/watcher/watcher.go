package mc_watcher

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	"github.com/solo-io/service-mesh-hub/services/common/multicluster"
	mc_manager "github.com/solo-io/service-mesh-hub/services/common/multicluster/manager"
	internal_watcher "github.com/solo-io/service-mesh-hub/services/common/multicluster/watcher/internal"
	core_v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

/*
	The following function is meant to serve as the secret watcher which will serve as the "entrypoint"
	for multi cluster applications. It will read in kube secrets with the `solo.io/kubeconfig` label
	and translate them into kube configs. The input to this function is the handler interface which
	will serve as the callbacks which receive the cluster config as they are added and deleted. This
	function is meant to be paired with the `AsyncManagerController` located in the `mcutils` package.

	The `Start()` function on the `AsyncManagerController` takes as an argument optional start functions,
	which this function returns. That way ensuring that this code is called anytime multi cluster watchers
	are necessary
*/
func StartLocalManager(handler mc_manager.KubeConfigHandler) mc_manager.AsyncManagerStartOptionsFunc {
	return func(ctx context.Context, mgr manager.Manager) error {
		secretCtrl := controller.NewSecretEventWatcher(multicluster.MultiClusterController, mgr)

		mcHandler := &multiClusterHandler{
			ctx:               ctx,
			clusterMembership: internal_watcher.NewClusterMembershipHandler(handler),
		}

		if err := secretCtrl.AddEventHandler(ctx, mcHandler, &internal_watcher.MultiClusterPredicate{}); err != nil {
			return err
		}
		return nil
	}
}

type multiClusterHandler struct {
	ctx               context.Context
	clusterMembership internal_watcher.ClusterSecretHandler
}

func NewMultiClusterHandler(ctx context.Context,
	clusterMembership internal_watcher.ClusterSecretHandler) controller.SecretEventHandler {
	return &multiClusterHandler{ctx: ctx, clusterMembership: clusterMembership}
}

func (c *multiClusterHandler) CreateSecret(s *core_v1.Secret) error {
	resync, err := c.clusterMembership.AddMemberCluster(c.ctx, s)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Errorf("error adding member cluster for secret %s.%s: %s", s.GetName(), s.GetNamespace(), err.Error())
		if resync {
			return err
		}
	}
	return nil
}

func (c *multiClusterHandler) UpdateSecret(old, new *core_v1.Secret) error {
	logger := contextutils.LoggerFrom(c.ctx)
	// If mc label has been removed from secret, remove from remote clusters
	if internal_watcher.HasRequiredMetadata(old.GetObjectMeta()) && !internal_watcher.HasRequiredMetadata(new.GetObjectMeta()) {
		resync, err := c.clusterMembership.DeleteMemberCluster(c.ctx, new)
		if err != nil {
			logger.Errorf("error deleting member cluster for "+
				"secret %s.%s", new.GetName(), new.GetNamespace())
			if resync {
				return err
			}
		}
	}
	// if mc label has been added to secret, add to remote cluster list
	if !internal_watcher.HasRequiredMetadata(old.GetObjectMeta()) && internal_watcher.HasRequiredMetadata(new.GetObjectMeta()) {
		resync, err := c.clusterMembership.AddMemberCluster(c.ctx, new)
		if err != nil {
			logger.Errorf("error adding member cluster for "+
				"secret %s.%s", new.GetName(), new.GetNamespace())
			if resync {
				return err
			}
		}
	}
	return nil
}

func (c *multiClusterHandler) DeleteSecret(s *core_v1.Secret) error {
	resync, err := c.clusterMembership.DeleteMemberCluster(c.ctx, s)
	if err != nil {
		contextutils.LoggerFrom(c.ctx).Errorf("error deleting member cluster for "+
			"secret %s.%s", s.GetName(), s.GetNamespace())
		if resync {
			return err
		}
	}
	return nil
}

func (c *multiClusterHandler) GenericSecret(obj *core_v1.Secret) error {
	contextutils.LoggerFrom(c.ctx).Warn("should never be called as generic events are not configured for secrets")
	return nil
}
