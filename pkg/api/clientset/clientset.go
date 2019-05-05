package clientset

import (
	"context"

	linkerdv1 "github.com/solo-io/supergloo/pkg/api/external/linkerd/v1"

	"github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned"
	"github.com/solo-io/solo-kit/pkg/api/external/kubernetes/pod"
	kubernetes2 "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"github.com/solo-io/supergloo/pkg/api/custom/clients/linkerd"

	"github.com/solo-io/supergloo/pkg/api/custom/clients/prometheus"
	promv1 "github.com/solo-io/supergloo/pkg/api/external/prometheus/v1"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	policyv1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
	"github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"
	rbacv1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/rbac/v1alpha1"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// initialize all resource clients here that will share a cache
func ClientsetFromContext(ctx context.Context) (*Clientset, error) {
	restConfig, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, err
	}
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	crdCache := kube.NewKubeCache(ctx)
	kubeCoreCache, err := cache.NewKubeCoreCache(ctx, kubeClient)
	if err != nil {
		return nil, err
	}

	promClient, err := promv1.NewPrometheusConfigClient(prometheus.ResourceClientFactory(kubeClient, kubeCoreCache))
	if err != nil {
		return nil, err
	}

	/*
		supergloo config clients
	*/
	install, err := v1.NewInstallClient(clientForCrd(v1.InstallCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := install.Register(); err != nil {
		return nil, err
	}

	mesh, err := v1.NewMeshClient(clientForCrd(v1.MeshCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := mesh.Register(); err != nil {
		return nil, err
	}

	meshIngress, err := v1.NewMeshIngressClient(clientForCrd(v1.MeshIngressCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := meshIngress.Register(); err != nil {
		return nil, err
	}

	meshGroup, err := v1.NewMeshGroupClient(clientForCrd(v1.MeshGroupCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := meshGroup.Register(); err != nil {
		return nil, err
	}

	upstream, err := gloov1.NewUpstreamClient(clientForCrd(gloov1.UpstreamCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := upstream.Register(); err != nil {
		return nil, err
	}

	routingRule, err := v1.NewRoutingRuleClient(clientForCrd(v1.RoutingRuleCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := routingRule.Register(); err != nil {
		return nil, err
	}

	securityRule, err := v1.NewSecurityRuleClient(clientForCrd(v1.SecurityRuleCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := securityRule.Register(); err != nil {
		return nil, err
	}

	// ilackarms: should we use Kube secret here? these secrets follow a different format (specific to istio)
	tlsSecret, err := v1.NewTlsSecretClient(&factory.KubeSecretClientFactory{
		Clientset:    kubeClient,
		PlainSecrets: true,
		Cache:        kubeCoreCache,
	})
	if err != nil {
		return nil, err
	}
	if err := tlsSecret.Register(); err != nil {
		return nil, err
	}

	secret, err := gloov1.NewSecretClient(&factory.KubeSecretClientFactory{
		Clientset: kubeClient,
		Cache:     kubeCoreCache,
	})
	if err != nil {
		return nil, err
	}
	if err := secret.Register(); err != nil {
		return nil, err
	}

	settings, err := gloov1.NewSettingsClient(clientForCrd(gloov1.SettingsCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := settings.Register(); err != nil {
		return nil, err
	}

	// special resource client wired up to kubernetes pods
	// used by the istio policy syncer to watch pods for service account info
	pods := pod.NewPodClient(kubeClient, kubeCoreCache)

	return newClientset(
		restConfig,
		kubeClient,
		promClient,
		newSuperglooClients(install, mesh, meshGroup, meshIngress, upstream,
			routingRule, securityRule, tlsSecret, secret, settings),
		newDiscoveryClients(pods),
	), nil
}

func IstioFromContext(ctx context.Context) (*IstioClients, error) {
	restConfig, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, err
	}
	crdCache := kube.NewKubeCache(ctx)
	/*
		istio clients
	*/

	rbacConfig, err := rbacv1alpha1.NewRbacConfigClient(clientForCrd(rbacv1alpha1.RbacConfigCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := rbacConfig.Register(); err != nil {
		return nil, err
	}

	serviceRole, err := rbacv1alpha1.NewServiceRoleClient(clientForCrd(rbacv1alpha1.ServiceRoleCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := serviceRole.Register(); err != nil {
		return nil, err
	}

	serviceRoleBinding, err := rbacv1alpha1.NewServiceRoleBindingClient(clientForCrd(rbacv1alpha1.ServiceRoleBindingCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := serviceRoleBinding.Register(); err != nil {
		return nil, err
	}

	meshPolicy, err := policyv1alpha1.NewMeshPolicyClient(clientForCrd(policyv1alpha1.MeshPolicyCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := meshPolicy.Register(); err != nil {
		return nil, err
	}

	destinationRule, err := v1alpha3.NewDestinationRuleClient(clientForCrd(v1alpha3.DestinationRuleCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := destinationRule.Register(); err != nil {
		return nil, err
	}

	virtualService, err := v1alpha3.NewVirtualServiceClient(clientForCrd(v1alpha3.VirtualServiceCrd, restConfig, crdCache))
	if err != nil {
		return nil, err
	}
	if err := virtualService.Register(); err != nil {
		return nil, err
	}
	return newIstioClients(rbacConfig, serviceRole, serviceRoleBinding, meshPolicy, destinationRule, virtualService), nil
}

func serviceProfileClientFromConfig(ctx context.Context, restConfig *rest.Config) (linkerdv1.ServiceProfileClient, error) {
	linkerdClient, err := versioned.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	cache, err := linkerd.NewLinkerdCache(ctx, linkerdClient)
	if err != nil {
		return nil, err
	}
	baseServiceProfileClient := linkerd.NewResourceClient(linkerdClient, cache)

	return linkerdv1.NewServiceProfileClientWithBase(baseServiceProfileClient), nil
}

func LinkerdFromContext(ctx context.Context) (*LinkerdClients, error) {
	restConfig, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, err
	}
	/*
		istio clients
	*/
	serviceProfile, err := serviceProfileClientFromConfig(ctx, restConfig)
	if err != nil {
		return nil, err
	}

	return newLinkerdClients(serviceProfile), nil
}

type Clientset struct {
	RestConfig *rest.Config

	Kube kubernetes.Interface

	Prometheus promv1.PrometheusConfigClient

	// config for supergloo
	Supergloo *SuperglooClients

	// discovery resources from kubernetes
	Discovery *discoveryClients
}

func newClientset(restConfig *rest.Config, kube kubernetes.Interface, prometheus promv1.PrometheusConfigClient, input *SuperglooClients, discovery *discoveryClients) *Clientset {
	return &Clientset{RestConfig: restConfig, Kube: kube, Prometheus: prometheus, Supergloo: input, Discovery: discovery}
}

func clientForCrd(crd crd.Crd, restConfig *rest.Config, kubeCache kube.SharedCache) factory.ResourceClientFactory {
	return &factory.KubeResourceClientFactory{Crd: crd, Cfg: restConfig, SharedCache: kubeCache}
}

type SuperglooClients struct {
	Install      v1.InstallClient
	Mesh         v1.MeshClient
	MeshGroup    v1.MeshGroupClient
	MeshIngress  v1.MeshIngressClient
	Upstream     gloov1.UpstreamClient
	RoutingRule  v1.RoutingRuleClient
	SecurityRule v1.SecurityRuleClient
	TlsSecret    v1.TlsSecretClient
	Secret       gloov1.SecretClient
	Settings     gloov1.SettingsClient
}

func newSuperglooClients(install v1.InstallClient, mesh v1.MeshClient, meshGroup v1.MeshGroupClient,
	meshIngress v1.MeshIngressClient, upstream gloov1.UpstreamClient, routingRule v1.RoutingRuleClient,
	securityRule v1.SecurityRuleClient, tlsSecret v1.TlsSecretClient, secret gloov1.SecretClient, settings gloov1.SettingsClient) *SuperglooClients {
	return &SuperglooClients{Install: install, Mesh: mesh, MeshGroup: meshGroup, MeshIngress: meshIngress,
		Upstream: upstream, RoutingRule: routingRule, SecurityRule: securityRule, TlsSecret: tlsSecret, Secret: secret, Settings: settings}
}

type discoveryClients struct {
	Pod kubernetes2.PodClient
}

func newDiscoveryClients(pod kubernetes2.PodClient) *discoveryClients {
	return &discoveryClients{Pod: pod}
}

type IstioClients struct {
	RbacConfig         rbacv1alpha1.RbacConfigClient
	ServiceRole        rbacv1alpha1.ServiceRoleClient
	ServiceRoleBinding rbacv1alpha1.ServiceRoleBindingClient
	MeshPolicy         policyv1alpha1.MeshPolicyClient
	DestinationRule    v1alpha3.DestinationRuleClient
	VirtualService     v1alpha3.VirtualServiceClient
}

func newIstioClients(rbacConfig rbacv1alpha1.RbacConfigClient, serviceRole rbacv1alpha1.ServiceRoleClient, serviceRoleBinding rbacv1alpha1.ServiceRoleBindingClient, meshPolicy policyv1alpha1.MeshPolicyClient, destinationRule v1alpha3.DestinationRuleClient, virtualService v1alpha3.VirtualServiceClient) *IstioClients {
	return &IstioClients{RbacConfig: rbacConfig, ServiceRole: serviceRole, ServiceRoleBinding: serviceRoleBinding, MeshPolicy: meshPolicy, DestinationRule: destinationRule, VirtualService: virtualService}
}

type LinkerdClients struct {
	ServiceProfile linkerdv1.ServiceProfileClient
}

func newLinkerdClients(serviceProfile linkerdv1.ServiceProfileClient) *LinkerdClients {
	return &LinkerdClients{ServiceProfile: serviceProfile}
}
