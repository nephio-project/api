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
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestCapacityGetMaxUplinkThroughput(t *testing.T) {
	tests := map[string]struct {
		input Capacity
		want  resource.Quantity
	}{
		"TestGetMaxUplinkThroughputEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: resource.Quantity{},
		},
		"TestGetMaxUplinkThroughput": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxUplinkThroughput: resource.MustParse("5G"),
				},
			},
			want: resource.MustParse("5G"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxUplinkThroughput()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestGetMaxUplinkThroughput: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestCapacityGetMaxDownlinkThroughput(t *testing.T) {
	tests := map[string]struct {
		input Capacity
		want  resource.Quantity
	}{
		"TestGetMaxDownlinkThroughputEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: resource.Quantity{},
		},
		"TestGetMaxDownlinkThroughput": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxDownlinkThroughput: resource.MustParse("5G"),
				},
			},
			want: resource.MustParse("5G"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxDownlinkThroughput()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestGetMaxDownlinkThroughput: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestCapacityGetMaxSessions(t *testing.T) {
	tests := map[string]struct {
		input Capacity
		want  int
	}{
		"TestGetMaxSessionsEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: 0,
		},
		"TestGetMaxSessions": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxSessions: 10,
				},
			},
			want: 10,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxDownlinkThroughput()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestGetMaxSessions: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestCapacityGetMaxSubscribers(t *testing.T) {
	tests := map[string]struct {
		input Capacity
		want  int
	}{
		"TestGetMaxSubscribersEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: 0,
		},
		"TestGetMaxSubscribers": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxSubscribers: 10,
				},
			},
			want: 10,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxDownlinkThroughput()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestGetMaxSubscribers: -want, +got:\n%s", diff)
			}
		})
	}
}
