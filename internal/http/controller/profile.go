package controller

import (
	"fmt"
	"github.com/apt-tool/apt-gateway/internal/http/response"

	"github.com/apt-tool/apt-gateway/internal/http/request"

	"github.com/gofiber/fiber/v2"
)

// GetProfile profile
func (c Controller) GetProfile(ctx *fiber.Ctx) error {
	record, err := c.Models.Users.GetByName(ctx.Locals("name").(string))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.user.Get] username and password don't match error=%w", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.UserResponse{}.DTO(record))
}

// UpdateProfile information
func (c Controller) UpdateProfile(ctx *fiber.Ctx) error {
	req := new(request.UserRegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.user.Update] failed to parse body error=%w", err))
	}

	if er := c.Models.Users.UpdateInfo(ctx.Locals("name").(string), req.Name); er != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.user.Update] failed to update user error=%w", er))
	}

	return ctx.SendStatus(fiber.StatusOK)
}
