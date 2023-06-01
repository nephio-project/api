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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NFDeployedConnectivity struct {
    // peer NF's Id (see NFInstance struct below)
    NeighborName string `json:"neighborName,omitempty" yaml:"neighborName,omitempty"`
}

// NFDeployedInstance defines an NF instance that is deployed
type NFDeployedInstance struct {
    // unique ID for this NF instance
    Id             string         `json:"id,omitempty" yaml:"id,omitempty"`
    // name of workload cluster where the NF instance is to be deployed
    ClusterName    string         `json:"clustername,omitempty" yaml:"clustername,omitempty"`
    // type of NF, example: amf, smf, upf
    NFType         string         `json:"nftype,omitempty" yaml:"nftype,omitempty"`
    // flavor of NF, example: small, medium, large
    NFFlavor       string         `json:"nfflavor,omitempty" yaml:"nfflavor,omitempty"`
    // NF vendor name
    NFVendor       string         `json:"nfvendor,omitempty" yaml:"nfvendor,omitempty"`
    // the software version of this NF vendor's NFType
    NFVersion      string         `json:"nfversion,omitempty" yaml:"nfversion,omitempty"`
    // list of connected NF instances to this NF instance
    Connectivities []NFDeployedConnectivity `json:"connectivities,omitempty" yaml:"connectivities,omitempty"`
}

// NFDeployedSpec defines list of deployed NF for Edge Status Aggregator to monitor
type NFDeployedSpec struct {
    NFInstances     []NFDeployedInstance `json:"nfinstances,omitempty" yaml:"nfinstances,omitempty"`
}

// NFDeployedStatus tracks the aggregated status of all the deployed NFs for this one deployment
type NFDeployedStatus struct {
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
    Conditions []NFDeploymentConditionType `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NfDeployed is the Schema for the tracking of the list of deployed NFs
type NFDeployed struct {
    metav1.TypeMeta   `json:",inline" yaml:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata"`

    Spec   NFDeployedSpec   `json:"spec,omitempty"`
    Status NFDeployedStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NFDeployedList contains a list of NfDeployed
type NFDeployedList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []NFDeployed `json:"items"`
}
