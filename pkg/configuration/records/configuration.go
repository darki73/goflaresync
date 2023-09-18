package records

// Configuration is the definition of a record configuration.
type Configuration struct {
	// Type is the type of the record.
	Type string `json:"type" yaml:"type" xml:"type" toml:"type" mapstructure:"type"`
	// Name is the name of the record.
	Name string `json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`
}

// GetType returns the type of the record.
func (configuration *Configuration) GetType() string {
	return configuration.Type
}

// GetName returns the name of the record.
func (configuration *Configuration) GetName() string {
	return configuration.Name
}
