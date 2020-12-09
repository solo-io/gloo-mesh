package gloomesh

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/helm"
)

type Uninstaller struct {
	KubeConfig  string
	KubeContext string
	Namespace   string
	ReleaseName string
	Verbose     bool
	DryRun      bool
}

func (i Uninstaller) UninstallGlooMesh(
	ctx context.Context,
) error {
	return i.uninstall(ctx, GlooMeshReleaseName)
}

func (i Uninstaller) UninstallAgentCrds(
	ctx context.Context,
) error {
	return i.uninstall(ctx, agentCrdsReleaseName)
}

func (i Uninstaller) UninstallCertAgent(
	ctx context.Context,
) error {
	return i.uninstall(ctx, certAgentReleaseName)
}

func (i Uninstaller) UninstallWasmAgent(
	ctx context.Context,
) error {
	return i.uninstall(ctx, wasmAgentReleaseName)
}

func (i Uninstaller) uninstall(
	ctx context.Context,
	releaseName string,
) error {
	if i.ReleaseName != "" {
		releaseName = i.ReleaseName
	}

	return helm.Uninstaller{
		KubeConfig:  i.KubeConfig,
		KubeContext: i.KubeContext,
		Namespace:   i.Namespace,
		ReleaseName: releaseName,
		Verbose:     i.Verbose,
		DryRun:      i.DryRun,
	}.UninstallChart(ctx)
}
