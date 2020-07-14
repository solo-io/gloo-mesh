// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1alpha1sets is a generated GoMock package.
package mock_v1alpha1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockKubernetesClusterSet is a mock of KubernetesClusterSet interface.
type MockKubernetesClusterSet struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterSetMockRecorder
}

// MockKubernetesClusterSetMockRecorder is the mock recorder for MockKubernetesClusterSet.
type MockKubernetesClusterSetMockRecorder struct {
	mock *MockKubernetesClusterSet
}

// NewMockKubernetesClusterSet creates a new mock instance.
func NewMockKubernetesClusterSet(ctrl *gomock.Controller) *MockKubernetesClusterSet {
	mock := &MockKubernetesClusterSet{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesClusterSet) EXPECT() *MockKubernetesClusterSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockKubernetesClusterSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockKubernetesClusterSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Keys))
}

// List mocks base method.
func (m *MockKubernetesClusterSet) List() []*v1alpha1.KubernetesCluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.KubernetesCluster)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockKubernetesClusterSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockKubernetesClusterSet)(nil).List))
}

// Map mocks base method.
func (m *MockKubernetesClusterSet) Map() map[string]*v1alpha1.KubernetesCluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.KubernetesCluster)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockKubernetesClusterSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockKubernetesClusterSet) Insert(kubernetesCluster ...*v1alpha1.KubernetesCluster) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range kubernetesCluster {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockKubernetesClusterSetMockRecorder) Insert(kubernetesCluster ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Insert), kubernetesCluster...)
}

// Equal mocks base method.
func (m *MockKubernetesClusterSet) Equal(kubernetesClusterSet v1alpha1sets.KubernetesClusterSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", kubernetesClusterSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockKubernetesClusterSetMockRecorder) Equal(kubernetesClusterSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Equal), kubernetesClusterSet)
}

// Has mocks base method.
func (m *MockKubernetesClusterSet) Has(kubernetesCluster *v1alpha1.KubernetesCluster) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", kubernetesCluster)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockKubernetesClusterSetMockRecorder) Has(kubernetesCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Has), kubernetesCluster)
}

// Delete mocks base method.
func (m *MockKubernetesClusterSet) Delete(kubernetesCluster *v1alpha1.KubernetesCluster) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", kubernetesCluster)
}

// Delete indicates an expected call of Delete.
func (mr *MockKubernetesClusterSetMockRecorder) Delete(kubernetesCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Delete), kubernetesCluster)
}

// Union mocks base method.
func (m *MockKubernetesClusterSet) Union(set v1alpha1sets.KubernetesClusterSet) v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockKubernetesClusterSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockKubernetesClusterSet) Difference(set v1alpha1sets.KubernetesClusterSet) v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockKubernetesClusterSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockKubernetesClusterSet) Intersection(set v1alpha1sets.KubernetesClusterSet) v1alpha1sets.KubernetesClusterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.KubernetesClusterSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockKubernetesClusterSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockKubernetesClusterSet) Find(id ezkube.ResourceId) (*v1alpha1.KubernetesCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.KubernetesCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockKubernetesClusterSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockKubernetesClusterSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockKubernetesClusterSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockKubernetesClusterSet)(nil).Length))
}

// MockMeshServiceSet is a mock of MeshServiceSet interface.
type MockMeshServiceSet struct {
	ctrl     *gomock.Controller
	recorder *MockMeshServiceSetMockRecorder
}

// MockMeshServiceSetMockRecorder is the mock recorder for MockMeshServiceSet.
type MockMeshServiceSetMockRecorder struct {
	mock *MockMeshServiceSet
}

// NewMockMeshServiceSet creates a new mock instance.
func NewMockMeshServiceSet(ctrl *gomock.Controller) *MockMeshServiceSet {
	mock := &MockMeshServiceSet{ctrl: ctrl}
	mock.recorder = &MockMeshServiceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshServiceSet) EXPECT() *MockMeshServiceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockMeshServiceSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockMeshServiceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockMeshServiceSet)(nil).Keys))
}

// List mocks base method.
func (m *MockMeshServiceSet) List() []*v1alpha1.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.MeshService)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockMeshServiceSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMeshServiceSet)(nil).List))
}

// Map mocks base method.
func (m *MockMeshServiceSet) Map() map[string]*v1alpha1.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.MeshService)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockMeshServiceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockMeshServiceSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockMeshServiceSet) Insert(meshService ...*v1alpha1.MeshService) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range meshService {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockMeshServiceSetMockRecorder) Insert(meshService ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMeshServiceSet)(nil).Insert), meshService...)
}

// Equal mocks base method.
func (m *MockMeshServiceSet) Equal(meshServiceSet v1alpha1sets.MeshServiceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", meshServiceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockMeshServiceSetMockRecorder) Equal(meshServiceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMeshServiceSet)(nil).Equal), meshServiceSet)
}

// Has mocks base method.
func (m *MockMeshServiceSet) Has(meshService *v1alpha1.MeshService) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", meshService)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockMeshServiceSetMockRecorder) Has(meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMeshServiceSet)(nil).Has), meshService)
}

// Delete mocks base method.
func (m *MockMeshServiceSet) Delete(meshService *v1alpha1.MeshService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", meshService)
}

// Delete indicates an expected call of Delete.
func (mr *MockMeshServiceSetMockRecorder) Delete(meshService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshServiceSet)(nil).Delete), meshService)
}

// Union mocks base method.
func (m *MockMeshServiceSet) Union(set v1alpha1sets.MeshServiceSet) v1alpha1sets.MeshServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshServiceSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockMeshServiceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockMeshServiceSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockMeshServiceSet) Difference(set v1alpha1sets.MeshServiceSet) v1alpha1sets.MeshServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshServiceSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockMeshServiceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockMeshServiceSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockMeshServiceSet) Intersection(set v1alpha1sets.MeshServiceSet) v1alpha1sets.MeshServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshServiceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockMeshServiceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockMeshServiceSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockMeshServiceSet) Find(id ezkube.ResourceId) (*v1alpha1.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMeshServiceSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMeshServiceSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockMeshServiceSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockMeshServiceSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMeshServiceSet)(nil).Length))
}

// MockMeshWorkloadSet is a mock of MeshWorkloadSet interface.
type MockMeshWorkloadSet struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadSetMockRecorder
}

// MockMeshWorkloadSetMockRecorder is the mock recorder for MockMeshWorkloadSet.
type MockMeshWorkloadSetMockRecorder struct {
	mock *MockMeshWorkloadSet
}

// NewMockMeshWorkloadSet creates a new mock instance.
func NewMockMeshWorkloadSet(ctrl *gomock.Controller) *MockMeshWorkloadSet {
	mock := &MockMeshWorkloadSet{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshWorkloadSet) EXPECT() *MockMeshWorkloadSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockMeshWorkloadSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockMeshWorkloadSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Keys))
}

// List mocks base method.
func (m *MockMeshWorkloadSet) List() []*v1alpha1.MeshWorkload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.MeshWorkload)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockMeshWorkloadSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMeshWorkloadSet)(nil).List))
}

// Map mocks base method.
func (m *MockMeshWorkloadSet) Map() map[string]*v1alpha1.MeshWorkload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.MeshWorkload)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockMeshWorkloadSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockMeshWorkloadSet) Insert(meshWorkload ...*v1alpha1.MeshWorkload) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range meshWorkload {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockMeshWorkloadSetMockRecorder) Insert(meshWorkload ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Insert), meshWorkload...)
}

// Equal mocks base method.
func (m *MockMeshWorkloadSet) Equal(meshWorkloadSet v1alpha1sets.MeshWorkloadSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", meshWorkloadSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockMeshWorkloadSetMockRecorder) Equal(meshWorkloadSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Equal), meshWorkloadSet)
}

// Has mocks base method.
func (m *MockMeshWorkloadSet) Has(meshWorkload *v1alpha1.MeshWorkload) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", meshWorkload)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockMeshWorkloadSetMockRecorder) Has(meshWorkload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Has), meshWorkload)
}

// Delete mocks base method.
func (m *MockMeshWorkloadSet) Delete(meshWorkload *v1alpha1.MeshWorkload) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", meshWorkload)
}

// Delete indicates an expected call of Delete.
func (mr *MockMeshWorkloadSetMockRecorder) Delete(meshWorkload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Delete), meshWorkload)
}

// Union mocks base method.
func (m *MockMeshWorkloadSet) Union(set v1alpha1sets.MeshWorkloadSet) v1alpha1sets.MeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshWorkloadSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockMeshWorkloadSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockMeshWorkloadSet) Difference(set v1alpha1sets.MeshWorkloadSet) v1alpha1sets.MeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshWorkloadSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockMeshWorkloadSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockMeshWorkloadSet) Intersection(set v1alpha1sets.MeshWorkloadSet) v1alpha1sets.MeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshWorkloadSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockMeshWorkloadSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockMeshWorkloadSet) Find(id ezkube.ResourceId) (*v1alpha1.MeshWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.MeshWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMeshWorkloadSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockMeshWorkloadSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockMeshWorkloadSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMeshWorkloadSet)(nil).Length))
}

// MockMeshSet is a mock of MeshSet interface.
type MockMeshSet struct {
	ctrl     *gomock.Controller
	recorder *MockMeshSetMockRecorder
}

// MockMeshSetMockRecorder is the mock recorder for MockMeshSet.
type MockMeshSetMockRecorder struct {
	mock *MockMeshSet
}

// NewMockMeshSet creates a new mock instance.
func NewMockMeshSet(ctrl *gomock.Controller) *MockMeshSet {
	mock := &MockMeshSet{ctrl: ctrl}
	mock.recorder = &MockMeshSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshSet) EXPECT() *MockMeshSetMockRecorder {
	return m.recorder
}

// Keys mocks base method.
func (m *MockMeshSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockMeshSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockMeshSet)(nil).Keys))
}

// List mocks base method.
func (m *MockMeshSet) List() []*v1alpha1.Mesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1alpha1.Mesh)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockMeshSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMeshSet)(nil).List))
}

// Map mocks base method.
func (m *MockMeshSet) Map() map[string]*v1alpha1.Mesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1alpha1.Mesh)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockMeshSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockMeshSet)(nil).Map))
}

// Insert mocks base method.
func (m *MockMeshSet) Insert(mesh ...*v1alpha1.Mesh) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range mesh {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert.
func (mr *MockMeshSetMockRecorder) Insert(mesh ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMeshSet)(nil).Insert), mesh...)
}

// Equal mocks base method.
func (m *MockMeshSet) Equal(meshSet v1alpha1sets.MeshSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", meshSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockMeshSetMockRecorder) Equal(meshSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMeshSet)(nil).Equal), meshSet)
}

// Has mocks base method.
func (m *MockMeshSet) Has(mesh *v1alpha1.Mesh) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", mesh)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockMeshSetMockRecorder) Has(mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMeshSet)(nil).Has), mesh)
}

// Delete mocks base method.
func (m *MockMeshSet) Delete(mesh *v1alpha1.Mesh) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", mesh)
}

// Delete indicates an expected call of Delete.
func (mr *MockMeshSetMockRecorder) Delete(mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshSet)(nil).Delete), mesh)
}

// Union mocks base method.
func (m *MockMeshSet) Union(set v1alpha1sets.MeshSet) v1alpha1sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshSet)
	return ret0
}

// Union indicates an expected call of Union.
func (mr *MockMeshSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockMeshSet)(nil).Union), set)
}

// Difference mocks base method.
func (m *MockMeshSet) Difference(set v1alpha1sets.MeshSet) v1alpha1sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshSet)
	return ret0
}

// Difference indicates an expected call of Difference.
func (mr *MockMeshSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockMeshSet)(nil).Difference), set)
}

// Intersection mocks base method.
func (m *MockMeshSet) Intersection(set v1alpha1sets.MeshSet) v1alpha1sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1alpha1sets.MeshSet)
	return ret0
}

// Intersection indicates an expected call of Intersection.
func (mr *MockMeshSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockMeshSet)(nil).Intersection), set)
}

// Find mocks base method.
func (m *MockMeshSet) Find(id ezkube.ResourceId) (*v1alpha1.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1alpha1.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMeshSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMeshSet)(nil).Find), id)
}

// Length mocks base method.
func (m *MockMeshSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockMeshSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMeshSet)(nil).Length))
}
