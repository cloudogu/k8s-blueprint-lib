package v2

// BlueprintMask describes an abstraction of CES components that should alter a blueprint definition before
// applying it to a CES system via a blueprint upgrade. The blueprint mask does not change the blueprint
// itself, but is applied to the information in it to generate a new, effective blueprint.
type BlueprintMask struct {
	// Dogus contains a set of dogus which alters the states of the dogus in the blueprint this mask is applied on.
	// The names and target states of all dogus must not be empty.
	Dogus []MaskDogu
}
