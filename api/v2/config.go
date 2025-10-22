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

// +kubebuilder:validation:XValidation:message="absent entries cannot have value or secretRef",rule="(has(self.absent) && self.absent) ? !has(self.value) && !has(self.secretRef) : true"
// +kubebuilder:validation:XValidation:message="config entries can have either a value or a secretRef",rule="(!has(self.absent) || !self.absent) ? has(self.value) != has(self.secretRef) : true"
// +kubebuilder:validation:XValidation:message="config entries with secret references have to be sensitive",rule="has(self.secretRef) ? has(self.sensitive) && self.sensitive : true"
// +kubebuilder:validation:XValidation:message="sensitive config entries are not allowed to have normal values",rule="(has(self.sensitive) && self.sensitive) ? !has(self.value) : true"

// ConfigEntry represents a single configuration entry that can be either regular or sensitive
type ConfigEntry struct {
	// Key is the configuration key name
	// +required
	Key string `json:"key"`

	// Absent indicates whether this key should be deleted (true) or set (false)
	// +optional
	Absent *bool `json:"absent,omitempty"`

	// Value is used for regular (non-sensitive) configuration entries
	// Mutually exclusive with SecretRef and ConfigRef
	// +optional
	Value *string `json:"value,omitempty"`

	// Sensitive indicates whether this config is sensitive and should be stored securely (true) or not (false)
	// +optional
	Sensitive *bool `json:"sensitive,omitempty"`

	// SecretRef is used for sensitive configuration entries
	// Mutually exclusive with Value
	// +optional
	SecretRef *Reference `json:"secretRef,omitempty"`

	// ConfigRef is used for configuration entries
	// Mutually exclusive with Value
	// +optional
	ConfigRef *Reference `json:"configRef,omitempty"`
}

// Reference points to a value in a Kubernetes secret
type Reference struct {
	// Name is the name of the secret or configmap in the same namespace
	// +required
	Name string `json:"name"`
	// Key is the key within the secret or configmap
	// +required
	Key string `json:"key"`
}

// Validate ensures ConfigEntry has valid state
func (c *ConfigEntry) Validate() error {
	if c.Absent != nil && *c.Absent {
		if c.Value != nil || c.SecretRef != nil || c.ConfigRef != nil {
			return fmt.Errorf("absent entries cannot have value, configRef or secretRef")
		}
		return nil
	}

	// For present entries, exactly one of Value or SecretRef must be set
	hasValue := c.Value != nil
	hasSecretRef := c.SecretRef != nil
	hasConfigRef := c.ConfigRef != nil

	if !c.hasOnlyOneValueConfiguration(hasValue, hasSecretRef, hasConfigRef) {
		return fmt.Errorf("config entries can have either a value, configRef or a secretRef")
	}

	return nil
}

func (c *ConfigEntry) hasOnlyOneValueConfiguration(bools ...bool) bool {
	hasValue := false
	for _, b := range bools {
		if b && !hasValue {
			hasValue = true
		} else if b {
			return false
		}
	}
	return hasValue
}
