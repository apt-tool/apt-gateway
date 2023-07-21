package handler

import (
	"strconv"

	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Client *client.Client
	Models *models.Interface
}

func (h Handler) Process(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("project_id", 0)

	return ctx.SendString(strconv.Itoa(id))
}

func (h Handler) Register(app *fiber.App) {
	app.Get("/:project_id", h.Process)
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
}
