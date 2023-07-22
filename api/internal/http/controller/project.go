package controller

import (
	"fmt"

	"github.com/automated-pen-testing/api/internal/http/request"

	"github.com/gofiber/fiber/v2"
)

// CreateProject into system
func (c Controller) CreateProject(ctx *fiber.Ctx) error {
	req := new(request.ProjectRequest)
	namespaceID, _ := ctx.ParamsInt("namespace_id", 0)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.project.Create] failed to parse body error=%w", err))
	}

	if err := c.Models.Projects.Create(req.ToModel(uint(namespaceID))); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.project.Create] failed to create project error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteProject by its id
func (c Controller) DeleteProject(ctx *fiber.Ctx) error {
	namespaceID, _ := ctx.ParamsInt("namespace_id", 0)

	if err := c.Models.Projects.Delete(uint(namespaceID)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.project.Create] failed to delete project error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}
