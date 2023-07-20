package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Auth middleware check user authentication.
func (m Middleware) Auth(ctx *fiber.Ctx) error {
	if token := ctx.Get("x-token", ""); token != "" {
		if name, err := m.JWTAuthenticator.ParseToken(token); err == nil {
			ctx.Locals("user", name)

			return ctx.Next()
		} else {
			log.Println(err)

			return ctx.SendStatus(http.StatusForbidden)
		}
	} else {
		return ctx.SendStatus(http.StatusUnauthorized)
	}
}
