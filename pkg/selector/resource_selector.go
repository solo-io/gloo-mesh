package selector

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/stringutils"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/config_lookup"
	core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	discovery_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/service-mesh-hub/pkg/clients"
	kubernetes_apps "github.com/solo-io/service-mesh-hub/pkg/clients/kubernetes/apps"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/discovery"
	"github.com/solo-io/service-mesh-hub/services/common/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	KubeServiceNotFound = func(name, namespace, cluster string) error {
		return eris.Errorf("Kubernetes Service with name: %s, namespace: %s, cluster: %s not found", name, namespace, cluster)
	}
	MultipleMeshServicesFound = func(name, namespace, clusterName string) error {
		return eris.Errorf("Multiple MeshServices found with labels %s=%s, %s=%s, %s=%s",
			constants.KUBE_SERVICE_NAME, name,
			constants.KUBE_SERVICE_NAMESPACE, namespace,
			constants.CLUSTER, clusterName)
	}
	MultipleMeshWorkloadsFound = func(name, namespace, clusterName string) error {
		return eris.Errorf("Multiple MeshWorkloads found with labels %s=%s, %s=%s, %s=%s",
			constants.KUBE_CONTROLLER_NAME, name,
			constants.KUBE_CONTROLLER_NAMESPACE, namespace,
			constants.CLUSTER, clusterName)
	}
	MeshServiceNotFound = func(name, namespace, clusterName string) error {
		return eris.Errorf("No MeshService found with labels %s=%s, %s=%s, %s=%s",
			constants.KUBE_SERVICE_NAME, name,
			constants.KUBE_SERVICE_NAMESPACE, namespace,
			constants.CLUSTER, clusterName)
	}
	MeshWorkloadNotFound = func(name, namespace, clusterName string) error {
		return eris.Errorf("No MeshWorkloads found with labels %s=%s, %s=%s, %s=%s",
			constants.KUBE_CONTROLLER_NAME, name,
			constants.KUBE_CONTROLLER_NAMESPACE, namespace,
			constants.CLUSTER, clusterName)
	}
	MustProvideClusterName = func(ref *core_types.ResourceRef) error {
		return eris.Errorf("Must provide cluster name in ref %+v", ref)
	}
	MissingClusterLabel = func(resourceName string) error {
		return eris.Errorf("Resource '%s' does not have a "+constants.CLUSTER+" label", resourceName)
	}
)

func NewResourceSelector(
	meshServiceClient zephyr_discovery.MeshServiceClient,
	meshWorkloadClient zephyr_discovery.MeshWorkloadClient,
	deploymentClientFactory kubernetes_apps.GeneratedDeploymentClientFactory,
	kubeConfigLookup config_lookup.KubeConfigLookup,
) ResourceSelector {
	return &resourceSelector{
		meshServiceClient:       meshServiceClient,
		meshWorkloadClient:      meshWorkloadClient,
		deploymentClientFactory: deploymentClientFactory,
		kubeConfigLookup:        kubeConfigLookup,
	}
}

type resourceSelector struct {
	meshServiceClient       zephyr_discovery.MeshServiceClient
	meshWorkloadClient      zephyr_discovery.MeshWorkloadClient
	deploymentClientFactory kubernetes_apps.GeneratedDeploymentClientFactory
	kubeConfigLookup        config_lookup.KubeConfigLookup
}

func (b *resourceSelector) GetMeshServiceByRefSelector(
	ctx context.Context,
	kubeServiceName string,
	kubeServiceNamespace string,
	kubeServiceCluster string,
) (*discovery_v1alpha1.MeshService, error) {
	if kubeServiceCluster == "" {
		return nil, MustProvideClusterName(&core_types.ResourceRef{Name: kubeServiceName, Namespace: kubeServiceNamespace})
	}
	destinationKey := client.MatchingLabels{
		constants.KUBE_SERVICE_NAME:      kubeServiceName,
		constants.KUBE_SERVICE_NAMESPACE: kubeServiceNamespace,
		constants.CLUSTER:                kubeServiceCluster,
	}
	meshServiceList, err := b.meshServiceClient.List(ctx, destinationKey)
	if err != nil {
		return nil, err
	}
	// there should only be a single MeshService with the kube Service name/namespace/cluster key
	if len(meshServiceList.Items) > 1 {
		return nil, MultipleMeshServicesFound(kubeServiceName, kubeServiceNamespace, kubeServiceCluster)
	} else if len(meshServiceList.Items) < 1 {
		return nil, MeshServiceNotFound(kubeServiceName, kubeServiceNamespace, kubeServiceCluster)
	}
	return &meshServiceList.Items[0], nil
}

// List all MeshServices and filter for the ones associated with the k8s Services specified in the selector
func (b *resourceSelector) GetMeshServicesByServiceSelector(
	ctx context.Context,
	selector *core_types.ServiceSelector,
) ([]*discovery_v1alpha1.MeshService, error) {
	var selectedMeshServices []*discovery_v1alpha1.MeshService
	meshServiceList, err := b.meshServiceClient.List(ctx)
	if err != nil {
		return nil, err
	}
	allMeshServices := convertServicesToPointerSlice(meshServiceList.Items)
	// select all MeshServices
	if selector.GetServiceSelectorType() == nil {
		return allMeshServices, nil
	}
	switch selectorType := selector.GetServiceSelectorType().(type) {
	case *core_types.ServiceSelector_Matcher_:
		selectedMeshServices = getMeshServicesBySelectorNamespace(
			selector.GetMatcher().GetLabels(),
			selector.GetMatcher().GetNamespaces(),
			selector.GetMatcher().GetClusters(),
			allMeshServices,
		)
	case *core_types.ServiceSelector_ServiceRefs_:
		for _, ref := range selector.GetServiceRefs().GetServices() {
			if ref.GetCluster() == "" {
				return nil, MustProvideClusterName(ref)
			}
			selectedMeshService := getMeshServiceByServiceKey(ref, allMeshServices)
			if selectedMeshService != nil {
				selectedMeshServices = append(selectedMeshServices, selectedMeshService)
			} else {
				// MeshService for referenced k8s Service not found
				return nil, KubeServiceNotFound(ref.GetName(), ref.GetNamespace(), ref.GetCluster())
			}
		}
	default:
		return nil, eris.Errorf("ServiceSelector has unexpected type %T", selectorType)
	}
	return selectedMeshServices, nil
}

func (b *resourceSelector) GetMeshWorkloadsByIdentitySelector(
	ctx context.Context,
	identitySelector *core_types.IdentitySelector,
) ([]*discovery_v1alpha1.MeshWorkload, error) {
	meshWorkloadList, err := b.meshWorkloadClient.List(ctx)
	if err != nil {
		return nil, err
	}

	if identitySelector == nil {
		return convertWorkloadsToPointerSlice(meshWorkloadList.Items), nil
	}

	var matches []*discovery_v1alpha1.MeshWorkload
	for _, workloadIter := range meshWorkloadList.Items {
		workload := workloadIter // careful not to close over the loop var
		switch identitySelector.GetIdentitySelectorType().(type) {
		case *core_types.IdentitySelector_Matcher_:
			namespaces := identitySelector.GetMatcher().GetNamespaces()
			clusters := identitySelector.GetMatcher().GetClusters()

			namespaceMatches := len(namespaces) == 0 || stringutils.ContainsString(workload.Spec.GetKubeController().GetKubeControllerRef().GetNamespace(), namespaces)
			clusterMatches := len(clusters) == 0 || stringutils.ContainsString(workload.Spec.GetKubeController().GetKubeControllerRef().GetCluster(), clusters)

			if namespaceMatches && clusterMatches {
				matches = append(matches, &workload)
			}

		case *core_types.IdentitySelector_ServiceAccountRefs_:
			for _, ref := range identitySelector.GetServiceAccountRefs().GetServiceAccounts() {
				if ref.GetCluster() == "" {
					return nil, MustProvideClusterName(ref)
				}

				if ref.GetNamespace() == workload.Spec.GetKubeController().GetKubeControllerRef().GetNamespace() && ref.GetName() == workload.Spec.GetKubeController().GetServiceAccountName() {
					matches = append(matches, &workload)
				}
			}
		default:
			return nil, eris.Errorf("IdentitySelector has unexpected type %T", identitySelector)
		}
	}

	return matches, nil
}

func (b *resourceSelector) GetMeshWorkloadsByWorkloadSelector(
	ctx context.Context,
	workloadSelector *core_types.WorkloadSelector,
) ([]*discovery_v1alpha1.MeshWorkload, error) {
	meshWorkloadList, err := b.meshWorkloadClient.List(ctx)
	if err != nil {
		return nil, err
	}

	// if a selector was not provided or if both of its field are empty, accept everything
	if workloadSelector == nil || (len(workloadSelector.Labels) == 0 && len(workloadSelector.Namespaces) == 0) {
		return convertWorkloadsToPointerSlice(meshWorkloadList.Items), nil
	}

	var matches []*discovery_v1alpha1.MeshWorkload

	// for each mesh workload we know about:
	//   - load its deployment
	//   - check whether the deployment labels match the selector labels
	//   - check whether the deployment namespace matches the selector namespaces
	for _, meshWorkloadIter := range meshWorkloadList.Items {
		meshWorkload := meshWorkloadIter // careful not to close over the loop var

		clusterName := meshWorkload.Labels[constants.CLUSTER]
		if clusterName == "" {
			return nil, MissingClusterLabel(meshWorkload.GetName())
		}

		config, err := b.kubeConfigLookup.FromCluster(ctx, clusterName)
		if err != nil {
			return nil, err
		}

		deploymentClient, err := b.deploymentClientFactory(config.RestConfig)
		if err != nil {
			return nil, err
		}

		workloadController, err := deploymentClient.Get(ctx, clients.ResourceRefToObjectKey(meshWorkload.Spec.GetKubeController().GetKubeControllerRef()))
		if err != nil {
			return nil, err
		}

		// consider the selector labels a subset of the controller labels if:
		//   the selector did not provide labels, OR
		//   every label in the selector appears with the same value in the controller
		labelsAreSubset := true
		if len(workloadSelector.Labels) > 0 {
			if len(workloadController.Labels) == 0 {
				labelsAreSubset = false
			} else {
				for k, v := range workloadSelector.Labels {
					if workloadController.Labels[k] != v {
						labelsAreSubset = false
						break
					}
				}
			}
		}

		namespaceMatches := len(workloadSelector.Namespaces) == 0 || stringutils.ContainsString(workloadController.GetNamespace(), workloadSelector.Namespaces)

		if labelsAreSubset && namespaceMatches {
			matches = append(matches, &meshWorkload)
		}
	}

	return matches, nil
}

func (b *resourceSelector) GetMeshWorkloadByRefSelector(
	ctx context.Context,
	podControllerName string,
	podControllerNamespace string,
	podControllerCluster string,
) (*discovery_v1alpha1.MeshWorkload, error) {
	if podControllerCluster == "" {
		return nil, MustProvideClusterName(&core_types.ResourceRef{Name: podControllerName, Namespace: podControllerNamespace})
	}
	destinationKey := client.MatchingLabels{
		constants.KUBE_CONTROLLER_NAME:      podControllerName,
		constants.KUBE_CONTROLLER_NAMESPACE: podControllerNamespace,
		constants.CLUSTER:                   podControllerCluster,
	}
	meshWorkloadList, err := b.meshWorkloadClient.List(ctx, destinationKey)
	if err != nil {
		return nil, err
	}
	// there should only be a single MeshService with the kube Service name/namespace/cluster key
	if len(meshWorkloadList.Items) > 1 {
		return nil, MultipleMeshWorkloadsFound(podControllerName, podControllerNamespace, podControllerCluster)
	} else if len(meshWorkloadList.Items) < 1 {
		return nil, MeshWorkloadNotFound(podControllerName, podControllerNamespace, podControllerCluster)
	}
	return &meshWorkloadList.Items[0], nil
}

func getMeshServiceByServiceKey(
	selectedRef *core_types.ResourceRef,
	meshServices []*discovery_v1alpha1.MeshService,
) *discovery_v1alpha1.MeshService {
	for _, meshService := range meshServices {
		kubeServiceRef := meshService.Spec.GetKubeService().GetRef()
		if selectedRef.GetName() == kubeServiceRef.GetName() &&
			selectedRef.GetNamespace() == kubeServiceRef.GetNamespace() &&
			selectedRef.GetCluster() == kubeServiceRef.GetCluster() {
			return meshService
		}
	}
	return nil
}

func getMeshServicesBySelectorNamespace(
	selectors map[string]string,
	namespaces []string,
	clusters []string,
	meshServices []*discovery_v1alpha1.MeshService,
) []*discovery_v1alpha1.MeshService {
	var selectedMeshServices []*discovery_v1alpha1.MeshService
	for _, meshService := range meshServices {
		kubeService := meshService.Spec.GetKubeService()
		if kubeServiceMatches(selectors, namespaces, clusters, kubeService) {
			selectedMeshServices = append(selectedMeshServices, meshService)
		}
	}
	return selectedMeshServices
}

/* For a k8s Service to match:
1) If labels is specified, all labels must exist on the k8s Service
2) If namespaces is specified, the k8s must be in one of those namespaces
3) The k8s Service must exist in the specified cluster. If cluster is empty, select across all clusters.
*/
func kubeServiceMatches(
	labels map[string]string,
	namespaces []string,
	clusters []string,
	kubeService *types.MeshServiceSpec_KubeService,
) bool {
	if len(namespaces) > 0 && !stringutils.ContainsString(kubeService.GetRef().GetNamespace(), namespaces) {
		return false
	}
	for k, v := range labels {
		serviceLabelValue, ok := kubeService.GetLabels()[k]
		if !ok || serviceLabelValue != v {
			return false
		}
	}
	if len(clusters) > 0 && !stringutils.ContainsString(kubeService.GetRef().GetCluster(), clusters) {
		return false
	}
	return true
}

func convertServicesToPointerSlice(meshServices []discovery_v1alpha1.MeshService) []*discovery_v1alpha1.MeshService {
	pointerSlice := make([]*discovery_v1alpha1.MeshService, 0, len(meshServices))
	for _, meshService := range meshServices {
		meshService := meshService
		pointerSlice = append(pointerSlice, &meshService)
	}
	return pointerSlice
}

func convertWorkloadsToPointerSlice(meshWorkloads []discovery_v1alpha1.MeshWorkload) []*discovery_v1alpha1.MeshWorkload {
	pointerSlice := []*discovery_v1alpha1.MeshWorkload{}
	for _, meshWorkloadIter := range meshWorkloads {
		meshWorkload := meshWorkloadIter
		pointerSlice = append(pointerSlice, &meshWorkload)
	}
	return pointerSlice
}
