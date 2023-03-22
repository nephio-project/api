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

package v1alpha1

import (
	"github.com/nephio-project/api/references"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	NetworkInstance *references.NamespaceReference `json:"networkInstance" yaml:"networkInstance"`
	// Cni defines the cni that is used to attach the interface to the pod
	// +kubebuilder:validation:Enum=sriov;ipvlan;macvlan
	Cni Cni `json:"cniType,omitempty" yaml:"cniType,omitempty"`
	// AttachmentType defines if the interface is attached using a vlan or not
	// +kubebuilder:validation:Enum=none;vlan
	AttachmentType AttachmentType `json:"attachementType,omitempty" yaml:"attachementType,omitempty"`
}

type InterfaceStatus struct {
}
