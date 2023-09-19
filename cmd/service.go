package cmd

import (
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/darki73/goflaresync/pkg/service"
	"github.com/spf13/cobra"
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
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.CreateService(); err != nil {
				log.FatalfWithFields(
					"failed to create the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}

			if err := systemManager.EnableService(); err != nil {
				log.FatalfWithFields(
					"failed to enable the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}

	},
}

// serviceUninstallCmd represents the service uninstall command.
var serviceUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the service",
	Long:  "Disables the the systemd service and uninstalls it",
	Run: func(cmd *cobra.Command, args []string) {
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.DisableService(); err != nil {
				log.FatalfWithFields(
					"failed to disable the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}

			if err := systemManager.DeleteService(); err != nil {
				log.FatalfWithFields(
					"failed to remove the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}
	},
}

// serviceStartCmd represents the service start command.
var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the service",
	Long:  "Starts the service",
	Run: func(cmd *cobra.Command, args []string) {
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.StartService(); err != nil {
				log.FatalfWithFields(
					"failed to start the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}
	},
}

// serviceStopCmd represents the service start command.
var serviceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the service",
	Long:  "Stops the service",
	Run: func(cmd *cobra.Command, args []string) {
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.StopService(); err != nil {
				log.FatalfWithFields(
					"failed to stop the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}
	},
}

// serviceRestartCmd represents the service restart command.
var serviceRestartCmd = &cobra.Command{
	Use:  "restart",
	Long: "Restarts the service",
	Run: func(cmd *cobra.Command, args []string) {
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.RestartService(); err != nil {
				log.FatalfWithFields(
					"failed to restart the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}
	},
}

// serviceEnableCmd represents the service enable command.
var serviceEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables the service",
	Long:  "Enables the service",
	Run: func(cmd *cobra.Command, args []string) {
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.EnableService(); err != nil {
				log.FatalfWithFields(
					"failed to enable the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}
	},
}

// serviceDisableCmd represents the service disable command.
var serviceDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables the service",
	Long:  "Disables the service",
	Run: func(cmd *cobra.Command, args []string) {
		manager := service.NewManager()
		systemManager := manager.GetSystemManager()

		if systemManager != nil {
			if err := systemManager.DisableService(); err != nil {
				log.FatalfWithFields(
					"failed to disable the service: %s",
					log.FieldsMap{
						"source": "main",
					},
					err.Error(),
				)
			}
		}
	},
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
