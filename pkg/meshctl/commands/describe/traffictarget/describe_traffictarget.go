package traffictarget

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
	discoveryv1alpha2 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha2"
	"github.com/solo-io/service-mesh-hub/pkg/meshctl/commands/describe/internal/flags"
	"github.com/solo-io/service-mesh-hub/pkg/meshctl/commands/describe/printing"
	"github.com/solo-io/service-mesh-hub/pkg/meshctl/utils"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Command(ctx context.Context, opts *flags.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "traffictarget",
		Short:   "Description of managed traffic targets",
		Aliases: []string{"traffictargets"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := utils.BuildClient(opts.Kubeconfig, opts.Kubecontext)
			if err != nil {
				return err
			}
			description, err := describeTrafficTargets(ctx, c)
			if err != nil {
				return err
			}
			fmt.Println(description)
			return nil
		},
	}

	cmd.SilenceUsage = true
	return cmd
}

func describeTrafficTargets(ctx context.Context, c client.Client) (string, error) {
	trafficTargetClient := discoveryv1alpha2.NewTrafficTargetClient(c)
	trafficTargetList, err := trafficTargetClient.ListTrafficTarget(ctx)
	if err != nil {
		return "", err
	}
	var trafficTargetDescriptions []trafficTargetDescription
	for _, trafficTarget := range trafficTargetList.Items {
		trafficTarget := trafficTarget // pike
		trafficTargetDescriptions = append(trafficTargetDescriptions, describeTrafficTarget(&trafficTarget))
	}

	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeader([]string{"Metadata", "Traffic_Policies", "Access_Policies"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	for _, description := range trafficTargetDescriptions {
		table.Append([]string{
			description.Metadata.string(),
			printing.FormattedObjectRefs(description.TrafficPolicies),
			printing.FormattedObjectRefs(description.AccessPolicies),
		})
	}
	table.Render()

	return buf.String(), nil
}

func (m trafficTargetMetadata) string() string {
	var s strings.Builder
	s.WriteString(printing.FormattedField("Name", m.Name))
	s.WriteString(printing.FormattedField("Namespace", m.Namespace))
	s.WriteString(printing.FormattedField("Cluster", m.Cluster))
	s.WriteString(printing.FormattedField("Type", m.Type))
	return s.String()
}

type trafficTargetDescription struct {
	Metadata        *trafficTargetMetadata
	TrafficPolicies []*v1.ObjectRef
	AccessPolicies  []*v1.ObjectRef
}

type trafficTargetMetadata struct {
	Type      string
	Name      string
	Namespace string
	Cluster   string
}

func describeTrafficTarget(trafficTarget *discoveryv1alpha2.TrafficTarget) trafficTargetDescription {
	meshMeta := getTrafficTargetMetadata(trafficTarget)

	var trafficPolicies []*v1.ObjectRef
	for _, fs := range trafficTarget.Status.AppliedTrafficPolicies {
		trafficPolicies = append(trafficPolicies, fs.Ref)
	}

	var accessPolicies []*v1.ObjectRef
	for _, vm := range trafficTarget.Status.AppliedAccessPolicies {
		accessPolicies = append(accessPolicies, vm.Ref)
	}

	return trafficTargetDescription{
		Metadata:        &meshMeta,
		TrafficPolicies: trafficPolicies,
		AccessPolicies:  accessPolicies,
	}
}

func getTrafficTargetMetadata(trafficTarget *discoveryv1alpha2.TrafficTarget) trafficTargetMetadata {
	switch trafficTarget.Spec.GetType().(type) {
	case *discoveryv1alpha2.TrafficTargetSpec_KubeService_:
		kubeServiceRef := trafficTarget.Spec.GetKubeService().Ref
		return trafficTargetMetadata{
			Type:      "kubernetes service",
			Name:      kubeServiceRef.Name,
			Namespace: kubeServiceRef.Namespace,
			Cluster:   kubeServiceRef.ClusterName,
		}
	}
	return trafficTargetMetadata{}
}
