package v3

// DoguDiff is the comparison of a Dogu's desired state vs. its cluster state.
// It contains the operation that needs to be done to achieve this desired state.
type DoguDiff struct {
	// +required
	Actual DoguDiffState `json:"actual"`
	// +required
	Expected DoguDiffState `json:"expected"`
	// +required
	NeededActions []DoguAction `json:"neededActions"`
}

// DoguDiffState is either the actual or desired state of a dogu in the cluster.
type DoguDiffState struct {
	// +required
	Namespace string `json:"namespace"`
	// +optional
	Version *string `json:"version,omitempty"`
	// +required
	Absent bool `json:"absent"`
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

const (
	// DoguActionInstall means the dogu is to be installed
	DoguActionInstall DoguAction = "install"
	// DoguActionUninstall means the dogu is to be uninstalled
	DoguActionUninstall DoguAction = "uninstall"
	// DoguActionUpgrade means an upgrade needs to be performed for the dogu
	DoguActionUpgrade DoguAction = "upgrade"
	// DoguActionDowngrade means a downgrade needs to be performed for the dogu
	DoguActionDowngrade DoguAction = "downgrade"
	// DoguActionSwitchNamespace means the dogu should be pulled from a different dogu registry namespace
	DoguActionSwitchNamespace DoguAction = "dogu namespace switch"
	// DoguActionUpdateReverseProxyConfig means the reverse proxy config of the dogu needs to be updated
	DoguActionUpdateReverseProxyConfig DoguAction = "update reverse proxy"
	// DoguActionUpdateResourceMinVolumeSize means the minimum volume size of the dogu needs to be changed
	DoguActionUpdateResourceMinVolumeSize DoguAction = "update resource minimum volume size"
	// DoguActionUpdateAdditionalMounts means the additional mounts should be updated for the dogu
	DoguActionUpdateAdditionalMounts DoguAction = "update additional mounts"
)
