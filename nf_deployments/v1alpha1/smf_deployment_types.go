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

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

type SMFDeployment struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   SMFDeploymentSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status SMFDeploymentStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SMFDeploymentList contains a list of SMFDeployments
type SMFDeploymentList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []SMFDeployment `json:"items" yaml:"items"`
}

type SMFDeploymentSpec struct {
	NFDeploymentSpec `json:",inline" yaml:",inline"`
}

type SMFDeploymentStatus struct {
	NFDeploymentStatus `json:",inline" yaml:",inline"`
}

// Interface type metadata.
var (
	SMFDeploymentKind              = reflect.TypeOf(SMFDeployment{}).Name()
	SMFDeploymentGroupKind         = schema.GroupKind{Group: Group, Kind: SMFDeploymentKind}.String()
	SMFDeploymentKindAPIVersion    = SMFDeploymentKind + "." + GroupVersion.String()
	SMFDeploymenteGroupVersionKind = GroupVersion.WithKind(SMFDeploymentKind)
)
