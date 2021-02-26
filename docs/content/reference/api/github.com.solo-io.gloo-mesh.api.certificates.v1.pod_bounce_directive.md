
---

title: "pod_bounce_directive.proto"

---

## Package : `certificates.mesh.gloo.solo.io`



<a name="top"></a>

<a name="API Reference for pod_bounce_directive.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pod_bounce_directive.proto


## Table of Contents
  - [PodBounceDirectiveSpec](#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec)
  - [PodBounceDirectiveSpec.PodSelector](#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector)
  - [PodBounceDirectiveSpec.PodSelector.LabelsEntry](#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry)
  - [PodBounceDirectiveSpec.PodSelector.RootCertSync](#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync)
  - [PodBounceDirectiveStatus](#certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus)
  - [PodBounceDirectiveStatus.BouncedPodSet](#certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet)







<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec"></a>

### PodBounceDirectiveSpec
When certificates are issued, Istio-controlled pods need to be bounced (restarted) to ensure they pick up the new certificates due to [this issue](https://github.com/istio/istio/issues/22993). The certificate issuer will create a PodBounceDirective containing the namespaces and labels of the pods that need to be bounced in order to pick up the new certs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| podsToBounce | [][certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.v1.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector" >}}) | repeated | A list of Kubernetes pods to bounce (delete and cause a restart) when the certificate is issued. This will include the control plane pods as well as any Pods which share a data plane with the target mesh. |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector"></a>

### PodBounceDirectiveSpec.PodSelector
pods that will be restarted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | string |  | The namespace in which the pods live. |
  | labels | [][certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.v1.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry" >}}) | repeated | Any labels shared by the Pods. |
  | waitForReplicas | uint32 |  | Wait for this number of replacement pods to reach be fully ready before deleting the next set of selected Pods. This is used to ensure the control plane pods are allowed to restart before sidecars and gateways are restarted. |
  | rootCertSync | [certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.v1.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync" >}}) |  | Wait for the control plane to have synced all root cert configmaps in data plane namespaces before bouncing these Pods. |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry"></a>

### PodBounceDirectiveSpec.PodSelector.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | string |  |  |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.RootCertSync"></a>

### PodBounceDirectiveSpec.PodSelector.RootCertSync
RootCertSync describes values in a secret and configmap which must be equal in order for a Pod to be bounced.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secretRef | [core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) |  |  |
  | secretKey | string |  |  |
  | configMapRef | [core.skv2.solo.io.ObjectRef]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.skv2.api.core.v1.core#core.skv2.solo.io.ObjectRef" >}}) |  |  |
  | configMapKey | string |  |  |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus"></a>

### PodBounceDirectiveStatus
PodBounceDirectiveStatus reports the status for stateful Pod bounces (when bouncing pods requires waiting for readiness).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| podsBounced | [][certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.v1.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet" >}}) | repeated | A list of Kubernetes pods to bounce (delete and cause a restart) when the certificate is issued. This will include the control plane pods as well as any Pods which share a data plane with the target mesh. |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet"></a>

### PodBounceDirectiveStatus.BouncedPodSet
A set of pods that were restarted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bouncedPods | []string | repeated | The names of the pods that were bounced for the corresponding selector specified in `PodBounceDirectiveSpec.PodSelector.labels`. |
  




 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

