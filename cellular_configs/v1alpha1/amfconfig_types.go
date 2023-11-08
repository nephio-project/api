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

// AMFConfigStatus defines the observed state of AmfConfig
type AMFConfigStatus struct {
}

// +kubebuilder:object:root=true
type AMFConfig struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   AMFConfigSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status AMFConfigStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AMFConfigList contains a list of amfConfigs
type AMFConfigList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []AMFConfig `json:"items" yaml:"items"`
}

// AMFConfigSpec defines the characteristics of a deployment of a network function
type AMFConfigSpec struct {
	// amfName defines name of the amf NF
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	AMFName string `json:"amfName" yaml:"amfName"`
	// ServedGuami defines information related to GUAMI
	ServedGuami []ServedGuamiConfig `json:"servedGUAMI" yaml:"servedGUAMI"`
}

// ServedGuami information related to GUAMI
type ServedGuamiConfig struct {
	// PlmnId defines the PLMN Identifier
	PLMNId PLMNId `json:"plmnId" yaml:"plmnId"`
	// AmfId defines the AMF Identifier
	AMFId AMFId `json:"amfId" yaml:"amfId"`
}

// AmfId defines the AMF Identifier
type AMFId struct {
	// amfRegionId identifies the region
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=8
	AMFRegionId string `json:"amfRegionId" yaml:"amfRegionId"`
	// amfSetId uniquely identifies the AMF Set within the AMF Region
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=8
	AMFSetId string `json:"amfSetId" yaml:"amfSetId"`
	// amfPointer uniquely identifies the AMF in AMF set
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=6
	AMFPointer string `json:"amfPointer" yaml:"amfPointer"`
}
