package controller

import (
	"fmt"

	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// CreateNamespace into system
func (c Controller) CreateNamespace(ctx *fiber.Ctx) error {
	req := new(request.NamespaceRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.namespace.Create] failed to parse body error: %w", err))
	}

	if err := c.Models.Namespaces.Create(req.ToModel()); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Create] failed to create model error: %w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteNamespace by its id
func (c Controller) DeleteNamespace(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	if err := c.Models.Namespaces.Delete(uint(id)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Delete] failed to delete model error: %w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetNamespaces of the system
func (c Controller) GetNamespaces(ctx *fiber.Ctx) error {
	req := new(request.NamespaceQueryRequest)

	if err := ctx.QueryParser(&req); err != nil {
		return c.ErrHandler.ErrQueryParser(ctx, fmt.Errorf("[controller.namespace.Get] failed to parse query error: %w", err))
	}

	list, err := c.Models.Namespaces.Get(req.Populate)
	if err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Get] failed to get models error: %w", err))
	}

	records := make([]*response.NamespaceResponse, 0)

	for _, item := range list {
		records = append(records, response.NamespaceResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}

// UserNamespace manages namespace users
func (c Controller) UserNamespace(ctx *fiber.Ctx) error {
	req := new(request.NamespaceUserRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.namespace.User] failed to parse body error: %w", err))
	}

	u, err := c.Models.Users.GetByID(req.UserID)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.namespace.User] failed to find model error: %w", err))
	}

	method := c.Models.Namespaces.RemoveUser
	if req.Add {
		method = c.Models.Namespaces.AddUser
	}

	if er := method(req.NamespaceID, u); er != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.User] failed to update error: %w", er))
	}

	return ctx.SendStatus(fiber.StatusOK)
}
