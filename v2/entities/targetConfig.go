package entities

// TargetConfig contains Dogu and Cloudogu EcoSystem specific configuration data which determine set-up and run
// behaviour respectively.
type TargetConfig struct {
	// Dogus contains Dogu specific configuration data which determine set-up and run behaviour.
	Dogus map[string]CombinedDoguConfig `json:"dogus,omitempty"`
	// Dogus contains Cloudogu EcoSystem specific configuration data which determine set-up and run behaviour.
	Global GlobalConfig `json:"global,omitempty"`
}

// CombinedDoguConfig contains configuration data of different sensitivity.
type CombinedDoguConfig struct {
	Config          DoguConfig          `json:"config,omitempty"`
	SensitiveConfig SensitiveDoguConfig `json:"sensitiveConfig,omitempty"`
}

type DoguConfig presentAbsentConfig

type SensitiveDoguConfig presentAbsentConfig

type GlobalConfig presentAbsentConfig

type presentAbsentConfig struct {
	// Present contains config keys that should be created (if they don't exist) or be updated with values populated by
	// the map.
	Present map[string]string `json:"present,omitempty"`
	// Absent contains config keys that should be removed if they exist. Not existing keys will be ignored.
	Absent []string `json:"absent,omitempty"`
}

type GlobalConfigKey string

type DoguConfigKey struct {
	DoguName string
	Key      string
}

type DoguConfigValue string
