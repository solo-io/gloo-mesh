---
title: "meshctl debug snapshot discovery input"
weight: 5
---
## meshctl debug snapshot discovery input

Input snapshot for the discovery pod

```
meshctl debug snapshot discovery input [flags]
```

### Options

```
  -h, --help   help for input
```

### Options inherited from parent commands

```
      --dir string           dir to write file outputs to
  -f, --file string          file to write output to
      --json                 display the entire json snapshot. The output can be piped into a command like jq (https://stedolan.github.io/jq/tutorial/). For example:
                              meshctl debug snapshot discovery input | jq '.'
      --kubeconfig string    Path to the kubeconfig from which the management cluster will be accessed
      --kubecontext string   Name of the kubeconfig context to use for the management cluster
  -v, --verbose              Enable verbose logging
      --zip string           zip file output
```

### SEE ALSO

* [meshctl debug snapshot discovery](../meshctl_debug_snapshot_discovery)	 - Input and output snapshots for the discovery pod

