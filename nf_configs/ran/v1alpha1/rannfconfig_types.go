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
	Sst uint8 `json:"sst,omitempty"`
	// +kubebuilder:validation:Pattern=`^[A-Fa-f0-9]{6}$`
	Sd string `json:"sd,omitempty"`
}

type PlmnInfo struct {
	// +kubebuilder:validation:Pattern=`[02-79][0-9][0-9]`
	Mcc string `json:"mcc,omitempty"`
	// +kubebuilder:validation:Pattern=`[0-9][0-9][0-9]|[0-9][0-9]`
	Mnc string `json:"mnc,omitempty"`
	// +kubebuilder:validation:Enum=2;3
	MncLength uint8 `json:"mncLength,omitempty"`
	//tac defines the tracking area code to be used by the cell
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=16777215
	Tac uint32 `json:"tac,omitempty"`
	//nssaiList defines the Nssai list to be configured for the cell
	NssaiList []Nssai `json:"nssaiList,omitempty"`
}

// RanNfConfigSpec defines the desired state of RanNfConfig
type RanNfConfigSpec struct {
	//gNB Identity
	GnbId string `json:"gnbId,omitempty"`
	//cellIdentity defines the cell identity of a cell
	CellIdentity string `json:"cellIdentity,omitempty"`
	//physicalCellId defines the physical cell identity of a cell
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=503
	PhysicalCellId uint32 `json:"physicalCellId,omitempty"`
	//plmn defines the plmn of a cell
	PlmnInfo            `json:"plmnInfo,omitempty"`
	DlFrequencyBand     uint32 `json:"dlFrequencyBand,omitempty"`
	DlSubCarrierSpacing uint16 `json:"dlSubCarrierSpacing,omitempty"`
	DlCarrierBandwidth  uint32 `json:"dlCarrierBandwidth,omitempty"`
	UlFrequencyBand     uint32 `json:"ulFrequencyBand,omitempty"`
	UlSubCarrierSpacing uint16 `json:"ulSubCarrierSpacing,omitempty"`
	UlCarrierBandwidth  uint32 `json:"ulCarrierBandwidth,omitempty"`
}

// RanNfConfigStatus defines the observed state of RanNfConfig
type RanNfConfigStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RanNfConfig is the Schema for the rannfconfigs API
type RanNfConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RanNfConfigSpec   `json:"spec,omitempty"`
	Status RanNfConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RanNfConfigList contains a list of RanNfConfig
type RanNfConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RanNfConfig `json:"items"`
}
