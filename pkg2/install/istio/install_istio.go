package istio

import (
	"bytes"
	"context"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/solo-io/supergloo/pkg2/install/helm"
)

const (
	IstioVersion103      = "1.0.3"
	IstioVersion103Chart = "https://s3.amazonaws.com/supergloo.solo.io/istio-1.0.3.tgz"

	IstioVersion105      = "1.0.5"
	IstioVersion105Chart = "https://s3.amazonaws.com/supergloo.solo.io/istio-1.0.5.tgz"
)

var supportedIstioVersions = map[string]versionedInstall{
	IstioVersion103: {
		chartPath:      IstioVersion103Chart,
		valuesTemplate: helmValues,
	},
	IstioVersion105: {
		chartPath:      IstioVersion105Chart,
		valuesTemplate: helmValues,
	},
}

type versionedInstall struct {
	chartPath      string
	valuesTemplate string
}

type InstallOptions struct {
	Version       string
	Namespace     string
	AutoInject    AutoInjectInstallOptions
	Mtls          MtlsInstallOptions
	Observability ObservabilityInstallOptions
	Gateway       GatewayInstallOptions
}

func (o InstallOptions) Validate() error {
	if o.Version == "" {
		return errors.Errorf("must provide istio install version")
	}
	if o.Namespace == "" {
		return errors.Errorf("must provide istio install namespace")
	}
	if o.Observability.EnableServiceGraph && !o.Observability.EnablePrometheus {
		return errors.Errorf("servicegraph can only be enabled with prometheus")
	}
	return nil
}

type AutoInjectInstallOptions struct {
	Enabled bool
}

type MtlsInstallOptions struct {
	Enabled        bool
	SelfSignedCert bool
}

type ObservabilityInstallOptions struct {
	EnableGrafana      bool
	EnablePrometheus   bool
	EnableJaeger       bool
	EnableServiceGraph bool
}

type GatewayInstallOptions struct {
	EnableIngress bool
	EnableEgress  bool
}

func releaseName(namespace, version string) string {
	return "supergloo-" + namespace + version
}

func InstallIstio(ctx context.Context, opts InstallOptions) error {
	if err := opts.Validate(); err != nil {
		return errors.Wrapf(err, "invalid install options")
	}
	version := opts.Version
	namespace := opts.Namespace
	installInfo, ok := supportedIstioVersions[version]
	if !ok {
		return errors.Errorf("%v is not a supported istio version. available versions and their chart locations: %v", version, supportedIstioVersions)
	}

	helmValueOverrides, err := template.New("istio-" + version).Parse(installInfo.valuesTemplate)
	if err != nil {
		return errors.Wrapf(err, "")
	}

	valuesBuf := &bytes.Buffer{}
	if err := helmValueOverrides.Execute(valuesBuf, opts); err != nil {
		return errors.Wrapf(err, "internal error: rendering helm values")
	}

	manifests, err := helm.RenderManifests(
		ctx,
		installInfo.chartPath,
		valuesBuf.String(),
		releaseName(namespace, version),
		namespace,
		"", // NOTE(ilackarms): use helm default
		true,
	)
	if err != nil {
		return errors.Wrapf(err, "rendering manifests")
	}

	for i, m := range manifests {
		// replace all instances of istio-system with the target namespace
		// based on instructions at https://istio.io/blog/2018/soft-multitenancy/#multiple-istio-control-planes
		m.Content = strings.Replace(m.Content, "istio-system", namespace, -1)
		manifests[i] = m
	}

	if err := helm.CreateFromManifests(ctx, namespace, manifests); err != nil {
		return errors.Wrapf(err, "creating istio from manifests")
	}

	return nil
}
