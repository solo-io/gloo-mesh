// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	context "context"
	reflect "reflect"

	appmesh "github.com/aws/aws-sdk-go/service/appmesh"
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
)

// MockAppmeshClient is a mock of AppmeshClient interface.
type MockAppmeshClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppmeshClientMockRecorder
}

// MockAppmeshClientMockRecorder is the mock recorder for MockAppmeshClient.
type MockAppmeshClientMockRecorder struct {
	mock *MockAppmeshClient
}

// NewMockAppmeshClient creates a new mock instance.
func NewMockAppmeshClient(ctrl *gomock.Controller) *MockAppmeshClient {
	mock := &MockAppmeshClient{ctrl: ctrl}
	mock.recorder = &MockAppmeshClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppmeshClient) EXPECT() *MockAppmeshClientMockRecorder {
	return m.recorder
}

// ListMeshes mocks base method.
func (m *MockAppmeshClient) ListMeshes(input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMeshes", input)
	ret0, _ := ret[0].(*appmesh.ListMeshesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMeshes indicates an expected call of ListMeshes.
func (mr *MockAppmeshClientMockRecorder) ListMeshes(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMeshes", reflect.TypeOf((*MockAppmeshClient)(nil).ListMeshes), input)
}

// ListTagsForResource mocks base method.
func (m *MockAppmeshClient) ListTagsForResource(arg0 *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*appmesh.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppmeshClientMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppmeshClient)(nil).ListTagsForResource), arg0)
}

// EnsureVirtualService mocks base method.
func (m *MockAppmeshClient) EnsureVirtualService(virtualServiceData *appmesh.VirtualServiceData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVirtualService", virtualServiceData)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVirtualService indicates an expected call of EnsureVirtualService.
func (mr *MockAppmeshClientMockRecorder) EnsureVirtualService(virtualServiceData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVirtualService", reflect.TypeOf((*MockAppmeshClient)(nil).EnsureVirtualService), virtualServiceData)
}

// EnsureVirtualRouter mocks base method.
func (m *MockAppmeshClient) EnsureVirtualRouter(virtualRouter *appmesh.VirtualRouterData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVirtualRouter", virtualRouter)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVirtualRouter indicates an expected call of EnsureVirtualRouter.
func (mr *MockAppmeshClientMockRecorder) EnsureVirtualRouter(virtualRouter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVirtualRouter", reflect.TypeOf((*MockAppmeshClient)(nil).EnsureVirtualRouter), virtualRouter)
}

// EnsureRoute mocks base method.
func (m *MockAppmeshClient) EnsureRoute(route *appmesh.RouteData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureRoute", route)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureRoute indicates an expected call of EnsureRoute.
func (mr *MockAppmeshClientMockRecorder) EnsureRoute(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureRoute", reflect.TypeOf((*MockAppmeshClient)(nil).EnsureRoute), route)
}

// EnsureVirtualNode mocks base method.
func (m *MockAppmeshClient) EnsureVirtualNode(virtualNode *appmesh.VirtualNodeData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVirtualNode", virtualNode)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVirtualNode indicates an expected call of EnsureVirtualNode.
func (mr *MockAppmeshClientMockRecorder) EnsureVirtualNode(virtualNode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVirtualNode", reflect.TypeOf((*MockAppmeshClient)(nil).EnsureVirtualNode), virtualNode)
}

// ReconcileVirtualRoutersAndRoutesAndVirtualServices mocks base method.
func (m *MockAppmeshClient) ReconcileVirtualRoutersAndRoutesAndVirtualServices(ctx context.Context, meshName *string, virtualRouters []*appmesh.VirtualRouterData, routes []*appmesh.RouteData, virtualServices []*appmesh.VirtualServiceData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualRoutersAndRoutesAndVirtualServices", ctx, meshName, virtualRouters, routes, virtualServices)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualRoutersAndRoutesAndVirtualServices indicates an expected call of ReconcileVirtualRoutersAndRoutesAndVirtualServices.
func (mr *MockAppmeshClientMockRecorder) ReconcileVirtualRoutersAndRoutesAndVirtualServices(ctx, meshName, virtualRouters, routes, virtualServices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualRoutersAndRoutesAndVirtualServices", reflect.TypeOf((*MockAppmeshClient)(nil).ReconcileVirtualRoutersAndRoutesAndVirtualServices), ctx, meshName, virtualRouters, routes, virtualServices)
}

// ReconcileVirtualNodes mocks base method.
func (m *MockAppmeshClient) ReconcileVirtualNodes(ctx context.Context, meshName *string, virtualNodes []*appmesh.VirtualNodeData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileVirtualNodes", ctx, meshName, virtualNodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileVirtualNodes indicates an expected call of ReconcileVirtualNodes.
func (mr *MockAppmeshClientMockRecorder) ReconcileVirtualNodes(ctx, meshName, virtualNodes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVirtualNodes", reflect.TypeOf((*MockAppmeshClient)(nil).ReconcileVirtualNodes), ctx, meshName, virtualNodes)
}

// MockSTSClient is a mock of STSClient interface.
type MockSTSClient struct {
	ctrl     *gomock.Controller
	recorder *MockSTSClientMockRecorder
}

// MockSTSClientMockRecorder is the mock recorder for MockSTSClient.
type MockSTSClientMockRecorder struct {
	mock *MockSTSClient
}

// NewMockSTSClient creates a new mock instance.
func NewMockSTSClient(ctrl *gomock.Controller) *MockSTSClient {
	mock := &MockSTSClient{ctrl: ctrl}
	mock.recorder = &MockSTSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSTSClient) EXPECT() *MockSTSClientMockRecorder {
	return m.recorder
}

// GetCallerIdentity mocks base method.
func (m *MockSTSClient) GetCallerIdentity() (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity")
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockSTSClientMockRecorder) GetCallerIdentity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockSTSClient)(nil).GetCallerIdentity))
}
