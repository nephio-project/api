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

	commonv1alpha1 "github.com/nephio-project/api/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	// Lifecycle determines the deletion lifecycle policies the resource
	// will follow
	Lifecycle commonv1alpha1.Lifecycle `json:"lifecycle,omitempty"`
}

// TokenStatus defines the observed state of the Token
type TokenStatus struct {
	// ConditionedStatus provides the status of the Readiness using conditions
	// if the condition is true the other attributes in the status are meaningful
	ConditionedStatus `json:",inline" yaml:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="REPO_TOKEN_STATUS",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"

// Token is the Schema for the repository token API
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TokenSpec   `json:"spec,omitempty"`
	Status TokenStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TokenList contains a list of Repository tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}

// Token type metadata.
var (
	TokenKind             = reflect.TypeOf(Token{}).Name()
	TokenGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: TokenKind}.String()
	TokenKindAPIVersion   = TokenKind + "." + GroupVersion.String()
	TokenGroupVersionKind = GroupVersion.WithKind(TokenKind)
)
