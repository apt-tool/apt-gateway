package controller

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) UserRegister(ctx *fiber.Ctx) error {
	req := new(request.UserRegister)

	if err := ctx.BodyParser(req); err != nil {
		log.Println(fmt.Errorf("[controller.user.Register] failed to parse body error=%w", err))

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		return err
	}

	userTmp := &user.User{
		Username: req.Name,
		Password: req.Pass,
	}

	if err := c.Models.Users.Create(userTmp); err != nil {
		log.Println(fmt.Errorf("[controller.user.Register] failed to create user error=%w", err))

		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *Controller) UserLogin(ctx *fiber.Ctx) error {
	req := new(request.UserRegister)

	if err := ctx.BodyParser(req); err != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to parse body error=%w", err))

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		return err
	}

	userTmp, err := c.Models.Users.Validate(req.Name, req.Pass)
	if err != nil {
		return err
	}

	token, err := c.JWTAuthenticator.GenerateToken(userTmp.Username)
	if err != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to create token error=%w", err))

		return fiber.ErrInternalServerError
	}

	if er := c.RedisConnector.Set(userTmp.Username, token, 0); er != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to save token error=%w", er))

		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).SendString(token)
}
