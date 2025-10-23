package v3

type CombinedDoguConfigDiff struct {
	// +optional
	DoguConfigDiff DoguConfigDiff `json:"doguConfigDiff,omitempty"`
	// +optional
	SensitiveDoguConfigDiff SensitiveDoguConfigDiff `json:"sensitiveDoguConfigDiff,omitempty"`
}

type DoguConfigValueState ConfigValueState

type DoguConfigDiff []ConfigEntryDiff

// +kubebuilder:object:generate:=false

type SensitiveDoguConfigDiff = DoguConfigDiff

// +kubebuilder:object:generate:=false

type SensitiveDoguConfigEntryDiff = ConfigEntryDiff
