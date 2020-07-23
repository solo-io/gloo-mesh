// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/failover_service.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
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
//A FailoverService creates a new hostname to which services can send requests.
//Requests will be routed based on a list of backing services ordered by
//decreasing priority. When outlier detection detects that a service in the list is
//in an unhealthy state, requests sent to the FailoverService will be routed
//to the next healthy service in the list. For each service referenced in the
//failover services list, outlier detection must be configured using a TrafficPolicy.
//
//Currently this feature only supports Services backed by Istio.
type FailoverServiceSpec struct {
	//
	//The DNS name of the failover service. Must be unique within the service mesh instance
	//since it is used as the hostname with which clients communicate.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The port on which the failover service listens.
	Port *FailoverServiceSpec_Port `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// The meshes that this failover service will be visible to.
	Meshes []*v1.ObjectRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	//
	//A list of services ordered by decreasing priority for failover.
	//All services must be backed by either the same service mesh instance or
	//backed by service meshes that are grouped under a common VirtualMesh.
	ComponentServices    []*FailoverServiceSpec_ComponentService `protobuf:"bytes,4,rep,name=component_services,json=componentServices,proto3" json:"component_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *FailoverServiceSpec) Reset()         { *m = FailoverServiceSpec{} }
func (m *FailoverServiceSpec) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec) ProtoMessage()    {}
func (*FailoverServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0}
}
func (m *FailoverServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec.Unmarshal(m, b)
}
func (m *FailoverServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec.Merge(m, src)
}
func (m *FailoverServiceSpec) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec.Size(m)
}
func (m *FailoverServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec proto.InternalMessageInfo

func (m *FailoverServiceSpec) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *FailoverServiceSpec) GetPort() *FailoverServiceSpec_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *FailoverServiceSpec) GetMeshes() []*v1.ObjectRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *FailoverServiceSpec) GetComponentServices() []*FailoverServiceSpec_ComponentService {
	if m != nil {
		return m.ComponentServices
	}
	return nil
}

// The port on which the failover service listens.
type FailoverServiceSpec_Port struct {
	// Port number.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Protocol of the requests sent to the failover service, must be one of HTTP, HTTPS, GRPC, HTTP2, MONGO, TCP, TLS.
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailoverServiceSpec_Port) Reset()         { *m = FailoverServiceSpec_Port{} }
func (m *FailoverServiceSpec_Port) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec_Port) ProtoMessage()    {}
func (*FailoverServiceSpec_Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0, 0}
}
func (m *FailoverServiceSpec_Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec_Port.Unmarshal(m, b)
}
func (m *FailoverServiceSpec_Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec_Port.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec_Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec_Port.Merge(m, src)
}
func (m *FailoverServiceSpec_Port) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec_Port.Size(m)
}
func (m *FailoverServiceSpec_Port) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec_Port.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec_Port proto.InternalMessageInfo

func (m *FailoverServiceSpec_Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FailoverServiceSpec_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type FailoverServiceSpec_ComponentService struct {
	// different service types can be selected as component services.
	//
	// Types that are valid to be assigned to ComposingServiceType:
	//	*FailoverServiceSpec_ComponentService_KubeService
	ComposingServiceType isFailoverServiceSpec_ComponentService_ComposingServiceType `protobuf_oneof:"composing_service_type"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *FailoverServiceSpec_ComponentService) Reset()         { *m = FailoverServiceSpec_ComponentService{} }
func (m *FailoverServiceSpec_ComponentService) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec_ComponentService) ProtoMessage()    {}
func (*FailoverServiceSpec_ComponentService) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0, 1}
}
func (m *FailoverServiceSpec_ComponentService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec_ComponentService.Unmarshal(m, b)
}
func (m *FailoverServiceSpec_ComponentService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec_ComponentService.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec_ComponentService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec_ComponentService.Merge(m, src)
}
func (m *FailoverServiceSpec_ComponentService) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec_ComponentService.Size(m)
}
func (m *FailoverServiceSpec_ComponentService) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec_ComponentService.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec_ComponentService proto.InternalMessageInfo

type isFailoverServiceSpec_ComponentService_ComposingServiceType interface {
	isFailoverServiceSpec_ComponentService_ComposingServiceType()
	Equal(interface{}) bool
}

type FailoverServiceSpec_ComponentService_KubeService struct {
	KubeService *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=kube_service,json=kubeService,proto3,oneof" json:"kube_service,omitempty"`
}

func (*FailoverServiceSpec_ComponentService_KubeService) isFailoverServiceSpec_ComponentService_ComposingServiceType() {
}

func (m *FailoverServiceSpec_ComponentService) GetComposingServiceType() isFailoverServiceSpec_ComponentService_ComposingServiceType {
	if m != nil {
		return m.ComposingServiceType
	}
	return nil
}

func (m *FailoverServiceSpec_ComponentService) GetKubeService() *v1.ClusterObjectRef {
	if x, ok := m.GetComposingServiceType().(*FailoverServiceSpec_ComponentService_KubeService); ok {
		return x.KubeService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FailoverServiceSpec_ComponentService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FailoverServiceSpec_ComponentService_KubeService)(nil),
	}
}

type FailoverServiceStatus struct {
	//
	//The most recent generation observed in the the FailoverService metadata.
	//If the observedGeneration does not match generation, the controller has not received the most
	//recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	//
	//The state of the overall resource, will only show accepted if it has been successfully
	//applied to all target meshes.
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.smh.solo.io.ApprovalState" json:"state,omitempty"`
	// The status of the FailoverService for each Mesh to which it has been applied.
	Meshes map[string]*ApprovalStatus `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// any errors observed which prevented the resource from being Accepted.
	ValidationErrors     []string `protobuf:"bytes,4,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailoverServiceStatus) Reset()         { *m = FailoverServiceStatus{} }
func (m *FailoverServiceStatus) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceStatus) ProtoMessage()    {}
func (*FailoverServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{1}
}
func (m *FailoverServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceStatus.Unmarshal(m, b)
}
func (m *FailoverServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceStatus.Marshal(b, m, deterministic)
}
func (m *FailoverServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceStatus.Merge(m, src)
}
func (m *FailoverServiceStatus) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceStatus.Size(m)
}
func (m *FailoverServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceStatus proto.InternalMessageInfo

func (m *FailoverServiceStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *FailoverServiceStatus) GetState() ApprovalState {
	if m != nil {
		return m.State
	}
	return ApprovalState_PENDING
}

func (m *FailoverServiceStatus) GetMeshes() map[string]*ApprovalStatus {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *FailoverServiceStatus) GetValidationErrors() []string {
	if m != nil {
		return m.ValidationErrors
	}
	return nil
}

func init() {
	proto.RegisterType((*FailoverServiceSpec)(nil), "networking.smh.solo.io.FailoverServiceSpec")
	proto.RegisterType((*FailoverServiceSpec_Port)(nil), "networking.smh.solo.io.FailoverServiceSpec.Port")
	proto.RegisterType((*FailoverServiceSpec_ComponentService)(nil), "networking.smh.solo.io.FailoverServiceSpec.ComponentService")
	proto.RegisterType((*FailoverServiceStatus)(nil), "networking.smh.solo.io.FailoverServiceStatus")
	proto.RegisterMapType((map[string]*ApprovalStatus)(nil), "networking.smh.solo.io.FailoverServiceStatus.MeshesEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/failover_service.proto", fileDescriptor_4c2a3822bd950167)
}

var fileDescriptor_4c2a3822bd950167 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x76, 0x7f, 0x5a, 0xba, 0xb3, 0x2a, 0xdb, 0xa9, 0x96, 0x25, 0x88, 0x2c, 0x15, 0x65, 0x41,
	0x76, 0x62, 0xa3, 0x88, 0x3f, 0x05, 0xb1, 0xb5, 0x5a, 0x10, 0xd1, 0x4e, 0xef, 0xbc, 0x59, 0x26,
	0xe9, 0x69, 0x12, 0x93, 0xcd, 0x09, 0x33, 0x93, 0xc8, 0xbe, 0x91, 0xcf, 0xe2, 0x63, 0x78, 0xe5,
	0x63, 0x48, 0x26, 0x3f, 0x5d, 0xc2, 0x0a, 0xdb, 0xab, 0xcc, 0x9c, 0x99, 0xf3, 0x9d, 0xef, 0x3b,
	0xe7, 0x9b, 0x10, 0xee, 0x87, 0x3a, 0xc8, 0x5c, 0xe6, 0xe1, 0xc2, 0x56, 0x18, 0xe3, 0x2c, 0x44,
	0x5b, 0x81, 0xcc, 0x43, 0x0f, 0x66, 0x0b, 0x50, 0xc1, 0x2c, 0xc8, 0x5c, 0x5b, 0xa4, 0xa1, 0x9d,
	0x80, 0xfe, 0x89, 0x32, 0x0a, 0x13, 0xdf, 0xce, 0x0f, 0x45, 0x9c, 0x06, 0xc2, 0xb1, 0xaf, 0x44,
	0x18, 0x63, 0x0e, 0x72, 0x5e, 0x65, 0xb0, 0x54, 0xa2, 0x46, 0xba, 0x7f, 0x7d, 0x97, 0xa9, 0x45,
	0xc0, 0x0a, 0x5c, 0x16, 0xa2, 0xc5, 0xd6, 0xd5, 0x8a, 0x72, 0xc7, 0xe0, 0x7b, 0x28, 0xc1, 0xce,
	0x0f, 0xcd, 0xb7, 0xc4, 0xb1, 0xde, 0x6d, 0x4c, 0x24, 0x17, 0x71, 0x78, 0x29, 0x74, 0x88, 0xc9,
	0x5c, 0x69, 0xa1, 0x6b, 0x80, 0x7b, 0x3e, 0xfa, 0x68, 0x96, 0x76, 0xb1, 0x2a, 0xa3, 0x07, 0xbf,
	0x7b, 0x64, 0xef, 0x63, 0xc5, 0xfc, 0xa2, 0xac, 0x70, 0x91, 0x82, 0x47, 0x2d, 0xb2, 0x13, 0xa0,
	0xd2, 0x89, 0x58, 0xc0, 0xb8, 0x33, 0xe9, 0x4c, 0x07, 0xbc, 0xd9, 0xd3, 0x0f, 0xa4, 0x9f, 0xa2,
	0xd4, 0xe3, 0xee, 0xa4, 0x33, 0x1d, 0x3a, 0xcf, 0xd8, 0x7a, 0x85, 0x6c, 0x0d, 0x2c, 0xfb, 0x86,
	0x52, 0x73, 0x93, 0x4d, 0x5f, 0x90, 0xed, 0x42, 0x0a, 0xa8, 0x71, 0x6f, 0xd2, 0x9b, 0x0e, 0x9d,
	0x07, 0xcc, 0xa8, 0x2d, 0x7a, 0xd0, 0x40, 0x7c, 0x75, 0x7f, 0x80, 0xa7, 0x39, 0x5c, 0xf1, 0xea,
	0x2e, 0x8d, 0x08, 0xf5, 0x70, 0x91, 0x62, 0x02, 0x89, 0xae, 0x3b, 0xad, 0xc6, 0x7d, 0x83, 0x70,
	0x74, 0x13, 0x26, 0x27, 0x35, 0x4a, 0x15, 0xe4, 0xbb, 0x5e, 0x2b, 0xa2, 0xac, 0x97, 0xa4, 0x5f,
	0x10, 0xa6, 0xb4, 0x12, 0x5c, 0x34, 0xe2, 0x4e, 0x45, 0xdf, 0x22, 0x3b, 0xa6, 0x83, 0x1e, 0xc6,
	0xa6, 0x11, 0x03, 0xde, 0xec, 0xad, 0x9c, 0x8c, 0xda, 0xf0, 0xf4, 0x8c, 0xdc, 0x8e, 0x32, 0x17,
	0x6a, 0xce, 0x06, 0x6b, 0xe8, 0x3c, 0x5a, 0x23, 0xfa, 0x24, 0xce, 0x94, 0x06, 0xd9, 0x68, 0x3f,
	0xbb, 0xc5, 0x87, 0x45, 0x6a, 0x85, 0x74, 0x3c, 0x26, 0xfb, 0x86, 0xaa, 0x0a, 0x13, 0xbf, 0x86,
	0x9b, 0xeb, 0x65, 0x0a, 0x07, 0x7f, 0xbb, 0xe4, 0x7e, 0x5b, 0xab, 0x16, 0x3a, 0x53, 0xd4, 0x26,
	0x7b, 0xe8, 0x16, 0x77, 0xe1, 0x72, 0xee, 0x43, 0x02, 0xd2, 0xf8, 0xc3, 0x90, 0xe8, 0x71, 0x5a,
	0x1f, 0x7d, 0x6a, 0x4e, 0xe8, 0x5b, 0xb2, 0x65, 0xcc, 0x63, 0xb4, 0xdd, 0x75, 0x1e, 0xff, 0xaf,
	0xb5, 0xef, 0xd3, 0x54, 0x62, 0x2e, 0xe2, 0xa2, 0x0e, 0xf0, 0x32, 0x87, 0x9e, 0xb7, 0x46, 0xfb,
	0x7a, 0xd3, 0xc1, 0x18, 0xb2, 0xec, 0x8b, 0xc9, 0x3d, 0x4d, 0xb4, 0x5c, 0x36, 0x73, 0x7f, 0x4a,
	0x76, 0x57, 0x7c, 0x0d, 0x52, 0xa2, 0x2c, 0xc7, 0x3e, 0xe0, 0xa3, 0xeb, 0x83, 0x53, 0x13, 0xb7,
	0x04, 0x19, 0xae, 0x60, 0xd0, 0x11, 0xe9, 0x45, 0xb0, 0xac, 0x6c, 0x5c, 0x2c, 0xe9, 0x11, 0xd9,
	0xca, 0x45, 0x9c, 0x41, 0x65, 0xe1, 0x27, 0x9b, 0xa8, 0xcb, 0x14, 0x2f, 0x93, 0xde, 0x74, 0x5f,
	0x75, 0x8e, 0xcf, 0x7f, 0xfd, 0x79, 0xd8, 0xf9, 0xfe, 0x79, 0x93, 0x1f, 0x46, 0x1a, 0xf9, 0xad,
	0xb7, 0xba, 0x5a, 0xa3, 0x79, 0xb7, 0xee, 0xb6, 0xf1, 0xcf, 0xf3, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x1f, 0x7e, 0x83, 0x86, 0x04, 0x00, 0x00,
}

func (this *FailoverServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec)
	if !ok {
		that2, ok := that.(FailoverServiceSpec)
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
	if this.Hostname != that1.Hostname {
		return false
	}
	if !this.Port.Equal(that1.Port) {
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
	if len(this.ComponentServices) != len(that1.ComponentServices) {
		return false
	}
	for i := range this.ComponentServices {
		if !this.ComponentServices[i].Equal(that1.ComponentServices[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_Port)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_Port)
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
	if this.Port != that1.Port {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_ComponentService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_ComponentService)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_ComponentService)
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
	if that1.ComposingServiceType == nil {
		if this.ComposingServiceType != nil {
			return false
		}
	} else if this.ComposingServiceType == nil {
		return false
	} else if !this.ComposingServiceType.Equal(that1.ComposingServiceType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_ComponentService_KubeService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_ComponentService_KubeService)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_ComponentService_KubeService)
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
	if !this.KubeService.Equal(that1.KubeService) {
		return false
	}
	return true
}
func (this *FailoverServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceStatus)
	if !ok {
		that2, ok := that.(FailoverServiceStatus)
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
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if len(this.ValidationErrors) != len(that1.ValidationErrors) {
		return false
	}
	for i := range this.ValidationErrors {
		if this.ValidationErrors[i] != that1.ValidationErrors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
