package cleanup

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/rotisserie/eris"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

func Command(ctx context.Context, clusters ...string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup",
		Short: "Clean up bootstrapped local resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cleanup(ctx, clusters...)
		},
	}

	cmd.SilenceUsage = true
	return cmd
}

func cleanup(ctx context.Context, clusters ...string) error {
	fmt.Println("Cleaning up clusters")

	box := packr.NewBox("./scripts")
	script, err := box.FindString("delete_clusters.sh")
	if err != nil {
		return eris.Wrap(err, "Error loading script")
	}

	args := []string{"-c", script}
	args = append(args, clusters...)
	cmd := exec.CommandContext(ctx, "bash", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}
