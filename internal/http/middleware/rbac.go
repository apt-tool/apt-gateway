package middleware

import (
	"errors"

	"github.com/apt-tool/apt-core/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

func (m Middleware) UserNamespace(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("namespace_id", 0)
	id := uint(tmp)

	u := ctx.Locals("users").(*user.User)

	namespaces, err := m.Models.UserNamespace.GetNamespaces(u.ID)
	if err != nil {
		return m.ErrHandler.ErrRecordNotFound(ctx, err)
	}

	for _, item := range namespaces {
		if item == id {
			ctx.Locals("namespace_id", id)

			return ctx.Next()
		}
	}

	return m.ErrHandler.ErrAccess(ctx, errors.New("user is not in this namespace"))
}
