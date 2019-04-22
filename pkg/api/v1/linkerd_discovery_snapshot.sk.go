// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type LinkerdDiscoverySnapshot struct {
	Meshes    MeshesByNamespace
	Installs  InstallsByNamespace
	Pods      PodsByNamespace
	Upstreams gloo_solo_io.UpstreamsByNamespace
}

func (s LinkerdDiscoverySnapshot) Clone() LinkerdDiscoverySnapshot {
	return LinkerdDiscoverySnapshot{
		Meshes:    s.Meshes.Clone(),
		Installs:  s.Installs.Clone(),
		Pods:      s.Pods.Clone(),
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s LinkerdDiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMeshes(),
		s.hashInstalls(),
		s.hashPods(),
		s.hashUpstreams(),
	)
}

func (s LinkerdDiscoverySnapshot) hashMeshes() uint64 {
	return hashutils.HashAll(s.Meshes.List().AsInterfaces()...)
}

func (s LinkerdDiscoverySnapshot) hashInstalls() uint64 {
	return hashutils.HashAll(s.Installs.List().AsInterfaces()...)
}

func (s LinkerdDiscoverySnapshot) hashPods() uint64 {
	return hashutils.HashAll(s.Pods.List().AsInterfaces()...)
}

func (s LinkerdDiscoverySnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.List().AsInterfaces()...)
}

func (s LinkerdDiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("meshes", s.hashMeshes()))
	fields = append(fields, zap.Uint64("installs", s.hashInstalls()))
	fields = append(fields, zap.Uint64("pods", s.hashPods()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type LinkerdDiscoverySnapshotStringer struct {
	Version   uint64
	Meshes    []string
	Installs  []string
	Pods      []string
	Upstreams []string
}

func (ss LinkerdDiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("LinkerdDiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Meshes %v\n", len(ss.Meshes))
	for _, name := range ss.Meshes {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Installs %v\n", len(ss.Installs))
	for _, name := range ss.Installs {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Pods %v\n", len(ss.Pods))
	for _, name := range ss.Pods {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s LinkerdDiscoverySnapshot) Stringer() LinkerdDiscoverySnapshotStringer {
	return LinkerdDiscoverySnapshotStringer{
		Version:   s.Hash(),
		Meshes:    s.Meshes.List().NamespacesDotNames(),
		Installs:  s.Installs.List().NamespacesDotNames(),
		Pods:      s.Pods.List().NamespacesDotNames(),
		Upstreams: s.Upstreams.List().NamespacesDotNames(),
	}
}
