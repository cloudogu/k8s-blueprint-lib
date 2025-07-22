package v2

type ConfigAction string

type ConfigValueState struct {
	Value  string `json:"value,omitempty"`
	Exists bool   `json:"exists"`
}
