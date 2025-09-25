package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//nolint:unused
//goland:noinspection GoUnusedConst
const (
	ConditionValid                = "Valid"
	ConditionExecutable           = "Executable"
	ConditionEcosystemHealthy     = "EcosystemHealthy"
	ConditionSelfUpgradeCompleted = "SelfUpgradeCompleted"
	ConditionCompleted            = "Completed"
	ConditionLastApplySucceeded   = "LastApplySucceeded"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=bp
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Display Name",type="string",JSONPath=".spec.displayName",description="The display name of the blueprint"
// +kubebuilder:printcolumn:name="Stopped",type="boolean",JSONPath=".spec.stopped",description="Whether the resource is started as a dry run"
// +kubebuilder:printcolumn:name="Valid",type="string",JSONPath=".status.conditions[?(@.type == 'Valid')].status",description="Whether the resource is valid in the current state"
// +kubebuilder:printcolumn:name="Executable",type="string",JSONPath=".status.conditions[?(@.type == 'Executable')].status",description="Whether the resource is executable in the current state"
// +kubebuilder:printcolumn:name="Ecosystem Healthy",type="string",JSONPath=".status.conditions[?(@.type == 'EcosystemHealthy')].status",description="Whether the ecosystem is healthy in the current state"
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=".status.conditions[?(@.type == 'Completed')].status",description="Whether the resource is completed in the current state"
// +kubebuilder:printcolumn:name="Last Apply Successful",type="string",JSONPath=".status.conditions[?(@.type == 'LastApplySucceeded')].status",description="Whether the last apply steps were successful"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the resource"

// Blueprint is the Schema for the blueprints API
type Blueprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the Blueprint.
	// +optional
	Spec *BlueprintSpec `json:"spec,omitempty"`
	// Status defines the observed state of the Blueprint.
	// +optional
	Status *BlueprintStatus `json:"status,omitempty"`
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
	// DisplayName is the name of the blueprint that will be shown in the UI.
	// +required
	DisplayName string `json:"displayName"`
	// Blueprint with the desired state of the ecosystem.
	// +required
	Blueprint BlueprintManifest `json:"blueprint"`
	// BlueprintMask can further restrict the desired state from the blueprint.
	// +optional
	BlueprintMask *BlueprintMask `json:"blueprintMask,omitempty"`
	// IgnoreDoguHealth lets the user execute the blueprint even if dogus are unhealthy at the moment.
	// +optional
	IgnoreDoguHealth *bool `json:"ignoreDoguHealth,omitempty"`
	// IgnoreComponentHealth lets the user execute the blueprint even if components are unhealthy at the moment.
	// +optional
	IgnoreComponentHealth *bool `json:"ignoreComponentHealth,omitempty"`
	// AllowDoguNamespaceSwitch lets the user switch the namespace of dogus in the blueprint mask
	// in comparison to the blueprint.
	// +optional
	AllowDoguNamespaceSwitch *bool `json:"allowDoguNamespaceSwitch,omitempty"`
	// Stopped lets the user stop the blueprint execution. The blueprint will still check if all attributes are correct and avoid a result with a failure state.
	// +optional
	Stopped *bool `json:"stopped,omitempty"`
}

// BlueprintStatus defines the observed state of Blueprint
type BlueprintStatus struct {
	// Conditions shows the current state of the blueprint
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// EffectiveBlueprint is the blueprint after applying the blueprint mask.
	// +optional
	EffectiveBlueprint *BlueprintManifest `json:"effectiveBlueprint,omitempty"`
	// StateDiff is the result of comparing the EffectiveBlueprint to the current cluster state.
	// It describes what operations need to be done to achieve the desired state of the blueprint.
	// +optional
	StateDiff *StateDiff `json:"stateDiff,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Blueprint{}, &BlueprintList{})
}
