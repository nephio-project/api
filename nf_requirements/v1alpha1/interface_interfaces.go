/*
Copyright 2023 The Nephio Authors.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/types"
)

type AttachmentType string

// AttachmentTypeNone defines an untagged attachement (no VLAN)
const AttachmentTypeNone AttachmentType = "none"

// AttachmentTypeVLAN defines a tagged/vlan attachement
const AttachmentTypeVLAN AttachmentType = "vlan"

type CNIType string

// CNITypeSRIOV defines the sriov cni
const CNITypeSRIOV CNIType = "sriov"

// CNITypeIPVLAN defines the ipvlan cni
const CNITypeIPVLAN CNIType = "ipvlan"

// CNITypeMACVLAN defines the macvlan cni
const CNITypeMACVLAN CNIType = "macvlan"

// GetNetworkInstance returns the networkinstance the interface belongs to
func (r *Interface) GetNetworkInstance() types.NamespacedName {
	nsn := types.NamespacedName{}
	if r.Spec.NetworkInstance != nil {
		nsn.Name = r.Spec.NetworkInstance.Name
		nsn.Namespace = r.Spec.NetworkInstance.Namespace
	}
	return nsn
}

// GetCNIType returns the cnitype of the interface
func (r *Interface) GetCNIType() *CNIType {
	return r.Spec.CNIType
}

// GetAttachmentType returns the attachment type of the interface
func (r *Interface) GetAttachmentType() *AttachmentType {
	return r.Spec.AttachmentType
}
