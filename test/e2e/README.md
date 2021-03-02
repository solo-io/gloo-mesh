# Integration tests
This directory contains specs that test the e2e workflows of Gloo Mesh deployed on Kubernetes.

## Cluster setup

The e2e tests currently run against a KIND cluster which is created via `ci/setup-kind.sh`. The entrypoint is contained in `e2e_suite_test.go`.

To run tests:

```shell script
RUN_E2E=true TEST_PKG=test/e2e make run-tests 
```

The e2e suite will automatically run `ci/setup-kind.sh` and teardown the cluster on test teardown.

To reuse an existing cluster (note that the tests attempt to reach the ingress on `localhost:32000`) and skip the 
setup/teardown steps, set env `USE_EXISTING=<name of kuybe context>`
