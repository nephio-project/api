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
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Reconciling implies that the deployment is progressing.
	// Reconciliation for a deployment is considered when a new version
	// is adopted, when new pods scale up or old pods scale down, or when
	// required peering is in progress.
	// Condition name follows Kpt guidelines.
	Reconciling NFDeploymentConditionType = "Reconciling"
	// Deployment is unable to make progress towards Reconciliation.
	// Reasons could be Pod creation failure, Peering failure etc.
	// Condition name follows Kpt guidelines.
	Stalled NFDeploymentConditionType = "Stalled"
	// The Deployment is considered available when following conditions hold:
	// 1. At-least the minimal set of Pods are up and running for at-least
	//minReadySeconds.
	// 2. The Deployment is ready for required peering.
	Available NFDeploymentConditionType = "Available"
	// The Deployment is making progress towards peering on the required
	// interfaces. A successful peering implies that the NF is reachable by
	// the required peers and is a able to reach them.
	Peering NFDeploymentConditionType = "Peering"
	// The Deployment is available and has peered successfully on required
	// interfaces.
	// At this stage, the deployment is ready to serve requests.
	Ready NFDeploymentConditionType = "Ready"
)

type InterfaceConfig struct {
	// Name of the interface
	Name string `json:"name,omitempty"`
	// IPv4 CIDRs for this interface
	IPs []string `json:"ips,omitempty"`
	// IPv4 addresses for gateway for this interface
	GatewayIPs []string `json:"gateway,omitempty"`
	// VLAN for this interface such that workload can tag, especially in case of sr-iov
	// +optional
	Vlan *string `json:"vlan:omitempty"`
}

type UPFCapacity struct {
	UplinkThroughput   resource.Quantity `json:"uplinkThroughput,omitempty"`
	DownlinkThroughput resource.Quantity `json:"downlinkThroughput,omitempty"`
}

type N6InterfaceConfig struct {
	Interface InterfaceConfig `json:"interface"`
	// DNN is the Data Network Name
	DNN      string `json:"dnn,omitempty"`
	UEIPPool string `json:"ueIPPool,omitempty"`
}

// UPFDeploymentSpec specifies config parameters for UPF
type UPFDeploymentSpec struct {
	Capacity     UPFCapacity         `json:"capacity"`
	N3Interfaces []InterfaceConfig   `json:"n3Interfaces"`
	N4Interfaces []InterfaceConfig   `json:"n4Interfaces"`
	N6Interfaces []N6InterfaceConfig `json:"n6Interfaces"`
	// +optional
	N9Interfaces []InterfaceConfig `json:"n9Interfaces,omitempty"`
}

type UPFDeploymentStatus struct {
	// The generation observed by the deployment controller.
	ObservedGeneration int32 `json:"observedGeneration"`
	// Current service state of the UPF.
	Conditions []NFCondition `json:"conditions,omitempty"`
}

type NFCondition struct {
	// Type of deployment condition.
	Type NFDeploymentConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// The last time this condition was probed for.
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

type NFDeploymentConditionType string

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UPFDeployment is the Schema for the upfdeployments API
type UPFDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UPFDeploymentSpec   `json:"spec,omitempty"`
	Status UPFDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UPFDeploymentList contains a list of UPFDeployment
type UPFDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UPFDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UPFDeployment{}, &UPFDeploymentList{})
}
