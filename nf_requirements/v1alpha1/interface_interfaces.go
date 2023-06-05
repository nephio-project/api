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

import "fmt"

const (
	errMissingNetworkInstance     = "missing networkInstance"
	errMissingNetworkInstanceName = "missing networkInstance name"
	errUnsupportedAttachmentType  = "unsupported attachmentType"
	errUnsupportedCNIType         = "unsupported cniType"
	errUnsupportedAddressing      = "unsupported addressing"
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

type Addressing string

// AddressingNone defines no L3 addressing, meaning L2
const AddressingNone Addressing = "none"

// AddressingIpv4Only defines L3 addressing as ipv4 only
const AddressingIpv4Only Addressing = "ipv4only"

// AddressingIpv6Only defines L3 addressing as ipv6 only
const AddressingIpv6Only Addressing = "ipv6only"

// AddressingDualStack defines L3 addressing as dual stack (ipv4 and ipv6)
const AddressingDualStack Addressing = "dualstack"

func IsCNITypeSupported(s string) bool {
	switch s {
	case string(CNITypeIPVLAN):
	case string(CNITypeMACVLAN):
	case string(CNITypeSRIOV):
	default:
		return false
	}
	return true
}

func IsAttachmentTypeSupported(s string) bool {
	switch s {
	case string(AttachmentTypeNone):
	case string(AttachmentTypeVLAN):
	default:
		return false
	}
	return true
}

func IsAddressingSupported(s string) bool {
	switch s {
	case string(AddressingNone):
	case string(AddressingIpv4Only):
	case string(AddressingIpv6Only):
	case string(AddressingDualStack):
	default:
		return false
	}
	return true
}

func (spec *InterfaceSpec) Validate() error {
	if spec == nil {
		return fmt.Errorf("spec invalid: %s", errMissingNetworkInstance)
	}
	if spec.AttachmentType != "" {
		if !IsAttachmentTypeSupported(string(spec.AttachmentType)) {
			return fmt.Errorf("spec invalid: %s, got: %s", errUnsupportedAttachmentType, string(spec.AttachmentType))
		}
	}
	if spec.CNIType != "" {
		if !IsCNITypeSupported(string(spec.CNIType)) {
			return fmt.Errorf("spec invalid: %s, got: %s", errUnsupportedCNIType, string(spec.CNIType))
		}
	}
	if spec.Addressing != "" {
		if !IsAddressingSupported(string(spec.Addressing)) {
			return fmt.Errorf("spec invalid: %s, got: %s", errUnsupportedAddressing, string(spec.CNIType))
		}
	}
	if spec.NetworkInstance == nil {
		return fmt.Errorf("spec invalid %s", errMissingNetworkInstance)
	}
	if spec.NetworkInstance.Name == "" {
		return fmt.Errorf("spec invalid: %s, got: %s", errMissingNetworkInstanceName, spec.NetworkInstance.Name)
	}
	return nil
}
