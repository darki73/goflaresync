package configuration

import (
	"github.com/darki73/goflaresync/pkg/configuration/cloudflare"
	"github.com/darki73/goflaresync/pkg/configuration/records"
	"github.com/darki73/goflaresync/pkg/configuration/watcher"
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	configuration *Configuration
	mutex         = &sync.RWMutex{}
	ChangeChannel = make(chan bool)
)

// Configuration is the definition of the configuration.
type Configuration struct {
	// Credentials is the Cloudflare Credentials configuration.
	Credentials *cloudflare.Configuration `json:"credentials" yaml:"credentials" xml:"credentials" toml:"credentials" mapstructure:"credentials"`
	// Records is the Cloudflare Records configuration.
	Records []*records.Configuration `json:"records" yaml:"records" xml:"records" toml:"records" mapstructure:"records"`
	// Watcher is the Watcher configuration.
	Watcher *watcher.Configuration `json:"watcher" yaml:"watcher" xml:"watcher" toml:"watcher" mapstructure:"watcher"`
	// LogLevel is the log level.
	LogLevel string `json:"log_level" yaml:"log_level" xml:"log_level" toml:"log_level" mapstructure:"log_level" env:"GOFLARESYNC_LOG_LEVEL"`
}

// GetCredentials returns the Cloudflare Credentials configuration.
func (configuration *Configuration) GetCredentials() *cloudflare.Configuration {
	return configuration.Credentials
}

// GetRecords returns the Cloudflare Records configuration.
func (configuration *Configuration) GetRecords() []*records.Configuration {
	return configuration.Records
}

// GetWatcher returns the Watcher configuration.
func (configuration *Configuration) GetWatcher() *watcher.Configuration {
	return configuration.Watcher
}

// GetLogLevel returns the log level.
func (configuration *Configuration) GetLogLevel() log.Level {
	logLevel, _ := log.ParseLevel(configuration.LogLevel)
	return logLevel
}

// LoadConfiguration loads the configuration from the given options.
func LoadConfiguration(options *ConfigurationOptions) error {
	viper.SetConfigName(options.GetName())
	viper.SetConfigType(options.GetExtension())
	viper.AddConfigPath(options.GetPath())
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	configuration = &Configuration{
		Credentials: cloudflare.InitializeWithDefaults(),
		Records:     []*records.Configuration{},
		Watcher:     watcher.InitializeWithDefaults(),
		LogLevel:    "i",
	}

	if err := viper.Unmarshal(configuration); err != nil {
		return err
	}

	if err := setLogLevel(); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		log.InfofWithFields(
			"configuration file changed: %s",
			log.FieldsMap{
				"source": "configuration",
			},
			event.Name,
		)

		mutex.Lock()
		defer mutex.Unlock()

		if err := viper.Unmarshal(configuration); err != nil {
			log.ErrorfWithFields(
				"error re-loading configuration: %s",
				log.FieldsMap{
					"source": "configuration",
				},
				err,
			)
		}

		if err := setLogLevel(); err != nil {
			log.ErrorfWithFields(
				"error setting log level: %s",
				log.FieldsMap{
					"source": "configuration",
				},
				err,
			)
		}

		ChangeChannel <- true
	})

	return nil
}

// GetConfiguration returns the configuration for the application.
func GetConfiguration() *Configuration {
	mutex.RLock()
	defer mutex.RUnlock()

	return configuration
}

// setLogLevel sets the log level.
func setLogLevel() error {
	desiredLogLevel, err := log.ParseLevel(configuration.LogLevel)

	if err != nil {
		return err
	}

	log.SetLevel(desiredLogLevel)

	return nil
}
