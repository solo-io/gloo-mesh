// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/core/v1alpha1/service_selector.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ServiceSelector
func (this *ServiceSelector) MarshalJSON() ([]byte, error) {
	str, err := ServiceSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceSelector
func (this *ServiceSelector) UnmarshalJSON(b []byte) error {
	return ServiceSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceSelector_Matcher
func (this *ServiceSelector_Matcher) MarshalJSON() ([]byte, error) {
	str, err := ServiceSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceSelector_Matcher
func (this *ServiceSelector_Matcher) UnmarshalJSON(b []byte) error {
	return ServiceSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ServiceSelector_ServiceRefs
func (this *ServiceSelector_ServiceRefs) MarshalJSON() ([]byte, error) {
	str, err := ServiceSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceSelector_ServiceRefs
func (this *ServiceSelector_ServiceRefs) UnmarshalJSON(b []byte) error {
	return ServiceSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceSelectorMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ServiceSelectorUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
