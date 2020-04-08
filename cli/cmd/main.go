package main

import (
	"context"
	"os"

	"github.com/solo-io/service-mesh-hub/cli/pkg/wire"
)

func main() {
	cliApp := wire.InitializeCLI(context.Background(), os.Stdout, os.Stdin)
	err := cliApp.Execute()
	if err != nil {
		os.Exit(1)
	}
}
