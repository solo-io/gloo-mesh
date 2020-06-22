// Code generated by MockGen. DO NOT EDIT.
// Source: connect_installation_scanner.go

// Package mock_consul is a generated GoMock package.
package mock_consul

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockConsulConnectInstallationScanner is a mock of ConsulConnectInstallationScanner interface.
type MockConsulConnectInstallationScanner struct {
	ctrl     *gomock.Controller
	recorder *MockConsulConnectInstallationScannerMockRecorder
}

// MockConsulConnectInstallationScannerMockRecorder is the mock recorder for MockConsulConnectInstallationScanner.
type MockConsulConnectInstallationScannerMockRecorder struct {
	mock *MockConsulConnectInstallationScanner
}

// NewMockConsulConnectInstallationScanner creates a new mock instance.
func NewMockConsulConnectInstallationScanner(ctrl *gomock.Controller) *MockConsulConnectInstallationScanner {
	mock := &MockConsulConnectInstallationScanner{ctrl: ctrl}
	mock.recorder = &MockConsulConnectInstallationScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsulConnectInstallationScanner) EXPECT() *MockConsulConnectInstallationScannerMockRecorder {
	return m.recorder
}

// IsConsulConnect mocks base method.
func (m *MockConsulConnectInstallationScanner) IsConsulConnect(arg0 v1.Container) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConsulConnect", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsConsulConnect indicates an expected call of IsConsulConnect.
func (mr *MockConsulConnectInstallationScannerMockRecorder) IsConsulConnect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConsulConnect", reflect.TypeOf((*MockConsulConnectInstallationScanner)(nil).IsConsulConnect), arg0)
}
