package cmd

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/core"

	"github.com/gofiber/fiber/v2"
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
	// create new fiber app
	app := fiber.New()

	// register core
	core.Register{
		Cfg: c.Cfg,
	}.Create(app)

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", c.Cfg.Core.Port)); err != nil {
		log.Fatal(err)
	}
}
