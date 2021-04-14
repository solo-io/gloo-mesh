// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets

import (
	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type WasmDeploymentSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Insert a resource into the set.
	Insert(wasmDeployment ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(wasmDeploymentSet WasmDeploymentSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(wasmDeployment ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(wasmDeployment ezkube.ResourceId)
	// Return the union with the provided set
	Union(set WasmDeploymentSet) WasmDeploymentSet
	// Return the difference with the provided set
	Difference(set WasmDeploymentSet) WasmDeploymentSet
	// Return the intersection with the provided set
	Intersection(set WasmDeploymentSet) WasmDeploymentSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another WasmDeploymentSet
	Delta(newSet WasmDeploymentSet) sksets.ResourceDelta
}

func makeGenericWasmDeploymentSet(wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range wasmDeploymentList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type wasmDeploymentSet struct {
	set sksets.ResourceSet
}

func NewWasmDeploymentSet(wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) WasmDeploymentSet {
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(wasmDeploymentList)}
}

func NewWasmDeploymentSetFromList(wasmDeploymentList *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentList) WasmDeploymentSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, 0, len(wasmDeploymentList.Items))
	for idx := range wasmDeploymentList.Items {
		list = append(list, &wasmDeploymentList.Items[idx])
	}
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(list)}
}

func (s *wasmDeploymentSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *wasmDeploymentSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
		})
	}

	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	for _, obj := range s.Generic().List(genericFilters...) {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return wasmDeploymentList
}

func (s *wasmDeploymentSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	}
	return newMap
}

func (s *wasmDeploymentSet) Insert(
	wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range wasmDeploymentList {
		s.Generic().Insert(obj)
	}
}

func (s *wasmDeploymentSet) Has(wasmDeployment ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(wasmDeployment)
}

func (s *wasmDeploymentSet) Equal(
	wasmDeploymentSet WasmDeploymentSet,
) bool {
	if s == nil {
		return wasmDeploymentSet == nil
	}
	return s.Generic().Equal(wasmDeploymentSet.Generic())
}

func (s *wasmDeploymentSet) Delete(WasmDeployment ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(WasmDeployment)
}

func (s *wasmDeploymentSet) Union(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	return NewWasmDeploymentSet(append(s.List(), set.List()...)...)
}

func (s *wasmDeploymentSet) Difference(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &wasmDeploymentSet{set: newSet}
}

func (s *wasmDeploymentSet) Intersection(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	for _, obj := range newSet.List() {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return NewWasmDeploymentSet(wasmDeploymentList...)
}

func (s *wasmDeploymentSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find WasmDeployment %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment), nil
}

func (s *wasmDeploymentSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *wasmDeploymentSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *wasmDeploymentSet) Delta(newSet WasmDeploymentSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type VirtualDestinationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Insert a resource into the set.
	Insert(virtualDestination ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualDestinationSet VirtualDestinationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualDestination ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualDestination ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualDestinationSet) VirtualDestinationSet
	// Return the difference with the provided set
	Difference(set VirtualDestinationSet) VirtualDestinationSet
	// Return the intersection with the provided set
	Intersection(set VirtualDestinationSet) VirtualDestinationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualDestinationSet
	Delta(newSet VirtualDestinationSet) sksets.ResourceDelta
}

func makeGenericVirtualDestinationSet(virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualDestinationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualDestinationSet struct {
	set sksets.ResourceSet
}

func NewVirtualDestinationSet(virtualDestinationList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) VirtualDestinationSet {
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(virtualDestinationList)}
}

func NewVirtualDestinationSetFromList(virtualDestinationList *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationList) VirtualDestinationSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, 0, len(virtualDestinationList.Items))
	for idx := range virtualDestinationList.Items {
		list = append(list, &virtualDestinationList.Items[idx])
	}
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(list)}
}

func (s *virtualDestinationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualDestinationSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
		})
	}

	var virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	for _, obj := range s.Generic().List(genericFilters...) {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	}
	return newMap
}

func (s *virtualDestinationSet) Insert(
	virtualDestinationList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualDestinationList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualDestinationSet) Has(virtualDestination ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualDestination)
}

func (s *virtualDestinationSet) Equal(
	virtualDestinationSet VirtualDestinationSet,
) bool {
	if s == nil {
		return virtualDestinationSet == nil
	}
	return s.Generic().Equal(virtualDestinationSet.Generic())
}

func (s *virtualDestinationSet) Delete(VirtualDestination ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualDestination)
}

func (s *virtualDestinationSet) Union(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	return NewVirtualDestinationSet(append(s.List(), set.List()...)...)
}

func (s *virtualDestinationSet) Difference(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualDestinationSet{set: newSet}
}

func (s *virtualDestinationSet) Intersection(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	for _, obj := range newSet.List() {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return NewVirtualDestinationSet(virtualDestinationList...)
}

func (s *virtualDestinationSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualDestination %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination), nil
}

func (s *virtualDestinationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualDestinationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualDestinationSet) Delta(newSet VirtualDestinationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type FederatedGatewaySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway
	// Insert a resource into the set.
	Insert(federatedGateway ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(federatedGatewaySet FederatedGatewaySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(federatedGateway ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(federatedGateway ezkube.ResourceId)
	// Return the union with the provided set
	Union(set FederatedGatewaySet) FederatedGatewaySet
	// Return the difference with the provided set
	Difference(set FederatedGatewaySet) FederatedGatewaySet
	// Return the intersection with the provided set
	Intersection(set FederatedGatewaySet) FederatedGatewaySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another FederatedGatewaySet
	Delta(newSet FederatedGatewaySet) sksets.ResourceDelta
}

func makeGenericFederatedGatewaySet(federatedGatewayList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range federatedGatewayList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type federatedGatewaySet struct {
	set sksets.ResourceSet
}

func NewFederatedGatewaySet(federatedGatewayList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway) FederatedGatewaySet {
	return &federatedGatewaySet{set: makeGenericFederatedGatewaySet(federatedGatewayList)}
}

func NewFederatedGatewaySetFromList(federatedGatewayList *networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGatewayList) FederatedGatewaySet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway, 0, len(federatedGatewayList.Items))
	for idx := range federatedGatewayList.Items {
		list = append(list, &federatedGatewayList.Items[idx])
	}
	return &federatedGatewaySet{set: makeGenericFederatedGatewaySet(list)}
}

func (s *federatedGatewaySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *federatedGatewaySet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway))
		})
	}

	var federatedGatewayList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway
	for _, obj := range s.Generic().List(genericFilters...) {
		federatedGatewayList = append(federatedGatewayList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway))
	}
	return federatedGatewayList
}

func (s *federatedGatewaySet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway)
	}
	return newMap
}

func (s *federatedGatewaySet) Insert(
	federatedGatewayList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range federatedGatewayList {
		s.Generic().Insert(obj)
	}
}

func (s *federatedGatewaySet) Has(federatedGateway ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(federatedGateway)
}

func (s *federatedGatewaySet) Equal(
	federatedGatewaySet FederatedGatewaySet,
) bool {
	if s == nil {
		return federatedGatewaySet == nil
	}
	return s.Generic().Equal(federatedGatewaySet.Generic())
}

func (s *federatedGatewaySet) Delete(FederatedGateway ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(FederatedGateway)
}

func (s *federatedGatewaySet) Union(set FederatedGatewaySet) FederatedGatewaySet {
	if s == nil {
		return set
	}
	return NewFederatedGatewaySet(append(s.List(), set.List()...)...)
}

func (s *federatedGatewaySet) Difference(set FederatedGatewaySet) FederatedGatewaySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &federatedGatewaySet{set: newSet}
}

func (s *federatedGatewaySet) Intersection(set FederatedGatewaySet) FederatedGatewaySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var federatedGatewayList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway
	for _, obj := range newSet.List() {
		federatedGatewayList = append(federatedGatewayList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway))
	}
	return NewFederatedGatewaySet(federatedGatewayList...)
}

func (s *federatedGatewaySet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find FederatedGateway %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.FederatedGateway), nil
}

func (s *federatedGatewaySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *federatedGatewaySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *federatedGatewaySet) Delta(newSet FederatedGatewaySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type RouteTableSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	// Insert a resource into the set.
	Insert(routeTable ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(routeTableSet RouteTableSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(routeTable ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(routeTable ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RouteTableSet) RouteTableSet
	// Return the difference with the provided set
	Difference(set RouteTableSet) RouteTableSet
	// Return the intersection with the provided set
	Intersection(set RouteTableSet) RouteTableSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RouteTableSet
	Delta(newSet RouteTableSet) sksets.ResourceDelta
}

func makeGenericRouteTableSet(routeTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range routeTableList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type routeTableSet struct {
	set sksets.ResourceSet
}

func NewRouteTableSet(routeTableList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) RouteTableSet {
	return &routeTableSet{set: makeGenericRouteTableSet(routeTableList)}
}

func NewRouteTableSetFromList(routeTableList *networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableList) RouteTableSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, 0, len(routeTableList.Items))
	for idx := range routeTableList.Items {
		list = append(list, &routeTableList.Items[idx])
	}
	return &routeTableSet{set: makeGenericRouteTableSet(list)}
}

func (s *routeTableSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *routeTableSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
		})
	}

	var routeTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	for _, obj := range s.Generic().List(genericFilters...) {
		routeTableList = append(routeTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
	}
	return routeTableList
}

func (s *routeTableSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable)
	}
	return newMap
}

func (s *routeTableSet) Insert(
	routeTableList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range routeTableList {
		s.Generic().Insert(obj)
	}
}

func (s *routeTableSet) Has(routeTable ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(routeTable)
}

func (s *routeTableSet) Equal(
	routeTableSet RouteTableSet,
) bool {
	if s == nil {
		return routeTableSet == nil
	}
	return s.Generic().Equal(routeTableSet.Generic())
}

func (s *routeTableSet) Delete(RouteTable ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(RouteTable)
}

func (s *routeTableSet) Union(set RouteTableSet) RouteTableSet {
	if s == nil {
		return set
	}
	return NewRouteTableSet(append(s.List(), set.List()...)...)
}

func (s *routeTableSet) Difference(set RouteTableSet) RouteTableSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &routeTableSet{set: newSet}
}

func (s *routeTableSet) Intersection(set RouteTableSet) RouteTableSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var routeTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	for _, obj := range newSet.List() {
		routeTableList = append(routeTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
	}
	return NewRouteTableSet(routeTableList...)
}

func (s *routeTableSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find RouteTable %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable), nil
}

func (s *routeTableSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *routeTableSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *routeTableSet) Delta(newSet RouteTableSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type DelegatedRouteTableSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable
	// Insert a resource into the set.
	Insert(delegatedRouteTable ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(delegatedRouteTableSet DelegatedRouteTableSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(delegatedRouteTable ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(delegatedRouteTable ezkube.ResourceId)
	// Return the union with the provided set
	Union(set DelegatedRouteTableSet) DelegatedRouteTableSet
	// Return the difference with the provided set
	Difference(set DelegatedRouteTableSet) DelegatedRouteTableSet
	// Return the intersection with the provided set
	Intersection(set DelegatedRouteTableSet) DelegatedRouteTableSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another DelegatedRouteTableSet
	Delta(newSet DelegatedRouteTableSet) sksets.ResourceDelta
}

func makeGenericDelegatedRouteTableSet(delegatedRouteTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range delegatedRouteTableList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type delegatedRouteTableSet struct {
	set sksets.ResourceSet
}

func NewDelegatedRouteTableSet(delegatedRouteTableList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable) DelegatedRouteTableSet {
	return &delegatedRouteTableSet{set: makeGenericDelegatedRouteTableSet(delegatedRouteTableList)}
}

func NewDelegatedRouteTableSetFromList(delegatedRouteTableList *networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTableList) DelegatedRouteTableSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable, 0, len(delegatedRouteTableList.Items))
	for idx := range delegatedRouteTableList.Items {
		list = append(list, &delegatedRouteTableList.Items[idx])
	}
	return &delegatedRouteTableSet{set: makeGenericDelegatedRouteTableSet(list)}
}

func (s *delegatedRouteTableSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *delegatedRouteTableSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable))
		})
	}

	var delegatedRouteTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable
	for _, obj := range s.Generic().List(genericFilters...) {
		delegatedRouteTableList = append(delegatedRouteTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable))
	}
	return delegatedRouteTableList
}

func (s *delegatedRouteTableSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable)
	}
	return newMap
}

func (s *delegatedRouteTableSet) Insert(
	delegatedRouteTableList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range delegatedRouteTableList {
		s.Generic().Insert(obj)
	}
}

func (s *delegatedRouteTableSet) Has(delegatedRouteTable ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(delegatedRouteTable)
}

func (s *delegatedRouteTableSet) Equal(
	delegatedRouteTableSet DelegatedRouteTableSet,
) bool {
	if s == nil {
		return delegatedRouteTableSet == nil
	}
	return s.Generic().Equal(delegatedRouteTableSet.Generic())
}

func (s *delegatedRouteTableSet) Delete(DelegatedRouteTable ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(DelegatedRouteTable)
}

func (s *delegatedRouteTableSet) Union(set DelegatedRouteTableSet) DelegatedRouteTableSet {
	if s == nil {
		return set
	}
	return NewDelegatedRouteTableSet(append(s.List(), set.List()...)...)
}

func (s *delegatedRouteTableSet) Difference(set DelegatedRouteTableSet) DelegatedRouteTableSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &delegatedRouteTableSet{set: newSet}
}

func (s *delegatedRouteTableSet) Intersection(set DelegatedRouteTableSet) DelegatedRouteTableSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var delegatedRouteTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable
	for _, obj := range newSet.List() {
		delegatedRouteTableList = append(delegatedRouteTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable))
	}
	return NewDelegatedRouteTableSet(delegatedRouteTableList...)
}

func (s *delegatedRouteTableSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find DelegatedRouteTable %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.DelegatedRouteTable), nil
}

func (s *delegatedRouteTableSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *delegatedRouteTableSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *delegatedRouteTableSet) Delta(newSet DelegatedRouteTableSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}
