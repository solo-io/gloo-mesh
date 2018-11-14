package setup

import (
	"context"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	gloov1 "github.com/solo-io/supergloo/pkg/api/external/gloo/v1"
	"github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"
	prometheusv1 "github.com/solo-io/supergloo/pkg/api/external/prometheus/v1"
	"github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/translator/istio"
	"github.com/solo-io/supergloo/pkg/translator/linkerd2"
	"k8s.io/client-go/kubernetes"
	"time"
)

func Main() error {
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

	prometheusClient, err := prometheusv1.NewConfigClient(&factory.KubeConfigMapClientFactory{
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

	secretClient, err := gloov1.NewSecretClient(&factory.KubeSecretClientFactory{
		Clientset: kubeClient,
	})
	if err != nil {
		return err
	}
	if err := secretClient.Register(); err != nil {
		return err
	}

	emitter := v1.NewTranslatorEmitter(meshClient, upstreamClient, secretClient)

	rpt := reporter.NewReporter("supergloo", meshClient.BaseClient())
	writeErrs := make(chan error)

	istioRoutingSyncer := &istio.RoutingSyncer{
		DestinationRuleReconciler: v1alpha3.NewDestinationRuleReconciler(destinationRuleClient),
		VirtualServiceReconciler:  v1alpha3.NewVirtualServiceReconciler(virtualServiceClient),
		Reporter:                  rpt,
		WriteSelector:             map[string]string{"supergloo.istio.routing": "configured"},
		WriteNamespace:            "supergloo-system",
	}

	linkerd2PrometheusSyncer := &linkerd2.PrometheusSyncer{
		Kube:             kubeClient,
		PrometheusClient: prometheusClient,
	}

	// TODO (rickducott: add consul syncer here)

	syncers := v1.TranslatorSyncers{
		istioRoutingSyncer,
		linkerd2PrometheusSyncer,
	}

	eventLoop := v1.NewTranslatorEventLoop(emitter, syncers)

	ctx := contextutils.WithLogger(context.Background(), "supergloo")
	watchOpts := clients.WatchOpts{
		Ctx:         ctx,
		RefreshRate: time.Second * 1,
	}

	eventLoopErrs, err := eventLoop.Run([]string{"supergloo-system", "gloo-system"}, watchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(watchOpts.Ctx, writeErrs, eventLoopErrs, "event_loop")

	logger := contextutils.LoggerFrom(watchOpts.Ctx)

	go func() {
		for {
			select {
			case err := <-writeErrs:
				logger.Errorf("error: %v", err)
			case <-watchOpts.Ctx.Done():
				close(writeErrs)
				return
			}
		}
	}()
	return nil
}
