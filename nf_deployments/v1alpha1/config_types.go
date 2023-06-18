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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ConfigSpec defines the desired state of the  Configuration
type ConfigSpec struct {
	// Config defines the configuration data, which is a json encapsulated resource
	//+kubebuilder:pruning:PreserveUnknownFields
	Config runtime.RawExtension `json:"config" yaml:"config"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories={nephio,config}

// Config is the Schema for the Config API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ConfigSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigList contains a list of Configs
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}

// Config type metadata.
var (
	ConfigKind             = reflect.TypeOf(Config{}).Name()
	ConfigGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: ConfigKind}.String()
	ConfigKindAPIVersion   = ConfigKind + "." + GroupVersion.String()
	ConfigGroupVersionKind = GroupVersion.WithKind(ConfigKind)
)
