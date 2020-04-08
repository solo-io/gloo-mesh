package operator

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
)

//go:generate mockgen -source ./manifest.go -destination mocks/mock_manifest_builder.go
type InstallerManifestBuilder interface {
	// Based on the pending installation config, generate an appropriate installation manifest
	Build(options *options.MeshInstallationConfig) (installationManifest string, err error)

	// Generate an IstioOperator spec that sets up Mesh with its demo profile
	GetOperatorSpecWithProfile(profile, installationNamespace string) (string, error)
}

var (
	InvalidProfileFound = func(profile string) error {
		return eris.Errorf(
			"invalid profile (%s) found, valid options are: [%s]",
			profile,
			strings.Join(ValidProfiles.List(), ","),
		)
	}
)

func NewInstallerManifestBuilder() InstallerManifestBuilder {
	return &installerManifestBuilder{}
}

type installerManifestBuilder struct{}

func (i *installerManifestBuilder) Build(options *options.MeshInstallationConfig) (string, error) {
	tmpl := template.New("")
	tmpl, err := tmpl.Parse(installationManifestTemplate)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, options)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (i *installerManifestBuilder) GetOperatorSpecWithProfile(profile, namespace string) (string, error) {
	if !ValidProfiles.Has(profile) {
		return "", InvalidProfileFound(profile)
	}

	tmpl := template.New("")
	tmpl, err := tmpl.Parse(istioOperatorWithProfile)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	err = tmpl.Execute(&buffer, &controlPlaneData{
		Profile:          profile,
		InstallNamespace: namespace,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

type controlPlaneData struct {
	Profile          string
	InstallNamespace string
}

// the raw yaml was obtained from `https://istio.io/operator.yaml` as suggested by https://preliminary.istio.io/docs/setup/install/standalone-operator/
var installationManifestTemplate = `
{{- if .CreateNamespace }}
---
apiVersion: v1
kind: Namespace
metadata:
  name: {{ .InstallNamespace }}
...
{{- end }}
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: istiooperators.install.istio.io
spec:
  group: install.istio.io
  versions:
    - name: v1alpha1
      served: true
      storage: true
  scope: Namespaced
  subresources:
    status: {}
  names:
    kind: IstioOperator
    listKind: IstioOperatorList
    plural: istiooperators
    singular: istiooperator
    shortNames:
    - iop
...
---
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{ .InstallNamespace }}
  name: istio-operator
...
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: istio-operator
rules:
# istio groups
- apiGroups:
  - authentication.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - config.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - install.istio.io
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
- apiGroups:
  - rbac.istio.io
  resources:
  - '*'
  verbs:
  - '*'
- apiGroups:
  - security.istio.io
  resources:
  - '*'
  verbs:
  - '*'
# k8s groups
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  - validatingwebhookconfigurations
  verbs:
  - '*'
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions.apiextensions.k8s.io
  - customresourcedefinitions
  verbs:
  - '*'
- apiGroups:
  - apps
  - extensions
  resources:
  - daemonsets
  - deployments
  - deployments/finalizers
  - ingresses
  - replicasets
  - statefulsets
  verbs:
  - '*'
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
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
  - policy
  resources:
  - poddisruptionbudgets
  verbs:
  - '*'
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - clusterrolebindings
  - clusterroles
  - roles
  - rolebindings
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - configmaps
  - endpoints
  - events
  - namespaces
  - pods
  - persistentvolumeclaims
  - secrets
  - services
  - serviceaccounts
  verbs:
  - '*'
...
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: istio-operator
subjects:
- kind: ServiceAccount
  name: istio-operator
  namespace: {{ .InstallNamespace }}
roleRef:
  kind: ClusterRole
  name: istio-operator
  apiGroup: rbac.authorization.k8s.io
...
---
apiVersion: v1
kind: Service
metadata:
  namespace: {{ .InstallNamespace }}
  labels:
    name: istio-operator
  name: istio-operator-metrics
spec:
  ports:
  - name: http-metrics
    port: 8383
    targetPort: 8383
  selector:
    name: istio-operator
...
---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: {{ .InstallNamespace }}
  name: istio-operator
spec:
  replicas: 1
  selector:
    matchLabels:
      name: istio-operator
  template:
    metadata:
      labels:
        name: istio-operator
    spec:
      serviceAccountName: istio-operator
      containers:
        - name: istio-operator
          image: docker.io/istio/operator:1.5.1
          command:
          - operator
          - server
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              cpu: 200m
              memory: 256Mi
            requests:
              cpu: 50m
              memory: 128Mi
          env:
            - name: WATCH_NAMESPACE
              value: {{ .InstallNamespace }}
            - name: LEADER_ELECTION_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: OPERATOR_NAME
              value: "istio-operator"
...

`

var istioOperatorWithProfile = `
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: {{ .InstallNamespace }}
  name: istiocontrolplane-{{ .Profile }}
spec:
  profile: {{ .Profile }}

`
