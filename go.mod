module github.com/solo-io/smh

go 1.14

replace (
	// github.com/Azure/go-autorest/autorest has different versions for the Go
	// modules than it does for releases on the repository. Note the correct
	// version when updating.
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	github.com/solo-io/service-mesh-hub => ../service-mesh-hub
	github.com/solo-io/skv2 => ../skv2

	k8s.io/api => k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.2
	k8s.io/apiserver => k8s.io/apiserver v0.17.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.2
	k8s.io/client-go => k8s.io/client-go v0.17.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.2
	k8s.io/code-generator => k8s.io/code-generator v0.17.2
	k8s.io/component-base => k8s.io/component-base v0.17.2
	k8s.io/cri-api => k8s.io/cri-api v0.17.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.2
	k8s.io/klog => github.com/stefanprodan/klog v0.0.0-20190418165334-9cbb78b20423
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.2
	k8s.io/kubectl => k8s.io/kubectl v0.17.2
	k8s.io/kubelet => k8s.io/kubelet v0.17.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.2
	k8s.io/metrics => k8s.io/metrics v0.17.2
	k8s.io/node-api => k8s.io/node-api v0.17.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.2
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.17.2
	k8s.io/sample-controller => k8s.io/sample-controller v0.17.2
)

require (
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/envoyproxy/go-control-plane v0.9.4
	github.com/go-delve/delve v1.4.1 // indirect
	github.com/go-logr/zapr v0.1.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4-0.20200406172829-6d816de489c1
	github.com/golang/protobuf v1.3.5
	github.com/hashicorp/consul v1.6.2
	github.com/hashicorp/go-multierror v1.0.0
	github.com/hashicorp/hcl v1.0.0
	github.com/linkerd/linkerd2 v0.5.1-0.20200402173539-fee70c064bc0
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.8.1
	github.com/peterh/liner v1.2.0 // indirect
	github.com/rotisserie/eris v0.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/solo-io/external-apis v0.0.5
	github.com/solo-io/go-utils v0.16.0
	github.com/solo-io/service-mesh-hub v0.5.1-0.20200630141134-62e606c0af98
	github.com/solo-io/skv2 v0.6.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	go.starlark.net v0.0.0-20200619143648-50ca820fafb9 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/arch v0.0.0-20200511175325-f7c78586839d // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	istio.io/api v0.0.0-20200610220835-a1a958746907
	istio.io/client-go v0.0.0-20200610222318-1cfead1f1938
	istio.io/istio v0.0.0-20200215010343-d9274c558175
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kubernetes v1.18.4
	sigs.k8s.io/controller-runtime v0.5.6
)
