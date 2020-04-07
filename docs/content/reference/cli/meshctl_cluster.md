---
title: "meshctl cluster"
weight: 5
---
## meshctl cluster

Register and perform operations on clusters

### Synopsis

Register and perform operations on clusters

```
meshctl cluster [flags]
```

### Options

```
  -h, --help   help for cluster
```

### Options inherited from parent commands

```
      --context string          Specify which context from the kubeconfig should be used; uses current context if none is specified
      --kube-timeout duration   Specify the default timeout for requests to kubernetes API servers (default 5s)
      --kubeconfig string       Specify the kubeconfig for the current command
  -n, --namespace string        Specify the namespace where Service Mesh Hub resources should be written (default "service-mesh-hub")
  -v, --verbose                 Enable verbose mode, which outputs additional execution details that may be helpful for debugging
```

### SEE ALSO

* [meshctl](../meshctl)	 - CLI for Service Mesh Hub
* [meshctl cluster register](../meshctl_cluster_register)	 - Register a new cluster by creating a service account token in that cluster through which to authorize Service Mesh Hub

