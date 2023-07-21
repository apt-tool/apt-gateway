package controller

import (
	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/pkg/models/namespace"

	"github.com/gofiber/fiber/v2"
)

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

func (c Controller) DeleteNamespace(ctx *fiber.Ctx) error {

}

func (c Controller) GetNamespaces(ctx *fiber.Ctx) error {

}

func (c Controller) UserNamespace(ctx *fiber.Ctx) error {

}
