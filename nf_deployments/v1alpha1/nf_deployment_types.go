/*
Copyright 2023 Copyright 2023 The Nephio Authors.

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
	nephioreqv1alpha1 "github.com/nephio-project/api/nf_requirements/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NFDeploymentSpec defines the characteristics of a deployment of a network function
type NFDeploymentSpec struct {
	// capacity defines the capacity characteristics of the NF deployment
	// +optional
	Capacity *nephioreqv1alpha1.CapacitySpec `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	// Interfaces defines the interfaces associated with the NF deployment
	// +optional
	Interfaces []*InterfaceConfig `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`
	// NetworkInstances defines the network instances associated with the NF deployment
	// +optional
	NetworkInstances []*NetworkInstance `json:"networkInstances,omitempty" yaml:"networkInstances,omitempty"`
	// configRef defines addiitonal configuration references the nf depends upon
	// +optional
	ConfigRef []*corev1.ObjectReference `json:"configRefs,omitempty" yaml:"configRefs,omitempty"`
}

// InterfaceConfig defines the configuration of the interface
type InterfaceConfig struct {
	// Name defines the name of the interface
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name" yaml:"name"`
	// IPv4Config defines the ipv4 configuration of the interface
	// +optional
	IPv4Config *IPv4Config `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
	// IPv6Config defines the ipv6 configuration of the interface
	// +optional
	IPv6Config *IPv6Config `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`
	// VlanID defines the specific vlan id associated on this interface
	// +optional
	VlanID uint16 `json:"vlanID,omitempty" yaml:"vlanID,omitempty"`
}

// IPv4Config defines the configuration parameters of an ipv4 interface or peer
type IPv4Config struct {
	// Address defines the IPv4 address and prefix length in CIDR notation
	// [IP prefix, range IPv4 with host bits]
	Address string `json:"address" yaml:"address"`
	// Gateway defines the IPv4 address associated to the interface as a gateway
	// +optional
	Gateway string `json:"gateway,omitempty" yaml:"gateway,omitempty"`
}

// IPv6Config defines the configuration parameters of an ipv6 interface or peer
type IPv6Config struct {
	// Address defines the IPv6 address and prefix length in CIDR notation
	// [IP prefix, range IPv6 with host bits]
	Address string `json:"address" yaml:"address"`
	// Gateway defines the IPv6 address associated to the interface as a gateway
	// +optional
	Gateway string `json:"gateway,omitempty" yaml:"gateway,omitempty"`
}

// A networkInstance is a Layer 3 forwarding construct
// such as a virtual routing and forwarding (VRF) instance,
type NetworkInstance struct {
	// Name defines the name of the network instance
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name" yaml:"name"`
	// interfaces defines the interfaces associated with the network instance
	// +optional
	Interfaces []string `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`
	// Peers defines the peer configuration associated with the network instance
	// +optional
	Peers []*PeerConfig `json:"peers,omitempty" yaml:"peers,omitempty"`
	// dnns defines the dnns assocated with the network instance
	// +optional
	Dnns []*Dnn `json:"dnns,omitempty" yaml:"dnns,omitempty"`
	// bgp defines the BGP configuration associated with the network instance
	// +optional
	Bgp *BGPConfig `json:"bgp,omitempty" yaml:"bgp,omitempty"`
}

// A Dnn defines the Data Network name defined by 3GPP
type PeerConfig struct {
	// Name defines the name of the data network
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name,omitempty"`
	// IPv4Config defines the ipv4 configuration of the peer
	// +optional
	IPv4Config *IPv4Config `json:"ipv4,omitempty" yaml:"ipv4,omitempty"`
	// IPv6Config defines the ipv6 configuration of the peer
	// +optional
	IPv6Config *IPv6Config `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`
}

// A Dnn defines the Data Network name defined by 3GPP
type Dnn struct {
	// Name defines the name of the data network
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name,omitempty"`
	// UEIPAddressPool defines the list of address pools associated with the dnn
	// +optional
	UEIPAddressPool []*IPAddressPool `json:"pool,omitempty" yaml:"pool,omitempty"`
}

type IPAddressPool struct {
	// Prefix defines the ip cidr in prefix notation. It is defines as a subnet
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix string `json:"prefix" yaml:"prefix"`
}

// BgpConfig specifies parameters for BGP related configuration for UPF and SMF
type BGPConfig struct {
	// RouterID defines the router ID of the bgp process
	RouterID string `json:"routerID" yaml:"prefix"`
	// ASNumber defines the AS number of the bgp process
	ASNumber int `json:"autonomousSystem" yaml:"autonomousSystem"`
	// BGPNeigbors defines the configuration of the BGP neighbor
	BGPNeigbors []*BGPNeighbor `json:"neighbors" yaml:"neighbors"`
}

type BGPNeighbor struct {
	// Address defines the IPv4 or IPv6 address of the BGP neighbor
	Address string `json:"address" yaml:"address"`
	// BGP interface name, MUST match the one use in InterfaceConfig
	Name string `json:"Name,omitempty"`
	// PeerAS defines the AS number of the bgp peer
	PeerAS int `json:"peerAS" yaml:"autonomousSystem"`
}

// NFDeploymentStatus defines the observed state of nf deployment
type NFDeploymentStatus struct {
	// The generation observed by the deployment controller.
	ObservedGeneration int32 `json:"observedGeneration"`
	// Current service state of the NF.
	Conditions []Condition `json:"conditions,omitempty"`
}

type Condition struct {
	// Type defines the type of the nf deployment condition.
	Type NFDeploymentConditionType `json:"type"`
	// Status defines the status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// LastProbeTime defines the last time this condition was probed for.
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// LastTransitionTime defines the last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Reason defines the reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Message defines a human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

type NFDeploymentConditionType string

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