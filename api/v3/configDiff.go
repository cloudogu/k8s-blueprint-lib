package v3

type ConfigAction string

const (
	// ConfigActionNone means that nothing is to do for this config key
	ConfigActionNone ConfigAction = "none"
	// ConfigActionSet means that the config key needs to be set as given
	ConfigActionSet ConfigAction = "set"
	// ConfigActionRemove means that the config key needs to be deleted
	ConfigActionRemove ConfigAction = "remove"
)

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
