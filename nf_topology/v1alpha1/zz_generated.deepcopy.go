//go:build !ignore_autogenerated

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
	nf_deploymentsv1alpha1 "github.com/nephio-project/api/nf_deployments/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFConnectivity) DeepCopyInto(out *NFConnectivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFConnectivity.
func (in *NFConnectivity) DeepCopy() *NFConnectivity {
	if in == nil {
		return nil
	}
	out := new(NFConnectivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFDeployedInstance) DeepCopyInto(out *NFDeployedInstance) {
	*out = *in
	if in.Connectivities != nil {
		in, out := &in.Connectivities, &out.Connectivities
		*out = make([]NFConnectivity, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFDeployedInstance.
func (in *NFDeployedInstance) DeepCopy() *NFDeployedInstance {
	if in == nil {
		return nil
	}
	out := new(NFDeployedInstance)
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
func (in *NFInterface) DeepCopyInto(out *NFInterface) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFInterface.
func (in *NFInterface) DeepCopy() *NFInterface {
	if in == nil {
		return nil
	}
	out := new(NFInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NFInterface) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFTemplate) DeepCopyInto(out *NFTemplate) {
	*out = *in
	out.NFPackageRef = in.NFPackageRef
	in.Capacity.DeepCopyInto(&out.Capacity)
	if in.NFInterfaces != nil {
		in, out := &in.NFInterfaces, &out.NFInterfaces
		*out = make([]NFInterface, len(*in))
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
	in.Status.DeepCopyInto(&out.Status)
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
		*out = make([]NFInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]nf_deploymentsv1alpha1.NFDeploymentConditionType, len(*in))
		copy(*out, *in)
	}
	if in.NFInstances != nil {
		in, out := &in.NFInstances, &out.NFInstances
		*out = make([]NFDeployedInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
