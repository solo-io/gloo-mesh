package istio

import (
	"context"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	kubev1 "k8s.io/api/core/v1"
	batchv1client "k8s.io/client-go/kubernetes/typed/batch/v1"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/supergloo/pkg/meshdiscovery/clientset"

	"github.com/solo-io/supergloo/pkg/translator/utils"
	"go.uber.org/zap"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

// the subsets of the kube api that we need
type CrdGetter interface {
	Get(name string, options metav1.GetOptions) (*v1beta1.CustomResourceDefinition, error)
}

type istioDiscoverySyncer struct {
	writeNamespace         string
	meshReconciler         v1.MeshReconciler
	meshPolicyClientLoader clientset.MeshPolicyClientLoader
	crdGetter              CrdGetter
	jobGetter              batchv1client.JobsGetter
}

func NewIstioDiscoverySyncer(writeNamespace string, meshReconciler v1.MeshReconciler, meshPolicyClient clientset.MeshPolicyClientLoader, crdGetter CrdGetter, jobGetter batchv1client.JobsGetter) v1.DiscoverySyncer {
	return &istioDiscoverySyncer{writeNamespace: writeNamespace, meshReconciler: meshReconciler, meshPolicyClientLoader: meshPolicyClient, crdGetter: crdGetter, jobGetter: jobGetter}
}

var discoveryLabels = map[string]string{
	"discovered_by": "istio-mesh-discovery",
}

func (i *istioDiscoverySyncer) ShouldSync(old, new *v1.DiscoverySnapshot) bool {
	if old == nil {
		return true
	}
	desired1, err1 := i.desiredMeshes(context.TODO(), old)
	desired2, err2 := i.desiredMeshes(context.TODO(), new)
	if err1 != nil || err2 != nil {
		return true
	}
	return hashutils.HashAll(desired1) != hashutils.HashAll(desired2)
}

func (i *istioDiscoverySyncer) Sync(ctx context.Context, snap *v1.DiscoverySnapshot) error {
	ctx = contextutils.WithLogger(ctx, "istio-mesh-discovery")
	logger := contextutils.LoggerFrom(ctx)
	logger.Infow("begin sync",
		zap.Int("Upstreams", len(snap.Upstreams)),
		zap.Int("Deployments", len(snap.Deployments)),
		zap.Int("Tlssecrets", len(snap.Tlssecrets)),
		zap.Int("Configmaps", len(snap.Configmaps)),
		zap.Int("Pods", len(snap.Pods)),
	)
	defer logger.Infow("end sync")
	logger.Debugf("full snapshot: %v", snap)

	istioMeshes, err := i.desiredMeshes(ctx, snap)
	if err != nil {
		return err
	}

	if err := i.meshReconciler.Reconcile(i.writeNamespace, istioMeshes, reconcileMeshDiscoveryMetadata, clients.ListOpts{Ctx: ctx, Selector: discoveryLabels}); err != nil {
		return err
	}

	return nil
}

type pilotDeployment struct {
	version, namespace string
}

func (i *istioDiscoverySyncer) desiredMeshes(ctx context.Context, snap *v1.DiscoverySnapshot) (v1.MeshList, error) {
	meshPolicyCrdRegistered, err := detectMeshPolicyCrd(i.crdGetter)
	if err != nil {
		return nil, newErrorDetectingMeshPolicy(err)
	}
	if !meshPolicyCrdRegistered {
		return nil, nil
	}

	pilots := detectPilotDeployments(ctx, snap.Deployments)

	if len(pilots) == 0 {
		return nil, nil
	}

	globalMtlsEnabled := func() bool {
		meshPolicyClient, err := i.meshPolicyClientLoader()
		if err != nil {
			return false
		}

		// https://istio.io/docs/tasks/security/authn-policy/#globally-enabling-istio-mutual-tls
		defaultMeshPolicy, err := meshPolicyClient.Read("default", clients.ReadOpts{Ctx: ctx})
		if err != nil {
			return false
		}
		for _, peer := range defaultMeshPolicy.GetPeers() {
			return peer.GetMtls() != nil
		}
		return false
	}()

	injectedPods := detectInjectedIstioPods(ctx, snap.Pods)

	var istioMeshes v1.MeshList
	for _, pilot := range pilots {

		// ensure that the post-install jobs have run for this pilot,
		// if not, we're not ready to detect it
		postInstallComplete, err := detectPostInstallJobComplete(i.jobGetter, pilot.namespace)
		if err != nil {
			contextutils.LoggerFrom(ctx).Errorf("failed to detect if post-install jobs have finished: %v", err)
			continue
		}

		if !postInstallComplete {
			continue
		}

		var autoInjectionEnabled bool
		sidecarInjector, err := snap.Deployments.Find(pilot.namespace, "istio-sidecar-injector")
		if err == nil && (sidecarInjector.Spec.Replicas == nil || *sidecarInjector.Spec.Replicas > 0) {
			autoInjectionEnabled = true
		}

		var smiEnabled bool
		smiAdapter, err := snap.Deployments.Find(pilot.namespace, "smi-adapter-istio")
		if err == nil && (smiAdapter.Spec.Replicas == nil || *smiAdapter.Spec.Replicas > 0) {
			smiEnabled = true
		}

		// https://istio.io/docs/tasks/security/plugin-ca-cert/#plugging-in-the-existing-certificate-and-key
		var rootCa *core.ResourceRef
		customRootCa, err := snap.Tlssecrets.Find(pilot.namespace, "cacerts")
		if err == nil {
			root := customRootCa.Metadata.Ref()
			rootCa = &root
		}

		var mtlsConfig *v1.MtlsConfig
		if globalMtlsEnabled {
			mtlsConfig = &v1.MtlsConfig{
				MtlsEnabled:     true,
				RootCertificate: rootCa,
			}
		}

		meshUpstreams := func() []*core.ResourceRef {
			injectedUpstreams := utils.UpstreamsForPods(injectedPods[pilot.namespace], snap.Upstreams)
			var usRefs []*core.ResourceRef
			for _, us := range injectedUpstreams {
				ref := us.Metadata.Ref()
				usRefs = append(usRefs, &ref)
			}
			return usRefs
		}()

		istioMesh := &v1.Mesh{
			Metadata: core.Metadata{
				Name:      "istio-" + pilot.namespace,
				Namespace: i.writeNamespace,
				Labels:    discoveryLabels,
			},
			MeshType: &v1.Mesh_Istio{
				Istio: &v1.IstioMesh{
					InstallationNamespace: pilot.namespace,
					Version:               pilot.version,
				},
			},
			MtlsConfig: mtlsConfig,
			SmiEnabled: smiEnabled,
			DiscoveryMetadata: &v1.DiscoveryMetadata{
				EnableAutoInject: autoInjectionEnabled,
				MtlsConfig:       mtlsConfig,
				Upstreams:        meshUpstreams,
			},
		}
		istioMeshes = append(istioMeshes, istioMesh)
	}

	return istioMeshes, nil
}

func detectPilotDeployments(ctx context.Context, deployments kubernetes.DeploymentList) []pilotDeployment {
	var pilots []pilotDeployment
	for _, deployment := range deployments {
		for _, container := range deployment.Spec.Template.Spec.Containers {
			if strings.Contains(container.Image, "istio/pilot") {
				split := strings.Split(container.Image, ":")
				if len(split) != 2 {
					contextutils.LoggerFrom(ctx).Errorf("invalid or unexpected image format for pilot: %v", container.Image)
					continue
				}
				pilots = append(pilots, pilotDeployment{version: split[1], namespace: deployment.Namespace})
			}
		}
	}
	return pilots
}

func detectMeshPolicyCrd(crdGetter CrdGetter) (bool, error) {
	_, err := crdGetter.Get(v1alpha1.MeshPolicyCrd.FullName(), metav1.GetOptions{})
	if err == nil {
		return true, nil
	}
	if kubeerrs.IsNotFound(err) {
		return false, nil
	}
	return false, err
}

var PostInstallJobs = []string{
	"istio-security-post-install",
}

func detectPostInstallJobComplete(jobGetter batchv1client.JobsGetter, pilotNamespace string) (bool, error) {
	for _, jobName := range PostInstallJobs {
		job, err := jobGetter.Jobs(pilotNamespace).Get(jobName, metav1.GetOptions{})
		if err != nil {
			if kubeerrs.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}
		var jobComplete bool
		for _, condition := range job.Status.Conditions {
			if condition.Type == batchv1.JobComplete && condition.Status == kubev1.ConditionTrue {
				jobComplete = true
				break
			}
		}
		if !jobComplete {
			return false, nil
		}
	}
	return true, nil
}

func detectInjectedIstioPods(ctx context.Context, pods kubernetes.PodList) map[string]kubernetes.PodList {
	injectedPods := make(map[string]kubernetes.PodList)
	for _, pod := range pods {
		discoveryNamespace, ok := detectInjectedIstioPod(ctx, pod)
		if ok {
			injectedPods[discoveryNamespace] = append(injectedPods[discoveryNamespace], pod)
		}
	}
	return injectedPods
}

func detectInjectedIstioPod(ctx context.Context, pod *kubernetes.Pod) (string, bool) {
	for _, container := range pod.Spec.Containers {
		if container.Name == "istio-proxy" {
			for i, arg := range container.Args {
				if arg == "--discoveryAddress" {
					if i == len(container.Args) {
						contextutils.LoggerFrom(ctx).Errorf("invalid args for istio-proxy sidecar for pod %v.%v: "+
							"expected to find --discoveryAddress with a parameter. instead, found %v", pod.Namespace, pod.Name,
							container.Args)
						return "", false
					}
					discoveryAddress := container.Args[i+1]
					discoveryService := strings.Split(discoveryAddress, ":")[0]
					discoveryNamespace := strings.TrimPrefix(discoveryService, "istio-pilot.")

					return discoveryNamespace, true
				}
			}
		}
	}
	return "", false
}

// after the first write, i.e. on any update
// we are only interested in updating discovery metadata at this point
// do not want to overwrite user-modified config settings
// smi is also included here as it's not intended to be user-writeable
// a return value of false indicates that, for the purposes of this syncer,
// the resource in storage matches the desired resource
func reconcileMeshDiscoveryMetadata(original, desired *v1.Mesh) (b bool, e error) {
	if desired.DiscoveryMetadata.Equal(original.DiscoveryMetadata) && desired.SmiEnabled == original.SmiEnabled {
		return false, nil
	}
	original.DiscoveryMetadata = desired.DiscoveryMetadata
	original.SmiEnabled = desired.SmiEnabled
	*desired = *original
	return true, nil
}
