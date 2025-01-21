package entities

import "encoding/json"

// TargetConfig contains Dogu and Cloudogu EcoSystem specific configuration data which determine set-up and run
// behaviour respectively.
type TargetConfig struct {
	// Dogus contains Dogu specific configuration data which determine set-up and run behaviour.
	Dogus DoguConfigMap `json:"dogus,omitempty"`
	// Dogus contains Cloudogu EcoSystem specific configuration data which determine set-up and run behaviour.
	Global GlobalConfig `json:"global,omitempty"`
}

type DoguConfigMap map[string]CombinedDoguConfig

func (in *DoguConfigMap) DeepCopy() *DoguConfigMap {
	out := new(DoguConfigMap)
	in.DeepCopyInto(out)
	return out
}

func (in *DoguConfigMap) DeepCopyInto(out *DoguConfigMap) {
	if out != nil {
		jsonStr, err := json.Marshal(in)
		if err != nil {
			return
		}
		err = json.Unmarshal(jsonStr, in)
		if err != nil {
			return
		}
	}
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

func (in TargetConfig) DeepCopyInto(out *TargetConfig) {
	if out != nil {
		out.Global = *in.Global.DeepCopy()
		out.Dogus = *in.Dogus.DeepCopy()
	}
}

func (in *GlobalConfig) DeepCopy() *GlobalConfig {
	out := new(GlobalConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *GlobalConfig) DeepCopyInto(out *GlobalConfig) {
	if out != nil {
		jsonStr, err := json.Marshal(in)
		if err != nil {
			return
		}
		err = json.Unmarshal(jsonStr, in)
		if err != nil {
			return
		}
	}
}
