package controller

import (
	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// CreateNamespace into system
func (c Controller) CreateNamespace(ctx *fiber.Ctx) error {
	req := new(request.NamespaceRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, err)
	}

	if err := c.Models.Namespaces.Create(req.ToModel()); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteNamespace by its id
func (c Controller) DeleteNamespace(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	if err := c.Models.Namespaces.Delete(uint(id)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetNamespaces of the system
func (c Controller) GetNamespaces(ctx *fiber.Ctx) error {
	req := new(request.NamespaceQueryRequest)

	if err := ctx.QueryParser(&req); err != nil {
		return c.ErrHandler.ErrQueryParser(ctx, err)
	}

	list, err := c.Models.Namespaces.Get(req.Populate)
	if err != nil {
		return c.ErrHandler.ErrDatabase(ctx, err)
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
		return c.ErrHandler.ErrBodyParser(ctx, err)
	}

	u, err := c.Models.Users.GetByID(req.UserID)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, err, errUserNotFound.Error())
	}

	method := c.Models.Namespaces.RemoveUser
	if req.Add {
		method = c.Models.Namespaces.AddUser
	}

	if er := method(req.NamespaceID, u); er != nil {
		return c.ErrHandler.ErrDatabase(ctx, er)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
