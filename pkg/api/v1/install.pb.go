// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/install.proto

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
//Installs represent a desired installation of a supported mesh.
//Supergloo watches for installs and synchronizes the managed installations
//with the desired configuration in the install object.
//
//Updating the configuration of an install object will cause supergloo to
//modify the corresponding mesh.
type Install struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// disables this install
	// setting this to true will cause supergloo to
	// not install this mesh, or uninstall an active install
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// The type of object the install handles
	// Currently support types are mesh, and ingress
	//
	// Types that are valid to be assigned to InstallType:
	//	*Install_Mesh
	//	*Install_Ingress
	InstallType isInstall_InstallType `protobuf_oneof:"install_type"`
	// which namespace to install to
	InstallationNamespace string `protobuf:"bytes,4,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// gzipped inline string containing the applied manifest
	// read-only, set by the server after successful installation.
	// TODO (ilackarms): make sure this is not too large for etcd (value size limit 1.5mb)
	InstalledManifest    string   `protobuf:"bytes,5,opt,name=installed_manifest,json=installedManifest,proto3" json:"installed_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{0}
}
func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (m *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(m, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

type isInstall_InstallType interface {
	isInstall_InstallType()
	Equal(interface{}) bool
}

type Install_Mesh struct {
	Mesh *MeshInstall `protobuf:"bytes,2,opt,name=mesh,proto3,oneof"`
}
type Install_Ingress struct {
	Ingress *MeshIngressInstall `protobuf:"bytes,3,opt,name=ingress,proto3,oneof"`
}

func (*Install_Mesh) isInstall_InstallType()    {}
func (*Install_Ingress) isInstall_InstallType() {}

func (m *Install) GetInstallType() isInstall_InstallType {
	if m != nil {
		return m.InstallType
	}
	return nil
}

func (m *Install) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Install) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Install) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Install) GetMesh() *MeshInstall {
	if x, ok := m.GetInstallType().(*Install_Mesh); ok {
		return x.Mesh
	}
	return nil
}

func (m *Install) GetIngress() *MeshIngressInstall {
	if x, ok := m.GetInstallType().(*Install_Ingress); ok {
		return x.Ingress
	}
	return nil
}

func (m *Install) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *Install) GetInstalledManifest() string {
	if m != nil {
		return m.InstalledManifest
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Install) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Install_OneofMarshaler, _Install_OneofUnmarshaler, _Install_OneofSizer, []interface{}{
		(*Install_Mesh)(nil),
		(*Install_Ingress)(nil),
	}
}

func _Install_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Install)
	// install_type
	switch x := m.InstallType.(type) {
	case *Install_Mesh:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mesh); err != nil {
			return err
		}
	case *Install_Ingress:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ingress); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Install.InstallType has unexpected type %T", x)
	}
	return nil
}

func _Install_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Install)
	switch tag {
	case 2: // install_type.mesh
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MeshInstall)
		err := b.DecodeMessage(msg)
		m.InstallType = &Install_Mesh{msg}
		return true, err
	case 3: // install_type.ingress
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MeshIngressInstall)
		err := b.DecodeMessage(msg)
		m.InstallType = &Install_Ingress{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Install_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Install)
	// install_type
	switch x := m.InstallType.(type) {
	case *Install_Mesh:
		s := proto.Size(x.Mesh)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Install_Ingress:
		s := proto.Size(x.Ingress)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Installation options for Istio
type IstioInstall struct {
	// which version of the istio helm chart to install
	// ignored if using custom helm chart
	IstioVersion string `protobuf:"bytes,2,opt,name=istio_version,json=istioVersion,proto3" json:"istio_version,omitempty"`
	// enable auto injection of pods
	EnableAutoInject bool `protobuf:"varint,3,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// enable mutual tls between pods
	EnableMtls bool `protobuf:"varint,4,opt,name=enable_mtls,json=enableMtls,proto3" json:"enable_mtls,omitempty"`
	// optional. set to use a custom root ca
	// to issue certificates for mtls
	// ignored if mtls is disabled
	CustomRootCert *core.ResourceRef `protobuf:"bytes,9,opt,name=custom_root_cert,json=customRootCert,proto3" json:"custom_root_cert,omitempty"`
	// install grafana with istio
	InstallGrafana bool `protobuf:"varint,6,opt,name=install_grafana,json=installGrafana,proto3" json:"install_grafana,omitempty"`
	// install prometheus with istio
	InstallPrometheus bool `protobuf:"varint,7,opt,name=install_prometheus,json=installPrometheus,proto3" json:"install_prometheus,omitempty"`
	// install jaeger with istio
	InstallJaeger        bool     `protobuf:"varint,8,opt,name=install_jaeger,json=installJaeger,proto3" json:"install_jaeger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioInstall) Reset()         { *m = IstioInstall{} }
func (m *IstioInstall) String() string { return proto.CompactTextString(m) }
func (*IstioInstall) ProtoMessage()    {}
func (*IstioInstall) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{1}
}
func (m *IstioInstall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioInstall.Unmarshal(m, b)
}
func (m *IstioInstall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioInstall.Marshal(b, m, deterministic)
}
func (m *IstioInstall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioInstall.Merge(m, src)
}
func (m *IstioInstall) XXX_Size() int {
	return xxx_messageInfo_IstioInstall.Size(m)
}
func (m *IstioInstall) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioInstall.DiscardUnknown(m)
}

var xxx_messageInfo_IstioInstall proto.InternalMessageInfo

func (m *IstioInstall) GetIstioVersion() string {
	if m != nil {
		return m.IstioVersion
	}
	return ""
}

func (m *IstioInstall) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *IstioInstall) GetEnableMtls() bool {
	if m != nil {
		return m.EnableMtls
	}
	return false
}

func (m *IstioInstall) GetCustomRootCert() *core.ResourceRef {
	if m != nil {
		return m.CustomRootCert
	}
	return nil
}

func (m *IstioInstall) GetInstallGrafana() bool {
	if m != nil {
		return m.InstallGrafana
	}
	return false
}

func (m *IstioInstall) GetInstallPrometheus() bool {
	if m != nil {
		return m.InstallPrometheus
	}
	return false
}

func (m *IstioInstall) GetInstallJaeger() bool {
	if m != nil {
		return m.InstallJaeger
	}
	return false
}

//
//Generic container for mesh installs handled by supergloo
//
//Holds all configuration shared between different mesh types
type MeshInstall struct {
	// The type of mesh to install
	// currently only istio is supported
	//
	// Types that are valid to be assigned to MeshInstallType:
	//	*MeshInstall_IstioMesh
	MeshInstallType isMeshInstall_MeshInstallType `protobuf_oneof:"mesh_install_type"`
	// reference to the Mesh crd that was created from this install
	// read-only, set by the server after successful installation.
	InstalledMesh        *core.ResourceRef `protobuf:"bytes,6,opt,name=installed_mesh,json=installedMesh,proto3" json:"installed_mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MeshInstall) Reset()         { *m = MeshInstall{} }
func (m *MeshInstall) String() string { return proto.CompactTextString(m) }
func (*MeshInstall) ProtoMessage()    {}
func (*MeshInstall) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{2}
}
func (m *MeshInstall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshInstall.Unmarshal(m, b)
}
func (m *MeshInstall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshInstall.Marshal(b, m, deterministic)
}
func (m *MeshInstall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshInstall.Merge(m, src)
}
func (m *MeshInstall) XXX_Size() int {
	return xxx_messageInfo_MeshInstall.Size(m)
}
func (m *MeshInstall) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshInstall.DiscardUnknown(m)
}

var xxx_messageInfo_MeshInstall proto.InternalMessageInfo

type isMeshInstall_MeshInstallType interface {
	isMeshInstall_MeshInstallType()
	Equal(interface{}) bool
}

type MeshInstall_IstioMesh struct {
	IstioMesh *IstioInstall `protobuf:"bytes,2,opt,name=istio_mesh,json=istioMesh,proto3,oneof"`
}

func (*MeshInstall_IstioMesh) isMeshInstall_MeshInstallType() {}

func (m *MeshInstall) GetMeshInstallType() isMeshInstall_MeshInstallType {
	if m != nil {
		return m.MeshInstallType
	}
	return nil
}

func (m *MeshInstall) GetIstioMesh() *IstioInstall {
	if x, ok := m.GetMeshInstallType().(*MeshInstall_IstioMesh); ok {
		return x.IstioMesh
	}
	return nil
}

func (m *MeshInstall) GetInstalledMesh() *core.ResourceRef {
	if m != nil {
		return m.InstalledMesh
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MeshInstall) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MeshInstall_OneofMarshaler, _MeshInstall_OneofUnmarshaler, _MeshInstall_OneofSizer, []interface{}{
		(*MeshInstall_IstioMesh)(nil),
	}
}

func _MeshInstall_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MeshInstall)
	// mesh_install_type
	switch x := m.MeshInstallType.(type) {
	case *MeshInstall_IstioMesh:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IstioMesh); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MeshInstall.MeshInstallType has unexpected type %T", x)
	}
	return nil
}

func _MeshInstall_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MeshInstall)
	switch tag {
	case 2: // mesh_install_type.istio_mesh
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IstioInstall)
		err := b.DecodeMessage(msg)
		m.MeshInstallType = &MeshInstall_IstioMesh{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MeshInstall_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MeshInstall)
	// mesh_install_type
	switch x := m.MeshInstallType.(type) {
	case *MeshInstall_IstioMesh:
		s := proto.Size(x.IstioMesh)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//
//Generic container for ingress installs handled by supergloo
//
//Holds all configuration shared between different ingress types
type MeshIngressInstall struct {
	// The type of mesh to install
	// currently only gloo is supported
	//
	// Types that are valid to be assigned to IngressInstallType:
	//	*MeshIngressInstall_Gloo
	IngressInstallType isMeshIngressInstall_IngressInstallType `protobuf_oneof:"ingress_install_type"`
	// reference to the Ingress crd that was created from this install
	// read-only, set by the server after successful installation.
	InstalledIngress     *core.ResourceRef `protobuf:"bytes,3,opt,name=installed_ingress,json=installedIngress,proto3" json:"installed_ingress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MeshIngressInstall) Reset()         { *m = MeshIngressInstall{} }
func (m *MeshIngressInstall) String() string { return proto.CompactTextString(m) }
func (*MeshIngressInstall) ProtoMessage()    {}
func (*MeshIngressInstall) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{3}
}
func (m *MeshIngressInstall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshIngressInstall.Unmarshal(m, b)
}
func (m *MeshIngressInstall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshIngressInstall.Marshal(b, m, deterministic)
}
func (m *MeshIngressInstall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshIngressInstall.Merge(m, src)
}
func (m *MeshIngressInstall) XXX_Size() int {
	return xxx_messageInfo_MeshIngressInstall.Size(m)
}
func (m *MeshIngressInstall) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshIngressInstall.DiscardUnknown(m)
}

var xxx_messageInfo_MeshIngressInstall proto.InternalMessageInfo

type isMeshIngressInstall_IngressInstallType interface {
	isMeshIngressInstall_IngressInstallType()
	Equal(interface{}) bool
}

type MeshIngressInstall_Gloo struct {
	Gloo *GlooInstall `protobuf:"bytes,1,opt,name=gloo,proto3,oneof"`
}

func (*MeshIngressInstall_Gloo) isMeshIngressInstall_IngressInstallType() {}

func (m *MeshIngressInstall) GetIngressInstallType() isMeshIngressInstall_IngressInstallType {
	if m != nil {
		return m.IngressInstallType
	}
	return nil
}

func (m *MeshIngressInstall) GetGloo() *GlooInstall {
	if x, ok := m.GetIngressInstallType().(*MeshIngressInstall_Gloo); ok {
		return x.Gloo
	}
	return nil
}

func (m *MeshIngressInstall) GetInstalledIngress() *core.ResourceRef {
	if m != nil {
		return m.InstalledIngress
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MeshIngressInstall) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MeshIngressInstall_OneofMarshaler, _MeshIngressInstall_OneofUnmarshaler, _MeshIngressInstall_OneofSizer, []interface{}{
		(*MeshIngressInstall_Gloo)(nil),
	}
}

func _MeshIngressInstall_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MeshIngressInstall)
	// ingress_install_type
	switch x := m.IngressInstallType.(type) {
	case *MeshIngressInstall_Gloo:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gloo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MeshIngressInstall.IngressInstallType has unexpected type %T", x)
	}
	return nil
}

func _MeshIngressInstall_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MeshIngressInstall)
	switch tag {
	case 1: // ingress_install_type.gloo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GlooInstall)
		err := b.DecodeMessage(msg)
		m.IngressInstallType = &MeshIngressInstall_Gloo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MeshIngressInstall_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MeshIngressInstall)
	// ingress_install_type
	switch x := m.IngressInstallType.(type) {
	case *MeshIngressInstall_Gloo:
		s := proto.Size(x.Gloo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Installation options for Gloo Ingress
type GlooInstall struct {
	// which version of the gloo helm chart to install
	// ignored if using custom helm chart
	GlooVersion string `protobuf:"bytes,2,opt,name=gloo_version,json=glooVersion,proto3" json:"gloo_version,omitempty"`
	// reference to the Mesh(s) that this ingress is acting upon
	Meshes               []*core.ResourceRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GlooInstall) Reset()         { *m = GlooInstall{} }
func (m *GlooInstall) String() string { return proto.CompactTextString(m) }
func (*GlooInstall) ProtoMessage()    {}
func (*GlooInstall) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{4}
}
func (m *GlooInstall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlooInstall.Unmarshal(m, b)
}
func (m *GlooInstall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlooInstall.Marshal(b, m, deterministic)
}
func (m *GlooInstall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlooInstall.Merge(m, src)
}
func (m *GlooInstall) XXX_Size() int {
	return xxx_messageInfo_GlooInstall.Size(m)
}
func (m *GlooInstall) XXX_DiscardUnknown() {
	xxx_messageInfo_GlooInstall.DiscardUnknown(m)
}

var xxx_messageInfo_GlooInstall proto.InternalMessageInfo

func (m *GlooInstall) GetGlooVersion() string {
	if m != nil {
		return m.GlooVersion
	}
	return ""
}

func (m *GlooInstall) GetMeshes() []*core.ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func init() {
	proto.RegisterType((*Install)(nil), "supergloo.solo.io.Install")
	proto.RegisterType((*IstioInstall)(nil), "supergloo.solo.io.IstioInstall")
	proto.RegisterType((*MeshInstall)(nil), "supergloo.solo.io.MeshInstall")
	proto.RegisterType((*MeshIngressInstall)(nil), "supergloo.solo.io.MeshIngressInstall")
	proto.RegisterType((*GlooInstall)(nil), "supergloo.solo.io.GlooInstall")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/install.proto", fileDescriptor_7b058d98c63047dc)
}

var fileDescriptor_7b058d98c63047dc = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x24, 0x37, 0x24, 0x27, 0x81, 0x4b, 0xe6, 0x02, 0x32, 0x2c, 0x08, 0x37, 0x57, 0x08,
	0x16, 0xe0, 0x88, 0xdb, 0x56, 0xaa, 0x58, 0x41, 0x90, 0x0a, 0x54, 0x4a, 0x55, 0xb9, 0x52, 0x17,
	0xdd, 0x58, 0x83, 0x73, 0xe2, 0x0c, 0xd8, 0x9e, 0x68, 0x66, 0x8c, 0xd4, 0x2d, 0x0f, 0x53, 0x55,
	0x5d, 0xf4, 0x31, 0xaa, 0x3e, 0x05, 0x8b, 0xbe, 0x01, 0x7d, 0x82, 0x6a, 0xc6, 0xe3, 0x90, 0x14,
	0x4a, 0xe9, 0xca, 0x9e, 0xf3, 0x9d, 0xef, 0xfc, 0x7c, 0xe7, 0xcc, 0xc0, 0x5e, 0xc4, 0xd4, 0x30,
	0x3b, 0xf3, 0x42, 0x9e, 0x74, 0x24, 0x8f, 0xf9, 0x2e, 0xe3, 0x1d, 0x99, 0x8d, 0x50, 0x44, 0x31,
	0xe7, 0x1d, 0x3a, 0x62, 0x9d, 0xcb, 0xbd, 0x0e, 0x4b, 0xa5, 0xa2, 0x71, 0xec, 0x8d, 0x04, 0x57,
	0x9c, 0x34, 0xc7, 0xb8, 0xa7, 0x19, 0x1e, 0xe3, 0x6b, 0x4b, 0x11, 0x8f, 0xb8, 0x41, 0x3b, 0xfa,
	0x2f, 0x77, 0x5c, 0xbb, 0x37, 0xb6, 0xfe, 0x5e, 0x30, 0x55, 0x84, 0x4e, 0x50, 0xd1, 0x3e, 0x55,
	0xd4, 0x52, 0x3a, 0x8f, 0xa0, 0x48, 0x45, 0x55, 0x26, 0x2d, 0x61, 0xe7, 0x11, 0x04, 0x81, 0x83,
	0x3f, 0xa8, 0xa8, 0x38, 0xe7, 0x94, 0xf6, 0xe7, 0x12, 0xcc, 0x9d, 0xe6, 0xfd, 0x93, 0x63, 0xa8,
	0xe4, 0xc9, 0xdd, 0xfe, 0x86, 0xb3, 0x5d, 0xff, 0x7f, 0xc9, 0x0b, 0xb9, 0xc0, 0x42, 0x05, 0xef,
	0x8d, 0xc1, 0xba, 0xab, 0x5f, 0xaf, 0x5b, 0x33, 0xdf, 0xaf, 0x5b, 0x4d, 0x85, 0x52, 0xf5, 0xd9,
	0x60, 0xb0, 0xdf, 0x66, 0x51, 0xca, 0x05, 0xb6, 0x7d, 0x4b, 0x27, 0xcf, 0xa1, 0x5a, 0x34, 0xee,
	0xa2, 0x09, 0xb5, 0x32, 0x1d, 0xaa, 0x67, 0xd1, 0x6e, 0x59, 0x07, 0xf3, 0xc7, 0xde, 0x64, 0x0d,
	0xaa, 0x7d, 0x26, 0xe9, 0x59, 0x8c, 0x7d, 0xd7, 0xd9, 0x70, 0xb6, 0xab, 0xfe, 0xf8, 0x4c, 0x9e,
	0x42, 0x39, 0x41, 0x39, 0x74, 0x67, 0x4d, 0xc4, 0x75, 0xef, 0xce, 0x9c, 0xbc, 0x1e, 0xca, 0xa1,
	0x6d, 0xe6, 0x64, 0xc6, 0x37, 0xde, 0xe4, 0x10, 0xe6, 0x58, 0x1a, 0x09, 0x94, 0xd2, 0x2d, 0x19,
	0xe2, 0xe6, 0x2f, 0x89, 0xc6, 0xeb, 0x96, 0x5f, 0xf0, 0xc8, 0x33, 0x58, 0xb1, 0x2b, 0x42, 0x15,
	0xe3, 0x69, 0x90, 0xd2, 0x04, 0xe5, 0x88, 0x86, 0xe8, 0x96, 0x37, 0x9c, 0xed, 0x9a, 0xbf, 0x3c,
	0x89, 0xbe, 0x2a, 0x40, 0xb2, 0x0b, 0xc4, 0x02, 0xd8, 0x0f, 0x12, 0x9a, 0xb2, 0x01, 0x4a, 0xe5,
	0xfe, 0x65, 0x28, 0xcd, 0x31, 0xd2, 0xb3, 0xc0, 0xfe, 0xf2, 0xd5, 0x4d, 0xb9, 0x04, 0x0e, 0xbb,
	0xba, 0x29, 0x03, 0xa9, 0x5a, 0x5c, 0x76, 0x17, 0xa0, 0x61, 0xff, 0x03, 0xf5, 0x7e, 0x84, 0xed,
	0x2f, 0xb3, 0xd0, 0x38, 0x95, 0x8a, 0xf1, 0x62, 0x6a, 0xff, 0xc1, 0x3c, 0xd3, 0xe7, 0xe0, 0x12,
	0x85, 0x64, 0x3c, 0x35, 0xfa, 0xd4, 0xfc, 0x86, 0x31, 0xbe, 0xcd, 0x6d, 0x64, 0x07, 0x08, 0xa6,
	0x5a, 0xc6, 0x80, 0x66, 0x8a, 0x07, 0x2c, 0x3d, 0xc7, 0x50, 0x19, 0x41, 0xaa, 0xfe, 0x62, 0x8e,
	0x1c, 0x66, 0x8a, 0x9f, 0x1a, 0x3b, 0x69, 0x41, 0xdd, 0x7a, 0x27, 0x2a, 0x96, 0xa6, 0xcb, 0xaa,
	0x0f, 0xb9, 0xa9, 0xa7, 0x62, 0x49, 0x8e, 0x60, 0x31, 0xcc, 0xa4, 0xe2, 0x49, 0x20, 0x38, 0x57,
	0x41, 0x88, 0x42, 0xb9, 0x35, 0xa3, 0xee, 0xea, 0xf4, 0xa0, 0x7d, 0x94, 0x3c, 0x13, 0x21, 0xfa,
	0x38, 0xf0, 0x17, 0x72, 0x8a, 0xcf, 0xb9, 0x3a, 0x42, 0xa1, 0xc8, 0x16, 0xfc, 0x5d, 0x74, 0x16,
	0x09, 0x3a, 0xa0, 0x29, 0x75, 0x2b, 0x26, 0xd3, 0x82, 0x35, 0x1f, 0xe7, 0xd6, 0x09, 0x21, 0x83,
	0x91, 0xe0, 0x09, 0xaa, 0x21, 0x66, 0xd2, 0x9d, 0x33, 0xbe, 0x85, 0x90, 0xaf, 0xc7, 0x00, 0xd9,
	0x84, 0x22, 0x40, 0x70, 0x4e, 0x31, 0x42, 0xe1, 0x56, 0x8d, 0xeb, 0xbc, 0xb5, 0xbe, 0x34, 0xc6,
	0xf6, 0x07, 0x07, 0xea, 0x13, 0x0b, 0x43, 0x0e, 0x00, 0x72, 0x1d, 0x27, 0x96, 0xac, 0x75, 0xcf,
	0xae, 0x4c, 0x8a, 0x7f, 0x32, 0xe3, 0xd7, 0x0c, 0x49, 0x07, 0x22, 0x07, 0xe3, 0xc4, 0x7a, 0xe0,
	0x3a, 0x4a, 0xe5, 0x77, 0x9a, 0xcc, 0xdf, 0xee, 0x01, 0xca, 0x61, 0xf7, 0x1f, 0x68, 0x6a, 0x5e,
	0x30, 0x35, 0xf1, 0x4f, 0x0e, 0x90, 0xbb, 0x0b, 0xaa, 0xaf, 0x83, 0xae, 0xcb, 0x5c, 0x93, 0xfb,
	0xaf, 0xc3, 0x71, 0xcc, 0x27, 0x0a, 0x35, 0xde, 0xe4, 0x05, 0xdc, 0xae, 0x5e, 0x30, 0x7d, 0x31,
	0x1e, 0x28, 0x73, 0x71, 0xcc, 0xb1, 0x45, 0x74, 0x57, 0x60, 0xc9, 0xb2, 0xa7, 0x8b, 0x0d, 0xa1,
	0x3e, 0x91, 0x96, 0xfc, 0x0b, 0x0d, 0x9d, 0xf6, 0xa7, 0xdd, 0xac, 0x6b, 0x5b, 0xb1, 0x9a, 0x7b,
	0x50, 0xd1, 0x3d, 0xa3, 0x2e, 0xa3, 0xf4, 0x70, 0x19, 0xd6, 0xb1, 0xbb, 0xfb, 0xf1, 0xdb, 0xba,
	0xf3, 0x6e, 0xeb, 0xc1, 0xb7, 0x7d, 0x74, 0x11, 0xd9, 0x27, 0xef, 0xac, 0x62, 0x9e, 0xba, 0x27,
	0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0xea, 0x8e, 0xf4, 0x0d, 0x06, 0x00, 0x00,
}

func (this *Install) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install)
	if !ok {
		that2, ok := that.(Install)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if that1.InstallType == nil {
		if this.InstallType != nil {
			return false
		}
	} else if this.InstallType == nil {
		return false
	} else if !this.InstallType.Equal(that1.InstallType) {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.InstalledManifest != that1.InstalledManifest {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Install_Mesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Mesh)
	if !ok {
		that2, ok := that.(Install_Mesh)
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
	if !this.Mesh.Equal(that1.Mesh) {
		return false
	}
	return true
}
func (this *Install_Ingress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Ingress)
	if !ok {
		that2, ok := that.(Install_Ingress)
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
	if !this.Ingress.Equal(that1.Ingress) {
		return false
	}
	return true
}
func (this *IstioInstall) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IstioInstall)
	if !ok {
		that2, ok := that.(IstioInstall)
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
	if this.IstioVersion != that1.IstioVersion {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if this.EnableMtls != that1.EnableMtls {
		return false
	}
	if !this.CustomRootCert.Equal(that1.CustomRootCert) {
		return false
	}
	if this.InstallGrafana != that1.InstallGrafana {
		return false
	}
	if this.InstallPrometheus != that1.InstallPrometheus {
		return false
	}
	if this.InstallJaeger != that1.InstallJaeger {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshInstall) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshInstall)
	if !ok {
		that2, ok := that.(MeshInstall)
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
	if that1.MeshInstallType == nil {
		if this.MeshInstallType != nil {
			return false
		}
	} else if this.MeshInstallType == nil {
		return false
	} else if !this.MeshInstallType.Equal(that1.MeshInstallType) {
		return false
	}
	if !this.InstalledMesh.Equal(that1.InstalledMesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshInstall_IstioMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshInstall_IstioMesh)
	if !ok {
		that2, ok := that.(MeshInstall_IstioMesh)
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
	if !this.IstioMesh.Equal(that1.IstioMesh) {
		return false
	}
	return true
}
func (this *MeshIngressInstall) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshIngressInstall)
	if !ok {
		that2, ok := that.(MeshIngressInstall)
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
	if that1.IngressInstallType == nil {
		if this.IngressInstallType != nil {
			return false
		}
	} else if this.IngressInstallType == nil {
		return false
	} else if !this.IngressInstallType.Equal(that1.IngressInstallType) {
		return false
	}
	if !this.InstalledIngress.Equal(that1.InstalledIngress) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshIngressInstall_Gloo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshIngressInstall_Gloo)
	if !ok {
		that2, ok := that.(MeshIngressInstall_Gloo)
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
	if !this.Gloo.Equal(that1.Gloo) {
		return false
	}
	return true
}
func (this *GlooInstall) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlooInstall)
	if !ok {
		that2, ok := that.(GlooInstall)
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
	if this.GlooVersion != that1.GlooVersion {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
