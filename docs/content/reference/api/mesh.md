
---
title: "mesh.proto"
---

## Package : `discovery.smh.solo.io`



<a name="top"></a>

<a name="API Reference for mesh.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mesh.proto


## Table of Contents
  - [MeshSpec](#discovery.smh.solo.io.MeshSpec)
  - [MeshSpec.AgentInfo](#discovery.smh.solo.io.MeshSpec.AgentInfo)
  - [MeshSpec.AwsAppMesh](#discovery.smh.solo.io.MeshSpec.AwsAppMesh)
  - [MeshSpec.ConsulConnectMesh](#discovery.smh.solo.io.MeshSpec.ConsulConnectMesh)
  - [MeshSpec.Istio](#discovery.smh.solo.io.MeshSpec.Istio)
  - [MeshSpec.Istio.CitadelInfo](#discovery.smh.solo.io.MeshSpec.Istio.CitadelInfo)
  - [MeshSpec.Istio.IngressGatewayInfo](#discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo)
  - [MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry](#discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry)
  - [MeshSpec.LinkerdMesh](#discovery.smh.solo.io.MeshSpec.LinkerdMesh)
  - [MeshSpec.MeshInstallation](#discovery.smh.solo.io.MeshSpec.MeshInstallation)
  - [MeshSpec.MeshInstallation.PodLabelsEntry](#discovery.smh.solo.io.MeshSpec.MeshInstallation.PodLabelsEntry)
  - [MeshSpec.OSM](#discovery.smh.solo.io.MeshSpec.OSM)
  - [MeshStatus](#discovery.smh.solo.io.MeshStatus)
  - [MeshStatus.AppliedFailoverService](#discovery.smh.solo.io.MeshStatus.AppliedFailoverService)
  - [MeshStatus.AppliedVirtualMesh](#discovery.smh.solo.io.MeshStatus.AppliedVirtualMesh)







<a name="discovery.smh.solo.io.MeshSpec"></a>

### MeshSpec
Meshes represent a currently registered service mesh.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| istio | [MeshSpec.Istio](#discovery.smh.solo.io.MeshSpec.Istio) |  |  |
| awsAppMesh | [MeshSpec.AwsAppMesh](#discovery.smh.solo.io.MeshSpec.AwsAppMesh) |  |  |
| linkerd | [MeshSpec.LinkerdMesh](#discovery.smh.solo.io.MeshSpec.LinkerdMesh) |  |  |
| consulConnect | [MeshSpec.ConsulConnectMesh](#discovery.smh.solo.io.MeshSpec.ConsulConnectMesh) |  |  |
| osm | [MeshSpec.OSM](#discovery.smh.solo.io.MeshSpec.OSM) |  |  |
| agentInfo | [MeshSpec.AgentInfo](#discovery.smh.solo.io.MeshSpec.AgentInfo) |  | Information about the SMH certificate agent if it has been installed to the remote cluster. |






<a name="discovery.smh.solo.io.MeshSpec.AgentInfo"></a>

### MeshSpec.AgentInfo
information about the SMH Cert-Agent which may be installed to the remote cluster which contains the Mesh control plane


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| agentNamespace | [string](#string) |  | the namespace in which the agent is installed |






<a name="discovery.smh.solo.io.MeshSpec.AwsAppMesh"></a>

### MeshSpec.AwsAppMesh
Mesh object representing AWS AppMesh.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | AWS name for the AppMesh instance, must be unique across the AWS account. |
| region | [string](#string) |  | The AWS region the AWS App Mesh control plane resources exist in. |
| awsAccountId | [string](#string) |  | The AWS Account ID associated with the Mesh. Populated at REST API registration time. |
| clusters | [][string](#string) | repeated | The k8s clusters on which sidecars for this AppMesh instance have been discovered. |






<a name="discovery.smh.solo.io.MeshSpec.ConsulConnectMesh"></a>

### MeshSpec.ConsulConnectMesh
Mesh object representing an installed ConsulConnect control plane.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| installation | [MeshSpec.MeshInstallation](#discovery.smh.solo.io.MeshSpec.MeshInstallation) |  |  |






<a name="discovery.smh.solo.io.MeshSpec.Istio"></a>

### MeshSpec.Istio
Mesh object representing an installed Istio control plane


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| installation | [MeshSpec.MeshInstallation](#discovery.smh.solo.io.MeshSpec.MeshInstallation) |  | Configuration metadata about the istio control plane installation. |
| citadelInfo | [MeshSpec.Istio.CitadelInfo](#discovery.smh.solo.io.MeshSpec.Istio.CitadelInfo) |  | Configuration metadata for Istio Citadel (Istio's security component). |
| ingressGateways | [][MeshSpec.Istio.IngressGatewayInfo](#discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo) | repeated | Configuration metadata for Istio IngressGateway (the Istio Ingress). |






<a name="discovery.smh.solo.io.MeshSpec.Istio.CitadelInfo"></a>

### MeshSpec.Istio.CitadelInfo
Configuration metadata for Istio Citadel (Istio's security component).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trustDomain | [string](#string) |  | Istio trust domain used for https/spiffe identity. https://spiffe.io/spiffe/concepts/#trust-domain https://istio.io/docs/reference/glossary/#identity<br>If empty will default to "cluster.local". |
| citadelServiceAccount | [string](#string) |  | istio-citadel service account, used to determine identity for the Istio CA cert. If empty will default to "istio-citadel". |






<a name="discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo"></a>

### MeshSpec.Istio.IngressGatewayInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| workloadLabels | [][MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry](#discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry) | repeated | Labels matching the workload which backs the gateway, defaults to `{"istio": "ingressgateway"}`. |
| externalAddress | [string](#string) |  | The externally-reachable address on which the gateway is listening for TLS connections. This will be the address used for cross-cluster connectivity. Defaults to the LoadBalancer Address (or NodeIP) of the Kubernetes Service (depending on its type). |
| externalTlsPort | [uint32](#uint32) |  | The externally-reachable port on which the gateway is listening for TLS connections. This will be the port used for cross-cluster connectivity. List of common ports: https://istio.io/latest/docs/ops/deployment/requirements/#ports-used-by-istio. Defaults to 15443 (or the NodePort) of the Kubernetes Service (depending on its type). |
| tlsContainerPort | [uint32](#uint32) |  | Container port on which the gateway is listening for TLS connections. Defaults to 15443. |






<a name="discovery.smh.solo.io.MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry"></a>

### MeshSpec.Istio.IngressGatewayInfo.WorkloadLabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="discovery.smh.solo.io.MeshSpec.LinkerdMesh"></a>

### MeshSpec.LinkerdMesh
Mesh object representing an installed Linkerd control plane.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| installation | [MeshSpec.MeshInstallation](#discovery.smh.solo.io.MeshSpec.MeshInstallation) |  |  |
| clusterDomain | [string](#string) |  | The cluster domain suffix this Linkerd mesh is configured with. See https://linkerd.io/2/tasks/using-custom-domain/ for info. |






<a name="discovery.smh.solo.io.MeshSpec.MeshInstallation"></a>

### MeshSpec.MeshInstallation
The cluster on which the control plane for this mesh is deployed. Not all MeshTypes have a MeshInstallation. Only self-hosted control planes such as Istio and Linkerd will have installation metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Namespace in which the control plane has been installed. |
| cluster | [string](#string) |  | Cluster in which the control plane has been installed. |
| podLabels | [][MeshSpec.MeshInstallation.PodLabelsEntry](#discovery.smh.solo.io.MeshSpec.MeshInstallation.PodLabelsEntry) | repeated | the labels on the control plane pods (read from the deployment) |
| version | [string](#string) |  | Version of the Mesh that has been installed. Determined using the image tag on the Mesh's primary control plane image (e.g. the istio-pilot image tag). |






<a name="discovery.smh.solo.io.MeshSpec.MeshInstallation.PodLabelsEntry"></a>

### MeshSpec.MeshInstallation.PodLabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="discovery.smh.solo.io.MeshSpec.OSM"></a>

### MeshSpec.OSM
https://github.com/openservicemesh/osm


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| installation | [MeshSpec.MeshInstallation](#discovery.smh.solo.io.MeshSpec.MeshInstallation) |  | Information about where OSM is installed in a managed Kubernetes Cluster. |






<a name="discovery.smh.solo.io.MeshStatus"></a>

### MeshStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| observedGeneration | [int64](#int64) |  | The observed generation of the Mesh. When this matches the Mesh's metadata.generation, it indicates that mesh-networking has reconciled the latest version of the Mesh. |
| appliedVirtualMesh | [MeshStatus.AppliedVirtualMesh](#discovery.smh.solo.io.MeshStatus.AppliedVirtualMesh) |  | The VirtualMesh, if any, which contains this mesh. |
| appliedFailoverServices | [][MeshStatus.AppliedFailoverService](#discovery.smh.solo.io.MeshStatus.AppliedFailoverService) | repeated | The FailoverServices, if any, which applies to this mesh. |






<a name="discovery.smh.solo.io.MeshStatus.AppliedFailoverService"></a>

### MeshStatus.AppliedFailoverService
AppliedFailoverService represents a FailoverService that has been applied to this Mesh. If an existing FailoverService becomes invalid the last applied FailoverService will be used.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref | [core.skv2.solo.io.ObjectRef](#core.skv2.solo.io.ObjectRef) |  | Reference to the FailoverService. |
| observedGeneration | [int64](#int64) |  | The observed generation of the accepted FailoverService. |
| spec | [networking.smh.solo.io.FailoverServiceSpec](#networking.smh.solo.io.FailoverServiceSpec) |  | The last known valid spec of the FailoverService. |






<a name="discovery.smh.solo.io.MeshStatus.AppliedVirtualMesh"></a>

### MeshStatus.AppliedVirtualMesh
AppliedVirtualMesh represents a VirtualMesh that has been applied to this Mesh. If an existing VirtualMesh becomes invalid, the last applied VirtualMesh will be used.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref | [core.skv2.solo.io.ObjectRef](#core.skv2.solo.io.ObjectRef) |  | Reference to the VirtualMesh. |
| observedGeneration | [int64](#int64) |  | The observed generation of the accepted VirtualMesh. |
| spec | [networking.smh.solo.io.VirtualMeshSpec](#networking.smh.solo.io.VirtualMeshSpec) |  | The last known valid spec of the VirtualMesh. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

