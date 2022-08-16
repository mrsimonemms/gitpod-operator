// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GitpodSpec defines the desired state of Gitpod
type GitpodSpec struct {
	//+kubebuilder:validation:Minimum=0
	// Size is the size of the Gitpod deployment
	Size int32 `json:"size"`
}

// GitpodStatus defines the observed state of Gitpod
type GitpodStatus struct {
	// Nodes are the names of the Gitpod pods
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Gitpod is the Schema for the gitpods API
type Gitpod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitpodSpec   `json:"spec,omitempty"`
	Status GitpodStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GitpodList contains a list of Gitpod
type GitpodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gitpod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Gitpod{}, &GitpodList{})
}
