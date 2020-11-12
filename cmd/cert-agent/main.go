package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/solo-io/gloo-mesh/pkg/certificates/agent"
	"github.com/solo-io/gloo-mesh/pkg/common/bootstrap"
	"github.com/solo-io/gloo-mesh/pkg/common/version"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	if err := rootCommand(ctx).Execute(); err != nil {
		contextutils.LoggerFrom(ctx).Fatal(err)
	}
	contextutils.LoggerFrom(ctx).Info("exiting...")
}

func rootCommand(ctx context.Context) *cobra.Command {
	opts := &bootstrap.Options{}
	cmd := &cobra.Command{
		Use:     "cert-agent [command]",
		Short:   "Start the Gloo Mesh Certificate Agent.",
		Long:    "The Gloo Mesh Certificate Agent is used to generate certificates signed by Gloo Mesh for use in managed clusters without requiring private keys to leave the managed cluster. For documentation on the actions taken by the Certificate Agent, see the generated documentation for the IssuedCertificate Custom Resource.",
		Version: version.Version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logrus.SetLevel(logrus.DebugLevel)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return startAgent(ctx, opts)
		},
	}

	opts.AddToFlags(cmd.PersistentFlags())

	return cmd
}

func startAgent(ctx context.Context, opts *bootstrap.Options) error {
	return agent.Start(ctx, *opts)
}
