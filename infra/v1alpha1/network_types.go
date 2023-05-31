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
	"reflect"

	reqv1alpha1 "github.com/nephio-project/api/nf_requirements/v1alpha1"
	ipamv1alpha1 "github.com/nokia/k8s-ipam/apis/alloc/ipam/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	// Topology defines the topology to which this network applies
	Topology string `json:"topology" yaml:"topology"`
	// BridgeDomains define a set of logical ports that share the same
	// flooding or broadcast characteristics. Like a virtual LAN (VLAN),
	// a bridge domain spans one or more ports of multiple devices.
	BridgeDomains []BridgeDomain `json:"bridgeDomains,omitempty" yaml:"bridgeDomains,omitempty"`
	// RoutingTables defines a set of routes belonging to a given routing instance
	// Multiple routing tables are also called virtual routing instances. Each virtual
	// routing instance can hold overlapping IP information
	// A routing table supports both ipv4 and ipv6
	RoutingTables []RoutingTable `json:"routingTables,omitempty" yaml:"routingTables,omitempty"`
}

type BridgeDomain struct {
	// Name defines the name of the bridge domain
	Name string `json:"name" yaml:"name"`
	// Interfaces defines the interfaces belonging to the bridge domain
	Interfaces []Interface `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`
}

type RoutingTable struct {
	// Name defines the name of the routing table
	Name string `json:"name" yaml:"name"`
	// Interfaces defines the interfaces belonging to the routing table
	Interfaces []Interface `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`
	// Prefixes defines the prefixes belonging to the routing table
	Prefixes []ipamv1alpha1.Prefix `json:"prefixes" yaml:"prefixes"`
}

type Interface struct {
	// Kind defines the kind of interface. Attached to a routing table both interface and
	// bridgedomain interfaces are allowed. In a BridgeDomain only regular interfaces are allowed
	// +kubebuilder:validation:Enum=`interface`;`bridgedomain`
	// +kubebuilder:default=interface
	Kind InterfaceKind `json:"kind" yaml:"kind"`
	// BridgeDomainName defines the name of the bridgeDomain belonging to the interface
	BridgeDomainName *string `json:"bridgeDomainName,omitempty" yaml:"bridgeDomainName,omitempty"`
	// InterfaceName defines the name of the interface
	InterfaceName *string `json:"interfaceName,omitempty" yaml:"interfaceName,omitempty"`
	// NodeName defines the name of the node the interface belongs to interface
	NodeName *string `json:"nodeName,omitempty" yaml:"nodeName,omitempty"`
	// Selector defines the selector criterias for the interface selection
	Selector *metav1.LabelSelector `json:"selector,omitempty" yaml:"selector,omitempty"`
	// AttachmentType defines the interface attachement: vlan or none
	// +kubebuilder:validation:Enum=none;vlan
	AttachmentType reqv1alpha1.AttachmentType `json:"attachmentType,omitempty" yaml:"attachmentType,omitempty"`
}

// NetworkStatus defines the observed state of Network
type NetworkStatus struct {
	// ConditionedStatus provides the status of the Readiness using conditions
	// if the condition is true the other attributes in the status are meaningful
	ConditionedStatus `json:",inline" yaml:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="REPO_STATUS",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:resource:categories={nephio,network}

// Network is the Schema for the Network API
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSpec   `json:"spec,omitempty"`
	Status NetworkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkList contains a list of Repositories
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}

// Network type metadata.
var (
	NetworkKind             = reflect.TypeOf(Network{}).Name()
	NetworkGroupKind        = schema.GroupKind{Group: GroupVersion.Group, Kind: NetworkKind}.String()
	NetworkKindAPIVersion   = NetworkKind + "." + GroupVersion.String()
	NetworkGroupVersionKind = GroupVersion.WithKind(NetworkKind)
)
