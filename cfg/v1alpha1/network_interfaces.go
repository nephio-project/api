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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// GetCondition returns the condition based on the condition kind
func (r *Network) GetCondition(t ConditionType) Condition {
	return r.Status.GetCondition(t)
}

// SetConditions sets the conditions on the resource. it allows for 0, 1 or more conditions
// to be set at once
func (r *Network) SetConditions(c ...Condition) {
	r.Status.SetConditions(c...)
}

// BuildNetworkConfig returns a Network from a client Object a crName and
// an Network Spec/Status
func BuildNetworkConfig(meta metav1.ObjectMeta, spec NetworkSpec, status NetworkStatus) *Network {
	return &Network{
		TypeMeta: metav1.TypeMeta{
			APIVersion: SchemeBuilder.GroupVersion.Identifier(),
			Kind:       NetworkKind,
		},
		ObjectMeta: meta,
		Spec:       spec,
		Status:     status,
	}
}
