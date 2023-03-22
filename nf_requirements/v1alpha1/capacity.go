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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
type Capacity struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec   CapacitySpec   `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status CapacityStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type CapacitySpec struct {
	// max uplink data
	MaxUplinkThroughput resource.Quantity `json:"maxUplinkThroughput,omitempty" yaml:"maxUplinkThroughput,omitempty"`
	// max downlink throughput
	MaxDownlinkThroughput resource.Quantity `json:"maxDownlinkThroughput,omitempty" yaml:"maxDownlinkThroughput,omitempty"`
	// max PDU sessions
	MaxSessions int `json:"maxSessions,omitempty" yaml:"maxSessions,omitempty"`
	// max Subscribers
	MaxSubscribers int `json:"maxSubscribers,omitempty" yaml:"maxSubscribers,omitempty"`
}

type CapacityStatus struct {
}
