package v2

type ConfigAction string

type ConfigValueState struct {
	// +optional
	Value  string `json:"value,omitempty"`
	Exists bool   `json:"exists"`
}
