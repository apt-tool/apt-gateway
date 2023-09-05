package middleware

import (
	"errors"

	"github.com/apt-tool/apt-core/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// UserProject checks to see if user can access this project or not
func (m Middleware) UserProject(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("namespace_id", 0)
	id := uint(tmp)

	// get login user
	u := ctx.Locals("users").(*user.User)

	// get namespaces
	namespaces, err := m.Models.UserNamespace.GetNamespaces(u.ID)
	if err != nil {
		return m.ErrHandler.ErrRecordNotFound(ctx, err)
	}

	// check to see if namespace exists
	for _, item := range namespaces {
		if item == id {
			ctx.Locals("namespace", id)

			tmp, _ = ctx.ParamsInt("id", 0)

			ctx.Locals("project", uint(tmp))

			return ctx.Next()
		}
	}

	return m.ErrHandler.ErrAccess(ctx, errors.New("user is not in this namespace"))
}

// UserNamespace checks to see if user belongs to a namespace or not
func (m Middleware) UserNamespace(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("id", 0)
	id := uint(tmp)

	// get login user
	u := ctx.Locals("users").(*user.User)

	// get namespaces
	namespaces, err := m.Models.UserNamespace.GetNamespaces(u.ID)
	if err != nil {
		return m.ErrHandler.ErrRecordNotFound(ctx, err)
	}

	// check to see if namespace exists
	for _, item := range namespaces {
		if item == id {
			ctx.Locals("namespace", id)

			return ctx.Next()
		}
	}

	return m.ErrHandler.ErrAccess(ctx, errors.New("user is not in this namespace"))
}
