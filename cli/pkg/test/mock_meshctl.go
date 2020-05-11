package cli_test

import (
	"bytes"
	"context"

	"github.com/golang/mock/gomock"
	"github.com/mattn/go-shellwords"
	. "github.com/onsi/ginkgo"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/exec"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/files"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/interactive"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/usage"
	usage_mocks "github.com/solo-io/service-mesh-hub/cli/pkg/common/usage/mocks"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/wire"
	"github.com/solo-io/service-mesh-hub/pkg/common/docker"
	"github.com/solo-io/service-mesh-hub/pkg/kubeconfig"
	"k8s.io/client-go/rest"
)

// Build and execute the CLI app using the given clients
type MockMeshctl struct {
	// MUST be non-nil - we always need to produce a mocked master cluster verification client and a mocked usage reporter
	MockController *gomock.Controller

	// Stdin which is turned into a *bytes.Buffer and passed into the command chain as an io.Reader
	Stdin string

	Ctx context.Context

	Clients common.Clients

	KubeClients common.KubeClients

	KubeLoader      kubeconfig.KubeLoader
	ImageNameParser docker.ImageNameParser
	FileReader      files.FileReader
	KubeConverter   kubeconfig.Converter

	Runner            exec.Runner
	Printers          common.Printers
	InteractivePrompt interactive.InteractivePrompt
}

// call with the same string you would pass to the meshctl binary, i.e. "cluster register --service-account-name test123"
// returns the output produced by meshctl on stdout as a string
func (m MockMeshctl) Invoke(argString string) (stdout string, err error) {
	if m.MockController == nil || m.Ctx == nil {
		Fail("Must provide both the ginkgo mock controller and a non-nil context to mock meshctl")
	}

	buffer := &bytes.Buffer{}

	usageReporter := usage_mocks.NewMockClient(m.MockController)
	usageReporter.
		EXPECT().
		StartReportingUsage(m.Ctx, usage.UsageReportingInterval).
		Return(nil)

	kubeFactory := func(masterConfig *rest.Config, writeNamespace string) (*common.KubeClients, error) {
		return &m.KubeClients, nil
	}

	clientsFactory := func(opts *options.Options) (*common.Clients, error) {
		return &m.Clients, nil
	}
	app := wire.InitializeCLIWithMocks(
		m.Ctx,
		buffer,
		bytes.NewBufferString(m.Stdin),
		usageReporter,
		kubeFactory,
		clientsFactory,
		m.KubeLoader,
		m.ImageNameParser,
		m.FileReader,
		m.KubeConverter,
		m.Printers,
		m.Runner,
		m.InteractivePrompt,
	)

	splitArgs, err := shellwords.Parse(argString)
	if err != nil {
		panic("Bad arg string: " + argString)
	}

	app.SetArgs(splitArgs)

	err = app.Execute()

	return buffer.String(), err
}
