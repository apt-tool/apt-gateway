package core

import (
	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/core/handler"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Register struct {
	Cfg config.Config
}

func (r Register) Create(app *fiber.App) {
	h := handler.Handler{
		Client: nil,
		Models: models.New(nil),
	}

	app.Get("/", h.Process)
}
