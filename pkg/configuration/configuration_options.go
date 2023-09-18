package configuration

// ConfigurationOptions is the options for the configuration.
type ConfigurationOptions struct {
	// Name is the name of the configuration file.
	Name string
	// Extension is the extension of the configuration file.
	Extension string
	// Path is the path to the configuration file.
	Path string
}

// NewConfiguration creates a new configuration options.
func NewConfiguration(name, extension, path string) *ConfigurationOptions {
	return &ConfigurationOptions{
		Name:      name,
		Extension: extension,
		Path:      path,
	}
}

// GetName returns the name of the configuration file.
func (options *ConfigurationOptions) GetName() string {
	return options.Name
}

// GetExtension returns the extension of the configuration file.
func (options *ConfigurationOptions) GetExtension() string {
	return options.Extension
}

// GetPath returns the path to the configuration file.
func (options *ConfigurationOptions) GetPath() string {
	return options.Path
}
