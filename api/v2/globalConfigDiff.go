package v2

// GlobalConfigDiff is a list of differences between the GlobalConfig in the Blueprint and the cluster state
type GlobalConfigDiff []ConfigEntryDiff

// GlobalConfigValueState represents either the actual or expected state of a global config key
type GlobalConfigValueState ConfigValueState
