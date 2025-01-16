package bpcore

// BlueprintApi is a string that contains a Blueprint API version identifier.
type BlueprintApi string

const (
	// V1 is the version 1 API identifier of Cloudogu EcoSystem blueprint mechanism.
	V1 BlueprintApi = "v1"
	// TestEmpty is a non-production, test-only API identifier of Cloudogu EcoSystem blueprint mechanism.
	TestEmpty BlueprintApi = "test/empty"
)

// GeneralBlueprint defines the minimum set to parse the blueprint API version string in order to select the right
// blueprint handling strategy. This is necessary in order to accommodate maximal changes in different blueprint API
// versions.
type GeneralBlueprint struct {
	// API is used to distinguish between different versions of the used API and impacts directly the interpretation of
	// this blueprint. Must not be empty.
	//
	// This field MUST NOT be MODIFIED or REMOVED because the API is paramount for distinguishing between different
	// blueprint version implementations.
	API BlueprintApi `json:"blueprintApi"`
}
