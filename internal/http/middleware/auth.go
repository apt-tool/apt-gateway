package middleware

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Auth middleware check user authentication.
func (m Middleware) Auth(ctx *fiber.Ctx) error {
	if token := ctx.Get("x-token", ""); token != "" {
		if name, err := m.JWTAuthenticator.ParseToken(token); err == nil {
			user, er := m.Models.Users.GetByName(name)
			if er != nil {
				return m.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("user not found"))
			}

			ctx.Locals("user", user)

			return ctx.Next()
		} else {
			return m.ErrHandler.ErrUnauthorized(ctx, err)
		}
	} else {
		return m.ErrHandler.ErrUnauthorized(ctx, errors.New("token not found"))
	}
}
