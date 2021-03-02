---
title: "meshctl uninstall"
weight: 5
---
## meshctl uninstall

Uninstall Gloo Mesh from the referenced cluster

```
meshctl uninstall [flags]
```

### Options

```
  -d, --dry-run               Output installation manifest
  -h, --help                  help for uninstall
      --kubeconfig string     path to the kubeconfig from which the management cluster will be accessed
      --kubecontext string    name of the kubeconfig context to use for the management cluster
      --namespace string      namespace in which to uninstall Gloo Mesh from (default "gloo-mesh")
      --release-name string   Helm release name (default "gloo-mesh")
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose logging
```

### SEE ALSO

* [meshctl](../meshctl)	 - The Command Line Interface for managing Gloo Mesh.

