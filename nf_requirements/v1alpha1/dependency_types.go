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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
type Dependency struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   DependencySpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status DependencyStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type DependencySpec struct {
	// PackageName defines the name of the packages this dependency is dependent upon
	PackageName string `json:"packageName" yaml:"packageName"`
	// Injectors define the object references the dependency wants to inject
	// relevant objects are APIVersion and Kind
	Injectors []corev1.ObjectReference `json:"injectors,omitempty" yaml:"injectors,omitempty"`
}

type DependencyStatus struct {
	Injected []corev1.ObjectReference `json:"injected,omitempty" yaml:"injected,omitempty"`
}

// Dependency type metadata.
var (
	DependencyKind             = reflect.TypeOf(Dependency{}).Name()
	DependencyGroupKind        = schema.GroupKind{Group: Group, Kind: DependencyKind}.String()
	DependencyKindAPIVersion   = DependencyKind + "." + GroupVersion.String()
	DependencyGroupVersionKind = GroupVersion.WithKind(DependencyKind)
)
