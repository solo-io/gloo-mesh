---
title: "meshctl demo istio-multicluster init"
weight: 5
---
## meshctl demo istio-multicluster init

Bootstrap a multicluster Istio demo with Gloo Mesh

### Synopsis


Bootstrap a multicluster Istio demo with Gloo Mesh.

Running the Gloo Mesh demo setup locally requires 4 tools to be installed and 
accessible via your PATH: kubectl >= v1.18.8, kind >= v0.8.1, istioctl, and docker.
We recommend allocating at least 8GB of RAM for Docker.

This command will bootstrap 2 clusters, one of which will run the Gloo Mesh
management-plane as well as Istio, and the other will just run Istio.


```
meshctl demo istio-multicluster init [flags]
```

### Options

```
      --enterprise       Install the enterprise features, requires a license key
  -h, --help             help for init
      --license string   Gloo Mesh Enterprise license key
      --version string   Gloo Mesh version to install.
                         Community defaults to meshctl version, enterprise defaults to latest
```

### Options inherited from parent commands

```
  -v, --verbose   Enable verbose logging
```

### SEE ALSO

* [meshctl demo istio-multicluster](../meshctl_demo_istio-multicluster)	 - Demo Gloo Mesh functionality with two Istio control planes deployed on separate k8s clusters

