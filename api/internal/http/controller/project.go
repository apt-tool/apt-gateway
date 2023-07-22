package controller

import (
	"fmt"

	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/internal/http/response"

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

// GetProject by its id
func (c Controller) GetProject(ctx *fiber.Ctx) error {
	projectID, _ := ctx.ParamsInt("project_id", 0)

	project, err := c.Models.Projects.GetByID(uint(projectID))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.project.Get] record not found error=%w", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ProjectResponse{}.DTO(project))
}

// DeleteProject by its id
func (c Controller) DeleteProject(ctx *fiber.Ctx) error {
	namespaceID, _ := ctx.ParamsInt("namespace_id", 0)

	if err := c.Models.Projects.Delete(uint(namespaceID)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.project.Create] failed to delete project error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}
