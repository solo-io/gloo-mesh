package extensions

import (
	"context"
	"sync"
	"time"

	"github.com/solo-io/go-utils/grpcutils"
	"github.com/solo-io/go-utils/hashutils"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/extensions/v1beta1"
	settingsv1 "github.com/solo-io/gloo-mesh/pkg/api/settings.mesh.gloo.solo.io/v1"
	"github.com/solo-io/go-utils/contextutils"
)

//go:generate mockgen -source ./clients.go -destination mocks/mock_clients.go

// A PushFunc handles push notifications from an Extensions Server
type PushFunc func(notification *v1beta1.PushNotification)

// Clients provides a convenience wrapper for a set of clients to communicate with multiple Extension Servers
type Clients []v1beta1.NetworkingExtensionsClient

func NewClientsFromSettings(ctx context.Context, extensionsServerOptions []*settingsv1.GrpcServer) (Clients, error) {
	var extensionsClients Clients
	for _, extensionsServer := range extensionsServerOptions {
		extensionsServerAddr := extensionsServer.GetAddress()
		if extensionsServerAddr == "" {
			return nil, eris.Errorf("must specify extensions server address")
		}
		dialOpts := grpcutils.DialOpts{
			Address:                    extensionsServerAddr,
			Insecure:                   extensionsServer.GetInsecure(),
			ReconnectOnNetworkFailures: extensionsServer.GetReconnectOnNetworkFailures(),
		}
		grpcConnection, err := dialOpts.Dial(ctx)
		if err != nil {
			return nil, err
		}
		extensionsClients = append(extensionsClients, v1beta1.NewNetworkingExtensionsClient(grpcConnection))
	}
	return extensionsClients, nil
}

// WatchPushNotifications watches push notifications from the available extension servers until the context is cancelled.
// Will call pushFn() when a notification is received.
func (c Clients) WatchPushNotifications(ctx context.Context, pushFn PushFunc) error {
	for _, exClient := range c {
		exClient := exClient // shadow for goroutine
		go handlePushesForever(ctx, exClient, pushFn)
	}
	return nil
}

// handles push notifications for an individual connection stream
func handlePushesForever(ctx context.Context, exClient v1beta1.NetworkingExtensionsClient, pushFn PushFunc) {
	for {
		if err := startNotificationWatch(ctx, exClient, pushFn); err != nil {
			contextutils.LoggerFrom(ctx).Errorf("failed to start push notification watch with client %+v", exClient)
			// retry after 5 sec
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second * 5):
			}
		}
	}
}

func startNotificationWatch(ctx context.Context, exClient v1beta1.NetworkingExtensionsClient, pushFn PushFunc) (err error) {
	notifications, err := exClient.WatchPushNotifications(ctx, &v1beta1.WatchPushNotificationsRequest{})
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		notification, err := notifications.Recv()
		if err != nil {
			return err
		}
		pushFn(notification)
	}
}

// Clientset provides a handle to fetching a set of cached gRPC clients
type Clientset interface {
	// ConfigureServers updates the set of servers this Clientset is configured with.
	// Restarts the notification watches if servers were updated, using the new pushFn to handle notification pushes.
	ConfigureServers(extensionsServerOptions []*settingsv1.GrpcServer, pushFn PushFunc) error

	// GetClients returns the set of Extension clients that are cached with this Clientset.
	// Must be called after UpdateServers
	GetClients() Clients
}

type clientset struct {
	rootCtx       context.Context
	cachedClients cachedClients
}

func NewClientset(rootCtx context.Context) *clientset {
	return &clientset{rootCtx: rootCtx}
}

type cachedClients struct {
	lock        sync.RWMutex
	ctx         context.Context
	cancel      context.CancelFunc
	optionsHash uint64 // this hash used to keep track of changes to the server options
	clients     Clients
}

func (c *clientset) ConfigureServers(extensionsServerOptions []*settingsv1.GrpcServer, pushFn PushFunc) error {
	optionsHash, err := hashutils.HashAllSafe(nil, extensionsServerOptions)
	if err != nil {
		return err
	}
	if c.cachedClients.optionsHash == optionsHash {
		// nothing to do, options have remained the same
		return nil
	}

	newContext, newCancel := context.WithCancel(c.rootCtx)
	newClients, err := NewClientsFromSettings(newContext, extensionsServerOptions)
	if err != nil {
		return err
	}

	c.cachedClients.lock.Lock()

	// cancel previous clients
	if c.cachedClients.cancel != nil {
		c.cachedClients.cancel()
	}
	// update latest clients
	c.cachedClients.clients = newClients
	c.cachedClients.ctx = newContext
	c.cachedClients.cancel = newCancel
	c.cachedClients.optionsHash = optionsHash

	c.cachedClients.lock.Unlock()

	return c.watchPushNotifications(pushFn)
}

func (c *clientset) GetClients() Clients {
	c.cachedClients.lock.RLock()
	clients := c.cachedClients.clients
	c.cachedClients.lock.RUnlock()

	return clients
}

func (c *clientset) watchPushNotifications(pushFn PushFunc) error {
	return c.GetClients().WatchPushNotifications(c.cachedClients.ctx, pushFn)
}
