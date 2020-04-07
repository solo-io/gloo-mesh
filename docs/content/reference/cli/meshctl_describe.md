---
title: "meshctl describe"
weight: 5
---
## meshctl describe

Describe policies affecting your Kubernetes services (kube-native services) or workloads (e.g., kube-native deployments). Format the `resource_name` arg as kube-name.kube-namespace.registered-cluster-name

### Synopsis

Describe policies affecting your Kubernetes services (kube-native services) or workloads (e.g., kube-native deployments). Format the `resource_name` arg as kube-name.kube-namespace.registered-cluster-name

```
meshctl describe (service|workload) resource_name [flags]
```

### Options

```
  -h, --help              help for describe
      --policies string   Policies to view. Valid policies: [all, access, traffic] (default "all")
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

