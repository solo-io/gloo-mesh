package e2e

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/aokoli/goutils"
	"github.com/golang/sync/errgroup"
	zephyr_core "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/kube/metadata"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	AwsAccountId          = "410461945957"
	Region                = "us-east-2"
	AppmeshArn            = "arn:aws:appmesh:us-east-2:410461945957:mesh/smh-e2e-test"
	EksArn                = "arn:aws:eks:us-east-2:410461945957:cluster/smh-e2e-test"
	AppmeshName           = "smh-e2e-test"
	EksClusterName        = "smh-e2e-test"
	SmhNamespace          = "service-mesh-hub"
	AppmeshInjectionLabel = "appmesh.k8s.aws/sidecarInjectorWebhook=enabled"
)

var (
	settingsObjKey    = client.ObjectKey{Name: metadata.GlobalSettingsName, Namespace: SmhNamespace}
	secretObjKey      = client.ObjectKey{Name: "e2e-appmesh-eks", Namespace: SmhNamespace}
	kubeClusterObjKey = client.ObjectKey{Name: metadata.BuildEksKubernetesClusterName(EksClusterName, Region), Namespace: SmhNamespace}
	appmeshObjKey     = client.ObjectKey{Name: metadata.BuildAppMeshName(AppmeshName, Region, AwsAccountId), Namespace: SmhNamespace}
	virtualMeshObjKey = client.ObjectKey{Name: "appmesh-vm", Namespace: SmhNamespace}

	// Populated during setup
	generatedNamespace string

	defaultSettingsYaml = fmt.Sprintf(`
apiVersion: core.zephyr.solo.io/zephyr_networking
kind: Settings
metadata:
  name: %s
  namespace: %s
spec:
  aws:
    disabled: true
`, settingsObjKey.Name, settingsObjKey.Namespace)

	settingsYaml = fmt.Sprintf(`
apiVersion: core.zephyr.solo.io/zephyr_networking
kind: Settings
metadata:
  name: %s
  namespace: %s
spec:
  aws:
    disabled: false
    accounts:
      - accountId: "%s"
        meshDiscovery:
          resourceSelectors:
          - arn: %s
        eksDiscovery:
          resourceSelectors:
          - arn: %s
`, settingsObjKey.Name, settingsObjKey.Namespace, AwsAccountId, AppmeshArn, EksArn)
	virtualMeshYaml = fmt.Sprintf(`
apiVersion: networking.zephyr.solo.io/zephyr_networking
kind: VirtualMesh
metadata:
  name: %s
  namespace: %s
spec:
  displayName: "Appmesh-VM"
  certificateAuthority:
    builtin:
      ttlDays: 356
      rsaKeySizeBytes: 4096
      orgName: "service-mesh-hub"
  federation: 
    mode: PERMISSIVE
  shared: {}
  enforceAccessControl: ENABLED
  meshes:
  - name: %s
    namespace: %s
`, virtualMeshObjKey.Name, virtualMeshObjKey.Namespace, appmeshObjKey.Name, appmeshObjKey.Namespace)

	buildSecretYaml = func(awsAccessKeyId, awsSecretAccessKey string) string {
		return fmt.Sprintf(`
apiVersion: v1
kind: Secret
type: solo.io/register/aws-credentials
metadata:
  name: %s
  namespace: %s
data:
  aws_access_key_id: %s
  aws_secret_access_key: %s
`, secretObjKey.Name, secretObjKey.Namespace, awsAccessKeyId, awsSecretAccessKey)
	}

	kubeClusterName = metadata.BuildEksKubernetesClusterName(EksClusterName, Region)
)

func getEksKubeContext(ctx context.Context) KubeContext {
	eg, ctx := errgroup.WithContext(ctx)

	r, w, err := os.Pipe()
	Expect(err).NotTo(HaveOccurred())
	defer r.Close()

	cmd := exec.CommandContext(ctx, "aws", "eks", "--region", Region, "update-kubeconfig", "--name", EksClusterName)
	cmd.Dir = "../.."
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	cmd.ExtraFiles = append(cmd.ExtraFiles, w)
	err = cmd.Start()
	// close this end after start, as we dont need it.
	w.Close()
	Expect(err).NotTo(HaveOccurred())

	eg.Go(cmd.Wait)

	err = eg.Wait()
	if err != nil {
		dumpState()
	}

	Expect(err).NotTo(HaveOccurred())

	// Use current context
	return NewKubeContext("")
}

func setupAppmeshEksEnvironment() string {
	// Deploy bookinfo into a new namespace on the EKS cluster
	ctx := context.Background()
	eksContext := getEksKubeContext(ctx)
	randomString, err := goutils.RandomAlphabetic(4)
	Expect(err).ToNot(HaveOccurred())
	generatedNamespace = strings.ToLower(randomString)
	eksContext.CreateNamespace(ctx, generatedNamespace)
	eksContext.LabelNamespace(ctx, generatedNamespace, AppmeshInjectionLabel)
	eksContext.DeployBookInfo(ctx, generatedNamespace)
	// Set SERVICES_DOMAIN env var on productpage-v1 to point at services within generated namespace
	// https://github.com/istio/istio/blob/1.5.0/samples/bookinfo/src/productpage/productpage.py#L60
	eksContext.SetDeploymentEnvVars(ctx, generatedNamespace, "productpage-v1", "productpage", map[string]string{"SERVICES_DOMAIN": generatedNamespace})
	eksContext.SetDeploymentEnvVars(ctx, generatedNamespace, "reviews-v1", "reviews", map[string]string{"SERVICES_DOMAIN": generatedNamespace})
	eksContext.SetDeploymentEnvVars(ctx, generatedNamespace, "reviews-v2", "reviews", map[string]string{"SERVICES_DOMAIN": generatedNamespace})
	eksContext.SetDeploymentEnvVars(ctx, generatedNamespace, "reviews-v3", "reviews", map[string]string{"SERVICES_DOMAIN": generatedNamespace})
	return generatedNamespace
}

func cleanupAppmeshEksEnvironment(ns string) {
	if useExisting := os.Getenv("USE_EXISTING"); useExisting != "" {
		// Cleans up discovery resources on management cluster
		// Reset back to default settings. This must be done before removing the AWS secret.
		settings, err := env.Management.SettingsClient.GetSettings(context.Background(), settingsObjKey)
		Expect(err).NotTo(HaveOccurred())
		var defaultSettings zephyr_core.Settings
		ParseYaml(defaultSettingsYaml, &defaultSettings)
		settings.Spec = defaultSettings.Spec
		err = env.Management.SettingsClient.UpdateSettings(context.Background(), settings)
		Expect(err).NotTo(HaveOccurred())
		// Wait for mesh-discovery to clean up discovered resources
		time.Sleep(10 * time.Second)
		// Delete AWS credentials
		err = env.Management.SecretClient.DeleteSecret(context.Background(), secretObjKey)
		if errors.IsNotFound(err) {
			err = nil
		}
		Expect(err).NotTo(HaveOccurred())
		// Delete remote cluster k8s credentials
		err = env.Management.SecretClient.DeleteSecret(context.Background(), kubeClusterObjKey)
		if errors.IsNotFound(err) {
			err = nil
		}
		Expect(err).NotTo(HaveOccurred())
		// Delete KubernetesCluster
		err = env.Management.KubeClusterClient.DeleteKubernetesCluster(context.Background(), kubeClusterObjKey)
		if errors.IsNotFound(err) {
			err = nil
		}
		Expect(err).NotTo(HaveOccurred())
	}

	// Clean up resources on remote EKS cluster
	ctx := context.Background()
	eksContext := getEksKubeContext(ctx)
	eksContext.DeleteNamespace(ctx, ns)
}

var _ = Describe("Appmesh EKS ", func() {
	AfterEach(func() {
		testLabels := map[string]string{"test": "true"}
		opts := &client.DeleteAllOfOptions{}
		opts.LabelSelector = labels.SelectorFromSet(testLabels)
		opts.Namespace = "service-mesh-hub"
	})

	// Fetch base64 encoded AWS credentials from environment
	var registerAwsSecret = func() {
		awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
		awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		awsAccessKeyIdBase64 := base64.StdEncoding.EncodeToString([]byte(awsAccessKeyId))
		awsSecretAccessKeyBase64 := base64.StdEncoding.EncodeToString([]byte(awsSecretAccessKey))
		var secret v1.Secret
		secretYaml := buildSecretYaml(awsAccessKeyIdBase64, awsSecretAccessKeyBase64)
		ParseYaml(secretYaml, &secret)
		err := env.Management.SecretClient.CreateSecret(context.Background(), &secret)
		Expect(err).NotTo(HaveOccurred())
	}

	var applySettings = func() {
		var newSettings zephyr_core.Settings
		ParseYaml(settingsYaml, &newSettings)
		existingSettings, err := env.Management.SettingsClient.GetSettings(context.Background(), settingsObjKey)
		Expect(err).NotTo(HaveOccurred())
		if !existingSettings.Spec.Equal(newSettings.Spec) {
			existingSettings.Spec = newSettings.Spec
			err = env.Management.SettingsClient.UpdateSettings(context.Background(), existingSettings)
			Expect(err).NotTo(HaveOccurred())
		}
	}

	var applyVirtualMesh = func() {
		var virtualMesh zephyr_networking.VirtualMesh
		ParseYaml(virtualMeshYaml, &virtualMesh)
		err := env.Management.VirtualMeshClient.CreateVirtualMesh(context.Background(), &virtualMesh)
		Expect(err).NotTo(HaveOccurred())
	}

	var deleteVirtualMesh = func() {
		err := env.Management.VirtualMeshClient.DeleteVirtualMesh(context.Background(), virtualMeshObjKey)
		// Give SMH time to make API calls to delete Appmesh resources
		time.Sleep(5 * time.Second)
		Expect(err).NotTo(HaveOccurred())
	}

	var expectGetKubeCluster = func(name string) {
		Eventually(
			KubeClusterShouldExist(client.ObjectKey{Name: name, Namespace: SmhNamespace}, env.Management),
			"30s", "1s").
			ShouldNot(BeNil())
	}

	var curlReviewsWithExpectedOutput = func(expectedString string, shouldExpect bool) {
		productPageDeployment := "productpage-v1"
		ctx := context.Background()
		eksContext := getEksKubeContext(ctx)
		eksContext.WaitForRollout(ctx, generatedNamespace, productPageDeployment)
		if shouldExpect {
			Eventually(func() string {
				return eksContext.Curl(
					context.Background(),
					generatedNamespace,
					productPageDeployment,
					"curl",
					fmt.Sprintf("http://reviews.%s:9080/reviews/1", generatedNamespace))
			}, "120s", "1s").Should(ContainSubstring(expectedString))
		} else {
			Eventually(func() string {
				return eksContext.Curl(
					context.Background(),
					generatedNamespace,
					productPageDeployment,
					"curl",
					fmt.Sprintf("http://reviews.%s:9080/reviews/1", generatedNamespace))
			}, "120s", "1s").ShouldNot(ContainSubstring(expectedString))
		}
	}

	It("should discover Appmesh mesh and EKS cluster", func() {
		// Register AWS account credentials
		registerAwsSecret()
		// Discover Appmesh mesh and EKS cluster
		applySettings()
		expectGetKubeCluster(kubeClusterName)
	})

	It("should translate Appmesh resources to enable all communication between workloads and services", func() {
		applyVirtualMesh()
		time.Sleep(200 * time.Second)
		curlReviewsWithExpectedOutput("The slapstick humour is refreshing", true)
	})

	It("should cleanup translated Appmesh resources and disable communication between workloads and services", func() {
		deleteVirtualMesh()
		curlReviewsWithExpectedOutput("The slapstick humour is refreshing", false)
	})
})
