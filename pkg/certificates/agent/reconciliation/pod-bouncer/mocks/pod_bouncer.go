// Code generated by MockGen. DO NOT EDIT.
// Source: ./pod_bouncer.go

// Package mock_podbouncer is a generated GoMock package.
package mock_podbouncer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
)

// MockRootCertMatcher is a mock of RootCertMatcher interface.
type MockRootCertMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockRootCertMatcherMockRecorder
}

// MockRootCertMatcherMockRecorder is the mock recorder for MockRootCertMatcher.
type MockRootCertMatcherMockRecorder struct {
	mock *MockRootCertMatcher
}

// NewMockRootCertMatcher creates a new mock instance.
func NewMockRootCertMatcher(ctrl *gomock.Controller) *MockRootCertMatcher {
	mock := &MockRootCertMatcher{ctrl: ctrl}
	mock.recorder = &MockRootCertMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRootCertMatcher) EXPECT() *MockRootCertMatcherMockRecorder {
	return m.recorder
}

// MatchesRootCert mocks base method.
func (m *MockRootCertMatcher) MatchesRootCert(ctx context.Context, rootCert []byte, selector *v1.PodBounceDirectiveSpec_PodSelector, allSecrets v1sets.SecretSet) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchesRootCert", ctx, rootCert, selector, allSecrets)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchesRootCert indicates an expected call of MatchesRootCert.
func (mr *MockRootCertMatcherMockRecorder) MatchesRootCert(ctx, rootCert, selector, allSecrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchesRootCert", reflect.TypeOf((*MockRootCertMatcher)(nil).MatchesRootCert), ctx, rootCert, selector, allSecrets)
}

// MockPodBouncer is a mock of PodBouncer interface.
type MockPodBouncer struct {
	ctrl     *gomock.Controller
	recorder *MockPodBouncerMockRecorder
}

// MockPodBouncerMockRecorder is the mock recorder for MockPodBouncer.
type MockPodBouncerMockRecorder struct {
	mock *MockPodBouncer
}

// NewMockPodBouncer creates a new mock instance.
func NewMockPodBouncer(ctrl *gomock.Controller) *MockPodBouncer {
	mock := &MockPodBouncer{ctrl: ctrl}
	mock.recorder = &MockPodBouncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodBouncer) EXPECT() *MockPodBouncerMockRecorder {
	return m.recorder
}

// BouncePods mocks base method.
func (m *MockPodBouncer) BouncePods(ctx context.Context, podBounceDirective *v1.PodBounceDirective, pods v1sets.PodSet, configMaps v1sets.ConfigMapSet, secrets v1sets.SecretSet) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BouncePods", ctx, podBounceDirective, pods, configMaps, secrets)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BouncePods indicates an expected call of BouncePods.
func (mr *MockPodBouncerMockRecorder) BouncePods(ctx, podBounceDirective, pods, configMaps, secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BouncePods", reflect.TypeOf((*MockPodBouncer)(nil).BouncePods), ctx, podBounceDirective, pods, configMaps, secrets)
}
