package core

import (
	"github.com/automated-pen-testing/api/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Register struct {
	Cfg config.Config
}

func (r Register) Create(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})
}
