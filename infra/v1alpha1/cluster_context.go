/*
Copyright 2022 The Nephio Authors.

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

type CNIConfig struct {
	CNIType         string `json:"cniType"`
	MasterInterface string `json:"masterInterface"`
}

// ClusterContextSpec defines the desired state of ClusterContext
type ClusterContextSpec struct {
	// SiteCode identifies this cluster's location
	// +optional
	SiteCode *string `json:"siteCode,omitempty"`

	// CNIConfig contains the details of CNI configuration for this cluster
	// +optional
	CNIConfig *CNIConfig `json:"cniConfig,omitempty"`
}

// ClusterContextStatus defines the observed state of ClusterContext
type ClusterContextStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterContext is the Schema for the clustercontexts API
type ClusterContext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterContextSpec   `json:"spec,omitempty"`
	Status ClusterContextStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterContextList contains a list of ClusterContext
type ClusterContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterContext `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterContext{}, &ClusterContextList{})
}
