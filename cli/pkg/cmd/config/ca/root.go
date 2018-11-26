package ca

import (
	"fmt"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/cli/pkg/cmd/options"
	"github.com/solo-io/supergloo/cli/pkg/common"
	"github.com/solo-io/supergloo/cli/pkg/nsutil"
	"github.com/spf13/cobra"
)

func Cmd(opts *options.Options) *cobra.Command {
	cOpts := &(opts.Config).Ca
	cmd := &cobra.Command{
		Use:   "ca",
		Short: `Update CA`,
		Long:  `Update CA`,
		Run: func(c *cobra.Command, args []string) {
			err := configureCa(opts)
			// TODO pass err upwards
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&cOpts.Mesh.Name, "mesh.name", "", "name of mesh to update")

	flags.StringVar(&cOpts.Mesh.Namespace, "mesh.namespace", "", "namespace of mesh to update")

	flags.StringVar(&cOpts.Secret.Name, "secret.name", "", "name of secret to apply")

	flags.StringVar(&cOpts.Secret.Namespace, "secret.namespace", "", "namespace of secret to apply")

	return cmd
}

func configureCa(opts *options.Options) error {

	// validate flags and gather interactive fields
	err := ensureFlags(opts)
	if err != nil {
		return err
	}

	meshClient, err := common.GetMeshClient()
	if err != nil {
		return err
	}
	mesh, err := (*meshClient).Read(opts.Config.Ca.Mesh.Namespace, opts.Config.Ca.Mesh.Name, clients.ReadOpts{})
	if err != nil {
		return err
	}

	if !mesh.Encryption.TlsEnabled {
		return fmt.Errorf("TLS is not enabled on mesh %v. You must first enable TLS before configuring CA.", opts.Config.Ca.Mesh.Name)
	}

	// TODO(mitchdraft) replace options.ResourceRef with core.ResourceRef
	// *(mesh.Encryption.Secret) = opts.Config.Ca.Secret
	mesh.Encryption.Secret = &core.ResourceRef{
		Name:      opts.Config.Ca.Secret.Name,
		Namespace: opts.Config.Ca.Secret.Namespace,
	}

	_, err = (*meshClient).Write(mesh, clients.WriteOpts{OverwriteExisting: true})
	if err != nil {
		return err
	}

	fmt.Printf("Configured mesh %v to use secret %v", opts.Config.Ca.Mesh.Name, opts.Config.Ca.Secret.Name)
	return nil
}

func ensureFlags(opts *options.Options) error {

	oMeshRef := &(opts.Config.Ca).Mesh
	if err := nsutil.EnsureMesh(oMeshRef, opts); err != nil {
		return err
	}

	oSecretRef := &(opts.Config.Ca).Secret
	if err := nsutil.EnsureCommonResource("secret", "secret", oSecretRef, opts); err != nil {
		return err
	}

	return nil
}
