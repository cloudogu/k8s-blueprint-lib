package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=bpm
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Valid",type="string",JSONPath=".status.conditions[?(@.type == 'Valid')].status",description="Whether the resource is valid in the current state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// BlueprintMaskResource is the Schema for the blueprint mask API
// TODO naming
type BlueprintMaskResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the BlueprintMaskResource.
	// +required
	Spec BlueprintMaskResourceSpec `json:"spec"`
	// Status defines the observed state of the BlueprintMaskResource.
	// +optional
	Status *BlueprintMaskResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlueprintMaskResourceList contains a list of BlueprintMaskResource
type BlueprintMaskResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlueprintMaskResource `json:"items"`
}

// BlueprintMaskResourceSpec defines the desired state of BlueprintMaskResource
type BlueprintMaskResourceSpec struct {
	// BlueprintMask contains a list of dogus which should modify a dogu set in a blueprint.
	// +required
	BlueprintMask
}

// BlueprintMaskResourceStatus defines the observed state of Blueprint
type BlueprintMaskResourceStatus struct {
	// Conditions shows the current state of the blueprint
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}
