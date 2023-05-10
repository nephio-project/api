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

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
}

// RepositoryStatus defines the observed state of Repository
type RepositoryStatus struct {
	// ConditionedStatus provides the status of the Readiness using conditions
	// if the condition is true the other attributes in the status are meaningful
	ConditionedStatus `json:",inline" yaml:",inline"`
	// URL is the url for the repo
	URL *string `json:"url,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="REPO_STATUS",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"

// Repository is the Schema for the repository API
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepositorySpec   `json:"spec,omitempty"`
	Status RepositoryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RepositoryList contains a list of Repositories
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}

// Repository type metadata.
var (
	RepositoryKind             = reflect.TypeOf(Repository{}).Name()
	RepositoryGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: RepositoryKind}.String()
	RepositoryKindAPIVersion   = RepositoryKind + "." + GroupVersion.String()
	RepositoryGroupVersionKind = GroupVersion.WithKind(RepositoryKind)
)
