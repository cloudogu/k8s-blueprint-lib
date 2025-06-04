package v1

// GlobalConfigDiff is a list of differences between the GlobalConfig in the EffectiveBlueprint and the cluster state
type GlobalConfigDiff []GlobalConfigEntryDiff

// GlobalConfigValueState represents either the actual or expected state of a global config key
type GlobalConfigValueState ConfigValueState

// GlobalConfigEntryDiff contains the difference and the needed actions for a single global config key
type GlobalConfigEntryDiff struct {
	Key          string                 `json:"key"`
	Actual       GlobalConfigValueState `json:"actual"`
	Expected     GlobalConfigValueState `json:"expected"`
	NeededAction ConfigAction           `json:"neededAction"`
}
