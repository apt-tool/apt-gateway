package controller

import (
	"fmt"

	"github.com/automated-pen-testing/api/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// GetInstructions returns all instructions
func (c Controller) GetInstructions(ctx *fiber.Ctx) error {
	list, err := c.Models.Instructions.Get()
	if err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.instruction.List] failed to get instructions error=%w", err))
	}

	records := make([]*response.InstructionResponse, 0)

	for _, item := range list {
		records = append(records, response.InstructionResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}
