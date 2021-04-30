// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for rbac.enterprise.mesh.gloo.solo.io/v1 resources

package v1

import (
    runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for Role

func (in *Role) DeepCopyInto(out *Role) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *Role) DeepCopy() *Role {
    if in == nil {
        return nil
    }
    out := new(Role)
    in.DeepCopyInto(out)
    return out
}

func (in *Role) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *RoleList) DeepCopyInto(out *RoleList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]Role, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *RoleList) DeepCopy() *RoleList {
    if in == nil {
        return nil
    }
    out := new(RoleList)
    in.DeepCopyInto(out)
    return out
}

func (in *RoleList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

// Generated Deepcopy methods for RoleBinding

func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
    out.TypeMeta = in.TypeMeta
    in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

    // deepcopy spec
    in.Spec.DeepCopyInto(&out.Spec)
    // deepcopy status
    in.Status.DeepCopyInto(&out.Status)

    return
}

func (in *RoleBinding) DeepCopy() *RoleBinding {
    if in == nil {
        return nil
    }
    out := new(RoleBinding)
    in.DeepCopyInto(out)
    return out
}

func (in *RoleBinding) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

func (in *RoleBindingList) DeepCopyInto(out *RoleBindingList) {
    *out = *in
    out.TypeMeta = in.TypeMeta
    in.ListMeta.DeepCopyInto(&out.ListMeta)
    if in.Items != nil {
        in, out := &in.Items, &out.Items
        *out = make([]RoleBinding, len(*in))
        for i := range *in {
            (*in)[i].DeepCopyInto(&(*out)[i])
        }
    }
    return
}

func (in *RoleBindingList) DeepCopy() *RoleBindingList {
    if in == nil {
        return nil
    }
    out := new(RoleBindingList)
    in.DeepCopyInto(out)
    return out
}

func (in *RoleBindingList) DeepCopyObject() runtime.Object {
    if c := in.DeepCopy(); c != nil {
        return c
    }
    return nil
}

