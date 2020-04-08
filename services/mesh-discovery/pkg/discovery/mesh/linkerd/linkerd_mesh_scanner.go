package linkerd

import (
	"context"
	"strings"

	linkerdconfig "github.com/linkerd/linkerd2/controller/gen/config"
	"github.com/linkerd/linkerd2/pkg/config"
	linkerdk8s "github.com/linkerd/linkerd2/pkg/k8s"
	v1 "k8s.io/api/core/v1"

	"github.com/google/wire"
	"github.com/rotisserie/eris"
	core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	discoveryv1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	"github.com/solo-io/service-mesh-hub/pkg/env"
	"github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/discovery/mesh"
	k8s_apps_v1 "k8s.io/api/apps/v1"
	k8s_meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	WireProviderSet = wire.NewSet(
		NewLinkerdMeshScanner,
	)
	DiscoveryLabels = map[string]string{
		"discovered_by": "linkerd-mesh-discovery",
	}
	UnexpectedControllerImageName = func(err error, imageName string) error {
		return eris.Wrapf(err, "invalid or unexpected image name format for linkerd controller: %s", imageName)
	}
	LinkerdConfigMapName = linkerdk8s.ConfigConfigMapName
	DefaultClusterDomain = "cluster.local"
)

// disambiguates this MeshScanner from the other MeshScanner implementations so that wire stays happy
type LinkerdMeshScanner mesh.MeshScanner

func NewLinkerdMeshScanner(imageNameParser docker.ImageNameParser) LinkerdMeshScanner {
	return &linkerdMeshScanner{
		imageNameParser: imageNameParser,
	}
}

type linkerdMeshScanner struct {
	imageNameParser docker.ImageNameParser
}

func getLinkerdConfig(ctx context.Context, name, namespace string, kube client.Client) (*linkerdconfig.All, error) {
	cm := &v1.ConfigMap{}
	key := client.ObjectKey{Name: name, Namespace: namespace}
	if err := kube.Get(ctx, key, cm); err != nil {
		return nil, err
	}
	cfg, err := config.FromConfigMap(cm.Data)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (l *linkerdMeshScanner) ScanDeployment(ctx context.Context, deployment *k8s_apps_v1.Deployment, kube client.Client) (*discoveryv1alpha1.Mesh, error) {

	linkerdController, err := l.detectLinkerdController(deployment)

	if err != nil {
		return nil, err
	}

	if linkerdController == nil {
		return nil, nil
	}

	linkerdConfig, err := getLinkerdConfig(ctx, LinkerdConfigMapName, linkerdController.namespace, kube)
	if err != nil {
		return nil, err
	}

	clusterDomain := linkerdConfig.GetGlobal().GetClusterDomain()
	if clusterDomain == "" {
		clusterDomain = DefaultClusterDomain
	}

	return &discoveryv1alpha1.Mesh{
		ObjectMeta: k8s_meta_v1.ObjectMeta{
			Name:      linkerdController.name(),
			Namespace: env.DefaultWriteNamespace,
			Labels:    DiscoveryLabels,
		},
		Spec: discovery_types.MeshSpec{
			MeshType: &discovery_types.MeshSpec_Linkerd{
				Linkerd: &discovery_types.MeshSpec_LinkerdMesh{
					Installation: &discovery_types.MeshSpec_MeshInstallation{
						InstallationNamespace: deployment.GetNamespace(),
						Version:               linkerdController.version,
					},
					ClusterDomain: clusterDomain,
				},
			},
			Cluster: &core_types.ResourceRef{
				Name:      deployment.GetClusterName(),
				Namespace: env.DefaultWriteNamespace,
			},
		},
	}, nil
}

func (l *linkerdMeshScanner) detectLinkerdController(deployment *k8s_apps_v1.Deployment) (*linkerdControllerDeployment, error) {
	var linkerdController *linkerdControllerDeployment

	for _, container := range deployment.Spec.Template.Spec.Containers {
		if strings.Contains(container.Image, "linkerd-io/controller") {
			// TODO there can be > 1 controller image per pod, do we care?
			parsedImage, err := l.imageNameParser.Parse(container.Image)
			if err != nil {
				return nil, UnexpectedControllerImageName(err, container.Image)
			}

			version := parsedImage.Tag
			if parsedImage.Digest != "" {
				version = parsedImage.Digest
			}
			linkerdController = &linkerdControllerDeployment{version: version, namespace: deployment.Namespace, cluster: deployment.ClusterName}
		}
	}

	return linkerdController, nil
}

type linkerdControllerDeployment struct {
	version, namespace, cluster string
}

func (c linkerdControllerDeployment) name() string {
	if c.cluster == "" {
		return "linkerd-" + c.namespace
	}
	// TODO cluster is not restricted to kube name scheme, kebab it
	return "linkerd-" + c.namespace + "-" + c.cluster
}
