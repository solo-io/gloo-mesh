// Code generated by MockGen. DO NOT EDIT.
// Source: ./snapshot.go

// Package mock_output is a generated GoMock package.
package mock_output

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	output "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/output"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	output0 "github.com/solo-io/skv2/contrib/pkg/output"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
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

// TrafficTargets mocks base method
func (m *MockSnapshot) TrafficTargets() []output.LabeledTrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficTargets")
	ret0, _ := ret[0].([]output.LabeledTrafficTargetSet)
	return ret0
}

// TrafficTargets indicates an expected call of TrafficTargets
func (mr *MockSnapshotMockRecorder) TrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficTargets", reflect.TypeOf((*MockSnapshot)(nil).TrafficTargets))
}

// MeshWorkloads mocks base method
func (m *MockSnapshot) MeshWorkloads() []output.LabeledMeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeshWorkloads")
	ret0, _ := ret[0].([]output.LabeledMeshWorkloadSet)
	return ret0
}

// MeshWorkloads indicates an expected call of MeshWorkloads
func (mr *MockSnapshotMockRecorder) MeshWorkloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeshWorkloads", reflect.TypeOf((*MockSnapshot)(nil).MeshWorkloads))
}

// Meshes mocks base method
func (m *MockSnapshot) Meshes() []output.LabeledMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meshes")
	ret0, _ := ret[0].([]output.LabeledMeshSet)
	return ret0
}

// Meshes indicates an expected call of Meshes
func (mr *MockSnapshotMockRecorder) Meshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meshes", reflect.TypeOf((*MockSnapshot)(nil).Meshes))
}

// ApplyLocalCluster mocks base method
func (m *MockSnapshot) ApplyLocalCluster(ctx context.Context, clusterClient client.Client, errHandler output0.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyLocalCluster", ctx, clusterClient, errHandler)
}

// ApplyLocalCluster indicates an expected call of ApplyLocalCluster
func (mr *MockSnapshotMockRecorder) ApplyLocalCluster(ctx, clusterClient, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLocalCluster", reflect.TypeOf((*MockSnapshot)(nil).ApplyLocalCluster), ctx, clusterClient, errHandler)
}

// ApplyMultiCluster mocks base method
func (m *MockSnapshot) ApplyMultiCluster(ctx context.Context, multiClusterClient multicluster.Client, errHandler output0.ErrorHandler) {
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

// MockLabeledTrafficTargetSet is a mock of LabeledTrafficTargetSet interface
type MockLabeledTrafficTargetSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledTrafficTargetSetMockRecorder
}

// MockLabeledTrafficTargetSetMockRecorder is the mock recorder for MockLabeledTrafficTargetSet
type MockLabeledTrafficTargetSetMockRecorder struct {
	mock *MockLabeledTrafficTargetSet
}

// NewMockLabeledTrafficTargetSet creates a new mock instance
func NewMockLabeledTrafficTargetSet(ctrl *gomock.Controller) *MockLabeledTrafficTargetSet {
	mock := &MockLabeledTrafficTargetSet{ctrl: ctrl}
	mock.recorder = &MockLabeledTrafficTargetSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledTrafficTargetSet) EXPECT() *MockLabeledTrafficTargetSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledTrafficTargetSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledTrafficTargetSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledTrafficTargetSet) Set() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledTrafficTargetSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledTrafficTargetSet) Generic() output0.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output0.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledTrafficTargetSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledTrafficTargetSet)(nil).Generic))
}

// MockLabeledMeshWorkloadSet is a mock of LabeledMeshWorkloadSet interface
type MockLabeledMeshWorkloadSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledMeshWorkloadSetMockRecorder
}

// MockLabeledMeshWorkloadSetMockRecorder is the mock recorder for MockLabeledMeshWorkloadSet
type MockLabeledMeshWorkloadSetMockRecorder struct {
	mock *MockLabeledMeshWorkloadSet
}

// NewMockLabeledMeshWorkloadSet creates a new mock instance
func NewMockLabeledMeshWorkloadSet(ctrl *gomock.Controller) *MockLabeledMeshWorkloadSet {
	mock := &MockLabeledMeshWorkloadSet{ctrl: ctrl}
	mock.recorder = &MockLabeledMeshWorkloadSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledMeshWorkloadSet) EXPECT() *MockLabeledMeshWorkloadSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledMeshWorkloadSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledMeshWorkloadSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledMeshWorkloadSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledMeshWorkloadSet) Set() v1alpha2sets.MeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.MeshWorkloadSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledMeshWorkloadSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledMeshWorkloadSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledMeshWorkloadSet) Generic() output0.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output0.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledMeshWorkloadSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledMeshWorkloadSet)(nil).Generic))
}

// MockLabeledMeshSet is a mock of LabeledMeshSet interface
type MockLabeledMeshSet struct {
	ctrl     *gomock.Controller
	recorder *MockLabeledMeshSetMockRecorder
}

// MockLabeledMeshSetMockRecorder is the mock recorder for MockLabeledMeshSet
type MockLabeledMeshSetMockRecorder struct {
	mock *MockLabeledMeshSet
}

// NewMockLabeledMeshSet creates a new mock instance
func NewMockLabeledMeshSet(ctrl *gomock.Controller) *MockLabeledMeshSet {
	mock := &MockLabeledMeshSet{ctrl: ctrl}
	mock.recorder = &MockLabeledMeshSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLabeledMeshSet) EXPECT() *MockLabeledMeshSetMockRecorder {
	return m.recorder
}

// Labels mocks base method
func (m *MockLabeledMeshSet) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockLabeledMeshSetMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockLabeledMeshSet)(nil).Labels))
}

// Set mocks base method
func (m *MockLabeledMeshSet) Set() v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set")
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLabeledMeshSetMockRecorder) Set() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLabeledMeshSet)(nil).Set))
}

// Generic mocks base method
func (m *MockLabeledMeshSet) Generic() output0.ResourceList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generic")
	ret0, _ := ret[0].(output0.ResourceList)
	return ret0
}

// Generic indicates an expected call of Generic
func (mr *MockLabeledMeshSetMockRecorder) Generic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generic", reflect.TypeOf((*MockLabeledMeshSet)(nil).Generic))
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

// AddTrafficTargets mocks base method
func (m *MockBuilder) AddTrafficTargets(trafficTargets ...*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range trafficTargets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddTrafficTargets", varargs...)
}

// AddTrafficTargets indicates an expected call of AddTrafficTargets
func (mr *MockBuilderMockRecorder) AddTrafficTargets(trafficTargets ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrafficTargets", reflect.TypeOf((*MockBuilder)(nil).AddTrafficTargets), trafficTargets...)
}

// GetTrafficTargets mocks base method
func (m *MockBuilder) GetTrafficTargets() v1alpha2sets.TrafficTargetSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrafficTargets")
	ret0, _ := ret[0].(v1alpha2sets.TrafficTargetSet)
	return ret0
}

// GetTrafficTargets indicates an expected call of GetTrafficTargets
func (mr *MockBuilderMockRecorder) GetTrafficTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrafficTargets", reflect.TypeOf((*MockBuilder)(nil).GetTrafficTargets))
}

// AddMeshWorkloads mocks base method
func (m *MockBuilder) AddMeshWorkloads(meshWorkloads ...*v1alpha2.MeshWorkload) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range meshWorkloads {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMeshWorkloads", varargs...)
}

// AddMeshWorkloads indicates an expected call of AddMeshWorkloads
func (mr *MockBuilderMockRecorder) AddMeshWorkloads(meshWorkloads ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMeshWorkloads", reflect.TypeOf((*MockBuilder)(nil).AddMeshWorkloads), meshWorkloads...)
}

// GetMeshWorkloads mocks base method
func (m *MockBuilder) GetMeshWorkloads() v1alpha2sets.MeshWorkloadSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshWorkloads")
	ret0, _ := ret[0].(v1alpha2sets.MeshWorkloadSet)
	return ret0
}

// GetMeshWorkloads indicates an expected call of GetMeshWorkloads
func (mr *MockBuilderMockRecorder) GetMeshWorkloads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshWorkloads", reflect.TypeOf((*MockBuilder)(nil).GetMeshWorkloads))
}

// AddMeshes mocks base method
func (m *MockBuilder) AddMeshes(meshes ...*v1alpha2.Mesh) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range meshes {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMeshes", varargs...)
}

// AddMeshes indicates an expected call of AddMeshes
func (mr *MockBuilderMockRecorder) AddMeshes(meshes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMeshes", reflect.TypeOf((*MockBuilder)(nil).AddMeshes), meshes...)
}

// GetMeshes mocks base method
func (m *MockBuilder) GetMeshes() v1alpha2sets.MeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshes")
	ret0, _ := ret[0].(v1alpha2sets.MeshSet)
	return ret0
}

// GetMeshes indicates an expected call of GetMeshes
func (mr *MockBuilderMockRecorder) GetMeshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshes", reflect.TypeOf((*MockBuilder)(nil).GetMeshes))
}

// BuildLabelPartitionedSnapshot mocks base method
func (m *MockBuilder) BuildLabelPartitionedSnapshot(labelKey string) (output.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildLabelPartitionedSnapshot", labelKey)
	ret0, _ := ret[0].(output.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildLabelPartitionedSnapshot indicates an expected call of BuildLabelPartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildLabelPartitionedSnapshot(labelKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildLabelPartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildLabelPartitionedSnapshot), labelKey)
}

// BuildSinglePartitionedSnapshot mocks base method
func (m *MockBuilder) BuildSinglePartitionedSnapshot(snapshotLabels map[string]string) (output.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSinglePartitionedSnapshot", snapshotLabels)
	ret0, _ := ret[0].(output.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSinglePartitionedSnapshot indicates an expected call of BuildSinglePartitionedSnapshot
func (mr *MockBuilderMockRecorder) BuildSinglePartitionedSnapshot(snapshotLabels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSinglePartitionedSnapshot", reflect.TypeOf((*MockBuilder)(nil).BuildSinglePartitionedSnapshot), snapshotLabels)
}
