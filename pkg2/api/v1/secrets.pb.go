// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/secrets.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

type TlsSecret struct {
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	RootCert             string        `protobuf:"bytes,1,opt,name=root_cert,json=root-cert.pem,proto3" json:"root_cert,omitempty"`
	CertChain            string        `protobuf:"bytes,2,opt,name=cert_chain,json=cert-chain.pem,proto3" json:"cert_chain,omitempty"`
	CaCert               string        `protobuf:"bytes,3,opt,name=ca_cert,json=ca-cert.pem,proto3" json:"ca_cert,omitempty"`
	CaKey                string        `protobuf:"bytes,4,opt,name=ca_key,json=ca-key.pem,proto3" json:"ca_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TlsSecret) Reset()         { *m = TlsSecret{} }
func (m *TlsSecret) String() string { return proto.CompactTextString(m) }
func (*TlsSecret) ProtoMessage()    {}
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c765e2d1cb5e906, []int{0}
}
func (m *TlsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSecret.Unmarshal(m, b)
}
func (m *TlsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSecret.Marshal(b, m, deterministic)
}
func (m *TlsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSecret.Merge(m, src)
}
func (m *TlsSecret) XXX_Size() int {
	return xxx_messageInfo_TlsSecret.Size(m)
}
func (m *TlsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSecret proto.InternalMessageInfo

func (m *TlsSecret) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *TlsSecret) GetRootCert() string {
	if m != nil {
		return m.RootCert
	}
	return ""
}

func (m *TlsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *TlsSecret) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *TlsSecret) GetCaKey() string {
	if m != nil {
		return m.CaKey
	}
	return ""
}

func init() {
	proto.RegisterType((*TlsSecret)(nil), "supergloo.solo.io.TlsSecret")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/secrets.proto", fileDescriptor_1c765e2d1cb5e906)
}

var fileDescriptor_1c765e2d1cb5e906 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xfd, 0xd2, 0xcf, 0x2a, 0xd4, 0x05, 0x24, 0x22, 0x84, 0xa2, 0x08, 0x95, 0xa8, 0x53, 0x97,
	0xd8, 0x6a, 0x59, 0x10, 0x63, 0x77, 0x96, 0xc2, 0xc4, 0x52, 0xb9, 0xae, 0x71, 0xad, 0xa4, 0xbd,
	0x96, 0xed, 0x80, 0xba, 0xf6, 0x69, 0x78, 0x14, 0x9e, 0x82, 0x01, 0x89, 0x07, 0xe8, 0x1b, 0x20,
	0xbb, 0x49, 0x59, 0x00, 0x89, 0xc9, 0xf7, 0x9e, 0x73, 0xee, 0xf1, 0xfd, 0xc1, 0x43, 0xa9, 0xdc,
	0xa2, 0x9a, 0x11, 0x0e, 0x4b, 0x6a, 0xa1, 0x84, 0x5c, 0x01, 0xb5, 0x95, 0x16, 0x46, 0x96, 0x00,
	0x94, 0x69, 0x45, 0x9f, 0x86, 0xd4, 0x0a, 0x6e, 0x84, 0xb3, 0x44, 0x1b, 0x70, 0x10, 0x9f, 0xee,
	0x79, 0xe2, 0x2b, 0x88, 0x82, 0xf4, 0x4c, 0x82, 0x84, 0xc0, 0x52, 0x1f, 0xed, 0x84, 0x69, 0x4f,
	0x02, 0xc8, 0x52, 0xd0, 0x90, 0xcd, 0xaa, 0x47, 0x3a, 0xaf, 0x0c, 0x73, 0x0a, 0x56, 0x3f, 0xf1,
	0xcf, 0x86, 0x69, 0x2d, 0x4c, 0xfd, 0x51, 0xfa, 0x6d, 0x6f, 0xfe, 0x2d, 0x94, 0x6b, 0x5a, 0x5b,
	0x0a, 0xc7, 0xe6, 0xcc, 0xb1, 0x3f, 0x94, 0x34, 0xf9, 0xae, 0xa4, 0xff, 0x11, 0xe1, 0xce, 0x7d,
	0x69, 0xef, 0xc2, 0x8c, 0xf1, 0x35, 0x3e, 0x6c, 0x2c, 0x13, 0x91, 0x45, 0x83, 0xee, 0xe8, 0x9c,
	0x70, 0x30, 0xa2, 0x19, 0x95, 0xdc, 0xd6, 0xec, 0x18, 0xbd, 0xbe, 0x5d, 0xfe, 0x9b, 0xec, 0xd5,
	0x71, 0x86, 0x3b, 0x06, 0xc0, 0x4d, 0xb9, 0x30, 0x2e, 0x89, 0xb2, 0x68, 0xd0, 0x99, 0x1c, 0x7b,
	0x20, 0xf7, 0x00, 0xd1, 0x62, 0x19, 0xf7, 0x31, 0xf6, 0xf1, 0x94, 0x2f, 0x98, 0x5a, 0x25, 0xad,
	0x20, 0x39, 0xf1, 0x48, 0x1e, 0x90, 0xa0, 0xb9, 0xc0, 0x07, 0x9c, 0xed, 0x3c, 0xfe, 0x07, 0x41,
	0x97, 0xb3, 0x2f, 0x87, 0x14, 0xb7, 0x39, 0x9b, 0x16, 0x62, 0x9d, 0xa0, 0x40, 0x62, 0xce, 0xf2,
	0x42, 0xac, 0x3d, 0x77, 0xd3, 0xdb, 0x6c, 0x11, 0xc2, 0x2d, 0x67, 0x37, 0x5b, 0x74, 0x14, 0x63,
	0x57, 0xda, 0xfa, 0x6c, 0x9b, 0x2d, 0x6a, 0x65, 0xd1, 0x98, 0xbc, 0xbc, 0xf7, 0xa2, 0x87, 0xc1,
	0xaf, 0xf7, 0xd6, 0x85, 0x1c, 0xd5, 0x6b, 0x9a, 0xb5, 0xc3, 0x7a, 0xae, 0x3e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0x17, 0x81, 0x1d, 0x22, 0x02, 0x00, 0x00,
}

func (this *TlsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TlsSecret)
	if !ok {
		that2, ok := that.(TlsSecret)
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
	if this.RootCert != that1.RootCert {
		return false
	}
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.CaCert != that1.CaCert {
		return false
	}
	if this.CaKey != that1.CaKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
