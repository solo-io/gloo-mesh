// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_appmesh is a generated GoMock package.
package mock_appmesh

import (
	context "context"
	v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	gomock "github.com/golang/mock/gomock"
	v1beta2sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"
	appmesh "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/appmesh"
	output "github.com/solo-io/skv2/contrib/pkg/output"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSnapshot is a mock of Snapshot interface
type MockSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMockRecorder
}

// MockSnapshotMockRecorder is the mock recorder for MockSnapshot
type MockSnapshotMockRecorder struct {
	mock *MockSnapshot
}

// NewMockSnapshot creates a new mock instance
func NewMockSnapshot(ctrl *gomock.Controller) *MockSnapshot {
	mock := &MockSnapshot{ctrl: ctrl}
	mock.recorder = &MockSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshot) EXPECT() *MockSnapshotMockRecorder {
	return m.recorder
}

// VirtualServices mocks base method
func (m *MockSnapshot) VirtualServices() []appmesh.LabeledVirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualServices")
	ret0, _ := ret[0].([]appmesh.LabeledVirtualServiceSet)
	return ret0
}

// VirtualServices indicates an expected call of VirtualServices
func (mr *MockSnapshotMockRecorder) VirtualServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualServices", reflect.TypeOf((*MockSnapshot)(nil).VirtualServices))
}

// VirtualNodes mocks base method
func (m *MockSnapshot) VirtualNodes() []appmesh.LabeledVirtualNodeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualNodes")
	ret0, _ := ret[0].([]appmesh.LabeledVirtualNodeSet)
	return ret0
}

// VirtualNodes indicates an expected call of VirtualNodes
func (mr *MockSnapshotMockRecorder) VirtualNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualNodes", reflect.TypeOf((*MockSnapshot)(nil).VirtualNodes))
}

// VirtualRouters mocks base method
func (m *MockSnapshot) VirtualRouters() []appmesh.LabeledVirtualRouterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualRouters")
	ret0, _ := ret[0].([]appmesh.LabeledVirtualRouterSet)
	return ret0
}

// VirtualRouters indicates an expected call of VirtualRouters
func (mr *MockSnapshotMockRecorder) VirtualRouters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualRouters", reflect.TypeOf((*MockSnapshot)(nil).VirtualRouters))
}

// ApplyLocalCluster mocks base method
func (m *MockSnapshot) ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyLocalCluster", ctx, clusterClient, errHandler)
}

// ApplyLocalCluster indicates an expected call of ApplyLocalCluster
func (mr *MockSnapshotMockRecorder) ApplyLocalCluster(ctx, clusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLocalCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyLocalCluster), ctx, clusterClient, errHandler)
}

// ApplyMultiCluster mocks base method
func (m *MockSnapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyMultiCluster", ctx, multiClusterClient, errHandler)
}

// ApplyMultiCluster indicates an expected call of ApplyMultiCluster
func (mr *MockSnapshotMockRecorder) ApplyMultiCluster(ctx, multiClusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyMultiCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyMultiCluster), ctx, multiClusterClient, errHandler)
}

// MarshalJSON mocks base method
func (m *MockSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockSnapshot)(nil).MarshalJSON))
}

// MockLabeledVirtualServiceSet is a mock of LabeledVirtualServiceSet interface
type MockLabeledVirtualServiceSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledVirtualServiceSetMockRecorder
}

// MockLabeledVirtualServiceSetMockRecorder is the mock recorder for MockLabeledVirtualServiceSet
type MockLabeledVirtualServiceSetMockRecorder struct {
	mock *MockLabeledVirtualServiceSet
}

// NewMockLabeledVirtualServiceSet creates a new mock instance
func NewMockLabeledVirtualServiceSet(ctrl *gomock.Controller) *MockLabeledVirtualServiceSet {
	mock := &MockLabeledVirtualServiceSet{ctrl: ctrl}
	mock.recorder = &MockLabeledVirtualServiceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledVirtualServiceSet) EXPECT() *MockLabeledVirtualServiceSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledVirtualServiceSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledVirtualServiceSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledVirtualServiceSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledVirtualServiceSet) Set() v1beta2sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1beta2sets.VirtualServiceSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledVirtualServiceSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledVirtualServiceSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledVirtualServiceSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledVirtualServiceSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledVirtualServiceSet)(nil).Generic))
}

// MockLabeledVirtualNodeSet is a mock of LabeledVirtualNodeSet interface
type MockLabeledVirtualNodeSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledVirtualNodeSetMockRecorder
}

// MockLabeledVirtualNodeSetMockRecorder is the mock recorder for MockLabeledVirtualNodeSet
type MockLabeledVirtualNodeSetMockRecorder struct {
	mock *MockLabeledVirtualNodeSet
}

// NewMockLabeledVirtualNodeSet creates a new mock instance
func NewMockLabeledVirtualNodeSet(ctrl *gomock.Controller) *MockLabeledVirtualNodeSet {
	mock := &MockLabeledVirtualNodeSet{ctrl: ctrl}
	mock.recorder = &MockLabeledVirtualNodeSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledVirtualNodeSet) EXPECT() *MockLabeledVirtualNodeSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledVirtualNodeSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledVirtualNodeSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledVirtualNodeSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledVirtualNodeSet) Set() v1beta2sets.VirtualNodeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1beta2sets.VirtualNodeSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledVirtualNodeSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledVirtualNodeSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledVirtualNodeSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledVirtualNodeSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledVirtualNodeSet)(nil).Generic))
}

// MockLabeledVirtualRouterSet is a mock of LabeledVirtualRouterSet interface
type MockLabeledVirtualRouterSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledVirtualRouterSetMockRecorder
}

// MockLabeledVirtualRouterSetMockRecorder is the mock recorder for MockLabeledVirtualRouterSet
type MockLabeledVirtualRouterSetMockRecorder struct {
	mock *MockLabeledVirtualRouterSet
}

// NewMockLabeledVirtualRouterSet creates a new mock instance
func NewMockLabeledVirtualRouterSet(ctrl *gomock.Controller) *MockLabeledVirtualRouterSet {
	mock := &MockLabeledVirtualRouterSet{ctrl: ctrl}
	mock.recorder = &MockLabeledVirtualRouterSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledVirtualRouterSet) EXPECT() *MockLabeledVirtualRouterSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledVirtualRouterSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledVirtualRouterSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledVirtualRouterSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledVirtualRouterSet) Set() v1beta2sets.VirtualRouterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1beta2sets.VirtualRouterSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledVirtualRouterSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledVirtualRouterSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledVirtualRouterSet) Generic() output.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledVirtualRouterSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledVirtualRouterSet)(nil).Generic))
}

// MockBuilder is a mock of Builder interface
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// AddVirtualServices mocks base method
func (m *MockBuilder) AddVirtualServices(virtualServices ...*v1beta2.VirtualService) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualServices {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddVirtualServices", varargs...)
}

// AddVirtualServices indicates an expected call of AddVirtualServices
func (mr *MockBuilderMockRecorder) AddVirtualServices(virtualServices ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVirtualServices", reflect.TypeOf((*MockBuilder)(nil).AddVirtualServices), virtualServices...)
}

// GetVirtualServices mocks base method
func (m *MockBuilder) GetVirtualServices() v1beta2sets.VirtualServiceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualServices")
	ret0, _ := ret[0].(v1beta2sets.VirtualServiceSet)
	return ret0
}

// GetVirtualServices indicates an expected call of GetVirtualServices
func (mr *MockBuilderMockRecorder) GetVirtualServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualServices", reflect.TypeOf((*MockBuilder)(nil).GetVirtualServices))
}

// AddVirtualNodes mocks base method
func (m *MockBuilder) AddVirtualNodes(virtualNodes ...*v1beta2.VirtualNode) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualNodes {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddVirtualNodes", varargs...)
}

// AddVirtualNodes indicates an expected call of AddVirtualNodes
func (mr *MockBuilderMockRecorder) AddVirtualNodes(virtualNodes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVirtualNodes", reflect.TypeOf((*MockBuilder)(nil).AddVirtualNodes), virtualNodes...)
}

// GetVirtualNodes mocks base method
func (m *MockBuilder) GetVirtualNodes() v1beta2sets.VirtualNodeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualNodes")
	ret0, _ := ret[0].(v1beta2sets.VirtualNodeSet)
	return ret0
}

// GetVirtualNodes indicates an expected call of GetVirtualNodes
func (mr *MockBuilderMockRecorder) GetVirtualNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualNodes", reflect.TypeOf((*MockBuilder)(nil).GetVirtualNodes))
}

// AddVirtualRouters mocks base method
func (m *MockBuilder) AddVirtualRouters(virtualRouters ...*v1beta2.VirtualRouter) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range virtualRouters {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddVirtualRouters", varargs...)
}

// AddVirtualRouters indicates an expected call of AddVirtualRouters
func (mr *MockBuilderMockRecorder) AddVirtualRouters(virtualRouters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVirtualRouters", reflect.TypeOf((*MockBuilder)(nil).AddVirtualRouters), virtualRouters...)
}

// GetVirtualRouters mocks base method
func (m *MockBuilder) GetVirtualRouters() v1beta2sets.VirtualRouterSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualRouters")
	ret0, _ := ret[0].(v1beta2sets.VirtualRouterSet)
	return ret0
}

// GetVirtualRouters indicates an expected call of GetVirtualRouters
func (mr *MockBuilderMockRecorder) GetVirtualRouters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualRouters", reflect.TypeOf((*MockBuilder)(nil).GetVirtualRouters))
}

// BuildLabelPartitionedSnapshot mocks base method
func (m *MockBuilder) BuildLabelPartitionedSnapshot(labelKey string) (appmesh.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildLabelPartitionedSnapshot", labelKey)
	ret0, _ := ret[0].(appmesh.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildLabelPartitionedSnapshot indicates an expected call of BuildLabelPartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildLabelPartitionedSnapshot(labelKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildLabelPartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildLabelPartitionedSnapshot), labelKey)
}

// BuildSinglePartitionedSnapshot mocks base method
func (m *MockBuilder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (appmesh.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSinglePartitionedSnapshot", snapshotLabels)
	ret0, _ := ret[0].(appmesh.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSinglePartitionedSnapshot indicates an expected call of BuildSinglePartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildSinglePartitionedSnapshot(snapshotLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSinglePartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSinglePartitionedSnapshot), snapshotLabels)
}

// AddCluster mocks base method
func (m *MockBuilder) AddCluster(cluster string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddCluster", cluster)
}

// AddCluster indicates an expected call of AddCluster
func (mr *MockBuilderMockRecorder) AddCluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockBuilder)(nil).AddCluster), cluster)
}

// Clusters mocks base method
func (m *MockBuilder) Clusters() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clusters")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Clusters indicates an expected call of Clusters
func (mr *MockBuilderMockRecorder) Clusters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clusters", reflect.TypeOf((*MockBuilder)(nil).Clusters))
}

// Merge mocks base method
func (m *MockBuilder) Merge(other appmesh.Builder) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Merge", other)
}

// Merge indicates an expected call of Merge
func (mr *MockBuilderMockRecorder) Merge(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockBuilder)(nil).Merge), other)
}

// Clone mocks base method
func (m *MockBuilder) Clone() appmesh.Builder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(appmesh.Builder)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockBuilderMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockBuilder)(nil).Clone))
}
