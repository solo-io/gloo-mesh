package install

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

var manifest = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: trafficsplits.split.smi-spec.io
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.service
    description: The service
    name: Service
    type: string
  group: split.smi-spec.io
  names:
    kind: TrafficSplit
    listKind: TrafficSplitList
    plural: trafficsplits
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: traffictargets.access.smi-spec.io
spec:
  group: access.smi-spec.io
  version: v1alpha1
  scope: Namespaced
  names:
    kind: TrafficTarget
    shortNames:
    - tt
    plural: traffictargets
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: httproutegroups.specs.smi-spec.io
spec:
  group: specs.smi-spec.io
  version: v1alpha1
  scope: Namespaced
  names:
    kind: HTTPRouteGroup
    shortNames:
    - htr
    plural: httproutegroups
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: tcproutes.specs.smi-spec.io
spec:
  group: specs.smi-spec.io
  version: v1alpha1
  scope: Namespaced
  names:
    kind: TCPRoute
    shortNames:
    - tr
    plural: tcproutes
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: smi-adapter-istio
  namespace: {{ .InstallNamespace }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: smi-adapter-istio
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - services
  - endpoints
  - persistentvolumeclaims
  - events
  - configmaps
  - secrets
  verbs:
  - '*'
- apiGroups:
  - apps
  resources:
  - deployments
  - daemonsets
  - replicasets
  - statefulsets
  verbs:
  - '*'
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - get
  - create
- apiGroups:
  - apps
  resourceNames:
  - smi-adapter-istio
  resources:
  - deployments/finalizers
  verbs:
  - update
- apiGroups:
  - split.smi-spec.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - networking.istio.io
  resources:
  - '*'
  verbs:
  - '*'
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: smi-adapter-istio
subjects:
- kind: ServiceAccount
  name: smi-adapter-istio
  namespace: {{ .InstallNamespace }}
roleRef:
  kind: ClusterRole
  name: smi-adapter-istio
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: smi-adapter-istio
  namespace: {{ .InstallNamespace }}
spec:
  replicas: 1
  selector:
    matchLabels:
      name: smi-adapter-istio
  template:
    metadata:
      labels:
        name: smi-adapter-istio
      annotations:
        sidecar.istio.io/inject: "false"
    spec:
      serviceAccountName: smi-adapter-istio
      containers:
      - name: smi-adapter-istio
        image: docker.io/stefanprodan/smi-adapter-istio:0.0.2-beta.1
        command:
        - smi-adapter-istio
        imagePullPolicy: Always
        env:
        - name: WATCH_NAMESPACE
          value: ""
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: OPERATOR_NAME
          value: "smi-adapter-istio"
`

func renderSmiAdapterManifest(installNamespace string) (string, error) {
	tmpl, err := template.New("smi-manifest-template").Parse(manifest)
	if err != nil {
		return "", errors.Wrapf(err, "parsing SMI manifest template")
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, struct{ InstallNamespace string }{installNamespace})
	if err != nil {
		return "", errors.Wrapf(err, "executing SMI manifest template")
	}

	return buf.String(), nil
}
