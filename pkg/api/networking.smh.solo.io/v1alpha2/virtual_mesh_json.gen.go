// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/virtual_mesh.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for VirtualMeshSpec
func (this *VirtualMeshSpec) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec
func (this *VirtualMeshSpec) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_MTLSConfig
func (this *VirtualMeshSpec_MTLSConfig) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_MTLSConfig
func (this *VirtualMeshSpec_MTLSConfig) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_MTLSConfig_SharedTrust
func (this *VirtualMeshSpec_MTLSConfig_SharedTrust) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_MTLSConfig_SharedTrust
func (this *VirtualMeshSpec_MTLSConfig_SharedTrust) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_MTLSConfig_LimitedTrust
func (this *VirtualMeshSpec_MTLSConfig_LimitedTrust) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_MTLSConfig_LimitedTrust
func (this *VirtualMeshSpec_MTLSConfig_LimitedTrust) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_RootCertificateAuthority
func (this *VirtualMeshSpec_RootCertificateAuthority) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_RootCertificateAuthority
func (this *VirtualMeshSpec_RootCertificateAuthority) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert
func (this *VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert
func (this *VirtualMeshSpec_RootCertificateAuthority_SelfSignedCert) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshSpec_Federation
func (this *VirtualMeshSpec_Federation) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshSpec_Federation
func (this *VirtualMeshSpec_Federation) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualMeshStatus
func (this *VirtualMeshStatus) MarshalJSON() ([]byte, error) {
	str, err := VirtualMeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualMeshStatus
func (this *VirtualMeshStatus) UnmarshalJSON(b []byte) error {
	return VirtualMeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	VirtualMeshMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	VirtualMeshUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
