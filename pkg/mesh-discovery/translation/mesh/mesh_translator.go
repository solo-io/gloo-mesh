package mesh

import (
	"context"

	settingsv1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"

	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"

	v1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh/detector"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
)

//go:generate mockgen -source ./mesh_translator.go -destination mocks/mesh_translator.go

// the mesh translator converts deployments with control plane images into Mesh CRs
type Translator interface {
	TranslateMeshes(in input.DiscoveryInputSnapshot, settings *settingsv1.DiscoverySettings) v1sets.MeshSet
}

type translator struct {
	ctx          context.Context
	meshDetector detector.MeshDetector
}

func NewTranslator(
	ctx context.Context,
	meshDetector detector.MeshDetector,
) Translator {
	ctx = contextutils.WithLogger(ctx, "mesh-translator")
	return &translator{ctx: ctx, meshDetector: meshDetector}
}

func (t *translator) TranslateMeshes(in input.DiscoveryInputSnapshot, settings *settingsv1.DiscoverySettings) v1sets.MeshSet {
	meshSet := v1sets.NewMeshSet()
	meshes, err := t.meshDetector.DetectMeshes(in, settings)
	if err != nil {
		contextutils.LoggerFrom(t.ctx).Warnw("ecnountered error discovering meshes", "err", err)
	}
	for _, mesh := range meshes {
		contextutils.LoggerFrom(t.ctx).Debugf("detected mesh %v", sets.Key(mesh))
		meshSet.Insert(mesh)
	}
	return meshSet
}
