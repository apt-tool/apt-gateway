package middleware

import (
	"errors"

	"github.com/ptaas-tool/base-api/pkg/enum"
	"github.com/ptaas-tool/base-api/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// Admin middleware checks the user admin role.
func (m Middleware) Admin(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(*user.User)

	if u.Role == enum.RoleAdmin {
		return ctx.Next()
	}

	return m.ErrHandler.ErrAccess(ctx, errors.New("user cannot access this endpoint"), "sorry, you cannot access this resource!")
}
