package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=bpm
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"
// +kubebuilder:gen

// BlueprintMask is the Schema for the blueprint mask API
type BlueprintMask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the BlueprintMask.
	// +required
	Spec BlueprintMaskSpec `json:"spec"`
	// Status defines the observed state of the BlueprintMask.
	// +optional
	Status *BlueprintMaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlueprintMaskList contains a list of BlueprintMask
type BlueprintMaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlueprintMask `json:"items"`
}

// BlueprintMaskSpec defines the desired state of BlueprintMask
type BlueprintMaskSpec struct {
	// BlueprintMaskManifest contains a list of dogus which should modify a dogu set in a blueprint.
	// +required
	*BlueprintMaskManifest `json:",inline"`
}

// BlueprintMaskStatus defines the observed state of BlueprintMask
type BlueprintMaskStatus struct {
	// Conditions shows the current state of the BlueprintMask
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}
