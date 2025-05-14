package v1

// DoguDiff is the comparison of a Dogu's desired state vs. its cluster state.
// It contains the operation that needs to be done to achieve this desired state.
type DoguDiff struct {
	Actual        DoguDiffState `json:"actual"`
	Expected      DoguDiffState `json:"expected"`
	NeededActions []DoguAction  `json:"neededActions"`
}

// DoguDiffState is either the actual or desired state of a dogu in the cluster.
type DoguDiffState struct {
	Namespace          string             `json:"namespace,omitempty"`
	Version            string             `json:"version,omitempty"`
	InstallationState  string             `json:"installationState"`
	ResourceConfig     ResourceConfig     `json:"resourceConfig,omitempty"`
	ReverseProxyConfig ReverseProxyConfig `json:"reverseProxyConfig,omitempty"`
}

type ResourceConfig struct {
	MinVolumeSize string `json:"minVolumeSize,omitempty"`
}

type ReverseProxyConfig struct {
	MaxBodySize      string `json:"maxBodySize,omitempty"`
	RewriteTarget    string `json:"rewriteTarget,omitempty"`
	AdditionalConfig string `json:"additionalConfig,omitempty"`
}

// DoguAction is the action that needs to be done for a dogu
// to achieve the desired state in the cluster.
type DoguAction string
