package v2

type ConfigAction string

// ConfigValueState represents either the actual or expected state of a config key
type ConfigValueState struct {
	// +optional
	Value *string `json:"value,omitempty"`
	// +required
	Exists bool `json:"exists"`
}

// ConfigDiff is a list of differences between Config in the Blueprint and the cluster state
type ConfigDiff []ConfigEntryDiff

// ConfigEntryDiff contains the difference and the needed actions for a single config key
type ConfigEntryDiff struct {
	// +required
	Key string `json:"key"`
	// +required
	Actual ConfigValueState `json:"actual"`
	// +required
	Expected ConfigValueState `json:"expected"`
	// +required
	NeededAction ConfigAction `json:"neededAction"`
}
