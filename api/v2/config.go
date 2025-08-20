package v2

import "fmt"

// Config contains Dogu and Cloudogu EcoSystem specific configuration data which determine set-up and run
// behavior respectively.
type Config struct {
	// Dogus contains Dogu specific configuration entries
	// +optional
	Dogus map[string][]ConfigEntry `json:"dogus,omitempty"`
	// Global contains EcoSystem specific configuration entries
	// +optional
	Global []ConfigEntry `json:"global,omitempty"`
}

// ConfigEntry represents a single configuration entry that can be either regular or sensitive
type ConfigEntry struct {
	// Key is the configuration key name
	// +required
	Key string `json:"key"`

	// Absent indicates whether this key should be deleted (true) or set (false)
	// +optional
	Absent *bool `json:"absent,omitempty"`

	// Value is used for regular (non-sensitive) configuration entries
	// Mutually exclusive with SecretRef
	// +optional
	Value *string `json:"value,omitempty"`

	// SecretRef is used for sensitive configuration entries
	// Mutually exclusive with Value
	// +optional
	SecretRef *SecretReference `json:"secretRef,omitempty"`
}

// SecretReference points to a value in a Kubernetes secret
type SecretReference struct {
	// Name is the name of the secret in the same namespace
	// +required
	Name string `json:"name"`
	// Key is the key within the secret
	// +required
	Key string `json:"key"`
}

// Validate ensures ConfigEntry has valid state
func (c *ConfigEntry) Validate() error {
	if *c.Absent {
		if c.Value != nil || c.SecretRef != nil {
			return fmt.Errorf("absent entries cannot have value or secretRef")
		}
		return nil
	}

	// For present entries, exactly one of Value or SecretRef must be set
	hasValue := c.Value != nil
	hasSecretRef := c.SecretRef != nil

	if hasValue != hasSecretRef {
		return fmt.Errorf("config entries can have either a value or a secretRef")
	}

	return nil
}

// IsSensitive returns true if this is a sensitive config entry
func (c *ConfigEntry) IsSensitive() bool {
	return c.SecretRef != nil
}
