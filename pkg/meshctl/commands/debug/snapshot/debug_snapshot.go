package snapshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/rotisserie/eris"

	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/utils"
	"github.com/solo-io/k8s-utils/debugutils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	filePermissions = 0644

	// filters for snapshots
	networking           = "networking"
	discovery            = "discovery"
	enterpriseNetworking = "enterprise-networking"
	enterpriseAgent      = "enterprise-agent"
	input                = "input"
	output               = "output"

	// jq query
	query = "to_entries | .[] | select(.key != \"clusters\") | select(.key != \"name\") | {kind: .key, list : [.value[]? | {name: .metadata.name, namespace: .metadata.namespace, cluster: .metadata.clusterName}]}"
)

type DebugSnapshotOpts struct {
	json    bool
	file    string
	dir     string
	zip     string
	verbose bool

	kubeconfig  string
	kubecontext string

	// hidden optional values
	metricsBindPort uint32
	namespace       string
}

func AddDebugSnapshotFlags(flags *pflag.FlagSet, opts *DebugSnapshotOpts) {
	utils.AddManagementKubeconfigFlags(&opts.kubeconfig, &opts.kubecontext, flags)
	flags.BoolVar(&opts.json, "json", false, "display the entire json snapshot. The output can be piped into a command like jq (https://stedolan.github.io/jq/tutorial/). For example:\n meshctl debug snapshot discovery input | jq '.'")
	flags.StringVarP(&opts.file, "file", "f", "", "file to write output to")
	flags.StringVar(&opts.dir, "dir", "", "dir to write file outputs to")
	flags.StringVar(&opts.zip, "zip", "", "zip file output")
	flags.Uint32Var(&opts.metricsBindPort, "port", defaults.MetricsPort, "metrics port")
	flags.StringVarP(&opts.namespace, "namespace", "n", defaults.GetPodNamespace(), "gloo-mesh namespace")
}

func Command(ctx context.Context, globalFlags *utils.GlobalFlags) *cobra.Command {
	opts := &DebugSnapshotOpts{}

	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "Input and Output snapshots for the discovery and networking pods. Requires jq to be installed if the --json flag is not being used.",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.verbose = globalFlags.Verbose
			return debugSnapshot(ctx, opts, []string{discovery, networking, enterpriseNetworking, enterpriseAgent}, []string{input, output})
		},
	}
	cmd.AddCommand(
		Networking(ctx, opts),
		Discovery(ctx, opts),
		EnterpriseNetworking(ctx, opts),
		EnterpriseAgent(ctx, opts),
	)
	AddDebugSnapshotFlags(cmd.PersistentFlags(), opts)

	cmd.PersistentFlags().Lookup("namespace").Hidden = true
	cmd.PersistentFlags().Lookup("port").Hidden = true
	return cmd
}

func EnterpriseNetworking(ctx context.Context, opts *DebugSnapshotOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enterprise-networking",
		Short: "Input and output snapshots for the enterprise networking pod",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugSnapshot(ctx, opts, []string{enterpriseNetworking}, []string{input, output})
		},
	}
	cmd.AddCommand(
		Input(ctx, opts, enterpriseNetworking),
		Output(ctx, opts, enterpriseNetworking),
	)
	return cmd
}

func EnterpriseAgent(ctx context.Context, opts *DebugSnapshotOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enterprise-agent",
		Short: "Input and output snapshots for the enterprise agent pod",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugSnapshot(ctx, opts, []string{enterpriseAgent}, []string{input, output})
		},
	}
	cmd.AddCommand(
		Input(ctx, opts, enterpriseAgent),
		Output(ctx, opts, enterpriseAgent),
	)
	return cmd
}

func Networking(ctx context.Context, opts *DebugSnapshotOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "networking",
		Short: "Input and output snapshots for the networking pod",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugSnapshot(ctx, opts, []string{networking}, []string{input, output})
		},
	}
	cmd.AddCommand(
		Input(ctx, opts, networking),
		Output(ctx, opts, networking),
	)
	return cmd
}

func Discovery(ctx context.Context, opts *DebugSnapshotOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discovery",
		Short: "Input and output snapshots for the discovery pod",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugSnapshot(ctx, opts, []string{discovery}, []string{input, output})
		},
	}
	cmd.AddCommand(
		Input(ctx, opts, discovery),
		Output(ctx, opts, discovery),
	)
	return cmd
}

func Input(ctx context.Context, opts *DebugSnapshotOpts, pod string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "input",
		Short: fmt.Sprintf("Input snapshot for the %s pod", pod),
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugSnapshot(ctx, opts, []string{pod}, []string{input})
		},
	}
	return cmd
}

func Output(ctx context.Context, opts *DebugSnapshotOpts, pod string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output",
		Short: fmt.Sprintf("Output snapshot for the %s pod", pod),
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugSnapshot(ctx, opts, []string{pod}, []string{output})
		},
	}
	return cmd
}

func debugSnapshot(ctx context.Context, opts *DebugSnapshotOpts, pods, types []string) error {
	// Check prerequisite jq is installed
	_, err := exec.Command("which", "jq").Output()
	if err != nil {
		fmt.Printf("Error: Could not find jq! Please install it from https://stedolan.github.io/jq/download/\n")
		return err
	}

	f, err := os.Create(opts.file)
	defer f.Close()

	fs := afero.NewOsFs()
	zipDir, err := afero.TempDir(fs, "", "")
	if err != nil {
		return err
	}
	defer fs.RemoveAll(zipDir)
	storageClient := debugutils.NewFileStorageClient(fs)
	for _, podName := range pods {
		for _, snapshotType := range types {
			fmt.Printf("%s snapshot for %s\n", snapshotType, podName)
			snapshot, snapshotErr := getSnapshot(ctx, opts, "", podName, snapshotType)
			if snapshotErr != nil {
				fmt.Println(snapshotErr.Error())
				continue
			}
			fileName := fmt.Sprintf("%s-%s-snapshot.json", podName, snapshotType)
			var snapshotStr string
			if opts.json {
				snapshotStr = snapshot
			} else {
				err = ioutil.WriteFile(fileName, []byte(snapshot), filePermissions)
				if err != nil {
					return err
				}
				pipedCmd := "cat " + fileName + " | jq '" + query + "'"
				cmdOut, err := exec.Command("bash", "-c", pipedCmd).Output()
				if err != nil {
					return err
				}
				snapshotStr = string(cmdOut)
				err = os.Remove(fileName)
				if err != nil {
					return err
				}
			}
			if opts.file != "" {
				_, err = f.WriteString(snapshotStr)
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				fmt.Printf("Written to %s\n", opts.file)
			} else if opts.zip != "" || opts.dir != "" {
				if len(snapshotStr) == 0 {
					continue
				}
				dir := zipDir
				if opts.dir != "" {
					dir = opts.dir
				}
				err = storageClient.Save(dir, &debugutils.StorageObject{
					Resource: strings.NewReader(snapshotStr),
					Name:     fileName,
				})
				if err != nil {
					return err
				}
				fmt.Printf("Written to %s\n", fileName)
			} else {
				fmt.Print(snapshotStr)
				fmt.Print("\n")
			}
		}
	}
	if opts.zip != "" {
		err = utils.Zip(fs, zipDir, opts.zip)
	}
	return nil
}

func getSnapshot(ctx context.Context, opts *DebugSnapshotOpts, localPort, podName, snapshotType string) (string, error) {
	kubeClient, err := utils.BuildClientset(opts.kubeconfig, opts.kubecontext)
	if err != nil {
		return "", err
	}
	_, err = kubeClient.AppsV1().Deployments(opts.namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return "", eris.Errorf("No %s.%s deployment found - skipping the %s snapshot\n", opts.namespace, podName, snapshotType)
	}

	portFwdContext, cancelPtFwd := context.WithCancel(ctx)
	mgmtDeployNamespace := opts.namespace
	mgmtDeployName := podName
	remotePort := strconv.Itoa(int(opts.metricsBindPort))
	// start port forward to mgmt server stats port
	localPort, err = utils.PortForwardFromDeployment(
		portFwdContext,
		opts.kubeconfig,
		opts.kubecontext,
		mgmtDeployName,
		mgmtDeployNamespace,
		fmt.Sprintf("%v", localPort),
		fmt.Sprintf("%v", remotePort),
	)
	if err != nil {
		return "", eris.Errorf("try verifying that `kubectl port-forward -n %v deployment/%v %v:%v` can be run successfully.", mgmtDeployNamespace, mgmtDeployName, localPort, remotePort)
	}
	// request snapshots page
	snapshotUrl := fmt.Sprintf("http://localhost:%v/snapshots/%s", localPort, snapshotType)
	resp, err := http.DefaultClient.Get(snapshotUrl)
	if err != nil {
		return "", eris.Errorf("try verifying that the %s pod is listening on port %v", podName, remotePort)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	snapshot := string(b)

	cancelPtFwd()

	return snapshot, nil
}
