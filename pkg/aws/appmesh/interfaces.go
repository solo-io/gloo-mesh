package appmesh

import (
	"context"

	appmesh2 "github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/sts"
)

//go:generate mockgen -source ./interfaces.go -destination mocks/mock_interfaces.go

type AppmeshMatcher interface {
	AreRoutesEqual(
		routeA *appmesh2.RouteData,
		routeB *appmesh2.RouteData,
	) bool

	AreVirtualNodesEqual(
		virtualNodeA *appmesh2.VirtualNodeData,
		virtualNodeB *appmesh2.VirtualNodeData,
	) bool

	AreVirtualServicesEqual(
		virtualServiceA *appmesh2.VirtualServiceData,
		virtualServiceB *appmesh2.VirtualServiceData,
	) bool

	AreVirtualRoutersEqual(
		virtualRouterA *appmesh2.VirtualRouterData,
		virtualRouterB *appmesh2.VirtualRouterData,
	) bool
}

/*
	Provide methods that ensure the existence of the given Appmesh resource.
	"Ensure" means create if the resource doesn't exist as identified by its name.
	If it does exist, update it if its spec doesn't match the provided resource spec. Otherwise do nothing.
*/
type AppmeshClient interface {
	EnsureVirtualService(virtualServiceData *appmesh2.VirtualServiceData) error
	EnsureVirtualRouter(virtualRouter *appmesh2.VirtualRouterData) error
	EnsureRoute(route *appmesh2.RouteData) error
	EnsureVirtualNode(virtualNode *appmesh2.VirtualNodeData) error

	ReconcileVirtualServices(
		ctx context.Context,
		meshName *string,
		virtualServices []*appmesh2.VirtualServiceData,
	) error
	ReconcileVirtualRouters(
		ctx context.Context,
		meshName *string,
		virtualRouters []*appmesh2.VirtualRouterData,
	) error
	ReconcileRoutes(
		ctx context.Context,
		meshName *string,
		routes []*appmesh2.RouteData,
	) error
	ReconcileVirtualNodes(
		ctx context.Context,
		meshName *string,
		virtualNodes []*appmesh2.VirtualNodeData,
	) error
}

type STSClient interface {
	// Retrieves caller identity metadata by making a request to AWS STS (Secure Token Service).
	GetCallerIdentity() (*sts.GetCallerIdentityOutput, error)
}
