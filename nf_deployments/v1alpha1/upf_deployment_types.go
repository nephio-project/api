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

type UPFDeployment struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   UPFDeploymentSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status UPFDeploymentStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UPFDeploymentList contains a list of UPFDeployments
type UPFDeploymentList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []UPFDeployment `json:"items" yaml:"items"`
}

type UPFDeploymentSpec struct {
	NFDeploymentSpec `json:",inline" yaml:",inline"`
}

type UPFDeploymentStatus struct {
	NFDeploymentStatus `json:",inline" yaml:",inline"`
}

// Implement NFDeployment interface

func (d *UPFDeployment) GetNFDeploymentSpec() *NFDeploymentSpec {
	return d.Spec.NFDeploymentSpec.DeepCopy()
}
func (d *UPFDeployment) GetNFDeploymentStatus() *NFDeploymentStatus {
	return d.Status.NFDeploymentStatus.DeepCopy()
}
func (d *UPFDeployment) SetNFDeploymentSpec(s *NFDeploymentSpec) {
	s.DeepCopyInto(&d.Spec.NFDeploymentSpec)
}
func (d *UPFDeployment) SetNFDeploymentStatus(s *NFDeploymentStatus) {
	s.DeepCopyInto(&d.Status.NFDeploymentStatus)
}

// Interface type metadata.
var (
	UPFDeploymentKind             = reflect.TypeOf(UPFDeployment{}).Name()
	UPFDeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: UPFDeploymentKind}.String()
	UPFDeploymentKindAPIVersion   = UPFDeploymentKind + "." + GroupVersion.String()
	UPFDeploymentGroupVersionKind = GroupVersion.WithKind(UPFDeploymentKind)
)
