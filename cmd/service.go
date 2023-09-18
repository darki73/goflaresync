package cmd

import (
	"github.com/darki73/goflaresync/pkg/helpers"
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// serviceCmd represents the service command.
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Handles the service commands",
}

// serviceInstallCmd represents the service install command.
var serviceInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the service",
	Long:  "Installs the systemd service and enables it",
	Run: func(cmd *cobra.Command, args []string) {
		checkForSystemd()
		checkIfRoot()

		servicePath := "/etc/systemd/system/goflaresync.service"

		if _, err := os.Stat(servicePath); err == nil {
			log.WarnWithFields(
				"service file already exists",
				log.FieldsMap{
					"source": "main",
				},
			)
			return
		}

		err := os.WriteFile(servicePath, []byte(serviceContent()), 0644)
		if err != nil {
			log.Fatalf(
				"failed to write the service file: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
			return
		}

		if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
			log.Fatalf(
				"failed to reload systemd: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully reloaded systemd",
			log.FieldsMap{
				"source": "main",
			},
		)

		if err := exec.Command("systemctl", "enable", "goflaresync").Run(); err != nil {
			log.Fatalf(
				"failed to enable the service: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully enabled the service",
			log.FieldsMap{
				"source": "main",
			},
		)
	},
}

// serviceUninstallCmd represents the service uninstall command.
var serviceUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the service",
	Long:  "Disables the the systemd service and uninstalls it",
	Run: func(cmd *cobra.Command, args []string) {
		checkForSystemd()
		checkIfRoot()

		servicePath := "/etc/systemd/system/goflaresync.service"
		if _, err := os.Stat(servicePath); err == nil {

			if err := exec.Command("systemctl", "stop", "goflaresync").Run(); err != nil {
				log.Fatalf(
					"failed to stop the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}

			log.InfoWithFields(
				"successfully stopped the service",
				log.FieldsMap{
					"source": "main",
				},
			)

			if err := exec.Command("systemctl", "disable", "goflaresync").Run(); err != nil {
				log.Fatalf(
					"failed to disable the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}

			log.InfoWithFields(
				"successfully disabled the service",
				log.FieldsMap{
					"source": "main",
				},
			)

			if err := os.Remove(servicePath); err != nil {
				log.FatalfWithFields(
					"failed to remove the service file: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}

			log.InfoWithFields(
				"successfully removed the service file",
				log.FieldsMap{
					"source": "main",
				},
			)

			if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
				log.Fatalf(
					"failed to reload systemd: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}

			log.InfoWithFields(
				"successfully reloaded systemd",
				log.FieldsMap{
					"source": "main",
				},
			)
		}
	},
}

// serviceStartCmd represents the service start command.
var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the service",
	Long:  "Starts the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("systemctl", "start", "goflaresync").Run(); err != nil {
			log.Fatalf(
				"failed to start the service: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully started the service",
			log.FieldsMap{
				"source": "main",
			},
		)
	},
}

// serviceStopCmd represents the service start command.
var serviceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the service",
	Long:  "Stops the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("systemctl", "stop", "goflaresync").Run(); err != nil {
			log.Fatalf(
				"failed to stop the service: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully stopped the service",
			log.FieldsMap{
				"source": "main",
			},
		)
	},
}

// serviceRestartCmd represents the service restart command.
var serviceRestartCmd = &cobra.Command{
	Use:  "restart",
	Long: "Restarts the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("systemctl", "restart", "goflaresync").Run(); err != nil {
			log.Fatalf(
				"failed to restart the service: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully restarted the service",
			log.FieldsMap{
				"source": "main",
			},
		)
	},
}

// serviceEnableCmd represents the service enable command.
var serviceEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables the service",
	Long:  "Enables the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("systemctl", "enable", "goflaresync").Run(); err != nil {
			log.Fatalf(
				"failed to enable the service: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully enabled the service",
			log.FieldsMap{
				"source": "main",
			},
		)
	},
}

// serviceDisableCmd represents the service disable command.
var serviceDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables the service",
	Long:  "Disables the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := exec.Command("systemctl", "disable", "goflaresync").Run(); err != nil {
			log.Fatalf(
				"failed to disable the service: %s",
				log.FieldsMap{
					"source": "main",
				},
				err.Error(),
			)
		}

		log.InfoWithFields(
			"successfully disabled the service",
			log.FieldsMap{
				"source": "main",
			},
		)
	},
}

// checkForSystemd checks if systemd is available.
func checkForSystemd() {
	if !helpers.HasSystemd() {
		log.FatalfWithFields(
			"this system does not have systemd",
			log.FieldsMap{
				"source": "main",
			},
		)
	}
}

// checkIfRoot checks if the current user is root.
func checkIfRoot() {
	if !helpers.IsRoot() {
		log.FatalfWithFields(
			"this command must be run as root",
			log.FieldsMap{
				"source": "main",
			},
		)
	}
}

// init registers the command and flags.
func init() {
	serviceCmd.AddCommand(serviceInstallCmd)
	serviceCmd.AddCommand(serviceUninstallCmd)
	serviceCmd.AddCommand(serviceStartCmd)
	serviceCmd.AddCommand(serviceStopCmd)
	serviceCmd.AddCommand(serviceRestartCmd)
	serviceCmd.AddCommand(serviceEnableCmd)
	serviceCmd.AddCommand(serviceDisableCmd)
	rootCmd.AddCommand(serviceCmd)
}

// serviceContent returns the content of the service file.
func serviceContent() string {
	return `[Unit]
Description=GoFlareSync Service
After=network.target

[Service]
Type=simple
ExecStart=/usr/bin/goflaresync start
Restart=on-failure

[Install]
WantedBy=multi-user.target
`
}
