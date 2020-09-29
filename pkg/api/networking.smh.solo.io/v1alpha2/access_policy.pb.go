// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/access_policy.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha2/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Access control policies apply ALLOW policies to communication in a mesh.
//Access control policies specify the following:
//ALLOW those requests that: originate from from **source workload**, target the **destination target**,
//and match the indicated request criteria (allowed_paths, allowed_methods, allowed_ports).
//Enforcement of access control is determined by the
//[VirtualMesh's GlobalAccessPolicy]({{% versioned_link_path fromRoot="/reference/api/virtual_mesh/#networking.smh.solo.io.VirtualMeshSpec.GlobalAccessPolicy" %}})
type AccessPolicySpec struct {
	//
	//Requests originating from these pods will have the rule applied.
	//Leave empty to have all pods in the mesh apply these policies.
	//
	//Note that access control policies are mapped to source pods by their
	//service account. If other pods share the same service account,
	//this access control rule will apply to those pods as well.
	//
	//For fine-grained access control policies, ensure that your
	//service accounts properly reflect the desired
	//boundary for your access control policies.
	SourceSelector []*IdentitySelector `protobuf:"bytes,2,rep,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty"`
	//
	//Requests destined for these pods will have the rule applied.
	//Leave empty to apply to all destination pods in the mesh.
	DestinationSelector []*TrafficTargetSelector `protobuf:"bytes,3,rep,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty"`
	//
	//Optional. A list of HTTP paths or gRPC methods to allow.
	//gRPC methods must be presented as fully-qualified name in the form of
	//"/packageName.serviceName/methodName" and are case sensitive.
	//Exact match, prefix match, and suffix match are supported for paths.
	//For example, the path "/books/review" matches
	//"/books/review" (exact match), "*books/" (suffix match), or "/books*" (prefix match).
	//
	//If not specified, allow any path.
	AllowedPaths []string `protobuf:"bytes,4,rep,name=allowed_paths,json=allowedPaths,proto3" json:"allowed_paths,omitempty"`
	//
	//Optional. A list of HTTP methods to allow (e.g., "GET", "POST").
	//It is ignored in gRPC case because the value is always "POST".
	//If not specified, allows any method.
	AllowedMethods []types.HttpMethodValue `protobuf:"varint,5,rep,packed,name=allowed_methods,json=allowedMethods,proto3,enum=networking.smh.solo.io.HttpMethodValue" json:"allowed_methods,omitempty"`
	//
	//Optional. A list of ports which to allow.
	//If not set any port is allowed.
	AllowedPorts         []uint32 `protobuf:"varint,6,rep,packed,name=allowed_ports,json=allowedPorts,proto3" json:"allowed_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicySpec) Reset()         { *m = AccessPolicySpec{} }
func (m *AccessPolicySpec) String() string { return proto.CompactTextString(m) }
func (*AccessPolicySpec) ProtoMessage()    {}
func (*AccessPolicySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a607a654bf8f02aa, []int{0}
}
func (m *AccessPolicySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicySpec.Unmarshal(m, b)
}
func (m *AccessPolicySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicySpec.Marshal(b, m, deterministic)
}
func (m *AccessPolicySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicySpec.Merge(m, src)
}
func (m *AccessPolicySpec) XXX_Size() int {
	return xxx_messageInfo_AccessPolicySpec.Size(m)
}
func (m *AccessPolicySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicySpec.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicySpec proto.InternalMessageInfo

func (m *AccessPolicySpec) GetSourceSelector() []*IdentitySelector {
	if m != nil {
		return m.SourceSelector
	}
	return nil
}

func (m *AccessPolicySpec) GetDestinationSelector() []*TrafficTargetSelector {
	if m != nil {
		return m.DestinationSelector
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedPaths() []string {
	if m != nil {
		return m.AllowedPaths
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedMethods() []types.HttpMethodValue {
	if m != nil {
		return m.AllowedMethods
	}
	return nil
}

func (m *AccessPolicySpec) GetAllowedPorts() []uint32 {
	if m != nil {
		return m.AllowedPorts
	}
	return nil
}

type AccessPolicyStatus struct {
	// The most recent generation observed in the the AccessPolicy metadata.
	// If the observedGeneration does not match generation, the controller has not received the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	// It will only show accepted if it has been successfully
	// applied to all target meshes.
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.smh.solo.io.ApprovalState" json:"state,omitempty"`
	// The status of the AccessPolicy for each TrafficTarget to which it has been applied.
	// An AccessPolicy may be Accepted for some TrafficTargets and rejected for others.
	TrafficTargets map[string]*ApprovalStatus `protobuf:"bytes,3,rep,name=traffic_targets,json=trafficTargets,proto3" json:"traffic_targets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The list of Workloads to which this policy has been applied.
	Workloads []string `protobuf:"bytes,4,rep,name=workloads,proto3" json:"workloads,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors               []string `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPolicyStatus) Reset()         { *m = AccessPolicyStatus{} }
func (m *AccessPolicyStatus) String() string { return proto.CompactTextString(m) }
func (*AccessPolicyStatus) ProtoMessage()    {}
func (*AccessPolicyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_a607a654bf8f02aa, []int{1}
}
func (m *AccessPolicyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPolicyStatus.Unmarshal(m, b)
}
func (m *AccessPolicyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPolicyStatus.Marshal(b, m, deterministic)
}
func (m *AccessPolicyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPolicyStatus.Merge(m, src)
}
func (m *AccessPolicyStatus) XXX_Size() int {
	return xxx_messageInfo_AccessPolicyStatus.Size(m)
}
func (m *AccessPolicyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPolicyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPolicyStatus proto.InternalMessageInfo

func (m *AccessPolicyStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *AccessPolicyStatus) GetState() ApprovalState {
	if m != nil {
		return m.State
	}
	return ApprovalState_PENDING
}

func (m *AccessPolicyStatus) GetTrafficTargets() map[string]*ApprovalStatus {
	if m != nil {
		return m.TrafficTargets
	}
	return nil
}

func (m *AccessPolicyStatus) GetWorkloads() []string {
	if m != nil {
		return m.Workloads
	}
	return nil
}

func (m *AccessPolicyStatus) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessPolicySpec)(nil), "networking.smh.solo.io.AccessPolicySpec")
	proto.RegisterType((*AccessPolicyStatus)(nil), "networking.smh.solo.io.AccessPolicyStatus")
	proto.RegisterMapType((map[string]*ApprovalStatus)(nil), "networking.smh.solo.io.AccessPolicyStatus.TrafficTargetsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/access_policy.proto", fileDescriptor_a607a654bf8f02aa)
}

var fileDescriptor_a607a654bf8f02aa = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x69, 0x63, 0x0b, 0xdd, 0xf3, 0xda, 0x63, 0x7b, 0x1c, 0xa1, 0x88, 0x84, 0x13, 0x35,
	0x2f, 0x4d, 0xb0, 0xf7, 0x72, 0xa8, 0x28, 0x27, 0x88, 0x8a, 0x88, 0xbd, 0xdc, 0xe1, 0x83, 0x2f,
	0x71, 0x9b, 0xcc, 0x25, 0x4b, 0xd3, 0xec, 0xb2, 0x3b, 0xe9, 0xd1, 0xef, 0xe0, 0x07, 0xf1, 0x73,
	0xe9, 0x17, 0x91, 0x6c, 0x92, 0xf6, 0xaa, 0x57, 0xe8, 0xdb, 0xe6, 0xbf, 0x33, 0xbf, 0x99, 0xfd,
	0x0f, 0x13, 0xf2, 0x35, 0xe1, 0x98, 0x16, 0x33, 0x2f, 0x12, 0x0b, 0x5f, 0x8b, 0x4c, 0x8c, 0xb9,
	0xf0, 0x35, 0xa8, 0x25, 0x8f, 0x60, 0xbc, 0x00, 0x9d, 0x8e, 0xd3, 0x62, 0xe6, 0x33, 0xc9, 0xfd,
	0x1c, 0xf0, 0x56, 0xa8, 0x39, 0xcf, 0x13, 0x7f, 0xf9, 0x82, 0x65, 0x32, 0x65, 0x13, 0x9f, 0x45,
	0x11, 0x68, 0x1d, 0x4a, 0x91, 0xf1, 0x68, 0xe5, 0x49, 0x25, 0x50, 0xd0, 0x93, 0x4d, 0xa0, 0xa7,
	0x17, 0xa9, 0x57, 0x42, 0x3d, 0x2e, 0x46, 0x67, 0x7b, 0x53, 0x53, 0x44, 0x59, 0xc1, 0x46, 0xe7,
	0x7b, 0x27, 0x69, 0xc8, 0x20, 0x42, 0xa1, 0x74, 0x9d, 0xf9, 0x76, 0xef, 0xcc, 0x25, 0xcb, 0x78,
	0xcc, 0x90, 0x8b, 0x3c, 0xd4, 0xc8, 0x10, 0x6a, 0xc0, 0x71, 0x22, 0x12, 0x61, 0x8e, 0x7e, 0x79,
	0xaa, 0xd4, 0xd3, 0x3f, 0x6d, 0x72, 0x74, 0x61, 0x5e, 0x3d, 0x35, 0x8f, 0xbe, 0x92, 0x10, 0xd1,
	0x4b, 0x32, 0xd0, 0xa2, 0x50, 0x11, 0x84, 0x4d, 0x17, 0x76, 0xdb, 0xb1, 0xdc, 0x83, 0x89, 0xeb,
	0xdd, 0x6f, 0x86, 0xf7, 0x29, 0x86, 0x1c, 0x39, 0xae, 0xae, 0xea, 0xf8, 0xa0, 0x5f, 0x01, 0x9a,
	0x6f, 0xfa, 0x83, 0x1c, 0xc7, 0xa0, 0x91, 0xe7, 0x75, 0x63, 0x0d, 0xd7, 0x32, 0xdc, 0xf1, 0x2e,
	0xee, 0xb5, 0x62, 0x37, 0x37, 0x3c, 0xba, 0x66, 0x2a, 0x01, 0x5c, 0xc3, 0x87, 0x77, 0x50, 0xeb,
	0x0a, 0x4f, 0xc8, 0x21, 0xcb, 0x32, 0x71, 0x0b, 0x71, 0x28, 0x19, 0xa6, 0xda, 0x7e, 0xe0, 0x58,
	0x6e, 0x2f, 0x78, 0x58, 0x8b, 0xd3, 0x52, 0xa3, 0x53, 0x32, 0x68, 0x82, 0x16, 0x80, 0xa9, 0x88,
	0xb5, 0xdd, 0x71, 0x2c, 0xb7, 0x3f, 0x79, 0xbe, 0xab, 0x83, 0x8f, 0x88, 0xf2, 0x8b, 0x09, 0xfd,
	0xc6, 0xb2, 0x02, 0x82, 0x7e, 0x9d, 0x5f, 0x69, 0x7a, 0xab, 0xac, 0x50, 0xa8, 0xed, 0xae, 0x63,
	0xb9, 0x87, 0x9b, 0xb2, 0xa5, 0x76, 0xfa, 0xd3, 0x22, 0x74, 0xcb, 0x65, 0x64, 0x58, 0x68, 0xea,
	0x93, 0xa1, 0x98, 0x95, 0x73, 0x85, 0x38, 0x4c, 0x20, 0x07, 0x65, 0x5e, 0x64, 0xb7, 0x9c, 0x96,
	0x6b, 0x05, 0xb4, 0xb9, 0xfa, 0xb0, 0xbe, 0xa1, 0xaf, 0x48, 0xc7, 0x8c, 0xd4, 0x6e, 0x3b, 0x2d,
	0xb7, 0x3f, 0x79, 0xba, 0xab, 0xe9, 0x0b, 0x29, 0x95, 0x58, 0xb2, 0xac, 0xac, 0x03, 0x41, 0x95,
	0x43, 0x13, 0x32, 0xc0, 0xca, 0xce, 0x10, 0x8d, 0x9f, 0xba, 0x76, 0xff, 0xcd, 0x4e, 0xcc, 0x7f,
	0x2d, 0x6f, 0x0f, 0x44, 0xbf, 0xcf, 0x51, 0xad, 0x82, 0x3e, 0x6e, 0x89, 0xf4, 0x11, 0xe9, 0x95,
	0xb4, 0x4c, 0xb0, 0xb8, 0x99, 0xc2, 0x46, 0xa0, 0x27, 0xa4, 0x0b, 0x4a, 0x09, 0x55, 0x39, 0xdf,
	0x0b, 0xea, 0xaf, 0x11, 0x27, 0xc3, 0x7b, 0xe0, 0xf4, 0x88, 0x58, 0x73, 0x58, 0x19, 0x4f, 0x7a,
	0x41, 0x79, 0xa4, 0xaf, 0x49, 0x67, 0x59, 0x8e, 0xc2, 0x98, 0x70, 0x30, 0x79, 0xb6, 0x8f, 0x09,
	0x85, 0x0e, 0xaa, 0xa4, 0x97, 0xed, 0xf3, 0xd6, 0xbb, 0xcb, 0x5f, 0xbf, 0x1f, 0xb7, 0xbe, 0x7f,
	0xde, 0xe7, 0x4f, 0x21, 0xe7, 0xc9, 0x3f, 0x8b, 0x76, 0xb7, 0xc6, 0x7a, 0xe9, 0x66, 0x5d, 0xb3,
	0x4e, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x66, 0x61, 0xb9, 0x7f, 0x04, 0x00, 0x00,
}

func (this *AccessPolicySpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicySpec)
	if !ok {
		that2, ok := that.(AccessPolicySpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.SourceSelector) != len(that1.SourceSelector) {
		return false
	}
	for i := range this.SourceSelector {
		if !this.SourceSelector[i].Equal(that1.SourceSelector[i]) {
			return false
		}
	}
	if len(this.DestinationSelector) != len(that1.DestinationSelector) {
		return false
	}
	for i := range this.DestinationSelector {
		if !this.DestinationSelector[i].Equal(that1.DestinationSelector[i]) {
			return false
		}
	}
	if len(this.AllowedPaths) != len(that1.AllowedPaths) {
		return false
	}
	for i := range this.AllowedPaths {
		if this.AllowedPaths[i] != that1.AllowedPaths[i] {
			return false
		}
	}
	if len(this.AllowedMethods) != len(that1.AllowedMethods) {
		return false
	}
	for i := range this.AllowedMethods {
		if this.AllowedMethods[i] != that1.AllowedMethods[i] {
			return false
		}
	}
	if len(this.AllowedPorts) != len(that1.AllowedPorts) {
		return false
	}
	for i := range this.AllowedPorts {
		if this.AllowedPorts[i] != that1.AllowedPorts[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessPolicyStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessPolicyStatus)
	if !ok {
		that2, ok := that.(AccessPolicyStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.TrafficTargets) != len(that1.TrafficTargets) {
		return false
	}
	for i := range this.TrafficTargets {
		if !this.TrafficTargets[i].Equal(that1.TrafficTargets[i]) {
			return false
		}
	}
	if len(this.Workloads) != len(that1.Workloads) {
		return false
	}
	for i := range this.Workloads {
		if this.Workloads[i] != that1.Workloads[i] {
			return false
		}
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if this.Errors[i] != that1.Errors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
