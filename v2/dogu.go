package v2

import (
	cescommons "github.com/cloudogu/ces-commons-lib/dogu"
	"github.com/cloudogu/cesapp-lib/core"
)

// Dogu defines a Dogu, its version, and the installation state in which it is supposed to be after a blueprint
// was applied.
type Dogu struct {
	// Name defines the name of the dogu, e.g. "official/postgresql"
	Name cescommons.QualifiedName
	// Version defines the version of the dogu that is to be installed. Must not be empty if the targetState is "present";
	// otherwise it is optional and is not going to be interpreted.
	Version core.Version
	// TargetState defines a state of installation of this dogu. Optional field, but defaults to "TargetStatePresent"
	TargetState TargetState
	// MinVolumeSize is the minimum storage of the dogu. This field is optional and can be nil to indicate that no
	// storage is needed.
	MinVolumeSize *VolumeSize
	// ReverseProxyConfig defines configuration for the ecosystem reverse proxy. This field is optional.
	ReverseProxyConfig ReverseProxyConfig
}
