package v2

// Config contains Dogu and Cloudogu EcoSystem specific configuration data which determine set-up and run
// behaviour respectively.
type Config struct {
	// Dogus contains Dogu specific configuration data which determine set-up and run behaviour.
	// +optional
	Dogus map[string]CombinedDoguConfig `json:"dogus,omitempty"`
	// Global contains EcoSystem specific configuration data which determine set-up and run behaviour.
	// +optional
	Global GlobalConfig `json:"global,omitempty"`
}

// CombinedDoguConfig states how a dogu should be configured.
type CombinedDoguConfig struct {
	// DoguConfig describes the config for this dogu.
	// Only use this config for entries, which need no special protection.
	// Otherwise, use SensitiveConfig instead.
	// +optional
	Config *DoguConfig `json:"config,omitempty"`
	// SensitiveConfig describes the sensitive config for this dogu.
	// In contrast to DoguConfig, this config can contain credentials, api keys licences etc.
	// +optional
	SensitiveConfig *SensitiveDoguConfig `json:"sensitiveConfig,omitempty"`
}

type DoguConfig presentAbsentConfig

type SensitiveDoguConfig PresentAbsentSensitiveConfig

type GlobalConfig presentAbsentConfig

type presentAbsentConfig struct {
	// Present describes config keys which should be present or should be set otherwise.
	// +optional
	Present map[string]string `json:"present,omitempty"`
	// Absent is a list of config keys which should be absent or get deleted.
	// +optional
	Absent []string `json:"absent,omitempty"`
}

// PresentAbsentSensitiveConfig describes which sensitive config should be present or absent.
// +optional
type PresentAbsentSensitiveConfig struct {
	// Present describes config keys which should be present or should be set otherwise.
	// +optional
	Present []SensitiveConfigEntry `json:"present,omitempty"`
	// Absent is a list of config keys which should be absent or get deleted.
	// +optional
	Absent []string `json:"absent,omitempty"`
}

type SensitiveConfigEntry struct {
	// Key is the name of the sensitive config key.
	// +required
	Key string `json:"key"`
	// SecretName is the name of the secret, from which the config key should be loaded.
	// The secret must be in the same namespace.
	// +required
	SecretName string `json:"secretName"`
	// SecretKey is the name of the key within the secret given by SecretName.
	// The value is used as the value for the sensitive config key.
	// +required
	SecretKey string `json:"secretKey"`
}
