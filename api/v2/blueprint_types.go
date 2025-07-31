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
	// StatusPhaseStaticallyValidated marks the given blueprint spec as validated.
	StatusPhaseStaticallyValidated StatusPhase = "staticallyValidated"
	// StatusPhaseValidated marks the given blueprint spec as validated.
	StatusPhaseValidated StatusPhase = "validated"
	// StatusPhaseEffectiveBlueprintGenerated marks that the effective blueprint was generated out of the blueprint and the mask.
	StatusPhaseEffectiveBlueprintGenerated StatusPhase = "effectiveBlueprintGenerated"
	// StatusPhaseStateDiffDetermined marks that the diff to the ecosystem state was successfully determined.
	StatusPhaseStateDiffDetermined StatusPhase = "stateDiffDetermined"
	// StatusPhaseInvalid marks the given blueprint spec is semantically incorrect.
	StatusPhaseInvalid StatusPhase = "invalid"
	// StatusPhaseEcosystemHealthyUpfront marks that all currently installed dogus are healthy.
	StatusPhaseEcosystemHealthyUpfront StatusPhase = "ecosystemHealthyUpfront"
	// StatusPhaseEcosystemUnhealthyUpfront marks that some currently installed dogus are unhealthy.
	StatusPhaseEcosystemUnhealthyUpfront StatusPhase = "ecosystemUnhealthyUpfront"
	// StatusPhaseBlueprintApplicationPreProcessed shows that all pre-processing steps for the blueprint application
	// were successful.
	StatusPhaseBlueprintApplicationPreProcessed StatusPhase = "blueprintApplicationPreProcessed"
	// StatusPhaseAwaitSelfUpgrade marks that the blueprint operator waits for termination for a self upgrade.
	StatusPhaseAwaitSelfUpgrade StatusPhase = "awaitSelfUpgrade"
	// StatusPhaseSelfUpgradeCompleted marks that the blueprint operator itself got successfully upgraded.
	StatusPhaseSelfUpgradeCompleted StatusPhase = "selfUpgradeCompleted"
	// StatusPhaseInProgress marks that the blueprint is currently being processed.
	StatusPhaseInProgress StatusPhase = "inProgress"
	// StatusPhaseBlueprintApplicationFailed shows that the blueprint application failed.
	StatusPhaseBlueprintApplicationFailed StatusPhase = "blueprintApplicationFailed"
	// StatusPhaseBlueprintApplied indicates that the blueprint was applied but the ecosystem is not healthy yet.
	StatusPhaseBlueprintApplied StatusPhase = "blueprintApplied"
	// StatusPhaseEcosystemHealthyAfterwards shows that the ecosystem got healthy again after applying the blueprint.
	StatusPhaseEcosystemHealthyAfterwards StatusPhase = "ecosystemHealthyAfterwards"
	// StatusPhaseEcosystemUnhealthyAfterwards shows that the ecosystem got not healthy again after applying the blueprint.
	StatusPhaseEcosystemUnhealthyAfterwards StatusPhase = "ecosystemUnhealthyAfterwards"
	// StatusPhaseFailed marks that an error occurred during processing of the blueprint.
	StatusPhaseFailed StatusPhase = "failed"
	// StatusPhaseCompleted marks the blueprint as successfully applied.
	StatusPhaseCompleted StatusPhase = "completed"
	// StatusPhaseApplyEcosystemConfig indicates that the apply ecosystem config phase is active.
	StatusPhaseApplyEcosystemConfig StatusPhase = "applyEcosystemConfig"
	// StatusPhaseApplyEcosystemConfigFailed indicates that the phase to apply ecosystem config failed.
	StatusPhaseApplyEcosystemConfigFailed StatusPhase = "applyEcosystemConfigFailed"
	// StatusPhaseEcosystemConfigApplied indicates that the phase to apply ecosystem config succeeded.
	StatusPhaseEcosystemConfigApplied StatusPhase = "ecosystemConfigApplied"
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
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the Blueprint.
	Spec BlueprintSpec `json:"spec,omitempty"`
	// Status defines the observed state of the Blueprint.
	Status BlueprintStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlueprintList contains a list of Blueprint
type BlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blueprint `json:"items"`
}

// BlueprintSpec defines the desired state of Blueprint
type BlueprintSpec struct {
	// Blueprint json with the desired state of the ecosystem.
	Blueprint BlueprintManifest `json:"blueprint"`
	// BlueprintMask json can further restrict the desired state from the blueprint.
	BlueprintMask BlueprintMask `json:"blueprintMask,omitempty"`
	// IgnoreDoguHealth lets the user execute the blueprint even if dogus are unhealthy at the moment.
	IgnoreDoguHealth bool `json:"ignoreDoguHealth,omitempty"`
	// IgnoreComponentHealth lets the user execute the blueprint even if components are unhealthy at the moment.
	IgnoreComponentHealth bool `json:"ignoreComponentHealth,omitempty"`
	// AllowDoguNamespaceSwitch lets the user switch the namespace of dogus in the blueprint mask
	// in comparison to the blueprint.
	AllowDoguNamespaceSwitch bool `json:"allowDoguNamespaceSwitch,omitempty"`
	// DryRun lets the user test a blueprint run to check if all attributes of the blueprint are correct and avoid a result with a failure state.
	DryRun bool `json:"dryRun,omitempty"`
}

// BlueprintStatus defines the observed state of Blueprint
type BlueprintStatus struct {
	// Phase represents the processing state of the blueprint
	Phase StatusPhase `json:"phase,omitempty"`
	// EffectiveBlueprint is the blueprint after applying the blueprint mask.
	EffectiveBlueprint BlueprintManifest `json:"effectiveBlueprint,omitempty"`
	// StateDiff is the result of comparing the EffectiveBlueprint to the current cluster state.
	// It describes what operations need to be done to achieve the desired state of the blueprint.
	StateDiff StateDiff `json:"stateDiff,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Blueprint{}, &BlueprintList{})
}
