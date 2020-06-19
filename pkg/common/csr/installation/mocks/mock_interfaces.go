// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_installation is a generated GoMock package.
package mock_installation

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	installation "github.com/solo-io/service-mesh-hub/pkg/common/csr/installation"
	reflect "reflect"
)

// MockCsrAgentInstaller is a mock of CsrAgentInstaller interface.
type MockCsrAgentInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockCsrAgentInstallerMockRecorder
}

// MockCsrAgentInstallerMockRecorder is the mock recorder for MockCsrAgentInstaller.
type MockCsrAgentInstallerMockRecorder struct {
	mock *MockCsrAgentInstaller
}

// NewMockCsrAgentInstaller creates a new mock instance.
func NewMockCsrAgentInstaller(ctrl *gomock.Controller) *MockCsrAgentInstaller {
	mock := &MockCsrAgentInstaller{ctrl: ctrl}
	mock.recorder = &MockCsrAgentInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCsrAgentInstaller) EXPECT() *MockCsrAgentInstallerMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockCsrAgentInstaller) Install(ctx context.Context, installOptions *installation.CsrAgentInstallOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, installOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockCsrAgentInstallerMockRecorder) Install(ctx, installOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockCsrAgentInstaller)(nil).Install), ctx, installOptions)
}

// Uninstall mocks base method.
func (m *MockCsrAgentInstaller) Uninstall(uninstallOptions *installation.CsrAgentUninstallOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", uninstallOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockCsrAgentInstallerMockRecorder) Uninstall(uninstallOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockCsrAgentInstaller)(nil).Uninstall), uninstallOptions)
}
