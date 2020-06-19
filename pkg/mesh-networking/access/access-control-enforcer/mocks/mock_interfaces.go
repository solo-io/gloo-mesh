// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_access_policy_enforcer is a generated GoMock package.
package mock_access_policy_enforcer

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	reflect "reflect"
)

// MockAccessPolicyEnforcerLoop is a mock of AccessPolicyEnforcerLoop interface.
type MockAccessPolicyEnforcerLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyEnforcerLoopMockRecorder
}

// MockAccessPolicyEnforcerLoopMockRecorder is the mock recorder for MockAccessPolicyEnforcerLoop.
type MockAccessPolicyEnforcerLoopMockRecorder struct {
	mock *MockAccessPolicyEnforcerLoop
}

// NewMockAccessPolicyEnforcerLoop creates a new mock instance.
func NewMockAccessPolicyEnforcerLoop(ctrl *gomock.Controller) *MockAccessPolicyEnforcerLoop {
	mock := &MockAccessPolicyEnforcerLoop{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyEnforcerLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyEnforcerLoop) EXPECT() *MockAccessPolicyEnforcerLoopMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockAccessPolicyEnforcerLoop) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockAccessPolicyEnforcerLoopMockRecorder) Start(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAccessPolicyEnforcerLoop)(nil).Start), ctx)
}

// MockAccessPolicyMeshEnforcer is a mock of AccessPolicyMeshEnforcer interface.
type MockAccessPolicyMeshEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessPolicyMeshEnforcerMockRecorder
}

// MockAccessPolicyMeshEnforcerMockRecorder is the mock recorder for MockAccessPolicyMeshEnforcer.
type MockAccessPolicyMeshEnforcerMockRecorder struct {
	mock *MockAccessPolicyMeshEnforcer
}

// NewMockAccessPolicyMeshEnforcer creates a new mock instance.
func NewMockAccessPolicyMeshEnforcer(ctrl *gomock.Controller) *MockAccessPolicyMeshEnforcer {
	mock := &MockAccessPolicyMeshEnforcer{ctrl: ctrl}
	mock.recorder = &MockAccessPolicyMeshEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessPolicyMeshEnforcer) EXPECT() *MockAccessPolicyMeshEnforcerMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockAccessPolicyMeshEnforcer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAccessPolicyMeshEnforcerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAccessPolicyMeshEnforcer)(nil).Name))
}

// ReconcileAccessControl mocks base method.
func (m *MockAccessPolicyMeshEnforcer) ReconcileAccessControl(ctx context.Context, mesh *v1alpha1.Mesh, virtualMesh *v1alpha10.VirtualMesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessControl", ctx, mesh, virtualMesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileAccessControl indicates an expected call of ReconcileAccessControl.
func (mr *MockAccessPolicyMeshEnforcerMockRecorder) ReconcileAccessControl(ctx, mesh, virtualMesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessControl", reflect.TypeOf((*MockAccessPolicyMeshEnforcer)(nil).ReconcileAccessControl), ctx, mesh, virtualMesh)
}
