package output

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/fieldutils"

	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/output/discovery"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
)

// the istio Applier applies a Snapshot of resources across clusters
type Applier interface {
	Apply(ctx context.Context, cli multicluster.Client, in input.Snapshot, out discovery.Snapshot) error
}

type applier struct {
	fieldOwners fieldutils.FieldOwnershipRegistry
}

func (a applier) HandleWriteError(resource ezkube.Object, err error) {
	//owners := a.fieldOwners.GetRegisteredOwnerships(resource)
	//for _, owner := range owners {
	//}
}

func (a applier) HandleDeleteError(resource ezkube.Object, err error) {
	panic("implement me")
}

func (a applier) HandleListError(err error) {
	panic("implement me")
}
