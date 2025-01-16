package v2

import (
	cescommons "github.com/cloudogu/ces-commons-lib/dogu"
	"golang.org/x/exp/maps"
)

// CensorValue is the value for censoring sensitive blueprint configuration data.
const CensorValue = "*****"

type Config struct {
	Dogus  map[cescommons.SimpleName]CombinedDoguConfig
	Global GlobalConfig
}

type CombinedDoguConfig struct {
	DoguName        cescommons.SimpleName
	Config          DoguConfig
	SensitiveConfig SensitiveDoguConfig
}

type DoguConfig struct {
	Present map[DoguConfigKey]DoguConfigValue
	Absent  []DoguConfigKey
}

type SensitiveDoguConfig = DoguConfig

type GlobalConfig struct {
	Present map[GlobalConfigKey]GlobalConfigValue
	Absent  []GlobalConfigKey
}

func (config GlobalConfig) GetGlobalConfigKeys() []GlobalConfigKey {
	var keys []GlobalConfigKey
	keys = append(keys, maps.Keys(config.Present)...)
	keys = append(keys, config.Absent...)
	return keys
}

func (config Config) GetDoguConfigKeys() []DoguConfigKey {
	var keys []DoguConfigKey
	for _, doguConfig := range config.Dogus {
		keys = append(keys, maps.Keys(doguConfig.Config.Present)...)
		keys = append(keys, doguConfig.Config.Absent...)
	}
	return keys
}

func (config Config) GetSensitiveDoguConfigKeys() []SensitiveDoguConfigKey {
	var keys []SensitiveDoguConfigKey
	for _, doguConfig := range config.Dogus {
		keys = append(keys, maps.Keys(doguConfig.SensitiveConfig.Present)...)
		keys = append(keys, doguConfig.SensitiveConfig.Absent...)
	}
	return keys
}

// GetDogusWithChangedConfig returns a list of dogus for which possible config changes are needed.
func (config Config) GetDogusWithChangedConfig() []cescommons.SimpleName {
	var dogus []cescommons.SimpleName
	for dogu, doguConfig := range config.Dogus {
		if len(doguConfig.Config.Present) != 0 || len(doguConfig.Config.Absent) != 0 {
			dogus = append(dogus, dogu)
		}
	}
	return dogus
}

// GetDogusWithChangedSensitiveConfig returns a list of dogus for which possible sensitive config changes are needed.
func (config Config) GetDogusWithChangedSensitiveConfig() []cescommons.SimpleName {
	var dogus []cescommons.SimpleName
	for dogu, doguConfig := range config.Dogus {
		if len(doguConfig.SensitiveConfig.Present) != 0 || len(doguConfig.SensitiveConfig.Absent) != 0 {
			dogus = append(dogus, dogu)
		}
	}
	return dogus
}

// censorValues censors all sensitive configuration data to make them unrecognisable.
func (config Config) censorValues() Config {
	for _, doguConfig := range config.Dogus {
		for k := range doguConfig.SensitiveConfig.Present {
			doguConfig.SensitiveConfig.Present[k] = CensorValue
		}
	}
	return config
}
