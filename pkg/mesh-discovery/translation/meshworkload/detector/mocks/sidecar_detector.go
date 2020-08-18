// Code generated by MockGen. DO NOT EDIT.
// Source: ./sidecar_detector.go

// Package mock_detector is a generated GoMock package.
package mock_detector

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	v1alpha2sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2/sets"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockSidecarDetector is a mock of SidecarDetector interface
type MockSidecarDetector struct {
	ctrl     *gomock.Controller
	recorder *MockSidecarDetectorMockRecorder
}

// MockSidecarDetectorMockRecorder is the mock recorder for MockSidecarDetector
type MockSidecarDetectorMockRecorder struct {
	mock *MockSidecarDetector
}

// NewMockSidecarDetector creates a new mock instance
func NewMockSidecarDetector(ctrl *gomock.Controller) *MockSidecarDetector {
	mock := &MockSidecarDetector{ctrl: ctrl}
	mock.recorder = &MockSidecarDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSidecarDetector) EXPECT() *MockSidecarDetectorMockRecorder {
	return m.recorder
}

// DetectMeshSidecar mocks base method
func (m *MockSidecarDetector) DetectMeshSidecar(pod *v1.Pod, meshes v1alpha2sets.MeshSet) *v1alpha2.Mesh {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectMeshSidecar", pod, meshes)
	ret0, _ := ret[0].(*v1alpha2.Mesh)
	return ret0
}

// DetectMeshSidecar indicates an expected call of DetectMeshSidecar
func (mr *MockSidecarDetectorMockRecorder) DetectMeshSidecar(pod, meshes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectMeshSidecar", reflect.TypeOf((*MockSidecarDetector)(nil).DetectMeshSidecar), pod, meshes)
}
