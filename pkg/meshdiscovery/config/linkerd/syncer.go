package linkerd

import (
	"context"
	"fmt"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/eventloop"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/clientset"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/mesh/linkerd"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/utils"
	"github.com/solo-io/supergloo/pkg/registration"
	"go.uber.org/zap"
)

const (
	injectionAnnotation = "linkerd.io/inject"
	enabled             = "enabled"
	disabled            = "disabled"

	proxyContainer = "linkerd-proxy"
)

func StartLinkerdDiscoveryConfigLoop(ctx context.Context, cs *clientset.Clientset, pubSub *registration.PubSub) {
	sgConfigLoop := &linkerdDiscoveryConfigLoop{cs: cs}
	sgListener := registration.NewSubscriber(ctx, pubSub, sgConfigLoop)
	sgListener.Listen(ctx)
}

type linkerdDiscoveryConfigLoop struct {
	cs *clientset.Clientset
}

func (cl *linkerdDiscoveryConfigLoop) Enabled(enabled registration.EnabledConfigLoops) bool {
	return enabled.Linkerd
}

func (cl *linkerdDiscoveryConfigLoop) Start(ctx context.Context, enabled registration.EnabledConfigLoops) (eventloop.EventLoop, error) {
	emitter := v1.NewLinkerdDiscoveryEmitter(
		cl.cs.Discovery.Mesh,
		cl.cs.Input.Install,
		cl.cs.Input.Pod,
		cl.cs.Input.Upstream,
	)
	syncer := newLinkerdConfigDiscoverSyncer(cl.cs)
	el := v1.NewLinkerdDiscoveryEventLoop(emitter, syncer)

	return el, nil
}

type linkerdConfigDiscoverSyncer struct {
	cs *clientset.Clientset
}

func newLinkerdConfigDiscoverSyncer(cs *clientset.Clientset) *linkerdConfigDiscoverSyncer {
	return &linkerdConfigDiscoverSyncer{cs: cs}
}

func (lcds *linkerdConfigDiscoverSyncer) Sync(ctx context.Context, snap *v1.LinkerdDiscoverySnapshot) error {
	ctx = contextutils.WithLogger(ctx, fmt.Sprintf("linkerd-config-discovery-sync-%v", snap.Hash()))
	logger := contextutils.LoggerFrom(ctx)
	fields := []interface{}{
		zap.Int("meshes", len(snap.Meshes.List())),
		zap.Int("installs", len(snap.Installs.List())),
		zap.Int("pods", len(snap.Pods.List())),
		zap.Int("upstreams", len(snap.Upstreams.List())),
	}

	logger.Infow("begin sync", fields...)
	defer logger.Infow("end sync", fields...)
	logger.Debugf("full snapshot: %v", snap)

	linkerdMeshes := utils.GetMeshes(snap.Meshes.List(), utils.LinkerdMeshFilterFunc, utils.FilterByLabels(linkerd.DiscoverySelector))
	linkerdInstalls := utils.GetActiveInstalls(snap.Installs.List(), utils.LinkerdInstallFilterFunc)
	injectedPods := utils.InjectedPodsByNamespace(snap.Pods.List(), proxyContainer)

	meshResources := organizeMeshes(
		linkerdMeshes,
		linkerdInstalls,
		injectedPods,
		snap.Upstreams.List(),
	)

	var updatedMeshes v1.MeshList
	for _, fullMesh := range meshResources {
		updatedMeshes = append(updatedMeshes, fullMesh.merge())
	}

	meshReconciler := v1.NewMeshReconciler(lcds.cs.Discovery.Mesh)
	listOpts := clients.ListOpts{
		Ctx:      ctx,
		Selector: linkerd.DiscoverySelector,
	}

	return meshReconciler.Reconcile("", updatedMeshes, nil, listOpts)
}

func organizeMeshes(meshes v1.MeshList, installs v1.InstallList, injectedPods kubernetes.PodsByNamespace,
	upstreams gloov1.UpstreamList) meshResourceList {
	result := make(meshResourceList, len(meshes))

	for i, mesh := range meshes {
		linkerdMesh := mesh.GetLinkerd()
		if linkerdMesh == nil {
			continue
		}
		fullMesh := &meshResources{
			Mesh: mesh,
		}
		for _, install := range installs {
			if install.InstallationNamespace == linkerdMesh.InstallationNamespace {
				fullMesh.Install = install
				break
			}
		}

		// Currently injection is a constant so there's no way to distinguish between
		// multiple istio deployments in a single cluster
		fullMesh.Upstreams = utils.GetUpstreamsForInjectedPods(injectedPods.List(), upstreams)

		result[i] = fullMesh
	}
	return result
}

type meshResourceList []*meshResources
type meshResources struct {
	Install   *v1.Install
	Mesh      *v1.Mesh
	Upstreams gloov1.UpstreamList
}

// Main merge method for discovered info
// Priority of data is as such Install > Mesh
func (fm *meshResources) merge() *v1.Mesh {
	result := fm.Mesh
	linkerdMesh := fm.Mesh.GetLinkerd()
	if linkerdMesh == nil {
		return fm.Mesh
	}
	mtlsConfig := &v1.MtlsConfig{}
	if result.DiscoveryMetadata == nil {
		result.DiscoveryMetadata = &v1.DiscoveryMetadata{}
	}

	var meshUpstreams []*core.ResourceRef
	for _, upstream := range fm.Upstreams {
		ref := upstream.Metadata.Ref()
		meshUpstreams = append(meshUpstreams, &ref)
	}
	result.DiscoveryMetadata.Upstreams = meshUpstreams

	result.DiscoveryMetadata.InjectedNamespaceLabel = injectionAnnotation

	if fm.Install != nil {
		mesh := fm.Install.GetMesh()
		if mesh != nil {
			linkerdMeshInstall := mesh.GetLinkerd()
			result.DiscoveryMetadata.MeshVersion = linkerdMeshInstall.GetVersion()
			result.DiscoveryMetadata.EnableAutoInject = linkerdMeshInstall.GetEnableAutoInject()
			mtlsConfig.MtlsEnabled = linkerdMeshInstall.GetEnableMtls()
		}
		result.DiscoveryMetadata.InstallationNamespace = fm.Install.InstallationNamespace
	}
	result.DiscoveryMetadata.MtlsConfig = mtlsConfig
	return result
}
