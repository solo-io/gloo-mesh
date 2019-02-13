package setup

import (
	"context"
	"time"

	kube_client "github.com/solo-io/supergloo/pkg2/kube"
	"github.com/solo-io/supergloo/pkg2/secret"

	"github.com/solo-io/supergloo/pkg2/translator/appmesh"

	factory2 "github.com/solo-io/supergloo/pkg2/factory"

	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/supergloo/pkg2/api/external/istio/networking/v1alpha3"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	prometheusv1 "github.com/solo-io/supergloo/pkg2/api/external/prometheus/v1"
	v1 "github.com/solo-io/supergloo/pkg2/api/v1"
	"github.com/solo-io/supergloo/pkg2/install"
	"github.com/solo-io/supergloo/pkg2/translator/consul"
	"github.com/solo-io/supergloo/pkg2/translator/istio"
	"github.com/solo-io/supergloo/pkg2/translator/linkerd2"
	apiexts "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
)

var defaultNamespaces = []string{"supergloo-system", "gloo-system", "default"}

func Main(errHandler func(error), namespaces ...string) error {
	if len(namespaces) == 0 {
		namespaces = defaultNamespaces
	}
	// TODO: ilackarms: suport options
	kubeCache := kube.NewKubeCache()
	restConfig, err := kubeutils.GetConfig("", "")
	if err != nil {
		return err
	}
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	destinationRuleClient, err := v1alpha3.NewDestinationRuleClient(&factory.KubeResourceClientFactory{
		Crd:         v1alpha3.DestinationRuleCrd,
		Cfg:         restConfig,
		SharedCache: kubeCache,
	})
	if err != nil {
		return err
	}
	if err := destinationRuleClient.Register(); err != nil {
		return err
	}

	virtualServiceClient, err := v1alpha3.NewVirtualServiceClient(&factory.KubeResourceClientFactory{
		Crd:         v1alpha3.VirtualServiceCrd,
		Cfg:         restConfig,
		SharedCache: kubeCache,
	})
	if err != nil {
		return err
	}
	if err := virtualServiceClient.Register(); err != nil {
		return err
	}

	prometheusClient, err := prometheusv1.NewPrometheusConfigClient(&factory.KubeConfigMapClientFactory{
		Clientset: kubeClient,
	})
	if err != nil {
		return err
	}
	if err := prometheusClient.Register(); err != nil {
		return err
	}

	meshClient, err := v1.NewMeshClient(&factory.KubeResourceClientFactory{
		Crd:         v1.MeshCrd,
		Cfg:         restConfig,
		SharedCache: kubeCache,
	})
	if err != nil {
		return err
	}
	if err := meshClient.Register(); err != nil {
		return err
	}

	installClient, err := v1.NewInstallClient(&factory.KubeResourceClientFactory{
		Crd:         v1.InstallCrd,
		Cfg:         restConfig,
		SharedCache: kubeCache,
	})
	if err != nil {
		return err
	}
	if err := installClient.Register(); err != nil {
		return err
	}

	routingRuleClient, err := v1.NewRoutingRuleClient(&factory.KubeResourceClientFactory{
		Crd:         v1.RoutingRuleCrd,
		Cfg:         restConfig,
		SharedCache: kubeCache,
	})
	if err != nil {
		return err
	}
	if err := routingRuleClient.Register(); err != nil {
		return err
	}

	upstreamClient, err := gloov1.NewUpstreamClient(&factory.KubeResourceClientFactory{
		Crd:         gloov1.UpstreamCrd,
		Cfg:         restConfig,
		SharedCache: kubeCache,
	})
	if err != nil {
		return err
	}
	if err := upstreamClient.Register(); err != nil {
		return err
	}

	istioSecretClient, err := factory2.GetIstioCacertsSecretClient(kubeClient)
	if err != nil {
		return err
	}
	if err := istioSecretClient.Register(); err != nil {
		return err
	}

	glooSecretClient, err := gloov1.NewSecretClient(&factory.KubeSecretClientFactory{
		Clientset: kubeClient,
	})
	if err != nil {
		return err
	}
	if err := glooSecretClient.Register(); err != nil {
		return err
	}

	installEmitter := v1.NewInstallEmitter(installClient, istioSecretClient)

	translatorEmitter := v1.NewTranslatorEmitter(
		glooSecretClient,
		upstreamClient,
		meshClient,
		routingRuleClient,
		istioSecretClient,
	)

	rpt := reporter.NewReporter("supergloo", meshClient.BaseClient())
	writeErrs := make(chan error)

	istioRoutingSyncer := istio.NewMeshRoutingSyncer(namespaces,
		nil, // if we run multiple syncers, set this to prevent a conflict / race
		v1alpha3.NewDestinationRuleReconciler(destinationRuleClient),
		v1alpha3.NewVirtualServiceReconciler(virtualServiceClient),
		rpt)

	linkerd2PrometheusSyncer := linkerd2.NewPrometheusSyncer(kubeClient, prometheusClient)
	istioPrometheusSyncer := istio.NewPrometheusSyncer(kubeClient, prometheusClient)

	consulEncryptionSyncer := &consul.ConsulSyncer{}
	consulPolicySyncer := &consul.PolicySyncer{}

	secretClient := kube_client.NewKubeSecretClient(kubeClient)
	podClient := kube_client.NewKubePodClient(kubeClient)

	secretSyncer := &secret.KubeSecretSyncer{
		SecretClient:      secretClient,
		PodClient:         podClient,
		IstioSecretClient: istioSecretClient,
	}
	istioEncryptionSyncer := &istio.EncryptionSyncer{
		SecretSyncer: secretSyncer,
	}
	istioPolicySyncer, err := istio.NewPolicySyncer("supergloo-system", kubeCache, restConfig)
	if err != nil {
		return err
	}

	appMeshSyncer := appmesh.NewSyncer()

	translatorSyncers := v1.TranslatorSyncers{
		appMeshSyncer,
		istioRoutingSyncer,
		istioPrometheusSyncer,
		linkerd2PrometheusSyncer,
		consulEncryptionSyncer,
		consulPolicySyncer,
		istioEncryptionSyncer,
		istioPolicySyncer,
	}

	apiExts, err := apiexts.NewForConfig(restConfig)
	if err != nil {
		return errors.Wrapf(err, "creating api extensions client")
	}
	installSyncer, err := install.NewKubeInstallSyncer(meshClient, istioSecretClient, kubeClient, apiExts)
	if err != nil {
		return errors.Wrapf(err, "creating kube install syncer")
	}
	installSyncers := v1.InstallSyncers{
		installSyncer,
	}

	translatorEventLoop := v1.NewTranslatorEventLoop(translatorEmitter, translatorSyncers)
	installEventLoop := v1.NewInstallEventLoop(installEmitter, installSyncers)

	ctx := contextutils.WithLogger(context.Background(), "supergloo")
	watchOpts := clients.WatchOpts{
		Ctx:         ctx,
		RefreshRate: time.Second * 1,
	}

	translatorEventLoopErrs, err := translatorEventLoop.Run(namespaces, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, writeErrs, translatorEventLoopErrs, "translator_event_loop")

	installEventLoopErrs, err := installEventLoop.Run(namespaces, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, writeErrs, installEventLoopErrs, "install_event_loop")

	logger := contextutils.LoggerFrom(watchOpts.Ctx)

	for {
		select {
		case err := <-writeErrs:
			logger.Errorf("error: %v", err)
			if errHandler != nil {
				errHandler(err)
			}
		case <-watchOpts.Ctx.Done():
			close(writeErrs)
			return nil
		}
	}
}
