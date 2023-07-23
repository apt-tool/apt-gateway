package http

import "github.com/gofiber/fiber/v2"

type Handler struct {
	AccessKey string
}

func (h Handler) Health(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	return nil
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
