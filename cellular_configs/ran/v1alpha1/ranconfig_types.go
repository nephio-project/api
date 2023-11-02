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

type Nssai struct {
	Sst uint8 `json:"sst"`
	// +kubebuilder:validation:Pattern=`^[A-Fa-f0-9]{6}$`
	Sd string `json:"sd"`
}

type PlmnInfo struct {
	// +kubebuilder:validation:Pattern=`[02-79][0-9][0-9]`
	Mcc string `json:"mcc"`
	// +kubebuilder:validation:Pattern=`[0-9][0-9][0-9]|[0-9][0-9]`
	Mnc string `json:"mnc"`
	// +kubebuilder:validation:Enum=2;3
	MncLength uint8 `json:"mncLength"`
	//tac defines the tracking area code to be used by the cell
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=16777215
	Tac uint32 `json:"tac"`
	//nssaiList defines the Nssai list to be configured for the cell
	NssaiList []Nssai `json:"nssaiList,omitempty"`
}

// RanConfigSpec defines the desired state of RanConfig
type RanConfigSpec struct {
	//cellIdentity defines the cell identity of a cell
	CellIdentity string `json:"cellIdentity"`
	//physicalCellId defines the physical cell identity of a cell
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=503
	PhysicalCellId uint32 `json:"physicalCellId"`
	//plmn defines the plmn of a cell
	PlmnInfo            `json:"plmnInfo,omitempty"`
	DlFrequencyBand     uint32 `json:"dlFrequencyBand"`
	DlSubCarrierSpacing uint16 `json:"dlSubCarrierSpacing"`
	DlCarrierBandwidth  uint32 `json:"dlCarrierBandwidth"`
	UlFrequencyBand     uint32 `json:"ulFrequencyBand"`
	UlSubCarrierSpacing uint16 `json:"ulSubCarrierSpacing"`
	UlCarrierBandwidth  uint32 `json:"ulCarrierBandwidth"`
}

// RanConfigStatus defines the observed state of RanConfig
type RanConfigStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RanConfig is the Schema for the ranconfigs API
type RanConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RanConfigSpec   `json:"spec,omitempty"`
	Status RanConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RanConfigList contains a list of RanConfig
type RanConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RanConfig `json:"items"`
}
