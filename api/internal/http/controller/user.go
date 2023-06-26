package controller

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/http/request"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {
	req := new(request.Login)

	if err := ctx.BodyParser(req); err != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to parse body error=%w", err))

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		return err
	}

	token, err := c.JWTAuthenticator.GenerateToken(req.Email)
	if err != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to create token error=%w", err))

		return fiber.ErrInternalServerError
	}

	if er := c.RedisConnector.Set(req.Email, token, 0); er != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to save token error=%w", er))

		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).SendString(token)
}
