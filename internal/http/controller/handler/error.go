package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct {
	DevMode bool
}

func (e ErrorHandler) Error(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.Error] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrBodyParser(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.BodyParser] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrQueryParser(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.QueryParser] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrRecordNotFound(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.RecordNotFound] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrValidation(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.Validation] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrDatabase(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.Database] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrLogical(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.Logical] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrUnauthorized(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.Unauthorized] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}

func (e ErrorHandler) ErrAccess(ctx *fiber.Ctx, err error, message string) error {
	if e.DevMode {
		log.Printf("[err.Access] error: %v\n", err)
	}

	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error":   err,
		"message": message,
	})
}
