package traffic_policy

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/solo-io/autopilot/pkg/utils"
	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	common_config "github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/interactive"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/resource_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create/prompts"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/create/validate"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_networking_types "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	"github.com/spf13/cobra"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

const (
	DoneSelectingOption = "DONE"
)

type CreateTrafficPolicyCmd *cobra.Command

func CreateTrafficPolicyCommand(
	ctx context.Context,
	out io.Writer,
	opts *options.Options,
	kubeLoader common_config.KubeLoader,
	kubeClientsFactory common.KubeClientsFactory,
	interactivePrompt interactive.InteractivePrompt,
	printers common.Printers,
) CreateTrafficPolicyCmd {
	cmd := cliconstants.CreateTrafficPolicyCommand
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return createTrafficPolicy(ctx, out, kubeLoader, kubeClientsFactory, opts, interactivePrompt, printers.ResourcePrinter)
	}
	return &cmd
}

func createTrafficPolicy(
	ctx context.Context,
	out io.Writer,
	kubeLoader common_config.KubeLoader,
	kubeClientsFactory common.KubeClientsFactory,
	opts *options.Options,
	interactivePrompt interactive.InteractivePrompt,
	resourcePrinter resource_printing.ResourcePrinter,
) error {
	var err error
	var masterCfg *rest.Config
	var masterKubeClients *common.KubeClients
	var meshServiceNames []string
	var meshServiceNamesToRefs map[string]*zephyr_core_types.ResourceRef
	var sourceSelector *zephyr_core_types.WorkloadSelector
	var targetSelector *zephyr_core_types.ServiceSelector
	var trafficShift *zephyr_networking_types.TrafficPolicySpec_MultiDestination
	if masterCfg, err = kubeLoader.GetRestConfigForContext(opts.Root.KubeConfig, opts.Root.KubeContext); err != nil {
		return err
	}
	if masterKubeClients, err = kubeClientsFactory(masterCfg, opts.Root.WriteNamespace); err != nil {
		return err
	}
	if meshServiceNames, meshServiceNamesToRefs, err = fetchMeshServiceRefs(ctx, masterKubeClients.MeshServiceClient); err != nil {
		return err
	}
	if sourceSelector, err = selectSourcesInteractively(interactivePrompt); err != nil {
		return err
	}
	if targetSelector, err = prompts.SelectServiceSelector(
		"Select the destination Services, omit to select all", meshServiceNames, meshServiceNamesToRefs, interactivePrompt); err != nil {
		return err
	}
	if trafficShift, err = selectTrafficShiftInteractively(meshServiceNames, meshServiceNamesToRefs, interactivePrompt); err != nil {
		return err
	}
	trafficPolicy := &zephyr_networking.TrafficPolicy{
		TypeMeta: k8s_meta_types.TypeMeta{
			Kind: "TrafficPolicy",
		},
		Spec: zephyr_networking_types.TrafficPolicySpec{
			SourceSelector:      sourceSelector,
			DestinationSelector: targetSelector,
			TrafficShift:        trafficShift,
		},
	}
	if !opts.Create.DryRun {
		return masterKubeClients.TrafficPolicyClient.CreateTrafficPolicy(ctx, trafficPolicy)
	} else {
		return resourcePrinter.Print(out, trafficPolicy, resource_printing.OutputFormat(opts.Create.OutputFormat))
	}
}

func selectSourcesInteractively(interactivePrompt interactive.InteractivePrompt) (*zephyr_core_types.WorkloadSelector, error) {
	labelSet, err := prompts.PromptLabels("Specify source workloads labels in the format (key1=value1, key2=value2), omit to permit workloads of with any labels", interactivePrompt)
	if err != nil {
		return nil, err
	}
	namespaces, err := prompts.PromptCommaDelimitedValues("Specify source workloads namespaces as comma-delimited list, omit to permit workloads of any namespace", interactivePrompt)
	if err != nil {
		return nil, err
	}
	return &zephyr_core_types.WorkloadSelector{Labels: labelSet, Namespaces: namespaces}, nil
}

func selectTrafficShiftInteractively(
	meshServiceNames []string,
	meshServiceNamesToRef map[string]*zephyr_core_types.ResourceRef,
	interactivePrompt interactive.InteractivePrompt,
) (*zephyr_networking_types.TrafficPolicySpec_MultiDestination, error) {
	var err error
	var meshServiceName, weightString, portString string
	var subsetLabels map[string]string
	var weightedDestinations []*zephyr_networking_types.TrafficPolicySpec_MultiDestination_WeightedDestination
	// Add select all option
	meshServiceNames = append([]string{DoneSelectingOption}, meshServiceNames...)
	for {
		weightedDest := &zephyr_networking_types.TrafficPolicySpec_MultiDestination_WeightedDestination{}
		if meshServiceName, err = interactivePrompt.SelectValue(
			"Select a Service to shift traffic to",
			meshServiceNames); err != nil {
			return nil, err
		}
		// User signal to finish selecting traffic shift
		if meshServiceName == DoneSelectingOption {
			break
		}
		// Remove item from subsequent selection
		meshServiceNames = utils.RemoveString(meshServiceNames, meshServiceName)
		// Prefix subsequent prompts with the current k8s Service the user is configuring
		messagePrefix := fmt.Sprintf("(%s)", meshServiceName)
		if portString, err = interactivePrompt.PromptValueWithValidator(
			fmt.Sprintf("%s Set Service port, omitting uses default Service port", messagePrefix), "",
			validate.AllowEmpty(validate.PositiveInteger)); err != nil {
			return nil, err
		}
		if weightString, err = interactivePrompt.PromptValueWithValidator(
			fmt.Sprintf("%s Set traffic shift weight, omitting defaults to value of 1", messagePrefix), "1",
			validate.AllowEmpty(validate.PositiveInteger)); err != nil {
			return nil, err
		}
		if subsetLabels, err = prompts.PromptLabels(
			fmt.Sprintf("%s If routing to subset, specify subset selector labels", messagePrefix), interactivePrompt); err != nil {
			return nil, err
		}
		weightedDest.Destination = meshServiceNamesToRef[meshServiceName]
		weightedDest.Subset = subsetLabels
		if portString != "" {
			portNum, err := strconv.Atoi(portString)
			if err != nil {
				return nil, err
			}
			weightedDest.Port = uint32(portNum)
		}
		weight, err := strconv.Atoi(weightString)
		if err != nil {
			return nil, err
		}
		weightedDest.Weight = uint32(weight)
		weightedDestinations = append(weightedDestinations, weightedDest)
	}
	return &zephyr_networking_types.TrafficPolicySpec_MultiDestination{Destinations: weightedDestinations}, nil
}

func fetchMeshServiceRefs(
	ctx context.Context,
	meshServiceClient zephyr_discovery.MeshServiceClient,
) ([]string,
	map[string]*zephyr_core_types.ResourceRef,
	error) {
	meshServices, err := meshServiceClient.ListMeshService(ctx)
	if err != nil {
		return nil, nil, err
	}
	var meshServiceNames []string
	meshServiceNamesToRef := map[string]*zephyr_core_types.ResourceRef{}
	for _, meshService := range meshServices.Items {
		meshService := meshService
		meshServiceName := buildMeshServiceName(&meshService)
		meshServiceNames = append(meshServiceNames, meshServiceName)
		meshServiceNamesToRef[meshServiceName] = &zephyr_core_types.ResourceRef{
			Name:      meshService.GetName(),
			Namespace: meshService.GetNamespace(),
			Cluster:   meshService.GetClusterName(),
		}
	}
	return meshServiceNames, meshServiceNamesToRef, nil
}

func buildMeshServiceName(meshService *zephyr_discovery.MeshService) string {
	return fmt.Sprintf("%s.%s.%s", meshService.GetName(), meshService.GetNamespace(), meshService.Spec.GetKubeService().GetRef().GetCluster())
}
