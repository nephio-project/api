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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	ipamv1alpha1 "github.com/nokia/k8s-ipam/apis/alloc/ipam/v1alpha1"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func TestIsCNISupported(t *testing.T) {
	cases := map[string]struct {
		input     string
		supported bool
	}{
		"TestIsCNISupportedSRIOV": {
			input:     "sriov",
			supported: true,
		},
		"TestIsCNISupportedIPVLAN": {
			input:     "ipvlan",
			supported: true,
		},
		"TestIsCNISupportedMACVLAN": {
			input:     "macvlan",
			supported: true,
		},
		"TestIsCNISupportedNOK": {
			input:     "a",
			supported: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			b := IsCNITypeSupported(tc.input)

			if diff := cmp.Diff(tc.supported, b); diff != "" {
				t.Errorf("TestIsCNISupported: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestIsAttachmentTypeSupported(t *testing.T) {
	cases := map[string]struct {
		input     string
		supported bool
	}{
		"TestIsAttachmentTypeSupportedNone": {
			input:     "none",
			supported: true,
		},
		"TestIsAttachmentTypeSupportedVLAN": {
			input:     "vlan",
			supported: true,
		},
		"TestIsAttachmentTypeSupportedNOK": {
			input:     "a",
			supported: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			b := IsAttachmentTypeSupported(tc.input)

			if diff := cmp.Diff(tc.supported, b); diff != "" {
				t.Errorf("TestIsCNISupported: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestValidateInterfaceSpec(t *testing.T) {
	cases := map[string]struct {
		input       *InterfaceSpec
		errExpected bool
		err         string
	}{
		"Normal": {
			input: &InterfaceSpec{
				NetworkInstance: &corev1.ObjectReference{
					Name: "a",
				},
				CNIType:        "ipvlan",
				AttachmentType: "none",
				IpFamilyPolicy: "dualstack",
			},
			errExpected: false,
		},
		"MissingNetworkInstanceName": {
			input: &InterfaceSpec{
				NetworkInstance: &corev1.ObjectReference{},
			},
			errExpected: true,
			err:         errMissingNetworkInstance,
		},
		"EmptySpec": {
			input:       &InterfaceSpec{},
			errExpected: true,
			err:         errMissingNetworkInstance,
		},
		"NilSpec": {
			input:       nil,
			errExpected: true,
			err:         errMissingNetworkInstance,
		},
		"UnsupportedCNIType": {
			input: &InterfaceSpec{
				NetworkInstance: &corev1.ObjectReference{
					Name: "a",
				},
				CNIType:        "a",
				AttachmentType: "none",
			},
			errExpected: true,
			err:         errUnsupportedCNIType,
		},
		"UnsupportedAttachmentType": {
			input: &InterfaceSpec{
				NetworkInstance: &corev1.ObjectReference{
					Name: "a",
				},
				CNIType:        "sriov",
				AttachmentType: "b",
			},
			errExpected: true,
			err:         errUnsupportedAttachmentType,
		},
		"UnsupportedAddressing": {
			input: &InterfaceSpec{
				NetworkInstance: &corev1.ObjectReference{
					Name: "a",
				},
				CNIType:        "sriov",
				AttachmentType: "vlan",
				IpFamilyPolicy: "a",
			},
			errExpected: true,
			err:         errUnsupportedAddressing,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.errExpected {
				assert.Error(t, err)
				if err != nil {
					if !strings.Contains(err.Error(), tc.err) {
						t.Errorf("err want: %s, got: %s", tc.err, err.Error())
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIsAddressingSupported(t *testing.T) {
	cases := map[string]struct {
		input     string
		supported bool
	}{
		"ipv4only": {
			input:     "ipv4only",
			supported: true,
		},
		"ipv6only": {
			input:     "ipv6only",
			supported: true,
		},
		"dualstack": {
			input:     "dualstack",
			supported: true,
		},
		"none": {
			input:     "none",
			supported: true,
		},
		"empty": {
			input:     "",
			supported: false,
		},
		"unknown": {
			input:     "a",
			supported: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			b := IsIPFamilyPolicySupported(tc.input)

			if diff := cmp.Diff(tc.supported, b); diff != "" {
				t.Errorf("-want, +got:\n%s", diff)
			}
		})
	}
}

func TestIsAddressFamilySupported(t *testing.T) {
	cases := map[string]struct {
		input     string
		supported bool
	}{
		"ipv4": {
			input:     "ipv4",
			supported: true,
		},
		"ipv6": {
			input:     "ipv6",
			supported: true,
		},
		"empty": {
			input:     "",
			supported: false,
		},
		"unknown": {
			input:     "a",
			supported: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			b := IsIPFamilySupported(tc.input)

			if diff := cmp.Diff(tc.supported, b); diff != "" {
				t.Errorf("-want, +got:\n%s", diff)
			}
		})
	}
}

func TestUpsertIPAllocation(t *testing.T) {
	cases := map[string]struct {
		interfaceStatus *InterfaceStatus
		allocStatus     ipamv1alpha1.IPAllocationStatus
		expectedLength  int
	}{
		"Init": {
			interfaceStatus: &InterfaceStatus{},
			allocStatus:     ipamv1alpha1.IPAllocationStatus{Prefix: pointer.String("a")},
			expectedLength:  1,
		},
		"InitWithNil": {
			interfaceStatus: &InterfaceStatus{},
			allocStatus:     ipamv1alpha1.IPAllocationStatus{},
			expectedLength:  0,
		},
		"Update": {
			interfaceStatus: &InterfaceStatus{
				IPAllocationStatus: []ipamv1alpha1.IPAllocationStatus{
					{Prefix: pointer.String("a")},
				},
			},
			allocStatus:    ipamv1alpha1.IPAllocationStatus{Prefix: pointer.String("a")},
			expectedLength: 1,
		},
		"Add": {
			interfaceStatus: &InterfaceStatus{
				IPAllocationStatus: []ipamv1alpha1.IPAllocationStatus{
					{Prefix: pointer.String("a")},
				},
			},
			allocStatus:    ipamv1alpha1.IPAllocationStatus{Prefix: pointer.String("b")},
			expectedLength: 2,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			tc.interfaceStatus.UpsertIPAllocation(tc.allocStatus)

			if len(tc.interfaceStatus.IPAllocationStatus) != tc.expectedLength {
				t.Errorf("-want: %d, +got: %d", tc.expectedLength, len(tc.interfaceStatus.IPAllocationStatus))
			}
		})
	}
}
