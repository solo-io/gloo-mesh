---
title: FailoverService
menuTitle: FailoverService
weight: 25
---

Service Mesh Hub provides the ability to configure a FailoverService. A FailoverService is
 a virtual traffic destination that is composed of a list of services ordered in decreasing
priority. The composing services are configured with outlier detection (the ability of the system 
to detect unresponsive services, [read more here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/outlier)), and traffic will automatically be shifted over to 
services next in the priority order. Currently this feature is only supported for Istio meshes.

## Before you begin
To illustrate these concepts, we will assume that:

* Service Mesh Hub is [installed and running on the `management-cluster`]({{% versioned_link_path fromRoot="/setup/#install-service-mesh-hub" %}})
* Istio is [installed on both `management-cluster` and `remote-cluster`]({{% versioned_link_path fromRoot="/guides/installing_istio" %}}) clusters
* Both `management-cluster` and `remote-cluster` clusters are [registered with Service Mesh Hub]({{% versioned_link_path fromRoot="/guides/#two-registered-clusters" %}})
under the names `management-cluster` and `remote-cluster` respectively
* The `bookinfo` app is [installed into two Istio clusters]({{% versioned_link_path fromRoot="/guides/#bookinfo-deployed-on-two-clusters" %}})


{{% notice note %}}
Be sure to review the assumptions and satisfy the pre-requisites from the [Guides]({{% versioned_link_path fromRoot="/guides" %}}) top-level document.
{{% /notice %}}

## Configure Outlier Detection

The services composing a FailoverService must be configured with outlier detection, 
which is done through the [TrafficPolicy]({{< versioned_link_path fromRoot="/reference/api/traffic_policy" >}}).
Apply the following config on the `management-cluster` cluster:

{{< highlight yaml "hl_lines=16-19" >}}
apiVersion: networking.smh.solo.io/v1alpha2
kind: TrafficPolicy
metadata:
  namespace: service-mesh-hub
  name: mgmt-reviews-outlier
spec:
  destinationSelector:
  - kubeServiceRefs:
      services:
      - name: reviews
        namespace: default
        clusterName: management-cluster
      - name: reviews
        namespace: default
        clusterName: new-remote-cluster
  outlierDetection:
    consecutiveErrors: 1
{{< /highlight >}}

For demonstration purposes, we're setting `consectiveErrors` to 1 for more easily
triggering the failover. Once applied, run the following:

```shell
kubectl -n service-mesh-hub get trafficpolicy/mgmt-reviews-outlier -oyaml
```

and you should see the following status indicating that the TrafficPolicy is valid and has been translated
into mesh-specific config:

```yaml
status:
  trafficTargets:
    reviews-default-management-cluster.service-mesh-hub.:
      state: ACCEPTED
    reviews-default-new-remote-cluster.service-mesh-hub.:
      state: ACCEPTED
  observedGeneration: "1"
  state: ACCEPTED
```

## Create the FailoverService

Now we will create the FailoverService for the `reviews` service, composed of the 
reviews service on the `management-cluster` cluster in first priority and on `new-remote-cluster`
in second priority. If the `reviews` service on the `management-cluster` cluster is unhealthy,
 requests will automatically be shifted over to the service on `new-remote-cluster`.
 Apply the following config to the `management-cluster` cluster:
 
```yaml
apiVersion: networking.smh.solo.io/v1alpha2
kind: FailoverService
metadata:
  name: reviews-failover
  namespace: service-mesh-hub
spec:
  hostname: reviews-failover.bookinfo.global
  port:
    number: 9080
    protocol: http
  meshes:
    - name: istiod-istio-system-management-cluster-cluster
      namespace: service-mesh-hub
  backingServices:
  - kubeService:
      name: reviews
      namespace: bookinfo
      clusterName: management-cluster
  - kubeService:
      name: reviews
      namespace: bookinfo
      clusterName: new-remote-cluster
```

{{% notice note %}}
The `.global` suffix is needed if you want the failoverService's hostname to be resolvable in the k8s context (e.g. with commands like `curl`), assuming
your Istio installation has istiocoredns enabled. This leverages Istio's default DNS setup, which creates DNS configuration for hostnames with `.global` suffix.
{{% /notice %}}

Notice that the services referenced under `failoverServices` matches the services
for which we just configured outlier detection. This is a requirement for all services composing
for a FailoverService.

The `meshes` field indicates the control planes that the FailoverService is visible to.
If multiple meshes are listed, they must be grouped under a common VirtualMesh.

Once applied, run the following

```shell
kubectl -n service-mesh-hub get failoverservice/reviews-failover -oyaml
```

and you should see the following status:

```yaml
status:
    meshes:
      istiod-istio-system-management-cluster.service-mesh-hub.:
        state: ACCEPTED
    observedGeneration: "1"
    state: ACCEPTED
```

## Demonstrating Failover Functionality

To demonstrate failover functionality, we configure a traffic shift such that all requests
targeting the `reviews` service will instead be routed to the reviews FailoverService
we created above.

{{% notice note %}}

As of Service Mesh Hub version v0.6.0, TrafficPolicies are unable to reference FailoverServices in their destination selector. 

We're currently working on extending the TrafficPolicy to enable this functionality.
 
For purposes of illustration, we'll manually create an Istio VirtualService to achieve the traffic shift.
{{% /notice %}}

Apply the following Istio VirtualService:

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: reviews-failover
  namespace: bookinfo
spec:
  hosts:
  - reviews
  http:
  - route:
    - destination:
        host: reviews-failover.bookinfo.global
        port:
          number: 9080
```

Port forward the productpage pod with the following command and open your web browser to
[localhost:9080](http://localhost:9080/productpage?u=normal).

```shell
kubectl -n bookinfo port-forward deployments/productpage-v1 9080
```

Reloading the page a few times should show the "Book Reviews" section showing either
no stars (for requests routed to the `reviews-v1` pod) and black stars 
(for requests routed to the `reviews-v2` pod). This shows that the productpage is routing
to the first service listed in the reviews-failover FailoverService. Recall from the
[multicluster setup guide]({{% versioned_link_path fromRoot="/guides/#two-registered-clusters" %}})
that `reviews-v1` and `reviews-v1` only exist on `management-cluster` and `reviews-v3` only
exists on `new-remote-cluster`, which we'll use to distinguish requests routing to either cluster.

Now, to trigger the failover, we'll modify the `reviews-v1` and `reviews-v2` deployment
to disable the web server. Run the following commands:

```shell
kubectl -n bookinfo patch deploy reviews-v1 --patch '{"spec": {"template": {"spec": {"containers": [{"name": "reviews","command": ["sleep", "20h"]}]}}}}'
```

```shell
kubectl -n bookinfo patch deploy reviews-v2 --patch '{"spec": {"template": {"spec": {"containers": [{"name": "reviews","command": ["sleep", "20h"]}]}}}}'
```

Once the modified deployment has rolled out, refresh the productpage and you should see
reviews with red stars, corresponding to `reviews-v3`, which only exists on `new-remote-cluster`,
demonstrating that the requests are indeed failing over to the second service in the list.

To restore the disabled `reviews-v1` and `reviews-v2`, run the following:

```shell
kubectl -n bookinfo patch deployment reviews-v1  --type json   -p '[{"op": "remove", "path": "/spec/template/spec/containers/0/command"}]'
```

```shell
kubectl -n bookinfo patch deployment reviews-v2  --type json   -p '[{"op": "remove", "path": "/spec/template/spec/containers/0/command"}]'
```

Once the deployment has rolled out, reloading the productpage should show reviews with no stars or black stars, indicating that
the failover service is routing requests to the first listed service in the `management-cluster`.
