package v2

type Config struct {
	Dogus  map[string]CombinedDoguConfig `json:"dogus,omitempty"`
	Global GlobalConfig                  `json:"global,omitempty"`
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
