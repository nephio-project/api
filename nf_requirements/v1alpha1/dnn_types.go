/*
Copyright 2023 Nephio.

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
	"k8s.io/apimachinery/pkg/types"
)

func (r *DataNetwork) GetNetworkInstance() types.NamespacedName {
	nsn := types.NamespacedName{}
	if r.Spec.NetworkInstance != nil {
		nsn.Name = r.Spec.NetworkInstance.Name
		nsn.Namespace = r.Spec.NetworkInstance.Namespace
	}
	return nsn
}

func (r *DataNetwork) GetPools() []*Pool {
	if r.Spec.Pools == nil {
		return []*Pool{}
	}
	return r.Spec.Pools
}

func (r *Pool) GetName() string {
	return r.Name
}

func (r *Pool) GetPrefixLength() uint8 {
	return r.PrefixLength
}
