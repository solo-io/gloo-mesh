---
title: "Install Gloo Mesh"
menuTitle: Install Gloo Mesh
description: Installing Gloo Mesh on a management cluster.
weight: 20
---

Gloo Mesh and Gloo Mesh Enterprise use a Kubernetes cluster to host the management plane (Gloo Mesh) while each service mesh can run on its own independent cluster. If you don't have access to multiple clusters, see the [Getting Started Guide]({{% versioned_link_path fromRoot="/getting_started/" %}}) to get started with Kubernetes in Docker, or refer to our [Using Kind]({{% versioned_link_path fromRoot="/setup/kind_setup" %}}) setup guide to provision two clusters.

{{% notice note %}}
Gloo Mesh Enterprise is the paid version of Gloo Mesh including the Gloo Mesh UI and multi-cluster role-based access control. To complete the installation you will need a license key. You can get a trial license by [requesting a demo](https://www.solo.io/products/gloo-mesh/) from the website.
{{% /notice %}}

![Gloo Mesh Architecture]({{% versioned_link_path fromRoot="/img/gloomesh-3clusters.png" %}})

You can install Gloo Mesh onto its own cluster and register remote clusters, or you can co-locate Gloo Mesh onto a cluster with a service mesh. The former (its own cluster) is the preferred deployment pattern, but for getting started, exploring, or to save resources, you can use the co-located deployment approach.

![Gloo Mesh Architecture]({{% versioned_link_path fromRoot="/img/gloomesh-2clusters.png" %}})

In this guide we will walk through the process of installing Gloo Mesh either through [meshctl](#installing-with-meshctl) or by using [Helm](#install-with-helm).

## Assumptions for setup

We will assume in this and following guides that we have access to two clusters and the following two contexts available in our `kubeconfig` file. 

Your actual context names will likely be different.

* `mgmt-cluster-context`
    - kubeconfig context pointing to a cluster where we will install and operate Gloo Mesh
* `remote-cluster-context`
    - kubeconfig context pointing to a cluster where we will install and manage a service mesh using Gloo Mesh 

To verify you're running the following commands in the correct context, run:

```shell
MGMT_CONTEXT=your_management_plane_context
REMOTE_CONTEXT=your_remote_context

kubectl config use-context $MGMT_CONTEXT
```

## Install Gloo Mesh

{{% notice note %}}
Note that these contexts need not be different; you may install and manage a service mesh in the same cluster as Gloo Mesh. For the purposes of this guide, though, we will assume they are different.
{{% /notice %}}

The following directions show how to install both Gloo Mesh and Gloo Mesh Enterprise. Select the tab that meets your scenario. If you are deploying Gloo Mesh Enterprise, you should also follow our guide on configuring Role-based API control.

### Installing with `meshctl`

`meshctl` is a CLI tool that helps bootstrap Gloo Mesh, register clusters, describe configured resources, and more. Get the latest `meshctl` from the [releases page on solo-io/gloo-mesh](https://github.com/solo-io/gloo-mesh/releases).

You can also quickly install like this:

```shell
curl -sL https://run.solo.io/meshctl/install | sh
```

Once you have the `meshctl` tool, you can install Gloo Mesh onto a cluster acting as the `mgmt-cluster` like this:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
meshctl install
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
meshctl install enterprise --license LICENSE_KEY_STRING
{{< /tab >}}
{{< /tabs >}}

If you're not using the context for the `mgmt-cluster`, you can explicitly specify it using the `--kubecontext` option:

```shell
meshctl install --kubecontext $MGMT_CONTEXT
```

You should see output similar to this:

```shell
Creating namespace gloo-mesh... Done.
Starting Gloo Mesh installation...
Gloo Mesh successfully installed!
Gloo Mesh has been installed to namespace gloo-mesh
```

To undo the installation, run `uninstall`:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
meshctl uninstall
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
meshctl uninstall enterprise
{{< /tab >}}
{{< /tabs >}}

### Installing with `kubectl apply`

If you prefer working directly with the Kubernetes resources, (either to use `kubectl apply` or to put into CI/CD), `meshctl` can output `yaml` from the `install` (or any) command with the `--dry-run` flag:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
meshctl install --dry-run
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
meshctl install enterprise --license LICENSE_KEY_STRING --dry-run
{{< /tab >}}
{{< /tabs >}}

You can use this output to later run `kubectl apply`:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
meshctl install --dry-run | kubectl apply -f -
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
meshctl install enterprise --license LICENSE_KEY_STRING --dry-run | kubectl apply -f -
{{< /tab >}}
{{< /tabs >}}

{{% notice note %}}
Note that the `--dry-run` outputs the entire `yaml`, but does not take care of proper ordering of resources. For example, there can be a *race* between Custom Resource Definitions being registered and any Custom Resources being created that may appear to be an error. If that happens, just re-run the `kubectl apply`.
{{% /notice %}}

To undo the installation, run:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
meshctl install --dry-run | kubectl delete -f -
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
meshctl install enterprise --license LICENSE_KEY_STRING --dry-run | kubectl delete -f -
{{< /tab >}}
{{< /tabs >}}

### Install with Helm

The Helm charts for Gloo Mesh support Helm 3. To install with Helm first add the Gloo Mesh or Gloo Mesh Enterprise Helm repository:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
helm repo add gloo-mesh https://storage.googleapis.com/gloo-mesh/gloo-mesh
helm repo update
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
helm repo add gloo-mesh-enterprise https://storage.googleapis.com/gloo-mesh-enterprise/gloo-mesh-enterprise
helm repo update
{{< /tab >}}
{{< /tabs >}}

{{% notice note %}}
Note that the location of the Gloo Mesh Helm charts is subject to change. When it finds a more permanent home, we'll remove this message.
{{% /notice %}}

Then install Gloo Mesh into the `gloo-mesh` namespace:


{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
kubectl create ns gloo-mesh
helm install gloo-mesh gloo-mesh/gloo-mesh --namespace gloo-mesh
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
kubectl create ns gloo-mesh
helm install gloo-mesh-enterprise gloo-mesh-enterprise/gloo-mesh-enterprise -n gloo-mesh --set licenseKey=LICENSE_KEY_STRING
{{< /tab >}}
{{< /tabs >}}


### Verify install
Once you've installed Gloo Mesh, verify what components were installed:

{{< tabs >}}
{{< tab name="Gloo Mesh" codelang="shell" >}}
kubectl get pods -n gloo-mesh

NAME                          READY   STATUS    RESTARTS   AGE
discovery-66675cf6fd-cdlpq    1/1     Running   0          32m
networking-6d7686564d-ngrdq   1/1     Running   0          32m
{{< /tab >}}
{{< tab name="Gloo Mesh Enterprise" codelang="shell">}}
kubectl get pods -n gloo-mesh

NAME                                   READY   STATUS    RESTARTS   AGE
discovery-67444fcf85-rdh7j             1/1     Running   0          4m5s
gloo-mesh-apiserver-6469d4b5f6-7b85g   3/3     Running   0          4m5s
networking-5dcb4d44d5-jzxk6            1/1     Running   0          4m5s
rbac-webhook-6fd7898df8-wkvpj          1/1     Running   0          4m5s
{{< /tab >}}
{{< /tabs >}}

Running the check command from meshctl will verify everything was installed correctly:

```shell
meshctl check
```

```shell
Gloo Mesh
-------------------
✅ Gloo Mesh pods are running

Management Configuration
---------------------------
✅ Gloo Mesh networking configuration resources are in a valid state
```

## Next steps

With Gloo Mesh or Gloo Mesh Enterprise installed, the next step is to [register clusters with Gloo Mesh]({{% versioned_link_path fromRoot="/setup/register_cluster" %}}) and discover the service meshes running on those clusters.



