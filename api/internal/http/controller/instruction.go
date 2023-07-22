package controller

import (
	"fmt"
	"github.com/automated-pen-testing/api/internal/http/request"

	"github.com/automated-pen-testing/api/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// CreateInstruction imports a new instruction
func (c Controller) CreateInstruction(ctx *fiber.Ctx) error {
	req := new(request.InstructionRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.instruction.Create] failed to parse body error=%w", err))
	}

	if err := c.Models.Instructions.Create(req.ToModel()); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.instruction.Create] failed to create instruction error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

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
