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
	"fmt"

	ipamv1alpha1 "github.com/nokia/k8s-ipam/apis/resource/ipam/v1alpha1"
)

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

type IpFamilyPolicy string

// IpFamilyPolicyNone defines no L3 addressing, meaning L2
const IpFamilyPolicyNone IpFamilyPolicy = "none"

// IpFamilyPolicyIPv4Only defines L3 IpFamilyPolicy as ipv4 only
const IpFamilyPolicyIPv4Only IpFamilyPolicy = "ipv4only"

// IpFamilyPolicyIPv6Only defines L3 IpFamilyPolicy as ipv6 only
const IpFamilyPolicyIPv6Only IpFamilyPolicy = "ipv6only"

// IpFamilyPolicyDualStack defines L3 IpFamilyPolicy as dual stack (ipv4 and ipv6)
const IpFamilyPolicyDualStack IpFamilyPolicy = "dualstack"

type IPFamily string

// IPFamilyIPv6 defines ipv4 ip address family
const IPFamilyIPv4 IPFamily = "ipv4"

// IPFamilyIPv6 defines ipv6 ip address family
const IPFamilyIPv6 IPFamily = "ipv6"

func IsIPFamilySupported(s string) bool {
	switch s {
	case string(IPFamilyIPv4):
	case string(IPFamilyIPv6):
	default:
		return false
	}
	return true
}

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

func IsIPFamilyPolicySupported(s string) bool {
	switch s {
	case string(IpFamilyPolicyNone):
	case string(IpFamilyPolicyIPv4Only):
	case string(IpFamilyPolicyIPv6Only):
	case string(IpFamilyPolicyDualStack):
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
	if spec.IpFamilyPolicy != "" {
		if !IsIPFamilyPolicySupported(string(spec.IpFamilyPolicy)) {
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

func (r *InterfaceStatus) UpsertIPClaim(newClaimStatus ipamv1alpha1.IPClaimStatus) {
	if newClaimStatus.Prefix == nil {
		return
	}
	if r.IPClaimStatus == nil {
		r.IPClaimStatus = []ipamv1alpha1.IPClaimStatus{}
	}
	for _, alloc := range r.IPClaimStatus {

		if alloc.Prefix != nil && *alloc.Prefix == *newClaimStatus.Prefix {
			alloc = newClaimStatus
			return
		}
	}
	r.IPClaimStatus = append(r.IPClaimStatus, newClaimStatus)
}

func (r *InterfaceSpec) GetAddressFamilies() []IPFamily {
	afs := []IPFamily{}
	switch r.IpFamilyPolicy {
	case IpFamilyPolicyDualStack:
		afs = append(afs, IPFamilyIPv4)
		afs = append(afs, IPFamilyIPv6)
	case IpFamilyPolicyIPv6Only:
		afs = append(afs, IPFamilyIPv6)
	case IpFamilyPolicyIPv4Only:
		afs = append(afs, IPFamilyIPv4)
	default:
		afs = append(afs, IPFamilyIPv4)
	}
	return afs
}
