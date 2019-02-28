// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/external/kubernetes/core/v1/kube_types.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//Intermediary proto representation of a kubernetes pod.
//Used to integrate solo-kit with kubernetes API
type Pod struct {
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// the kubernetes pod spec as an inline json string
	Spec string `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	// the kubernetes pod status as an inline json string
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_c148cb2834f27f59, []int{0}
}
func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (m *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(m, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Pod) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *Pod) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Pod)(nil), "core.kubernetes.io.Pod")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/external/kubernetes/core/v1/kube_types.proto", fileDescriptor_c148cb2834f27f59)
}

var fileDescriptor_c148cb2834f27f59 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x2f,
	0x2e, 0x2d, 0x48, 0x2d, 0x4a, 0xcf, 0xc9, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xad, 0x28,
	0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x2e, 0x4d, 0x4a, 0x2d, 0xca, 0x4b, 0x2d, 0x49, 0x2d,
	0xd6, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x33, 0x04, 0x0b, 0xc5, 0x97, 0x54, 0x16, 0xa4, 0x16,
	0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x81, 0x64, 0xf4, 0x10, 0x2a, 0xf5, 0x32, 0xf3,
	0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xd2, 0xfa, 0x20, 0x16, 0x44, 0xa5, 0x94, 0x21, 0x36,
	0x8b, 0x41, 0x74, 0x76, 0x66, 0x09, 0xd8, 0xde, 0x32, 0x43, 0xfd, 0xdc, 0xd4, 0x92, 0xc4, 0x94,
	0xc4, 0x92, 0x44, 0x12, 0xb4, 0xc0, 0xf8, 0x10, 0x2d, 0x4a, 0x0d, 0x8c, 0x5c, 0xcc, 0x01, 0xf9,
	0x29, 0x42, 0x16, 0x5c, 0x1c, 0x30, 0xc3, 0x24, 0x52, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xc4, 0xf4,
	0xc0, 0x4e, 0x05, 0xa9, 0xd7, 0xcb, 0xcc, 0xd7, 0xf3, 0x85, 0xca, 0x3a, 0xb1, 0x9c, 0xb8, 0x27,
	0xcf, 0x10, 0x04, 0x57, 0x2d, 0x24, 0xc4, 0xc5, 0x52, 0x5c, 0x90, 0x9a, 0x2c, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b,
	0x30, 0x81, 0x45, 0xa1, 0x3c, 0x2b, 0xe1, 0xa6, 0x8f, 0x2c, 0xac, 0x5c, 0xcc, 0x05, 0xf9, 0x29,
	0x4d, 0x1f, 0x59, 0xd8, 0x84, 0x58, 0x0a, 0xf2, 0x53, 0x8a, 0x9d, 0xdc, 0x56, 0x3c, 0x92, 0x63,
	0x8c, 0x72, 0xc0, 0x1b, 0xce, 0x05, 0xd9, 0xe9, 0x84, 0xc2, 0x3a, 0x89, 0x0d, 0xec, 0x23, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x5f, 0x36, 0x43, 0xb2, 0x01, 0x00, 0x00,
}

func (this *Pod) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pod)
	if !ok {
		that2, ok := that.(Pod)
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
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.Spec != that1.Spec {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
