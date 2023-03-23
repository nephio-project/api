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
	"reflect"

	"github.com/nephio-project/api/references"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
type DataNetworkName struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   DataNetworkNameSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status DataNetworkNameStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type DataNetworkNameSpec struct {
	// Pools defines the parameters of the IP pool associated with the DNN
	Pools []*Pool `json:"pools,omitempty"`
	// NetworkInstance defines the networkInstance context to which this DNN belongs
	NetworkInstance *references.NamespaceReference `json:"networkInstanceReference" yaml:"networkReference"`
}

type Pool struct {
	// Name defines the name of the pool
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name uint8 `json:"name,omitempty" yaml:"name,omitempty"`
	// PrefixLength define the size of the pool
	PrefixLength uint8 `json:"prefixLength,omitempty" yaml:"prefixLength,omitempty"`
}

type DataNetworkNameStatus struct {
}

// DataNetworkName type metadata.
var (
	DataNetworkNameKind             = reflect.TypeOf(DataNetworkName{}).Name()
	DataNetworkNameGroupKind        = schema.GroupKind{Group: Group, Kind: DataNetworkNameKind}.String()
	DataNetworkNameKindAPIVersion   = DataNetworkNameKind + "." + GroupVersion.String()
	DataNetworkNameGroupVersionKind = GroupVersion.WithKind(DataNetworkNameKind)
)
