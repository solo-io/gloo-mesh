// Code generated by MockGen. DO NOT EDIT.
// Source: ./local_snapshot.go

// Package mock_input is a generated GoMock package.
package mock_input

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	input "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
	v1alpha2sets "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2/sets"
	v1alpha2sets0 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1alpha2/sets"
	multicluster "github.com/solo-io/skv2/pkg/multicluster"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSettingsSnapshot is a mock of SettingsSnapshot interface
type MockSettingsSnapshot struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsSnapshotMockRecorder
}

// MockSettingsSnapshotMockRecorder is the mock recorder for MockSettingsSnapshot
type MockSettingsSnapshotMockRecorder struct {
	mock *MockSettingsSnapshot
}

// NewMockSettingsSnapshot creates a new mock instance
func NewMockSettingsSnapshot(ctrl *gomock.Controller) *MockSettingsSnapshot {
	mock := &MockSettingsSnapshot{ctrl: ctrl}
	mock.recorder = &MockSettingsSnapshotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsSnapshot) EXPECT() *MockSettingsSnapshotMockRecorder {
	return m.recorder
}

// Settings mocks base method
func (m *MockSettingsSnapshot) Settings() v1alpha2sets0.SettingsSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Settings")
	ret0, _ := ret[0].(v1alpha2sets0.SettingsSet)
	return ret0
}

// Settings indicates an expected call of Settings
func (mr *MockSettingsSnapshotMockRecorder) Settings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Settings", reflect.TypeOf((*MockSettingsSnapshot)(nil).Settings))
}

// VirtualMeshes mocks base method
func (m *MockSettingsSnapshot) VirtualMeshes() v1alpha2sets.VirtualMeshSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMeshes")
	ret0, _ := ret[0].(v1alpha2sets.VirtualMeshSet)
	return ret0
}

// VirtualMeshes indicates an expected call of VirtualMeshes
func (mr *MockSettingsSnapshotMockRecorder) VirtualMeshes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMeshes", reflect.TypeOf((*MockSettingsSnapshot)(nil).VirtualMeshes))
}

// SyncStatusesMultiCluster mocks base method
func (m *MockSettingsSnapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts input.SettingsSyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusesMultiCluster", ctx, mcClient, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatusesMultiCluster indicates an expected call of SyncStatusesMultiCluster
func (mr *MockSettingsSnapshotMockRecorder) SyncStatusesMultiCluster(ctx, mcClient, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusesMultiCluster", reflect.TypeOf((*MockSettingsSnapshot)(nil).SyncStatusesMultiCluster), ctx, mcClient, opts)
}

// SyncStatuses mocks base method
func (m *MockSettingsSnapshot) SyncStatuses(ctx context.Context, c client.Client, opts input.SettingsSyncStatusOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatuses", ctx, c, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatuses indicates an expected call of SyncStatuses
func (mr *MockSettingsSnapshotMockRecorder) SyncStatuses(ctx, c, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatuses", reflect.TypeOf((*MockSettingsSnapshot)(nil).SyncStatuses), ctx, c, opts)
}

// MarshalJSON mocks base method
func (m *MockSettingsSnapshot) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockSettingsSnapshotMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockSettingsSnapshot)(nil).MarshalJSON))
}

// MockSettingsBuilder is a mock of SettingsBuilder interface
type MockSettingsBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsBuilderMockRecorder
}

// MockSettingsBuilderMockRecorder is the mock recorder for MockSettingsBuilder
type MockSettingsBuilderMockRecorder struct {
	mock *MockSettingsBuilder
}

// NewMockSettingsBuilder creates a new mock instance
func NewMockSettingsBuilder(ctrl *gomock.Controller) *MockSettingsBuilder {
	mock := &MockSettingsBuilder{ctrl: ctrl}
	mock.recorder = &MockSettingsBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettingsBuilder) EXPECT() *MockSettingsBuilderMockRecorder {
	return m.recorder
}

// BuildSnapshot mocks base method
func (m *MockSettingsBuilder) BuildSnapshot(ctx context.Context, name string, opts input.SettingsBuildOptions) (input.SettingsSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSnapshot", ctx, name, opts)
	ret0, _ := ret[0].(input.SettingsSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildSnapshot indicates an expected call of BuildSnapshot
func (mr *MockSettingsBuilderMockRecorder) BuildSnapshot(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSnapshot", reflect.TypeOf((*MockSettingsBuilder)(nil).BuildSnapshot), ctx, name, opts)
}
