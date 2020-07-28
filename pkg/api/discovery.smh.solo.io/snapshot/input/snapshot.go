// Code generated by skv2. DO NOT EDIT.

// The Input Snapshot contains the set of all:
// * ConfigMaps
// * Services
// * Pods
// * Nodes
// * Deployments
// * ReplicaSets
// * DaemonSets
// * StatefulSets
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the SnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1_client "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	v1 "k8s.io/api/core/v1"

	apps_v1_client "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	apps_v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
	apps_v1 "k8s.io/api/apps/v1"
)

// the snapshot of input resources consumed by translation
type Snapshot interface {

	// return the set of input ConfigMaps
	ConfigMaps() v1_sets.ConfigMapSet
	// return the set of input Services
	Services() v1_sets.ServiceSet
	// return the set of input Pods
	Pods() v1_sets.PodSet
	// return the set of input Nodes
	Nodes() v1_sets.NodeSet

	// return the set of input Deployments
	Deployments() apps_v1_sets.DeploymentSet
	// return the set of input ReplicaSets
	ReplicaSets() apps_v1_sets.ReplicaSetSet
	// return the set of input DaemonSets
	DaemonSets() apps_v1_sets.DaemonSetSet
	// return the set of input StatefulSets
	StatefulSets() apps_v1_sets.StatefulSetSet
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name string

	configMaps v1_sets.ConfigMapSet
	services   v1_sets.ServiceSet
	pods       v1_sets.PodSet
	nodes      v1_sets.NodeSet

	deployments  apps_v1_sets.DeploymentSet
	replicaSets  apps_v1_sets.ReplicaSetSet
	daemonSets   apps_v1_sets.DaemonSetSet
	statefulSets apps_v1_sets.StatefulSetSet
}

func NewSnapshot(
	name string,

	configMaps v1_sets.ConfigMapSet,
	services v1_sets.ServiceSet,
	pods v1_sets.PodSet,
	nodes v1_sets.NodeSet,

	deployments apps_v1_sets.DeploymentSet,
	replicaSets apps_v1_sets.ReplicaSetSet,
	daemonSets apps_v1_sets.DaemonSetSet,
	statefulSets apps_v1_sets.StatefulSetSet,

) Snapshot {
	return &snapshot{
		name: name,

		configMaps:   configMaps,
		services:     services,
		pods:         pods,
		nodes:        nodes,
		deployments:  deployments,
		replicaSets:  replicaSets,
		daemonSets:   daemonSets,
		statefulSets: statefulSets,
	}
}

func (s snapshot) ConfigMaps() v1_sets.ConfigMapSet {
	return s.configMaps
}

func (s snapshot) Services() v1_sets.ServiceSet {
	return s.services
}

func (s snapshot) Pods() v1_sets.PodSet {
	return s.pods
}

func (s snapshot) Nodes() v1_sets.NodeSet {
	return s.nodes
}

func (s snapshot) Deployments() apps_v1_sets.DeploymentSet {
	return s.deployments
}

func (s snapshot) ReplicaSets() apps_v1_sets.ReplicaSetSet {
	return s.replicaSets
}

func (s snapshot) DaemonSets() apps_v1_sets.DaemonSetSet {
	return s.daemonSets
}

func (s snapshot) StatefulSets() apps_v1_sets.StatefulSetSet {
	return s.statefulSets
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["configMaps"] = s.configMaps.List()
	snapshotMap["services"] = s.services.List()
	snapshotMap["pods"] = s.pods.List()
	snapshotMap["nodes"] = s.nodes.List()
	snapshotMap["deployments"] = s.deployments.List()
	snapshotMap["replicaSets"] = s.replicaSets.List()
	snapshotMap["daemonSets"] = s.daemonSets.List()
	snapshotMap["statefulSets"] = s.statefulSets.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
// Two types of builders are available:
// a builder for snapshots of resources across multiple clusters
// a builder for snapshots of resources within a single cluster
type Builder interface {
	BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error)
}

// Options for building a snapshot
type BuildOptions struct {

	// List options for composing a snapshot from ConfigMaps
	ConfigMaps []client.ListOption
	// List options for composing a snapshot from Services
	Services []client.ListOption
	// List options for composing a snapshot from Pods
	Pods []client.ListOption
	// List options for composing a snapshot from Nodes
	Nodes []client.ListOption

	// List options for composing a snapshot from Deployments
	Deployments []client.ListOption
	// List options for composing a snapshot from ReplicaSets
	ReplicaSets []client.ListOption
	// List options for composing a snapshot from DaemonSets
	DaemonSets []client.ListOption
	// List options for composing a snapshot from StatefulSets
	StatefulSets []client.ListOption
}

// build a snapshot from resources across multiple clusters
type multiClusterBuilder struct {
	clusters multicluster.ClusterSet

	configMaps v1_client.MulticlusterConfigMapClient
	services   v1_client.MulticlusterServiceClient
	pods       v1_client.MulticlusterPodClient
	nodes      v1_client.MulticlusterNodeClient

	deployments  apps_v1_client.MulticlusterDeploymentClient
	replicaSets  apps_v1_client.MulticlusterReplicaSetClient
	daemonSets   apps_v1_client.MulticlusterDaemonSetClient
	statefulSets apps_v1_client.MulticlusterStatefulSetClient
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterBuilder(
	clusters multicluster.ClusterSet,
	client multicluster.Client,
) Builder {
	return &multiClusterBuilder{
		clusters: clusters,

		configMaps: v1_client.NewMulticlusterConfigMapClient(client),
		services:   v1_client.NewMulticlusterServiceClient(client),
		pods:       v1_client.NewMulticlusterPodClient(client),
		nodes:      v1_client.NewMulticlusterNodeClient(client),

		deployments:  apps_v1_client.NewMulticlusterDeploymentClient(client),
		replicaSets:  apps_v1_client.NewMulticlusterReplicaSetClient(client),
		daemonSets:   apps_v1_client.NewMulticlusterDaemonSetClient(client),
		statefulSets: apps_v1_client.NewMulticlusterStatefulSetClient(client),
	}
}

func (b *multiClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	configMaps := v1_sets.NewConfigMapSet()
	services := v1_sets.NewServiceSet()
	pods := v1_sets.NewPodSet()
	nodes := v1_sets.NewNodeSet()

	deployments := apps_v1_sets.NewDeploymentSet()
	replicaSets := apps_v1_sets.NewReplicaSetSet()
	daemonSets := apps_v1_sets.NewDaemonSetSet()
	statefulSets := apps_v1_sets.NewStatefulSetSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertConfigMapsFromCluster(ctx, cluster, configMaps, opts.ConfigMaps...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertServicesFromCluster(ctx, cluster, services, opts.Services...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertPodsFromCluster(ctx, cluster, pods, opts.Pods...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertNodesFromCluster(ctx, cluster, nodes, opts.Nodes...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDeploymentsFromCluster(ctx, cluster, deployments, opts.Deployments...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertReplicaSetsFromCluster(ctx, cluster, replicaSets, opts.ReplicaSets...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDaemonSetsFromCluster(ctx, cluster, daemonSets, opts.DaemonSets...); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertStatefulSetsFromCluster(ctx, cluster, statefulSets, opts.StatefulSets...); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewSnapshot(
		name,

		configMaps,
		services,
		pods,
		nodes,
		deployments,
		replicaSets,
		daemonSets,
		statefulSets,
	)

	return outputSnap, errs
}

func (b *multiClusterBuilder) insertConfigMapsFromCluster(ctx context.Context, cluster string, configMaps v1_sets.ConfigMapSet, opts ...client.ListOption) error {
	configMapClient, err := b.configMaps.Cluster(cluster)
	if err != nil {
		return err
	}

	configMapList, err := configMapClient.ListConfigMap(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range configMapList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		configMaps.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertServicesFromCluster(ctx context.Context, cluster string, services v1_sets.ServiceSet, opts ...client.ListOption) error {
	serviceClient, err := b.services.Cluster(cluster)
	if err != nil {
		return err
	}

	serviceList, err := serviceClient.ListService(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range serviceList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		services.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertPodsFromCluster(ctx context.Context, cluster string, pods v1_sets.PodSet, opts ...client.ListOption) error {
	podClient, err := b.pods.Cluster(cluster)
	if err != nil {
		return err
	}

	podList, err := podClient.ListPod(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range podList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		pods.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertNodesFromCluster(ctx context.Context, cluster string, nodes v1_sets.NodeSet, opts ...client.ListOption) error {
	nodeClient, err := b.nodes.Cluster(cluster)
	if err != nil {
		return err
	}

	nodeList, err := nodeClient.ListNode(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range nodeList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		nodes.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertDeploymentsFromCluster(ctx context.Context, cluster string, deployments apps_v1_sets.DeploymentSet, opts ...client.ListOption) error {
	deploymentClient, err := b.deployments.Cluster(cluster)
	if err != nil {
		return err
	}

	deploymentList, err := deploymentClient.ListDeployment(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range deploymentList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		deployments.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertReplicaSetsFromCluster(ctx context.Context, cluster string, replicaSets apps_v1_sets.ReplicaSetSet, opts ...client.ListOption) error {
	replicaSetClient, err := b.replicaSets.Cluster(cluster)
	if err != nil {
		return err
	}

	replicaSetList, err := replicaSetClient.ListReplicaSet(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range replicaSetList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		replicaSets.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertDaemonSetsFromCluster(ctx context.Context, cluster string, daemonSets apps_v1_sets.DaemonSetSet, opts ...client.ListOption) error {
	daemonSetClient, err := b.daemonSets.Cluster(cluster)
	if err != nil {
		return err
	}

	daemonSetList, err := daemonSetClient.ListDaemonSet(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range daemonSetList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		daemonSets.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertStatefulSetsFromCluster(ctx context.Context, cluster string, statefulSets apps_v1_sets.StatefulSetSet, opts ...client.ListOption) error {
	statefulSetClient, err := b.statefulSets.Cluster(cluster)
	if err != nil {
		return err
	}

	statefulSetList, err := statefulSetClient.ListStatefulSet(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range statefulSetList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		statefulSets.Insert(&item)
	}

	return nil
}

// build a snapshot from resources in a single cluster
type singleClusterBuilder struct {
	configMaps v1_client.ConfigMapClient
	services   v1_client.ServiceClient
	pods       v1_client.PodClient
	nodes      v1_client.NodeClient

	deployments  apps_v1_client.DeploymentClient
	replicaSets  apps_v1_client.ReplicaSetClient
	daemonSets   apps_v1_client.DaemonSetClient
	statefulSets apps_v1_client.StatefulSetClient
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewSingleClusterBuilder(
	client client.Client,
) Builder {
	return &singleClusterBuilder{

		configMaps: v1_client.NewConfigMapClient(client),
		services:   v1_client.NewServiceClient(client),
		pods:       v1_client.NewPodClient(client),
		nodes:      v1_client.NewNodeClient(client),

		deployments:  apps_v1_client.NewDeploymentClient(client),
		replicaSets:  apps_v1_client.NewReplicaSetClient(client),
		daemonSets:   apps_v1_client.NewDaemonSetClient(client),
		statefulSets: apps_v1_client.NewStatefulSetClient(client),
	}
}

func (b *singleClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	configMaps := v1_sets.NewConfigMapSet()
	services := v1_sets.NewServiceSet()
	pods := v1_sets.NewPodSet()
	nodes := v1_sets.NewNodeSet()

	deployments := apps_v1_sets.NewDeploymentSet()
	replicaSets := apps_v1_sets.NewReplicaSetSet()
	daemonSets := apps_v1_sets.NewDaemonSetSet()
	statefulSets := apps_v1_sets.NewStatefulSetSet()

	var errs error

	if err := b.insertConfigMaps(ctx, configMaps, opts.ConfigMaps...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertServices(ctx, services, opts.Services...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertPods(ctx, pods, opts.Pods...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertNodes(ctx, nodes, opts.Nodes...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertDeployments(ctx, deployments, opts.Deployments...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertReplicaSets(ctx, replicaSets, opts.ReplicaSets...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertDaemonSets(ctx, daemonSets, opts.DaemonSets...); err != nil {
		errs = multierror.Append(errs, err)
	}
	if err := b.insertStatefulSets(ctx, statefulSets, opts.StatefulSets...); err != nil {
		errs = multierror.Append(errs, err)
	}

	outputSnap := NewSnapshot(
		name,

		configMaps,
		services,
		pods,
		nodes,
		deployments,
		replicaSets,
		daemonSets,
		statefulSets,
	)

	return outputSnap, errs
}

func (b *singleClusterBuilder) insertConfigMaps(ctx context.Context, configMaps v1_sets.ConfigMapSet, opts ...client.ListOption) error {
	configMapList, err := b.configMaps.ListConfigMap(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range configMapList.Items {
		item := item // pike
		configMaps.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertServices(ctx context.Context, services v1_sets.ServiceSet, opts ...client.ListOption) error {
	serviceList, err := b.services.ListService(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range serviceList.Items {
		item := item // pike
		services.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertPods(ctx context.Context, pods v1_sets.PodSet, opts ...client.ListOption) error {
	podList, err := b.pods.ListPod(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range podList.Items {
		item := item // pike
		pods.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertNodes(ctx context.Context, nodes v1_sets.NodeSet, opts ...client.ListOption) error {
	nodeList, err := b.nodes.ListNode(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range nodeList.Items {
		item := item // pike
		nodes.Insert(&item)
	}

	return nil
}

func (b *singleClusterBuilder) insertDeployments(ctx context.Context, deployments apps_v1_sets.DeploymentSet, opts ...client.ListOption) error {
	deploymentList, err := b.deployments.ListDeployment(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range deploymentList.Items {
		item := item // pike
		deployments.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertReplicaSets(ctx context.Context, replicaSets apps_v1_sets.ReplicaSetSet, opts ...client.ListOption) error {
	replicaSetList, err := b.replicaSets.ListReplicaSet(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range replicaSetList.Items {
		item := item // pike
		replicaSets.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertDaemonSets(ctx context.Context, daemonSets apps_v1_sets.DaemonSetSet, opts ...client.ListOption) error {
	daemonSetList, err := b.daemonSets.ListDaemonSet(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range daemonSetList.Items {
		item := item // pike
		daemonSets.Insert(&item)
	}

	return nil
}
func (b *singleClusterBuilder) insertStatefulSets(ctx context.Context, statefulSets apps_v1_sets.StatefulSetSet, opts ...client.ListOption) error {
	statefulSetList, err := b.statefulSets.ListStatefulSet(ctx, opts...)
	if err != nil {
		return err
	}

	for _, item := range statefulSetList.Items {
		item := item // pike
		statefulSets.Insert(&item)
	}

	return nil
}

// Utility for manually building input snapshots.
type InputSnapshotBuilder struct {
	name string

	configMaps v1_sets.ConfigMapSet
	services   v1_sets.ServiceSet
	pods       v1_sets.PodSet
	nodes      v1_sets.NodeSet

	deployments  apps_v1_sets.DeploymentSet
	replicaSets  apps_v1_sets.ReplicaSetSet
	daemonSets   apps_v1_sets.DaemonSetSet
	statefulSets apps_v1_sets.StatefulSetSet
}

func NewInputSnapshotBuilder(name string) *InputSnapshotBuilder {
	return &InputSnapshotBuilder{
		name: name,

		configMaps: v1_sets.NewConfigMapSet(),
		services:   v1_sets.NewServiceSet(),
		pods:       v1_sets.NewPodSet(),
		nodes:      v1_sets.NewNodeSet(),

		deployments:  apps_v1_sets.NewDeploymentSet(),
		replicaSets:  apps_v1_sets.NewReplicaSetSet(),
		daemonSets:   apps_v1_sets.NewDaemonSetSet(),
		statefulSets: apps_v1_sets.NewStatefulSetSet(),
	}
}

func (i *InputSnapshotBuilder) Build() Snapshot {
	return NewSnapshot(
		i.name,

		i.configMaps,
		i.services,
		i.pods,
		i.nodes,

		i.deployments,
		i.replicaSets,
		i.daemonSets,
		i.statefulSets,
	)
}
func (i *InputSnapshotBuilder) AddConfigMaps(configMaps []*v1.ConfigMap) *InputSnapshotBuilder {
	i.configMaps.Insert(configMaps...)
	return i
}
func (i *InputSnapshotBuilder) AddServices(services []*v1.Service) *InputSnapshotBuilder {
	i.services.Insert(services...)
	return i
}
func (i *InputSnapshotBuilder) AddPods(pods []*v1.Pod) *InputSnapshotBuilder {
	i.pods.Insert(pods...)
	return i
}
func (i *InputSnapshotBuilder) AddNodes(nodes []*v1.Node) *InputSnapshotBuilder {
	i.nodes.Insert(nodes...)
	return i
}
func (i *InputSnapshotBuilder) AddDeployments(deployments []*apps_v1.Deployment) *InputSnapshotBuilder {
	i.deployments.Insert(deployments...)
	return i
}
func (i *InputSnapshotBuilder) AddReplicaSets(replicaSets []*apps_v1.ReplicaSet) *InputSnapshotBuilder {
	i.replicaSets.Insert(replicaSets...)
	return i
}
func (i *InputSnapshotBuilder) AddDaemonSets(daemonSets []*apps_v1.DaemonSet) *InputSnapshotBuilder {
	i.daemonSets.Insert(daemonSets...)
	return i
}
func (i *InputSnapshotBuilder) AddStatefulSets(statefulSets []*apps_v1.StatefulSet) *InputSnapshotBuilder {
	i.statefulSets.Insert(statefulSets...)
	return i
}
