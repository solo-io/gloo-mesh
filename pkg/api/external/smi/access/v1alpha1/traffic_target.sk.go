// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"sort"

	github_com_solo_io_supergloo_api_external_smi_access "github.com/solo-io/supergloo/api/external/smi/access"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
)

func NewTrafficTarget(namespace, name string) *TrafficTarget {
	traffictarget := &TrafficTarget{}
	traffictarget.TrafficTarget.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return traffictarget
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableTrafficTarget interface {
	resources.Resource
	Clone() *github_com_solo_io_supergloo_api_external_smi_access.TrafficTarget
}

var _ CloneableTrafficTarget = &github_com_solo_io_supergloo_api_external_smi_access.TrafficTarget{}

type TrafficTarget struct {
	github_com_solo_io_supergloo_api_external_smi_access.TrafficTarget
}

func (r *TrafficTarget) Clone() resources.Resource {
	return &TrafficTarget{TrafficTarget: *r.TrafficTarget.Clone()}
}

func (r *TrafficTarget) Hash() uint64 {
	clone := r.TrafficTarget.Clone()

	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})

	return hashutils.HashAll(clone)
}

type TrafficTargetList []*TrafficTarget

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list TrafficTargetList) Find(namespace, name string) (*TrafficTarget, error) {
	for _, trafficTarget := range list {
		if trafficTarget.GetMetadata().Name == name {
			if namespace == "" || trafficTarget.GetMetadata().Namespace == namespace {
				return trafficTarget, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find trafficTarget %v.%v", namespace, name)
}

func (list TrafficTargetList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, trafficTarget := range list {
		ress = append(ress, trafficTarget)
	}
	return ress
}

func (list TrafficTargetList) Names() []string {
	var names []string
	for _, trafficTarget := range list {
		names = append(names, trafficTarget.GetMetadata().Name)
	}
	return names
}

func (list TrafficTargetList) NamespacesDotNames() []string {
	var names []string
	for _, trafficTarget := range list {
		names = append(names, trafficTarget.GetMetadata().Namespace+"."+trafficTarget.GetMetadata().Name)
	}
	return names
}

func (list TrafficTargetList) Sort() TrafficTargetList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list TrafficTargetList) Clone() TrafficTargetList {
	var trafficTargetList TrafficTargetList
	for _, trafficTarget := range list {
		trafficTargetList = append(trafficTargetList, resources.Clone(trafficTarget).(*TrafficTarget))
	}
	return trafficTargetList
}

func (list TrafficTargetList) Each(f func(element *TrafficTarget)) {
	for _, trafficTarget := range list {
		f(trafficTarget)
	}
}

func (list TrafficTargetList) EachResource(f func(element resources.Resource)) {
	for _, trafficTarget := range list {
		f(trafficTarget)
	}
}

func (list TrafficTargetList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *TrafficTarget) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}
