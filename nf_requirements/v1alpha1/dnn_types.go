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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
type DataNetwork struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   DataNetworkSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status DataNetworkStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type DataNetworkSpec struct {
	// Pools defines the parameters of the IP pool associated with the DNN
	Pools []*Pool `json:"pools,omitempty"`
	// NetworkInstance defines the networkInstance context to which this DNN belongs
	// Name and optionally Namespace is used here
	NetworkInstance corev1.ObjectReference `json:"networkInstance" yaml:"networkInstance"`
}

type Pool struct {
	// Name defines the name of the pool
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// PrefixLength define the size of the pool
	PrefixLength uint8 `json:"prefixLength,omitempty" yaml:"prefixLength,omitempty"`
}

type DataNetworkStatus struct {
	IPAllocationStatus *ipamv1alpha1.IPAllocationStatus `json:"ipAllocationStatus,omitempty" yaml:"ipAllocationStatus,omitempty"`
}

// DataNetworkName type metadata.
var (
	DataNetworkKind             = reflect.TypeOf(DataNetwork{}).Name()
	DataNetworkGroupKind        = schema.GroupKind{Group: Group, Kind: DataNetworkKind}.String()
	DataNetworkKindAPIVersion   = DataNetworkKind + "." + GroupVersion.String()
	DataNetworkGroupVersionKind = GroupVersion.WithKind(DataNetworkKind)
)
