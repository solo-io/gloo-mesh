package register

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	common_config "github.com/solo-io/service-mesh-hub/cli/pkg/common/config"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	cluster_internal "github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/internal"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/cluster/register/csr"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	k8s_core "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	"github.com/solo-io/service-mesh-hub/pkg/kubeconfig"
	"github.com/spf13/pflag"
	k8s_core_types "k8s.io/api/core/v1"
	k8s_errs "k8s.io/apimachinery/pkg/api/errors"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	FailedToCheckForPreviousKubeCluster = "Could not get KubernetesCluster resource from master cluster"
)

var (
	FailedLoadingRemoteConfig = func(err error) error {
		return eris.Wrap(err, "Failed to load the kube config for the remote cluster")
	}
	FailedToCreateAuthToken = func(saRef *zephyr_core_types.ResourceRef, remoteKubeConfig, remoteContext string) string {
		return fmt.Sprintf("Failed to create an auth token for service account %s.%s in cluster "+
			"pointed to by kube config %s with context %s. This operation is not atomic, so the service account may "+
			"have been created and left in the cluster while a later step failed. \n",
			saRef.GetNamespace(), saRef.GetName(), remoteKubeConfig, remoteContext)
	}
	FailedToConvertToSecret = func(err error) error {
		return eris.Wrap(err, "Could not convert kube config for new service account into secret")
	}
	FailedToWriteSecret = func(err error) error {
		return eris.Wrap(err, "Could not write secret to master cluster")
	}
	FailedToWriteKubeCluster = func(err error) error {
		return eris.Wrap(err, "Could not write KubernetesCluster resource to master cluster")
	}
	FailedToEnsureWriteNamespace = func(err error, namespace string) error {
		return eris.Wrapf(err, "Failed to find or create namespace %s on the cluster being registered", namespace)
	}
)

// write a new kube config secret to the master cluster containing creds for talking to the remote cluster as the given service account

func RegisterCluster(
	ctx context.Context,
	binaryName string,
	flags *pflag.FlagSet,
	clientsFactory common.ClientsFactory,
	kubeClientsFactory common.KubeClientsFactory,
	opts *options.Options,
	out io.Writer,
	kubeLoader common_config.KubeLoader,
) error {

	if err := cluster_internal.VerifyRemoteContextFlags(opts); err != nil {
		return err
	}

	if err := cluster_internal.VerifyMasterCluster(clientsFactory, opts); err != nil {
		return err
	}

	registerOpts := opts.Cluster.Register

	// set up kube clients for the master cluster
	masterCfg, err := kubeLoader.GetRestConfigForContext(opts.Root.KubeConfig, opts.Root.KubeContext)
	if err != nil {
		return common.FailedLoadingMasterConfig(err)
	}
	masterKubeClients, err := kubeClientsFactory(masterCfg, opts.Root.WriteNamespace)
	if err != nil {
		return err
	}

	// set up kube clients for the remote cluster

	// default the remote kube config/context to the root settings
	remoteConfigPath, remoteContext := opts.Root.KubeConfig, opts.Root.KubeContext
	if registerOpts.RemoteKubeConfig != "" {
		// if we specified a kube config for the remote cluster, use that instead
		remoteConfigPath = registerOpts.RemoteKubeConfig
	}

	// if we didn't have a context from the root, or if we had an override for the
	// remote context, use the remote context instead
	if remoteContext == "" || registerOpts.RemoteContext != "" {
		remoteContext = registerOpts.RemoteContext
	}

	remoteCfg, err := kubeLoader.GetRestConfigForContext(remoteConfigPath, remoteContext)
	if err != nil {
		return FailedLoadingRemoteConfig(err)
	}
	remoteKubeClients, err := kubeClientsFactory(remoteCfg, registerOpts.RemoteWriteNamespace)
	if err != nil {
		return err
	}

	// if overwrite returns ok than the program should continue, else return
	// The reason for the 2 return vars is that err may be nil and returned anyway
	if ok, err := shouldOverwrite(ctx, binaryName, flags, out, opts, masterKubeClients); !ok {
		return err
	}

	if err = ensureRemoteNamespace(ctx, registerOpts.RemoteWriteNamespace, remoteKubeClients); err != nil {
		return err
	}

	configForServiceAccount, err := generateServiceAccountConfig(ctx, out, remoteKubeClients, remoteCfg, registerOpts)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Successfully wrote service account to remote cluster...\n")

	clients, err := clientsFactory(opts)
	if err != nil {
		return eris.Errorf("Unexpected error: Clients should have already been built, but failed to build again")
	}
	// Install CRDs to remote cluster. This must happen BEFORE kubeconfig Secret is written to master cluster because
	// relevant CRDs must exist before SMH attempts any cross cluster functionality.
	csrAgentInstaller := clients.ClusterRegistrationClients.CsrAgentInstallerFactory(remoteKubeClients.HelmInstaller, masterKubeClients.DeployedVersionFinder)
	err = csrAgentInstaller.Install(
		ctx,
		&csr.CsrAgentInstallOptions{
			KubeConfig:           remoteConfigPath,
			KubeContext:          remoteContext,
			ClusterName:          registerOpts.RemoteClusterName,
			SmhInstallNamespace:  opts.Root.WriteNamespace,
			UseDevCsrAgentChart:  registerOpts.UseDevCsrAgentChart,
			ReleaseName:          cliconstants.CsrAgentReleaseName,
			RemoteWriteNamespace: registerOpts.RemoteWriteNamespace,
		},
	)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Successfully set up CSR agent...\n")

	// Write kubeconfig Secret and KubeCluster CRD to master cluster
	secret, err := writeKubeConfigToMaster(
		ctx,
		opts.Root.WriteNamespace,
		registerOpts,
		remoteConfigPath,
		configForServiceAccount,
		masterKubeClients,
		kubeLoader,
	)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "Successfully wrote kube config secret to master cluster...\n")
	err = writeKubeClusterToMaster(ctx, masterKubeClients, opts.Root.WriteNamespace, registerOpts, secret)
	if err != nil {
		return err
	}

	fmt.Fprintf(
		out,
		"\nCluster %s is now registered in your Service Mesh Hub installation\n",
		registerOpts.RemoteClusterName,
	)

	return nil
}

func ensureRemoteNamespace(ctx context.Context, writeNamespace string, remote *common.KubeClients) error {
	_, err := remote.NamespaceClient.GetNamespace(ctx, client.ObjectKey{Name: writeNamespace})
	if k8s_errs.IsNotFound(err) {
		return remote.NamespaceClient.CreateNamespace(ctx, &k8s_core_types.Namespace{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name: writeNamespace,
			},
		})
	} else if err != nil {
		return FailedToEnsureWriteNamespace(err, writeNamespace)
	}
	return nil
}

func shouldOverwrite(
	ctx context.Context,
	binaryName string,
	flags *pflag.FlagSet,
	out io.Writer,
	opts *options.Options,
	masterKubeClients *common.KubeClients,
) (ok bool, err error) {
	if !opts.Cluster.Register.Overwrite {
		_, err = masterKubeClients.KubeClusterClient.GetKubernetesCluster(
			ctx,
			client.ObjectKey{
				Name:      opts.Cluster.Register.RemoteClusterName,
				Namespace: opts.Root.WriteNamespace,
			},
		)
		if err != nil && !k8s_errs.IsNotFound(err) {
			// if kube cluster does not exist for the given name, continue
			fmt.Fprintf(out, FailedToCheckForPreviousKubeCluster)
			return false, err
		} else if err == nil {
			// nil error signifying the object exists
			exampleCommand := []string{binaryName}

			// using flags.Visit rather than mucking around with os.Args because we may be
			// running in a test environment, where os.Args is the test invocation rather than meshctl
			flags.Visit(func(flag *pflag.Flag) {
				exampleCommand = append(exampleCommand, fmt.Sprintf("--%s %s", flag.Name, flag.Value))
			})
			exampleCommand = append(exampleCommand, fmt.Sprintf("--%s", options.ClusterRegisterOverwriteFlag))
			fmt.Fprintf(out, "Cluster already registered; if you would like to update this cluster please run the previous command with the --%s flag: \n\n"+
				"$ %s\n", options.ClusterRegisterOverwriteFlag, strings.Join(exampleCommand, " "))
			return false, nil
		}
	}
	return true, nil
}

func writeKubeConfigToMaster(
	ctx context.Context,
	writeNamespace string,
	registerOpts options.Register,
	remoteKubeConfig string,
	serviceAccountConfig *rest.Config,
	masterKubeClients *common.KubeClients,
	kubeLoader common_config.KubeLoader,
) (*k8s_core_types.Secret, error) {

	// now we need the cluster/context information from that config
	remoteKubeCtx, err := kubeLoader.GetRawConfigForContext(remoteKubeConfig, registerOpts.RemoteContext)
	if err != nil {
		return nil, common_config.FailedToParseContext(err)
	}

	remoteContextName := remoteKubeCtx.CurrentContext
	if registerOpts.RemoteContext != "" {
		remoteContextName = registerOpts.RemoteContext
	}

	remoteContext := remoteKubeCtx.Contexts[remoteContextName]
	remoteCluster := remoteKubeCtx.Clusters[remoteContext.Cluster]

	// hacky step for running locally in KIND
	err = hackClusterConfigForLocalTestingInKIND(remoteCluster, remoteContextName, registerOpts.LocalClusterDomainOverride)
	if err != nil {
		return nil, err
	}

	secret, err := kubeconfig.KubeConfigToSecret(
		registerOpts.RemoteClusterName,
		writeNamespace,
		&kubeconfig.KubeConfig{
			Config: api.Config{
				Kind:        "Secret",
				APIVersion:  "kubernetes_core",
				Preferences: api.Preferences{},
				Clusters: map[string]*api.Cluster{
					registerOpts.RemoteClusterName: remoteCluster,
				},
				AuthInfos: map[string]*api.AuthInfo{
					registerOpts.RemoteClusterName: {
						Token: serviceAccountConfig.BearerToken,
					},
				},
				Contexts: map[string]*api.Context{
					registerOpts.RemoteClusterName: {
						LocationOfOrigin: remoteContext.LocationOfOrigin,
						Cluster:          registerOpts.RemoteClusterName,
						AuthInfo:         registerOpts.RemoteClusterName,
						Namespace:        remoteContext.Namespace,
						Extensions:       remoteContext.Extensions,
					},
				},
				CurrentContext: registerOpts.RemoteClusterName,
			},
			Cluster: registerOpts.RemoteClusterName,
		})

	if err != nil {
		return nil, FailedToConvertToSecret(err)
	}

	err = upsertSecretData(ctx, secret, masterKubeClients.SecretClient)
	if err != nil {
		return nil, FailedToWriteSecret(err)
	}

	return secret, nil
}

// write the KubernetesCluster resource to the master cluster
func writeKubeClusterToMaster(
	ctx context.Context,
	masterKubeClients *common.KubeClients,
	writeNamespace string,
	registerOpts options.Register,
	secret *k8s_core_types.Secret,
) error {
	cluster := &zephyr_discovery.KubernetesCluster{
		ObjectMeta: k8s_meta_types.ObjectMeta{
			Name:      registerOpts.RemoteClusterName,
			Namespace: writeNamespace,
		},
		Spec: zephyr_discovery_types.KubernetesClusterSpec{
			SecretRef: &zephyr_core_types.ResourceRef{
				Name:      secret.GetName(),
				Namespace: secret.GetNamespace(),
			},
			WriteNamespace: registerOpts.RemoteWriteNamespace,
		},
	}
	err := masterKubeClients.KubeClusterClient.UpsertKubernetesClusterSpec(ctx, cluster)
	if err != nil {
		return FailedToWriteKubeCluster(err)
	}
	return nil
}

func generateServiceAccountConfig(
	ctx context.Context,
	out io.Writer,
	kubeClients *common.KubeClients,
	remoteAuthConfig *rest.Config,
	registerOpts options.Register,
) (*rest.Config, error) {

	// the new cluster name doubles as the name for the service account we will auth as
	serviceAccountRef := &zephyr_core_types.ResourceRef{
		Name:      registerOpts.RemoteClusterName,
		Namespace: registerOpts.RemoteWriteNamespace,
	}
	configForServiceAccount, err := kubeClients.ClusterAuthorization.
		CreateAuthConfigForCluster(ctx, remoteAuthConfig, serviceAccountRef)
	if err != nil {
		fmt.Fprintf(out, FailedToCreateAuthToken(
			serviceAccountRef,
			registerOpts.RemoteKubeConfig,
			registerOpts.RemoteContext,
		))
		return nil, err
	}

	return configForServiceAccount, nil
}

// if:
//   * we are operating against a context named "kind-", AND
//   * the server appears to point to localhost, AND
//   * the --local-cluster-domain-override flag is populated with a value
//
// then we rewrite the server config to communicate over the value of `--local-cluster-domain-override`, which
// resolves to the host machine of docker. We also need to skip TLS verification
// and zero-out the cert data, because the cert on the remote cluster's API server wasn't
// issued for the domain contained in the value of `--local-cluster-domain-override`.
//
// this function call is a no-op if those conditions are not met
func hackClusterConfigForLocalTestingInKIND(
	remoteCluster *api.Cluster,
	remoteContextName, clusterDomainOverride string,
) error {
	serverUrl, err := url.Parse(remoteCluster.Server)
	if err != nil {
		return err
	}

	if strings.HasPrefix(remoteContextName, "kind-") &&
		(serverUrl.Hostname() == "127.0.0.1" || serverUrl.Hostname() == "localhost") &&
		clusterDomainOverride != "" {

		remoteCluster.Server = fmt.Sprintf("https://%s:%s", clusterDomainOverride, serverUrl.Port())
		remoteCluster.InsecureSkipTLSVerify = true
		remoteCluster.CertificateAuthority = ""
		remoteCluster.CertificateAuthorityData = []byte("")
	}

	return nil
}

func upsertSecretData(ctx context.Context, secret *k8s_core_types.Secret, secretClient k8s_core.SecretClient) error {
	existing, err := secretClient.GetSecret(ctx, client.ObjectKey{Name: secret.Name, Namespace: secret.Namespace})
	if err != nil {
		if k8s_errs.IsNotFound(err) {
			return secretClient.CreateSecret(ctx, secret)
		}
		return err
	}
	existing.Data = secret.Data
	existing.StringData = secret.StringData
	return secretClient.UpdateSecret(ctx, existing)
}
