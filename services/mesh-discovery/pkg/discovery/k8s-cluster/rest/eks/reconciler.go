package eks

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/clients"
	"github.com/solo-io/service-mesh-hub/pkg/env"
	"github.com/solo-io/service-mesh-hub/pkg/metadata"
	"github.com/solo-io/service-mesh-hub/services/common/constants"
	compute_target_aws "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/compute-target/aws"
	eks_client "github.com/solo-io/service-mesh-hub/services/mesh-discovery/pkg/compute-target/aws/clients/eks"
	"github.com/solo-io/skv2/pkg/multicluster/discovery/cloud"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	MaxResults = 100
)

var (
	DiscoveryLabels = map[string]string{constants.DISCOVERED_BY: constants.EKS_CLUSTER_DISCOVERY}
)

type eksReconciler struct {
	kubeClusterClient         v1alpha1.KubernetesClusterClient
	eksClientFactory          eks_client.EksClientFactory
	eksConfigBuilderFactory   eks_client.EksConfigBuilderFactory
	clusterRegistrationClient clients.ClusterRegistrationClient
}

func NewEksDiscoveryReconciler(
	kubeClusterClient v1alpha1.KubernetesClusterClient,
	eksClientFactory eks_client.EksClientFactory,
	eksConfigBuilderFactory eks_client.EksConfigBuilderFactory,
	clusterRegistrationClient clients.ClusterRegistrationClient,
) compute_target_aws.EksDiscoveryReconciler {
	return &eksReconciler{
		kubeClusterClient:         kubeClusterClient,
		eksClientFactory:          eksClientFactory,
		eksConfigBuilderFactory:   eksConfigBuilderFactory,
		clusterRegistrationClient: clusterRegistrationClient,
	}
}

func (e *eksReconciler) Reconcile(ctx context.Context, creds *credentials.Credentials, region string) error {
	eksClient, err := e.eksClientFactory(creds, region)
	if err != nil {
		return err
	}
	clusterNamesOnAWS, err := e.fetchEksClustersOnAWS(ctx, eksClient, region)
	if err != nil {
		return err
	}
	clusterNamesOnSMH, err := e.fetchEksClustersOnSMH(ctx)
	if err != nil {
		return err
	}
	clustersToRegister := clusterNamesOnAWS.Difference(clusterNamesOnSMH)
	// TODO deregister clusters that are no longer on AWS
	//clustersToDeregister := clusterNamesOnSMH.Difference(clusterNamesOnAWS)
	for clusterName, _ := range clustersToRegister {
		err := e.registerCluster(ctx, eksClient, clusterName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *eksReconciler) fetchEksClustersOnAWS(
	ctx context.Context,
	eksClient cloud.EksClient,
	region string,
) (sets.String, error) {
	var clusterNames []string
	input := &eks.ListClustersInput{
		MaxResults: aws.Int64(MaxResults),
	}
	for {
		listClustersOutput, err := eksClient.ListClusters(ctx, input)
		if err != nil {
			return nil, err
		}
		for _, clusterNameRef := range listClustersOutput.Clusters {
			clusterName := metadata.BuildEksClusterName(aws.StringValue(clusterNameRef), region)
			clusterNames = append(clusterNames, clusterName)
		}
		if listClustersOutput.NextToken == nil {
			break
		}
		input.NextToken = listClustersOutput.NextToken
	}
	return sets.NewString(clusterNames...), nil
}

func (e *eksReconciler) fetchEksClustersOnSMH(ctx context.Context) (sets.String, error) {
	clusters, err := e.kubeClusterClient.ListKubernetesCluster(ctx, client.MatchingLabels(DiscoveryLabels))
	if err != nil {
		return nil, err
	}
	var clusterNames []string
	for _, cluster := range clusters.Items {
		cluster := cluster
		clusterNames = append(clusterNames, cluster.GetName())
	}
	return sets.NewString(clusterNames...), nil
}

func (e *eksReconciler) registerCluster(
	ctx context.Context,
	eksClient cloud.EksClient,
	clusterName string,
) error {
	cluster, err := eksClient.DescribeCluster(ctx, clusterName)
	if err != nil {
		return err
	}
	config, err := e.eksConfigBuilderFactory(eksClient).ConfigForCluster(ctx, cluster)
	if err != nil {
		return err
	}
	return e.clusterRegistrationClient.Register(
		ctx,
		config,
		clusterName,
		env.GetWriteNamespace(), // TODO make this configurable
		false,
		false,
		"",
		"",
		DiscoveryLabels,
	)
}
