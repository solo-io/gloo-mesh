package authorizationpolicy

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/go-multierror"
	"github.com/rotisserie/eris"
	commonv1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discoveryv1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/input"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	securityv1beta1spec "istio.io/api/security/v1beta1"
	typesv1beta1 "istio.io/api/type/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
)

//go:generate mockgen -source ./authorization_policy_translator.go -destination mocks/authorization_policy_translator.go

const (
	translatorName = "authorization-policy-translator"
)

var (
	trustDomainNotFound = func(clusterName string) error {
		return eris.Errorf("Trust domain not found for cluster %s", clusterName)
	}
)

// the AuthorizationPolicy translator translates a Destination into a AuthorizationPolicy.
type Translator interface {
	// Translate translates an appropriate AuthorizationPolicy for the given Destination.
	// returns nil if no AuthorizationPolicy is required for the Destination (i.e. if no AuthorizationPolicy features are required, such access control).
	//
	// Errors caused by invalid user config will be reported using the Reporter.
	//
	// Note that the input snapshot DestinationSet contains the given Destination.
	Translate(
		in input.LocalSnapshot,
		destination *discoveryv1.Destination,
		reporter reporting.Reporter,
	) *securityv1beta1.AuthorizationPolicy
}

type translator struct{}

func NewTranslator() Translator {
	return &translator{}
}

func (t *translator) Translate(
	in input.LocalSnapshot,
	destination *discoveryv1.Destination,
	reporter reporting.Reporter,
) *securityv1beta1.AuthorizationPolicy {
	kubeService := destination.Spec.GetKubeService()

	if kubeService == nil {
		// TODO(harveyxia): non kube services currently unsupported
		return nil
	}

	authPolicy := t.initializeAuthorizationPolicy(destination)

	for _, policy := range destination.Status.AppliedAccessPolicies {
		rule, err := t.translateAccessPolicy(policy.Spec, in.Meshes())
		if err != nil {
			reporter.ReportAccessPolicyToDestination(destination, policy.Ref, eris.Wrapf(err, "%v", translatorName))
			continue
		}
		authPolicy.Spec.Rules = append(authPolicy.Spec.Rules, rule)
	}

	// don't output an AuthPolicy with no matching rules, which semantically denies all requests
	// reference: https://istio.io/latest/docs/reference/config/security/authorization-policy/#AuthorizationPolicy
	if len(authPolicy.Spec.Rules) == 0 {
		return nil
	}

	return authPolicy
}

func (t *translator) initializeAuthorizationPolicy(
	destination *discoveryv1.Destination,
) *securityv1beta1.AuthorizationPolicy {
	meta := metautils.TranslatedObjectMeta(
		destination.Spec.GetKubeService().Ref,
		destination.Annotations,
	)
	authPolicy := &securityv1beta1.AuthorizationPolicy{
		ObjectMeta: meta,
		Spec: securityv1beta1spec.AuthorizationPolicy{
			Selector: &typesv1beta1.WorkloadSelector{
				MatchLabels: destination.Spec.GetKubeService().WorkloadSelectorLabels,
			},
			Action: securityv1beta1spec.AuthorizationPolicy_ALLOW,
		},
	}
	return authPolicy
}

/*
	Translate an AccessPolicy instance into a Rule consisting of a Rule_From for each SourceSelector
	and a single Rule_To containing the rules specified in the AccessPolicy.
*/
func (t *translator) translateAccessPolicy(
	accessPolicy *v1.AccessPolicySpec,
	meshes discoveryv1sets.MeshSet,
) (*securityv1beta1spec.Rule, error) {
	var fromRules []*securityv1beta1spec.Rule_From
	for _, sourceSelector := range accessPolicy.SourceSelector {
		fromRule, err := t.buildSource(sourceSelector, meshes)
		if err != nil {
			return nil, err
		}
		fromRules = append(fromRules, fromRule)
	}
	toRules := buildToRules(accessPolicy)
	return &securityv1beta1spec.Rule{
		From: fromRules,
		To:   toRules,
	}, nil
}

func buildToRules(accessPolicy *v1.AccessPolicySpec) []*securityv1beta1spec.Rule_To {
	allowedPaths := accessPolicy.AllowedPaths
	allowedMethods := accessPolicy.AllowedMethods
	allowedPorts := convertIntsToStrings(accessPolicy.AllowedPorts)
	var ruleTo []*securityv1beta1spec.Rule_To
	if allowedPaths != nil || allowedMethods != nil || allowedPorts != nil {
		ruleTo = append(ruleTo, &securityv1beta1spec.Rule_To{
			Operation: &securityv1beta1spec.Operation{
				Paths:   accessPolicy.AllowedPaths,
				Methods: allowedMethods,
				Ports:   allowedPorts,
			},
		})
	}
	return ruleTo
}

// Generate all fully qualified principal names for specified service accounts.
// Reference: https://istio.io/docs/reference/config/security/authorization-policy/#Source
func (t *translator) buildSource(
	sources *commonv1.IdentitySelector,
	meshes discoveryv1sets.MeshSet,
) (*securityv1beta1spec.Rule_From, error) {
	if sources.GetKubeIdentityMatcher() == nil && sources.GetKubeServiceAccountRefs() == nil {
		// allow any source identity
		return &securityv1beta1spec.Rule_From{
			Source: &securityv1beta1spec.Source{},
		}, nil
	}
	// Select by identity matcher.
	wildcardPrincipals, namespaces, err := parseIdentityMatcher(sources.KubeIdentityMatcher, meshes)
	if err != nil {
		return nil, err
	}
	// Select by direct reference to ServiceAccounts
	serviceAccountPrincipals, err := parseServiceAccountRefs(sources.KubeServiceAccountRefs, meshes)
	if err != nil {
		return nil, err
	}

	return &securityv1beta1spec.Rule_From{
		Source: &securityv1beta1spec.Source{
			Principals: append(wildcardPrincipals, serviceAccountPrincipals...),
			Namespaces: namespaces,
		},
	}, nil
}

// Parse a list of principals and namespaces from a KubeIdentityMatcher.
func parseIdentityMatcher(
	kubeIdentityMatcher *commonv1.IdentitySelector_KubeIdentityMatcher,
	meshes discoveryv1sets.MeshSet,
) ([]string, []string, error) {
	var principals []string
	var namespaces []string
	if kubeIdentityMatcher == nil {
		return nil, nil, nil
	}
	if len(kubeIdentityMatcher.Clusters) > 0 {
		// select by clusters and specifiedNamespaces
		trustDomains, err := getTrustDomainsForClusters(kubeIdentityMatcher.Clusters, meshes)
		if err != nil {
			return nil, nil, err
		}
		specifiedNamespaces := kubeIdentityMatcher.Namespaces
		// Permit any namespace if unspecified.
		if len(specifiedNamespaces) == 0 {
			specifiedNamespaces = []string{""}
		}
		for _, trustDomain := range trustDomains {
			for _, namespace := range specifiedNamespaces {
				// Use empty string for service account to permit any.
				uri, err := buildSpiffeURI(trustDomain, namespace, "")
				if err != nil {
					return nil, nil, err
				}
				principals = append(principals, uri)
			}
		}
	} else if len(kubeIdentityMatcher.Namespaces) > 0 {
		// select by namespaces, permit any cluster
		namespaces = kubeIdentityMatcher.Namespaces
	}
	return principals, namespaces, nil
}

func parseServiceAccountRefs(
	kubeServiceAccountRefs *commonv1.IdentitySelector_KubeServiceAccountRefs,
	meshes discoveryv1sets.MeshSet,
) ([]string, error) {
	if kubeServiceAccountRefs == nil {
		return nil, nil
	}
	var principals []string
	for _, serviceAccountRef := range kubeServiceAccountRefs.ServiceAccounts {
		trustDomains, err := getTrustDomainsForClusters([]string{serviceAccountRef.ClusterName}, meshes)
		if err != nil {
			return nil, err
		}
		for _, trustDomain := range trustDomains {
			uri, err := buildSpiffeURI(trustDomain, serviceAccountRef.Namespace, serviceAccountRef.Name)
			if err != nil {
				return nil, err
			}
			principals = append(principals, uri)
		}
	}
	return principals, nil
}

/*
	Fetch trust domains for the Istio mesh of the given cluster.
	Multiple mesh installations of the same type on the same cluster are unsupported, simply use the first Mesh encountered.
*/
func getTrustDomainsForClusters(
	clusterNames []string,
	meshes discoveryv1sets.MeshSet,
) ([]string, error) {
	var errs *multierror.Error
	var trustDomains []string
	// omitted cluster name denotes all clusters
	if len(clusterNames) == 0 || len(clusterNames) == 1 && clusterNames[0] == "" {
		meshes.List(func(mesh *discoveryv1.Mesh) (_ bool) {
			if domain := mesh.Spec.GetIstio().GetTrustDomain(); domain != "" {
				trustDomains = append(trustDomains, domain)
			}
			return
		})
	} else {
		for _, clusterName := range clusterNames {
			trustDomain, err := getTrustDomainForCluster(clusterName, meshes)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			trustDomains = append(trustDomains, trustDomain)
		}
	}
	return trustDomains, errs.ErrorOrNil()
}

// Fetch trust domains by cluster so we can attribute missing trust domains to the problematic clusterName and report back to user.
func getTrustDomainForCluster(
	clusterName string,
	meshes discoveryv1sets.MeshSet,
) (string, error) {
	var trustDomain string
	for _, mesh := range meshes.List(func(mesh *discoveryv1.Mesh) bool {
		istio := mesh.Spec.GetIstio()
		return istio == nil || istio.GetInstallation().GetCluster() != clusterName
	}) {
		trustDomain = mesh.Spec.GetIstio().GetTrustDomain()
	}
	if trustDomain == "" {
		return "", trustDomainNotFound(clusterName)
	}
	return trustDomain, nil
}

/*
	The principal string format is described here: https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md#2-spiffe-identity
	Testing shows that the "spiffe://" prefix cannot be included.
	Istio only respects prefix or suffix wildcards, https://github.com/istio/istio/blob/9727308b3dadbfc8151cf70a045d1c7c52ab222b/pilot/pkg/security/authz/model/matcher/string.go#L45
*/
func buildSpiffeURI(trustDomain, namespace, serviceAccount string) (string, error) {
	if trustDomain == "" {
		return "", eris.New("trustDomain cannot be empty")
	}
	if namespace == "" {
		return fmt.Sprintf("%s/ns/*", trustDomain), nil
	} else if serviceAccount == "" {
		return fmt.Sprintf("%s/ns/%s/sa/*", trustDomain, namespace), nil
	} else {
		return fmt.Sprintf("%s/ns/%s/sa/%s", trustDomain, namespace, serviceAccount), nil
	}
}

func convertIntsToStrings(ints []uint32) []string {
	var strings []string
	for _, i := range ints {
		strings = append(strings, strconv.Itoa(int(i)))
	}
	return strings
}
