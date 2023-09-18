package cmd

import (
	"fmt"
	"github.com/darki73/goflaresync/pkg/version"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		tmpl, err := template.New("version").Parse(version.GetVersionTemplate())
		if err != nil {
			fmt.Println("Error parsing version template:", err)
			return
		}
		err = tmpl.Execute(os.Stdout, version.GetBuildInfo())
		if err != nil {
			fmt.Println("Error executing version template:", err)
		}
	},
}

// init initializes the version command.
func init() {
	rootCmd.AddCommand(versionCmd)
}
