package watcher

import "time"

// Configuration is the definition of the configuration of the watcher.
type Configuration struct {
	// Interval is the interval at which the watcher will check for changes.
	Interval time.Duration `json:"interval" yaml:"interval" xml:"interval" toml:"interval" mapstructure:"interval" env:"GOFLARESYNC_WATCHER_INTERVAL"`
	// AddressSource is the source of real ip address.
	AddressSource string `json:"address_source" yaml:"address_source" xml:"address_source" toml:"address_source" mapstructure:"address_source" env:"GOFLARESYNC_ADDRESS_SOURCE"`
}

// InitializeWithDefaults initializes the configuration with default values.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		Interval:      5 * time.Minute,
		AddressSource: "https://api.ipify.org",
	}
}

// GetInterval returns the interval at which the watcher will check for changes.
func (configuration *Configuration) GetInterval() time.Duration {
	return configuration.Interval
}

// GetAddressSource returns the source of real ip address.
func (configuration *Configuration) GetAddressSource() string {
	return configuration.AddressSource
}
