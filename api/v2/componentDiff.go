package v2

// ComponentDiff is the comparison of a Component's desired state vs. its cluster state.
// It contains the operation that needs to be done to achieve this desired state.
type ComponentDiff struct {
	// Actual contains the component's state in the current system.
	// +required
	Actual ComponentDiffState `json:"actual"`
	// Expected contains the desired component's target state.
	// +required
	Expected ComponentDiffState `json:"expected"`
	// NeededActions contains the necessary actions to apply the wanted state for this component.
	// +required
	NeededActions []ComponentAction `json:"neededActions"`
}

// ComponentDiffState is either the actual or desired state of a component in the cluster. The fields will be used to
// determine the kind of changed if there is a drift between actual or desired state.
type ComponentDiffState struct {
	// Namespace is part of the address under which the component will be obtained. This namespace must NOT
	// to be confused with the K8s cluster namespace.
	// +optional
	Namespace *string `json:"distributionNamespace,omitempty"`
	// Version contains the component's version.
	// +optional
	Version *string `json:"version,omitempty"`
	// InstallationState contains the component's installation state. Such a state correlate with the domain Actions:
	//
	//  - domain.ActionInstall
	//  - domain.ActionUninstall
	//  - and so on
	// +required
	InstallationState string `json:"installationState"`
	// DeployConfig contains generic properties for the component.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	// +optional
	DeployConfig DeployConfig `json:"deployConfig,omitempty"`
}

// ComponentAction is the action that needs to be done for a component
// to achieve the desired state in the cluster.
type ComponentAction string
