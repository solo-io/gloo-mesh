package smi

import (
	"sort"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	"github.com/solo-io/supergloo/api/external/smi/access"
	"github.com/solo-io/supergloo/api/external/smi/specs"
	"github.com/solo-io/supergloo/imported/deislabs/smi-sdk-go/pkg/apis/access/v1alpha1"
	specv1alpha1 "github.com/solo-io/supergloo/imported/deislabs/smi-sdk-go/pkg/apis/specs/v1alpha1"
	sgaccess "github.com/solo-io/supergloo/pkg/api/external/smi/access/v1alpha1"
	sgspec "github.com/solo-io/supergloo/pkg/api/external/smi/specs/v1alpha1"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	"github.com/solo-io/supergloo/pkg/translator/utils"
)

type SecurityConfig struct {
	TrafficTargets  sgaccess.TrafficTargetList
	HTTPRouteGroups sgspec.HTTPRouteGroupList
}

func (c *SecurityConfig) Sort() {
	sort.SliceStable(c.TrafficTargets, func(i, j int) bool {
		return c.TrafficTargets[i].GetMetadata().Less(c.TrafficTargets[j].GetMetadata())
	})
	sort.SliceStable(c.HTTPRouteGroups, func(i, j int) bool {
		return c.HTTPRouteGroups[i].GetMetadata().Less(c.HTTPRouteGroups[j].GetMetadata())
	})
}
func createSecurityConfig(rules v1.SecurityRuleList,
	upstreams gloov1.UpstreamList,
	pods kubernetes.PodList,
	services kubernetes.ServiceList,
	resourceErrs reporter.ResourceErrors) SecurityConfig {

	var sc SecurityConfig

	// for each rule:
	for _, r := range rules {
		// create a route group and corresponding traffic targets for that rule
		trafficTargets, routeGroup, err := createTrafficTargetsFromRule(r, upstreams, services, pods)
		if err != nil {
			resourceErrs.AddError(r, err)
			continue
		}
		sc.HTTPRouteGroups = append(sc.HTTPRouteGroups, routeGroup)
		sc.TrafficTargets = append(sc.TrafficTargets, trafficTargets...)
	}

	return sc
}

const serviceAccountKind = "ServiceAccount"
const httpRouteGroupKind = "HTTPRouteGroup"

func createTrafficTargetsFromRule(rule *v1.SecurityRule, upstreams gloov1.UpstreamList, services kubernetes.ServiceList, pods kubernetes.PodList) (sgaccess.TrafficTargetList, *sgspec.HTTPRouteGroup, error) {
	sourcePods, err := utils.PodsForSelector(rule.SourceSelector, upstreams, pods)
	if err != nil {
		return nil, nil, err
	}
	var sourceIdentites []v1alpha1.IdentityBindingSubject
	for _, sourcePod := range sourcePods {
		serviceAcct := sourcePod.Spec.ServiceAccountName

		var alreadyAdded bool
		for _, identity := range sourceIdentites {
			if identity.Name == serviceAcct && identity.Namespace == sourcePod.Namespace {
				alreadyAdded = true
				break
			}
		}
		if alreadyAdded {
			continue
		}
		sourceIdentites = append(sourceIdentites, v1alpha1.IdentityBindingSubject{
			Kind:      serviceAccountKind,
			Name:      serviceAcct,
			Namespace: sourcePod.Namespace,
		})
	}

	destPods, err := utils.PodsForSelector(rule.DestinationSelector, upstreams, pods)
	if err != nil {
		return nil, nil, err
	}
	destServices, err := utils.ServicesForSelector(rule.DestinationSelector, upstreams, services)
	if err != nil {
		return nil, nil, err
	}
	var destIdentites []v1alpha1.IdentityBindingSubject
	for _, destPod := range destPods {
		serviceAcct := destPod.Spec.ServiceAccountName
		for _, destSvc := range destServices {
			for _, port := range destSvc.Spec.Ports {
				var alreadyAdded bool
				for _, identity := range destIdentites {
					if identity.Namespace == destPod.Namespace &&
						identity.Name == serviceAcct &&
						identity.Port == port.Name {
						alreadyAdded = true
						break
					}
				}
				if alreadyAdded {
					continue
				}
				destIdentites = append(destIdentites, v1alpha1.IdentityBindingSubject{
					Kind:      serviceAccountKind,
					Name:      serviceAcct,
					Namespace: destPod.Namespace,
					Port:      port.Name,
				})
			}
		}
	}

	var allowedMatches []specv1alpha1.HTTPMatch
	var matchNames []string
	for _, path := range rule.AllowedPaths {
		allowedMatches = append(allowedMatches, specv1alpha1.HTTPMatch{
			Name:      path,
			Methods:   rule.AllowedMethods,
			PathRegex: path,
		})
		matchNames = append(matchNames, path)
	}

	routeGroupName := rule.Metadata.Name
	targetSpec := []v1alpha1.TrafficTargetSpec{{
		Name:    routeGroupName,
		Kind:    httpRouteGroupKind,
		Matches: matchNames,
	}}

	var trafficTargets sgaccess.TrafficTargetList
	for _, destIdentiy := range destIdentites {
		trafficTargets = append(trafficTargets, &sgaccess.TrafficTarget{
			TrafficTarget: access.TrafficTarget{
				ObjectMeta: kubeutils.ToKubeMeta(core.Metadata{
					Namespace: rule.Metadata.Namespace,
					Name:      destIdentiy.Name,
				}),
				Destination: destIdentiy,
				Sources:     sourceIdentites,
				Specs:       targetSpec,
			},
		})
	}

	routeGroup := &sgspec.HTTPRouteGroup{
		HTTPRouteGroup: specs.HTTPRouteGroup{
			ObjectMeta: kubeutils.ToKubeMeta(core.Metadata{
				Namespace: rule.Metadata.Namespace,
				Name:      routeGroupName,
			}),
			Matches: allowedMatches,
		},
	}

	return trafficTargets, routeGroup, nil
}
