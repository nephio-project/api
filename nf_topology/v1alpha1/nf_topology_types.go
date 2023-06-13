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
	nephioreqv1alpha1 "github.com/nephio-project/api/nf_requirements/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// NFType defines the type of network functions
type NFType string

const (
	NFTypeUPF NFType = "upf"
	NFTypeSMF NFType = "smf"
	NFTypeAMF NFType = "amf"
)

// NFInterface defines the specification of network attachment points of a NF
type NFInterface struct {
	// Name of the network attachment point
	Name string `json:"name" yaml:"name"`

	// NetworkInstanceRef is a reference to NetworkInstance. Two NF with attachment to
	// the same NetworkInstance is considered connected neighbors
	NetworkInstanceName string `json:"networkInstanceName" yaml:"networkInstanceName"`
}

// NFTemplate defines the template for deployment of an instance of a NF
type NFTemplate struct {
	// NFType specifies the type of NF this template is specifying
	NFType NFType `json:"nfType" yaml:"nfType"`

	// ClassName --- for now, the NFClass this NF template will derive from
	ClassName string `json:"className" yaml:"className"`

	// Capacity specifies the NF capacity profile for this NF instance
	Capacity nephioreqv1alpha1.CapacitySpec `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// NFInterfaces
	NFInterfaces []NFInterface `json:"nfInterfaces,omitempty" yaml:"nfInterfaces,omitempty"`
}

type NFInstance struct {
	// Name specifies the name of this NFInstance
	Name string `json:"name" yaml:"name"`

	// ClusterSelector specifies the matching labels for the NF instance to be instantiated
	ClusterSelector metav1.LabelSelector `json:"clusterSelector" yaml:"clusterSelector"`

	// NFTemplate specifies the template of the NF to be deployed when a cluster matches
	// the selector above
	NFTemplate NFTemplate `json:"nfTemplate" yaml:"nfTemplate"`
}

type NFTopologySpec struct {
	NFInstances []NFInstance `json:"nfInstances" yaml:"nfInstances"`
}

// NFTopologyStatus defines the observed state of NFTopology
type NFTopologyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NFTopology is the Schema for the nfTopology API
type NFTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NFTopologySpec   `json:"spec,omitempty"`
	Status NFTopologyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NFTopologyList contains a list of NFTopology
type NFTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NFTopology `json:"items"`
}
