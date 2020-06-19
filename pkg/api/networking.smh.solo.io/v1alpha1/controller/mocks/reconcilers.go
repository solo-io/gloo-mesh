// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockTrafficPolicyReconciler is a mock of TrafficPolicyReconciler interface.
type MockTrafficPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyReconcilerMockRecorder
}

// MockTrafficPolicyReconcilerMockRecorder is the mock recorder for MockTrafficPolicyReconciler.
type MockTrafficPolicyReconcilerMockRecorder struct {
	mock *MockTrafficPolicyReconciler
}

// NewMockTrafficPolicyReconciler creates a new mock instance.
func NewMockTrafficPolicyReconciler(ctrl *gomock.Controller) *MockTrafficPolicyReconciler {
	mock := &MockTrafficPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyReconciler) EXPECT() *MockTrafficPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicy mocks base method.
func (m *MockTrafficPolicyReconciler) ReconcileTrafficPolicy(obj *v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MockTrafficPolicyReconcilerMockRecorder) ReconcileTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyReconciler)(nil).ReconcileTrafficPolicy), obj)
}

// MockTrafficPolicyDeletionReconciler is a mock of TrafficPolicyDeletionReconciler interface.
type MockTrafficPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyDeletionReconcilerMockRecorder
}

// MockTrafficPolicyDeletionReconcilerMockRecorder is the mock recorder for MockTrafficPolicyDeletionReconciler.
type MockTrafficPolicyDeletionReconcilerMockRecorder struct {
	mock *MockTrafficPolicyDeletionReconciler
}

// NewMockTrafficPolicyDeletionReconciler creates a new mock instance.
func NewMockTrafficPolicyDeletionReconciler(ctrl *gomock.Controller) *MockTrafficPolicyDeletionReconciler {
	mock := &MockTrafficPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyDeletionReconciler) EXPECT() *MockTrafficPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicyDeletion mocks base method.
func (m *MockTrafficPolicyDeletionReconciler) ReconcileTrafficPolicyDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileTrafficPolicyDeletion", req)
}

// ReconcileTrafficPolicyDeletion indicates an expected call of ReconcileTrafficPolicyDeletion.
func (mr *MockTrafficPolicyDeletionReconcilerMockRecorder) ReconcileTrafficPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicyDeletion", reflect.TypeOf((*MockTrafficPolicyDeletionReconciler)(nil).ReconcileTrafficPolicyDeletion), req)
}

// MockTrafficPolicyFinalizer is a mock of TrafficPolicyFinalizer interface.
type MockTrafficPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyFinalizerMockRecorder
}

// MockTrafficPolicyFinalizerMockRecorder is the mock recorder for MockTrafficPolicyFinalizer.
type MockTrafficPolicyFinalizerMockRecorder struct {
	mock *MockTrafficPolicyFinalizer
}

// NewMockTrafficPolicyFinalizer creates a new mock instance.
func NewMockTrafficPolicyFinalizer(ctrl *gomock.Controller) *MockTrafficPolicyFinalizer {
	mock := &MockTrafficPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyFinalizer) EXPECT() *MockTrafficPolicyFinalizerMockRecorder {
	return m.recorder
}

// ReconcileTrafficPolicy mocks base method.
func (m *MockTrafficPolicyFinalizer) ReconcileTrafficPolicy(obj *v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficPolicy indicates an expected call of ReconcileTrafficPolicy.
func (mr *MockTrafficPolicyFinalizerMockRecorder) ReconcileTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyFinalizer)(nil).ReconcileTrafficPolicy), obj)
}

// TrafficPolicyFinalizerName mocks base method.
func (m *MockTrafficPolicyFinalizer) TrafficPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrafficPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrafficPolicyFinalizerName indicates an expected call of TrafficPolicyFinalizerName.
func (mr *MockTrafficPolicyFinalizerMockRecorder) TrafficPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrafficPolicyFinalizerName", reflect.TypeOf((*MockTrafficPolicyFinalizer)(nil).TrafficPolicyFinalizerName))
}

// FinalizeTrafficPolicy mocks base method.
func (m *MockTrafficPolicyFinalizer) FinalizeTrafficPolicy(obj *v1alpha1.TrafficPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeTrafficPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeTrafficPolicy indicates an expected call of FinalizeTrafficPolicy.
func (mr *MockTrafficPolicyFinalizerMockRecorder) FinalizeTrafficPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeTrafficPolicy", reflect.TypeOf((*MockTrafficPolicyFinalizer)(nil).FinalizeTrafficPolicy), obj)
}

// MockTrafficPolicyReconcileLoop is a mock of TrafficPolicyReconcileLoop interface.
type MockTrafficPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockTrafficPolicyReconcileLoopMockRecorder
}

// MockTrafficPolicyReconcileLoopMockRecorder is the mock recorder for MockTrafficPolicyReconcileLoop.
type MockTrafficPolicyReconcileLoopMockRecorder struct {
	mock *MockTrafficPolicyReconcileLoop
}

// NewMockTrafficPolicyReconcileLoop creates a new mock instance.
func NewMockTrafficPolicyReconcileLoop(ctrl *gomock.Controller) *MockTrafficPolicyReconcileLoop {
	mock := &MockTrafficPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockTrafficPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrafficPolicyReconcileLoop) EXPECT() *MockTrafficPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunTrafficPolicyReconciler mocks base method.
func (m *MockTrafficPolicyReconcileLoop) RunTrafficPolicyReconciler(ctx context.Context, rec controller.TrafficPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunTrafficPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTrafficPolicyReconciler indicates an expected call of RunTrafficPolicyReconciler.
func (mr *MockTrafficPolicyReconcileLoopMockRecorder) RunTrafficPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTrafficPolicyReconciler", reflect.TypeOf((*MockTrafficPolicyReconcileLoop)(nil).RunTrafficPolicyReconciler), varargs...)
}

// MockAccessControlPolicyReconciler is a mock of AccessControlPolicyReconciler interface.
type MockAccessControlPolicyReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlPolicyReconcilerMockRecorder
}

// MockAccessControlPolicyReconcilerMockRecorder is the mock recorder for MockAccessControlPolicyReconciler.
type MockAccessControlPolicyReconcilerMockRecorder struct {
	mock *MockAccessControlPolicyReconciler
}

// NewMockAccessControlPolicyReconciler creates a new mock instance.
func NewMockAccessControlPolicyReconciler(ctrl *gomock.Controller) *MockAccessControlPolicyReconciler {
	mock := &MockAccessControlPolicyReconciler{ctrl: ctrl}
	mock.recorder = &MockAccessControlPolicyReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlPolicyReconciler) EXPECT() *MockAccessControlPolicyReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessControlPolicy mocks base method.
func (m *MockAccessControlPolicyReconciler) ReconcileAccessControlPolicy(obj *v1alpha1.AccessControlPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessControlPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessControlPolicy indicates an expected call of ReconcileAccessControlPolicy.
func (mr *MockAccessControlPolicyReconcilerMockRecorder) ReconcileAccessControlPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessControlPolicy", reflect.TypeOf((*MockAccessControlPolicyReconciler)(nil).ReconcileAccessControlPolicy), obj)
}

// MockAccessControlPolicyDeletionReconciler is a mock of AccessControlPolicyDeletionReconciler interface.
type MockAccessControlPolicyDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlPolicyDeletionReconcilerMockRecorder
}

// MockAccessControlPolicyDeletionReconcilerMockRecorder is the mock recorder for MockAccessControlPolicyDeletionReconciler.
type MockAccessControlPolicyDeletionReconcilerMockRecorder struct {
	mock *MockAccessControlPolicyDeletionReconciler
}

// NewMockAccessControlPolicyDeletionReconciler creates a new mock instance.
func NewMockAccessControlPolicyDeletionReconciler(ctrl *gomock.Controller) *MockAccessControlPolicyDeletionReconciler {
	mock := &MockAccessControlPolicyDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockAccessControlPolicyDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlPolicyDeletionReconciler) EXPECT() *MockAccessControlPolicyDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileAccessControlPolicyDeletion mocks base method.
func (m *MockAccessControlPolicyDeletionReconciler) ReconcileAccessControlPolicyDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileAccessControlPolicyDeletion", req)
}

// ReconcileAccessControlPolicyDeletion indicates an expected call of ReconcileAccessControlPolicyDeletion.
func (mr *MockAccessControlPolicyDeletionReconcilerMockRecorder) ReconcileAccessControlPolicyDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessControlPolicyDeletion", reflect.TypeOf((*MockAccessControlPolicyDeletionReconciler)(nil).ReconcileAccessControlPolicyDeletion), req)
}

// MockAccessControlPolicyFinalizer is a mock of AccessControlPolicyFinalizer interface.
type MockAccessControlPolicyFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlPolicyFinalizerMockRecorder
}

// MockAccessControlPolicyFinalizerMockRecorder is the mock recorder for MockAccessControlPolicyFinalizer.
type MockAccessControlPolicyFinalizerMockRecorder struct {
	mock *MockAccessControlPolicyFinalizer
}

// NewMockAccessControlPolicyFinalizer creates a new mock instance.
func NewMockAccessControlPolicyFinalizer(ctrl *gomock.Controller) *MockAccessControlPolicyFinalizer {
	mock := &MockAccessControlPolicyFinalizer{ctrl: ctrl}
	mock.recorder = &MockAccessControlPolicyFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlPolicyFinalizer) EXPECT() *MockAccessControlPolicyFinalizerMockRecorder {
	return m.recorder
}

// ReconcileAccessControlPolicy mocks base method.
func (m *MockAccessControlPolicyFinalizer) ReconcileAccessControlPolicy(obj *v1alpha1.AccessControlPolicy) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileAccessControlPolicy", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileAccessControlPolicy indicates an expected call of ReconcileAccessControlPolicy.
func (mr *MockAccessControlPolicyFinalizerMockRecorder) ReconcileAccessControlPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileAccessControlPolicy", reflect.TypeOf((*MockAccessControlPolicyFinalizer)(nil).ReconcileAccessControlPolicy), obj)
}

// AccessControlPolicyFinalizerName mocks base method.
func (m *MockAccessControlPolicyFinalizer) AccessControlPolicyFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessControlPolicyFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// AccessControlPolicyFinalizerName indicates an expected call of AccessControlPolicyFinalizerName.
func (mr *MockAccessControlPolicyFinalizerMockRecorder) AccessControlPolicyFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessControlPolicyFinalizerName", reflect.TypeOf((*MockAccessControlPolicyFinalizer)(nil).AccessControlPolicyFinalizerName))
}

// FinalizeAccessControlPolicy mocks base method.
func (m *MockAccessControlPolicyFinalizer) FinalizeAccessControlPolicy(obj *v1alpha1.AccessControlPolicy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAccessControlPolicy", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeAccessControlPolicy indicates an expected call of FinalizeAccessControlPolicy.
func (mr *MockAccessControlPolicyFinalizerMockRecorder) FinalizeAccessControlPolicy(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAccessControlPolicy", reflect.TypeOf((*MockAccessControlPolicyFinalizer)(nil).FinalizeAccessControlPolicy), obj)
}

// MockAccessControlPolicyReconcileLoop is a mock of AccessControlPolicyReconcileLoop interface.
type MockAccessControlPolicyReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockAccessControlPolicyReconcileLoopMockRecorder
}

// MockAccessControlPolicyReconcileLoopMockRecorder is the mock recorder for MockAccessControlPolicyReconcileLoop.
type MockAccessControlPolicyReconcileLoopMockRecorder struct {
	mock *MockAccessControlPolicyReconcileLoop
}

// NewMockAccessControlPolicyReconcileLoop creates a new mock instance.
func NewMockAccessControlPolicyReconcileLoop(ctrl *gomock.Controller) *MockAccessControlPolicyReconcileLoop {
	mock := &MockAccessControlPolicyReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockAccessControlPolicyReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessControlPolicyReconcileLoop) EXPECT() *MockAccessControlPolicyReconcileLoopMockRecorder {
	return m.recorder
}

// RunAccessControlPolicyReconciler mocks base method.
func (m *MockAccessControlPolicyReconcileLoop) RunAccessControlPolicyReconciler(ctx context.Context, rec controller.AccessControlPolicyReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAccessControlPolicyReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAccessControlPolicyReconciler indicates an expected call of RunAccessControlPolicyReconciler.
func (mr *MockAccessControlPolicyReconcileLoopMockRecorder) RunAccessControlPolicyReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAccessControlPolicyReconciler", reflect.TypeOf((*MockAccessControlPolicyReconcileLoop)(nil).RunAccessControlPolicyReconciler), varargs...)
}

// MockVirtualMeshReconciler is a mock of VirtualMeshReconciler interface.
type MockVirtualMeshReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshReconcilerMockRecorder
}

// MockVirtualMeshReconcilerMockRecorder is the mock recorder for MockVirtualMeshReconciler.
type MockVirtualMeshReconcilerMockRecorder struct {
	mock *MockVirtualMeshReconciler
}

// NewMockVirtualMeshReconciler creates a new mock instance.
func NewMockVirtualMeshReconciler(ctrl *gomock.Controller) *MockVirtualMeshReconciler {
	mock := &MockVirtualMeshReconciler{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshReconciler) EXPECT() *MockVirtualMeshReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMesh mocks base method.
func (m *MockVirtualMeshReconciler) ReconcileVirtualMesh(obj *v1alpha1.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MockVirtualMeshReconcilerMockRecorder) ReconcileVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MockVirtualMeshReconciler)(nil).ReconcileVirtualMesh), obj)
}

// MockVirtualMeshDeletionReconciler is a mock of VirtualMeshDeletionReconciler interface.
type MockVirtualMeshDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshDeletionReconcilerMockRecorder
}

// MockVirtualMeshDeletionReconcilerMockRecorder is the mock recorder for MockVirtualMeshDeletionReconciler.
type MockVirtualMeshDeletionReconcilerMockRecorder struct {
	mock *MockVirtualMeshDeletionReconciler
}

// NewMockVirtualMeshDeletionReconciler creates a new mock instance.
func NewMockVirtualMeshDeletionReconciler(ctrl *gomock.Controller) *MockVirtualMeshDeletionReconciler {
	mock := &MockVirtualMeshDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshDeletionReconciler) EXPECT() *MockVirtualMeshDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMeshDeletion mocks base method.
func (m *MockVirtualMeshDeletionReconciler) ReconcileVirtualMeshDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVirtualMeshDeletion", req)
}

// ReconcileVirtualMeshDeletion indicates an expected call of ReconcileVirtualMeshDeletion.
func (mr *MockVirtualMeshDeletionReconcilerMockRecorder) ReconcileVirtualMeshDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMeshDeletion", reflect.TypeOf((*MockVirtualMeshDeletionReconciler)(nil).ReconcileVirtualMeshDeletion), req)
}

// MockVirtualMeshFinalizer is a mock of VirtualMeshFinalizer interface.
type MockVirtualMeshFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshFinalizerMockRecorder
}

// MockVirtualMeshFinalizerMockRecorder is the mock recorder for MockVirtualMeshFinalizer.
type MockVirtualMeshFinalizerMockRecorder struct {
	mock *MockVirtualMeshFinalizer
}

// NewMockVirtualMeshFinalizer creates a new mock instance.
func NewMockVirtualMeshFinalizer(ctrl *gomock.Controller) *MockVirtualMeshFinalizer {
	mock := &MockVirtualMeshFinalizer{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshFinalizer) EXPECT() *MockVirtualMeshFinalizerMockRecorder {
	return m.recorder
}

// ReconcileVirtualMesh mocks base method.
func (m *MockVirtualMeshFinalizer) ReconcileVirtualMesh(obj *v1alpha1.VirtualMesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualMesh", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileVirtualMesh indicates an expected call of ReconcileVirtualMesh.
func (mr *MockVirtualMeshFinalizerMockRecorder) ReconcileVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualMesh", reflect.TypeOf((*MockVirtualMeshFinalizer)(nil).ReconcileVirtualMesh), obj)
}

// VirtualMeshFinalizerName mocks base method.
func (m *MockVirtualMeshFinalizer) VirtualMeshFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMeshFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// VirtualMeshFinalizerName indicates an expected call of VirtualMeshFinalizerName.
func (mr *MockVirtualMeshFinalizerMockRecorder) VirtualMeshFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMeshFinalizerName", reflect.TypeOf((*MockVirtualMeshFinalizer)(nil).VirtualMeshFinalizerName))
}

// FinalizeVirtualMesh mocks base method.
func (m *MockVirtualMeshFinalizer) FinalizeVirtualMesh(obj *v1alpha1.VirtualMesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeVirtualMesh", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeVirtualMesh indicates an expected call of FinalizeVirtualMesh.
func (mr *MockVirtualMeshFinalizerMockRecorder) FinalizeVirtualMesh(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeVirtualMesh", reflect.TypeOf((*MockVirtualMeshFinalizer)(nil).FinalizeVirtualMesh), obj)
}

// MockVirtualMeshReconcileLoop is a mock of VirtualMeshReconcileLoop interface.
type MockVirtualMeshReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshReconcileLoopMockRecorder
}

// MockVirtualMeshReconcileLoopMockRecorder is the mock recorder for MockVirtualMeshReconcileLoop.
type MockVirtualMeshReconcileLoopMockRecorder struct {
	mock *MockVirtualMeshReconcileLoop
}

// NewMockVirtualMeshReconcileLoop creates a new mock instance.
func NewMockVirtualMeshReconcileLoop(ctrl *gomock.Controller) *MockVirtualMeshReconcileLoop {
	mock := &MockVirtualMeshReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshReconcileLoop) EXPECT() *MockVirtualMeshReconcileLoopMockRecorder {
	return m.recorder
}

// RunVirtualMeshReconciler mocks base method.
func (m *MockVirtualMeshReconcileLoop) RunVirtualMeshReconciler(ctx context.Context, rec controller.VirtualMeshReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunVirtualMeshReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunVirtualMeshReconciler indicates an expected call of RunVirtualMeshReconciler.
func (mr *MockVirtualMeshReconcileLoopMockRecorder) RunVirtualMeshReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunVirtualMeshReconciler", reflect.TypeOf((*MockVirtualMeshReconcileLoop)(nil).RunVirtualMeshReconciler), varargs...)
}
