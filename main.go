package main

import (
	"github.com/darki73/goflaresync/cmd"
	"github.com/darki73/goflaresync/pkg/log"
)

// main is the entry point of the application.
func main() {
	if err := cmd.Execute(); err != nil {
		log.FatalfWithFields(
			"failed to execute command: %s",
			log.FieldsMap{
				"source": "main",
			},
			err.Error(),
		)
	}
}
