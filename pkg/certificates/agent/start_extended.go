package agent

import (
	"context"

	corev1clients "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	pod_bouncer "github.com/solo-io/gloo-mesh/pkg/certificates/agent/reconciliation/pod-bouncer"
	"github.com/solo-io/gloo-mesh/pkg/certificates/agent/translation"
	"github.com/solo-io/skv2/pkg/bootstrap"
)

// Options for extending the functionality of the Networking controller
type ExtensionOpts struct {
	CertAgentReconciler CertAgentReconcilerExtensionOpts
}

type MakeExtensionOpts func(ctx context.Context, parameters bootstrap.StartParameters) ExtensionOpts

func (opts *ExtensionOpts) initDefaults(parameters bootstrap.StartParameters) {
	opts.CertAgentReconciler.initDefaults(parameters)
}

// Options for overriding functionality of the Networking Reconciler
type CertAgentReconcilerExtensionOpts struct {

	// Hook to override Translator used by Networking Reconciler
	MakeTranslator func(translator translation.Translator) translation.Translator
	// Pod Bouncer to be used by translator, allows overriding the dependency
	PodBouncer pod_bouncer.PodBouncer
}

func (opts *CertAgentReconcilerExtensionOpts) initDefaults(parameters bootstrap.StartParameters) {

	if opts.MakeTranslator == nil {
		// use default translator
		opts.MakeTranslator = func(translator translation.Translator) translation.Translator {
			return translator
		}
	}

	if opts.PodBouncer == nil {
		opts.PodBouncer = pod_bouncer.NewPodBouncer(
			corev1clients.NewPodClient(parameters.MasterManager.GetClient()),
			pod_bouncer.NewSecretRootCertMatcher(),
		)
	}
}
