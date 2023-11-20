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

// ServedGUAMIStatus defines the observed state of ServedGUAMI
type ServedGUAMIStatus struct {
}

// +kubebuilder:object:root=true
type ServedGUAMI struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   ServedGUAMISpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status ServedGUAMIStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServedGUAMIList contains a list of amf
type ServedGUAMIList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []ServedGUAMI `json:"items" yaml:"items"`
}

// ServedGUAMISpec defines the servedGUAMI for an AMF
type ServedGUAMISpec struct {
	// PlmnID defines the PLMN Identifier
	PLMNID PLMNID `json:"plmnID" yaml:"plmnID"`
	// AMFID defines the AMF Identifier
	AMFID AMFID `json:"amfID" yaml:"amfID"`
}

// AMFID defines the AMF Identifier
type AMFID struct {
	// AMFRegionID identifies the region
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=8
	AMFRegionID string `json:"amfRegionID" yaml:"amfRegionID"`
	// AMFfSetID uniquely identifies the AMF Set within the AMF Region
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=8
	AMFSetID string `json:"amfSetID" yaml:"amfSetID"`
	// AMFPointer uniquely identifies the AMF in AMF set
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=6
	AMFPointer string `json:"amfPointer" yaml:"amfPointer"`
}
