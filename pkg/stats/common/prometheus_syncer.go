package common

import (
	"context"
	"fmt"

	"github.com/prometheus/prometheus/config"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/go-utils/stringutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/supergloo/pkg/api/custom/clients/prometheus"
	prometheusv1 "github.com/solo-io/supergloo/pkg/api/external/prometheus/v1"
	v1 "github.com/solo-io/supergloo/pkg/api/v1"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// registration-level syncer

type prometheusSyncer struct {
	syncerName                 string
	client                     prometheusv1.PrometheusConfigClient
	kube                       kubernetes.Interface
	chooseMesh                 func(mesh *v1.Mesh) bool
	scrapeConfigs              func(mesh *v1.Mesh) ([]*config.ScrapeConfig, error)
	skipBouncingPrometheusPods bool
}

func NewPrometheusSyncer(syncerName string, client prometheusv1.PrometheusConfigClient, kube kubernetes.Interface, chooseMesh func(mesh *v1.Mesh) bool, scrapeConfigs func(mesh *v1.Mesh) ([]*config.ScrapeConfig, error), skipBouncingPrometheusPods bool) *prometheusSyncer {
	return &prometheusSyncer{syncerName: syncerName, client: client, kube: kube, chooseMesh: chooseMesh, scrapeConfigs: scrapeConfigs, skipBouncingPrometheusPods: skipBouncingPrometheusPods}
}

// Ensure all prometheus configs contain scrape configs for the meshes which target them
//
// TODO (ilackarms): figure out a way to have multiple syncers report on the same resource
// currently prometheusSyncer returns errors instead of reporting
func (s *prometheusSyncer) Sync(ctx context.Context, snap *v1.RegistrationSnapshot) error {
	ctx = contextutils.WithLogger(ctx, fmt.Sprintf("%v-prometheus-syncer-%v", s.syncerName, snap.Hash()))
	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("begin sync %v", snap.Stringer())
	defer logger.Infof("end sync %v", snap.Stringer())

	// a map of the prometheus configmap to the meshes it should scrape
	promConfigsWithMeshes := s.getPromCfgsWithMeshes(snap.Meshes)

	logger.Infof("syncing %v prometheus configs", len(promConfigsWithMeshes))

	if err := s.syncPrometheusConfigsWithMeshes(ctx, promConfigsWithMeshes); err != nil {
		return errors.Wrapf(err, "failed syncing prometheus configs with mesh scrape jobs")
	}

	return nil
}

// get all the prometheus configs that must scrape one or more meshes
func (s *prometheusSyncer) getPromCfgsWithMeshes(meshes v1.MeshList) map[core.ResourceRef][]*v1.Mesh {

	// a map of the prometheus configmap to the meshes it should scrape
	promConfigsWithMeshes := make(map[core.ResourceRef][]*v1.Mesh)

	for _, mesh := range meshes {
		if !s.chooseMesh(mesh) {
			continue
		}

		// we only care about meshes with monitoring configs
		monitoring := mesh.GetMonitoringConfig()
		if monitoring == nil {
			continue
		}
		promCfgs := monitoring.GetPrometheusConfigmaps()

		for _, promCfg := range promCfgs {
			// add this mesh to the set this prometheus should scrape
			promConfigsWithMeshes[promCfg] = append(promConfigsWithMeshes[promCfg], mesh)
		}
	}

	return promConfigsWithMeshes
}

const superGlooScrapePrefix = "supergloo"

func scrapePrefix(mesh core.ResourceRef, syncerName string) string {
	return fmt.Sprintf("%v-%v-%v-", superGlooScrapePrefix, syncerName, mesh.Key())
}

func (s *prometheusSyncer) syncPrometheusConfigsWithMeshes(ctx context.Context, configsWithMeshes map[core.ResourceRef][]*v1.Mesh) error {
	// list all configs
	allPromConfigs, err := s.client.List("", clients.ListOpts{Ctx: ctx})
	if err != nil {
		return err
	}

	contextutils.LoggerFrom(ctx).Debugf("syncing prometheus configs with"+
		" mesh registrations: %#v", configsWithMeshes)

	// sync each prom config with the jobs it needs
	// we'll want to know which prom cfgs were updated so we can bounce their pods
	var updatedPromConfigs []core.ResourceRef
	for _, originalCfg := range allPromConfigs {
		cfgRef := originalCfg.Metadata.Ref()

		// get the scrape jobs we want to ensure for this config
		meshesForThisConfig := configsWithMeshes[cfgRef]

		// generate scrape configs
		var scrapeConfigsToAdd []*config.ScrapeConfig
		for _, mesh := range meshesForThisConfig {
			scrapesForMesh, err := s.scrapeConfigs(mesh)
			if err != nil {
				return errors.Wrapf(err, "internal error: generating scrape configs for mesh %v", mesh.Metadata.Ref())
			}
			// use mesh key as job prefix
			scrapesForMesh = prometheus.AddPrefix(scrapesForMesh, scrapePrefix(mesh.Metadata.Ref(), s.syncerName))
			scrapeConfigsToAdd = append(scrapesForMesh, scrapeConfigsToAdd...)
		}

		promCfg := originalCfg.Clone().(*prometheusv1.PrometheusConfig)

		// remove all scrape configs and start fresh
		removed := promCfg.RemoveScrapeConfigs(fmt.Sprintf("%v-%v", superGlooScrapePrefix, s.syncerName))

		// add all scrape configs
		added := promCfg.AddScrapeConfigs(scrapeConfigsToAdd)

		// this is not a config we manage
		if removed == 0 && added == 0 {
			continue
		}

		if hashutils.HashAll(promCfg.ScrapeConfigs) == hashutils.HashAll(originalCfg.PrometheusConfig) {
			continue
		}

		contextutils.LoggerFrom(ctx).Infof("prometheus %v syncing %v prometheus scrape configs", cfgRef.Key(), added)

		if _, err := s.client.Write(promCfg, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true}); err != nil {
			return errors.Wrapf(err, "writing updated prometheus config to storage")
		}

		updatedPromConfigs = append(updatedPromConfigs, promCfg.Metadata.Ref())
	}

	if s.skipBouncingPrometheusPods {
		return nil
	}

	return s.bouncePodsWithConfigs(ctx, updatedPromConfigs)
}

// bounce pods will delete pods which use any of the updated configs specified
// this is required to kick prometheus and ensure it receives the latest config

func (s *prometheusSyncer) bouncePodsWithConfigs(ctx context.Context, updatedConfigs []core.ResourceRef) error {
	contextutils.LoggerFrom(ctx).Infof("bouncing prometheus pods with updated configmaps")

	// collect pods across namespaces in which we've updated configmaps
	var namespacesWithUpdatedPrometheusConfigs []string
	for _, cfg := range updatedConfigs {
		namespacesWithUpdatedPrometheusConfigs = append(namespacesWithUpdatedPrometheusConfigs, cfg.Namespace)
	}
	namespacesWithUpdatedPrometheusConfigs = stringutils.Unique(namespacesWithUpdatedPrometheusConfigs)

	// list all pods, bounce any that mount any of the updated configmaps
	var allPods []kubev1.Pod
	for _, ns := range namespacesWithUpdatedPrometheusConfigs {
		pods, err := s.kube.CoreV1().Pods(ns).List(metav1.ListOptions{})
		if err != nil {
			return errors.Wrapf(err, "listing pods in ns %v", ns)
		}
		allPods = append(allPods, pods.Items...)
	}

	// for each pod, if it contains a volume mount with one of the updated configmaps, bounce it
	var podsToBounce []kubev1.Pod
findPodsToBounce:
	for _, pod := range allPods {
		for _, vol := range pod.Spec.Volumes {
			if vol.ConfigMap == nil {
				continue
			}
			// see if the pod namespace + the configmap name match one of our refs
			for _, updatedConfigmap := range updatedConfigs {
				if updatedConfigmap.Name == vol.ConfigMap.Name && updatedConfigmap.Namespace == pod.Namespace {
					podsToBounce = append(podsToBounce, pod)
					continue findPodsToBounce
				}
			}
		}
	}

	for _, pod := range podsToBounce {
		contextutils.LoggerFrom(ctx).Infof("bouncing prometheus pod %v.%v", pod.Namespace, pod.Name)
		if err := s.kube.CoreV1().Pods(pod.Namespace).Delete(pod.Name, nil); err != nil {
			return errors.Wrapf(err, "bouncing prometheus pod with updated prometheus configmap")
		}
	}

	return nil
}
