package appmesh

import (
	"context"
	"fmt"

	aws_v1beta2 "github.com/aws/aws-app-mesh-controller-for-k8s/apis/appmesh/v1beta2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/input"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/mesh-discovery/translation/mesh/detector"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("AppMesh MeshDetector", func() {

	var (
		cluster1  = "cluster1"
		meshName1 = "mesh1"
		region    = "us-east-2"
		accountID = "1234"
		meshArn1  = fmt.Sprintf("arn:aws:appmesh:%s:%s:mesh/%s", region, accountID, meshName1)

		meshDetector detector.MeshDetector
	)

	BeforeEach(func() {
		meshDetector = NewMeshDetector(context.Background())
	})

	It("detects one app mesh instance across two clusters", func() {
		awsMesh1 := aws_v1beta2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:        meshName1,
				ClusterName: cluster1,
			},
			Spec: aws_v1beta2.MeshSpec{
				AWSName: &meshName1,
				NamespaceSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"mesh": meshName1},
				},
			},
			Status: aws_v1beta2.MeshStatus{
				MeshARN: &meshArn1,
			},
		}

		cluster2 := "cluster2"
		awsMesh2 := aws_v1beta2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:        meshName1,
				ClusterName: cluster2,
			},
			Spec: aws_v1beta2.MeshSpec{
				AWSName: &meshName1,
				NamespaceSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"mesh": meshName1},
				},
			},
			Status: aws_v1beta2.MeshStatus{
				MeshARN: &meshArn1,
			},
		}

		expected := v1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", meshName1, "us-east-2", "1234"),
				Namespace: "gloo-mesh",
				Labels: map[string]string{
					"owner.discovery.mesh.gloo.solo.io": "gloo-mesh",
				},
			},
			Spec: v1.MeshSpec{
				Type: &v1.MeshSpec_AwsAppMesh_{
					AwsAppMesh: &v1.MeshSpec_AwsAppMesh{
						AwsName:      meshName1,
						Region:       "us-east-2",
						AwsAccountId: "1234",
						Arn:          meshArn1,
						Clusters:     []string{cluster1, cluster2},
					},
				},
			},
		}

		builder := input.NewInputDiscoveryInputSnapshotManualBuilder("app mesh test")
		builder.AddMeshes([]*aws_v1beta2.Mesh{&awsMesh1, &awsMesh2})

		actual, err := meshDetector.DetectMeshes(builder.Build(), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(HaveLen(1))
		Expect(actual).To(ContainElement(&expected))
	})

	It("detects disparate app mesh instances on one cluster", func() {
		awsMesh1 := aws_v1beta2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:        meshName1,
				ClusterName: cluster1,
			},
			Spec: aws_v1beta2.MeshSpec{
				AWSName: &meshName1,
				NamespaceSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"mesh": meshName1},
				},
			},
			Status: aws_v1beta2.MeshStatus{
				MeshARN: &meshArn1,
			},
		}

		meshName2 := "mesh2"
		meshArn2 := fmt.Sprintf("arn:aws:appmesh:%s:%s:mesh/%s", region, accountID, meshName2)
		awsMesh2 := aws_v1beta2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:        meshName2,
				ClusterName: cluster1,
			},
			Spec: aws_v1beta2.MeshSpec{
				AWSName: &meshName2,
				NamespaceSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"mesh": meshName2},
				},
			},
			Status: aws_v1beta2.MeshStatus{
				MeshARN: &meshArn2,
			},
		}

		expected1 := v1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", meshName1, "us-east-2", "1234"),
				Namespace: "gloo-mesh",
				Labels: map[string]string{
					"owner.discovery.mesh.gloo.solo.io": "gloo-mesh",
				},
			},
			Spec: v1.MeshSpec{
				Type: &v1.MeshSpec_AwsAppMesh_{
					AwsAppMesh: &v1.MeshSpec_AwsAppMesh{
						AwsName:      meshName1,
						Region:       "us-east-2",
						AwsAccountId: "1234",
						Arn:          meshArn1,
						Clusters:     []string{cluster1},
					},
				},
			},
		}
		expected2 := v1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", meshName2, "us-east-2", "1234"),
				Namespace: "gloo-mesh",
				Labels: map[string]string{
					"owner.discovery.mesh.gloo.solo.io": "gloo-mesh",
				},
			},
			Spec: v1.MeshSpec{
				Type: &v1.MeshSpec_AwsAppMesh_{
					AwsAppMesh: &v1.MeshSpec_AwsAppMesh{
						AwsName:      meshName2,
						Region:       "us-east-2",
						AwsAccountId: "1234",
						Arn:          meshArn2,
						Clusters:     []string{cluster1},
					},
				},
			},
		}

		builder := input.NewInputDiscoveryInputSnapshotManualBuilder("app mesh test")
		builder.AddMeshes([]*aws_v1beta2.Mesh{&awsMesh1, &awsMesh2})

		actual, err := meshDetector.DetectMeshes(builder.Build(), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(HaveLen(2))
		Expect(actual).To(ContainElement(&expected1))
		Expect(actual).To(ContainElement(&expected2))
	})

	It("does not detect meshes that haven't been assigned an ARN", func() {
		awsMesh := aws_v1beta2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:        meshName1,
				ClusterName: cluster1,
			},
			Spec: aws_v1beta2.MeshSpec{
				AWSName: &meshName1,
				NamespaceSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"mesh": meshName1},
				},
			},
			Status: aws_v1beta2.MeshStatus{},
		}

		builder := input.NewInputDiscoveryInputSnapshotManualBuilder("app mesh test")
		builder.AddMeshes([]*aws_v1beta2.Mesh{&awsMesh})

		actual, err := meshDetector.DetectMeshes(builder.Build(), nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(HaveLen(0))
	})

	It("errors when an ARN is malformed", func() {
		badArn := "this can't be parsed as a Mesh ARN"

		awsMesh := aws_v1beta2.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name:        meshName1,
				ClusterName: cluster1,
			},
			Spec: aws_v1beta2.MeshSpec{
				AWSName: &meshName1,
				NamespaceSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"mesh": meshName1},
				},
			},
			Status: aws_v1beta2.MeshStatus{
				MeshARN: &badArn,
			},
		}

		builder := input.NewInputDiscoveryInputSnapshotManualBuilder("app mesh test")
		builder.AddMeshes([]*aws_v1beta2.Mesh{&awsMesh})

		_, err := meshDetector.DetectMeshes(builder.Build(), nil)
		Expect(err).To(HaveOccurred())
	})

})
