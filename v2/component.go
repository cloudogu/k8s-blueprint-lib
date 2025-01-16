package v2

import (
	"github.com/Masterminds/semver/v3"
)

// Component represents a CES component (e.g. operators), its version, and the installation state in which it is supposed to be
// after a blueprint was applied.
type Component struct {
	// Name defines the name and namespace of the component. Must not be empty.
	Name QualifiedComponentName
	// Version defines the version of the package that is to be installed. Must not be empty if the targetState is
	// "present"; otherwise it is optional and is not going to be interpreted.
	Version *semver.Version
	// TargetState defines a state of installation of this package. Optional field, but defaults to "TargetStatePresent"
	TargetState TargetState
	// DeployConfig defines generic properties for the component. This field is optional.
	DeployConfig map[string]interface{}
}
