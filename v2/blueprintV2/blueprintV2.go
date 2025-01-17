package blueprintV2

import (
	"github.com/cloudogu/blueprint-lib/bpcore"
	"github.com/cloudogu/blueprint-lib/v2/entities"
)

// BlueprintV2 describes an abstraction of CES components that should be absent or present within one or more CES
// instances. When the same Blueprint is applied to two different CES instances it is required to leave two equal
// instances in terms of the components.
//
// In general additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint API version.
type BlueprintV2 struct {
	bpcore.GeneralBlueprint
	// Dogus contains a set of exact dogu versions, which should be present or absent
	// in the CES instance after this blueprint was applied.
	// Optional.
	Dogus []entities.TargetDogu `json:"dogus,omitempty"`
	// Components are a set of exact package versions,
	// which should be present or absent in the CES instance after which this blueprint was applied.
	// The packages must correspond to the used package manager.
	// Optional.
	Components []entities.TargetComponent `json:"components,omitempty"`
	// Config is used for ecosystem configuration to be applied.
	// Optional.
	Config entities.TargetConfig `json:"config,omitempty"`
}

// RegistryConfig contains hierarchically organized key-value data configuration data on how a Cloudogu EcoSystem or its
// dogus are supposed to be set-up and run.
