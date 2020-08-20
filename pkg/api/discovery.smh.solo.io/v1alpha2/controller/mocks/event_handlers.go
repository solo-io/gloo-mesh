// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTrafficTargetEventHandler is a mock of TrafficTargetEventHandler interface
type MockTrafficTargetEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetEventHandlerMockRecorder
}

// MockTrafficTargetEventHandlerMockRecorder is the mock recorder for MockTrafficTargetEventHandler
type MockTrafficTargetEventHandlerMockRecorder struct {
	mock *MockTrafficTargetEventHandler
}

// NewMockTrafficTargetEventHandler creates a new mock instance
func NewMockTrafficTargetEventHandler(ctrl *gomock.Controller) *MockTrafficTargetEventHandler {
	mock := &MockTrafficTargetEventHandler{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrafficTargetEventHandler) EXPECT() *MockTrafficTargetEventHandlerMockRecorder {
	return m.recorder
}

// CreateTrafficTarget mocks base method
func (m *MockTrafficTargetEventHandler) CreateTrafficTarget(obj *v1alpha2.TrafficTarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrafficTarget", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrafficTarget indicates an expected call of CreateTrafficTarget
func (mr *MockTrafficTargetEventHandlerMockRecorder) CreateTrafficTarget(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrafficTarget", reflect.TypeOf((*MockTrafficTargetEventHandler)(nil).CreateTrafficTarget), obj)
}

// UpdateTrafficTarget mocks base method
func (m *MockTrafficTargetEventHandler) UpdateTrafficTarget(old, new *v1alpha2.TrafficTarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrafficTarget", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrafficTarget indicates an expected call of UpdateTrafficTarget
func (mr *MockTrafficTargetEventHandlerMockRecorder) UpdateTrafficTarget(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrafficTarget", reflect.TypeOf((*MockTrafficTargetEventHandler)(nil).UpdateTrafficTarget), old, new)
}

// DeleteTrafficTarget mocks base method
func (m *MockTrafficTargetEventHandler) DeleteTrafficTarget(obj *v1alpha2.TrafficTarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrafficTarget", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrafficTarget indicates an expected call of DeleteTrafficTarget
func (mr *MockTrafficTargetEventHandlerMockRecorder) DeleteTrafficTarget(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrafficTarget", reflect.TypeOf((*MockTrafficTargetEventHandler)(nil).DeleteTrafficTarget), obj)
}

// GenericTrafficTarget mocks base method
func (m *MockTrafficTargetEventHandler) GenericTrafficTarget(obj *v1alpha2.TrafficTarget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericTrafficTarget", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericTrafficTarget indicates an expected call of GenericTrafficTarget
func (mr *MockTrafficTargetEventHandlerMockRecorder) GenericTrafficTarget(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericTrafficTarget", reflect.TypeOf((*MockTrafficTargetEventHandler)(nil).GenericTrafficTarget), obj)
}

// MockTrafficTargetEventWatcher is a mock of TrafficTargetEventWatcher interface
type MockTrafficTargetEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficTargetEventWatcherMockRecorder
}

// MockTrafficTargetEventWatcherMockRecorder is the mock recorder for MockTrafficTargetEventWatcher
type MockTrafficTargetEventWatcherMockRecorder struct {
	mock *MockTrafficTargetEventWatcher
}

// NewMockTrafficTargetEventWatcher creates a new mock instance
func NewMockTrafficTargetEventWatcher(ctrl *gomock.Controller) *MockTrafficTargetEventWatcher {
	mock := &MockTrafficTargetEventWatcher{ctrl: ctrl}
	mock.recorder = &MockTrafficTargetEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrafficTargetEventWatcher) EXPECT() *MockTrafficTargetEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockTrafficTargetEventWatcher) AddEventHandler(ctx context.Context, h controller.TrafficTargetEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockTrafficTargetEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockTrafficTargetEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshWorkloadEventHandler is a mock of MeshWorkloadEventHandler interface
type MockMeshWorkloadEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadEventHandlerMockRecorder
}

// MockMeshWorkloadEventHandlerMockRecorder is the mock recorder for MockMeshWorkloadEventHandler
type MockMeshWorkloadEventHandlerMockRecorder struct {
	mock *MockMeshWorkloadEventHandler
}

// NewMockMeshWorkloadEventHandler creates a new mock instance
func NewMockMeshWorkloadEventHandler(ctrl *gomock.Controller) *MockMeshWorkloadEventHandler {
	mock := &MockMeshWorkloadEventHandler{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshWorkloadEventHandler) EXPECT() *MockMeshWorkloadEventHandlerMockRecorder {
	return m.recorder
}

// CreateMeshWorkload mocks base method
func (m *MockMeshWorkloadEventHandler) CreateMeshWorkload(obj *v1alpha2.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMeshWorkload indicates an expected call of CreateMeshWorkload
func (mr *MockMeshWorkloadEventHandlerMockRecorder) CreateMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).CreateMeshWorkload), obj)
}

// UpdateMeshWorkload mocks base method
func (m *MockMeshWorkloadEventHandler) UpdateMeshWorkload(old, new *v1alpha2.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeshWorkload", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeshWorkload indicates an expected call of UpdateMeshWorkload
func (mr *MockMeshWorkloadEventHandlerMockRecorder) UpdateMeshWorkload(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).UpdateMeshWorkload), old, new)
}

// DeleteMeshWorkload mocks base method
func (m *MockMeshWorkloadEventHandler) DeleteMeshWorkload(obj *v1alpha2.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMeshWorkload indicates an expected call of DeleteMeshWorkload
func (mr *MockMeshWorkloadEventHandlerMockRecorder) DeleteMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).DeleteMeshWorkload), obj)
}

// GenericMeshWorkload mocks base method
func (m *MockMeshWorkloadEventHandler) GenericMeshWorkload(obj *v1alpha2.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericMeshWorkload indicates an expected call of GenericMeshWorkload
func (mr *MockMeshWorkloadEventHandlerMockRecorder) GenericMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).GenericMeshWorkload), obj)
}

// MockMeshWorkloadEventWatcher is a mock of MeshWorkloadEventWatcher interface
type MockMeshWorkloadEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadEventWatcherMockRecorder
}

// MockMeshWorkloadEventWatcherMockRecorder is the mock recorder for MockMeshWorkloadEventWatcher
type MockMeshWorkloadEventWatcherMockRecorder struct {
	mock *MockMeshWorkloadEventWatcher
}

// NewMockMeshWorkloadEventWatcher creates a new mock instance
func NewMockMeshWorkloadEventWatcher(ctrl *gomock.Controller) *MockMeshWorkloadEventWatcher {
	mock := &MockMeshWorkloadEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshWorkloadEventWatcher) EXPECT() *MockMeshWorkloadEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockMeshWorkloadEventWatcher) AddEventHandler(ctx context.Context, h controller.MeshWorkloadEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockMeshWorkloadEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshWorkloadEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshEventHandler is a mock of MeshEventHandler interface
type MockMeshEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshEventHandlerMockRecorder
}

// MockMeshEventHandlerMockRecorder is the mock recorder for MockMeshEventHandler
type MockMeshEventHandlerMockRecorder struct {
	mock *MockMeshEventHandler
}

// NewMockMeshEventHandler creates a new mock instance
func NewMockMeshEventHandler(ctrl *gomock.Controller) *MockMeshEventHandler {
	mock := &MockMeshEventHandler{ctrl: ctrl}
	mock.recorder = &MockMeshEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshEventHandler) EXPECT() *MockMeshEventHandlerMockRecorder {
	return m.recorder
}

// CreateMesh mocks base method
func (m *MockMeshEventHandler) CreateMesh(obj *v1alpha2.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMesh indicates an expected call of CreateMesh
func (mr *MockMeshEventHandlerMockRecorder) CreateMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).CreateMesh), obj)
}

// UpdateMesh mocks base method
func (m *MockMeshEventHandler) UpdateMesh(old, new *v1alpha2.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMesh", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMesh indicates an expected call of UpdateMesh
func (mr *MockMeshEventHandlerMockRecorder) UpdateMesh(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).UpdateMesh), old, new)
}

// DeleteMesh mocks base method
func (m *MockMeshEventHandler) DeleteMesh(obj *v1alpha2.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMesh indicates an expected call of DeleteMesh
func (mr *MockMeshEventHandlerMockRecorder) DeleteMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).DeleteMesh), obj)
}

// GenericMesh mocks base method
func (m *MockMeshEventHandler) GenericMesh(obj *v1alpha2.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericMesh indicates an expected call of GenericMesh
func (mr *MockMeshEventHandlerMockRecorder) GenericMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).GenericMesh), obj)
}

// MockMeshEventWatcher is a mock of MeshEventWatcher interface
type MockMeshEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshEventWatcherMockRecorder
}

// MockMeshEventWatcherMockRecorder is the mock recorder for MockMeshEventWatcher
type MockMeshEventWatcherMockRecorder struct {
	mock *MockMeshEventWatcher
}

// NewMockMeshEventWatcher creates a new mock instance
func NewMockMeshEventWatcher(ctrl *gomock.Controller) *MockMeshEventWatcher {
	mock := &MockMeshEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshEventWatcher) EXPECT() *MockMeshEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockMeshEventWatcher) AddEventHandler(ctx context.Context, h controller.MeshEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockMeshEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshEventWatcher)(nil).AddEventHandler), varargs...)
}
