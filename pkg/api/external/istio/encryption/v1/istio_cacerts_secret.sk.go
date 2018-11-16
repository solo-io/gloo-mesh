// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewIstioCacertsSecret(namespace, name string) *IstioCacertsSecret {
	return &IstioCacertsSecret{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *IstioCacertsSecret) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type IstioCacertsSecretList []*IstioCacertsSecret
type IstiocertsByNamespace map[string]IstioCacertsSecretList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list IstioCacertsSecretList) Find(namespace, name string) (*IstioCacertsSecret, error) {
	for _, istioCacertsSecret := range list {
		if istioCacertsSecret.Metadata.Name == name {
			if namespace == "" || istioCacertsSecret.Metadata.Namespace == namespace {
				return istioCacertsSecret, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find istioCacertsSecret %v.%v", namespace, name)
}

func (list IstioCacertsSecretList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, istioCacertsSecret := range list {
		ress = append(ress, istioCacertsSecret)
	}
	return ress
}

func (list IstioCacertsSecretList) Names() []string {
	var names []string
	for _, istioCacertsSecret := range list {
		names = append(names, istioCacertsSecret.Metadata.Name)
	}
	return names
}

func (list IstioCacertsSecretList) NamespacesDotNames() []string {
	var names []string
	for _, istioCacertsSecret := range list {
		names = append(names, istioCacertsSecret.Metadata.Namespace+"."+istioCacertsSecret.Metadata.Name)
	}
	return names
}

func (list IstioCacertsSecretList) Sort() IstioCacertsSecretList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list IstioCacertsSecretList) Clone() IstioCacertsSecretList {
	var istioCacertsSecretList IstioCacertsSecretList
	for _, istioCacertsSecret := range list {
		istioCacertsSecretList = append(istioCacertsSecretList, proto.Clone(istioCacertsSecret).(*IstioCacertsSecret))
	}
	return istioCacertsSecretList
}

func (list IstioCacertsSecretList) ByNamespace() IstiocertsByNamespace {
	byNamespace := make(IstiocertsByNamespace)
	for _, istioCacertsSecret := range list {
		byNamespace.Add(istioCacertsSecret)
	}
	return byNamespace
}

func (byNamespace IstiocertsByNamespace) Add(istioCacertsSecret ...*IstioCacertsSecret) {
	for _, item := range istioCacertsSecret {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace IstiocertsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace IstiocertsByNamespace) List() IstioCacertsSecretList {
	var list IstioCacertsSecretList
	for _, istioCacertsSecretList := range byNamespace {
		list = append(list, istioCacertsSecretList...)
	}
	return list.Sort()
}

func (byNamespace IstiocertsByNamespace) Clone() IstiocertsByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &IstioCacertsSecret{}

// Kubernetes Adapter for IstioCacertsSecret

func (o *IstioCacertsSecret) GetObjectKind() schema.ObjectKind {
	t := IstioCacertsSecretCrd.TypeMeta()
	return &t
}

func (o *IstioCacertsSecret) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*IstioCacertsSecret)
}

var IstioCacertsSecretCrd = crd.NewCrd("encryption.istio.io",
	"istiocerts",
	"encryption.istio.io",
	"v1",
	"IstioCacertsSecret",
	"ics",
	&IstioCacertsSecret{})
