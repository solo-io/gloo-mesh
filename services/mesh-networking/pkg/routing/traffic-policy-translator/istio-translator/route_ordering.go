package istio_translator

import (
	smh_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	istio_networking "istio.io/api/networking/v1alpha3"
)

type SpecificitySortableRoutes []*istio_networking.HTTPRoute

func (b SpecificitySortableRoutes) Len() int {
	return len(b)
}

func (b SpecificitySortableRoutes) Less(i, j int) bool {
	// if the first HTTPRoute matches more specific criteria than the second HTTPRoute,
	// order them such that the first precedes the second (i.e. takes precedence over)
	return isHttpRouteMatcherMoreSpecific(b[i], b[j])
}

func (b SpecificitySortableRoutes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

/* Order decreasing by specificity of matcher. Specifically this means ordering by the following matcher fields.
Note that the precedence of the matcher fields listed below is somewhat arbitrary (e.g. matching on query parameters
is less specific than matching on headers), and the precedence of the fields can be safely rearranged if the resulting
ordering needs to be changed.

1. Headers, number of items decreasing
2. QueryParams, number of items decreasing
3. SourceLabels, number of items decreasing
4. SourceNamespace, alphabetical decreasing
5. Uri, according to StringMatch specificity defined below
6. Method, according to ordering declared in HTTPMethodOrdering
7. WithoutHeaders, number of items decreasing
*/
func isHttpRouteMatcherMoreSpecific(httpRouteA, httpRouteB *istio_networking.HTTPRoute) bool {
	// each HttpRoute is guaranteed to only have a single HttpMatchRequest
	a := httpRouteA.GetMatch()[0]
	b := httpRouteB.GetMatch()[0]
	if len(a.GetHeaders()) > len(b.GetHeaders()) {
		return true
	} else if len(a.GetHeaders()) < len(b.GetHeaders()) {
		return false
	}
	if len(a.GetQueryParams()) > len(b.GetQueryParams()) {
		return true
	} else if len(a.GetQueryParams()) < len(b.GetQueryParams()) {
		return false
	}
	if len(a.GetSourceLabels()) > len(b.GetSourceLabels()) {
		return true
	} else if len(a.GetSourceLabels()) < len(b.GetSourceLabels()) {
		return false
	}
	if a.GetSourceNamespace() > b.GetSourceNamespace() {
		return true
	} else if a.GetSourceNamespace() < b.GetSourceNamespace() {
		return false
	}
	if isStringMatchMoreSpecific(a.GetUri(), b.GetUri()) {
		return true
	} else if isStringMatchMoreSpecific(b.GetUri(), a.GetUri()) {
		return false
	}
	if isMethodMoreSpecific(a.GetMethod(), b.GetMethod()) {
		return true
	} else if isMethodMoreSpecific(b.GetMethod(), a.GetMethod()) {
		return false
	}
	if len(a.GetWithoutHeaders()) > len(b.GetWithoutHeaders()) {
		return true
	} else if len(a.GetWithoutHeaders()) < len(b.GetWithoutHeaders()) {
		return false
	}
	return false
}

// In decreasing order of specificity: exact, regex, prefix
// If both StringMatch objects are of same type, compare by length (longer being more specific)
func isStringMatchMoreSpecific(a, b *istio_networking.StringMatch) bool {
	if len(a.GetExact()) > len(b.GetExact()) {
		return true
	} else if len(a.GetExact()) < len(b.GetExact()) {
		return false
	}
	// the notion of specificity doesn't apply to this regex string ordering, but this is needed for determinism
	if len(a.GetRegex()) > len(b.GetRegex()) {
		return true
	} else if len(a.GetRegex()) < len(b.GetRegex()) {
		return false
	}
	if len(a.GetPrefix()) > len(b.GetPrefix()) {
		return true
	} else if len(a.GetPrefix()) < len(b.GetPrefix()) {
		return false
	}
	return false
}

// SMH API currently only supports exact method matches
func isMethodMoreSpecific(a, b *istio_networking.StringMatch) bool {
	if a.GetExact() == "" {
		return false
	} else if b.GetExact() == "" {
		return true
	} else {
		// if both method matchers exist, ordering by specificity is irrelevant because they match in a mutually exclusively manner
		// we apply an ordering here purely for determinism
		return smh_core_types.HttpMethodValue_value[a.GetExact()] > smh_core_types.HttpMethodValue_value[b.GetExact()]
	}
}
