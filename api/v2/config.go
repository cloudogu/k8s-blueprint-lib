package v2

// Config contains Dogu and Cloudogu EcoSystem specific configuration data which determine set-up and run
// behaviour respectively.
type Config struct {
	// Dogus contains Dogu specific configuration data which determine set-up and run behaviour.
	Dogus map[string]CombinedDoguConfig `json:"dogus,omitempty"`
	// Dogus contains EcoSystem specific configuration data which determine set-up and run behaviour.
	Global GlobalConfig `json:"global,omitempty"`
}

type CombinedDoguConfig struct {
	Config          DoguConfig          `json:"config,omitempty"`
	SensitiveConfig SensitiveDoguConfig `json:"sensitiveConfig,omitempty"`
}

type DoguConfig presentAbsentConfig

type SensitiveDoguConfig presentAbsentConfig

type GlobalConfig presentAbsentConfig

type presentAbsentConfig struct {
	Present map[string]string `json:"present,omitempty"`
	Absent  []string          `json:"absent,omitempty"`
}
