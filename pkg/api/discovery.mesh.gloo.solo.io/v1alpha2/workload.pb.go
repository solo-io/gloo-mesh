// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/discovery/v1alpha2/workload.proto

package v1alpha2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Describes a workload controlled by a discovered service mesh.
type WorkloadSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes platform specific properties of the workload.
	//
	// Types that are assignable to Type:
	//	*WorkloadSpec_Kubernetes
	Type isWorkloadSpec_Type `protobuf_oneof:"type"`
	// The Mesh with which this Workload is associated.
	Mesh *v1.ObjectRef `protobuf:"bytes,4,opt,name=mesh,proto3" json:"mesh,omitempty"`
	// Metadata specific to an App Mesh controlled workload.
	AppMesh *WorkloadSpec_AppMesh `protobuf:"bytes,5,opt,name=app_mesh,json=appMesh,proto3" json:"app_mesh,omitempty"`
}

func (x *WorkloadSpec) Reset() {
	*x = WorkloadSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadSpec) ProtoMessage() {}

func (x *WorkloadSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadSpec.ProtoReflect.Descriptor instead.
func (*WorkloadSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{0}
}

func (m *WorkloadSpec) GetType() isWorkloadSpec_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *WorkloadSpec) GetKubernetes() *WorkloadSpec_KubernetesWorkload {
	if x, ok := x.GetType().(*WorkloadSpec_Kubernetes); ok {
		return x.Kubernetes
	}
	return nil
}

func (x *WorkloadSpec) GetMesh() *v1.ObjectRef {
	if x != nil {
		return x.Mesh
	}
	return nil
}

func (x *WorkloadSpec) GetAppMesh() *WorkloadSpec_AppMesh {
	if x != nil {
		return x.AppMesh
	}
	return nil
}

type isWorkloadSpec_Type interface {
	isWorkloadSpec_Type()
}

type WorkloadSpec_Kubernetes struct {
	// Information describing workloads backed by Kubernetes Pods.
	Kubernetes *WorkloadSpec_KubernetesWorkload `protobuf:"bytes,1,opt,name=kubernetes,proto3,oneof"`
}

func (*WorkloadSpec_Kubernetes) isWorkloadSpec_Type() {}

type WorkloadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The observed generation of the Workload.
	// When this matches the Workload's `metadata.generation` it indicates that Gloo Mesh
	// has processed the latest version of the Workload.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The set of AccessLogRecords that have been applied to this Workload.
	AppliedAccessLogRecords []*WorkloadStatus_AppliedAccessLogRecord `protobuf:"bytes,2,rep,name=applied_access_log_records,json=appliedAccessLogRecords,proto3" json:"applied_access_log_records,omitempty"`
	// The set of WasmDeployments that have been applied to this Workload.
	AppliedWasmDeployments []*WorkloadStatus_AppliedWasmDeployment `protobuf:"bytes,3,rep,name=applied_wasm_deployments,json=appliedWasmDeployments,proto3" json:"applied_wasm_deployments,omitempty"`
}

func (x *WorkloadStatus) Reset() {
	*x = WorkloadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadStatus) ProtoMessage() {}

func (x *WorkloadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadStatus.ProtoReflect.Descriptor instead.
func (*WorkloadStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *WorkloadStatus) GetAppliedAccessLogRecords() []*WorkloadStatus_AppliedAccessLogRecord {
	if x != nil {
		return x.AppliedAccessLogRecords
	}
	return nil
}

func (x *WorkloadStatus) GetAppliedWasmDeployments() []*WorkloadStatus_AppliedWasmDeployment {
	if x != nil {
		return x.AppliedWasmDeployments
	}
	return nil
}

// Describes a Kubernetes workload (e.g. a Deployment or DaemonSet).
type WorkloadSpec_KubernetesWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource reference to the Kubernetes Pod controller (i.e. Deployment, ReplicaSet, DaemonSet) for this Workload..
	Controller *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	// Labels on the Pod itself (read from `metadata.labels`), which are used to determine which Services front this workload.
	PodLabels map[string]string `protobuf:"bytes,2,rep,name=pod_labels,json=podLabels,proto3" json:"pod_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Service account associated with the Pods owned by this controller.
	ServiceAccountName string `protobuf:"bytes,3,opt,name=service_account_name,json=serviceAccountName,proto3" json:"service_account_name,omitempty"`
}

func (x *WorkloadSpec_KubernetesWorkload) Reset() {
	*x = WorkloadSpec_KubernetesWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadSpec_KubernetesWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadSpec_KubernetesWorkload) ProtoMessage() {}

func (x *WorkloadSpec_KubernetesWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadSpec_KubernetesWorkload.ProtoReflect.Descriptor instead.
func (*WorkloadSpec_KubernetesWorkload) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WorkloadSpec_KubernetesWorkload) GetController() *v1.ClusterObjectRef {
	if x != nil {
		return x.Controller
	}
	return nil
}

func (x *WorkloadSpec_KubernetesWorkload) GetPodLabels() map[string]string {
	if x != nil {
		return x.PodLabels
	}
	return nil
}

func (x *WorkloadSpec_KubernetesWorkload) GetServiceAccountName() string {
	if x != nil {
		return x.ServiceAccountName
	}
	return ""
}

// Metadata specific to an App Mesh controlled workload.
type WorkloadSpec_AppMesh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value of the env var APPMESH_VIRTUAL_NODE_NAME on the App Mesh envoy proxy container.
	VirtualNodeName string `protobuf:"bytes,1,opt,name=virtual_node_name,json=virtualNodeName,proto3" json:"virtual_node_name,omitempty"`
	// Ports exposed by this workload. Needed for declaring App Mesh VirtualNode listeners.
	Ports []*WorkloadSpec_AppMesh_ContainerPort `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *WorkloadSpec_AppMesh) Reset() {
	*x = WorkloadSpec_AppMesh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadSpec_AppMesh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadSpec_AppMesh) ProtoMessage() {}

func (x *WorkloadSpec_AppMesh) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadSpec_AppMesh.ProtoReflect.Descriptor instead.
func (*WorkloadSpec_AppMesh) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WorkloadSpec_AppMesh) GetVirtualNodeName() string {
	if x != nil {
		return x.VirtualNodeName
	}
	return ""
}

func (x *WorkloadSpec_AppMesh) GetPorts() []*WorkloadSpec_AppMesh_ContainerPort {
	if x != nil {
		return x.Ports
	}
	return nil
}

// Kubernetes application container ports.
type WorkloadSpec_AppMesh_ContainerPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port     uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *WorkloadSpec_AppMesh_ContainerPort) Reset() {
	*x = WorkloadSpec_AppMesh_ContainerPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadSpec_AppMesh_ContainerPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadSpec_AppMesh_ContainerPort) ProtoMessage() {}

func (x *WorkloadSpec_AppMesh_ContainerPort) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadSpec_AppMesh_ContainerPort.ProtoReflect.Descriptor instead.
func (*WorkloadSpec_AppMesh_ContainerPort) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *WorkloadSpec_AppMesh_ContainerPort) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *WorkloadSpec_AppMesh_ContainerPort) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

// Describes an [AccessLogRecord]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.enterprise.observability.v1alpha1.access_logging/" >}}) that applies to this Workload.
type WorkloadStatus_AppliedAccessLogRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to the AccessLogRecord object.
	Ref *v1.ObjectRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// The observed generation of the accepted AccessLogRecord.
	ObservedGeneration int64 `protobuf:"varint,2,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	// Any errors encountered while processing the AccessLogRecord object
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *WorkloadStatus_AppliedAccessLogRecord) Reset() {
	*x = WorkloadStatus_AppliedAccessLogRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadStatus_AppliedAccessLogRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadStatus_AppliedAccessLogRecord) ProtoMessage() {}

func (x *WorkloadStatus_AppliedAccessLogRecord) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadStatus_AppliedAccessLogRecord.ProtoReflect.Descriptor instead.
func (*WorkloadStatus_AppliedAccessLogRecord) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{1, 0}
}

func (x *WorkloadStatus_AppliedAccessLogRecord) GetRef() *v1.ObjectRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *WorkloadStatus_AppliedAccessLogRecord) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *WorkloadStatus_AppliedAccessLogRecord) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

// Describes a [WasmDeployment]({{< versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.enterprise.networking.v1alpha1.wasm_deployment/" >}}) that applies to this Workload.
type WorkloadStatus_AppliedWasmDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to the WasmDeployment object.
	Ref *v1.ObjectRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// The observed generation of the WasmDeployment.
	ObservedGeneration int64 `protobuf:"varint,2,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	// Any errors encountered while processing the WasmDeployment object.
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *WorkloadStatus_AppliedWasmDeployment) Reset() {
	*x = WorkloadStatus_AppliedWasmDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadStatus_AppliedWasmDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadStatus_AppliedWasmDeployment) ProtoMessage() {}

func (x *WorkloadStatus_AppliedWasmDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadStatus_AppliedWasmDeployment.ProtoReflect.Descriptor instead.
func (*WorkloadStatus_AppliedWasmDeployment) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP(), []int{1, 1}
}

func (x *WorkloadStatus_AppliedWasmDeployment) GetRef() *v1.ObjectRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *WorkloadStatus_AppliedWasmDeployment) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *WorkloadStatus_AppliedWasmDeployment) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x05, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x12, 0x4c, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f,
	0x6d, 0x65, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x4d, 0x65, 0x73, 0x68, 0x1a, 0xb5, 0x02, 0x0a, 0x12, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x43, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x6a, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x70, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x30,
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x3c, 0x0a, 0x0e, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xcd,
	0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x3f, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x06,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe4, 0x04, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7f, 0x0a, 0x1a, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x17, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x7b, 0x0a, 0x18, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x57, 0x61, 0x73, 0x6d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x16, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x57, 0x61, 0x73, 0x6d, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x90, 0x01, 0x0a, 0x16, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x2e, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x03,
	0x72, 0x65, 0x66, 0x12, 0x2e, 0x0a, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x8f, 0x01, 0x0a, 0x15,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x57, 0x61, 0x73, 0x6d, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66,
	0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x2e, 0x0a, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x4f, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_goTypes = []interface{}{
	(*WorkloadSpec)(nil),                          // 0: discovery.mesh.gloo.solo.io.WorkloadSpec
	(*WorkloadStatus)(nil),                        // 1: discovery.mesh.gloo.solo.io.WorkloadStatus
	(*WorkloadSpec_KubernetesWorkload)(nil),       // 2: discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload
	(*WorkloadSpec_AppMesh)(nil),                  // 3: discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh
	nil,                                           // 4: discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload.PodLabelsEntry
	(*WorkloadSpec_AppMesh_ContainerPort)(nil),    // 5: discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh.ContainerPort
	(*WorkloadStatus_AppliedAccessLogRecord)(nil), // 6: discovery.mesh.gloo.solo.io.WorkloadStatus.AppliedAccessLogRecord
	(*WorkloadStatus_AppliedWasmDeployment)(nil),  // 7: discovery.mesh.gloo.solo.io.WorkloadStatus.AppliedWasmDeployment
	(*v1.ObjectRef)(nil),                          // 8: core.skv2.solo.io.ObjectRef
	(*v1.ClusterObjectRef)(nil),                   // 9: core.skv2.solo.io.ClusterObjectRef
}
var file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_depIdxs = []int32{
	2,  // 0: discovery.mesh.gloo.solo.io.WorkloadSpec.kubernetes:type_name -> discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload
	8,  // 1: discovery.mesh.gloo.solo.io.WorkloadSpec.mesh:type_name -> core.skv2.solo.io.ObjectRef
	3,  // 2: discovery.mesh.gloo.solo.io.WorkloadSpec.app_mesh:type_name -> discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh
	6,  // 3: discovery.mesh.gloo.solo.io.WorkloadStatus.applied_access_log_records:type_name -> discovery.mesh.gloo.solo.io.WorkloadStatus.AppliedAccessLogRecord
	7,  // 4: discovery.mesh.gloo.solo.io.WorkloadStatus.applied_wasm_deployments:type_name -> discovery.mesh.gloo.solo.io.WorkloadStatus.AppliedWasmDeployment
	9,  // 5: discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload.controller:type_name -> core.skv2.solo.io.ClusterObjectRef
	4,  // 6: discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload.pod_labels:type_name -> discovery.mesh.gloo.solo.io.WorkloadSpec.KubernetesWorkload.PodLabelsEntry
	5,  // 7: discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh.ports:type_name -> discovery.mesh.gloo.solo.io.WorkloadSpec.AppMesh.ContainerPort
	8,  // 8: discovery.mesh.gloo.solo.io.WorkloadStatus.AppliedAccessLogRecord.ref:type_name -> core.skv2.solo.io.ObjectRef
	8,  // 9: discovery.mesh.gloo.solo.io.WorkloadStatus.AppliedWasmDeployment.ref:type_name -> core.skv2.solo.io.ObjectRef
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadSpec_KubernetesWorkload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadSpec_AppMesh); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadSpec_AppMesh_ContainerPort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadStatus_AppliedAccessLogRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadStatus_AppliedWasmDeployment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WorkloadSpec_Kubernetes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_discovery_v1alpha2_workload_proto_depIdxs = nil
}
