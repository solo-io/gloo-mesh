// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_strategies is a generated GoMock package.
package mock_strategies

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	strategies "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/decider/strategies"
)

// MockFederationStrategy is a mock of FederationStrategy interface.
type MockFederationStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockFederationStrategyMockRecorder
}

// MockFederationStrategyMockRecorder is the mock recorder for MockFederationStrategy.
type MockFederationStrategyMockRecorder struct {
	mock *MockFederationStrategy
}

// NewMockFederationStrategy creates a new mock instance.
func NewMockFederationStrategy(ctrl *gomock.Controller) *MockFederationStrategy {
	mock := &MockFederationStrategy{ctrl: ctrl}
	mock.recorder = &MockFederationStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederationStrategy) EXPECT() *MockFederationStrategyMockRecorder {
	return m.recorder
}

// WriteFederationToServices mocks base method.
func (m *MockFederationStrategy) WriteFederationToServices(ctx context.Context, vm *v1alpha1.VirtualMesh, meshNameToMetadata strategies.MeshNameToMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFederationToServices", ctx, vm, meshNameToMetadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFederationToServices indicates an expected call of WriteFederationToServices.
func (mr *MockFederationStrategyMockRecorder) WriteFederationToServices(ctx, vm, meshNameToMetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFederationToServices", reflect.TypeOf((*MockFederationStrategy)(nil).WriteFederationToServices), ctx, vm, meshNameToMetadata)
}
