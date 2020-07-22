package groups

import (
	"github.com/solo-io/service-mesh-hub/codegen/constants"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	smhModule       = "github.com/solo-io/service-mesh-hub"
	v1alpha2Version = "v1alpha2"
	apiRoot         = "pkg/api"
)

var SMHGroups = []model.Group{
	//TODO(ilackarms): we may need to restore settings at some point, so this is left commented
	//makeGroup("core", v1alpha2Version, []resourceToGenerate{
	//	{kind: "Settings", noStatus: true},
	//}),
	makeGroup("discovery", v1alpha2Version, []resourceToGenerate{
		{kind: "MeshService"},
		{kind: "MeshWorkload"},
		{kind: "Mesh"},
	}),
	makeGroup("networking", v1alpha2Version, []resourceToGenerate{
		{kind: "TrafficPolicy"},
		{kind: "AccessPolicy"},
		{kind: "VirtualMesh"},
		{kind: "FailoverService"},
	}),
}

var CSRGroups = []model.Group{
	makeGroup("security", v1alpha2Version, []resourceToGenerate{
		{kind: "VirtualMeshCertificateSigningRequest"},
	}),
}

type resourceToGenerate struct {
	kind     string
	noStatus bool // don't put a status on this resource
}

func makeGroup(groupPrefix, version string, resourcesToGenerate []resourceToGenerate) model.Group {
	var resources []model.Resource
	for _, resource := range resourcesToGenerate {
		res := model.Resource{
			Kind: resource.kind,
			Spec: model.Field{
				Type: model.Type{
					Name: resource.kind + "Spec",
				},
			},
		}
		if !resource.noStatus {
			res.Status = &model.Field{Type: model.Type{
				Name: resource.kind + "Status",
			}}
		}
		resources = append(resources, res)
	}

	return model.Group{
		GroupVersion: schema.GroupVersion{
			Group:   groupPrefix + "." + constants.ServiceMeshHubApiGroupSuffix,
			Version: version,
		},
		Module:           smhModule,
		Resources:        resources,
		RenderManifests:  true,
		RenderTypes:      true,
		RenderClients:    true,
		RenderController: true,
		MockgenDirective: true,
		CustomTemplates:  contrib.AllGroupCustomTemplates,
		ApiRoot:          apiRoot,
	}
}
