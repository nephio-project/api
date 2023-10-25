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
type AmfConfig struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   AmfConfigSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`

}

//+kubebuilder:object:root=true

// AmfConfigList contains a list of amfConfigs
type AmfConfigList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []amfConfig `json:"items" yaml:"items"`
}

// AmfConfigSpec defines the characteristics of a deployment of a network function
type AmfConfigSpec struct {
	// amfName defines name of the amf NF
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	AmfName string `json:"amfName" yaml:"amfName"`
	// ServedGuami defines information related to GUAMI
	ServedGuami []ServedGuamiConfig `json:"servedGuami" yaml:"servedGuami"`
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

// ServedGuami information related to GUAMI
type ServedGuamiConfig struct {
	// PlmnId defines the PLMN Identifier
	PlmnId PlmnId `json:"plmnId" yaml:"plmnId"`
	// AmfId defines the AMF Identifier
	AmfId AmfId `json:"amfId" yaml:"amfId"`
}

// PlmnId defines the Public Land Mobile Network Identifier
type PlmnId struct {
	// mcc defines the mobile country code 
	//
	// +kubebuilder:validation:Pattern=`[02-79][0-9][0-9]`
	Mcc string `json:"mcc" yaml:"mcc"`
	// mnc defines the mobile network code 
	//
	// +kubebuilder:validation:Pattern=`[0-9][0-9][0-9]|[0-9][0-9]`
	Mnc string `json:"mnc" yaml:"mnc"`
}

// AmfId defines the AMF Identifier
type AmfId struct {
	// amfRegionId identifies the region
	//
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=8
	AmfRegionId string  `json:"amfRegionId" yaml:"amfRegionId"`
	// amfSetId uniquely identifies the AMF Set within the AMF Region
	//
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=8
	AmfSetId string  `json:"amfSetId" yaml:"amfSetId"`
	// amfPointer uniquely identifies the AMF in AMF set
	//
	// +kubebuilder:validation:Pattern=`[01]*`
	// +kubebuilder:validation:MaxLength=6
	AmfPointer string  `json:"amfPointer" yaml:"amfPointer"`
}