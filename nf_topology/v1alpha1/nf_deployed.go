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

type NFConnectivity struct {
	// peer NF's Id (see NFInstance struct below)
	NeighborName string `json:"neighborName,omitempty" yaml:"neighborName,omitempty"`
}

// NFDeployedInstance defines an NF instance that is deployed
type NFDeployedInstance struct {
	// unique ID for this NF instance
	ID string `json:"id,omitempty" yaml:"id,omitempty"`
	// type of NF, example: amf, smf, upf
	NFType string `json:"nfType,omitempty" yaml:"nfType,omitempty"`
	// corresponding NFInstance name
	NFInstaceName string `json:"nfInstanceName,omitempty" yaml:"nfInstanceName,omitempty"`
	// list of connected NF instances to this NF instance
	Connectivities []NFConnectivity `json:"connectivities,omitempty" yaml:"connectivities,omitempty"`
}
