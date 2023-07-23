package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AccessKey  string
	PrivateKey string
}

func (h Handler) Health(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	path := ctx.Query("path")

	return ctx.Download(fmt.Sprintf("./data/docs/%s.txt", path), fmt.Sprintf("%s.txt", path))
}

func (h Handler) Upload(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) List(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Execute(ctx *fiber.Ctx) error {
	return nil
}
