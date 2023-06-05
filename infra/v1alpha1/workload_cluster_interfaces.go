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

import "fmt"

const (
	errMissingWorkloadClusterSpec            = "missing WorkloadCluster spec"
	errMissingClusterName                    = "mandatory field ClusterName is missing from WorkloadCluster"
	errWorkloadClusterrNotInitialized        = "workload cluster not initialized"
	errMissingWorkloadClusterMasterInterface = "mandatory field MasterInterface is missing from ClusterContext"
	errMissingWorkloadClusterCNIConfig       = "mandatory field CNIConfig is missing from ClusterContext"
)

func (spec *WorkloadClusterSpec) Validate() error {
	if spec == nil {
		return fmt.Errorf("spec invalid: %s", errMissingWorkloadClusterSpec)
	}
	if spec.ClusterName == "" || spec.ClusterName == "empty" {
		return fmt.Errorf("spec invalid: %s", errMissingClusterName)
	}
	if len(spec.CNIs) == 0 {
		return fmt.Errorf("spec invalid: %s", errMissingWorkloadClusterCNIConfig)
	}
	if spec.MasterInterface == nil {
		return fmt.Errorf("spec invalid: %s", errMissingWorkloadClusterMasterInterface)
	}
	return nil
}
