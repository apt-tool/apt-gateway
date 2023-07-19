package cmd

import (
	"github.com/automated-pen-testing/api/internal/config"

	"github.com/spf13/cobra"
)

// Core is the processing logic of the apt
type Core struct {
	Cfg config.Config
}

func (c Core) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "core",
		Short: "start apt core processor",
		Run: func(_ *cobra.Command, _ []string) {
			c.main()
		},
	}
}

func (c Core) main() {

}
