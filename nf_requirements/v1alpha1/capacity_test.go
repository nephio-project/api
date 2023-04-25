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
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestCapacityGetMaxUplinkThroughput(t *testing.T) {
	res := resource.MustParse("5G")

	tests := map[string]struct {
		input Capacity
		want  *resource.Quantity
	}{
		"TestGetMaxUplinkThroughputEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: nil,
		},
		"TestGetMaxUplinkThroughput": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxUplinkThroughput: &res,
				},
			},
			want: &res,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxUplinkThroughput()
			switch {
			case got == nil && tc.want != nil:
				t.Errorf("TestCapacityGetMaxUplinkThroughput: -want %v, +got: nil\n", *tc.want)
			case got != nil && tc.want == nil:
				t.Errorf("TestCapacityGetMaxUplinkThroughput: -want nil, +got: %v\n", *got)
			case got != nil && tc.want != nil:
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("TestCapacityGetMaxUplinkThroughput: -want, +got:\n%s", diff)
				}
			default:
				// got == nil and want == nil ok
			}
		})
	}
}

func TestCapacityGetMaxDownlinkThroughput(t *testing.T) {
	res := resource.MustParse("5G")

	tests := map[string]struct {
		input Capacity
		want  *resource.Quantity
	}{
		"TestGetMaxDownlinkThroughputEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: nil,
		},
		"TestGetMaxDownlinkThroughput": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxDownlinkThroughput: &res,
				},
			},
			want: &res,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxDownlinkThroughput()
			switch {
			case got == nil && tc.want != nil:
				t.Errorf("TestCapacityGetMaxDownlinkThroughput: -want %v, +got: nil\n", *tc.want)
			case got != nil && tc.want == nil:
				t.Errorf("TestCapacityGetMaxDownlinkThroughput: -want nil, +got: %v\n", *got)
			case got != nil && tc.want != nil:
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("TestCapacityGetMaxDownlinkThroughput: -want, +got:\n%s", diff)
				}
			default:
				// got == nil and want == nil ok
			}
		})
	}
}

func TestCapacityGetMaxSessions(t *testing.T) {
	maxSessions := 10

	tests := map[string]struct {
		input Capacity
		want  *int
	}{
		"TestGetMaxSessionsEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: nil,
		},
		"TestGetMaxSessions": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxSessions: &maxSessions,
				},
			},
			want: &maxSessions,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxSessions()
			switch {
			case got == nil && tc.want != nil:
				t.Errorf("TestCapacityGetMaxSessions: -want %v, +got: nil\n", *tc.want)
			case got != nil && tc.want == nil:
				t.Errorf("TestCapacityGetMaxSessions: -want nil, +got: %v\n", *got)
			case got != nil && tc.want != nil:
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("TestCapacityGetMaxSessions: -want, +got:\n%s", diff)
				}
			default:
				// got == nil and want == nil ok
			}
		})
	}
}

func TestCapacityGetMaxSubscribers(t *testing.T) {
	maxSubscribers := 10

	tests := map[string]struct {
		input Capacity
		want  *int
	}{
		"TestGetMaxSubscribersEmpty": {
			input: Capacity{
				Spec: CapacitySpec{},
			},
			want: nil,
		},
		"TestGetMaxSubscribers": {
			input: Capacity{
				Spec: CapacitySpec{
					MaxSubscribers: &maxSubscribers,
				},
			},
			want: &maxSubscribers,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.GetMaxSubscribers()
			switch {
			case got == nil && tc.want != nil:
				t.Errorf("TestCapacityGetMaxSubscribers: -want %v, +got: nil\n", *tc.want)
			case got != nil && tc.want == nil:
				t.Errorf("TestCapacityGetMaxSubscribers: -want nil, +got: %v\n", *got)
			case got != nil && tc.want != nil:
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("TestCapacityGetMaxSubscribers: -want, +got:\n%s", diff)
				}
			default:
				// got == nil and want == nil ok
			}
		})
	}
}
