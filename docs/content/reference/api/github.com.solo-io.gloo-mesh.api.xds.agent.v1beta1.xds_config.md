
---

title: "xds_config.proto"

---

## Package : `xds.agent.enterprise.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for xds_config.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## xds_config.proto


## Table of Contents
  - [XdsConfigSpec](#xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec)
  - [XdsConfigSpec.Resource](#xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.Resource)
  - [XdsConfigSpec.TypedResources](#xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.TypedResources)
  - [XdsConfigStatus](#xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigStatus)







<a name="xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec"></a>

### XdsConfigSpec
XdsConfigs are used to issue xDS Configuration Resources to running Envoy instances. They are created by Gloo Mesh for processing by an agent running on managed clusters.<br>The agent will serve the specified xDS configuration resources on its grpc-xds port (default 9977) to the Envoy instances (nodes) defined in the XDSConfigSpec.<br>This feature is currently only available in Gloo Mesh Enterprise.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| workloads | [][core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) | repeated | The Workloads that will receive this xDS Configuration. |
  | types | [][xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.TypedResources]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.xds.agent.v1beta1.xds_config#xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.TypedResources" >}}) | repeated | The xDS resources to serve to the nodes. Mapped by type URL. |
  





<a name="xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.Resource"></a>

### XdsConfigSpec.Resource
A single named resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | string |  | Name of the resource as referenced by xDS. |
  | compressedData | bytes |  | Stored as compressed, base-64 encoded raw bytes. |
  





<a name="xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.TypedResources"></a>

### XdsConfigSpec.TypedResources
A set of resources of a single type (typeURL).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| typeUrl | string |  | The type URL of the resources in the given set. |
  | resources | [][xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.Resource]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.xds.agent.v1beta1.xds_config#xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigSpec.Resource" >}}) | repeated | Stored as compressed, base-64 encoded raw bytes. |
  





<a name="xds.agent.enterprise.mesh.gloo.solo.io.XdsConfigStatus"></a>

### XdsConfigStatus
The XdsConfig status is written by the CertificateRequesting agent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | int64 |  | The most recent generation observed in the the XdsConfig metadata. If the `observedGeneration` does not match `metadata.generation`, the Gloo Mesh agent has not processed the most recent version of this XdsConfig. |
  | error | string |  | Any error observed which prevented the XdsConfig from being processed. If the error is empty, the request has been processed successfully. |
  




 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

