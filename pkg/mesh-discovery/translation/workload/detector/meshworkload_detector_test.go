package detector_test

import (
	"context"

	"github.com/solo-io/skv2/pkg/ezkube"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	corev1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	v1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/utils"
	. "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/detector"
	mock_detector "github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/detector/mocks"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/workload/types"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

var _ = Describe("WorkloadDetector", func() {

	var (
		ctrl                *gomock.Controller
		mockSidecarDetector *mock_detector.MockSidecarDetector
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockSidecarDetector = mock_detector.NewMockSidecarDetector(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	deploymentName := "name"
	deploymentNs := "namespace"
	deploymentCluster := "cluster"
	serviceAccountName := "service-account-name"
	podLabels := map[string]string{"a": "b"}

	makeDeployment := func() *appsv1.Deployment {
		return &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   deploymentNs,
				ClusterName: deploymentCluster,
				Name:        deploymentName,
			},
			Spec: appsv1.DeploymentSpec{
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: podLabels,
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Image: "istio-pilot:latest",
							},
						},
						ServiceAccountName: serviceAccountName,
					},
				},
			},
		}
	}

	makeReplicaSet := func(dep *appsv1.Deployment) *appsv1.ReplicaSet {
		rs := &appsv1.ReplicaSet{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   deploymentNs,
				ClusterName: deploymentCluster,
				Name:        "replicaset",
			},
		}
		err := controllerutil.SetControllerReference(dep, rs, scheme.Scheme)
		Expect(err).NotTo(HaveOccurred())
		return rs
	}

	makePod := func(rs *appsv1.ReplicaSet) *corev1.Pod {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   deploymentNs,
				Name:        "pod",
				ClusterName: deploymentCluster,
			},
			Status: corev1.PodStatus{
				PodIP: "ip-1",
			},
		}
		err := controllerutil.SetControllerReference(rs, pod, scheme.Scheme)
		Expect(err).NotTo(HaveOccurred())
		return pod
	}

	mesh := &v1.Mesh{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mesh",
			Namespace: "gloo-mesh",
		},
	}

	It("translates a deployment with a detected sidecar to a workload", func() {

		deployment := makeDeployment()
		rs := makeReplicaSet(deployment)
		pod := makePod(rs)

		pods := corev1sets.NewPodSet(pod)
		replicaSets := appsv1sets.NewReplicaSetSet(rs)
		detector := NewWorkloadDetector(
			context.TODO(),
			pods,
			replicaSets,
			mockSidecarDetector,
		)

		meshes := v1sets.NewMeshSet()

		mockSidecarDetector.EXPECT().DetectMeshSidecar(pod, meshes).Return(mesh)

		workload := detector.DetectWorkload(types.ToWorkload(deployment), meshes)

		outputMeta := utils.DiscoveredObjectMeta(deployment)
		// expect appended workload kind
		outputMeta.Name += "-deployment"

		Expect(workload).To(Equal(&v1.Workload{
			ObjectMeta: outputMeta,
			Spec: v1.WorkloadSpec{
				Type: &v1.WorkloadSpec_Kubernetes{
					Kubernetes: &v1.WorkloadSpec_KubernetesWorkload{
						Controller:         ezkube.MakeClusterObjectRef(deployment),
						PodLabels:          podLabels,
						ServiceAccountName: serviceAccountName,
					},
				},
				Mesh: ezkube.MakeObjectRef(mesh),
			},
		}))
	})

})
