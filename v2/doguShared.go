package v2

import "k8s.io/apimachinery/pkg/api/resource"

type VolumeSize struct {
	resource.Quantity
}

// String returns the string representation of this volume size quantity.
func (vs VolumeSize) String() string {
	return vs.Quantity.String()
}

type ReverseProxyConfig struct {
	MaxBodySize      *BodySize
	RewriteTarget    RewriteTarget
	AdditionalConfig AdditionalConfig
}

type BodySize = resource.Quantity
type RewriteTarget string
type AdditionalConfig string
