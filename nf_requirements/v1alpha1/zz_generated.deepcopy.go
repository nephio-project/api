//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Nephio.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	ipamv1alpha1 "github.com/nokia/k8s-ipam/apis/alloc/ipam/v1alpha1"
	vlanv1alpha1 "github.com/nokia/k8s-ipam/apis/alloc/vlan/v1alpha1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capacity) DeepCopyInto(out *Capacity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capacity.
func (in *Capacity) DeepCopy() *Capacity {
	if in == nil {
		return nil
	}
	out := new(Capacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Capacity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacitySpec) DeepCopyInto(out *CapacitySpec) {
	*out = *in
	out.MaxUplinkThroughput = in.MaxUplinkThroughput.DeepCopy()
	out.MaxDownlinkThroughput = in.MaxDownlinkThroughput.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacitySpec.
func (in *CapacitySpec) DeepCopy() *CapacitySpec {
	if in == nil {
		return nil
	}
	out := new(CapacitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityStatus) DeepCopyInto(out *CapacityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityStatus.
func (in *CapacityStatus) DeepCopy() *CapacityStatus {
	if in == nil {
		return nil
	}
	out := new(CapacityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNetwork) DeepCopyInto(out *DataNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNetwork.
func (in *DataNetwork) DeepCopy() *DataNetwork {
	if in == nil {
		return nil
	}
	out := new(DataNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNetworkSpec) DeepCopyInto(out *DataNetworkSpec) {
	*out = *in
	if in.Pools != nil {
		in, out := &in.Pools, &out.Pools
		*out = make([]*Pool, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Pool)
				**out = **in
			}
		}
	}
	out.NetworkInstance = in.NetworkInstance
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNetworkSpec.
func (in *DataNetworkSpec) DeepCopy() *DataNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(DataNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNetworkStatus) DeepCopyInto(out *DataNetworkStatus) {
	*out = *in
	if in.Pools != nil {
		in, out := &in.Pools, &out.Pools
		*out = make([]PoolStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNetworkStatus.
func (in *DataNetworkStatus) DeepCopy() *DataNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(DataNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interface) DeepCopyInto(out *Interface) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Interface) DeepCopy() *Interface {
	if in == nil {
		return nil
	}
	out := new(Interface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Interface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSpec) DeepCopyInto(out *InterfaceSpec) {
	*out = *in
	if in.NetworkInstance != nil {
		in, out := &in.NetworkInstance, &out.NetworkInstance
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSpec.
func (in *InterfaceSpec) DeepCopy() *InterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(InterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceStatus) DeepCopyInto(out *InterfaceStatus) {
	*out = *in
	if in.IPAllocationStatus != nil {
		in, out := &in.IPAllocationStatus, &out.IPAllocationStatus
		*out = new(ipamv1alpha1.IPAllocationStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.VLANAllocationStatus != nil {
		in, out := &in.VLANAllocationStatus, &out.VLANAllocationStatus
		*out = new(vlanv1alpha1.VLANAllocationStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceStatus.
func (in *InterfaceStatus) DeepCopy() *InterfaceStatus {
	if in == nil {
		return nil
	}
	out := new(InterfaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFAttachment) DeepCopyInto(out *NFAttachment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFAttachment.
func (in *NFAttachment) DeepCopy() *NFAttachment {
	if in == nil {
		return nil
	}
	out := new(NFAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClass) DeepCopyInto(out *NFClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClass.
func (in *NFClass) DeepCopy() *NFClass {
	if in == nil {
		return nil
	}
	out := new(NFClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClassList) DeepCopyInto(out *NFClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClassList.
func (in *NFClassList) DeepCopy() *NFClassList {
	if in == nil {
		return nil
	}
	out := new(NFClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClassSpec) DeepCopyInto(out *NFClassSpec) {
	*out = *in
	out.PackageRef = in.PackageRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClassSpec.
func (in *NFClassSpec) DeepCopy() *NFClassSpec {
	if in == nil {
		return nil
	}
	out := new(NFClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFClassStatus) DeepCopyInto(out *NFClassStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFClassStatus.
func (in *NFClassStatus) DeepCopy() *NFClassStatus {
	if in == nil {
		return nil
	}
	out := new(NFClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFInstance) DeepCopyInto(out *NFInstance) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	in.NFTemplate.DeepCopyInto(&out.NFTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFInstance.
func (in *NFInstance) DeepCopy() *NFInstance {
	if in == nil {
		return nil
	}
	out := new(NFInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTemplate) DeepCopyInto(out *NFTemplate) {
	*out = *in
	in.Capacity.DeepCopyInto(&out.Capacity)
	if in.NFAttachments != nil {
		in, out := &in.NFAttachments, &out.NFAttachments
		*out = make([]NFAttachment, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTemplate.
func (in *NFTemplate) DeepCopy() *NFTemplate {
	if in == nil {
		return nil
	}
	out := new(NFTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopology) DeepCopyInto(out *NFTopology) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopology.
func (in *NFTopology) DeepCopy() *NFTopology {
	if in == nil {
		return nil
	}
	out := new(NFTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFTopology) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopologyList) DeepCopyInto(out *NFTopologyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NFTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopologyList.
func (in *NFTopologyList) DeepCopy() *NFTopologyList {
	if in == nil {
		return nil
	}
	out := new(NFTopologyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFTopologyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopologySpec) DeepCopyInto(out *NFTopologySpec) {
	*out = *in
	if in.NFInstances != nil {
		in, out := &in.NFInstances, &out.NFInstances
		*out = make([]*NFInstance, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NFInstance)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopologySpec.
func (in *NFTopologySpec) DeepCopy() *NFTopologySpec {
	if in == nil {
		return nil
	}
	out := new(NFTopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTopologyStatus) DeepCopyInto(out *NFTopologyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFTopologyStatus.
func (in *NFTopologyStatus) DeepCopy() *NFTopologyStatus {
	if in == nil {
		return nil
	}
	out := new(NFTopologyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstance) DeepCopyInto(out *NetworkInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstance.
func (in *NetworkInstance) DeepCopy() *NetworkInstance {
	if in == nil {
		return nil
	}
	out := new(NetworkInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceList) DeepCopyInto(out *NetworkInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceList.
func (in *NetworkInstanceList) DeepCopy() *NetworkInstanceList {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceSpec) DeepCopyInto(out *NetworkInstanceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceSpec.
func (in *NetworkInstanceSpec) DeepCopy() *NetworkInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInstanceStatus) DeepCopyInto(out *NetworkInstanceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInstanceStatus.
func (in *NetworkInstanceStatus) DeepCopy() *NetworkInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRevisionReference) DeepCopyInto(out *PackageRevisionReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRevisionReference.
func (in *PackageRevisionReference) DeepCopy() *PackageRevisionReference {
	if in == nil {
		return nil
	}
	out := new(PackageRevisionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pool) DeepCopyInto(out *Pool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pool.
func (in *Pool) DeepCopy() *Pool {
	if in == nil {
		return nil
	}
	out := new(Pool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolStatus) DeepCopyInto(out *PoolStatus) {
	*out = *in
	in.IPAllocation.DeepCopyInto(&out.IPAllocation)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolStatus.
func (in *PoolStatus) DeepCopy() *PoolStatus {
	if in == nil {
		return nil
	}
	out := new(PoolStatus)
	in.DeepCopyInto(out)
	return out
}
