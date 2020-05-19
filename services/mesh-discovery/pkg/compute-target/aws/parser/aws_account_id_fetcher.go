package aws_utils

import (
	"context"

	"github.com/keikoproj/aws-auth/pkg/mapper"
	k8s_core "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

var (
	AwsAuthConfigMapKey = client.ObjectKey{Name: "aws-auth", Namespace: "kube-system"}
)

type awsAccountIdFetcher struct {
	arnParser              ArnParser
	configMapClientFactory k8s_core.ConfigMapClientFactory
}

func NewAwsAccountIdFetcher(
	arnParser ArnParser,
	configMapClientFactory k8s_core.ConfigMapClientFactory,
) AwsAccountIdFetcher {
	return &awsAccountIdFetcher{
		arnParser:              arnParser,
		configMapClientFactory: configMapClientFactory,
	}
}

func (a *awsAccountIdFetcher) GetEksAccountId(
	ctx context.Context,
	clusterScopedClient client.Client,
) (AwsAccountId, error) {
	configMap, err := a.configMapClientFactory(clusterScopedClient).GetConfigMap(ctx, AwsAuthConfigMapKey)
	if err != nil && !errors.IsNotFound(err) {
		return "", err
	}
	if configMap == nil {
		return "", nil
	}
	var authData mapper.AwsAuthData
	mapRoles, ok := configMap.Data["mapRoles"]
	if !ok {
		return "", nil
	}
	err = yaml.Unmarshal([]byte(mapRoles), &authData.MapRoles)
	if err != nil {
		return "", err
	}
	for _, mapRole := range authData.MapRoles {
		if mapRole.Username == "system:node:{{EC2PrivateDNSName}}" {
			accountID, err := a.arnParser.ParseAccountID(mapRole.RoleARN)
			if err != nil {
				return "", err
			}
			return AwsAccountId(accountID), nil
		}
	}
	return "", nil
}
