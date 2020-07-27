// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha2

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the TrafficPolicy.Spec
func (in *TrafficPolicySpec) DeepCopyInto(out *TrafficPolicySpec) {
	p := proto.Clone(in).(*TrafficPolicySpec)
	*out = *p
}

// DeepCopyInto for the TrafficPolicy.Status
func (in *TrafficPolicyStatus) DeepCopyInto(out *TrafficPolicyStatus) {
	p := proto.Clone(in).(*TrafficPolicyStatus)
	*out = *p
}

// DeepCopyInto for the AccessPolicy.Spec
func (in *AccessPolicySpec) DeepCopyInto(out *AccessPolicySpec) {
	p := proto.Clone(in).(*AccessPolicySpec)
	*out = *p
}

// DeepCopyInto for the AccessPolicy.Status
func (in *AccessPolicyStatus) DeepCopyInto(out *AccessPolicyStatus) {
	p := proto.Clone(in).(*AccessPolicyStatus)
	*out = *p
}

// DeepCopyInto for the VirtualMesh.Spec
func (in *VirtualMeshSpec) DeepCopyInto(out *VirtualMeshSpec) {
	p := proto.Clone(in).(*VirtualMeshSpec)
	*out = *p
}

// DeepCopyInto for the VirtualMesh.Status
func (in *VirtualMeshStatus) DeepCopyInto(out *VirtualMeshStatus) {
	p := proto.Clone(in).(*VirtualMeshStatus)
	*out = *p
}

// DeepCopyInto for the FailoverService.Spec
func (in *FailoverServiceSpec) DeepCopyInto(out *FailoverServiceSpec) {
	p := proto.Clone(in).(*FailoverServiceSpec)
	*out = *p
}

// DeepCopyInto for the FailoverService.Status
func (in *FailoverServiceStatus) DeepCopyInto(out *FailoverServiceStatus) {
	p := proto.Clone(in).(*FailoverServiceStatus)
	*out = *p
}
