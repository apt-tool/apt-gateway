package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
	DevMode bool
}

func (e ErrorHandler) ErrBodyParser(ctx *fiber.Ctx, err error, message ...string) error {
	if e.DevMode {
		log.Printf("[err.BodyParser] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("failed to parse body: %s", message[0]))
}

func (e ErrorHandler) ErrQueryParser(ctx *fiber.Ctx, err error, message ...string) error {
	if e.DevMode {
		log.Printf("[err.QueryParser] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("faield to parse query: %s", message[0]))
}

func (e ErrorHandler) ErrRecordNotFound(ctx *fiber.Ctx, err error, message ...string) error {
	if e.DevMode {
		log.Printf("[err.RecordNotFound] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("record not found: %s", message[0]))
}

func (e ErrorHandler) ErrValidation(ctx *fiber.Ctx, err error, message ...string) error {
	if e.DevMode {
		log.Printf("[err.Validation] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("failed to validate: %s", message[0]))
}

func (e ErrorHandler) ErrDatabase(ctx *fiber.Ctx, err error) error {
	if e.DevMode {
		log.Printf("[err.Database] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusInternalServerError).SendString("database error occurred")
}

func (e ErrorHandler) ErrLogical(ctx *fiber.Ctx, err error) error {
	if e.DevMode {
		log.Printf("[err.Logical] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusInternalServerError).SendString("server error")
}

func (e ErrorHandler) ErrUnauthorized(ctx *fiber.Ctx, err error) error {
	if e.DevMode {
		log.Printf("[err.Unauthorized] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusUnauthorized).SendString("user is unauthorized")
}

func (e ErrorHandler) ErrAccess(ctx *fiber.Ctx, err error) error {
	if e.DevMode {
		log.Printf("[err.Access] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusForbidden).SendString("user cannot access this resource")
}
