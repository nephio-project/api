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

// PLMNConfigStatus defines the observed state of PLMNConfig
type PLMNConfigStatus struct {
}

// +kubebuilder:object:root=true
type PLMNConfig struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   PLMNConfigSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status PLMNConfigStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PLMNConfigList contains a list of PLMNConfigs
type PLMNConfigList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []PLMNConfig `json:"items" yaml:"items"`
}

// PLMNConfigSpec defines the characteristics of a deployment of a network function
type PLMNConfigSpec struct {
	// PLMNInfo defines the list of PLMN to configure for core NFs
	PLMNInfo []PLMNInfo `json:"PLMNInfo" yaml:"PLMNInfo"`
}

// PLMNInfo defines the structure of PLMN
type PLMNInfo struct {
	// PLMNId defines the PLMN Identifier
	PLMNId PLMNId `json:"plmnId" yaml:"plmnId"`
	// tac defines the tracking area code
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=16777215
	TAC uint32 `json:"tac" yaml:"tac"`
	// NSSAI defines the Network Slice Selection Assistance Information
	NSSAI []NSSAI `json:"nssai" yaml:"nssai"`
}

// PLMNId defines the Public Land Mobile Network Identifier
type PLMNId struct {
	// mcc defines the mobile country code
	// +kubebuilder:validation:Pattern=`[02-79][0-9][0-9]`
	MCC string `json:"mcc" yaml:"mcc"`
	// mnc defines the mobile network code
	// +kubebuilder:validation:Pattern=`[0-9][0-9][0-9]|[0-9][0-9]`
	MNC string `json:"mnc" yaml:"mnc"`
}

// NSSAI defines the Network Slice Selection Assistance Information
type NSSAI struct {
	// Sst defines Slice/Service Type
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	SST int `json:"sst" yaml:"sst"`
	// Sd defines Service Differentiator
	// +optional
	// +kubebuilder:validation:Pattern=`^[A-Fa-f0-9]{6}$`
	SD string `json:"sd,omitempty" yaml:"sd,omitempty"`
	// dnnInfo defines the Data Network Names information
	DNNInfo DnnInfo `json:"dnnInfo" yaml:"dnnInfo"`
}

// NOTE: For R2 the dnn.name here should match with NF deployment dnn name. This is a hack for R2
// DNN defines the Data Network Names Information
type DnnInfo struct {
	// name defines dnn
	Name string `json:"name" yaml:"name"`
	// SessionType defines session type
	// +kubebuilder:validation:Enum=ipv4;ipv6;ipv4v6;unstructured;ethernet
	SessionType string `json:"sessionType" yaml:"sessionType"`
	// Dns defines Domain Network Service Name
	// +optional
	DNS string `json:"dns" yaml:"dns"`
}
