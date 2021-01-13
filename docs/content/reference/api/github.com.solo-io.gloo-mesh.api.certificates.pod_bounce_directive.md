
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
  - [PodBounceDirectiveStatus](#certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus)
  - [PodBounceDirectiveStatus.BouncedPodSet](#certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet)







<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec"></a>

### PodBounceDirectiveSpec
When certificates are issued, pods may need to be bounced (restarted) to ensure they pick up the new certificates. If so, the certificate Issuer will create a PodBounceDirective containing the namespaces and labels of the pods that need to be bounced in order to pick up the new certs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| podsToBounce | [][certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector" >}}) | repeated | A list of k8s pods to bounce (delete and cause a restart) when the certificate is issued. This will include the control plane pods as well as any pods which share a data plane with the target mesh. |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector"></a>

### PodBounceDirectiveSpec.PodSelector
Pods that will be restarted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | string |  | The namespace in which the pods live. |
  | labels | [][certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry" >}}) | repeated | Any labels shared by the pods. |
  | waitForReplicas | uint32 |  | Wait for this number of replacement pods to reach be fully Ready before deleting the next set of selected pods. This is used to ensure the control plane pods are allowed to restart before sidecars and gateways are restarted. |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry"></a>

### PodBounceDirectiveSpec.PodSelector.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | string |  |  |
  | value | string |  |  |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus"></a>

### PodBounceDirectiveStatus
PodBounceDirectiveStatus reports the status for stateful pod bounces (when bouncing pods requires waiting for readiness)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| podsBounced | [][certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.certificates.pod_bounce_directive#certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet" >}}) | repeated | A list of k8s pods to bounce (delete and cause a restart) when the certificate is issued. This will include the control plane pods as well as any pods which share a data plane with the target mesh. |
  





<a name="certificates.mesh.gloo.solo.io.PodBounceDirectiveStatus.BouncedPodSet"></a>

### PodBounceDirectiveStatus.BouncedPodSet
A set of Pods that were restarted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bouncedPods | []string | repeated | The names of the pods that were bounced for the corresponding selector. |
  




 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->

