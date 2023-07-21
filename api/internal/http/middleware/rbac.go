package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (m Middleware) UserNamespace(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("namespace_id", 0)
	id := uint(tmp)

	u, err := m.Models.Users.GetByName(ctx.Locals("name").(string), true)
	if err != nil {
		return m.ErrHandler.ErrRecordNotFound(ctx, err)
	}

	for _, item := range u.Namespaces {
		if item.ID == id {
			return ctx.Next()
		}
	}

	return m.ErrHandler.ErrAccess(ctx, errors.New("user is not in this namespace"))
}
