package cmd

import (
	"github.com/darki73/goflaresync/pkg/configuration"
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/darki73/goflaresync/pkg/watcher"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the application",
	Long:  "Starts the application and begins the synchronization process",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initializeConfiguration(); err != nil {
			log.Fatal(err.Error())
		}

		instance := watcher.New()
		if err := instance.Start(); err != nil {
			log.Fatal(err.Error())
		}

		for {
			select {
			case <-configuration.ChangeChannel:
				if err := instance.Restart(); err != nil {
					log.Fatal(err.Error())
				}
			}
		}
	},
}

// init initializes the start command.
func init() {
	rootCmd.AddCommand(startCmd)
}
