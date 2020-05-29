---
title: Multi-cluster Traffic
menuTitle: Multi-cluster Traffic
weight: 75
---

In the [previous guides]({{% versioned_link_path fromRoot="/guides/federate_identity/" %}}), we federated multiple meshes and established a [shared root CA for a shared identity]({{% versioned_link_path fromRoot="/guides/federate_identity/#understanding-the-shared-root-process" %}}) domain. Now that we have a logical [VirtualMesh]({{% versioned_link_path fromRoot="/reference/api/virtual_mesh/" %}}), we need a way to establish **traffic** policies across the multiple meshes, without treating each of them individually. Service Mesh Hub helps by establishing a single, unified API that understands the logical VirtualMesh construct.

## Before you begin
To illustrate these concepts, we will assume that:

* Service Mesh Hub is [installed and running on the `management-plane-context`]({{% versioned_link_path fromRoot="/setup/#install-service-mesh-hub" %}})
* Istio is [installed on both `management-plane-context` and `remote-cluster-context`]({{% versioned_link_path fromRoot="/guides/installing_istio" %}}) clusters
* Both `management-plane-context` and `remote-cluster-context` clusters are [registered with Service Mesh Hub]({{% versioned_link_path fromRoot="/guides/#two-registered-clusters" %}})
* The `bookinfo` app is [installed into two Istio clusters]({{% versioned_link_path fromRoot="/guides/#bookinfo-deployed-on-two-clusters" %}})


{{% notice note %}}
Be sure to review the assumptions and satisfy the pre-requisites from the [Guides]({{% versioned_link_path fromRoot="/guides" %}}) top-level document.
{{% /notice %}}

## Controlling cross-cluster traffic

We will now perform a *multi-cluster traffic split*, splitting traffic from the `productpage` service in the `management-plane-context` cluster between `reviews-v1`, `reviews-v2`, and `reviews-v3` running in the `remote-cluster-context` cluster.

{{< tabs >}}
{{< tab name="YAML file" codelang="shell">}}
apiVersion: networking.zephyr.solo.io/v1alpha1
kind: TrafficPolicy
metadata:
  namespace: service-mesh-hub
  name: simple
spec:
  destinationSelector:
    serviceRefs:
      services:
        - cluster: management-plane
          name: reviews
          namespace: default
  trafficShift:
    destinations:
      - destination:
          cluster: new-remote-cluster
          name: reviews
          namespace: default
        weight: 75
      - destination:
          cluster: management-plane
          name: reviews
          namespace: default
        weight: 15
        subset:
          version: v1
      - destination:
          cluster: management-plane
          name: reviews
          namespace: default
        weight: 10
        subset:
          version: v2
{{< /tab >}}
{{< tab name="CLI inline" codelang="shell" >}}
kubectl apply --context management-plane-context -f - << EOF
apiVersion: networking.zephyr.solo.io/v1alpha1
kind: TrafficPolicy
metadata:
  namespace: service-mesh-hub
  name: simple
spec:
  destinationSelector:
    serviceRefs:
      services:
        - cluster: management-plane
          name: reviews
          namespace: default
  trafficShift:
    destinations:
      - destination:
          cluster: new-remote-cluster
          name: reviews
          namespace: default
        weight: 75
      - destination:
          cluster: management-plane
          name: reviews
          namespace: default
        weight: 15
        subset:
          version: v1
      - destination:
          cluster: management-plane
          name: reviews
          namespace: default
        weight: 10
        subset:
          version: v2
EOF
{{< /tab >}}
{{< /tabs >}}

{{% notice warning %}}
You may need to restart your workloads to ensure that they pick up the newly-distributed certs. Istio currently does not have support
for workloads detecting a change in root cert and re-issuing an SDS request.
{{% /notice %}}

Once you apply this resource to the `management-plane-context` cluster, you should occasionally see traffic being routed to the reviews-v3 service, which will produce red-colored stars on the product page.

To go into slightly more detail here: The above TrafficPolicy says that:

* Any traffic destined for the *reviews service* in the *management plane cluster*
* Should be split across several different destinations
* 25% of traffic gets split between the v1 and v2 instances of the reviews service in the management plane cluster
* 75% of traffic gets sent to the v3 instance of the reviews service on the remote cluster

We have successfully set up multi-cluster traffic across our application! Note that this has been done transparently to the application. The application can continue talking to what it believes is the local instance of the service, but, behind the scenes, we have instead routed that traffic to an entirely different cluster. 

Note that this is interesting in its own right, that we have easily
achieved multi-cluster communication, but also in other scenarios like in fast disaster recovery: We can quickly route traffic to healthy instances of a service in an entirely different data center.


## See it in action

Check out "Part Four" of the ["Dive into Service Mesh Hub" video series](https://www.youtube.com/watch?v=4sWikVELr5M&list=PLBOtlFtGznBjr4E9xYHH9eVyiOwnk1ciK):

<iframe width="560" height="315" src="https://www.youtube.com/embed/HAr1Mw1bxB4" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>


