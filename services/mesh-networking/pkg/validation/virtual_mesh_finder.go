package vm_validation

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
)

var (
	InvalidMeshRefsError = func(refs []string) error {
		return eris.Errorf("The following mesh refs are invalid: %v", refs)
	}
)

type virtualMeshFinder struct {
	meshClient zephyr_discovery.MeshClient
}

func NewVirtualMeshFinder(meshClient zephyr_discovery.MeshClient) VirtualMeshFinder {
	return &virtualMeshFinder{meshClient: meshClient}
}

func (v *virtualMeshFinder) GetMeshesForVirtualMesh(
	ctx context.Context,
	virtualMesh *zephyr_networking.VirtualMesh,
) ([]*zephyr_discovery.Mesh, error) {
	meshList, err := v.meshClient.ListMesh(ctx)
	if err != nil {
		return nil, err
	}
	var result []*zephyr_discovery.Mesh
	var invalidRefs []string
	for _, ref := range virtualMesh.Spec.GetMeshes() {
		var foundMesh *zephyr_discovery.Mesh
		for _, mesh := range meshList.Items {
			// thankx rob pike
			mesh := mesh
			if mesh.GetName() == ref.GetName() && mesh.GetNamespace() == ref.GetNamespace() {
				foundMesh = &mesh
			}
		}
		if foundMesh == nil {
			invalidRefs = append(invalidRefs, fmt.Sprintf("%s.%s", ref.GetName(), ref.GetNamespace()))
			continue
		}
		result = append(result, foundMesh)
	}
	if len(invalidRefs) > 0 {
		return nil, InvalidMeshRefsError(invalidRefs)
	}
	return result, nil
}
