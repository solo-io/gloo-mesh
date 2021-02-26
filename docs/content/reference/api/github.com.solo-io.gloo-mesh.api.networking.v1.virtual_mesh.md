
---

title: "virtual_mesh.proto"

---

## Package : `networking.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for virtual_mesh.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## virtual_mesh.proto


## Table of Contents
  - [VirtualMeshSpec](#networking.mesh.gloo.solo.io.VirtualMeshSpec)
  - [VirtualMeshSpec.Federation](#networking.mesh.gloo.solo.io.VirtualMeshSpec.Federation)
  - [VirtualMeshSpec.MTLSConfig](#networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig)
  - [VirtualMeshSpec.MTLSConfig.LimitedTrust](#networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.LimitedTrust)
  - [VirtualMeshSpec.MTLSConfig.SharedTrust](#networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.SharedTrust)
  - [VirtualMeshSpec.RootCertificateAuthority](#networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority)
  - [VirtualMeshSpec.RootCertificateAuthority.SelfSignedCert](#networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority.SelfSignedCert)
  - [VirtualMeshStatus](#networking.mesh.gloo.solo.io.VirtualMeshStatus)
  - [VirtualMeshStatus.MeshesEntry](#networking.mesh.gloo.solo.io.VirtualMeshStatus.MeshesEntry)

  - [VirtualMeshSpec.GlobalAccessPolicy](#networking.mesh.gloo.solo.io.VirtualMeshSpec.GlobalAccessPolicy)






<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec"></a>

### VirtualMeshSpec
Represents a logical grouping of Meshes for shared configuration and cross-mesh interoperability.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| meshes | [][core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) | repeated | Specify the Meshes configured by this VirtualMesh. |
  | mtlsConfig | [networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig" >}}) |  | Specify mTLS options. |
  | federation | [networking.mesh.gloo.solo.io.VirtualMeshSpec.Federation]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.Federation" >}}) |  | Specify how to federate Destinations across service mesh boundaries. |
  | globalAccessPolicy | [networking.mesh.gloo.solo.io.VirtualMeshSpec.GlobalAccessPolicy]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.GlobalAccessPolicy" >}}) |  | Specify a global access policy for all Workloads and Destinations associated with this VirtualMesh. |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.Federation"></a>

### VirtualMeshSpec.Federation
"Federation" refers to the ability to expose Destinations across service mesh boundaries, i.e. to traffic originating from Workloads external to the Destination's Mesh.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissive | [google.protobuf.Empty]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.protoc-gen-ext.external.google.protobuf.empty#google.protobuf.Empty" >}}) |  | Expose all Destinations to all Workloads in this VirtualMesh. |
  | flatNetwork | bool |  | If true, all multicluster traffic will be routed directly to the Kubernetes service endpoints of the Destinations, rather than through an ingress gateway. This mode requires a flat network environment. |
  | hostnameSuffix | string |  | Configure the suffix for hostnames of Destinations federated within this VirtualMesh. Currently this is only supported for Istio with [smart DNS proxying enabled](https://istio.io/latest/blog/2020/dns-proxy/), otherwise setting this field results in an error. If omitted, the hostname suffix defaults to "global". |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig"></a>

### VirtualMeshSpec.MTLSConfig
Specify mTLS options. This includes options for configuring Mutual TLS within an individual mesh, as well as enabling mTLS across Meshes by establishing cross-mesh trust.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shared | [networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.SharedTrust]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.SharedTrust" >}}) |  | Shared trust (allow communication between any pair of Workloads and Destinations in the grouped Meshes). |
  | limited | [networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.LimitedTrust]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.LimitedTrust" >}}) |  | Limited trust (selectively allow communication between Workloads and Destinations in the grouped Meshes). *Currently not available.* |
  | autoRestartPods | bool |  | Specify whether to allow Gloo Mesh to restart Kubernetes Pods when certificates are rotated when establishing shared trust. If this option is not explicitly enabled, users must restart Pods manually for the new certificates to be picked up. `meshctl` provides the command `meshctl mesh restart` to simplify this process, see [here]({{< versioned_link_path fromRoot="reference/cli/meshctl_mesh_restart/" >}}) for more info. |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.LimitedTrust"></a>

### VirtualMeshSpec.MTLSConfig.LimitedTrust
Limited trust is a trust model which does not require trusting Meshes to share the same root certificate or identity. Instead, trust is established between different Meshes by connecting their ingress/egress gateways with a common certificate/identity. In this model all requests between different have the following request path when communicating between clusters ```                cluster 1 MTLS               shared MTLS                  cluster 2 MTLS client/workload <-----------> egress gateway <----------> ingress gateway <--------------> server ``` This approach has the downside of not maintaining identity from client to server, but allows for ad-hoc addition of additional Meshes into a VirtualMesh.






<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.MTLSConfig.SharedTrust"></a>

### VirtualMeshSpec.MTLSConfig.SharedTrust
Shared trust is a trust model requiring a common root certificate shared between trusting Meshes, as well as shared identity between all Workloads and Destinations which wish to communicate within the VirtualMesh.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rootCertificateAuthority | [networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority" >}}) |  | Configure a Root Certificate Authority which will be shared by all Meshes associated with this VirtualMesh. If this is not provided, a self-signed certificate will be generated by Gloo Mesh. |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority"></a>

### VirtualMeshSpec.RootCertificateAuthority
Specify parameters for configuring the root certificate authority for a VirtualMesh.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| generated | [networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority.SelfSignedCert]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority.SelfSignedCert" >}}) |  | Generate a self-signed root certificate with the given options. |
  | secret | [core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) |  | Reference to a Kubernetes Secret containing the root certificate authority. Provided secrets must conform to a specified format, [documented here]({{% versioned_link_path fromRoot="/guides/federate_identity/" %}}). |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.RootCertificateAuthority.SelfSignedCert"></a>

### VirtualMeshSpec.RootCertificateAuthority.SelfSignedCert
Configuration for generating a self-signed root certificate. Uses the X.509 format, RFC5280.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ttlDays | uint32 |  | Number of days before root cert expires. Defaults to 365. |
  | rsaKeySizeBytes | uint32 |  | Size in bytes of the root cert's private key. Defaults to 4096. |
  | orgName | string |  | Root cert organization name. Defaults to "gloo-mesh". |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshStatus"></a>

### VirtualMeshStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  | The most recent generation observed in the the VirtualMesh metadata. If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most recent version of this resource. |
  | state | [common.mesh.gloo.solo.io.ApprovalState]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.common.v1.validation_state#common.mesh.gloo.solo.io.ApprovalState" >}}) |  | The state of the overall resource. It will only show accepted if it has been successfully applied to all selected Meshes. |
  | errors | []string | repeated | Any errors found while processing this generation of the resource. |
  | meshes | [][networking.mesh.gloo.solo.io.VirtualMeshStatus.MeshesEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh#networking.mesh.gloo.solo.io.VirtualMeshStatus.MeshesEntry" >}}) | repeated | The status of the VirtualMesh for each Mesh to which it has been applied. A VirtualMesh may be Accepted for some Meshes and rejected for others. |
  





<a name="networking.mesh.gloo.solo.io.VirtualMeshStatus.MeshesEntry"></a>

### VirtualMeshStatus.MeshesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | [networking.mesh.gloo.solo.io.ApprovalStatus]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.status#networking.mesh.gloo.solo.io.ApprovalStatus" >}}) |  |  |
  




 <!-- end messages -->


<a name="networking.mesh.gloo.solo.io.VirtualMeshSpec.GlobalAccessPolicy"></a>

### VirtualMeshSpec.GlobalAccessPolicy
Specify a global access policy for all Workloads and Destinations associated with this VirtualMesh.

| Name | Number | Description |
| ---- | ------ | ----------- |
| MESH_DEFAULT | 0 | Assume the default for the service mesh type. Istio defaults to `false`, App Mesh defaults to `true`. |
| ENABLED | 1 | Disallow traffic to all Destinations in the VirtualMesh unless explicitly allowed through [AccessPolicies]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.access_policy/" >}}). |
| DISABLED | 2 | Allow traffic to all Destinations in the VirtualMesh unless explicitly disallowed through [AccessPolicies]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.access_policy/" >}}). |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

