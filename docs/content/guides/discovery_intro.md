---
title: Mesh Discovery
menuTitle: Mesh Discovery
weight: 20
---

Gloo Mesh can automatically discover service mesh installations on registered clusters using control plane and sidecar discovery, as well as workloads and services exposed through the service mesh.

In this guide we will learn about the four main discovery capabilities in the context of Kubernetes as the compute platform:

1. **Kubernetes Clusters**
    - Representation of a cluster that Gloo Mesh is aware of and is authorized to talk to its Kubernetes API server
    - *note*: this resource is created by `meshctl` at cluster registration time
2. **Meshes**
    - Representation of a service mesh control plane that has been discovered 
3. **Workloads**
    - Representation of a pod that is a member of a service mesh; this is often determined by the presence of an injected proxy sidecar
4. **Destinations**
    - Representation of a Kubernetes service that is backed by Workload pods, e.g. pods that are a member of the service mesh


## Before you begin
To illustrate these concepts, we will assume that:

* Gloo Mesh is [installed and running on the `mgmt-cluster`]({{% versioned_link_path fromRoot="/setup/#install-gloo-mesh" %}})
* Istio is [installed on both the `mgmt-cluster` and `remote-cluster`]({{% versioned_link_path fromRoot="/guides/installing_istio" %}})
* Both `mgmt-cluster` and `remote-cluster` clusters are [registered with Gloo Mesh]({{% versioned_link_path fromRoot="/guides/#two-registered-clusters" %}})
* The `bookinfo` app is [installed into the two clusters]({{% versioned_link_path fromRoot="/guides/#bookinfo-deployment" %}})


{{% notice note %}}
Be sure to review the assumptions and satisfy the pre-requisites from the [Guides]({{% versioned_link_path fromRoot="/guides" %}}) top-level document.
{{% /notice %}}

### Discover Kubernetes Clusters

Ensure that your `kubeconfig` has the correct context set as its `currentContext`:

```shell
MGMT_CONTEXT=your_management_plane_context
REMOTE_CONTEXT=your_remote_context
kubectl config use-context $MGMT_CONTEXT
```

Validate that the cluster have been registered by checking for `KubernetesClusters` custom resources:

```shell
kubectl get kubernetesclusters -n gloo-mesh
```

```shell
NAME                 AGE
mgmt-cluster         23h
remote-cluster       23h
```

### Discover Meshes

Check to see that Istio has been discovered:

```shell
meshctl describe mesh
```

```
+-----------------------------+----------------+-------------------+
|           METADATA          | VIRTUAL MESHES | FAILOVER SERVICES |
+-----------------------------+----------------+-------------------+
| Namespace: istio-system     |                |                   |
| Cluster: mgmt-cluster       |                |                   |
| Type: istio                 |                |                   |
| Version: 1.8.1              |                |                   |
|                             |                |                   |
+-----------------------------+----------------+-------------------+
| Namespace: istio-system     |                |                   |
| Cluster: remote-cluster     |                |                   |
| Type: istio                 |                |                   |
| Version: 1.8.1              |                |                   |
|                             |                |                   |
+-----------------------------+----------------+-------------------+
```

We can print it in YAML form to see all the information we discovered:

```shell
kubectl -n gloo-mesh get mesh istiod-istio-system-remote-cluster -oyaml
```

(snipped for brevity)

{{< highlight yaml >}}
apiVersion: discovery.mesh.gloo.solo.io/v1
kind: Mesh
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
    [...]
  generation: 2
  labels:
    cluster.discovery.mesh.gloo.solo.io: remote-cluster
    cluster.multicluster.solo.io: ""
    owner.discovery.mesh.gloo.solo.io: gloo-mesh
  name: istiod-istio-system-remote-cluster
  namespace: gloo-mesh
  resourceVersion: "3218"
  selfLink: /apis/discovery.mesh.gloo.solo.io/v1/namespaces/gloo-mesh/meshes/istiod-istio-system-remote-cluster
  uid: 7c079983-3ece-4aed-b71a-bf56c8cd6267
spec:
  agentInfo:
    agentNamespace: gloo-mesh
  istio:
    citadelInfo:
      IstiodServiceAccount: istiod-service-account
      trustDomain: cluster.local
    ingressGateways:
    - externalAddress: 172.20.0.3
      externalTlsPort: 32000
      tlsContainerPort: 15443
      workloadLabels:
        istio: ingressgateway
    installation:
      cluster: remote-cluster
      namespace: istio-system
      podLabels:
        istio: pilot
      version: 1.8.1
status:
  observedGeneration: 2

{{< /highlight >}}

### Discover Workloads

Check to see that the `bookinfo` pods have been correctly identified as Workloads:

```shell
kubectl -n gloo-mesh get workloads
```

```
NAME                                                            AGE
details-v1-bookinfo-mgmt-cluster-deployment                     3m54s
istio-ingressgateway-istio-system-mgmt-cluster-deployment       23h
istio-ingressgateway-istio-system-remote-cluster-deployment     23h
productpage-v1-bookinfo-mgmt-cluster-deployment                 3m54s
ratings-v1-bookinfo-mgmt-cluster-deployment                     3m53s
ratings-v1-bookinfo-remote-cluster-deployment                   3m25s
reviews-v1-bookinfo-mgmt-cluster-deployment                     3m53s
reviews-v2-bookinfo-mgmt-cluster-deployment                     3m53s
reviews-v3-bookinfo-remote-cluster-deployment                   2m
```

### Discover Destinations

Similarly for the `bookinfo` services:

```shell
kubectl -n gloo-mesh get destinations
```

```
NAME                                                 AGE
details-bookinfo-mgmt-cluster                        4m23s
istio-ingressgateway-istio-system-mgmt-cluster       23h
istio-ingressgateway-istio-system-remote-cluster     23h
productpage-bookinfo-mgmt-cluster                    4m23s
ratings-bookinfo-mgmt-cluster                        4m22s
ratings-bookinfo-remote-cluster                      3m54s
reviews-bookinfo-mgmt-cluster                        4m22s
reviews-bookinfo-remote-cluster                      2m29s
```

## See it in action

Check out "Part One" of the ["Dive into Gloo Mesh" video series](https://www.youtube.com/watch?v=4sWikVELr5M&list=PLBOtlFtGznBjr4E9xYHH9eVyiOwnk1ciK)
(note that the video content reflects Gloo Mesh <b>v0.6.1</b>):

<iframe width="560" height="315" src="https://www.youtube.com/embed/4sWikVELr5M" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

## Next Steps

Now that we have Istio installed, and we've seen Gloo Mesh discover the meshes across different clusters, we can now unify them into a single [Virtual Mesh]({{% versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1alpha2.virtual_mesh/" %}}). See the guide on [establishing shared trust domain for multiple meshes]({{% versioned_link_path fromRoot="/guides/federate_identity" %}}).
