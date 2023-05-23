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

// PackageRevisionReference is a temporary replica of PackageRevisionReference used for the
// ONE Summit
type PackageRevisionReference struct {
	// Namespace is the namespace for both the repository and package revision
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Repository is the name of the repository containing the package
	RepositoryName string `json:"repository"`

	// PackageName is the name of the package for the revision
	PackageName string `json:"packageName"`

	// Revision is the specific version number of the revision of the package
	Revision string `json:"revision"`
}

// NFClassSpec defines the reusable class struct for NF in general
type NFClassSpec struct {
    // Vendor is the name of the NF vendor
    Vendor      string `json:"vendor,omitempty" yaml:"vendor,omitempty"`

    // Version of the software version for this NF vendor's NFType
    Version     string `json:"version,omitempty" yaml:"version,omitempty"`

    // PackageRef is a reference to the upstream package used for this NF deployment
    PackageRef  PackageRevisionReference `json:"packageRef,omitempty" yaml:"packageRef,omitempty"`
}

// NFClassStatus defines the observed state of NFClass
type NFClassStatus struct {
    // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
    // Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=nfclasses,scope=Cluster

// NFClass is the Schema for the nfclasses API
type NFClass struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   NFClassSpec   `json:"spec,omitempty"`
    Status NFClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NFClassList contains a list of NFClass
type NFClassList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []NFClass `json:"items"`
}
