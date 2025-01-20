package entities

type TargetComponent struct {
	// Name defines the name of the component including its distribution namespace, f. i. "k8s/k8s-dogu-operator". Must not be empty.
	Name string `json:"name"`
	// Version defines the version of the component that is to be installed. Must not be empty if the targetState is "present";
	// otherwise it is optional and is not going to be interpreted.
	Version string `json:"version"`
	// TargetState defines a state of installation of this component. Optional field, but defaults to "TargetStatePresent"
	TargetState string `json:"targetState"`
	// DeployConfig defines a generic property map for the component configuration. This field is optional.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	DeployConfig DeployConfig `json:"deployConfig,omitempty"`
}

type DeployConfig map[string]interface{}
