/*
Copyright 2023 Nephio.

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
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestDNNGetNetworkInstance(t *testing.T) {
	tests := map[string]struct {
		input DataNetwork
		want  types.NamespacedName
	}{
		"GetNetworkInstanceEmpty": {
			input: DataNetwork{
				Spec: DataNetworkSpec{},
			},
			want: types.NamespacedName{},
		},
		"GetNetworkInstanceName": {
			input: DataNetwork{
				Spec: DataNetworkSpec{
					NetworkInstance: &corev1.ObjectReference{
						Name: "a",
					},
				},
			},
			want: types.NamespacedName{
				Name: "a",
			},
		},
		"GetNetworkInstanceNameSpace": {
			input: DataNetwork{
				Spec: DataNetworkSpec{
					NetworkInstance: &corev1.ObjectReference{
						Namespace: "a",
					},
				},
			},
			want: types.NamespacedName{
				Namespace: "a",
			},
		},
		"GetNetworkInstanceNameSpaceName": {
			input: DataNetwork{
				Spec: DataNetworkSpec{
					NetworkInstance: &corev1.ObjectReference{
						Namespace: "a",
						Name:      "a",
					},
				},
			},
			want: types.NamespacedName{
				Namespace: "a",
				Name:      "a",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetNetworkInstance()

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestGetNetworkInstanc: -want, +got:\n%s", diff)
			}

		})
	}
}
