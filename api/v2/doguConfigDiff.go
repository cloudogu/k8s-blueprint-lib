package v2

type CombinedDoguConfigDiff struct {
	DoguConfigDiff          DoguConfigDiff          `json:"doguConfigDiff,omitempty"`
	SensitiveDoguConfigDiff SensitiveDoguConfigDiff `json:"sensitiveDoguConfigDiff,omitempty"`
}

type DoguConfigValueState ConfigValueState

type DoguConfigDiff []DoguConfigEntryDiff
type DoguConfigEntryDiff struct {
	Key          string               `json:"key"`
	Actual       DoguConfigValueState `json:"actual"`
	Expected     DoguConfigValueState `json:"expected"`
	NeededAction ConfigAction         `json:"neededAction"`
}

// +kubebuilder:object:generate:=false

type SensitiveDoguConfigDiff = DoguConfigDiff

// +kubebuilder:object:generate:=false

type SensitiveDoguConfigEntryDiff = DoguConfigEntryDiff
