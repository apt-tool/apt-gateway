package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// UserRegister will create a new user into system.
func (c Controller) UserRegister(ctx *fiber.Ctx) error {
	req := new(request.UserRegister)

	if err := ctx.BodyParser(req); err != nil {
		log.Println(fmt.Errorf("[controller.user.Register] failed to parse body error=%w", err))

		return ctx.Status(fiber.StatusBadRequest).SendString(errBadRequest.Error())
	}

	if err := req.Validate(); err != nil {
		log.Println(fmt.Errorf("[controller.user.Register] failed to validate request error=%w", err))

		return ctx.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%v:%v", errValidRequest.Error(), err))
	}

	userTmp := &user.User{
		Username: req.Name,
		Password: req.Pass,
	}

	if err := c.Models.Users.Create(userTmp); err != nil {
		log.Println(fmt.Errorf("[controller.user.Register] failed to create user error=%w", err))

		return ctx.Status(fiber.StatusInternalServerError).SendString(errDatabase.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// UserLogin logs in a user.
func (c Controller) UserLogin(ctx *fiber.Ctx) error {
	req := new(request.UserRegister)

	if err := ctx.BodyParser(req); err != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to parse body error=%w", err))

		return ctx.Status(fiber.StatusBadRequest).SendString(errBadRequest.Error())
	}

	if err := req.Validate(); err != nil {
		log.Println(fmt.Errorf("[controller.user.Login] failed to validate request error=%w", err))

		return ctx.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("%v:%v", errValidRequest.Error(), err))
	}

	userTmp, err := c.Models.Users.Validate(req.Name, req.Pass)
	if err != nil {
		log.Println(fmt.Errorf("[controller.user.Login] username and password don't match error=%w", err))

		return ctx.Status(fiber.StatusNotFound).SendString(errUserNotFound.Error())
	}

	token, etime, err := c.JWTAuthenticator.GenerateToken(userTmp.Username, userTmp.Role)
	if err != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to create token error=%w", err))

		return ctx.Status(fiber.StatusInternalServerError).SendString(errToken.Error())
	}

	if er := c.RedisConnector.Set(userTmp.Username, strconv.Itoa(int(userTmp.Role)), etime.Sub(time.Now())); er != nil {
		log.Println(fmt.Errorf("[controller.Loing] failed to save token error=%w", er))

		return ctx.Status(fiber.StatusInternalServerError).SendString(errToken.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
