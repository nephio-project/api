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

// NephioClusterPackageConfigSpec provides the provider-independent
// Nephio configuration information needed by a Nephio cluster package
// to provision clusters. This is information about the Nephio platform
// configuration, not the cluster configuration.
type NephioClusterPackageConfigSpec struct {
	// NephioCatalogRepo is the name of the Porch blueprints repository
	// containing the Nephio system packages.
	// +kubebuilder:default=nephio-catalog
	// +optional
	NephioCatalogRepo *string `json:"nephioCatalogRepo,omitempty" yaml:"nephioCatalogRepo,omitempty"`
	// NephioMgmtRepo is the name of the Porch deployment repository for the
	// Nephio management cluster.
	// +kubebuilder:default=nephio-mgmt
	// +optional
	NephioMgmtRepo *string `json:"nephioMgmtRepo,omitempty" yaml:"nephioMgmtRepo,omitempty"`
	// NephioStagingRepo is the name of the Porch deployment repository
	// used for staging packages for new clusters.
	// +kubebuilder:default=nephio-staging
	// +optional
	NephioStagingRepo *string `json:"nephioStagingRepo,omitempty" yaml:"nephioStagingRepo,omitempty"`
}

// NephioClusterPackageConfigStatus defines the observed state of NephioClusterPackageConfig
type NephioClusterPackageConfigStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NephioClusterPackageConfig is the Schema for the Nephio Cluster Package Config API
type NephioClusterPackageConfig struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   NephioClusterPackageConfigSpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status NephioClusterPackageConfigStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NephioClusterPackageConfigList contains a list of Repositories
type NephioClusterPackageConfigList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Items           []NephioClusterPackageConfig `json:"items" yaml:"items"`
}

func init() {
	SchemeBuilder.Register(&NephioClusterPackageConfig{}, &NephioClusterPackageConfigList{})
}

// NephioClusterPackageConfig type metadata.
var (
	NephioClusterPackageConfigKind             = reflect.TypeOf(NephioClusterPackageConfig{}).Name()
	NephioClusterPackageConfigGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: NephioClusterPackageConfigKind}.String()
	NephioClusterPackageConfigKindAPIVersion   = NephioClusterPackageConfigKind + "." + GroupVersion.String()
	NephioClusterPackageConfigGroupVersionKind = GroupVersion.WithKind(NephioClusterPackageConfigKind)
)
