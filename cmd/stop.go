package cmd

import (
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/darki73/goflaresync/pkg/service"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command.
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the service",
	Long:  "Stops the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initializeConfiguration(); err != nil {
			log.Fatal(err.Error())
		}

		manager := service.NewManager()
		systemManager := manager.GetSystemManager()
		if systemManager != nil {
			if err := systemManager.GetProcessIdentifierHandler().HandleApplicationStop(); err != nil {
				log.Fatal(err.Error())
			}
		}
	},
}

// init initializes the command.
func init() {
	rootCmd.AddCommand(stopCmd)
}
