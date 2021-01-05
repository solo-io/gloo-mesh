
---

title: "failover_service.proto"

---

## Package : `networking.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for failover_service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## failover_service.proto


## Table of Contents
  - [FailoverServiceSpec](#networking.mesh.gloo.solo.io.FailoverServiceSpec)
  - [FailoverServiceSpec.BackingService](#networking.mesh.gloo.solo.io.FailoverServiceSpec.BackingService)
  - [FailoverServiceSpec.Port](#networking.mesh.gloo.solo.io.FailoverServiceSpec.Port)
  - [FailoverServiceStatus](#networking.mesh.gloo.solo.io.FailoverServiceStatus)
  - [FailoverServiceStatus.MeshesEntry](#networking.mesh.gloo.solo.io.FailoverServiceStatus.MeshesEntry)







<a name="networking.mesh.gloo.solo.io.FailoverServiceSpec"></a>

### FailoverServiceSpec
A FailoverService creates a new hostname to which services can send requests. Requests will be routed based on a list of backing traffic targets ordered by decreasing priority. When outlier detection detects that a traffic target in the list is in an unhealthy state, requests sent to the FailoverService will be routed to the next healthy traffic target in the list. For each traffic target referenced in the FailoverService's BackingServices list, outlier detection must be configured using a TrafficPolicy.<br>Currently this feature only supports Services backed by Istio.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hostname | string |  | The DNS name of the FailoverService. Must be unique within the service mesh instance since it is used as the hostname with which clients communicate. |
  | port | [networking.mesh.gloo.solo.io.FailoverServiceSpec.Port]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.failover_service#networking.mesh.gloo.solo.io.FailoverServiceSpec.Port" >}}) |  | The port on which the FailoverService listens. |
  | meshes | [][core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) | repeated | The meshes that this FailoverService will be visible to. |
  | backingServices | [][networking.mesh.gloo.solo.io.FailoverServiceSpec.BackingService]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.failover_service#networking.mesh.gloo.solo.io.FailoverServiceSpec.BackingService" >}}) | repeated | The list of services backing the FailoverService, ordered by decreasing priority. All services must be backed by either the same service mesh instance or backed by service meshes that are grouped under a common VirtualMesh. |
  





<a name="networking.mesh.gloo.solo.io.FailoverServiceSpec.BackingService"></a>

### FailoverServiceSpec.BackingService
The traffic targets that comprise the FailoverService.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kubeService | [core.skv2.solo.io.ClusterObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ClusterObjectRef" >}}) |  | Name/namespace/cluster of a kubernetes service. |
  





<a name="networking.mesh.gloo.solo.io.FailoverServiceSpec.Port"></a>

### FailoverServiceSpec.Port
The port on which the FailoverService listens.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | uint32 |  | Port number. |
  | protocol | string |  | Protocol of the requests sent to the FailoverService, must be one of HTTP, HTTPS, GRPC, HTTP2, MONGO, TCP, TLS. |
  





<a name="networking.mesh.gloo.solo.io.FailoverServiceStatus"></a>

### FailoverServiceStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  | The most recent generation observed in the the FailoverService metadata. If the observedGeneration does not match generation, the controller has not received the most recent version of this resource. |
  | state | [networking.mesh.gloo.solo.io.ApprovalState]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.validation_state#networking.mesh.gloo.solo.io.ApprovalState" >}}) |  | The state of the overall resource, will only show accepted if it has been successfully applied to all target meshes. |
  | meshes | [][networking.mesh.gloo.solo.io.FailoverServiceStatus.MeshesEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.failover_service#networking.mesh.gloo.solo.io.FailoverServiceStatus.MeshesEntry" >}}) | repeated | The status of the FailoverService for each Mesh to which it has been applied. |
  | errors | []string | repeated | Any errors found while processing this generation of the resource. |
  





<a name="networking.mesh.gloo.solo.io.FailoverServiceStatus.MeshesEntry"></a>

### FailoverServiceStatus.MeshesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | [networking.mesh.gloo.solo.io.ApprovalStatus]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.validation_state#networking.mesh.gloo.solo.io.ApprovalStatus" >}}) |  |  |
  




 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

