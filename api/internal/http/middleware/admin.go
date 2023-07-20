package middleware

import (
	"fmt"
	"log"
	"strconv"

	"github.com/automated-pen-testing/api/pkg/enum"

	"github.com/gofiber/fiber/v2"
)

// Admin middleware checks the user admin role.
func (m Middleware) Admin(ctx *fiber.Ctx) error {
	tmp, err := m.RedisConnector.Get(ctx.Locals("name").(string))
	if err != nil {
		log.Println(fmt.Errorf("token expired error=%v", err))

		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	role, _ := strconv.Atoi(tmp)

	if role == int(enum.RoleAdmin) {
		return ctx.Next()
	}

	return ctx.SendStatus(fiber.StatusForbidden)
}
