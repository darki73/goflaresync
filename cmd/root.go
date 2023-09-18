package cmd

import (
	"github.com/darki73/goflaresync/pkg/configuration"
	"github.com/spf13/cobra"
)

var (
	// configurationPath is the path to the configuration file.
	configurationPath string
	// configurationName is the name of the configuration file.
	configurationName string
	// configurationExtension is the extension of the configuration file.
	configurationExtension string
)

// rootCmd is the root command.
var rootCmd = &cobra.Command{
	Use:   "goflaresync",
	Short: "Tool to update Cloudflare DNS records based on IP changes.",
	Long:  "This tool is designed to provide easy to use and flexible way to update Cloudflare DNS records based on IP changes.",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// init initializes the root command.
func init() {
	rootCmd.PersistentFlags().StringVar(&configurationPath, "configuration-path", "/etc/goflaresync", "Path to the configuration file")
	rootCmd.PersistentFlags().StringVar(&configurationName, "configuration-name", "config", "Name of the configuration file")
	rootCmd.PersistentFlags().StringVar(&configurationExtension, "configuration-extension", "yaml", "Extension of the configuration file")
}

// getConfigurationOptions returns the configuration options.
func getConfigurationOptions() *configuration.ConfigurationOptions {
	return &configuration.ConfigurationOptions{
		Path:      configurationPath,
		Name:      configurationName,
		Extension: configurationExtension,
	}
}

// initializeConfiguration initializes the configuration.
func initializeConfiguration() error {
	return configuration.LoadConfiguration(getConfigurationOptions())
}
