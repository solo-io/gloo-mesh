package setup

import (
	"context"

	"github.com/solo-io/supergloo/pkg/meshdiscovery/pkg/clientset"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/pkg/config"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/pkg/config/istio"
	"github.com/solo-io/supergloo/pkg/registration"
)

func NewDiscoveryConfigLoopStarters(clientset *clientset.Clientset) registration.ConfigLoopStarters {
	starters := createConfigStarters(clientset)
	return starters
}

func createConfigStarters(cs *clientset.Clientset) registration.ConfigLoopStarters {
	return registration.ConfigLoopStarters{
		createIstioConfigLoopStarter(cs),
	}
}

func createIstioConfigLoopStarter(cs *clientset.Clientset) registration.ConfigLoopStarter {
	return func(ctx context.Context, enabled registration.EnabledConfigLoops) (config.EventLoop, error) {
		if enabled.Istio {
			return istio.NewIstioConfigDiscoveryRunner(ctx, cs)
		}
		return nil, nil
	}
}
