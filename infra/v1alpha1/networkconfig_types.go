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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NetworkConfigSpec defines the desired state of NetworkConfig
type NetworkConfigSpec struct {
	// PrefixLengths define the prefix lengths for ipv4 and ipv6 configuration elements
	PrefixLengths *PrefixLengths `json:"prefixLengths,omitempty" yaml:"prefixLengths,omitempty"`
}

type PrefixLengths struct {
	// IPv4 defines the ipv4 prefixlengths
	// +optional
	IPv4 *IPv4PrefixLengths `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
	// IPv6 defines the ipv6 prefixlengths
	// +optional
	IPv6 *IPv6PrefixLengths `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`
}

type IPv4PrefixLengths struct {
	// +kubebuilder:default:=31
	InterfaceInternal *uint8 `json:"interfaceInternal,omitempty" yaml:"interfaceInternal,omitempty"`
	// +kubebuilder:default:=24
	InterfaceExternal *uint8 `json:"interfaceExternal,omitempty" yaml:"interfaceExternal,omitempty"`
	// +kubebuilder:default:=16
	Pool *uint8 `json:"pool,omitempty" yaml:"pool,omitempty"`
}

type IPv6PrefixLengths struct {
	// +kubebuilder:default:=127
	InterfaceInternal *uint8 `json:"interfaceInternal,omitempty" yaml:"interfaceInternal,omitempty"`
	// +kubebuilder:default:=64
	InterfaceExternal *uint8 `json:"interfaceExternal,omitempty" yaml:"interfaceExternal,omitempty"`
	// +kubebuilder:default:=48
	Pool *uint8 `json:"pool,omitempty" yaml:"pool,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="REPO_STATUS",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:resource:categories={nephio,network}

// NetworkConfig is the Schema for the Network API
type NetworkConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConfigSpec `json:"spec,omitempty"`
	Status NetworkStatus     `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkList contains a list of Repositories
type NetworkConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConfig{}, &NetworkConfigList{})
}

// NetworkConfig type metadata.
var (
	NetworkConfigKind             = reflect.TypeOf(NetworkConfig{}).Name()
	NetworkConfigGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: NetworkConfigKind}.String()
	NetworkConfigKindAPIVersion   = NetworkConfigKind + "." + GroupVersion.String()
	NetworkConfigGroupVersionKind = GroupVersion.WithKind(NetworkConfigKind)
)
