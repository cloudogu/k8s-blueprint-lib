package v2

// DoguDiff is the comparison of a Dogu's desired state vs. its cluster state.
// It contains the operation that needs to be done to achieve this desired state.
type DoguDiff struct {
	Actual        DoguDiffState `json:"actual"`
	Expected      DoguDiffState `json:"expected"`
	NeededActions []DoguAction  `json:"neededActions"`
}

// DoguDiffState is either the actual or desired state of a dogu in the cluster.
type DoguDiffState struct {
	// +optional
	Namespace *string `json:"namespace,omitempty"`
	// +optional
	Version           *string `json:"version,omitempty"`
	InstallationState string  `json:"installationState"`
	// +optional
	ResourceConfig *ResourceConfig `json:"resourceConfig,omitempty"`
	// +optional
	ReverseProxyConfig *ReverseProxyConfig `json:"reverseProxyConfig,omitempty"`
	// +optional
	AdditionalMounts []AdditionalMount `json:"additionalMounts,omitempty" patchStrategy:"replace"`
}

// DoguAction is the action that needs to be done for a dogu
// to achieve the desired state in the cluster.
type DoguAction string
