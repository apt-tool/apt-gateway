package cmd

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/core"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// Core is the processing logic of the apt
type Core struct {
	Cfg config.Config
	Db  *gorm.DB
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

	// create new models interface
	m := models.New(c.Db)

	// register core
	core.Register{
		Cfg:    c.Cfg,
		Models: m,
	}.Create(app)

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", c.Cfg.Core.Port)); err != nil {
		log.Fatal(err)
	}
}
