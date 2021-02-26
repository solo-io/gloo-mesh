package extensions

import (
	"context"
	"fmt"
	"net"

	extensionutils "github.com/solo-io/gloo-mesh/pkg/mesh-networking/extensions"

	"go.uber.org/atomic"

	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/extensions"
	"google.golang.org/grpc"

	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/extensions/v1beta1"
)

const ExtensionsServerPort = 2345

type testExtensionsServer struct {
	createMeshPatches func(ctx context.Context, mesh *v1.MeshSpec) (istio.Builder, error)
	hasConnected      *atomic.Bool
}

func NewTestExtensionsServer() *testExtensionsServer {
	return &testExtensionsServer{createMeshPatches: getCreateMeshPatchesFunc(), hasConnected: &atomic.Bool{}}
}

// Runs an e2e implementation of a grpc extensions service for Networking
// that adds a route to a simple "HelloWorld" server running on the local machine (reachable via `host.docker.internal` from inside KinD)
func (t *testExtensionsServer) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", ExtensionsServerPort))
	if err != nil {
		return err
	}
	grpcSrv := grpc.NewServer()
	v1beta1.RegisterNetworkingExtensionsServer(grpcSrv, t)
	return grpcSrv.Serve(l)
}

func (t *testExtensionsServer) GetExtensionPatches(ctx context.Context, request *v1beta1.ExtensionPatchRequest) (*v1beta1.ExtensionPatchResponse, error) {
	inputs := extensionutils.InputSnapshotFromProto("test-server", request.Inputs)

	var patches []*v1beta1.GeneratedObject
	for _, mesh := range inputs.Meshes().List() {
		mesh := mesh // shadow for pointer
		outputs, err := t.createMeshPatches(ctx, &mesh.Spec)
		if err != nil {
			return nil, err
		}
		patches = append(patches, extensions.OutputsToProto(outputs)...)
	}
	return &v1beta1.ExtensionPatchResponse{PatchedOutputs: patches}, nil
}

func (t *testExtensionsServer) WatchPushNotifications(request *v1beta1.WatchPushNotificationsRequest, server v1beta1.NetworkingExtensions_WatchPushNotificationsServer) error {

	// one to start
	if err := server.Send(&v1beta1.PushNotification{}); err != nil {
		return err
	}

	// client has connected
	t.hasConnected.Store(true)

	// sleep forever
	select {
	case <-server.Context().Done():
		return nil
	}
}

// returns true if a client has connected to this server
func (t *testExtensionsServer) HasConnected() bool {
	return t.hasConnected.Load()
}
