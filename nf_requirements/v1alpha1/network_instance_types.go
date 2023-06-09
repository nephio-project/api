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

// NetworkInstanceSpec is meant for defining a type of network's characteristics
// It is a placeholder for now
type NetworkInstanceSpec struct {
	// NetworkName is just a dummy for now
	NetworkName string `json:"networkName,omitempty" yaml:"networkName,omitempty"`
}

// NetworkInstanceStatus defines the observed state of Networkinstance
type NetworkInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=networkinstances,scope=Cluster

// NetworkInstance is the Schema for the networkinstance API
type NetworkInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkInstanceSpec   `json:"spec,omitempty"`
	Status NetworkInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkInstanceList contains a list of NetworkInstance
type NetworkInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInstance `json:"items"`
}
