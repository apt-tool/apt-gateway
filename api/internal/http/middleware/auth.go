package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Auth(ctx *fiber.Ctx) error {
	if token := ctx.Get("x-token", ""); token != "" {
		if email, err := m.JWTAuthenticator.ParseToken(token); err == nil {
			ctx.Locals("email", email)

			return ctx.Next()
		} else {
			log.Println(err)

			return ctx.SendStatus(http.StatusForbidden)
		}
	} else {
		return ctx.SendStatus(http.StatusUnauthorized)
	}
}
