// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/linkerd_discovery_snapshot_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	v10 "github.com/solo-io/supergloo/pkg/api/v1"
)

// MockLinkerdDiscoveryEmitter is a mock of LinkerdDiscoveryEmitter interface
type MockLinkerdDiscoveryEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockLinkerdDiscoveryEmitterMockRecorder
}

// MockLinkerdDiscoveryEmitterMockRecorder is the mock recorder for MockLinkerdDiscoveryEmitter
type MockLinkerdDiscoveryEmitterMockRecorder struct {
	mock *MockLinkerdDiscoveryEmitter
}

// NewMockLinkerdDiscoveryEmitter creates a new mock instance
func NewMockLinkerdDiscoveryEmitter(ctrl *gomock.Controller) *MockLinkerdDiscoveryEmitter {
	mock := &MockLinkerdDiscoveryEmitter{ctrl: ctrl}
	mock.recorder = &MockLinkerdDiscoveryEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLinkerdDiscoveryEmitter) EXPECT() *MockLinkerdDiscoveryEmitterMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockLinkerdDiscoveryEmitter) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockLinkerdDiscoveryEmitterMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockLinkerdDiscoveryEmitter)(nil).Register))
}

// Mesh mocks base method
func (m *MockLinkerdDiscoveryEmitter) Mesh() v10.MeshClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mesh")
	ret0, _ := ret[0].(v10.MeshClient)
	return ret0
}

// Mesh indicates an expected call of Mesh
func (mr *MockLinkerdDiscoveryEmitterMockRecorder) Mesh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mesh", reflect.TypeOf((*MockLinkerdDiscoveryEmitter)(nil).Mesh))
}

// Install mocks base method
func (m *MockLinkerdDiscoveryEmitter) Install() v10.InstallClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install")
	ret0, _ := ret[0].(v10.InstallClient)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockLinkerdDiscoveryEmitterMockRecorder) Install() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockLinkerdDiscoveryEmitter)(nil).Install))
}

// Pod mocks base method
func (m *MockLinkerdDiscoveryEmitter) Pod() kubernetes.PodClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pod")
	ret0, _ := ret[0].(kubernetes.PodClient)
	return ret0
}

// Pod indicates an expected call of Pod
func (mr *MockLinkerdDiscoveryEmitterMockRecorder) Pod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pod", reflect.TypeOf((*MockLinkerdDiscoveryEmitter)(nil).Pod))
}

// Upstream mocks base method
func (m *MockLinkerdDiscoveryEmitter) Upstream() v1.UpstreamClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upstream")
	ret0, _ := ret[0].(v1.UpstreamClient)
	return ret0
}

// Upstream indicates an expected call of Upstream
func (mr *MockLinkerdDiscoveryEmitterMockRecorder) Upstream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upstream", reflect.TypeOf((*MockLinkerdDiscoveryEmitter)(nil).Upstream))
}

// Snapshots mocks base method
func (m *MockLinkerdDiscoveryEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *v10.LinkerdDiscoverySnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", watchNamespaces, opts)
	ret0, _ := ret[0].(<-chan *v10.LinkerdDiscoverySnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockLinkerdDiscoveryEmitterMockRecorder) Snapshots(watchNamespaces, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockLinkerdDiscoveryEmitter)(nil).Snapshots), watchNamespaces, opts)
}
