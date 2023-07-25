package http

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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

	file, err := ctx.FormFile("script")
	if err != nil {
		log.Println(fmt.Errorf("[handler.Upload] failed to get file error=%w", err))

		return fiber.ErrBadRequest
	}

	return ctx.SaveFile(file, fmt.Sprintf("./data/attacks/%s.sh", name))
}

func (h Handler) List(ctx *fiber.Ctx) error {
	entries, err := os.ReadDir("./data/attacks/")
	if err != nil {
		log.Println(fmt.Errorf("[handler.List] failed to get files error=%w", err))

		return fiber.ErrInternalServerError
	}

	list := make([]string, 0)

	for _, e := range entries {
		if e.Name() != "NOTE.txt" {
			list = append(list, strings.Replace(e.Name(), ".sh", "", 1))
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

func (h Handler) Execute(ctx *fiber.Ctx) error {
	req := new(ExecuteRequest)

	if err := ctx.BodyParser(&req); err != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to parse body error=%w", err))

		return fiber.ErrBadRequest
	}

	path := fmt.Sprintf("./data/attacks/%s.sh", req.Path)

	cmd, err := exec.Command("/bin/sh", path, req.Param).Output()
	if err != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to get files error=%w", err))

		return fiber.ErrInternalServerError
	}

	f, err := os.Create(fmt.Sprintf("./data/docs/%d.txt", req.DocumentID))
	if err != nil {
		log.Println(fmt.Errorf("[handler.Execute] failed to store log file error=%w", err))

		return fiber.ErrInternalServerError
	}

	defer f.Close()

	_, _ = f.Write(cmd)

	return ctx.SendStatus(fiber.StatusOK)
}
