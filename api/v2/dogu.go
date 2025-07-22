package v2

import (
	"encoding/json"
	"fmt"
)

// Dogu defines a Dogu, its version, and the installation state in which it is supposed to be after a blueprint
// was applied.
type Dogu struct {
	// Name defines the name of the dogu including its namespace, f. i. "official/nginx". Must not be empty.
	Name string `json:"name"`
	// Version defines the version of the dogu that is to be installed. Must not be empty if the targetState is "present";
	// otherwise it is optional and is not going to be interpreted.
	Version string `json:"version"`
	// Absent defines if the dogu should be absent in the ecosystem. Defaults to false.
	Absent string `json:"absent,omitempty"`
	// PlatformConfig defines infrastructure configuration around the dogu, such as reverse proxy config, volume size etc.
	PlatformConfig PlatformConfig `json:"platformConfig,omitempty"`
}

func (in *Dogu) DeepCopy() *Dogu {
	out := new(Dogu)
	in.DeepCopyInto(out)
	return out
}

func (in *Dogu) DeepCopyInto(out *Dogu) {
	if out != nil {
		jsonStr, err := json.Marshal(in)
		if err != nil {
			panic(fmt.Errorf("error marshaling Dogu: %w", err))
		}
		err = json.Unmarshal(jsonStr, out)
		if err != nil {
			panic(fmt.Errorf("error unmarshaling Dogu: %w", err))
		}
	}
}

type ResourceConfig struct {
	MinVolumeSize string `json:"minVolumeSize,omitempty"`
}

type ReverseProxyConfig struct {
	MaxBodySize      string `json:"maxBodySize,omitempty"`
	RewriteTarget    string `json:"rewriteTarget,omitempty"`
	AdditionalConfig string `json:"additionalConfig,omitempty"`
}

type DataSourceType string

//goland:noinspection GoUnusedConst
const (
	// DataSourceConfigMap mounts a config map as a data source.
	DataSourceConfigMap DataSourceType = "ConfigMap"
	// DataSourceSecret mounts a secret as a data source.
	DataSourceSecret DataSourceType = "Secret"
)

// AdditionalMount is a description of what data should be mounted to a specific Dogu volume (already defined in dogu.json).
type AdditionalMount struct {
	// SourceType defines where the data is coming from.
	// Valid options are:
	//   ConfigMap - data stored in a kubernetes ConfigMap.
	//   Secret - data stored in a kubernetes Secret.
	SourceType DataSourceType `json:"sourceType"`
	// Name is the name of the data source.
	Name string `json:"name"`
	// Volume is the name of the volume to which the data should be mounted. It is defined in the respective dogu.json.
	Volume string `json:"volume"`
	// Subfolder defines a subfolder in which the data should be put within the volume.
	// +optional
	Subfolder string `json:"subfolder,omitempty"`
}

type PlatformConfig struct {
	ResourceConfig         ResourceConfig     `json:"resource,omitempty"`
	ReverseProxyConfig     ReverseProxyConfig `json:"reverseProxy,omitempty"`
	AdditionalMountsConfig []AdditionalMount  `json:"additionalMounts,omitempty" patchStrategy:"replace"`
}
