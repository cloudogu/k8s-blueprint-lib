package entities

import (
	"encoding/json"
	"fmt"
)

type TargetComponent struct {
	// Name defines the name of the component including its distribution namespace, f. i. "k8s/k8s-dogu-operator". Must not be empty.
	Name string `json:"name"`
	// Version defines the version of the component that is to be installed. Must not be empty if the targetState is "present";
	// otherwise it is optional and is not going to be interpreted.
	Version string `json:"version"`
	// TargetState defines a state of installation of this component. Optional field, but defaults to "TargetStatePresent"
	TargetState string `json:"targetState,omitempty"`
	// DeployConfig defines a generic property map for the component configuration. This field is optional.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	DeployConfig DeployConfig `json:"deployConfig,omitempty"`
}

func (in TargetComponent) DeepCopyInto(out *TargetComponent) {
	if out != nil {
		out.Name = in.Name
		out.Version = in.Version
		out.TargetState = in.TargetState
		out.DeployConfig = *in.DeployConfig.DeepCopy()
	}
}

type DeployConfig map[string]interface{}

func (in *DeployConfig) DeepCopy() *DeployConfig {
	out := new(DeployConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *DeployConfig) DeepCopyInto(out *DeployConfig) {
	if out != nil {
		jsonStr, err := json.Marshal(in)
		if err != nil {
			panic(fmt.Errorf("error unmarshaling DeployConfig: %w", err))
		}
		err = json.Unmarshal(jsonStr, out)
		if err != nil {
			panic(fmt.Errorf("error unmarshaling DeployConfig: %w", err))
		}
	}
}
