package linkerd2

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	prometheusv1 "github.com/solo-io/supergloo/pkg2/api/external/prometheus/v1"
	v1 "github.com/solo-io/supergloo/pkg2/api/v1"
	"github.com/solo-io/supergloo/pkg2/translator/shared"
	"k8s.io/client-go/kubernetes"
)

func NewPrometheusSyncer(kube kubernetes.Interface, prometheusClient prometheusv1.PrometheusConfigClient) v1.TranslatorSyncer {
	return &shared.PrometheusSyncer{
		Kube:                 kube,
		PrometheusClient:     prometheusClient,
		DesiredScrapeConfigs: LinkerdScrapeConfigs,
		GetConfigMap: func(mesh *v1.Mesh) *core.ResourceRef {
			linkerdMesh, ok := mesh.MeshType.(*v1.Mesh_Linkerd2)
			if !ok {
				// not our mesh, we don't care
				return nil
			}
			return linkerdMesh.Linkerd2.PrometheusConfigmap
		},
	}
}
