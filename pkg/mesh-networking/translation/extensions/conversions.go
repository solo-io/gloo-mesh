package extensions

import (
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/extensions/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ObjectMetaToProto constructs a proto-compatible version of a k8s ObjectMeta
func ObjectMetaToProto(meta metav1.ObjectMeta) *v1alpha1.ObjectMeta {
	return &v1alpha1.ObjectMeta{
		Name:        meta.Name,
		Namespace:   meta.Namespace,
		ClusterName: meta.ClusterName,
		Labels:      meta.Labels,
		Annotations: meta.Annotations,
	}
}

// ObjectMetaToProto constructs a k8s ObjectMeta from a proto-compatible version
func ObjectMetaFromProto(meta *v1alpha1.ObjectMeta) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        meta.GetName(),
		Namespace:   meta.GetNamespace(),
		ClusterName: meta.GetClusterName(),
		Labels:      meta.GetLabels(),
		Annotations: meta.GetAnnotations(),
	}
}
