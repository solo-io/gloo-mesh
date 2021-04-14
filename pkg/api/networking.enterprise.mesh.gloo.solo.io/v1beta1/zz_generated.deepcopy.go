// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for networking.enterprise.mesh.gloo.solo.io/v1beta1 resources

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for WasmDeployment

func (in *WasmDeployment) DeepCopyInto(out *WasmDeployment) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *WasmDeployment) DeepCopy() *WasmDeployment {
	if in == nil {
		return nil
	}
	out := new(WasmDeployment)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *WasmDeploymentList) DeepCopyInto(out *WasmDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WasmDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *WasmDeploymentList) DeepCopy() *WasmDeploymentList {
	if in == nil {
		return nil
	}
	out := new(WasmDeploymentList)
	in.DeepCopyInto(out)
	return out
}

func (in *WasmDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for VirtualDestination

func (in *VirtualDestination) DeepCopyInto(out *VirtualDestination) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *VirtualDestination) DeepCopy() *VirtualDestination {
	if in == nil {
		return nil
	}
	out := new(VirtualDestination)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualDestination) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualDestinationList) DeepCopyInto(out *VirtualDestinationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualDestinationList) DeepCopy() *VirtualDestinationList {
	if in == nil {
		return nil
	}
	out := new(VirtualDestinationList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualDestinationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for FederatedGateway

func (in *FederatedGateway) DeepCopyInto(out *FederatedGateway) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *FederatedGateway) DeepCopy() *FederatedGateway {
	if in == nil {
		return nil
	}
	out := new(FederatedGateway)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *FederatedGatewayList) DeepCopyInto(out *FederatedGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *FederatedGatewayList) DeepCopy() *FederatedGatewayList {
	if in == nil {
		return nil
	}
	out := new(FederatedGatewayList)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for RouteTable

func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

func (in *RouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *RouteTableList) DeepCopyInto(out *RouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *RouteTableList) DeepCopy() *RouteTableList {
	if in == nil {
		return nil
	}
	out := new(RouteTableList)
	in.DeepCopyInto(out)
	return out
}

func (in *RouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for DelegatedRouteTable

func (in *DelegatedRouteTable) DeepCopyInto(out *DelegatedRouteTable) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *DelegatedRouteTable) DeepCopy() *DelegatedRouteTable {
	if in == nil {
		return nil
	}
	out := new(DelegatedRouteTable)
	in.DeepCopyInto(out)
	return out
}

func (in *DelegatedRouteTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *DelegatedRouteTableList) DeepCopyInto(out *DelegatedRouteTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DelegatedRouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *DelegatedRouteTableList) DeepCopy() *DelegatedRouteTableList {
	if in == nil {
		return nil
	}
	out := new(DelegatedRouteTableList)
	in.DeepCopyInto(out)
	return out
}

func (in *DelegatedRouteTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
