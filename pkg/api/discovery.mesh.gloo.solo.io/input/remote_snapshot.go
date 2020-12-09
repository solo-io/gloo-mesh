// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./remote_snapshot.go -destination mocks/remote_snapshot.go

// The Input RemoteSnapshot contains the set of all:
// * Meshes
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
// or a ClusterWatcher (for multiple clusters) using the RemoteSnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appmesh_k8s_aws_v1beta2 "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2"
	appmesh_k8s_aws_v1beta2_sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.k8s.aws/v1beta2/sets"

	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"

	apps_v1 "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	apps_v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"
)

// the snapshot of input resources consumed by translation
type RemoteSnapshot interface {

	// return the set of input Meshes
	Meshes() appmesh_k8s_aws_v1beta2_sets.MeshSet

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

// options for syncing input object statuses
type RemoteSyncStatusOptions struct {

	// sync status of Mesh objects
	Mesh bool

	// sync status of ConfigMap objects
	ConfigMap bool
	// sync status of Service objects
	Service bool
	// sync status of Pod objects
	Pod bool
	// sync status of Node objects
	Node bool

	// sync status of Deployment objects
	Deployment bool
	// sync status of ReplicaSet objects
	ReplicaSet bool
	// sync status of DaemonSet objects
	DaemonSet bool
	// sync status of StatefulSet objects
	StatefulSet bool
}

type snapshotRemote struct {
	name string

	meshes appmesh_k8s_aws_v1beta2_sets.MeshSet

	configMaps v1_sets.ConfigMapSet
	services   v1_sets.ServiceSet
	pods       v1_sets.PodSet
	nodes      v1_sets.NodeSet

	deployments  apps_v1_sets.DeploymentSet
	replicaSets  apps_v1_sets.ReplicaSetSet
	daemonSets   apps_v1_sets.DaemonSetSet
	statefulSets apps_v1_sets.StatefulSetSet
}

func NewRemoteSnapshot(
	name string,

	meshes appmesh_k8s_aws_v1beta2_sets.MeshSet,

	configMaps v1_sets.ConfigMapSet,
	services v1_sets.ServiceSet,
	pods v1_sets.PodSet,
	nodes v1_sets.NodeSet,

	deployments apps_v1_sets.DeploymentSet,
	replicaSets apps_v1_sets.ReplicaSetSet,
	daemonSets apps_v1_sets.DaemonSetSet,
	statefulSets apps_v1_sets.StatefulSetSet,

) RemoteSnapshot {
	return &snapshotRemote{
		name: name,

		meshes:       meshes,
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

func (s snapshotRemote) Meshes() appmesh_k8s_aws_v1beta2_sets.MeshSet {
	return s.meshes
}

func (s snapshotRemote) ConfigMaps() v1_sets.ConfigMapSet {
	return s.configMaps
}

func (s snapshotRemote) Services() v1_sets.ServiceSet {
	return s.services
}

func (s snapshotRemote) Pods() v1_sets.PodSet {
	return s.pods
}

func (s snapshotRemote) Nodes() v1_sets.NodeSet {
	return s.nodes
}

func (s snapshotRemote) Deployments() apps_v1_sets.DeploymentSet {
	return s.deployments
}

func (s snapshotRemote) ReplicaSets() apps_v1_sets.ReplicaSetSet {
	return s.replicaSets
}

func (s snapshotRemote) DaemonSets() apps_v1_sets.DaemonSetSet {
	return s.daemonSets
}

func (s snapshotRemote) StatefulSets() apps_v1_sets.StatefulSetSet {
	return s.statefulSets
}

func (s snapshotRemote) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["meshes"] = s.meshes.List()
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
type RemoteBuilder interface {
	BuildSnapshot(ctx context.Context, name string, opts RemoteBuildOptions) (RemoteSnapshot, error)
}

// Options for building a snapshot
type RemoteBuildOptions struct {

	// List options for composing a snapshot from Meshes
	Meshes ResourceRemoteBuildOptions

	// List options for composing a snapshot from ConfigMaps
	ConfigMaps ResourceRemoteBuildOptions
	// List options for composing a snapshot from Services
	Services ResourceRemoteBuildOptions
	// List options for composing a snapshot from Pods
	Pods ResourceRemoteBuildOptions
	// List options for composing a snapshot from Nodes
	Nodes ResourceRemoteBuildOptions

	// List options for composing a snapshot from Deployments
	Deployments ResourceRemoteBuildOptions
	// List options for composing a snapshot from ReplicaSets
	ReplicaSets ResourceRemoteBuildOptions
	// List options for composing a snapshot from DaemonSets
	DaemonSets ResourceRemoteBuildOptions
	// List options for composing a snapshot from StatefulSets
	StatefulSets ResourceRemoteBuildOptions
}

// Options for reading resources of a given type
type ResourceRemoteBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources across multiple clusters
type multiClusterRemoteBuilder struct {
	clusters multicluster.Interface
	client   multicluster.Client
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterRemoteBuilder(
	clusters multicluster.Interface,
	client multicluster.Client,
) RemoteBuilder {
	return &multiClusterRemoteBuilder{
		clusters: clusters,
		client:   client,
	}
}

func (b *multiClusterRemoteBuilder) BuildSnapshot(ctx context.Context, name string, opts RemoteBuildOptions) (RemoteSnapshot, error) {

	meshes := appmesh_k8s_aws_v1beta2_sets.NewMeshSet()

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

		if err := b.insertMeshesFromCluster(ctx, cluster, meshes, opts.Meshes); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertConfigMapsFromCluster(ctx, cluster, configMaps, opts.ConfigMaps); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertServicesFromCluster(ctx, cluster, services, opts.Services); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertPodsFromCluster(ctx, cluster, pods, opts.Pods); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertNodesFromCluster(ctx, cluster, nodes, opts.Nodes); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDeploymentsFromCluster(ctx, cluster, deployments, opts.Deployments); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertReplicaSetsFromCluster(ctx, cluster, replicaSets, opts.ReplicaSets); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDaemonSetsFromCluster(ctx, cluster, daemonSets, opts.DaemonSets); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertStatefulSetsFromCluster(ctx, cluster, statefulSets, opts.StatefulSets); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewRemoteSnapshot(
		name,

		meshes,
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

func (b *multiClusterRemoteBuilder) insertMeshesFromCluster(ctx context.Context, cluster string, meshes appmesh_k8s_aws_v1beta2_sets.MeshSet, opts ResourceRemoteBuildOptions) error {
	meshClient, err := appmesh_k8s_aws_v1beta2.NewMulticlusterMeshClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "appmesh.k8s.aws",
			Version: "v1beta2",
			Kind:    "Mesh",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	meshList, err := meshClient.ListMesh(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range meshList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		meshes.Insert(&item)
	}

	return nil
}

func (b *multiClusterRemoteBuilder) insertConfigMapsFromCluster(ctx context.Context, cluster string, configMaps v1_sets.ConfigMapSet, opts ResourceRemoteBuildOptions) error {
	configMapClient, err := v1.NewMulticlusterConfigMapClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "ConfigMap",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	configMapList, err := configMapClient.ListConfigMap(ctx, opts.ListOptions...)
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
func (b *multiClusterRemoteBuilder) insertServicesFromCluster(ctx context.Context, cluster string, services v1_sets.ServiceSet, opts ResourceRemoteBuildOptions) error {
	serviceClient, err := v1.NewMulticlusterServiceClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Service",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	serviceList, err := serviceClient.ListService(ctx, opts.ListOptions...)
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
func (b *multiClusterRemoteBuilder) insertPodsFromCluster(ctx context.Context, cluster string, pods v1_sets.PodSet, opts ResourceRemoteBuildOptions) error {
	podClient, err := v1.NewMulticlusterPodClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Pod",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	podList, err := podClient.ListPod(ctx, opts.ListOptions...)
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
func (b *multiClusterRemoteBuilder) insertNodesFromCluster(ctx context.Context, cluster string, nodes v1_sets.NodeSet, opts ResourceRemoteBuildOptions) error {
	nodeClient, err := v1.NewMulticlusterNodeClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Node",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	nodeList, err := nodeClient.ListNode(ctx, opts.ListOptions...)
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

func (b *multiClusterRemoteBuilder) insertDeploymentsFromCluster(ctx context.Context, cluster string, deployments apps_v1_sets.DeploymentSet, opts ResourceRemoteBuildOptions) error {
	deploymentClient, err := apps_v1.NewMulticlusterDeploymentClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "Deployment",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	deploymentList, err := deploymentClient.ListDeployment(ctx, opts.ListOptions...)
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
func (b *multiClusterRemoteBuilder) insertReplicaSetsFromCluster(ctx context.Context, cluster string, replicaSets apps_v1_sets.ReplicaSetSet, opts ResourceRemoteBuildOptions) error {
	replicaSetClient, err := apps_v1.NewMulticlusterReplicaSetClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "ReplicaSet",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	replicaSetList, err := replicaSetClient.ListReplicaSet(ctx, opts.ListOptions...)
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
func (b *multiClusterRemoteBuilder) insertDaemonSetsFromCluster(ctx context.Context, cluster string, daemonSets apps_v1_sets.DaemonSetSet, opts ResourceRemoteBuildOptions) error {
	daemonSetClient, err := apps_v1.NewMulticlusterDaemonSetClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "DaemonSet",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	daemonSetList, err := daemonSetClient.ListDaemonSet(ctx, opts.ListOptions...)
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
func (b *multiClusterRemoteBuilder) insertStatefulSetsFromCluster(ctx context.Context, cluster string, statefulSets apps_v1_sets.StatefulSetSet, opts ResourceRemoteBuildOptions) error {
	statefulSetClient, err := apps_v1.NewMulticlusterStatefulSetClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "StatefulSet",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	statefulSetList, err := statefulSetClient.ListStatefulSet(ctx, opts.ListOptions...)
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
