// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/networking/v1alpha2/access_policy.proto

package v1alpha2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1alpha2"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Grants communication permission between selected identities (i.e. traffic sources) and Destinations (i.e. destinations).
// Explicitly granted access permission is required if a
// [VirtualMesh's GlobalAccessPolicy]({{% versioned_link_path fromRoot="/reference/api/virtual_mesh/#networking.mesh.gloo.solo.io.VirtualMeshSpec.GlobalAccessPolicy" %}})
// is set to `ENABLED`.
type AccessPolicySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the identities of Workloads (i.e. traffic sources) for which to apply this AccessPolicy.
	// Leave empty to apply the AccessPolicy to all Workloads colocated in the destination's Mesh.
	SourceSelector []*v1alpha2.IdentitySelector `protobuf:"bytes,1,rep,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty"`
	// Specify the Destinations for which to apply this AccessPolicy.
	// Leave empty to apply the AccessPolicy to all Destinations.
	DestinationSelector []*v1alpha2.DestinationSelector `protobuf:"bytes,2,rep,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty"`
	//
	//Optional.* A list of HTTP paths or gRPC methods to allow.
	//gRPC methods must be presented as fully-qualified name in the form of
	//"/packageName.serviceName/methodName" and are case sensitive.
	//Exact match, prefix match, and suffix match are supported for paths.
	//For example, the path "/books/review" matches
	//"/books/review" (exact match), "*books/" (suffix match), or "/books*" (prefix match).
	//
	//If not specified, allow any path.
	AllowedPaths []string `protobuf:"bytes,3,rep,name=allowed_paths,json=allowedPaths,proto3" json:"allowed_paths,omitempty"`
	//
	//Optional.* A list of HTTP methods to allow (e.g., "GET", "POST").
	//It is ignored in gRPC case because the value is always "POST".
	//If not specified, allows any method.
	AllowedMethods []string `protobuf:"bytes,4,rep,name=allowed_methods,json=allowedMethods,proto3" json:"allowed_methods,omitempty"`
	//
	//Optional.* A list of ports which to allow.
	//If not set any port is allowed.
	AllowedPorts []uint32 `protobuf:"varint,5,rep,packed,name=allowed_ports,json=allowedPorts,proto3" json:"allowed_ports,omitempty"`
}

func (x *AccessPolicySpec) Reset() {
	*x = AccessPolicySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPolicySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPolicySpec) ProtoMessage() {}

func (x *AccessPolicySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPolicySpec.ProtoReflect.Descriptor instead.
func (*AccessPolicySpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescGZIP(), []int{0}
}

func (x *AccessPolicySpec) GetSourceSelector() []*v1alpha2.IdentitySelector {
	if x != nil {
		return x.SourceSelector
	}
	return nil
}

func (x *AccessPolicySpec) GetDestinationSelector() []*v1alpha2.DestinationSelector {
	if x != nil {
		return x.DestinationSelector
	}
	return nil
}

func (x *AccessPolicySpec) GetAllowedPaths() []string {
	if x != nil {
		return x.AllowedPaths
	}
	return nil
}

func (x *AccessPolicySpec) GetAllowedMethods() []string {
	if x != nil {
		return x.AllowedMethods
	}
	return nil
}

func (x *AccessPolicySpec) GetAllowedPorts() []uint32 {
	if x != nil {
		return x.AllowedPorts
	}
	return nil
}

type AccessPolicyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the AccessPolicy metadata.
	// If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	// It will only show accepted if it has been successfully applied to selected Destinations.
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// The status of the AccessPolicy for each Destination to which it has been applied.
	// An AccessPolicy may be accepted for some Destinations and rejected for others.
	Destinations map[string]*ApprovalStatus `protobuf:"bytes,3,rep,name=destinations,proto3" json:"destinations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The list of Workloads to which this policy has been applied.
	Workloads []string `protobuf:"bytes,4,rep,name=workloads,proto3" json:"workloads,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors []string `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *AccessPolicyStatus) Reset() {
	*x = AccessPolicyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPolicyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPolicyStatus) ProtoMessage() {}

func (x *AccessPolicyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPolicyStatus.ProtoReflect.Descriptor instead.
func (*AccessPolicyStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescGZIP(), []int{1}
}

func (x *AccessPolicyStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *AccessPolicyStatus) GetState() ApprovalState {
	if x != nil {
		return x.State
	}
	return ApprovalState_PENDING
}

func (x *AccessPolicyStatus) GetDestinations() map[string]*ApprovalStatus {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *AccessPolicyStatus) GetWorkloads() []string {
	if x != nil {
		return x.Workloads
	}
	return nil
}

func (x *AccessPolicyStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x10,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x53, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x60, 0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x95, 0x03, 0x0a, 0x12, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x1a, 0x6d, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x50, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_goTypes = []interface{}{
	(*AccessPolicySpec)(nil),             // 0: networking.mesh.gloo.solo.io.AccessPolicySpec
	(*AccessPolicyStatus)(nil),           // 1: networking.mesh.gloo.solo.io.AccessPolicyStatus
	nil,                                  // 2: networking.mesh.gloo.solo.io.AccessPolicyStatus.DestinationsEntry
	(*v1alpha2.IdentitySelector)(nil),    // 3: common.mesh.gloo.solo.io.IdentitySelector
	(*v1alpha2.DestinationSelector)(nil), // 4: common.mesh.gloo.solo.io.DestinationSelector
	(ApprovalState)(0),                   // 5: networking.mesh.gloo.solo.io.ApprovalState
	(*ApprovalStatus)(nil),               // 6: networking.mesh.gloo.solo.io.ApprovalStatus
}
var file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_depIdxs = []int32{
	3, // 0: networking.mesh.gloo.solo.io.AccessPolicySpec.source_selector:type_name -> common.mesh.gloo.solo.io.IdentitySelector
	4, // 1: networking.mesh.gloo.solo.io.AccessPolicySpec.destination_selector:type_name -> common.mesh.gloo.solo.io.DestinationSelector
	5, // 2: networking.mesh.gloo.solo.io.AccessPolicyStatus.state:type_name -> networking.mesh.gloo.solo.io.ApprovalState
	2, // 3: networking.mesh.gloo.solo.io.AccessPolicyStatus.destinations:type_name -> networking.mesh.gloo.solo.io.AccessPolicyStatus.DestinationsEntry
	6, // 4: networking.mesh.gloo.solo.io.AccessPolicyStatus.DestinationsEntry.value:type_name -> networking.mesh.gloo.solo.io.ApprovalStatus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_validation_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPolicySpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPolicyStatus); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1alpha2_access_policy_proto_depIdxs = nil
}
