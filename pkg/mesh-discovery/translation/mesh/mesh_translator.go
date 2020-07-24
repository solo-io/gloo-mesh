package mesh

import (
	"context"

	appsv1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	"github.com/solo-io/go-utils/contextutils"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	"github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/translation/mesh/detector"
	"github.com/solo-io/skv2/contrib/pkg/sets"
)

//go:generate mockgen -source ./mesh_translator.go -destination mocks/mesh_translator.go

// the mesh translator converts deployments with control plane images into Mesh CRs
type Translator interface {
	TranslateMeshes(deployments appsv1sets.DeploymentSet) v1alpha2sets.MeshSet
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

func (t *translator) TranslateMeshes(deployments appsv1sets.DeploymentSet) v1alpha2sets.MeshSet {
	meshSet := v1alpha2sets.NewMeshSet()
	for _, deployment := range deployments.List() {
		mesh, err := t.meshDetector.DetectMesh(deployment)
		if err != nil {
			contextutils.LoggerFrom(t.ctx).Warnw("failed to discover mesh for deployment ", "deployment", sets.Key(deployment))
		}
		if mesh == nil {
			continue
		}
		contextutils.LoggerFrom(t.ctx).Debugf("detected mesh service %v", sets.Key(mesh))
		meshSet.Insert(mesh)
	}
	return meshSet
}
