package controller

import (
	"fmt"

	"github.com/apt-tool/apt-gateway/internal/http/request"
	"github.com/apt-tool/apt-gateway/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// CreateUser will create a new user into system
func (c Controller) CreateUser(ctx *fiber.Ctx) error {
	req := new(request.UserRegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return c.ErrHandler.ErrBodyParser(
			ctx,
			fmt.Errorf("[controller.user.Register] failed to parse body error=%w", err),
			MessageRequestBody,
		)
	}

	if err := req.Validate(); err != nil {
		return c.ErrHandler.ErrValidation(
			ctx,
			fmt.Errorf("[controller.user.Register] failed to validate request error=%w", err),
			MessageLoginErrValidation,
		)
	}

	if err := c.Models.Users.Create(req.ToModel()); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.user.Register] failed to create user error=%w", err),
			MessageFailedEntityCreate,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetUser profile
func (c Controller) GetUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	record, err := c.Models.Users.GetByID(uint(id))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.user.Get] username and password don't match error=%w", err),
			MessageFailedEntityList,
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.UserResponse{}.DTO(record))
}

// GetUsersList returns the list of users
func (c Controller) GetUsersList(ctx *fiber.Ctx) error {
	list, err := c.Models.Users.GetAll()
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.user.List] failed to get users error=%w", err),
			MessageFailedEntityList,
		)
	}

	records := make([]*response.UserResponse, 0)

	for _, item := range list {
		records = append(records, response.UserResponse{}.DTO(item))
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}

// UpdateUser changes the users role
func (c Controller) UpdateUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id", 0)

	req := new(request.UserRegisterRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(
			ctx,
			fmt.Errorf("[controller.user.Update] failed to parse body error=%w", err),
			MessageRequestBody,
		)
	}

	u, err := c.Models.Users.GetByID(uint(id))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.user.Update] failed to find user error=%w", err),
			MessageFailedEntityUpdate,
		)
	}

	u.Username = req.Name
	u.Role = req.Role
	u.Password = req.Pass

	if er := c.Models.Users.Update(uint(id), u); er != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.user.Update] failed to update user error=%w", err),
			MessageFailedEntityUpdate,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DeleteUser removes user
func (c Controller) DeleteUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	if err := c.Models.Users.Delete(uint(id)); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.user.Delete] failed to delete user error=%w", err),
			MessageFailedEntityRemove,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
