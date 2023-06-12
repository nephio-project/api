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

type InterfaceUsageKind string

const (
	InterfaceUsageKindInternal InterfaceUsageKind = "internal"
	InterfaceUsageKindExternal InterfaceUsageKind = "external"
	InterfaceUsageKindPool     InterfaceUsageKind = "pool"
)

func (r *NetworkConfig) GetInterfacePrefixLength(interfaceKind InterfaceUsageKind, isIpv6 bool) uint8 {
	switch interfaceKind {
	case InterfaceUsageKindInternal:
		if isIpv6 {
			if r.Spec.PrefixLengths != nil && r.Spec.PrefixLengths.IPv6 != nil && r.Spec.PrefixLengths.IPv6.InterfaceInternal != nil {
				return *r.Spec.PrefixLengths.IPv6.InterfaceInternal
			}
			return 127
		}
		// has to be ipv4
		if r.Spec.PrefixLengths != nil && r.Spec.PrefixLengths.IPv4 != nil && r.Spec.PrefixLengths.IPv4.InterfaceInternal != nil {
			return *r.Spec.PrefixLengths.IPv4.InterfaceInternal
		}
		return 31
	case InterfaceUsageKindExternal:
		if isIpv6 {
			if r.Spec.PrefixLengths != nil && r.Spec.PrefixLengths.IPv6 != nil && r.Spec.PrefixLengths.IPv6.InterfaceExternal != nil {
				return *r.Spec.PrefixLengths.IPv6.InterfaceExternal
			}
			return 64
		}
		// has to be ipv4
		if r.Spec.PrefixLengths != nil && r.Spec.PrefixLengths.IPv4 != nil && r.Spec.PrefixLengths.IPv4.InterfaceExternal != nil {
			return *r.Spec.PrefixLengths.IPv4.InterfaceExternal
		}
		return 24
	case InterfaceUsageKindPool:
		if isIpv6 {
			if r.Spec.PrefixLengths != nil && r.Spec.PrefixLengths.IPv6 != nil && r.Spec.PrefixLengths.IPv6.Pool != nil {
				return *r.Spec.PrefixLengths.IPv6.Pool
			}
			return 48
		}
		// has to be ipv4
		if r.Spec.PrefixLengths != nil && r.Spec.PrefixLengths.IPv4 != nil && r.Spec.PrefixLengths.IPv4.Pool != nil {
			return *r.Spec.PrefixLengths.IPv4.Pool
		}
		return 16
	default:
		// invalid value
		return 0
	}
}
