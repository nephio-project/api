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

type AttachmentType string

// AttachmentTypeNone defines an untagged attachement (no VLAN)
const AttachmentTypeNone AttachmentType = "none"
// AttachmentTypeVLAN defines a tagged/vlan attachement
const AttachmentTypeVLAN AttachmentType = "vlan"

type CNIType string

// CNITypeSRIOV defines the sriov cni
const CNITypeSRIOV CNIType = "sriov"
// CNITypeIPVLAN defines the ipvlan cni
const CNITypeIPVLAN CNIType = "ipvlan"
// CNITypeMACVLAN defines the macvlan cni
const CNITypeMACVLAN CNIType = "macvlan"