package install

import (
	"github.com/pkg/errors"
	"github.com/solo-io/supergloo/cli/pkg/constants"
	"github.com/solo-io/supergloo/cli/pkg/flagutils"
	"github.com/solo-io/supergloo/cli/pkg/options"
	"github.com/solo-io/supergloo/cli/pkg/surveyutils"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/spf13/cobra"
)

func installIstioCmd(opts *options.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "istio",
		Short: "install the Istio control plane",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Interactive {
				if err := surveyutils.SurveyMetadata("installation", &opts.Metadata); err != nil {
					return err
				}
				if err := surveyutils.SurveyIstioInstall(&opts.Install); err != nil {
					return err
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return createIstioInstall(opts)
		},
	}
	flagutils.AddMetadataFlags(cmd.PersistentFlags(), &opts.Metadata)
	flagutils.AddOutputFlag(cmd.PersistentFlags(), &opts.OutputType)
	flagutils.AddInteractiveFlag(cmd.PersistentFlags(), &opts.Interactive)
	flagutils.AddIstioInstallFlags(cmd.PersistentFlags(), &opts.Install)
	return cmd
}

func createIstioInstall(opts *options.Options) error {
	install, err := installIstioFromOpts(opts)
	if err != nil {
		return err
	}
	return createInstall(opts, install)
}

func installIstioFromOpts(opts *options.Options) (*v1.Install, error) {
	if err := validateIstioInstall(opts.Install); err != nil {
		return nil, err
	}
	in := &v1.Install{
		Metadata:              opts.Metadata,
		InstallationNamespace: opts.Install.InstallationNamespace.Istio,
		InstallType: &v1.Install_Mesh{
			Mesh: &v1.MeshInstall{
				InstallType: &v1.MeshInstall_IstioMesh{
					IstioMesh: &opts.Install.IstioInstall,
				},
			},
		},
	}

	return in, nil
}

func validateIstioInstall(in options.Install) error {
	var validVersion bool
	for _, ver := range constants.SupportedIstioVersions {
		if in.IstioInstall.IstioVersion == ver {
			validVersion = true
			break
		}
	}
	if !validVersion {
		return errors.Errorf("%v is not a supported "+
			"istio version", in.IstioInstall.IstioVersion)
	}

	return nil
}
