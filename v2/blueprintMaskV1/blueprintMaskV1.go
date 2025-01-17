package blueprintMaskV1

import (
	"github.com/cloudogu/blueprint-lib/bpcore"
)

// BlueprintMaskV1 describes an abstraction of CES components that should alter a blueprint definition before
// applying it to a CES system via a blueprint upgrade. The blueprint mask should not change the blueprint JSON file
// itself, but is applied to the information in it to generate a new effective blueprint.
//
// In general additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint mask API version.
type BlueprintMaskV1 struct {
	bpcore.GeneralBlueprintMask
	// ID is the unique name of the set over all components. This blueprint mask ID should be used to distinguish
	// from similar blueprint masks between humans in an easy way. Must not be empty.
	ID string `json:"blueprintMaskId"`
	// Dogus contains a set of dogus which alters the states of the dogus in the blueprint this mask is applied on.
	// The names and target states of all dogus must not be empty.
	Dogus []MaskTargetDogu `json:"dogus"`
}

// MaskTargetDogu defines a Dogu, its version, and the installation state in which it is supposed to be after a blueprint
// was applied for a blueprintMask.
type MaskTargetDogu struct {
	// Name defines the name of the dogu including its namespace, f. i. "official/nginx". Must not be empty. If you set another namespace than in the normal blueprint, a
	Name string `json:"name"`
	// Version defines the version of the dogu that is to be installed. This version is optional and overrides
	// the version of the dogu from the blueprint.
	Version string `json:"version"`
	// TargetState defines a state of installation of this dogu. Optional field, but defaults to "TargetStatePresent"
	TargetState bpcore.TargetState `json:"targetState"`
}
