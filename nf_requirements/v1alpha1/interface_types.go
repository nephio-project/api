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
	"reflect"

	ipamv1alpha1 "github.com/nokia/k8s-ipam/apis/alloc/ipam/v1alpha1"
	vlanv1alpha1 "github.com/nokia/k8s-ipam/apis/alloc/vlan/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
type Interface struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   InterfaceSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status InterfaceStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// TBD how do we distinguish the loopback from the vnic(s)
type InterfaceSpec struct {
	// NetworkInstance defines the networkInstance to which this interface belongs
	// Name and optionally Namespace is used here
	NetworkInstance *corev1.ObjectReference `json:"networkInstance" yaml:"networkInstance"`
	// CNIType defines the cniType that is used to attach the interface to the pod
	// +kubebuilder:validation:Enum=sriov;ipvlan;macvlan
	CNIType CNIType `json:"cniType,omitempty" yaml:"cniType,omitempty"`
	// AttachmentType defines if the interface is attached using a vlan or not
	// +kubebuilder:validation:Enum=none;vlan
	AttachmentType AttachmentType `json:"attachmentType,omitempty" yaml:"attachmentType,omitempty"`
}

type InterfaceStatus struct {
	IPAllocationStatus   *ipamv1alpha1.IPAllocationStatus   `json:"ipAllocationStatus,omitempty" yaml:"ipAllocationStatus,omitempty"`
	VLANAllocationStatus *vlanv1alpha1.VLANAllocationStatus `json:"vlanAllocationStatus,omitempty" yaml:"vlanAllocationStatus,omitempty"`
}

// Interface type metadata.
var (
	InterfaceKind             = reflect.TypeOf(Interface{}).Name()
	InterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceKind}.String()
	InterfaceKindAPIVersion   = InterfaceKind + "." + GroupVersion.String()
	InterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceKind)
)
