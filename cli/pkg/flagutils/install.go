package flagutils

import (
	"fmt"

	"github.com/solo-io/supergloo/cli/pkg/constants"

	"github.com/solo-io/supergloo/cli/pkg/options"
	"github.com/solo-io/supergloo/pkg/install/istio"
	"github.com/spf13/pflag"
)

func AddInstallFlags(set *pflag.FlagSet, in *options.Install) {

	set.BoolVar(&in.Update,
		"update",
		false,
		"update an existing install?")

}

func AddIstioInstallFlags(set *pflag.FlagSet, in *options.Install) {
	set.StringVar(&in.IstioInstall.InstallationNamespace,
		"installation-namespace",
		"istio-system",
		"which namespace to install Istio into?")

	set.StringVar(&in.IstioInstall.IstioVersion,
		"version",
		istio.IstioVersion106,
		fmt.Sprintf("version of istio to install? available: %v", constants.SupportedIstioVersions))

	set.BoolVar(&in.IstioInstall.EnableMtls,
		"mtls",
		true,
		"enable mtls?")

	set.BoolVar(&in.IstioInstall.EnableAutoInject,
		"auto-inject",
		true,
		"enable auto-injection?")

	set.BoolVar(&in.IstioInstall.InstallGrafana,
		"grafana",
		true,
		"add grafana to the install?")

	set.BoolVar(&in.IstioInstall.InstallPrometheus,
		"prometheus",
		true,
		"add prometheus to the install?")

	set.BoolVar(&in.IstioInstall.InstallJaeger,
		"jaeger",
		true,
		"add jaeger to the install?")

}
