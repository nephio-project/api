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
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetBridgeDomainName(t *testing.T) {
	cases := map[string]struct {
		itfce        *Interface
		bdName       string
		selectorName string
		expected     string
	}{
		"Selector": {
			itfce:        &Interface{Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{},
			}},
			bdName:       "a",
			selectorName: "b",
			expected:     "a-b-bd",
		},
		"NoSelector": {
			itfce:        &Interface{},
			bdName:       "a",
			selectorName: "b",
			expected:     "a-bd",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.itfce.GetBridgeDomainName(tc.bdName, tc.selectorName)

			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("-want, +got:\n%s", diff)
			}
		})
	}
}
