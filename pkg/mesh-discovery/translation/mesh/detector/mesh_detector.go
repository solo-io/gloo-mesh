package detector

import (
	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
)

// a MeshDetector detects Mesh control plane instances (e.g. Pilot)
// in a snapshot of discovery resources.
// If detection fails, an error is returned
// If no mesh is detected, an empty list is returned
// Separate Detectors are implemented for different Mesh types / versions.
type MeshDetector interface {
	DetectMeshes(in input.RemoteSnapshot) (v1alpha2.MeshSlice, error)
}

// wrapper for multiple mesh detectors.
// returns all detected meshes
type MeshDetectors []MeshDetector

func (d MeshDetectors) DetectMeshes(in input.RemoteSnapshot) (v1alpha2.MeshSlice, error) {
	var allMeshes v1alpha2.MeshSlice
	var errs error
	for _, detector := range d {
		meshes, err := detector.DetectMeshes(in)
		if err != nil {
			errs = multierror.Append(errs, err)
		}
		allMeshes = append(allMeshes, meshes...)
	}
	return allMeshes, errs
}
