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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CommonCoreRANConfig struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   CommonCoreRANConfigSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`

}

//+kubebuilder:object:root=true

// CommonCoreRANConfigList contains a list of CommonCoreRANConfigs
type CommonCoreRANConfigList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []CommonCoreRANConfig `json:"items" yaml:"items"`
}

// CommonCoreRANConfigSpec defines the characteristics of a deployment of a network function
type CommonCoreRANConfigSpec struct {
	// plmnInfo defines the list of PLMNs to configure for core NFs
	PlmnInfo []PlmnInfoConfig `json:"plmnInfo,omitempty" yaml:"plmnInfo,omitempty"`
	// communicationModel defines the communcation strategy between core components, using NRF or without NRF
	// +optional
	CommunicationModel []CommunicationModelConfig `json:"communicationModel,omitempty" yaml:"communicationModel,omitempty"`
}

type ObjectReference struct {
	// APIVersion of the target resources
	APIVersion string `yaml:"apiVersion,omitempty" json:"apiVersion,omitempty"`

	// Kind of the target resources
	Kind string `yaml:"kind,omitempty" json:"kind,omitempty"`

	// Name of the target resource
	// +optional
	Name *string `yaml:"name" json:"name"`

	// Note: Namespace is not allowed; the namespace
	// must match the namespace of the PackageVariantSet resource
}

// PlmnInfoConfig defines the structure of PLMN
type PlmnInfoConfig struct {
	// mcc defines the mobile country code 
	//
	// +kubebuilder:validation:Pattern=`[02-79][0-9][0-9]`
	Mcc string `json:"mcc" yaml:"mcc"`
	// mnc defines the mobile network code 
	//
	// +kubebuilder:validation:Pattern=`[0-9][0-9][0-9]|[0-9][0-9]`
	Mnc string `json:"mnc" yaml:"mnc"`
	// tac defines the tracking area code 
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=16777215
	Tac int `json:"tac" yaml:"tac"`
	// NSSAI defines the Network Slice Selection Assistance Information
	// +optional
	Nssai []Nssai `json:"nssai,omitempty" yaml:"nssai,omitempty"`

}

// NSSAI defines the Network Slice Selection Assistance Information
type Nssai struct {
	// Sst defines Slice/Service Type
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	Sst int `json:"sst" yaml:"sst"`
	// Sd defines Service Differentiator
	// +optional
	// +kubebuilder:validation:Pattern=`^[A-Fa-f0-9]{6}$`
	Sd *string `json:"sd,omitempty" yaml:"sd,omitempty"`
	// dnnInfo defines the Data Network Names information 
	DnnInfo []DnnInfo `json:"dnnInfo,omitempty" yaml:"dnnInfo,omitempty"`
}


// DNN defines the Data Network Names Information
type DnnInfo struct {
	// name defines dnn
	Name string `json:"name" yaml:"name"`
	// SessionType defines session type 
	// +kubebuilder:validation:Enum=ipv4;ipv6;ipv4v6;unstructured;ethernet
	SessionType *string `json:"sessionType" yaml:"sessionType"`
	// Dns defines Domain Network Service Name
	// +optional
	Dns *string `json:"dns" yaml:"dns"`
}

// CommunicationModelConfig defines the Communication Model
type CommunicationModelConfig struct {
	// name defines name of the communication model
	// +kubebuilder:validation:Enum=withNrf;woNrf
	Name string `json:"name" yaml:"name"`
	// nrfUri defines NRF Uniform Resource Identifier
	NrfUri *string `json:"nrfUri" yaml:"nrfUri"`
}