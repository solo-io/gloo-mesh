// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/certificates/v1/ca_options.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// Configuration for generating a self-signed root certificate.
// Uses the X.509 format, RFC5280.
type CommonCertOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of days before root cert expires. Defaults to 365.
	TtlDays uint32 `protobuf:"varint,1,opt,name=ttl_days,json=ttlDays,proto3" json:"ttl_days,omitempty"`
	// Size in bytes of the root cert's private key. Defaults to 4096.
	RsaKeySizeBytes uint32 `protobuf:"varint,2,opt,name=rsa_key_size_bytes,json=rsaKeySizeBytes,proto3" json:"rsa_key_size_bytes,omitempty"`
	// Root cert organization name. Defaults to "gloo-mesh".
	OrgName string `protobuf:"bytes,3,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	// The ratio of cert lifetime to refresh a cert. For example, at 0.10 and 1 hour TTL,
	// we would refresh 6 minutes before expiration
	SecretRotationGracePeriodRatio float32 `protobuf:"fixed32,4,opt,name=secret_rotation_grace_period_ratio,json=secretRotationGracePeriodRatio,proto3" json:"secret_rotation_grace_period_ratio,omitempty"`
}

func (x *CommonCertOptions) Reset() {
	*x = CommonCertOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonCertOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonCertOptions) ProtoMessage() {}

func (x *CommonCertOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonCertOptions.ProtoReflect.Descriptor instead.
func (*CommonCertOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP(), []int{0}
}

func (x *CommonCertOptions) GetTtlDays() uint32 {
	if x != nil {
		return x.TtlDays
	}
	return 0
}

func (x *CommonCertOptions) GetRsaKeySizeBytes() uint32 {
	if x != nil {
		return x.RsaKeySizeBytes
	}
	return 0
}

func (x *CommonCertOptions) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *CommonCertOptions) GetSecretRotationGracePeriodRatio() float32 {
	if x != nil {
		return x.SecretRotationGracePeriodRatio
	}
	return 0
}

// Specify parameters for configuring the root certificate authority for a VirtualMesh.
type IntermediateCertificateAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the source of the Root CA data which Gloo Mesh will use for the VirtualMesh.
	//
	// Types that are assignable to CaSource:
	//	*IntermediateCertificateAuthority_Vault
	CaSource isIntermediateCertificateAuthority_CaSource `protobuf_oneof:"ca_source"`
}

func (x *IntermediateCertificateAuthority) Reset() {
	*x = IntermediateCertificateAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntermediateCertificateAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntermediateCertificateAuthority) ProtoMessage() {}

func (x *IntermediateCertificateAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntermediateCertificateAuthority.ProtoReflect.Descriptor instead.
func (*IntermediateCertificateAuthority) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP(), []int{1}
}

func (m *IntermediateCertificateAuthority) GetCaSource() isIntermediateCertificateAuthority_CaSource {
	if m != nil {
		return m.CaSource
	}
	return nil
}

func (x *IntermediateCertificateAuthority) GetVault() *VaultCA {
	if x, ok := x.GetCaSource().(*IntermediateCertificateAuthority_Vault); ok {
		return x.Vault
	}
	return nil
}

type isIntermediateCertificateAuthority_CaSource interface {
	isIntermediateCertificateAuthority_CaSource()
}

type IntermediateCertificateAuthority_Vault struct {
	// Use vault as the intermediate CA source
	Vault *VaultCA `protobuf:"bytes,1,opt,name=vault,proto3,oneof"`
}

func (*IntermediateCertificateAuthority_Vault) isIntermediateCertificateAuthority_CaSource() {}

var File_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x74, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x74, 0x74, 0x6c, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x73, 0x61,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x4a, 0x0a, 0x22, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61,
	0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x70, 0x0a,
	0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x41, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x4c, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_goTypes = []interface{}{
	(*CommonCertOptions)(nil),                // 0: certificates.mesh.gloo.solo.io.CommonCertOptions
	(*IntermediateCertificateAuthority)(nil), // 1: certificates.mesh.gloo.solo.io.IntermediateCertificateAuthority
	(*VaultCA)(nil),                          // 2: certificates.mesh.gloo.solo.io.VaultCA
}
var file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_depIdxs = []int32{
	2, // 0: certificates.mesh.gloo.solo.io.IntermediateCertificateAuthority.vault:type_name -> certificates.mesh.gloo.solo.io.VaultCA
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_vault_ca_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonCertOptions); i {
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
		file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntermediateCertificateAuthority); i {
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
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*IntermediateCertificateAuthority_Vault)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_certificates_v1_ca_options_proto_depIdxs = nil
}
