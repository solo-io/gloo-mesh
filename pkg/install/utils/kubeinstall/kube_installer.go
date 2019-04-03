package kubeinstall

import (
	"context"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/supergloo/pkg/install/utils/kuberesource"
	"golang.org/x/sync/errgroup"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiexts "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// an interface allowing these methods to be mocked
type Installer interface {
	ReconcilleResources(ctx context.Context, installNamespace string, resources kuberesource.UnstructuredResources, installLabels map[string]string) error
	PurgeResources(ctx context.Context, withLabels map[string]string) error
}

type KubeInstaller struct {
	cache         *Cache
	cfg           *rest.Config
	dynamic       dynamic.Interface
	client        client.Client
	core          kubernetes.Interface
	apiExtensions apiexts.Interface
	callbacks     []CallbackOptions
}

var _ Installer = &KubeInstaller{}

/*
NewKubeInstaller does not initialize the cache.
Should be one once globally
*/
func NewKubeInstaller(cfg *rest.Config, cache *Cache, callbacks ...CallbackOptions) (*KubeInstaller, error) {
	apiExts, err := apiexts.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	core, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{})
	if err != nil {
		return nil, err
	}

	return &KubeInstaller{
		cache:         cache,
		cfg:           cfg,
		apiExtensions: apiExts,
		client:        client,
		dynamic:       dynamicClient,
		core:          core,
		callbacks:     callbacks,
	}, nil
}

func (r *KubeInstaller) preInstall() error {
	for _, cb := range r.callbacks {
		if err := cb.PreInstall(); err != nil {
			return errors.Wrapf(err, "error in pre-install hook")
		}
	}
	return nil
}

func (r *KubeInstaller) postInstall() error {
	for _, cb := range r.callbacks {
		if err := cb.PostInstall(); err != nil {
			return errors.Wrapf(err, "error in post-install hook")
		}
	}
	return nil
}

func (r *KubeInstaller) preCreate(res *unstructured.Unstructured) error {
	for _, cb := range r.callbacks {
		if err := cb.PreCreate(res); err != nil {
			return errors.Wrapf(err, "error in pre-create hook")
		}
	}
	return nil
}

func (r *KubeInstaller) postCreate(res *unstructured.Unstructured) error {
	for _, cb := range r.callbacks {
		if err := cb.PostCreate(res); err != nil {
			return errors.Wrapf(err, "error in post-create hook")
		}
	}
	return nil
}

func (r *KubeInstaller) preUpdate(res *unstructured.Unstructured) error {
	for _, cb := range r.callbacks {
		if err := cb.PreUpdate(res); err != nil {
			return errors.Wrapf(err, "error in pre-update hook")
		}
	}
	return nil
}

func (r *KubeInstaller) postUpdate(res *unstructured.Unstructured) error {
	for _, cb := range r.callbacks {
		if err := cb.PostUpdate(res); err != nil {
			return errors.Wrapf(err, "error in post-update hook")
		}
	}
	return nil
}

func (r *KubeInstaller) preDelete(res *unstructured.Unstructured) error {
	for _, cb := range r.callbacks {
		if err := cb.PreDelete(res); err != nil {
			return errors.Wrapf(err, "error in pre-delete hook")
		}
	}
	return nil
}

func (r *KubeInstaller) postDelete(res *unstructured.Unstructured) error {
	for _, cb := range r.callbacks {
		if err := cb.PostDelete(res); err != nil {
			return errors.Wrapf(err, "error in post-delete hook")
		}
	}
	return nil
}

type CallbackOptions interface {
	PreInstall() error
	PostInstall() error
	PreCreate(res *unstructured.Unstructured) error
	PostCreate(res *unstructured.Unstructured) error
	PreUpdate(res *unstructured.Unstructured) error
	PostUpdate(res *unstructured.Unstructured) error
	PreDelete(res *unstructured.Unstructured) error
	PostDelete(res *unstructured.Unstructured) error
}

func (r *KubeInstaller) ReconcilleResources(ctx context.Context, installNamespace string, resources kuberesource.UnstructuredResources, ownerLabels map[string]string) error {
	if err := r.preInstall(); err != nil {
		return errors.Wrapf(err, "error in pre-install hook")
	}

	if err := r.reconcileResources(ctx, installNamespace, resources.ByKey(), ownerLabels); err != nil {
		return err
	}

	if err := r.postInstall(); err != nil {
		return errors.Wrapf(err, "error in pre-install hook")
	}

	return nil
}

func (r *KubeInstaller) reconcileResources(ctx context.Context, installNamespace string, desiredResources kuberesource.UnstructuredResourcesByKey, ownerLabels map[string]string) error {
	cachedResources := r.cache.List().WithLabels(ownerLabels).ByKey()

	logger := contextutils.LoggerFrom(ctx)

	logger.Infow("reconciling desired resources against cached resources",
		"desired", len(desiredResources),
		"cached_with_label", len(cachedResources),
		"labels", ownerLabels,
		"cache_total", len(r.cache.resources),
	)

	// set labels for writing
	for _, res := range desiredResources {
		labels := res.GetLabels()
		if labels == nil {
			labels = make(map[string]string)
		}
		for k, v := range ownerLabels {
			labels[k] = v
		}
		res.SetLabels(labels)
		res.SetNamespace(installNamespace)
	}

	// determine what must be created, deleted, updated
	var resourcesToDelete, resourcesToCreate, resourcesToUpdate kuberesource.UnstructuredResources
	for key, res := range desiredResources {
		if _, exists := cachedResources[key]; exists {
			resourcesToUpdate = append(resourcesToUpdate, res)
		} else {
			resourcesToCreate = append(resourcesToCreate, res)
		}
	}
	// undesired resources with labels get deleted
	for key, res := range cachedResources {
		if _, desired := desiredResources[key]; !desired {
			resourcesToDelete = append(resourcesToDelete, res)
		}
	}

	logger.Infof("preparing to create %v, update %v, and delete %v resources", len(resourcesToCreate), len(resourcesToUpdate), len(resourcesToDelete))

	// delete in reverse order of install
	groupedResourcesToDelete := resourcesToDelete.GroupedByGVK()
	for i := len(groupedResourcesToDelete); i > 0; i-- {
		group := groupedResourcesToDelete[i-1]
		g := errgroup.Group{}
		for _, res := range group.Resources {
			res := res
			g.Go(func() error {
				if err := r.preDelete(res); err != nil {
					return err
				}
				resKey := fmt.Sprintf("%v %v.%v", res.GroupVersionKind().Kind, res.GetNamespace(), res.GetName())
				logger.Infof("deleting resource %v", resKey)

				if err := retry.Do(func() error { return r.client.Delete(ctx, res) }); err != nil && !kubeerrs.IsNotFound(err) {
					return errors.Wrapf(err, "deleting  %v", resKey)
				}
				if err := r.postDelete(res); err != nil {
					return err
				}
				r.cache.Delete(res)
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
	}

	// create
	for _, group := range resourcesToCreate.GroupedByGVK() {
		// batch create for each resource group
		g := errgroup.Group{}
		for _, res := range group.Resources {
			res := res
			g.Go(func() error {
				if err := r.preCreate(res); err != nil {
					return err
				}
				resKey := fmt.Sprintf("%v %v.%v", res.GroupVersionKind().Kind, res.GetNamespace(), res.GetName())
				logger.Infof("creating resource %v", resKey)

				if err := retry.Do(func() error { return r.client.Create(ctx, res) }); err != nil {
					return errors.Wrapf(err, "creating %v", resKey)
				}
				if err := r.postCreate(res); err != nil {
					return err
				}
				if err := r.waitForResourceReady(ctx, res); err != nil {
					return errors.Wrapf(err, "waiting for resource to become ready %v", resKey)
				}
				r.cache.Set(res)
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
	}

	// update
	for _, group := range resourcesToUpdate.GroupedByGVK() {
		g := errgroup.Group{}
		for _, res := range group.Resources {
			desired := res
			g.Go(func() error {
				key := kuberesource.Key(desired)
				original, ok := cachedResources[key]
				if !ok {
					return errors.Errorf("internal error: could not find original resource for desired key %v", key)
				}
				// don't update the object if there is a match
				if kuberesource.Match(ctx, original, desired) {
					return nil
				}
				if err := r.updateResourceVersion(ctx, desired); err != nil {
					return err
				}
				resKey := fmt.Sprintf("%v %v.%v", desired.GroupVersionKind().Kind, desired.GetNamespace(), desired.GetName())
				logger.Infof("updating resource %v", resKey)

				if err := retry.Do(func() error { return r.client.Update(ctx, desired) }); err != nil {
					return errors.Wrapf(err, "updating %v", resKey)
				}
				if err := r.waitForResourceReady(ctx, desired); err != nil {
					return errors.Wrapf(err, "waiting for resource to become ready %v", resKey)
				}
				r.cache.Set(desired)
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
	}

	logger.Infof("created %v, updated %v, and deleted %v resources", len(resourcesToCreate), len(resourcesToUpdate), len(resourcesToDelete))

	return nil
}

// do an HTTP GET to update the resource version of the desired object
func (r *KubeInstaller) updateResourceVersion(ctx context.Context, res *unstructured.Unstructured) error {
	currentFromServer := res.DeepCopyObject().(*unstructured.Unstructured)
	objectKey := client.ObjectKey{Namespace: res.GetNamespace(), Name: res.GetName()}
	if err := r.client.Get(ctx, objectKey, currentFromServer); err != nil {
		return err
	}
	res.SetResourceVersion(currentFromServer.GetResourceVersion())
	return nil
}

func (r *KubeInstaller) PurgeResources(ctx context.Context, withLabels map[string]string) error {
	return r.reconcileResources(ctx, "", nil, withLabels)
}

func (r *KubeInstaller) waitForResourceReady(ctx context.Context, res *unstructured.Unstructured) error {
	runtimeObject, err := kuberesource.ConvertUnstructured(res)
	if err != nil {
		return nil // not a handled type, possibly a crd
	}
	switch obj := runtimeObject.(type) {
	case *v1beta1.CustomResourceDefinition:
		if err := r.waitForCrd(ctx, obj.Name); err != nil {
			return err
		}
		// refresh the client to get the new rest mappings for the crd
		r.client, err = client.New(r.cfg, client.Options{})
		if err != nil {
			return err
		}
	case *extensionsv1beta1.Deployment:
		return r.waitForDeploymentReplica(ctx, obj.Name, obj.Namespace)
	case *appsv1.Deployment:
		return r.waitForDeploymentReplica(ctx, obj.Name, obj.Namespace)
	case *appsv1beta2.Deployment:
		return r.waitForDeploymentReplica(ctx, obj.Name, obj.Namespace)
	}
	return nil
}

func (r *KubeInstaller) waitForCrd(ctx context.Context, crdName string) error {
	return retry.Do(func() error {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		crd, err := r.apiExtensions.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, v1.GetOptions{})
		if err != nil {
			return errors.Wrapf(err, "lookup crd %v", crdName)
		}

		var established bool
		for _, status := range crd.Status.Conditions {
			if status.Type == v1beta1.Established {
				established = true
				break
			}
		}

		if !established {
			return errors.Errorf("crd %v exists but not yet established by kube", crdName)
		}

		// attempt to do a list on the crd's resources. the above can still give false positives
		_, err = r.dynamic.Resource(schema.GroupVersionResource{
			Group:    crd.Spec.Group,
			Version:  crd.Spec.Version,
			Resource: crd.Spec.Names.Plural,
		}).List(v1.ListOptions{})
		if err != nil {
			return err
		}

		contextutils.LoggerFrom(ctx).Infof("registered crd %v", crd.ObjectMeta)

		return nil
	},
		retry.Delay(time.Millisecond*250),
		retry.DelayType(retry.FixedDelay),
		retry.Attempts(500), // give a considerable amount of time to pull the image
	)
}

func (r *KubeInstaller) waitForDeploymentReplica(ctx context.Context, name, namespace string) error {
	return retry.Do(func() error {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		deployment, err := r.core.AppsV1().Deployments(namespace).Get(name, v1.GetOptions{})
		if err != nil {
			return errors.Wrapf(err, "lookup deployment %v.%v", name, namespace)
		}

		// no replicas to wait for
		if deployment.Spec.Replicas != nil && *deployment.Spec.Replicas == 0 {
			return nil
		}

		// wait for at least one replica to become ready
		if deployment.Status.ReadyReplicas < 1 {
			var condition appsv1.DeploymentCondition
			if len(deployment.Status.Conditions) > 0 {
				condition = deployment.Status.Conditions[0]
			}
			return errors.Errorf("no ready replicas for deployment %v.%v with condition %#v", namespace, name,
				condition)
		}

		contextutils.LoggerFrom(ctx).Infof("deployment %v.%v ready", namespace, name)
		return nil
	},
		retry.Delay(time.Millisecond*250),
		retry.DelayType(retry.FixedDelay),
		retry.Attempts(100),
	)
}

// consider moving to kube utils/errs package?

func IsNoKindMatch(err error) bool {
	_, ok := err.(*meta.NoKindMatchError)
	return ok
}
