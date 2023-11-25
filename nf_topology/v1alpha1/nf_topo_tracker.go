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
	nephiodeployv1alpha1 "github.com/nephio-project/api/workload/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NFConnectivity struct {
	// peer NF's Id (see NFInstance struct below)
	NeighborName string `json:"neighborName,omitempty" yaml:"neighborName,omitempty"`
}

// NFTopoTrackerInstance defines an NF instance that is deployed
type NFTopoTrackerInstance struct {
	// unique ID for this NF instance
	ID string `json:"id,omitempty" yaml:"id,omitempty"`
	// name of workload cluster where the NF instance is to be deployed
	ClusterName string `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	// type of NF, example: amf, smf, upf
	NFType string `json:"nfType,omitempty" yaml:"nfType,omitempty"`
	// flavor of NF, example: small, medium, large
	NFFlavor string `json:"nfFlavor,omitempty" yaml:"nfFlavor,omitempty"`
	// NF vendor name
	NFVendor string `json:"nfVendor,omitempty" yaml:"nfVendor,omitempty"`
	// the software version of this NF vendor's NFType
	NFVersion string `json:"nfVersion,omitempty" yaml:"nfVersion,omitempty"`
	// list of connected NF instances to this NF instance
	Connectivities []NFConnectivity `json:"connectivities,omitempty" yaml:"connectivities,omitempty"`
}

// NFTopoTrackerSpec defines list of deployed NF for Edge Status Aggregator to monitor
type NFTopoTrackerSpec struct {
	NFInstances []NFTopoTrackerInstance `json:"nfInstances,omitempty" yaml:"nfInstances,omitempty"`
}

// NFTopoTrackerStatus tracks the aggregated status of all the deployed NFs for this one deployment
type NFTopoTrackerStatus struct {
	// The generation observed by the deployment controller.
	ObservedGeneration int32 `json:"observedGeneration,omitempty"`

	// Total number of NFs targeted by this deployment
	TargetedNFs int32 `json:"targetedNFs,omitempty"`

	// Total number of NFs targeted by this deployment with a Ready Condition set.
	ReadyNFs int32 `json:"readyNFs,omitempty"`

	// Total number of NFs targeted by this deployment with an Available Condition set.
	AvailableNFs int32 `json:"availableNFs,omitempty"`

	// Total number of NFs targeted by this deployment with a Stalled Condition set.
	StalledNFs int32 `json:"stalledNFs,omitempty"`

	// Current service state of the UPF.
	Conditions []nephiodeployv1alpha1.NFDeploymentConditionType `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NfTopoTracker is the Schema for the tracking of the list of deployed NFs
type NFTopoTracker struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata"`

	Spec   NFTopoTrackerSpec   `json:"spec,omitempty"`
	Status NFTopoTrackerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NFTopoTrackerList contains a list of NfTopoTracker
type NFTopoTrackerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NFTopoTracker `json:"items"`
}
