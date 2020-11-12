package failoverservice

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
	networkingv1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/describe/printing"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/utils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Command(ctx context.Context) *cobra.Command {
	opts := &options{}
	cmd := &cobra.Command{
		Use:     "failoverservice",
		Short:   "Description of failover services",
		Aliases: []string{"failoverservices"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := utils.BuildClient(opts.kubeconfig, opts.kubecontext)
			if err != nil {
				return err
			}
			description, err := describeFailoverServices(ctx, c, opts.searchTerms)
			if err != nil {
				return err
			}
			fmt.Println(description)
			return nil
		},
	}
	opts.addToFlags(cmd.Flags())

	cmd.SilenceUsage = true
	return cmd
}

type options struct {
	kubeconfig  string
	kubecontext string
	searchTerms []string
}

func (o *options) addToFlags(flags *pflag.FlagSet) {
	utils.AddManagementKubeconfigFlags(&o.kubeconfig, &o.kubecontext, flags)
	flags.StringSliceVarP(&o.searchTerms, "search", "s", []string{}, "A list of terms to match failover service names against")
}

func describeFailoverServices(ctx context.Context, c client.Client, searchTerms []string) (string, error) {
	failoverServiceClient := networkingv1alpha2.NewFailoverServiceClient(c)
	failoverServiceList, err := failoverServiceClient.ListFailoverService(ctx)
	if err != nil {
		return "", err
	}
	var failoverServiceDescription []failoverServiceDescription
	for _, failoverService := range failoverServiceList.Items {
		failoverService := failoverService // pike
		if matchFailoverService(failoverService, searchTerms) {
			failoverServiceDescription = append(failoverServiceDescription, describeFailoverService(&failoverService))
		}
	}

	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Metadata", "Net", "Meshes", "Backing_Services"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	for _, description := range failoverServiceDescription {
		table.Append([]string{
			printing.FormattedClusterObjectRef(description.Metadata),
			description.Net.string(),
			printing.FormattedObjectRefs(description.Meshes),
			printing.FormattedClusterObjectRefs(description.BackingServices),
		})
	}
	table.Render()

	return buf.String(), nil
}

func (n failoverServiceNet) string() string {
	var s strings.Builder
	s.WriteString(printing.FormattedField("Hostname", n.Hostname))
	s.WriteString(printing.FormattedField("Port", fmt.Sprint(n.Port)))
	return s.String()
}

type failoverServiceDescription struct {
	Metadata        *v1.ClusterObjectRef
	Net             *failoverServiceNet
	Meshes          []*v1.ObjectRef
	BackingServices []*v1.ClusterObjectRef
}

type failoverServiceNet struct {
	Hostname string
	Port     uint32
}

func matchFailoverService(failoverService networkingv1alpha2.FailoverService, searchTerms []string) bool {
	// do not apply matching when there are no search strings
	if len(searchTerms) == 0 {
		return true
	}

	for _, s := range searchTerms {
		if strings.Contains(failoverService.Name, s) {
			return true
		}
	}

	return false
}

func describeFailoverService(failoverService *networkingv1alpha2.FailoverService) failoverServiceDescription {
	failoverServiceMeta := getFailoverServiceMetadata(failoverService)
	failoverServiceNet := getFailoverServiceNet(failoverService)

	var backingServices []*v1.ClusterObjectRef
	for _, bs := range failoverService.Spec.GetBackingServices() {
		switch bst := bs.GetBackingServiceType().(type) {
		case *networkingv1alpha2.FailoverServiceSpec_BackingService_KubeService:
			backingServices = append(backingServices, bst.KubeService)
		}
	}

	return failoverServiceDescription{
		Metadata:        &failoverServiceMeta,
		Net:             &failoverServiceNet,
		Meshes:          failoverService.Spec.GetMeshes(),
		BackingServices: backingServices,
	}
}

func getFailoverServiceMetadata(failoverService *networkingv1alpha2.FailoverService) v1.ClusterObjectRef {
	return v1.ClusterObjectRef{
		Name:        failoverService.Name,
		Namespace:   failoverService.Namespace,
		ClusterName: failoverService.ClusterName,
	}
}

func getFailoverServiceNet(failoverService *networkingv1alpha2.FailoverService) failoverServiceNet {
	return failoverServiceNet{
		Hostname: failoverService.Spec.Hostname,
		Port:     failoverService.Spec.Port.Number,
	}
}
