// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockKubernetesClusterEventHandler is a mock of KubernetesClusterEventHandler interface.
type MockKubernetesClusterEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterEventHandlerMockRecorder
}

// MockKubernetesClusterEventHandlerMockRecorder is the mock recorder for MockKubernetesClusterEventHandler.
type MockKubernetesClusterEventHandlerMockRecorder struct {
	mock *MockKubernetesClusterEventHandler
}

// NewMockKubernetesClusterEventHandler creates a new mock instance.
func NewMockKubernetesClusterEventHandler(ctrl *gomock.Controller) *MockKubernetesClusterEventHandler {
	mock := &MockKubernetesClusterEventHandler{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesClusterEventHandler) EXPECT() *MockKubernetesClusterEventHandlerMockRecorder {
	return m.recorder
}

// CreateKubernetesCluster mocks base method.
func (m *MockKubernetesClusterEventHandler) CreateKubernetesCluster(obj *v1alpha1.KubernetesCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKubernetesCluster", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateKubernetesCluster indicates an expected call of CreateKubernetesCluster.
func (mr *MockKubernetesClusterEventHandlerMockRecorder) CreateKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterEventHandler)(nil).CreateKubernetesCluster), obj)
}

// UpdateKubernetesCluster mocks base method.
func (m *MockKubernetesClusterEventHandler) UpdateKubernetesCluster(old, new *v1alpha1.KubernetesCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKubernetesCluster", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateKubernetesCluster indicates an expected call of UpdateKubernetesCluster.
func (mr *MockKubernetesClusterEventHandlerMockRecorder) UpdateKubernetesCluster(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterEventHandler)(nil).UpdateKubernetesCluster), old, new)
}

// DeleteKubernetesCluster mocks base method.
func (m *MockKubernetesClusterEventHandler) DeleteKubernetesCluster(obj *v1alpha1.KubernetesCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKubernetesCluster", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteKubernetesCluster indicates an expected call of DeleteKubernetesCluster.
func (mr *MockKubernetesClusterEventHandlerMockRecorder) DeleteKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterEventHandler)(nil).DeleteKubernetesCluster), obj)
}

// GenericKubernetesCluster mocks base method.
func (m *MockKubernetesClusterEventHandler) GenericKubernetesCluster(obj *v1alpha1.KubernetesCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericKubernetesCluster", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericKubernetesCluster indicates an expected call of GenericKubernetesCluster.
func (mr *MockKubernetesClusterEventHandlerMockRecorder) GenericKubernetesCluster(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericKubernetesCluster", reflect.TypeOf((*MockKubernetesClusterEventHandler)(nil).GenericKubernetesCluster), obj)
}

// MockKubernetesClusterEventWatcher is a mock of KubernetesClusterEventWatcher interface.
type MockKubernetesClusterEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClusterEventWatcherMockRecorder
}

// MockKubernetesClusterEventWatcherMockRecorder is the mock recorder for MockKubernetesClusterEventWatcher.
type MockKubernetesClusterEventWatcherMockRecorder struct {
	mock *MockKubernetesClusterEventWatcher
}

// NewMockKubernetesClusterEventWatcher creates a new mock instance.
func NewMockKubernetesClusterEventWatcher(ctrl *gomock.Controller) *MockKubernetesClusterEventWatcher {
	mock := &MockKubernetesClusterEventWatcher{ctrl: ctrl}
	mock.recorder = &MockKubernetesClusterEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesClusterEventWatcher) EXPECT() *MockKubernetesClusterEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockKubernetesClusterEventWatcher) AddEventHandler(ctx context.Context, h controller.KubernetesClusterEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockKubernetesClusterEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockKubernetesClusterEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshServiceEventHandler is a mock of MeshServiceEventHandler interface.
type MockMeshServiceEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshServiceEventHandlerMockRecorder
}

// MockMeshServiceEventHandlerMockRecorder is the mock recorder for MockMeshServiceEventHandler.
type MockMeshServiceEventHandlerMockRecorder struct {
	mock *MockMeshServiceEventHandler
}

// NewMockMeshServiceEventHandler creates a new mock instance.
func NewMockMeshServiceEventHandler(ctrl *gomock.Controller) *MockMeshServiceEventHandler {
	mock := &MockMeshServiceEventHandler{ctrl: ctrl}
	mock.recorder = &MockMeshServiceEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshServiceEventHandler) EXPECT() *MockMeshServiceEventHandlerMockRecorder {
	return m.recorder
}

// CreateMeshService mocks base method.
func (m *MockMeshServiceEventHandler) CreateMeshService(obj *v1alpha1.MeshService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeshService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMeshService indicates an expected call of CreateMeshService.
func (mr *MockMeshServiceEventHandlerMockRecorder) CreateMeshService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeshService", reflect.TypeOf((*MockMeshServiceEventHandler)(nil).CreateMeshService), obj)
}

// UpdateMeshService mocks base method.
func (m *MockMeshServiceEventHandler) UpdateMeshService(old, new *v1alpha1.MeshService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeshService", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeshService indicates an expected call of UpdateMeshService.
func (mr *MockMeshServiceEventHandlerMockRecorder) UpdateMeshService(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeshService", reflect.TypeOf((*MockMeshServiceEventHandler)(nil).UpdateMeshService), old, new)
}

// DeleteMeshService mocks base method.
func (m *MockMeshServiceEventHandler) DeleteMeshService(obj *v1alpha1.MeshService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeshService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMeshService indicates an expected call of DeleteMeshService.
func (mr *MockMeshServiceEventHandlerMockRecorder) DeleteMeshService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeshService", reflect.TypeOf((*MockMeshServiceEventHandler)(nil).DeleteMeshService), obj)
}

// GenericMeshService mocks base method.
func (m *MockMeshServiceEventHandler) GenericMeshService(obj *v1alpha1.MeshService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericMeshService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericMeshService indicates an expected call of GenericMeshService.
func (mr *MockMeshServiceEventHandlerMockRecorder) GenericMeshService(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericMeshService", reflect.TypeOf((*MockMeshServiceEventHandler)(nil).GenericMeshService), obj)
}

// MockMeshServiceEventWatcher is a mock of MeshServiceEventWatcher interface.
type MockMeshServiceEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshServiceEventWatcherMockRecorder
}

// MockMeshServiceEventWatcherMockRecorder is the mock recorder for MockMeshServiceEventWatcher.
type MockMeshServiceEventWatcherMockRecorder struct {
	mock *MockMeshServiceEventWatcher
}

// NewMockMeshServiceEventWatcher creates a new mock instance.
func NewMockMeshServiceEventWatcher(ctrl *gomock.Controller) *MockMeshServiceEventWatcher {
	mock := &MockMeshServiceEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshServiceEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshServiceEventWatcher) EXPECT() *MockMeshServiceEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockMeshServiceEventWatcher) AddEventHandler(ctx context.Context, h controller.MeshServiceEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockMeshServiceEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshServiceEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshWorkloadEventHandler is a mock of MeshWorkloadEventHandler interface.
type MockMeshWorkloadEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadEventHandlerMockRecorder
}

// MockMeshWorkloadEventHandlerMockRecorder is the mock recorder for MockMeshWorkloadEventHandler.
type MockMeshWorkloadEventHandlerMockRecorder struct {
	mock *MockMeshWorkloadEventHandler
}

// NewMockMeshWorkloadEventHandler creates a new mock instance.
func NewMockMeshWorkloadEventHandler(ctrl *gomock.Controller) *MockMeshWorkloadEventHandler {
	mock := &MockMeshWorkloadEventHandler{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshWorkloadEventHandler) EXPECT() *MockMeshWorkloadEventHandlerMockRecorder {
	return m.recorder
}

// CreateMeshWorkload mocks base method.
func (m *MockMeshWorkloadEventHandler) CreateMeshWorkload(obj *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMeshWorkload indicates an expected call of CreateMeshWorkload.
func (mr *MockMeshWorkloadEventHandlerMockRecorder) CreateMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).CreateMeshWorkload), obj)
}

// UpdateMeshWorkload mocks base method.
func (m *MockMeshWorkloadEventHandler) UpdateMeshWorkload(old, new *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeshWorkload", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeshWorkload indicates an expected call of UpdateMeshWorkload.
func (mr *MockMeshWorkloadEventHandlerMockRecorder) UpdateMeshWorkload(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).UpdateMeshWorkload), old, new)
}

// DeleteMeshWorkload mocks base method.
func (m *MockMeshWorkloadEventHandler) DeleteMeshWorkload(obj *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMeshWorkload indicates an expected call of DeleteMeshWorkload.
func (mr *MockMeshWorkloadEventHandlerMockRecorder) DeleteMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).DeleteMeshWorkload), obj)
}

// GenericMeshWorkload mocks base method.
func (m *MockMeshWorkloadEventHandler) GenericMeshWorkload(obj *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericMeshWorkload indicates an expected call of GenericMeshWorkload.
func (mr *MockMeshWorkloadEventHandlerMockRecorder) GenericMeshWorkload(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericMeshWorkload", reflect.TypeOf((*MockMeshWorkloadEventHandler)(nil).GenericMeshWorkload), obj)
}

// MockMeshWorkloadEventWatcher is a mock of MeshWorkloadEventWatcher interface.
type MockMeshWorkloadEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadEventWatcherMockRecorder
}

// MockMeshWorkloadEventWatcherMockRecorder is the mock recorder for MockMeshWorkloadEventWatcher.
type MockMeshWorkloadEventWatcherMockRecorder struct {
	mock *MockMeshWorkloadEventWatcher
}

// NewMockMeshWorkloadEventWatcher creates a new mock instance.
func NewMockMeshWorkloadEventWatcher(ctrl *gomock.Controller) *MockMeshWorkloadEventWatcher {
	mock := &MockMeshWorkloadEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshWorkloadEventWatcher) EXPECT() *MockMeshWorkloadEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
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

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockMeshWorkloadEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshWorkloadEventWatcher)(nil).AddEventHandler), varargs...)
}

// MockMeshEventHandler is a mock of MeshEventHandler interface.
type MockMeshEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshEventHandlerMockRecorder
}

// MockMeshEventHandlerMockRecorder is the mock recorder for MockMeshEventHandler.
type MockMeshEventHandlerMockRecorder struct {
	mock *MockMeshEventHandler
}

// NewMockMeshEventHandler creates a new mock instance.
func NewMockMeshEventHandler(ctrl *gomock.Controller) *MockMeshEventHandler {
	mock := &MockMeshEventHandler{ctrl: ctrl}
	mock.recorder = &MockMeshEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshEventHandler) EXPECT() *MockMeshEventHandlerMockRecorder {
	return m.recorder
}

// CreateMesh mocks base method.
func (m *MockMeshEventHandler) CreateMesh(obj *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMesh indicates an expected call of CreateMesh.
func (mr *MockMeshEventHandlerMockRecorder) CreateMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).CreateMesh), obj)
}

// UpdateMesh mocks base method.
func (m *MockMeshEventHandler) UpdateMesh(old, new *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMesh", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMesh indicates an expected call of UpdateMesh.
func (mr *MockMeshEventHandlerMockRecorder) UpdateMesh(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).UpdateMesh), old, new)
}

// DeleteMesh mocks base method.
func (m *MockMeshEventHandler) DeleteMesh(obj *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMesh indicates an expected call of DeleteMesh.
func (mr *MockMeshEventHandlerMockRecorder) DeleteMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).DeleteMesh), obj)
}

// GenericMesh mocks base method.
func (m *MockMeshEventHandler) GenericMesh(obj *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericMesh indicates an expected call of GenericMesh.
func (mr *MockMeshEventHandlerMockRecorder) GenericMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericMesh", reflect.TypeOf((*MockMeshEventHandler)(nil).GenericMesh), obj)
}

// MockMeshEventWatcher is a mock of MeshEventWatcher interface.
type MockMeshEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMeshEventWatcherMockRecorder
}

// MockMeshEventWatcherMockRecorder is the mock recorder for MockMeshEventWatcher.
type MockMeshEventWatcherMockRecorder struct {
	mock *MockMeshEventWatcher
}

// NewMockMeshEventWatcher creates a new mock instance.
func NewMockMeshEventWatcher(ctrl *gomock.Controller) *MockMeshEventWatcher {
	mock := &MockMeshEventWatcher{ctrl: ctrl}
	mock.recorder = &MockMeshEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshEventWatcher) EXPECT() *MockMeshEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
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

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockMeshEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockMeshEventWatcher)(nil).AddEventHandler), varargs...)
}
