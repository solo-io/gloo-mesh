package cluster_registration_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	mock_crd_uninstall "github.com/solo-io/service-mesh-hub/cli/pkg/tree/uninstall/crd/mocks"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	v1 "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	zephyr_security_scheme "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/auth"
	"github.com/solo-io/service-mesh-hub/pkg/clients"
	cluster_registration "github.com/solo-io/service-mesh-hub/pkg/clients/cluster-registration"
	"github.com/solo-io/service-mesh-hub/pkg/constants"
	"github.com/solo-io/service-mesh-hub/pkg/env"
	"github.com/solo-io/service-mesh-hub/pkg/factories"
	"github.com/solo-io/service-mesh-hub/pkg/installers/csr"
	mock_csr "github.com/solo-io/service-mesh-hub/pkg/installers/csr/mocks"
	"github.com/solo-io/service-mesh-hub/pkg/kubeconfig"
	mock_kubeconfig "github.com/solo-io/service-mesh-hub/pkg/kubeconfig/mocks"
	cert_secrets "github.com/solo-io/service-mesh-hub/pkg/security/secrets"
	mock_mc_manager "github.com/solo-io/service-mesh-hub/services/common/compute-target/k8s/mocks"
	mock_zephyr_discovery_clients "github.com/solo-io/service-mesh-hub/test/mocks/clients/discovery.zephyr.solo.io/v1alpha1"
	mock_k8s_core_clients "github.com/solo-io/service-mesh-hub/test/mocks/clients/kubernetes/core/v1"
	v12 "k8s.io/api/core/v1"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Cluster Deregistration", func() {
	var (
		ctrl             *gomock.Controller
		ctx              context.Context
		remoteRestConfig = &rest.Config{
			Host: "remote-cluster.com",
		}
		remoteClientConfig = &clientcmd.DirectClientConfig{}
	)

	BeforeEach(func() {
		ctx = context.TODO()
		ctrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("can deregister a cluster", func() {
		mockCsrAgentInstaller := mock_csr.NewMockCsrAgentInstaller(ctrl)
		crdRemover := mock_crd_uninstall.NewMockCrdRemover(ctrl)
		configLookup := mock_kubeconfig.NewMockKubeConfigLookup(ctrl)
		kubeClusterClient := mock_zephyr_discovery_clients.NewMockKubernetesClusterClient(ctrl)
		localSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		remoteSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		dynamicClientGetter := mock_mc_manager.NewMockDynamicClientGetter(ctrl)
		remoteServiceAccountClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)

		kubeConfigSecretRef := &zephyr_core_types.ResourceRef{
			Name:      "kube-config-secret",
			Namespace: env.GetWriteNamespace(),
		}
		kubeConfigSecret := &v12.Secret{
			ObjectMeta: clients.ResourceRefToObjectMeta(kubeConfigSecretRef),
		}
		remoteClusterName := "remote-cluster-name"
		remoteWriteNamespace := "remote-write-namespace"
		clusterToDeregister := &zephyr_discovery.KubernetesCluster{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      remoteClusterName,
				Namespace: env.GetWriteNamespace(),
			},
			Spec: zephyr_discovery_types.KubernetesClusterSpec{
				SecretRef:      kubeConfigSecretRef,
				WriteNamespace: remoteWriteNamespace,
			},
		}
		intermediateCertSecret := &v12.Secret{
			Type: cert_secrets.IntermediateCertSecretType,
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      "intermediate-cert",
				Namespace: remoteWriteNamespace,
			},
		}

		kubeRestConfig := &kubeconfig.ConvertedConfigs{
			RestConfig:   remoteRestConfig,
			ClientConfig: remoteClientConfig,
		}
		mockCsrAgentInstaller.EXPECT().Uninstall(&csr.CsrAgentUninstallOptions{
			KubeConfig:       csr.KubeConfig{KubeConfig: remoteClientConfig},
			ReleaseName:      cliconstants.CsrAgentReleaseName,
			ReleaseNamespace: clusterToDeregister.Spec.GetWriteNamespace(),
		}).Return(nil)
		configLookup.EXPECT().
			FromCluster(ctx, clusterToDeregister.GetName()).
			Return(kubeRestConfig, nil)
		dynamicClientGetter.EXPECT().
			GetClientForCluster(ctx, remoteClusterName).
			Return(nil, nil)
		remoteSecretClient.EXPECT().
			ListSecret(ctx, client.InNamespace(remoteWriteNamespace)).
			Return(&v12.SecretList{
				Items: []v12.Secret{*intermediateCertSecret},
			}, nil)
		remoteSecretClient.EXPECT().
			DeleteSecret(ctx, clients.ObjectMetaToObjectKey(intermediateCertSecret.ObjectMeta)).
			Return(nil)
		localSecretClient.EXPECT().
			GetSecret(ctx, clients.ResourceRefToObjectKey(clusterToDeregister.Spec.GetSecretRef())).
			Return(kubeConfigSecret, nil)
		localSecretClient.EXPECT().
			DeleteSecret(ctx, clients.ObjectMetaToObjectKey(kubeConfigSecret.ObjectMeta)).
			Return(nil)
		kubeClusterClient.EXPECT().
			DeleteKubernetesCluster(ctx, clients.ObjectMetaToObjectKey(clusterToDeregister.ObjectMeta)).
			Return(nil)
		remoteServiceAccountClient.EXPECT().
			DeleteAllOfServiceAccount(
				ctx,
				client.InNamespace(remoteWriteNamespace),
				client.MatchingLabels{
					constants.ManagedByLabel:        constants.ServiceMeshHubApplicationName,
					auth.RegistrationServiceAccount: auth.RegistrationServiceAccountValue,
				},
			).
			Return(nil)
		crdRemover.EXPECT().
			RemoveCrdGroup(ctx, clusterToDeregister.GetName(), kubeRestConfig.RestConfig, zephyr_security_scheme.SchemeGroupVersion).
			Return(false, nil)

		clusterDeregistrationClient := cluster_registration.NewClusterDeregistrationClient(
			crdRemover,
			func(helmInstallerFactory factories.HelmInstallerFactory) csr.CsrAgentInstaller {
				return mockCsrAgentInstaller
			},
			configLookup,
			kubeClusterClient,
			localSecretClient,
			func(_ client.Client) v1.SecretClient {
				return remoteSecretClient
			},
			dynamicClientGetter,
			func(_ client.Client) v1.ServiceAccountClient {
				return remoteServiceAccountClient
			},
		)

		err := clusterDeregistrationClient.Deregister(ctx, clusterToDeregister)
		Expect(err).NotTo(HaveOccurred())
	})

	It("responds with the appropriate error if the config lookup fails", func() {
		testErr := eris.New("test-err")
		crdRemover := mock_crd_uninstall.NewMockCrdRemover(ctrl)
		mockCsrAgentInstaller := mock_csr.NewMockCsrAgentInstaller(ctrl)
		configLookup := mock_kubeconfig.NewMockKubeConfigLookup(ctrl)
		kubeClusterClient := mock_zephyr_discovery_clients.NewMockKubernetesClusterClient(ctrl)
		localSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		dynamicClientGetter := mock_mc_manager.NewMockDynamicClientGetter(ctrl)
		kubeConfigSecretRef := &zephyr_core_types.ResourceRef{
			Name:      "kube-config-secret",
			Namespace: env.GetWriteNamespace(),
		}
		remoteWriteNamespace := "remote-write-namespace"
		remoteClusterName := "remote-cluster-name"
		clusterToDeregister := &zephyr_discovery.KubernetesCluster{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      remoteClusterName,
				Namespace: env.GetWriteNamespace(),
			},
			Spec: zephyr_discovery_types.KubernetesClusterSpec{
				SecretRef:      kubeConfigSecretRef,
				WriteNamespace: remoteWriteNamespace,
			},
		}

		configLookup.EXPECT().
			FromCluster(ctx, clusterToDeregister.GetName()).
			Return(nil, testErr)

		clusterDeregistrationClient := cluster_registration.NewClusterDeregistrationClient(
			crdRemover,
			func(helmInstallerFactory factories.HelmInstallerFactory) csr.CsrAgentInstaller {
				return mockCsrAgentInstaller
			},
			configLookup,
			kubeClusterClient,
			localSecretClient,
			func(_ client.Client) v1.SecretClient {
				return nil
			},
			dynamicClientGetter,
			func(_ client.Client) v1.ServiceAccountClient {
				return nil
			},
		)

		err := clusterDeregistrationClient.Deregister(ctx, clusterToDeregister)
		Expect(err).To(testutils.HaveInErrorChain(cluster_registration.FailedToFindClusterCredentials(testErr, remoteClusterName)))
	})

	It("responds with the appropriate error if CSR uninstallation fails", func() {
		testErr := eris.New("test-err")
		mockCsrAgentInstaller := mock_csr.NewMockCsrAgentInstaller(ctrl)
		crdRemover := mock_crd_uninstall.NewMockCrdRemover(ctrl)
		configLookup := mock_kubeconfig.NewMockKubeConfigLookup(ctrl)
		kubeClusterClient := mock_zephyr_discovery_clients.NewMockKubernetesClusterClient(ctrl)
		localSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		remoteSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		dynamicClientGetter := mock_mc_manager.NewMockDynamicClientGetter(ctrl)
		remoteServiceAccountClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		kubeConfigSecretRef := &zephyr_core_types.ResourceRef{
			Name:      "kube-config-secret",
			Namespace: env.GetWriteNamespace(),
		}
		remoteClusterName := "remote-cluster-name"
		remoteWriteNamespace := "remote-write-namespace"
		clusterToDeregister := &zephyr_discovery.KubernetesCluster{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      remoteClusterName,
				Namespace: env.GetWriteNamespace(),
			},
			Spec: zephyr_discovery_types.KubernetesClusterSpec{
				SecretRef:      kubeConfigSecretRef,
				WriteNamespace: remoteWriteNamespace,
			},
		}
		configLookup.EXPECT().
			FromCluster(ctx, clusterToDeregister.GetName()).
			Return(&kubeconfig.ConvertedConfigs{
				RestConfig:   remoteRestConfig,
				ClientConfig: remoteClientConfig,
			}, nil)
		mockCsrAgentInstaller.EXPECT().Uninstall(&csr.CsrAgentUninstallOptions{
			KubeConfig:       csr.KubeConfig{KubeConfig: remoteClientConfig},
			ReleaseName:      cliconstants.CsrAgentReleaseName,
			ReleaseNamespace: clusterToDeregister.Spec.GetWriteNamespace(),
		}).Return(testErr)

		clusterDeregistrationClient := cluster_registration.NewClusterDeregistrationClient(
			crdRemover,
			func(helmInstallerFactory factories.HelmInstallerFactory) csr.CsrAgentInstaller {
				return mockCsrAgentInstaller
			},
			configLookup,
			kubeClusterClient,
			localSecretClient,
			func(_ client.Client) v1.SecretClient {
				return remoteSecretClient
			},
			dynamicClientGetter,
			func(_ client.Client) v1.ServiceAccountClient {
				return remoteServiceAccountClient
			},
		)

		err := clusterDeregistrationClient.Deregister(ctx, clusterToDeregister)
		Expect(err).To(testutils.HaveInErrorChain(cluster_registration.FailedToUninstallCsrAgent(testErr, remoteClusterName)))
	})

	It("responds with the appropriate error if CRD removal fails", func() {
		testErr := eris.New("test-err")
		mockCsrAgentInstaller := mock_csr.NewMockCsrAgentInstaller(ctrl)
		crdRemover := mock_crd_uninstall.NewMockCrdRemover(ctrl)
		configLookup := mock_kubeconfig.NewMockKubeConfigLookup(ctrl)
		kubeClusterClient := mock_zephyr_discovery_clients.NewMockKubernetesClusterClient(ctrl)
		localSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		remoteSecretClient := mock_k8s_core_clients.NewMockSecretClient(ctrl)
		dynamicClientGetter := mock_mc_manager.NewMockDynamicClientGetter(ctrl)
		remoteServiceAccountClient := mock_k8s_core_clients.NewMockServiceAccountClient(ctrl)
		kubeConfigSecretRef := &zephyr_core_types.ResourceRef{
			Name:      "kube-config-secret",
			Namespace: env.GetWriteNamespace(),
		}
		kubeConfigSecret := &v12.Secret{
			ObjectMeta: clients.ResourceRefToObjectMeta(kubeConfigSecretRef),
		}
		remoteClusterName := "remote-cluster-name"
		remoteWriteNamespace := "remote-write-namespace"
		clusterToDeregister := &zephyr_discovery.KubernetesCluster{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      remoteClusterName,
				Namespace: env.GetWriteNamespace(),
			},
			Spec: zephyr_discovery_types.KubernetesClusterSpec{
				SecretRef:      kubeConfigSecretRef,
				WriteNamespace: remoteWriteNamespace,
			},
		}

		intermediateCertSecret := &v12.Secret{
			Type: cert_secrets.IntermediateCertSecretType,
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      "intermediate-cert",
				Namespace: remoteWriteNamespace,
			},
		}
		configLookup.EXPECT().
			FromCluster(ctx, clusterToDeregister.GetName()).
			Return(&kubeconfig.ConvertedConfigs{
				RestConfig:   remoteRestConfig,
				ClientConfig: remoteClientConfig,
			}, nil)
		mockCsrAgentInstaller.EXPECT().Uninstall(&csr.CsrAgentUninstallOptions{
			KubeConfig:       csr.KubeConfig{KubeConfig: remoteClientConfig},
			ReleaseName:      cliconstants.CsrAgentReleaseName,
			ReleaseNamespace: clusterToDeregister.Spec.GetWriteNamespace(),
		}).Return(nil)
		crdRemover.EXPECT().
			RemoveCrdGroup(ctx, clusterToDeregister.GetName(), remoteRestConfig, zephyr_security_scheme.SchemeGroupVersion).
			Return(false, testErr)
		dynamicClientGetter.EXPECT().
			GetClientForCluster(ctx, remoteClusterName).
			Return(nil, nil)
		remoteSecretClient.EXPECT().
			ListSecret(ctx, client.InNamespace(remoteWriteNamespace)).
			Return(&v12.SecretList{
				Items: []v12.Secret{*intermediateCertSecret},
			}, nil)
		remoteSecretClient.EXPECT().
			DeleteSecret(ctx, clients.ObjectMetaToObjectKey(intermediateCertSecret.ObjectMeta)).
			Return(nil)
		localSecretClient.EXPECT().
			GetSecret(ctx, clients.ResourceRefToObjectKey(clusterToDeregister.Spec.GetSecretRef())).
			Return(kubeConfigSecret, nil)
		localSecretClient.EXPECT().
			DeleteSecret(ctx, clients.ObjectMetaToObjectKey(kubeConfigSecret.ObjectMeta)).
			Return(nil)
		kubeClusterClient.EXPECT().
			DeleteKubernetesCluster(ctx, clients.ObjectMetaToObjectKey(clusterToDeregister.ObjectMeta)).
			Return(nil)

		clusterDeregistrationClient := cluster_registration.NewClusterDeregistrationClient(
			crdRemover,
			func(helmInstallerFactory factories.HelmInstallerFactory) csr.CsrAgentInstaller {
				return mockCsrAgentInstaller
			},
			configLookup,
			kubeClusterClient,
			localSecretClient,
			func(_ client.Client) v1.SecretClient {
				return remoteSecretClient
			},
			dynamicClientGetter,
			func(_ client.Client) v1.ServiceAccountClient {
				return remoteServiceAccountClient
			},
		)

		err := clusterDeregistrationClient.Deregister(ctx, clusterToDeregister)
		Expect(err).To(testutils.HaveInErrorChain(cluster_registration.FailedToRemoveCrds(testErr, remoteClusterName)))
	})
})
