// Code generated by MockGen. DO NOT EDIT.
// Source: ./reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2"
	controller "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1alpha2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	reflect "reflect"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockIssuedCertificateReconciler is a mock of IssuedCertificateReconciler interface
type MockIssuedCertificateReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockIssuedCertificateReconcilerMockRecorder
}

// MockIssuedCertificateReconcilerMockRecorder is the mock recorder for MockIssuedCertificateReconciler
type MockIssuedCertificateReconcilerMockRecorder struct {
	mock *MockIssuedCertificateReconciler
}

// NewMockIssuedCertificateReconciler creates a new mock instance
func NewMockIssuedCertificateReconciler(ctrl *gomock.Controller) *MockIssuedCertificateReconciler {
	mock := &MockIssuedCertificateReconciler{ctrl: ctrl}
	mock.recorder = &MockIssuedCertificateReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIssuedCertificateReconciler) EXPECT() *MockIssuedCertificateReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIssuedCertificate mocks base method
func (m *MockIssuedCertificateReconciler) ReconcileIssuedCertificate(obj *v1alpha2.IssuedCertificate) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIssuedCertificate", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileIssuedCertificate indicates an expected call of ReconcileIssuedCertificate
func (mr *MockIssuedCertificateReconcilerMockRecorder) ReconcileIssuedCertificate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateReconciler)(nil).ReconcileIssuedCertificate), obj)
}

// MockIssuedCertificateDeletionReconciler is a mock of IssuedCertificateDeletionReconciler interface
type MockIssuedCertificateDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockIssuedCertificateDeletionReconcilerMockRecorder
}

// MockIssuedCertificateDeletionReconcilerMockRecorder is the mock recorder for MockIssuedCertificateDeletionReconciler
type MockIssuedCertificateDeletionReconcilerMockRecorder struct {
	mock *MockIssuedCertificateDeletionReconciler
}

// NewMockIssuedCertificateDeletionReconciler creates a new mock instance
func NewMockIssuedCertificateDeletionReconciler(ctrl *gomock.Controller) *MockIssuedCertificateDeletionReconciler {
	mock := &MockIssuedCertificateDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockIssuedCertificateDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIssuedCertificateDeletionReconciler) EXPECT() *MockIssuedCertificateDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileIssuedCertificateDeletion mocks base method
func (m *MockIssuedCertificateDeletionReconciler) ReconcileIssuedCertificateDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIssuedCertificateDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileIssuedCertificateDeletion indicates an expected call of ReconcileIssuedCertificateDeletion
func (mr *MockIssuedCertificateDeletionReconcilerMockRecorder) ReconcileIssuedCertificateDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIssuedCertificateDeletion", reflect.TypeOf((*MockIssuedCertificateDeletionReconciler)(nil).ReconcileIssuedCertificateDeletion), req)
}

// MockIssuedCertificateFinalizer is a mock of IssuedCertificateFinalizer interface
type MockIssuedCertificateFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockIssuedCertificateFinalizerMockRecorder
}

// MockIssuedCertificateFinalizerMockRecorder is the mock recorder for MockIssuedCertificateFinalizer
type MockIssuedCertificateFinalizerMockRecorder struct {
	mock *MockIssuedCertificateFinalizer
}

// NewMockIssuedCertificateFinalizer creates a new mock instance
func NewMockIssuedCertificateFinalizer(ctrl *gomock.Controller) *MockIssuedCertificateFinalizer {
	mock := &MockIssuedCertificateFinalizer{ctrl: ctrl}
	mock.recorder = &MockIssuedCertificateFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIssuedCertificateFinalizer) EXPECT() *MockIssuedCertificateFinalizerMockRecorder {
	return m.recorder
}

// ReconcileIssuedCertificate mocks base method
func (m *MockIssuedCertificateFinalizer) ReconcileIssuedCertificate(obj *v1alpha2.IssuedCertificate) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileIssuedCertificate", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileIssuedCertificate indicates an expected call of ReconcileIssuedCertificate
func (mr *MockIssuedCertificateFinalizerMockRecorder) ReconcileIssuedCertificate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateFinalizer)(nil).ReconcileIssuedCertificate), obj)
}

// IssuedCertificateFinalizerName mocks base method
func (m *MockIssuedCertificateFinalizer) IssuedCertificateFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssuedCertificateFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// IssuedCertificateFinalizerName indicates an expected call of IssuedCertificateFinalizerName
func (mr *MockIssuedCertificateFinalizerMockRecorder) IssuedCertificateFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssuedCertificateFinalizerName", reflect.TypeOf((*MockIssuedCertificateFinalizer)(nil).IssuedCertificateFinalizerName))
}

// FinalizeIssuedCertificate mocks base method
func (m *MockIssuedCertificateFinalizer) FinalizeIssuedCertificate(obj *v1alpha2.IssuedCertificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeIssuedCertificate", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeIssuedCertificate indicates an expected call of FinalizeIssuedCertificate
func (mr *MockIssuedCertificateFinalizerMockRecorder) FinalizeIssuedCertificate(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeIssuedCertificate", reflect.TypeOf((*MockIssuedCertificateFinalizer)(nil).FinalizeIssuedCertificate), obj)
}

// MockIssuedCertificateReconcileLoop is a mock of IssuedCertificateReconcileLoop interface
type MockIssuedCertificateReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockIssuedCertificateReconcileLoopMockRecorder
}

// MockIssuedCertificateReconcileLoopMockRecorder is the mock recorder for MockIssuedCertificateReconcileLoop
type MockIssuedCertificateReconcileLoopMockRecorder struct {
	mock *MockIssuedCertificateReconcileLoop
}

// NewMockIssuedCertificateReconcileLoop creates a new mock instance
func NewMockIssuedCertificateReconcileLoop(ctrl *gomock.Controller) *MockIssuedCertificateReconcileLoop {
	mock := &MockIssuedCertificateReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockIssuedCertificateReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIssuedCertificateReconcileLoop) EXPECT() *MockIssuedCertificateReconcileLoopMockRecorder {
	return m.recorder
}

// RunIssuedCertificateReconciler mocks base method
func (m *MockIssuedCertificateReconcileLoop) RunIssuedCertificateReconciler(ctx context.Context, rec controller.IssuedCertificateReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunIssuedCertificateReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunIssuedCertificateReconciler indicates an expected call of RunIssuedCertificateReconciler
func (mr *MockIssuedCertificateReconcileLoopMockRecorder) RunIssuedCertificateReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunIssuedCertificateReconciler", reflect.TypeOf((*MockIssuedCertificateReconcileLoop)(nil).RunIssuedCertificateReconciler), varargs...)
}

// MockCertificateRequestReconciler is a mock of CertificateRequestReconciler interface
type MockCertificateRequestReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateRequestReconcilerMockRecorder
}

// MockCertificateRequestReconcilerMockRecorder is the mock recorder for MockCertificateRequestReconciler
type MockCertificateRequestReconcilerMockRecorder struct {
	mock *MockCertificateRequestReconciler
}

// NewMockCertificateRequestReconciler creates a new mock instance
func NewMockCertificateRequestReconciler(ctrl *gomock.Controller) *MockCertificateRequestReconciler {
	mock := &MockCertificateRequestReconciler{ctrl: ctrl}
	mock.recorder = &MockCertificateRequestReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateRequestReconciler) EXPECT() *MockCertificateRequestReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCertificateRequest mocks base method
func (m *MockCertificateRequestReconciler) ReconcileCertificateRequest(obj *v1alpha2.CertificateRequest) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateRequest", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCertificateRequest indicates an expected call of ReconcileCertificateRequest
func (mr *MockCertificateRequestReconcilerMockRecorder) ReconcileCertificateRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateRequest", reflect.TypeOf((*MockCertificateRequestReconciler)(nil).ReconcileCertificateRequest), obj)
}

// MockCertificateRequestDeletionReconciler is a mock of CertificateRequestDeletionReconciler interface
type MockCertificateRequestDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateRequestDeletionReconcilerMockRecorder
}

// MockCertificateRequestDeletionReconcilerMockRecorder is the mock recorder for MockCertificateRequestDeletionReconciler
type MockCertificateRequestDeletionReconcilerMockRecorder struct {
	mock *MockCertificateRequestDeletionReconciler
}

// NewMockCertificateRequestDeletionReconciler creates a new mock instance
func NewMockCertificateRequestDeletionReconciler(ctrl *gomock.Controller) *MockCertificateRequestDeletionReconciler {
	mock := &MockCertificateRequestDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockCertificateRequestDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateRequestDeletionReconciler) EXPECT() *MockCertificateRequestDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileCertificateRequestDeletion mocks base method
func (m *MockCertificateRequestDeletionReconciler) ReconcileCertificateRequestDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateRequestDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileCertificateRequestDeletion indicates an expected call of ReconcileCertificateRequestDeletion
func (mr *MockCertificateRequestDeletionReconcilerMockRecorder) ReconcileCertificateRequestDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateRequestDeletion", reflect.TypeOf((*MockCertificateRequestDeletionReconciler)(nil).ReconcileCertificateRequestDeletion), req)
}

// MockCertificateRequestFinalizer is a mock of CertificateRequestFinalizer interface
type MockCertificateRequestFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateRequestFinalizerMockRecorder
}

// MockCertificateRequestFinalizerMockRecorder is the mock recorder for MockCertificateRequestFinalizer
type MockCertificateRequestFinalizerMockRecorder struct {
	mock *MockCertificateRequestFinalizer
}

// NewMockCertificateRequestFinalizer creates a new mock instance
func NewMockCertificateRequestFinalizer(ctrl *gomock.Controller) *MockCertificateRequestFinalizer {
	mock := &MockCertificateRequestFinalizer{ctrl: ctrl}
	mock.recorder = &MockCertificateRequestFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateRequestFinalizer) EXPECT() *MockCertificateRequestFinalizerMockRecorder {
	return m.recorder
}

// ReconcileCertificateRequest mocks base method
func (m *MockCertificateRequestFinalizer) ReconcileCertificateRequest(obj *v1alpha2.CertificateRequest) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileCertificateRequest", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileCertificateRequest indicates an expected call of ReconcileCertificateRequest
func (mr *MockCertificateRequestFinalizerMockRecorder) ReconcileCertificateRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileCertificateRequest", reflect.TypeOf((*MockCertificateRequestFinalizer)(nil).ReconcileCertificateRequest), obj)
}

// CertificateRequestFinalizerName mocks base method
func (m *MockCertificateRequestFinalizer) CertificateRequestFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificateRequestFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CertificateRequestFinalizerName indicates an expected call of CertificateRequestFinalizerName
func (mr *MockCertificateRequestFinalizerMockRecorder) CertificateRequestFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificateRequestFinalizerName", reflect.TypeOf((*MockCertificateRequestFinalizer)(nil).CertificateRequestFinalizerName))
}

// FinalizeCertificateRequest mocks base method
func (m *MockCertificateRequestFinalizer) FinalizeCertificateRequest(obj *v1alpha2.CertificateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeCertificateRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeCertificateRequest indicates an expected call of FinalizeCertificateRequest
func (mr *MockCertificateRequestFinalizerMockRecorder) FinalizeCertificateRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeCertificateRequest", reflect.TypeOf((*MockCertificateRequestFinalizer)(nil).FinalizeCertificateRequest), obj)
}

// MockCertificateRequestReconcileLoop is a mock of CertificateRequestReconcileLoop interface
type MockCertificateRequestReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateRequestReconcileLoopMockRecorder
}

// MockCertificateRequestReconcileLoopMockRecorder is the mock recorder for MockCertificateRequestReconcileLoop
type MockCertificateRequestReconcileLoopMockRecorder struct {
	mock *MockCertificateRequestReconcileLoop
}

// NewMockCertificateRequestReconcileLoop creates a new mock instance
func NewMockCertificateRequestReconcileLoop(ctrl *gomock.Controller) *MockCertificateRequestReconcileLoop {
	mock := &MockCertificateRequestReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockCertificateRequestReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateRequestReconcileLoop) EXPECT() *MockCertificateRequestReconcileLoopMockRecorder {
	return m.recorder
}

// RunCertificateRequestReconciler mocks base method
func (m *MockCertificateRequestReconcileLoop) RunCertificateRequestReconciler(ctx context.Context, rec controller.CertificateRequestReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCertificateRequestReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCertificateRequestReconciler indicates an expected call of RunCertificateRequestReconciler
func (mr *MockCertificateRequestReconcileLoopMockRecorder) RunCertificateRequestReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCertificateRequestReconciler", reflect.TypeOf((*MockCertificateRequestReconcileLoop)(nil).RunCertificateRequestReconciler), varargs...)
}

// MockPodBounceDirectiveReconciler is a mock of PodBounceDirectiveReconciler interface
type MockPodBounceDirectiveReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPodBounceDirectiveReconcilerMockRecorder
}

// MockPodBounceDirectiveReconcilerMockRecorder is the mock recorder for MockPodBounceDirectiveReconciler
type MockPodBounceDirectiveReconcilerMockRecorder struct {
	mock *MockPodBounceDirectiveReconciler
}

// NewMockPodBounceDirectiveReconciler creates a new mock instance
func NewMockPodBounceDirectiveReconciler(ctrl *gomock.Controller) *MockPodBounceDirectiveReconciler {
	mock := &MockPodBounceDirectiveReconciler{ctrl: ctrl}
	mock.recorder = &MockPodBounceDirectiveReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodBounceDirectiveReconciler) EXPECT() *MockPodBounceDirectiveReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePodBounceDirective mocks base method
func (m *MockPodBounceDirectiveReconciler) ReconcilePodBounceDirective(obj *v1alpha2.PodBounceDirective) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePodBounceDirective", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePodBounceDirective indicates an expected call of ReconcilePodBounceDirective
func (mr *MockPodBounceDirectiveReconcilerMockRecorder) ReconcilePodBounceDirective(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveReconciler)(nil).ReconcilePodBounceDirective), obj)
}

// MockPodBounceDirectiveDeletionReconciler is a mock of PodBounceDirectiveDeletionReconciler interface
type MockPodBounceDirectiveDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockPodBounceDirectiveDeletionReconcilerMockRecorder
}

// MockPodBounceDirectiveDeletionReconcilerMockRecorder is the mock recorder for MockPodBounceDirectiveDeletionReconciler
type MockPodBounceDirectiveDeletionReconcilerMockRecorder struct {
	mock *MockPodBounceDirectiveDeletionReconciler
}

// NewMockPodBounceDirectiveDeletionReconciler creates a new mock instance
func NewMockPodBounceDirectiveDeletionReconciler(ctrl *gomock.Controller) *MockPodBounceDirectiveDeletionReconciler {
	mock := &MockPodBounceDirectiveDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockPodBounceDirectiveDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodBounceDirectiveDeletionReconciler) EXPECT() *MockPodBounceDirectiveDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcilePodBounceDirectiveDeletion mocks base method
func (m *MockPodBounceDirectiveDeletionReconciler) ReconcilePodBounceDirectiveDeletion(req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePodBounceDirectiveDeletion", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcilePodBounceDirectiveDeletion indicates an expected call of ReconcilePodBounceDirectiveDeletion
func (mr *MockPodBounceDirectiveDeletionReconcilerMockRecorder) ReconcilePodBounceDirectiveDeletion(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePodBounceDirectiveDeletion", reflect.TypeOf((*MockPodBounceDirectiveDeletionReconciler)(nil).ReconcilePodBounceDirectiveDeletion), req)
}

// MockPodBounceDirectiveFinalizer is a mock of PodBounceDirectiveFinalizer interface
type MockPodBounceDirectiveFinalizer struct {
	ctrl     *gomock.Controller
	recorder *MockPodBounceDirectiveFinalizerMockRecorder
}

// MockPodBounceDirectiveFinalizerMockRecorder is the mock recorder for MockPodBounceDirectiveFinalizer
type MockPodBounceDirectiveFinalizerMockRecorder struct {
	mock *MockPodBounceDirectiveFinalizer
}

// NewMockPodBounceDirectiveFinalizer creates a new mock instance
func NewMockPodBounceDirectiveFinalizer(ctrl *gomock.Controller) *MockPodBounceDirectiveFinalizer {
	mock := &MockPodBounceDirectiveFinalizer{ctrl: ctrl}
	mock.recorder = &MockPodBounceDirectiveFinalizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodBounceDirectiveFinalizer) EXPECT() *MockPodBounceDirectiveFinalizerMockRecorder {
	return m.recorder
}

// ReconcilePodBounceDirective mocks base method
func (m *MockPodBounceDirectiveFinalizer) ReconcilePodBounceDirective(obj *v1alpha2.PodBounceDirective) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcilePodBounceDirective", obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcilePodBounceDirective indicates an expected call of ReconcilePodBounceDirective
func (mr *MockPodBounceDirectiveFinalizerMockRecorder) ReconcilePodBounceDirective(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcilePodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveFinalizer)(nil).ReconcilePodBounceDirective), obj)
}

// PodBounceDirectiveFinalizerName mocks base method
func (m *MockPodBounceDirectiveFinalizer) PodBounceDirectiveFinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodBounceDirectiveFinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// PodBounceDirectiveFinalizerName indicates an expected call of PodBounceDirectiveFinalizerName
func (mr *MockPodBounceDirectiveFinalizerMockRecorder) PodBounceDirectiveFinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodBounceDirectiveFinalizerName", reflect.TypeOf((*MockPodBounceDirectiveFinalizer)(nil).PodBounceDirectiveFinalizerName))
}

// FinalizePodBounceDirective mocks base method
func (m *MockPodBounceDirectiveFinalizer) FinalizePodBounceDirective(obj *v1alpha2.PodBounceDirective) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizePodBounceDirective", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizePodBounceDirective indicates an expected call of FinalizePodBounceDirective
func (mr *MockPodBounceDirectiveFinalizerMockRecorder) FinalizePodBounceDirective(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizePodBounceDirective", reflect.TypeOf((*MockPodBounceDirectiveFinalizer)(nil).FinalizePodBounceDirective), obj)
}

// MockPodBounceDirectiveReconcileLoop is a mock of PodBounceDirectiveReconcileLoop interface
type MockPodBounceDirectiveReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockPodBounceDirectiveReconcileLoopMockRecorder
}

// MockPodBounceDirectiveReconcileLoopMockRecorder is the mock recorder for MockPodBounceDirectiveReconcileLoop
type MockPodBounceDirectiveReconcileLoopMockRecorder struct {
	mock *MockPodBounceDirectiveReconcileLoop
}

// NewMockPodBounceDirectiveReconcileLoop creates a new mock instance
func NewMockPodBounceDirectiveReconcileLoop(ctrl *gomock.Controller) *MockPodBounceDirectiveReconcileLoop {
	mock := &MockPodBounceDirectiveReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockPodBounceDirectiveReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodBounceDirectiveReconcileLoop) EXPECT() *MockPodBounceDirectiveReconcileLoopMockRecorder {
	return m.recorder
}

// RunPodBounceDirectiveReconciler mocks base method
func (m *MockPodBounceDirectiveReconcileLoop) RunPodBounceDirectiveReconciler(ctx context.Context, rec controller.PodBounceDirectiveReconciler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunPodBounceDirectiveReconciler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunPodBounceDirectiveReconciler indicates an expected call of RunPodBounceDirectiveReconciler
func (mr *MockPodBounceDirectiveReconcileLoopMockRecorder) RunPodBounceDirectiveReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPodBounceDirectiveReconciler", reflect.TypeOf((*MockPodBounceDirectiveReconcileLoop)(nil).RunPodBounceDirectiveReconciler), varargs...)
}
