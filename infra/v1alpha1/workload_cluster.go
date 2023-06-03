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
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"reflect"
)

const (
	errMissingWorkloadClusterSpec = "missing WorkloadCluster spec"
	errMissingClusterName         = "mandatory field ClusterName is missing from WorkloadCluster"
)

// WorkloadClusterSpec defines the desired state of WorkloadCluster
type WorkloadClusterSpec struct {
	// ClusterName is the unique name for this cluster
	ClusterName string `json:"clusterName,omitempty"`
	// CNIs defines the CNIs required for the workload cluster
	CNIs []string `json:"cnis,omitempty"`
	// MasterInterface define the master interface for secondary networking in the nodes
	// on the cluster
	MasterInterface *string `json:"masterInterface,omitempty"`
}

// WorkloadClusterStatus defines the observed state of WorkloadCluster
type WorkloadClusterStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WorkloadCluster is the Schema for the clustercontexts API
type WorkloadCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadClusterSpec   `json:"spec,omitempty"`
	Status WorkloadClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkloadClusterList contains a list of WorkloadCluster
type WorkloadClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkloadCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkloadCluster{}, &WorkloadClusterList{})
}

func (spec *WorkloadClusterSpec) Validate() error {
	if spec == nil {
		return fmt.Errorf("spec invalid: %s", errMissingWorkloadClusterSpec)
	}
	if spec.ClusterName == "" {
		return fmt.Errorf("spec invalid: %s", errMissingClusterName)
	}
	return nil
}

var (
	WorkloadClusterKind             = reflect.TypeOf(WorkloadCluster{}).Name()
	WorkloadClusterGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: WorkloadClusterKind}.String()
	WorkloadClusterKindAPIVersion   = WorkloadClusterKind + "." + GroupVersion.String()
	WorkloadClusterGroupVersionKind = GroupVersion.WithKind(WorkloadClusterKind)
)
