package istio

import (
	"context"
	"fmt"

	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/extensions"

	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/local"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	istioextensions "github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/extensions"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/internal"
	"github.com/solo-io/go-utils/contextutils"
)

// the istio translator translates an input networking snapshot to an output snapshot of Istio resources
type Translator interface {
	// Translate translates the appropriate resources to apply input configuration resources for all Istio meshes contained in the input snapshot.
	// Output resources will be added to the output.Builder
	// Errors caused by invalid user config will be reported using the Reporter.
	Translate(
		ctx context.Context,
		in input.LocalSnapshot,
		userSupplied input.RemoteSnapshot,
		istioOutputs istio.Builder,
		localOutputs local.Builder,
		reporter reporting.Reporter,
	)
}

type istioTranslator struct {
	totalTranslates int // TODO(ilackarms): metric

	// note: these interfaces are set directly in unit tests, but not exposed in the Translator's constructor
	dependencies internal.DependencyFactory
	extender     istioextensions.IstioExtender
}

func NewIstioTranslator(extensionClients extensions.Clientset) Translator {
	return &istioTranslator{
		dependencies: internal.NewDependencyFactory(),
		extender:     istioextensions.NewIstioExtender(extensionClients),
	}
}

func (t *istioTranslator) Translate(
	ctx context.Context,
	in input.LocalSnapshot,
	userSupplied input.RemoteSnapshot,
	istioOutputs istio.Builder,
	localOutputs local.Builder,
	reporter reporting.Reporter,
) {
	ctx = contextutils.WithLogger(ctx, fmt.Sprintf("istio-translator-%v", t.totalTranslates))

	destinationTranslator := t.dependencies.MakeDestinationTranslator(
		ctx,
		userSupplied,
		in.KubernetesClusters(),
		in.Destinations(),
	)

	for _, destination := range in.Destinations().List() {
		destinationTranslator.Translate(in, destination, istioOutputs, reporter)
	}

	meshTranslator := t.dependencies.MakeMeshTranslator(
		ctx,
		userSupplied,
		in.KubernetesClusters(),
		in.Secrets(),
		in.Workloads(),
		in.Destinations(),
	)

	for _, mesh := range in.Meshes().List() {
		meshTranslator.Translate(in, mesh, istioOutputs, localOutputs, reporter)
	}

	if err := t.extender.PatchOutputs(ctx, in, istioOutputs); err != nil {
		// TODO(ilackarms): consider providing/checking user option to fail here when the extender server is unavailable.
		// currently we just log the error and continue.
		contextutils.LoggerFrom(ctx).Errorf("failed to apply extension patches: %v", err)
	}

	t.totalTranslates++
}
