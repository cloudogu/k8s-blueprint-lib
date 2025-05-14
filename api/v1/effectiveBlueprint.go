package v1

import (
	"github.com/cloudogu/k8s-blueprint-lib/json/entities"
)

// EffectiveBlueprint describes an abstraction of CES components that should be absent or present within one or more CES
// instances after combining the blueprint with the blueprint mask.
//
// In general additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint API version.
type EffectiveBlueprint struct {
	// Dogus contains a set of exact dogu versions which should be present or absent in the CES instance after which this
	// blueprint was applied. Optional.
	Dogus []entities.TargetDogu `json:"dogus,omitempty"`
	// Components contains a set of exact component versions which should be present or absent in the CES instance after which
	// this blueprint was applied. Optional.
	Components []entities.TargetComponent `json:"components,omitempty"`
	// Config is used for ecosystem configuration to be applied.
	// Optional.
	Config entities.TargetConfig `json:"config,omitempty"`
}
