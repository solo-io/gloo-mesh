// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock_v1alpha1 is a generated GoMock package.
package mock_v1alpha1

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/xds.enterprise.agent.mesh.gloo.solo.io/v1alpha1"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMulticlusterClientset is a mock of MulticlusterClientset interface
type MockMulticlusterClientset struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterClientsetMockRecorder
}

// MockMulticlusterClientsetMockRecorder is the mock recorder for MockMulticlusterClientset
type MockMulticlusterClientsetMockRecorder struct {
	mock *MockMulticlusterClientset
}

// NewMockMulticlusterClientset creates a new mock instance
func NewMockMulticlusterClientset(ctrl *gomock.Controller) *MockMulticlusterClientset {
	mock := &MockMulticlusterClientset{ctrl: ctrl}
	mock.recorder = &MockMulticlusterClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterClientset) EXPECT() *MockMulticlusterClientsetMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterClientset) Cluster(cluster string) (v1alpha1.Clientset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.Clientset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterClientsetMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterClientset)(nil).Cluster), cluster)
}

// MockClientset is a mock of Clientset interface
type MockClientset struct {
	ctrl     *gomock.Controller
	recorder *MockClientsetMockRecorder
}

// MockClientsetMockRecorder is the mock recorder for MockClientset
type MockClientsetMockRecorder struct {
	mock *MockClientset
}

// NewMockClientset creates a new mock instance
func NewMockClientset(ctrl *gomock.Controller) *MockClientset {
	mock := &MockClientset{ctrl: ctrl}
	mock.recorder = &MockClientsetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientset) EXPECT() *MockClientsetMockRecorder {
	return m.recorder
}

// XdsConfigs mocks base method
func (m *MockClientset) XdsConfigs() v1alpha1.XdsConfigClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XdsConfigs")
	ret0, _ := ret[0].(v1alpha1.XdsConfigClient)
	return ret0
}

// XdsConfigs indicates an expected call of XdsConfigs
func (mr *MockClientsetMockRecorder) XdsConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XdsConfigs", reflect.TypeOf((*MockClientset)(nil).XdsConfigs))
}

// MockXdsConfigReader is a mock of XdsConfigReader interface
type MockXdsConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockXdsConfigReaderMockRecorder
}

// MockXdsConfigReaderMockRecorder is the mock recorder for MockXdsConfigReader
type MockXdsConfigReaderMockRecorder struct {
	mock *MockXdsConfigReader
}

// NewMockXdsConfigReader creates a new mock instance
func NewMockXdsConfigReader(ctrl *gomock.Controller) *MockXdsConfigReader {
	mock := &MockXdsConfigReader{ctrl: ctrl}
	mock.recorder = &MockXdsConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXdsConfigReader) EXPECT() *MockXdsConfigReaderMockRecorder {
	return m.recorder
}

// GetXdsConfig mocks base method
func (m *MockXdsConfigReader) GetXdsConfig(ctx context.Context, key client.ObjectKey) (*v1alpha1.XdsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetXdsConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.XdsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetXdsConfig indicates an expected call of GetXdsConfig
func (mr *MockXdsConfigReaderMockRecorder) GetXdsConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXdsConfig", reflect.TypeOf((*MockXdsConfigReader)(nil).GetXdsConfig), ctx, key)
}

// ListXdsConfig mocks base method
func (m *MockXdsConfigReader) ListXdsConfig(ctx context.Context, opts ...client.ListOption) (*v1alpha1.XdsConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListXdsConfig", varargs...)
	ret0, _ := ret[0].(*v1alpha1.XdsConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListXdsConfig indicates an expected call of ListXdsConfig
func (mr *MockXdsConfigReaderMockRecorder) ListXdsConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListXdsConfig", reflect.TypeOf((*MockXdsConfigReader)(nil).ListXdsConfig), varargs...)
}

// MockXdsConfigWriter is a mock of XdsConfigWriter interface
type MockXdsConfigWriter struct {
	ctrl     *gomock.Controller
	recorder *MockXdsConfigWriterMockRecorder
}

// MockXdsConfigWriterMockRecorder is the mock recorder for MockXdsConfigWriter
type MockXdsConfigWriterMockRecorder struct {
	mock *MockXdsConfigWriter
}

// NewMockXdsConfigWriter creates a new mock instance
func NewMockXdsConfigWriter(ctrl *gomock.Controller) *MockXdsConfigWriter {
	mock := &MockXdsConfigWriter{ctrl: ctrl}
	mock.recorder = &MockXdsConfigWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXdsConfigWriter) EXPECT() *MockXdsConfigWriterMockRecorder {
	return m.recorder
}

// CreateXdsConfig mocks base method
func (m *MockXdsConfigWriter) CreateXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateXdsConfig indicates an expected call of CreateXdsConfig
func (mr *MockXdsConfigWriterMockRecorder) CreateXdsConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateXdsConfig", reflect.TypeOf((*MockXdsConfigWriter)(nil).CreateXdsConfig), varargs...)
}

// DeleteXdsConfig mocks base method
func (m *MockXdsConfigWriter) DeleteXdsConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteXdsConfig indicates an expected call of DeleteXdsConfig
func (mr *MockXdsConfigWriterMockRecorder) DeleteXdsConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteXdsConfig", reflect.TypeOf((*MockXdsConfigWriter)(nil).DeleteXdsConfig), varargs...)
}

// UpdateXdsConfig mocks base method
func (m *MockXdsConfigWriter) UpdateXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateXdsConfig indicates an expected call of UpdateXdsConfig
func (mr *MockXdsConfigWriterMockRecorder) UpdateXdsConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateXdsConfig", reflect.TypeOf((*MockXdsConfigWriter)(nil).UpdateXdsConfig), varargs...)
}

// PatchXdsConfig mocks base method
func (m *MockXdsConfigWriter) PatchXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchXdsConfig indicates an expected call of PatchXdsConfig
func (mr *MockXdsConfigWriterMockRecorder) PatchXdsConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchXdsConfig", reflect.TypeOf((*MockXdsConfigWriter)(nil).PatchXdsConfig), varargs...)
}

// DeleteAllOfXdsConfig mocks base method
func (m *MockXdsConfigWriter) DeleteAllOfXdsConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfXdsConfig indicates an expected call of DeleteAllOfXdsConfig
func (mr *MockXdsConfigWriterMockRecorder) DeleteAllOfXdsConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfXdsConfig", reflect.TypeOf((*MockXdsConfigWriter)(nil).DeleteAllOfXdsConfig), varargs...)
}

// UpsertXdsConfig mocks base method
func (m *MockXdsConfigWriter) UpsertXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, transitionFuncs ...v1alpha1.XdsConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertXdsConfig indicates an expected call of UpsertXdsConfig
func (mr *MockXdsConfigWriterMockRecorder) UpsertXdsConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertXdsConfig", reflect.TypeOf((*MockXdsConfigWriter)(nil).UpsertXdsConfig), varargs...)
}

// MockXdsConfigStatusWriter is a mock of XdsConfigStatusWriter interface
type MockXdsConfigStatusWriter struct {
	ctrl     *gomock.Controller
	recorder *MockXdsConfigStatusWriterMockRecorder
}

// MockXdsConfigStatusWriterMockRecorder is the mock recorder for MockXdsConfigStatusWriter
type MockXdsConfigStatusWriterMockRecorder struct {
	mock *MockXdsConfigStatusWriter
}

// NewMockXdsConfigStatusWriter creates a new mock instance
func NewMockXdsConfigStatusWriter(ctrl *gomock.Controller) *MockXdsConfigStatusWriter {
	mock := &MockXdsConfigStatusWriter{ctrl: ctrl}
	mock.recorder = &MockXdsConfigStatusWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXdsConfigStatusWriter) EXPECT() *MockXdsConfigStatusWriterMockRecorder {
	return m.recorder
}

// UpdateXdsConfigStatus mocks base method
func (m *MockXdsConfigStatusWriter) UpdateXdsConfigStatus(ctx context.Context, obj *v1alpha1.XdsConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateXdsConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateXdsConfigStatus indicates an expected call of UpdateXdsConfigStatus
func (mr *MockXdsConfigStatusWriterMockRecorder) UpdateXdsConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateXdsConfigStatus", reflect.TypeOf((*MockXdsConfigStatusWriter)(nil).UpdateXdsConfigStatus), varargs...)
}

// PatchXdsConfigStatus mocks base method
func (m *MockXdsConfigStatusWriter) PatchXdsConfigStatus(ctx context.Context, obj *v1alpha1.XdsConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchXdsConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchXdsConfigStatus indicates an expected call of PatchXdsConfigStatus
func (mr *MockXdsConfigStatusWriterMockRecorder) PatchXdsConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchXdsConfigStatus", reflect.TypeOf((*MockXdsConfigStatusWriter)(nil).PatchXdsConfigStatus), varargs...)
}

// MockXdsConfigClient is a mock of XdsConfigClient interface
type MockXdsConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockXdsConfigClientMockRecorder
}

// MockXdsConfigClientMockRecorder is the mock recorder for MockXdsConfigClient
type MockXdsConfigClientMockRecorder struct {
	mock *MockXdsConfigClient
}

// NewMockXdsConfigClient creates a new mock instance
func NewMockXdsConfigClient(ctrl *gomock.Controller) *MockXdsConfigClient {
	mock := &MockXdsConfigClient{ctrl: ctrl}
	mock.recorder = &MockXdsConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockXdsConfigClient) EXPECT() *MockXdsConfigClientMockRecorder {
	return m.recorder
}

// GetXdsConfig mocks base method
func (m *MockXdsConfigClient) GetXdsConfig(ctx context.Context, key client.ObjectKey) (*v1alpha1.XdsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetXdsConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.XdsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetXdsConfig indicates an expected call of GetXdsConfig
func (mr *MockXdsConfigClientMockRecorder) GetXdsConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).GetXdsConfig), ctx, key)
}

// ListXdsConfig mocks base method
func (m *MockXdsConfigClient) ListXdsConfig(ctx context.Context, opts ...client.ListOption) (*v1alpha1.XdsConfigList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListXdsConfig", varargs...)
	ret0, _ := ret[0].(*v1alpha1.XdsConfigList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListXdsConfig indicates an expected call of ListXdsConfig
func (mr *MockXdsConfigClientMockRecorder) ListXdsConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).ListXdsConfig), varargs...)
}

// CreateXdsConfig mocks base method
func (m *MockXdsConfigClient) CreateXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, opts ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateXdsConfig indicates an expected call of CreateXdsConfig
func (mr *MockXdsConfigClientMockRecorder) CreateXdsConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).CreateXdsConfig), varargs...)
}

// DeleteXdsConfig mocks base method
func (m *MockXdsConfigClient) DeleteXdsConfig(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteXdsConfig indicates an expected call of DeleteXdsConfig
func (mr *MockXdsConfigClientMockRecorder) DeleteXdsConfig(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).DeleteXdsConfig), varargs...)
}

// UpdateXdsConfig mocks base method
func (m *MockXdsConfigClient) UpdateXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateXdsConfig indicates an expected call of UpdateXdsConfig
func (mr *MockXdsConfigClientMockRecorder) UpdateXdsConfig(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).UpdateXdsConfig), varargs...)
}

// PatchXdsConfig mocks base method
func (m *MockXdsConfigClient) PatchXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchXdsConfig indicates an expected call of PatchXdsConfig
func (mr *MockXdsConfigClientMockRecorder) PatchXdsConfig(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).PatchXdsConfig), varargs...)
}

// DeleteAllOfXdsConfig mocks base method
func (m *MockXdsConfigClient) DeleteAllOfXdsConfig(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfXdsConfig indicates an expected call of DeleteAllOfXdsConfig
func (mr *MockXdsConfigClientMockRecorder) DeleteAllOfXdsConfig(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).DeleteAllOfXdsConfig), varargs...)
}

// UpsertXdsConfig mocks base method
func (m *MockXdsConfigClient) UpsertXdsConfig(ctx context.Context, obj *v1alpha1.XdsConfig, transitionFuncs ...v1alpha1.XdsConfigTransitionFunction) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range transitionFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertXdsConfig", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertXdsConfig indicates an expected call of UpsertXdsConfig
func (mr *MockXdsConfigClientMockRecorder) UpsertXdsConfig(ctx, obj interface{}, transitionFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, transitionFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertXdsConfig", reflect.TypeOf((*MockXdsConfigClient)(nil).UpsertXdsConfig), varargs...)
}

// UpdateXdsConfigStatus mocks base method
func (m *MockXdsConfigClient) UpdateXdsConfigStatus(ctx context.Context, obj *v1alpha1.XdsConfig, opts ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateXdsConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateXdsConfigStatus indicates an expected call of UpdateXdsConfigStatus
func (mr *MockXdsConfigClientMockRecorder) UpdateXdsConfigStatus(ctx, obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateXdsConfigStatus", reflect.TypeOf((*MockXdsConfigClient)(nil).UpdateXdsConfigStatus), varargs...)
}

// PatchXdsConfigStatus mocks base method
func (m *MockXdsConfigClient) PatchXdsConfigStatus(ctx context.Context, obj *v1alpha1.XdsConfig, patch client.Patch, opts ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, obj, patch}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchXdsConfigStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchXdsConfigStatus indicates an expected call of PatchXdsConfigStatus
func (mr *MockXdsConfigClientMockRecorder) PatchXdsConfigStatus(ctx, obj, patch interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, obj, patch}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchXdsConfigStatus", reflect.TypeOf((*MockXdsConfigClient)(nil).PatchXdsConfigStatus), varargs...)
}

// MockMulticlusterXdsConfigClient is a mock of MulticlusterXdsConfigClient interface
type MockMulticlusterXdsConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterXdsConfigClientMockRecorder
}

// MockMulticlusterXdsConfigClientMockRecorder is the mock recorder for MockMulticlusterXdsConfigClient
type MockMulticlusterXdsConfigClientMockRecorder struct {
	mock *MockMulticlusterXdsConfigClient
}

// NewMockMulticlusterXdsConfigClient creates a new mock instance
func NewMockMulticlusterXdsConfigClient(ctrl *gomock.Controller) *MockMulticlusterXdsConfigClient {
	mock := &MockMulticlusterXdsConfigClient{ctrl: ctrl}
	mock.recorder = &MockMulticlusterXdsConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterXdsConfigClient) EXPECT() *MockMulticlusterXdsConfigClientMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockMulticlusterXdsConfigClient) Cluster(cluster string) (v1alpha1.XdsConfigClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster", cluster)
	ret0, _ := ret[0].(v1alpha1.XdsConfigClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cluster indicates an expected call of Cluster
func (mr *MockMulticlusterXdsConfigClientMockRecorder) Cluster(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockMulticlusterXdsConfigClient)(nil).Cluster), cluster)
}
