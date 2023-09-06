package controller

import (
	"fmt"

	"github.com/apt-tool/apt-gateway/internal/http/request"
	"github.com/apt-tool/apt-gateway/internal/http/response"

	"github.com/apt-tool/apt-core/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// CreateNamespace into system
func (c Controller) CreateNamespace(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(*user.User)

	req := new(request.NamespaceRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(
			ctx,
			fmt.Errorf("[controller.namespace.Create] failed to parse body error: %w", err),
			MessageRequestBody,
		)
	}

	if err := c.Models.Namespaces.Create(req.ToModel(u.Username)); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.Create] failed to create model error: %w", err),
			MessageFailedEntityCreate,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteNamespace by its id
func (c Controller) DeleteNamespace(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	if err := c.Models.Namespaces.Delete(uint(id)); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.Delete] failed to delete model error: %w", err),
			MessageFailedEntityRemove,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetNamespacesList of the system
func (c Controller) GetNamespacesList(ctx *fiber.Ctx) error {
	list, err := c.Models.Namespaces.Get()
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.GetList] failed to get models error: %w", err),
			MessageFailedEntityList,
		)
	}

	records := make([]*response.NamespaceResponse, 0)

	for _, item := range list {
		records = append(records, response.NamespaceResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}

// UpdateNamespace manages namespace users
func (c Controller) UpdateNamespace(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	req := new(request.NamespaceUpdateRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(
			ctx,
			fmt.Errorf("[controller.namespace.Update] failed to parse body error= %w", err),
			MessageRequestBody,
		)
	}

	if err := c.Models.UserNamespace.Clear(uint(id)); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.Update] failed to remove records error=%w", err),
			MessageFailedEntityUpdate,
		)
	}

	if er := c.Models.UserNamespace.Create(uint(id), req.UserIDs); er != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.Update] failed to update error= %w", er),
			MessageFailedEntityUpdate,
		)
	}

	// todo: update namespace

	return ctx.SendStatus(fiber.StatusOK)
}

// GetUserNamespacesList returns list of user namespaces
func (c Controller) GetUserNamespacesList(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(*user.User)

	ids, err := c.Models.UserNamespace.GetNamespaces(u.ID)
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.User] failed to get ids error= %w", err),
			MessageFailedEntityList,
		)
	}

	namespaces, err := c.Models.Namespaces.GetByIDs(ids)
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.namespace.User] failed to get namespaces error= %w", err),
			MessageFailedEntityList,
		)
	}

	list := make([]*response.NamespaceResponse, 0)

	for _, item := range namespaces {
		list = append(list, response.NamespaceResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(list)
}

// GetUserNamespace by namespace id
func (c Controller) GetUserNamespace(ctx *fiber.Ctx) error {
	namespace, err := c.Models.Namespaces.GetByID(ctx.Locals("namespace").(uint))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.namespace.GetUserNamespace] failed to get projects error=%w", err),
			MessageFailedEntityList,
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NamespaceResponse{}.DTO(namespace))
}

// GetNamespace returns namespace with its projects
func (c Controller) GetNamespace(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	namespace, err := c.Models.Namespaces.GetByID(uint(id))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.namespace.Get] failed to get projects error=%w", err),
			MessageFailedEntityList,
		)
	}

	ids, err := c.Models.UserNamespace.GetUsers(uint(id))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.namespace.Get] failed to get ids error=%w", err),
			MessageFailedEntityList,
		)
	}

	users, err := c.Models.Users.GetByIDs(ids)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.namespace.Get] failed to get users error=%w", err),
			MessageFailedEntityList,
		)
	}

	namespace.Users = users

	return ctx.Status(fiber.StatusOK).JSON(response.NamespaceResponse{}.DTO(namespace))
}
