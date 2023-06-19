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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GVKC: group, kind, config
type GVKC struct {
	// Group specifies the Kubernetes API group (api version) for this ref
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// Version specifies the version for this ref
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Kind specifies the resource kind for this ref
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	// Config specifies the resource content (spec) in JSON string format for this ref
	Config string `json:"config,omitempty" yaml:"config,omitempty"`
}

// ConfigRefSpec defines the structure for config reference specification
type ConfigRefSpec struct {
	// GVK specifies the referenced data embedded inside ConfigRef
	GVKC GVKC `json:"gvkc,omitempty" yaml:"gvkc,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:path=configref

// ConfigRef is the Schema for the ConfigRef API
type ConfigRef struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigRefSpec   `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigRefList contains a list of NFClass
type ConfigRefList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigRef `json:"items"`
}
