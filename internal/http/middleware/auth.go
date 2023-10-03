package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Auth middleware check user authentication.
func (m Middleware) Auth(ctx *fiber.Ctx) error {
	if token := ctx.Get("x-token", ""); token != "" {
		if _, err := m.JWTAuthenticator.ParseToken(token); err == nil {
			return ctx.Next()
		} else {
			return m.ErrHandler.ErrUnauthorized(ctx, err, "you are not logged in!")
		}
	} else {
		return m.ErrHandler.ErrUnauthorized(ctx, errors.New("token not found"), "login please!")
	}
}
