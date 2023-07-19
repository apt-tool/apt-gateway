package handler

import (
	"github.com/automated-pen-testing/api/internal/core/request"
	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Client *client.Client
	Models *models.Interface
}

func (h Handler) Process(ctx *fiber.Ctx) error {
	req := new(request.Request)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
