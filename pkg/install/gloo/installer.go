package gloo

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/supergloo/pkg/install/utils/helm"

	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

type Installer interface {
	EnsureGlooInstall(ctx context.Context, install *v1.Install, meshes v1.MeshList, meshIngresses v1.MeshIngressList) (*v1.MeshIngress, error)
}

type defaultInstaller struct {
	helmInstaller helm.Installer
}

func NewDefaultInstaller(helmInstaller helm.Installer) *defaultInstaller {
	return &defaultInstaller{helmInstaller: helmInstaller}
}

func (installer *defaultInstaller) EnsureGlooInstall(ctx context.Context, install *v1.Install, meshes v1.MeshList, meshIngresses v1.MeshIngressList) (*v1.MeshIngress, error) {
	ctx = contextutils.WithLogger(ctx, "gloo-ingress-installer")
	logger := contextutils.LoggerFrom(ctx)

	installIngress, ok := install.InstallType.(*v1.Install_Ingress)
	if !ok {
		return nil, errors.Errorf("non ingress install detected in ingress install, %v", install.Metadata.Ref())
	}

	logger.Infof("beginning gloo install sync %v", installIngress)

	glooInstall, ok := installIngress.Ingress.IngressInstallType.(*v1.MeshIngressInstall_Gloo)
	if !ok {
		return nil, errors.Errorf("%v: invalid install type, only gloo ingress supported currently", install.Metadata.Ref())
	}

	var previousInstall helm.Manifests
	if install.InstalledManifest != "" {
		logger.Infof("detected previous install of gloo ingress")
		manifests, err := helm.NewManifestsFromGzippedString(install.InstalledManifest)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing previously installed manifest")
		}
		previousInstall = manifests
	}

	installNamespace := install.InstallationNamespace

	if install.Disabled {
		if len(previousInstall) > 0 {
			logger.Infof("deleting previous gloo ingress install")
			if err := installer.helmInstaller.DeleteFromManifests(ctx, installNamespace, previousInstall); err != nil {
				return nil, errors.Wrapf(err, "uninstalling gloo ingress")
			}
			install.InstalledManifest = ""
			installIngress.Ingress.InstalledIngress = nil
		}
		return nil, nil
	}

	opts := NewInstallOptions(previousInstall, installer.helmInstaller, installNamespace, glooInstall.Gloo.GlooVersion)

	logger.Infof("installing gloo-ingress with options: %#v", opts)

	manifests, err := helm.InstallOrUpdate(ctx, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "installing gloo ingress")
	}

	gzipped, err := manifests.Gzipped()
	if err != nil {
		return nil, errors.Wrapf(err, "converting installed manifests to gzipped string")
	}

	var meshRefs []*core.ResourceRef
	for _, glooMesh := range glooInstall.Gloo.Meshes {
		mesh, err := meshes.Find(glooMesh.Namespace, glooMesh.Name)
		if err == nil && mesh != nil {
			ref := mesh.Metadata.Ref()
			meshRefs = append(meshRefs, &ref)
		}
	}

	var meshIngress *v1.MeshIngress
	if installIngress.Ingress.InstalledIngress != nil {
		var err error
		meshIngress, err = meshIngresses.Find(installIngress.Ingress.InstalledIngress.Strings())
		if err != nil {
			return nil, errors.Wrapf(err, "installed ingress not found")
		}
	}

	if meshIngress != nil {
		meshIngress.Meshes = meshRefs
		meshIngress.MeshIngressType = &v1.MeshIngress_Gloo{
			Gloo: &v1.GlooMeshIngress{
				InstallationNamespace: install.InstallationNamespace,
			},
		}
	} else {

		meshIngress = &v1.MeshIngress{
			Metadata: core.Metadata{
				Namespace: install.Metadata.Namespace,
				Name:      install.Metadata.Name,
			},
			MeshIngressType: &v1.MeshIngress_Gloo{
				Gloo: &v1.GlooMeshIngress{
					InstallationNamespace: install.InstallationNamespace,
				},
			},
			Meshes: meshRefs,
		}
	}

	// caller should expect the install to have been modified
	install.InstalledManifest = gzipped
	ref := meshIngress.Metadata.Ref()
	installIngress.Ingress.InstalledIngress = &ref

	return meshIngress, nil
}
