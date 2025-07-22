package v2

// Blueprint describes an abstraction of CES components which should be absent or present within one or more CES
// instances.
//
// In general, additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint API version.
type Blueprint struct {
	// Dogus contains a set of exact dogu versions which should be present or absent in the CES instance after which this
	// blueprint was applied. Optional.
	Dogus []Dogu `json:"dogus,omitempty"`
	// Components contains a set of exact component versions which should be present or absent in the CES instance after which
	// this blueprint was applied. Optional.
	Components []Component `json:"components,omitempty"`
	// Config is used for ecosystem configuration to be applied.
	// Optional.
	Config Config `json:"config,omitempty"`
}

// BlueprintMask describes changes to the given blueprint. Often customers use the same blueprint for multiple instances
// and use the blueprint mask to remove dogus from it.
//
// In general additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint API version.
type BlueprintMask struct {
	// Dogus contains a set of exact dogu with their versions which should be present or absent.
	// Optional.
	Dogus []Dogu `json:"dogus,omitempty"`
}
