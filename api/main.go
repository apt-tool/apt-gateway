package main

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/cmd"
	"github.com/automated-pen-testing/api/internal/config"

	"github.com/spf13/cobra"
)

func main() {
	// load configs
	cfg := config.Load("config.yml")

	// create root command
	root := cobra.Command{}

	// add sub commands to root
	root.AddCommand(
		cmd.API{
			Cfg: cfg,
		}.Command(),
		cmd.Core{}.Command(),
	)

	// execute root command
	if err := root.Execute(); err != nil {
		log.Fatal(fmt.Errorf("failed to execute command: %w", err))
	}
}
