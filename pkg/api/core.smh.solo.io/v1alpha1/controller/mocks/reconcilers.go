// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockSettingsReconciler is a mock of SettingsReconciler interface.
type MockSettingsReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReconcilerMockRecorder
}

// MockSettingsReconcilerMockRecorder is the mock recorder for MockSettingsReconciler.
type MockSettingsReconcilerMockRecorder struct {
	mock *MockSettingsReconciler
}

// NewMockSettingsReconciler creates a new mock instance.
func NewMockSettingsReconciler(ctrl *gomock.Controller) *MockSettingsReconciler {
	mock := &MockSettingsReconciler{ctrl: ctrl}
	mock.recorder = &MockSettingsReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsReconciler) EXPECT() *MockSettingsReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method.
func (m *MockSettingsReconciler) ReconcileSettings(obj *v1alpha1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings.
func (mr *MockSettingsReconcilerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockSettingsReconciler)(nil).ReconcileSettings), obj)
}

// MockSettingsDeletionReconciler is a mock of SettingsDeletionReconciler interface.
type MockSettingsDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsDeletionReconcilerMockRecorder
}

// MockSettingsDeletionReconcilerMockRecorder is the mock recorder for MockSettingsDeletionReconciler.
type MockSettingsDeletionReconcilerMockRecorder struct {
	mock *MockSettingsDeletionReconciler
}

// NewMockSettingsDeletionReconciler creates a new mock instance.
func NewMockSettingsDeletionReconciler(ctrl *gomock.Controller) *MockSettingsDeletionReconciler {
	mock := &MockSettingsDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockSettingsDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsDeletionReconciler) EXPECT() *MockSettingsDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileSettingsDeletion mocks base method.
func (m *MockSettingsDeletionReconciler) ReconcileSettingsDeletion(req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileSettingsDeletion", req)
}

// ReconcileSettingsDeletion indicates an expected call of ReconcileSettingsDeletion.
func (mr *MockSettingsDeletionReconcilerMockRecorder) ReconcileSettingsDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettingsDeletion", reflect.TypeOf((*MockSettingsDeletionReconciler)(nil).ReconcileSettingsDeletion), req)
}

// MockSettingsFinalizer is a mock of SettingsFinalizer interface.
type MockSettingsFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsFinalizerMockRecorder
}

// MockSettingsFinalizerMockRecorder is the mock recorder for MockSettingsFinalizer.
type MockSettingsFinalizerMockRecorder struct {
	mock *MockSettingsFinalizer
}

// NewMockSettingsFinalizer creates a new mock instance.
func NewMockSettingsFinalizer(ctrl *gomock.Controller) *MockSettingsFinalizer {
	mock := &MockSettingsFinalizer{ctrl: ctrl}
	mock.recorder = &MockSettingsFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsFinalizer) EXPECT() *MockSettingsFinalizerMockRecorder {
	return m.recorder
}

// ReconcileSettings mocks base method.
func (m *MockSettingsFinalizer) ReconcileSettings(obj *v1alpha1.Settings) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSettings", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileSettings indicates an expected call of ReconcileSettings.
func (mr *MockSettingsFinalizerMockRecorder) ReconcileSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSettings", reflect.TypeOf((*MockSettingsFinalizer)(nil).ReconcileSettings), obj)
}

// SettingsFinalizerName mocks base method.
func (m *MockSettingsFinalizer) SettingsFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettingsFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// SettingsFinalizerName indicates an expected call of SettingsFinalizerName.
func (mr *MockSettingsFinalizerMockRecorder) SettingsFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettingsFinalizerName", reflect.TypeOf((*MockSettingsFinalizer)(nil).SettingsFinalizerName))
}

// FinalizeSettings mocks base method.
func (m *MockSettingsFinalizer) FinalizeSettings(obj *v1alpha1.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeSettings", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeSettings indicates an expected call of FinalizeSettings.
func (mr *MockSettingsFinalizerMockRecorder) FinalizeSettings(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeSettings", reflect.TypeOf((*MockSettingsFinalizer)(nil).FinalizeSettings), obj)
}

// MockSettingsReconcileLoop is a mock of SettingsReconcileLoop interface.
type MockSettingsReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsReconcileLoopMockRecorder
}

// MockSettingsReconcileLoopMockRecorder is the mock recorder for MockSettingsReconcileLoop.
type MockSettingsReconcileLoopMockRecorder struct {
	mock *MockSettingsReconcileLoop
}

// NewMockSettingsReconcileLoop creates a new mock instance.
func NewMockSettingsReconcileLoop(ctrl *gomock.Controller) *MockSettingsReconcileLoop {
	mock := &MockSettingsReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockSettingsReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsReconcileLoop) EXPECT() *MockSettingsReconcileLoopMockRecorder {
	return m.recorder
}

// RunSettingsReconciler mocks base method.
func (m *MockSettingsReconcileLoop) RunSettingsReconciler(ctx context.Context, rec controller.SettingsReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunSettingsReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunSettingsReconciler indicates an expected call of RunSettingsReconciler.
func (mr *MockSettingsReconcileLoopMockRecorder) RunSettingsReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSettingsReconciler", reflect.TypeOf((*MockSettingsReconcileLoop)(nil).RunSettingsReconciler), varargs...)
}
