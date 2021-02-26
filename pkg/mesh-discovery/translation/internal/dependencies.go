package translation

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh/detector/appmesh"
	appmeshsidecar "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/detector/appmesh"

	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/destination"
	destinationdetector "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/destination/detector"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh"
	meshdetector "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh/detector"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh/detector/istio"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh/detector/osm"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload"
	workloaddetector "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/detector"
	istiosidecar "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/detector/istio"
	osmsidecar "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/detector/osm"
)

// we must generate in the same package because the interface is private
//go:generate mockgen -source ./dependencies.go -destination mocks/dependencies.go

// the dependencyFactory creates dependencies for the translator from a given snapshot
// NOTE(ilackarms): private interface used here as it's not expected we'll need to
// define our dependencyFactory anywhere else
type DependencyFactory interface {
	MakeMeshTranslator(
		ctx context.Context,
	) mesh.Translator

	MakeWorkloadTranslator(
		ctx context.Context,
		in input.DiscoveryInputSnapshot,
	) workload.Translator

	MakeDestinationTranslator() destination.Translator
}

type DependencyFactoryImpl struct{}

func (d DependencyFactoryImpl) MakeMeshTranslator(
	ctx context.Context,
) mesh.Translator {

	detectors := meshdetector.MeshDetectors{
		istio.NewMeshDetector(ctx),
		appmesh.NewMeshDetector(ctx),
		osm.NewMeshDetector(ctx),
	}

	return mesh.NewTranslator(ctx, detectors)
}

func (d DependencyFactoryImpl) MakeWorkloadTranslator(
	ctx context.Context,
	in input.DiscoveryInputSnapshot,
) workload.Translator {
	sidecarDetectors := workloaddetector.SidecarDetectors{
		istiosidecar.NewSidecarDetector(ctx),
		appmeshsidecar.NewSidecarDetector(ctx),
		osmsidecar.NewSidecarDetector(ctx),
	}

	workloadDetector := workloaddetector.NewWorkloadDetector(
		ctx,
		in.Pods(),
		in.ReplicaSets(),
		sidecarDetectors,
	)
	return workload.NewTranslator(ctx, workloadDetector)
}

func (d DependencyFactoryImpl) MakeDestinationTranslator() destination.Translator {
	return destination.NewTranslator(destinationdetector.NewDestinationDetector())

}
