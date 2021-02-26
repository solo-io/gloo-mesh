
---

title: "traffic_policy.proto"

---

## Package : `networking.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for traffic_policy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## traffic_policy.proto


## Table of Contents
  - [TrafficPolicySpec](#networking.mesh.gloo.solo.io.TrafficPolicySpec)
  - [TrafficPolicySpec.HttpMatcher](#networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher)
  - [TrafficPolicySpec.HttpMatcher.QueryParameterMatcher](#networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher.QueryParameterMatcher)
  - [TrafficPolicySpec.Policy](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy)
  - [TrafficPolicySpec.Policy.CorsPolicy](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy)
  - [TrafficPolicySpec.Policy.CorsPolicy.StringMatch](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy.StringMatch)
  - [TrafficPolicySpec.Policy.FaultInjection](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection)
  - [TrafficPolicySpec.Policy.FaultInjection.Abort](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection.Abort)
  - [TrafficPolicySpec.Policy.HeaderManipulation](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation)
  - [TrafficPolicySpec.Policy.HeaderManipulation.AppendRequestHeadersEntry](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendRequestHeadersEntry)
  - [TrafficPolicySpec.Policy.HeaderManipulation.AppendResponseHeadersEntry](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendResponseHeadersEntry)
  - [TrafficPolicySpec.Policy.MTLS](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS)
  - [TrafficPolicySpec.Policy.MTLS.Istio](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio)
  - [TrafficPolicySpec.Policy.Mirror](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.Mirror)
  - [TrafficPolicySpec.Policy.MultiDestination](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination)
  - [TrafficPolicySpec.Policy.MultiDestination.WeightedDestination](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination)
  - [TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination)
  - [TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination.SubsetEntry](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination.SubsetEntry)
  - [TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference)
  - [TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference.SubsetEntry](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference.SubsetEntry)
  - [TrafficPolicySpec.Policy.OutlierDetection](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.OutlierDetection)
  - [TrafficPolicySpec.Policy.RetryPolicy](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.RetryPolicy)
  - [TrafficPolicyStatus](#networking.mesh.gloo.solo.io.TrafficPolicyStatus)
  - [TrafficPolicyStatus.DestinationsEntry](#networking.mesh.gloo.solo.io.TrafficPolicyStatus.DestinationsEntry)

  - [TrafficPolicySpec.Policy.MTLS.Istio.TLSmode](#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio.TLSmode)






<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec"></a>

### TrafficPolicySpec
Applies L7 routing and post-routing configuration on selected network edges.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sourceSelector | [][common.mesh.gloo.solo.io.WorkloadSelector]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.common.v1.selectors#common.mesh.gloo.solo.io.WorkloadSelector" >}}) | repeated | Specify the Workloads (traffic sources) this TrafficPolicy applies to. Omit to apply to all Workloads. |
  | destinationSelector | [][common.mesh.gloo.solo.io.DestinationSelector]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.common.v1.selectors#common.mesh.gloo.solo.io.DestinationSelector" >}}) | repeated | Specify the Destinations (destinations) this TrafficPolicy applies to. Omit to apply to all Destinations. |
  | httpRequestMatchers | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher" >}}) | repeated | Specify criteria that HTTP requests must satisfy for the TrafficPolicy to apply. Conditions defined within a single matcher are conjunctive, i.e. all conditions must be satisfied for a match to occur. Conditions defined between different matchers are disjunctive, i.e. at least one matcher must be satisfied for the TrafficPolicy to apply. Omit to apply to any HTTP request. |
  | policy | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy" >}}) |  | Specify L7 routing and post-routing configuration. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher"></a>

### TrafficPolicySpec.HttpMatcher
Specify HTTP request level match criteria. All specified conditions must be satisfied for a match to occur.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| prefix | string |  | If specified, the targeted path must begin with the prefix. |
  | exact | string |  | If specified, the targeted path must exactly match the value. |
  | regex | string |  | If specified, the targeted path must match the regex. |
  | headers | [][common.mesh.gloo.solo.io.HeaderMatcher]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.common.v1.request_matchers#common.mesh.gloo.solo.io.HeaderMatcher" >}}) | repeated | Specify a set of headers which requests must match in entirety (all headers must match). |
  | queryParameters | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher.QueryParameterMatcher]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher.QueryParameterMatcher" >}}) | repeated | Specify a set of URL query parameters which requests must match in entirety (all query params must match). |
  | method | string |  | Specify an HTTP method to match against. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.HttpMatcher.QueryParameterMatcher"></a>

### TrafficPolicySpec.HttpMatcher.QueryParameterMatcher
Specify match criteria against the target URL's query parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | string |  | Specify the name of a key that must be present in the requested path's query string. |
  | value | string |  | Specify the value of the query parameter keyed on `name`. |
  | regex | bool |  | If true, treat `value` as a regular expression. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy"></a>

### TrafficPolicySpec.Policy
Specify L7 routing and post-routing configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trafficShift | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination" >}}) |  | Shift traffic to a different destination. |
  | faultInjection | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection" >}}) |  | Inject faulty responses. |
  | requestTimeout | [google.protobuf.Duration]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.duration#google.protobuf.Duration" >}}) |  | Set a timeout on requests. |
  | retries | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.RetryPolicy]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.RetryPolicy" >}}) |  | Set a retry policy on requests. |
  | corsPolicy | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy" >}}) |  | Set a Cross-Origin Resource Sharing policy (CORS) for requests. Refer to [this link](https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS) for further details about cross origin resource sharing. |
  | mirror | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.Mirror]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.Mirror" >}}) |  | Mirror traffic to a another destination (traffic will be sent to its original destination in addition to the mirrored destinations). |
  | headerManipulation | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation" >}}) |  | Manipulate request and response headers. |
  | outlierDetection | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.OutlierDetection]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.OutlierDetection" >}}) |  | Configure [outlier detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/outlier) on the selected destinations. Specifying this field requires an empty `source_selector` because it must apply to all traffic. |
  | mtls | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS" >}}) |  | Configure mTLS settings. If specified will override global default defined in Settings. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy"></a>

### TrafficPolicySpec.Policy.CorsPolicy
Specify Cross-Origin Resource Sharing policy (CORS) for requests. Refer to [this link](https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS) for further details about cross origin resource sharing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allowOrigins | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy.StringMatch]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy.StringMatch" >}}) | repeated | String patterns that match allowed origins. An origin is allowed if any of the string matchers match. |
  | allowMethods | []string | repeated | List of HTTP methods allowed to access the resource. The content will be serialized to the `Access-Control-Allow-Methods` header. |
  | allowHeaders | []string | repeated | List of HTTP headers that can be used when requesting the resource. Serialized to the `Access-Control-Allow-Headers` header. |
  | exposeHeaders | []string | repeated | A list of HTTP headers that browsers are allowed to access. Serialized to the `Access-Control-Expose-Headers` header. |
  | maxAge | [google.protobuf.Duration]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.duration#google.protobuf.Duration" >}}) |  | Specify how long the results of a preflight request can be cached. Serialized to the `Access-Control-Max-Age` header. |
  | allowCredentials | [google.protobuf.BoolValue]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.wrappers#google.protobuf.BoolValue" >}}) |  | Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials. Translates to the `Access-Control-Allow-Credentials` header. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.CorsPolicy.StringMatch"></a>

### TrafficPolicySpec.Policy.CorsPolicy.StringMatch
Describes how to match a given string in HTTP headers. Match is case-sensitive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exact | string |  | Exact string match. |
  | prefix | string |  | Prefix-based match. |
  | regex | string |  | ECMAscript style regex-based match. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection"></a>

### TrafficPolicySpec.Policy.FaultInjection
Specify one or more faults to inject for the selected network edge.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fixedDelay | [google.protobuf.Duration]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.duration#google.protobuf.Duration" >}}) |  | Add a delay of a fixed duration before sending the request. Format: `1h`/`1m`/`1s`/`1ms`. MUST be >=1ms. |
  | abort | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection.Abort]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection.Abort" >}}) |  | Abort the request and return the specified error code back to traffic source. |
  | percentage | double |  | Percentage of requests to be faulted. Values range between 0 and 100. If omitted all requests will be faulted. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.FaultInjection.Abort"></a>

### TrafficPolicySpec.Policy.FaultInjection.Abort
Abort the request and return the specified error code back to traffic source.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| httpStatus | int32 |  | Required. HTTP status code to use to abort the request. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation"></a>

### TrafficPolicySpec.Policy.HeaderManipulation
Specify modifications to request and response headers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| removeResponseHeaders | []string | repeated | HTTP headers to remove before returning a response to the caller. |
  | appendResponseHeaders | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendResponseHeadersEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendResponseHeadersEntry" >}}) | repeated | Additional HTTP headers to add before returning a response to the caller. |
  | removeRequestHeaders | []string | repeated | HTTP headers to remove before forwarding a request to the destination service. |
  | appendRequestHeaders | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendRequestHeadersEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendRequestHeadersEntry" >}}) | repeated | Additional HTTP headers to add before forwarding a request to the destination service. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendRequestHeadersEntry"></a>

### TrafficPolicySpec.Policy.HeaderManipulation.AppendRequestHeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | string |  |  |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.HeaderManipulation.AppendResponseHeadersEntry"></a>

### TrafficPolicySpec.Policy.HeaderManipulation.AppendResponseHeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | string |  |  |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS"></a>

### TrafficPolicySpec.Policy.MTLS
Configure mTLS settings on destinations. If specified this overrides the global default defined in Settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| istio | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio" >}}) |  | Istio TLS settings. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio"></a>

### TrafficPolicySpec.Policy.MTLS.Istio
Istio TLS settings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tlsMode | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio.TLSmode]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio.TLSmode" >}}) |  | TLS connection mode |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.Mirror"></a>

### TrafficPolicySpec.Policy.Mirror
Mirror traffic to a another destination (traffic will be sent to its original destination in addition to the mirrored destinations).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubeService | [core.skv2.solo.io.ClusterObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ClusterObjectRef" >}}) |  | Reference (name, namespace, Gloo Mesh cluster) to a Kubernetes service. |
  | percentage | double |  | Percentage of traffic to mirror. If omitted all traffic will be mirrored. Values must be between 0 and 100. |
  | port | uint32 |  | Port on the destination to receive traffic. Required if the destination exposes multiple ports. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination"></a>

### TrafficPolicySpec.Policy.MultiDestination
Specify a traffic shift destination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destinations | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination" >}}) | repeated | Specify weighted traffic shift destinations. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination"></a>

### TrafficPolicySpec.Policy.MultiDestination.WeightedDestination
Specify a traffic shift destination along with a weight.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| weight | uint32 |  | Specify the proportion of traffic to be forwarded to this destination. Weights across all of the `destinations` must sum to 100. |
  | kubeService | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination" >}}) |  | Specify a Kubernetes Service. |
  | virtualDestination | [networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference" >}}) |  | Specify a VirtualDestination. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination"></a>

### TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination
A Kubernetes destination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | string |  | The name of the service. |
  | namespace | string |  | The namespace of the service. |
  | clusterName | string |  | The Gloo Mesh cluster name (registration name) of the service. |
  | subset | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination.SubsetEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination.SubsetEntry" >}}) | repeated | Specify, by labels, a subset of service instances to route to. |
  | port | uint32 |  | Port on the service to receive traffic. Required if the service exposes more than one port. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination.SubsetEntry"></a>

### TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.KubeDestination.SubsetEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | string |  |  |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference"></a>

### TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference
Specify a VirtualDestination traffic shift destination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | string |  | The name of the VirtualDestination object. |
  | namespace | string |  | The namespace of the VirtualDestination object. |
  | subset | [][networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference.SubsetEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference.SubsetEntry" >}}) | repeated | Specify, by labels, a subset of service instances backing the VirtualDestination to route to. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference.SubsetEntry"></a>

### TrafficPolicySpec.Policy.MultiDestination.WeightedDestination.VirtualDestinationReference.SubsetEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | string |  |  |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.OutlierDetection"></a>

### TrafficPolicySpec.Policy.OutlierDetection
Configure [outlier detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/outlier) on the selected destinations. Specifying this field requires an empty `source_selector` because it must apply to all traffic.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consecutiveErrors | uint32 |  | The number of errors before a host is ejected from the connection pool. A default will be used if not set. |
  | interval | [google.protobuf.Duration]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.duration#google.protobuf.Duration" >}}) |  | The time interval between ejection sweep analysis. Format: `1h`/`1m`/`1s`/`1ms`. Must be >= `1ms`. A default will be used if not set. |
  | baseEjectionTime | [google.protobuf.Duration]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.duration#google.protobuf.Duration" >}}) |  | The minimum ejection duration. Format: `1h`/`1m`/`1s`/`1ms`. Must be >= `1ms`. A default will be used if not set. |
  | maxEjectionPercent | uint32 |  | The maximum percentage of hosts that can be ejected from the load balancing pool. At least one host will be ejected regardless of the value. Must be between 0 and 100. A default will be used if not set. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.RetryPolicy"></a>

### TrafficPolicySpec.Policy.RetryPolicy
Specify retries for failed requests.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attempts | int32 |  | Number of retries for a given request |
  | perTryTimeout | [google.protobuf.Duration]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.duration#google.protobuf.Duration" >}}) |  | Timeout per retry attempt for a given request. Format: `1h`/`1m`/`1s`/`1ms`. *Must be >= 1ms*. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicyStatus"></a>

### TrafficPolicyStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  | The most recent generation observed in the the TrafficPolicy metadata. If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most recent version of this resource. |
  | state | [common.mesh.gloo.solo.io.ApprovalState]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.common.v1.validation_state#common.mesh.gloo.solo.io.ApprovalState" >}}) |  | The state of the overall resource. It will only show accepted if it has been successfully applied to all selected Destinations. |
  | destinations | [][networking.mesh.gloo.solo.io.TrafficPolicyStatus.DestinationsEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.traffic_policy#networking.mesh.gloo.solo.io.TrafficPolicyStatus.DestinationsEntry" >}}) | repeated | The status of the TrafficPolicy for each selected Destination. A TrafficPolicy may be Accepted for some Destinations and rejected for others. |
  | workloads | []string | repeated | The list of selected Workloads for which this policy has been applied. |
  | errors | []string | repeated | Any errors found while processing this generation of the resource. |
  





<a name="networking.mesh.gloo.solo.io.TrafficPolicyStatus.DestinationsEntry"></a>

### TrafficPolicyStatus.DestinationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | [networking.mesh.gloo.solo.io.ApprovalStatus]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.status#networking.mesh.gloo.solo.io.ApprovalStatus" >}}) |  |  |
  




 <!-- end messages -->


<a name="networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS.Istio.TLSmode"></a>

### TrafficPolicySpec.Policy.MTLS.Istio.TLSmode
TLS connection mode. Enums correspond to those [defined here](https://github.com/istio/api/blob/00636152b9d9254b614828a65723840282a177d3/networking/v1beta1/destination_rule.proto#L886)

| Name | Number | Description |
| ---- | ------ | ----------- |
| DISABLE | 0 | Do not originate a TLS connection to the upstream endpoint. |
| SIMPLE | 1 | Originate a TLS connection to the upstream endpoint. |
| ISTIO_MUTUAL | 2 | Secure connections to the upstream using mutual TLS by presenting client certificates for authentication. This mode uses certificates generated automatically by Istio for mTLS authentication. When this mode is used, all other fields in `ClientTLSSettings` should be empty. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

