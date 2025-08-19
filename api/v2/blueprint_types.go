package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type StatusPhase string

const (
	// StatusPhaseNew marks a newly created blueprint-CR.
	StatusPhaseNew StatusPhase = ""
	// StatusPhaseBlueprintApplicationFailed shows that the blueprint application failed.
	StatusPhaseBlueprintApplicationFailed StatusPhase = "blueprintApplicationFailed"
	// StatusPhaseBlueprintApplied indicates that the blueprint was applied but the ecosystem is not healthy yet.
	StatusPhaseBlueprintApplied StatusPhase = "blueprintApplied"
	// StatusPhaseFailed marks that an error occurred during processing of the blueprint.
	StatusPhaseFailed StatusPhase = "failed"
	// StatusPhaseCompleted marks the blueprint as successfully applied.
	StatusPhaseCompleted StatusPhase = "completed"
	// StatusPhaseRestartsTriggered indicates that a restart has been triggered for all Dogus that needed a restart.
	// Restarts are needed when the Dogu config changes.
	StatusPhaseRestartsTriggered StatusPhase = "restartsTriggered"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=bp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="The current status of the resource"
// +kubebuilder:printcolumn:name="DryRun",type="boolean",JSONPath=".spec.dryRun",description="Whether the resource is started as a dry run"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// Blueprint is the Schema for the blueprints API
type Blueprint struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the Blueprint.
	// +optional
	Spec BlueprintSpec `json:"spec,omitempty"`
	// Status defines the observed state of the Blueprint.
	// +optional
	Status BlueprintStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlueprintList contains a list of Blueprint
type BlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blueprint `json:"items"`
}

// BlueprintSpec defines the desired state of Blueprint
type BlueprintSpec struct {
	// Blueprint json with the desired state of the ecosystem.
	Blueprint BlueprintManifest `json:"blueprint"`
	// BlueprintMask json can further restrict the desired state from the blueprint.
	// +optional
	BlueprintMask BlueprintMask `json:"blueprintMask,omitempty"`
	// IgnoreDoguHealth lets the user execute the blueprint even if dogus are unhealthy at the moment.
	// +optional
	IgnoreDoguHealth bool `json:"ignoreDoguHealth,omitempty"`
	// IgnoreComponentHealth lets the user execute the blueprint even if components are unhealthy at the moment.
	// +optional
	IgnoreComponentHealth bool `json:"ignoreComponentHealth,omitempty"`
	// AllowDoguNamespaceSwitch lets the user switch the namespace of dogus in the blueprint mask
	// in comparison to the blueprint.
	// +optional
	AllowDoguNamespaceSwitch bool `json:"allowDoguNamespaceSwitch,omitempty"`
	// DryRun lets the user test a blueprint run to check if all attributes of the blueprint are correct and avoid a result with a failure state.
	// +optional
	DryRun bool `json:"dryRun,omitempty"`
}

// BlueprintStatus defines the observed state of Blueprint
type BlueprintStatus struct {
	// Conditions shows the current state of the blueprint
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Phase represents the processing state of the blueprint
	// +optional
	Phase StatusPhase `json:"phase,omitempty"`
	// EffectiveBlueprint is the blueprint after applying the blueprint mask.
	// +optional
	EffectiveBlueprint BlueprintManifest `json:"effectiveBlueprint,omitempty"`
	// StateDiff is the result of comparing the EffectiveBlueprint to the current cluster state.
	// It describes what operations need to be done to achieve the desired state of the blueprint.
	// +optional
	StateDiff StateDiff `json:"stateDiff,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Blueprint{}, &BlueprintList{})
}
