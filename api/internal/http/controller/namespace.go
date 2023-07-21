package controller

import (
	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/internal/http/response"
	"github.com/automated-pen-testing/api/pkg/models/namespace"

	"github.com/gofiber/fiber/v2"
)

// CreateNamespace into system
func (c Controller) CreateNamespace(ctx *fiber.Ctx) error {
	req := new(request.NamespaceRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	tmp := namespace.Namespace{
		Name: req.Name,
	}

	if err := c.Models.Namespaces.Create(&tmp); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteNamespace by its id
func (c Controller) DeleteNamespace(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	if err := c.Models.Namespaces.Delete(uint(id)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetNamespaces of the system
func (c Controller) GetNamespaces(ctx *fiber.Ctx) error {
	req := new(request.NamespaceQueryRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	list, err := c.Models.Namespaces.Get(req.Populate)
	if err != nil {
		return err
	}

	records := make([]*response.NamespaceResponse, 0)

	for _, item := range list {
		tmp := &response.NamespaceResponse{
			ID:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
		}

		records = append(records, tmp)
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}

// UserNamespace manages namespace users
func (c Controller) UserNamespace(ctx *fiber.Ctx) error {
	req := new(request.NamespaceUserRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	u, err := c.Models.Users.GetByID(req.UserID)
	if err != nil {
		return err
	}

	method := c.Models.Namespaces.RemoveUser
	if req.Add {
		method = c.Models.Namespaces.AddUser
	}

	if err := method(req.NamespaceID, u); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
