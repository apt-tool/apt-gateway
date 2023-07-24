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
	id, _ := ctx.ParamsInt("namespace_id", 0)

	if err := c.Models.Namespaces.Delete(uint(id)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Delete] failed to delete model error: %w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetNamespaces of the system
func (c Controller) GetNamespaces(ctx *fiber.Ctx) error {
	list, err := c.Models.Namespaces.Get()
	if err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Get] failed to get models error: %w", err))
	}

	records := make([]*response.NamespaceResponse, 0)

	for _, item := range list {
		records = append(records, response.NamespaceResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}

// UpdateNamespace manages namespace users
func (c Controller) UpdateNamespace(ctx *fiber.Ctx) error {
	req := new(request.NamespaceUpdateRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.namespace.Update] failed to parse body error= %w", err))
	}

	if err := c.Models.Namespaces.RemoveUsers(req.NamespaceID); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Update] failed to remove records error=%w", err))
	}

	list, err := c.Models.Users.GetByIDs(req.UserIDs)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.namespace.Update] failed to find users error= %w", err))
	}

	if er := c.Models.Namespaces.AddUser(req.NamespaceID, list); er != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.Update] failed to update error= %w", er))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetUserNamespaces by users name
func (c Controller) GetUserNamespaces(ctx *fiber.Ctx) error {
	name := ctx.Locals("name").(string)

	u, err := c.Models.Users.GetByName(name, true)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.namespace.User] failed to get records error= %w", err))
	}

	namespaces, err := c.Models.Namespaces.GetUserNamespaces(u.ID)
	if err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.namespace.User] failed to get namespaces error= %w", err))
	}

	list := make([]*response.NamespaceResponse, 0)

	for _, item := range namespaces {
		list = append(list, response.NamespaceResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

// GetNamespace returns namespace with its projects
func (c Controller) GetNamespace(ctx *fiber.Ctx) error {
	namespaceID, _ := ctx.ParamsInt("namespace_id", 0)

	namespace, err := c.Models.Namespaces.GetByID(uint(namespaceID), false)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.namespace.Get] failed to get projects error=%w", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NamespaceResponse{}.DTO(namespace))
}

// GetNamespaceUsers returns namespace with its users
func (c Controller) GetNamespaceUsers(ctx *fiber.Ctx) error {
	namespaceID, _ := ctx.ParamsInt("namespace_id", 0)

	namespace, err := c.Models.Namespaces.GetByID(uint(namespaceID), true)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.namespace.GetUsers] failed to get users error=%w", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NamespaceResponse{}.DTO(namespace))
}
