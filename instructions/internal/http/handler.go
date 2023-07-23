package http

import (
	"fmt"
	"log"
	"os"

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
	name := ctx.Query("name")

	file, err := ctx.FormFile("document")
	if err != nil {
		log.Println(fmt.Errorf("[handler.Upload] failed to get file error=%w", err))

		return fiber.ErrBadRequest
	}

	return ctx.SaveFile(file, fmt.Sprintf("./data/attacks/%s.txt", name))
}

func (h Handler) List(ctx *fiber.Ctx) error {
	entries, err := os.ReadDir("./data/attacks/")
	if err != nil {
		log.Println(fmt.Errorf("[handler.List] failed to get files error=%w", err))

		return fiber.ErrInternalServerError
	}

	list := make([]string, 0)

	for _, e := range entries {
		list = append(list, e.Name())
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

func (h Handler) Execute(ctx *fiber.Ctx) error {
	return nil
}
