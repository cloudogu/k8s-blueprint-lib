package v1

// StateDiff is the result of comparing the EffectiveBlueprint to the current cluster state.
// It describes what operations need to be done to achieve the desired state of the blueprint.
type StateDiff struct {
	// DoguDiffs maps simple dogu names to the determined diff.
	DoguDiffs map[string]DoguDiff `json:"doguDiffs,omitempty"`
	// ComponentDiffs maps simple component names to the determined diff.
	ComponentDiffs map[string]ComponentDiff `json:"componentDiffs,omitempty"`
	// DoguConfigDiffs maps simple dogu names to the determined config diff.
	DoguConfigDiffs map[string]CombinedDoguConfigDiff `json:"doguConfigDiffs,omitempty"`
	// GlobalConfigDiff is the difference between the GlobalConfig in the EffectiveBlueprint and the cluster state.
	GlobalConfigDiff GlobalConfigDiff `json:"globalConfigDiff,omitempty"`
}
