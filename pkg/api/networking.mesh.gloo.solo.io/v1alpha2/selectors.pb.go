// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1alpha2/selectors.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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
//Select TrafficTargets using one or more platform-specific selection objects.
type TrafficTargetSelector struct {
	// A KubeServiceMatcher matches kubernetes services by their labels, namespaces, and/or clusters.
	KubeServiceMatcher *TrafficTargetSelector_KubeServiceMatcher `protobuf:"bytes,1,opt,name=kube_service_matcher,json=kubeServiceMatcher,proto3" json:"kube_service_matcher,omitempty"`
	// Match individual k8s Services by direct reference.
	KubeServiceRefs      *TrafficTargetSelector_KubeServiceRefs `protobuf:"bytes,2,opt,name=kube_service_refs,json=kubeServiceRefs,proto3" json:"kube_service_refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *TrafficTargetSelector) Reset()         { *m = TrafficTargetSelector{} }
func (m *TrafficTargetSelector) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetSelector) ProtoMessage()    {}
func (*TrafficTargetSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{0}
}
func (m *TrafficTargetSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetSelector.Unmarshal(m, b)
}
func (m *TrafficTargetSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetSelector.Marshal(b, m, deterministic)
}
func (m *TrafficTargetSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetSelector.Merge(m, src)
}
func (m *TrafficTargetSelector) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetSelector.Size(m)
}
func (m *TrafficTargetSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetSelector.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetSelector proto.InternalMessageInfo

func (m *TrafficTargetSelector) GetKubeServiceMatcher() *TrafficTargetSelector_KubeServiceMatcher {
	if m != nil {
		return m.KubeServiceMatcher
	}
	return nil
}

func (m *TrafficTargetSelector) GetKubeServiceRefs() *TrafficTargetSelector_KubeServiceRefs {
	if m != nil {
		return m.KubeServiceRefs
	}
	return nil
}

type TrafficTargetSelector_KubeServiceMatcher struct {
	//
	//If specified, all labels must exist on k8s Service.
	//When used in a networking policy, omission matches any labels.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any label key and/or value.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//
	//If specified, match k8s Services if they exist in one of the specified namespaces.
	//When used in a networking policy, omission matches any namespace.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any namespace.
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//If specified, match k8s Services if they exist in one of the specified clusters.
	//When used in a networking policy, omission matches any cluster.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any cluster.
	Clusters             []string `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficTargetSelector_KubeServiceMatcher) Reset() {
	*m = TrafficTargetSelector_KubeServiceMatcher{}
}
func (m *TrafficTargetSelector_KubeServiceMatcher) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetSelector_KubeServiceMatcher) ProtoMessage()    {}
func (*TrafficTargetSelector_KubeServiceMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{0, 0}
}
func (m *TrafficTargetSelector_KubeServiceMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetSelector_KubeServiceMatcher.Unmarshal(m, b)
}
func (m *TrafficTargetSelector_KubeServiceMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetSelector_KubeServiceMatcher.Marshal(b, m, deterministic)
}
func (m *TrafficTargetSelector_KubeServiceMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetSelector_KubeServiceMatcher.Merge(m, src)
}
func (m *TrafficTargetSelector_KubeServiceMatcher) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetSelector_KubeServiceMatcher.Size(m)
}
func (m *TrafficTargetSelector_KubeServiceMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetSelector_KubeServiceMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetSelector_KubeServiceMatcher proto.InternalMessageInfo

func (m *TrafficTargetSelector_KubeServiceMatcher) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TrafficTargetSelector_KubeServiceMatcher) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *TrafficTargetSelector_KubeServiceMatcher) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TrafficTargetSelector_KubeServiceRefs struct {
	//
	//Match k8s Services by direct reference.
	//When used in a networking policy, omission of any field (name, namespace, or clusterName) allows matching any value for that field.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any value for the given field.
	Services             []*v1.ClusterObjectRef `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TrafficTargetSelector_KubeServiceRefs) Reset()         { *m = TrafficTargetSelector_KubeServiceRefs{} }
func (m *TrafficTargetSelector_KubeServiceRefs) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetSelector_KubeServiceRefs) ProtoMessage()    {}
func (*TrafficTargetSelector_KubeServiceRefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{0, 1}
}
func (m *TrafficTargetSelector_KubeServiceRefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetSelector_KubeServiceRefs.Unmarshal(m, b)
}
func (m *TrafficTargetSelector_KubeServiceRefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetSelector_KubeServiceRefs.Marshal(b, m, deterministic)
}
func (m *TrafficTargetSelector_KubeServiceRefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetSelector_KubeServiceRefs.Merge(m, src)
}
func (m *TrafficTargetSelector_KubeServiceRefs) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetSelector_KubeServiceRefs.Size(m)
}
func (m *TrafficTargetSelector_KubeServiceRefs) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetSelector_KubeServiceRefs.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetSelector_KubeServiceRefs proto.InternalMessageInfo

func (m *TrafficTargetSelector_KubeServiceRefs) GetServices() []*v1.ClusterObjectRef {
	if m != nil {
		return m.Services
	}
	return nil
}

//
//Select Kubernetes workloads directly using label namespace and/or cluster criteria. See comments on the fields for
//detailed semantics.
type WorkloadSelector struct {
	//
	//If specified, all labels must exist on k8s workload.
	//When used in a networking policy, omission matches any labels.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any label key and/or value.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//
	//If specified, match k8s workloads if they exist in one of the specified namespaces.
	//When used in a networking policy, omission matches any namespace.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any namespace.
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//If specified, match k8s workloads if they exist in one of the specified clusters.
	//When used in a networking policy, omission matches any cluster.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any cluster.
	Clusters             []string `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadSelector) Reset()         { *m = WorkloadSelector{} }
func (m *WorkloadSelector) String() string { return proto.CompactTextString(m) }
func (*WorkloadSelector) ProtoMessage()    {}
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{1}
}
func (m *WorkloadSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSelector.Unmarshal(m, b)
}
func (m *WorkloadSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSelector.Marshal(b, m, deterministic)
}
func (m *WorkloadSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSelector.Merge(m, src)
}
func (m *WorkloadSelector) XXX_Size() int {
	return xxx_messageInfo_WorkloadSelector.Size(m)
}
func (m *WorkloadSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSelector.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSelector proto.InternalMessageInfo

func (m *WorkloadSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WorkloadSelector) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *WorkloadSelector) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

//
//Selector capable of selecting specific service identities. Useful for binding policy rules.
//Either (namespaces, cluster, service_account_names) or service_accounts can be specified.
//If all fields are omitted, any source identity is permitted.
type IdentitySelector struct {
	// A KubeIdentityMatcher matches request identities based on the k8s namespace and cluster.
	KubeIdentityMatcher *IdentitySelector_KubeIdentityMatcher `protobuf:"bytes,1,opt,name=kube_identity_matcher,json=kubeIdentityMatcher,proto3" json:"kube_identity_matcher,omitempty"`
	// KubeServiceAccountRefs matches request identities based on the k8s service account of request.
	KubeServiceAccountRefs *IdentitySelector_KubeServiceAccountRefs `protobuf:"bytes,2,opt,name=kube_service_account_refs,json=kubeServiceAccountRefs,proto3" json:"kube_service_account_refs,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                                 `json:"-"`
	XXX_unrecognized       []byte                                   `json:"-"`
	XXX_sizecache          int32                                    `json:"-"`
}

func (m *IdentitySelector) Reset()         { *m = IdentitySelector{} }
func (m *IdentitySelector) String() string { return proto.CompactTextString(m) }
func (*IdentitySelector) ProtoMessage()    {}
func (*IdentitySelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{2}
}
func (m *IdentitySelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentitySelector.Unmarshal(m, b)
}
func (m *IdentitySelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentitySelector.Marshal(b, m, deterministic)
}
func (m *IdentitySelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentitySelector.Merge(m, src)
}
func (m *IdentitySelector) XXX_Size() int {
	return xxx_messageInfo_IdentitySelector.Size(m)
}
func (m *IdentitySelector) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentitySelector.DiscardUnknown(m)
}

var xxx_messageInfo_IdentitySelector proto.InternalMessageInfo

func (m *IdentitySelector) GetKubeIdentityMatcher() *IdentitySelector_KubeIdentityMatcher {
	if m != nil {
		return m.KubeIdentityMatcher
	}
	return nil
}

func (m *IdentitySelector) GetKubeServiceAccountRefs() *IdentitySelector_KubeServiceAccountRefs {
	if m != nil {
		return m.KubeServiceAccountRefs
	}
	return nil
}

type IdentitySelector_KubeIdentityMatcher struct {
	//
	//If specified, match k8s identity if it exists in one of the specified namespaces.
	//When used in a networking policy, omission matches any namespace.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any namespace.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	//
	//If specified, match k8s identity if it exists in one of the specified clusters.
	//When used in a networking policy, omission matches any cluster.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any cluster.
	Clusters             []string `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentitySelector_KubeIdentityMatcher) Reset()         { *m = IdentitySelector_KubeIdentityMatcher{} }
func (m *IdentitySelector_KubeIdentityMatcher) String() string { return proto.CompactTextString(m) }
func (*IdentitySelector_KubeIdentityMatcher) ProtoMessage()    {}
func (*IdentitySelector_KubeIdentityMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{2, 0}
}
func (m *IdentitySelector_KubeIdentityMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentitySelector_KubeIdentityMatcher.Unmarshal(m, b)
}
func (m *IdentitySelector_KubeIdentityMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentitySelector_KubeIdentityMatcher.Marshal(b, m, deterministic)
}
func (m *IdentitySelector_KubeIdentityMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentitySelector_KubeIdentityMatcher.Merge(m, src)
}
func (m *IdentitySelector_KubeIdentityMatcher) XXX_Size() int {
	return xxx_messageInfo_IdentitySelector_KubeIdentityMatcher.Size(m)
}
func (m *IdentitySelector_KubeIdentityMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentitySelector_KubeIdentityMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_IdentitySelector_KubeIdentityMatcher proto.InternalMessageInfo

func (m *IdentitySelector_KubeIdentityMatcher) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *IdentitySelector_KubeIdentityMatcher) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type IdentitySelector_KubeServiceAccountRefs struct {
	//
	//Match k8s ServiceAccounts by direct reference.
	//When used in a networking policy, omission of any field (name, namespace, or clusterName) allows matching any value for that field.
	//When used in a Role, a wildcard `"*"` must be explicitly used to match any value for the given field.
	ServiceAccounts      []*v1.ClusterObjectRef `protobuf:"bytes,1,rep,name=service_accounts,json=serviceAccounts,proto3" json:"service_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IdentitySelector_KubeServiceAccountRefs) Reset() {
	*m = IdentitySelector_KubeServiceAccountRefs{}
}
func (m *IdentitySelector_KubeServiceAccountRefs) String() string { return proto.CompactTextString(m) }
func (*IdentitySelector_KubeServiceAccountRefs) ProtoMessage()    {}
func (*IdentitySelector_KubeServiceAccountRefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8340097348d3751, []int{2, 1}
}
func (m *IdentitySelector_KubeServiceAccountRefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentitySelector_KubeServiceAccountRefs.Unmarshal(m, b)
}
func (m *IdentitySelector_KubeServiceAccountRefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentitySelector_KubeServiceAccountRefs.Marshal(b, m, deterministic)
}
func (m *IdentitySelector_KubeServiceAccountRefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentitySelector_KubeServiceAccountRefs.Merge(m, src)
}
func (m *IdentitySelector_KubeServiceAccountRefs) XXX_Size() int {
	return xxx_messageInfo_IdentitySelector_KubeServiceAccountRefs.Size(m)
}
func (m *IdentitySelector_KubeServiceAccountRefs) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentitySelector_KubeServiceAccountRefs.DiscardUnknown(m)
}

var xxx_messageInfo_IdentitySelector_KubeServiceAccountRefs proto.InternalMessageInfo

func (m *IdentitySelector_KubeServiceAccountRefs) GetServiceAccounts() []*v1.ClusterObjectRef {
	if m != nil {
		return m.ServiceAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*TrafficTargetSelector)(nil), "networking.mesh.gloo.solo.io.TrafficTargetSelector")
	proto.RegisterType((*TrafficTargetSelector_KubeServiceMatcher)(nil), "networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher")
	proto.RegisterMapType((map[string]string)(nil), "networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceMatcher.LabelsEntry")
	proto.RegisterType((*TrafficTargetSelector_KubeServiceRefs)(nil), "networking.mesh.gloo.solo.io.TrafficTargetSelector.KubeServiceRefs")
	proto.RegisterType((*WorkloadSelector)(nil), "networking.mesh.gloo.solo.io.WorkloadSelector")
	proto.RegisterMapType((map[string]string)(nil), "networking.mesh.gloo.solo.io.WorkloadSelector.LabelsEntry")
	proto.RegisterType((*IdentitySelector)(nil), "networking.mesh.gloo.solo.io.IdentitySelector")
	proto.RegisterType((*IdentitySelector_KubeIdentityMatcher)(nil), "networking.mesh.gloo.solo.io.IdentitySelector.KubeIdentityMatcher")
	proto.RegisterType((*IdentitySelector_KubeServiceAccountRefs)(nil), "networking.mesh.gloo.solo.io.IdentitySelector.KubeServiceAccountRefs")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/networking/v1alpha2/selectors.proto", fileDescriptor_f8340097348d3751)
}

var fileDescriptor_f8340097348d3751 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x13, 0xa8, 0x9a, 0xcd, 0x21, 0x61, 0x9b, 0x56, 0xc1, 0xa0, 0xa8, 0x0a, 0x97, 0x5c,
	0xba, 0x56, 0xc3, 0x05, 0x7a, 0x41, 0xb4, 0x14, 0x89, 0x7f, 0xd8, 0x56, 0x42, 0xe2, 0x52, 0xad,
	0xb7, 0x63, 0xc7, 0xb5, 0xe3, 0xb5, 0x76, 0xd7, 0x2e, 0xb9, 0xf1, 0x38, 0x3c, 0x0b, 0x27, 0x5e,
	0x80, 0x0b, 0x2f, 0xc1, 0x15, 0x79, 0xed, 0x98, 0xc4, 0x89, 0x22, 0xf1, 0x73, 0xe0, 0xe4, 0xf5,
	0xec, 0xce, 0x37, 0x33, 0xdf, 0x7c, 0x33, 0xe8, 0x89, 0x1f, 0xe8, 0x49, 0xea, 0x12, 0x2e, 0xa6,
	0x8e, 0x12, 0x91, 0x38, 0x08, 0x84, 0xe3, 0x47, 0x42, 0x1c, 0x4c, 0x41, 0x4d, 0x1c, 0x96, 0x04,
	0x4e, 0x0c, 0xfa, 0x5a, 0xc8, 0x30, 0x88, 0x7d, 0x27, 0x3b, 0x64, 0x51, 0x32, 0x61, 0x63, 0x47,
	0x41, 0x04, 0x5c, 0x0b, 0xa9, 0x48, 0x22, 0x85, 0x16, 0xf8, 0xee, 0xaf, 0x47, 0x24, 0x77, 0x24,
	0x39, 0x04, 0xc9, 0xf1, 0x48, 0x20, 0xec, 0x81, 0x2f, 0x84, 0x1f, 0x81, 0x63, 0xde, 0xba, 0xa9,
	0xe7, 0x5c, 0x4b, 0x96, 0x24, 0x30, 0xf7, 0xb6, 0xef, 0xa8, 0x30, 0x1b, 0x9b, 0x58, 0x5c, 0x48,
	0x70, 0xb2, 0x43, 0xf3, 0x2d, 0x2f, 0x7b, 0xbe, 0xf0, 0x85, 0x39, 0x3a, 0xf9, 0xa9, 0xb0, 0x0e,
	0xbf, 0xde, 0x40, 0xbb, 0xe7, 0x92, 0x79, 0x5e, 0xc0, 0xcf, 0x99, 0xf4, 0x41, 0x9f, 0x95, 0x19,
	0xe1, 0x8f, 0xa8, 0x17, 0xa6, 0x2e, 0x5c, 0x28, 0x90, 0x59, 0xc0, 0xe1, 0x62, 0xca, 0x34, 0x9f,
	0x80, 0xec, 0x5b, 0xfb, 0xd6, 0xa8, 0x3d, 0x7e, 0x4a, 0x36, 0x65, 0x4a, 0xd6, 0x42, 0x92, 0x17,
	0xa9, 0x0b, 0x67, 0x05, 0xdc, 0xab, 0x02, 0x8d, 0xe2, 0x70, 0xc5, 0x86, 0x05, 0xba, 0xb5, 0x14,
	0x59, 0x82, 0xa7, 0xfa, 0x0d, 0x13, 0xf6, 0xe4, 0x2f, 0xc3, 0x52, 0xf0, 0x14, 0xed, 0x84, 0xcb,
	0x06, 0xfb, 0x87, 0x85, 0xf0, 0x6a, 0x6e, 0xf8, 0x0a, 0x6d, 0x45, 0xcc, 0x85, 0x48, 0xf5, 0xad,
	0xfd, 0xe6, 0xa8, 0x3d, 0xa6, 0xff, 0xa6, 0x66, 0xf2, 0xd2, 0x80, 0x9e, 0xc6, 0x5a, 0xce, 0x68,
	0x19, 0x01, 0x0f, 0x10, 0x8a, 0xd9, 0x14, 0x54, 0xc2, 0x38, 0xe4, 0xc5, 0x36, 0x47, 0x2d, 0xba,
	0x60, 0xc1, 0x36, 0xda, 0xe6, 0x51, 0xaa, 0x34, 0x48, 0xd5, 0x6f, 0x9a, 0xdb, 0xea, 0xdf, 0x7e,
	0x88, 0xda, 0x0b, 0x90, 0xb8, 0x8b, 0x9a, 0x21, 0xcc, 0x4c, 0x9f, 0x5a, 0x34, 0x3f, 0xe2, 0x1e,
	0xba, 0x99, 0xb1, 0x28, 0x05, 0x43, 0x62, 0x8b, 0x16, 0x3f, 0x47, 0x8d, 0x07, 0x96, 0x4d, 0x51,
	0xa7, 0xc6, 0x0e, 0x7e, 0x84, 0xb6, 0x4b, 0xe2, 0xe7, 0x75, 0xdf, 0x23, 0x46, 0x46, 0xb9, 0xb8,
	0xaa, 0x62, 0x4f, 0x8a, 0xe0, 0x6f, 0xdc, 0x2b, 0xe0, 0x9a, 0x82, 0x47, 0x2b, 0xa7, 0xe1, 0x37,
	0x0b, 0x75, 0xdf, 0x0b, 0x19, 0x46, 0x82, 0x5d, 0x56, 0x6a, 0xa2, 0x35, 0x2e, 0x8f, 0x36, 0x73,
	0x59, 0xf7, 0xff, 0x8f, 0x38, 0x1b, 0x7e, 0x69, 0xa2, 0xee, 0xb3, 0x4b, 0x88, 0x75, 0xa0, 0x67,
	0x55, 0x7d, 0x19, 0xda, 0x35, 0x9a, 0x0d, 0xca, 0x8b, 0xda, 0xb8, 0x1c, 0x6f, 0x2e, 0xb7, 0x0e,
	0x67, 0x54, 0x33, 0x37, 0xce, 0x47, 0x65, 0x27, 0x5c, 0x35, 0xe2, 0x4f, 0x16, 0xba, 0xbd, 0x34,
	0x2c, 0x8c, 0x73, 0x91, 0xc6, 0x7a, 0x71, 0x68, 0x4e, 0xff, 0x20, 0x78, 0xa9, 0x88, 0xc7, 0x05,
	0x9a, 0x19, 0x9b, 0xbd, 0x70, 0xad, 0xdd, 0x7e, 0x87, 0x76, 0xd6, 0xa4, 0x5b, 0xeb, 0x8e, 0xb5,
	0xb1, 0x3b, 0x8d, 0x5a, 0x77, 0x26, 0x68, 0x6f, 0x7d, 0x12, 0xf8, 0x35, 0xea, 0xd6, 0x2a, 0xfd,
	0x2d, 0x95, 0x76, 0xd4, 0x12, 0xa4, 0x3a, 0x7e, 0xfb, 0xf9, 0xfb, 0xc0, 0xfa, 0xf0, 0x7c, 0xe3,
	0xf2, 0x4e, 0x42, 0xbf, 0xb6, 0xc0, 0x57, 0x59, 0xac, 0x56, 0xba, 0xbb, 0x65, 0x16, 0xeb, 0xfd,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x51, 0xda, 0xff, 0x11, 0x06, 0x00, 0x00,
}

func (this *TrafficTargetSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficTargetSelector)
	if !ok {
		that2, ok := that.(TrafficTargetSelector)
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
	if !this.KubeServiceMatcher.Equal(that1.KubeServiceMatcher) {
		return false
	}
	if !this.KubeServiceRefs.Equal(that1.KubeServiceRefs) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficTargetSelector_KubeServiceMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficTargetSelector_KubeServiceMatcher)
	if !ok {
		that2, ok := that.(TrafficTargetSelector_KubeServiceMatcher)
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
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TrafficTargetSelector_KubeServiceRefs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TrafficTargetSelector_KubeServiceRefs)
	if !ok {
		that2, ok := that.(TrafficTargetSelector_KubeServiceRefs)
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
	if len(this.Services) != len(that1.Services) {
		return false
	}
	for i := range this.Services {
		if !this.Services[i].Equal(that1.Services[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *WorkloadSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSelector)
	if !ok {
		that2, ok := that.(WorkloadSelector)
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
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IdentitySelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector)
	if !ok {
		that2, ok := that.(IdentitySelector)
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
	if !this.KubeIdentityMatcher.Equal(that1.KubeIdentityMatcher) {
		return false
	}
	if !this.KubeServiceAccountRefs.Equal(that1.KubeServiceAccountRefs) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IdentitySelector_KubeIdentityMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector_KubeIdentityMatcher)
	if !ok {
		that2, ok := that.(IdentitySelector_KubeIdentityMatcher)
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
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IdentitySelector_KubeServiceAccountRefs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector_KubeServiceAccountRefs)
	if !ok {
		that2, ok := that.(IdentitySelector_KubeServiceAccountRefs)
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
	if len(this.ServiceAccounts) != len(that1.ServiceAccounts) {
		return false
	}
	for i := range this.ServiceAccounts {
		if !this.ServiceAccounts[i].Equal(that1.ServiceAccounts[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
