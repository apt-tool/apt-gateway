package controller

import (
	"fmt"

	"github.com/apt-tool/apt-gateway/internal/http/request"
	"github.com/apt-tool/apt-gateway/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// Login logs in a user into our system
func (c Controller) Login(ctx *fiber.Ctx) error {
	req := new(request.UserRegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.user.Loing] failed to parse body error=%w", err))
	}

	if err := req.Validate(); err != nil {
		return c.ErrHandler.ErrValidation(ctx, fmt.Errorf("[controller.user.Login] failed to validate request error=%w", err))
	}

	userTmp, err := c.Models.Users.Validate(req.Name, req.Pass)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.user.Login] username and password don't match error=%w", err))
	}

	token, _, err := c.JWTAuthenticator.GenerateToken(userTmp.Username, userTmp.Role)
	if err != nil {
		return c.ErrHandler.ErrLogical(ctx, fmt.Errorf("[controller.user.Loing] failed to create token error=%w", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Token{
		Token: token,
	})
}
