package v2

// StateDiff is the result of comparing the Blueprint to the current cluster state.
// It describes what operations need to be done to achieve the desired state of the blueprint.
type StateDiff struct {
	// DoguDiffs maps simple dogu names to the determined diff.
	// +optional
	DoguDiffs map[string]DoguDiff `json:"doguDiffs,omitempty"`
	// ComponentDiffs maps simple component names to the determined diff.
	// +optional
	ComponentDiffs map[string]ComponentDiff `json:"componentDiffs,omitempty"`
	// DoguConfigDiffs maps simple dogu names to the determined config diff.
	// +optional
	DoguConfigDiffs map[string]ConfigDiff `json:"doguConfigDiffs,omitempty"`
	// GlobalConfigDiff is the difference between the GlobalConfig in the Blueprint and the cluster state.
	// +optional
	GlobalConfigDiff ConfigDiff `json:"globalConfigDiff,omitempty"`
}
