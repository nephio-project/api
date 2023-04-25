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

import "k8s.io/apimachinery/pkg/api/resource"

// GetMaxUplinkThroughput return the max uplinkthroughput
func (r *Capacity) GetMaxUplinkThroughput() *resource.Quantity {
	return r.Spec.MaxUplinkThroughput
}

// GetMaxDownlinkThroughput return the max downlinkthroughput
func (r *Capacity) GetMaxDownlinkThroughput() *resource.Quantity {
	return r.Spec.MaxDownlinkThroughput
}

// GetMaxSessions returns the max sessions
func (r *Capacity) GetMaxSessions() *int {
	return r.Spec.MaxSessions
}

// GetMaxSubscribers returns the max subscribers
func (r *Capacity) GetMaxSubscribers() *int {
	return r.Spec.MaxSubscribers
}
