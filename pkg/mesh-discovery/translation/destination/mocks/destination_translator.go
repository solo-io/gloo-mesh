// Code generated by MockGen. DO NOT EDIT.
// Source: ./destination_translator.go

// Package mock_destination is a generated GoMock package.
package mock_destination

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1sets0 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
)

// MockTranslator is a mock of Translator interface
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorMockRecorder
}

// MockTranslatorMockRecorder is the mock recorder for MockTranslator
type MockTranslatorMockRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTranslator) EXPECT() *MockTranslatorMockRecorder {
	return m.recorder
}

// TranslateDestinations mocks base method
func (m *MockTranslator) TranslateDestinations(ctx context.Context, services v1sets.ServiceSet, pods v1sets.PodSet, nodes v1sets.NodeSet, workloads v1sets0.WorkloadSet, meshes v1sets0.MeshSet, endpoints v1sets.EndpointsSet) v1sets0.DestinationSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateDestinations", ctx, services, pods, nodes, workloads, meshes, endpoints)
	ret0, _ := ret[0].(v1sets0.DestinationSet)
	return ret0
}

// TranslateDestinations indicates an expected call of TranslateDestinations
func (mr *MockTranslatorMockRecorder) TranslateDestinations(ctx, services, pods, nodes, workloads, meshes, endpoints interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateDestinations", reflect.TypeOf((*MockTranslator)(nil).TranslateDestinations), ctx, services, pods, nodes, workloads, meshes, endpoints)
}
