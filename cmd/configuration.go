package cmd

import (
	"fmt"
	"github.com/darki73/goflaresync/pkg/configuration"
	"github.com/darki73/goflaresync/pkg/log"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

// configurationCmd represents the configuration command.
var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Manage configuration",
}

// configurationDisplayCmd represents the display subcommand.
var configurationDisplayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display configuration",
	Long:  "Display configuration but omit sensitive information such as API key.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initializeConfiguration(); err != nil {
			log.Fatal(err.Error())
		}

		config := configuration.GetConfiguration()
		if config.GetCredentials().GetToken() != "" {
			config.Credentials.Token = "********"
		}

		if config.GetCredentials().GetEmail() != "" {
			config.Credentials.Email = "********"
		}

		displayConfiguration(config)
	},
}

// configurationDisplayFullCmd represents the display-full subcommand.
var configurationDisplayFullCmd = &cobra.Command{
	Use:   "display-full",
	Short: "Display configuration",
	Long:  "Display configuration including sensitive information such as API key.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initializeConfiguration(); err != nil {
			log.Fatal(err.Error())
		}

		config := configuration.GetConfiguration()
		displayConfiguration(config)
	},
}

// init registers the commands and subcommands.
func init() {
	configurationCmd.AddCommand(configurationDisplayCmd)
	configurationCmd.AddCommand(configurationDisplayFullCmd)
	rootCmd.AddCommand(configurationCmd)
}

// displayConfiguration displays the configuration.
func displayConfiguration(config *configuration.Configuration) {
	tmplStr := `Credentials:
  Email: {{ .Credentials.Email }}
  Token: {{ .Credentials.Token }}
Records:
{{- range .Records }}
  - Type: {{ .Type }}
    Name: {{ .Name }}
{{- end }}
Watcher:
  Interval: {{ .Watcher.Interval }}
Log Level: {{ .LogLevel }}
`
	tmpl, err := template.New("config").Parse(tmplStr)
	if err != nil {
		fmt.Println("Error parsing configuration template:", err)
		return
	}

	err = tmpl.Execute(os.Stdout, config)
	if err != nil {
		fmt.Println("Error executing configuration template:", err)
	}
}
