package http

import (
	"fmt"

	"github.com/automated-pen-testing/instructions/internal/crypto"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) AccessMiddleware(ctx *fiber.Ctx) error {
	path := ctx.Query("path", "")
	cypher := crypto.GetMD5Hash(fmt.Sprintf("%s%s", h.AccessKey, path))

	if cypher != ctx.Get("x-token") {
		return fiber.ErrForbidden
	}

	return ctx.Next()
}
