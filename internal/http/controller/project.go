package controller

import (
	"fmt"

	"github.com/apt-tool/apt-gateway/internal/http/request"
	"github.com/apt-tool/apt-gateway/internal/http/response"
	"github.com/apt-tool/apt-gateway/internal/utils/crypto"

	"github.com/apt-tool/apt-core/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// CreateProject into system
func (c Controller) CreateProject(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(*user.User)

	req := new(request.ProjectRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.project.Create] failed to parse body error=%w", err))
	}

	if err := c.Models.Projects.Create(req.ToModel(ctx.Locals("namespace").(uint), u.Username)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.project.Create] failed to create project error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetProject by its id
func (c Controller) GetProject(ctx *fiber.Ctx) error {
	project, err := c.Models.Projects.GetByID(ctx.Locals("project").(uint))
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.project.Get] record not found error=%w", err))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ProjectResponse{}.DTO(project))
}

// DeleteProject by its id
func (c Controller) DeleteProject(ctx *fiber.Ctx) error {
	if err := c.Models.Projects.Delete(ctx.Locals("project").(uint)); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.project.Create] failed to delete project error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// ExecuteProject will send http request to core
func (c Controller) ExecuteProject(ctx *fiber.Ctx) error {
	projectID := ctx.Locals("project").(uint)
	url := fmt.Sprintf("%s/%d", c.Config.HTTP.Core, projectID)

	rsp, err := c.Client.Get(url, fmt.Sprintf("x-secure:%s", c.Config.HTTP.CoreSecret))
	if err != nil {
		return c.ErrHandler.ErrLogical(ctx, fmt.Errorf("[controller.project.Execute] failed to execute project error=%w", err))
	}

	if rsp.StatusCode != 200 {
		return c.ErrHandler.ErrLogical(ctx, fmt.Errorf("[controller.project.Execute] failed to execute project error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// DownloadProjectDocument will download the project document
func (c Controller) DownloadProjectDocument(ctx *fiber.Ctx) error {
	documentID, _ := ctx.ParamsInt("document_id", 0)

	cypher := crypto.GetMD5Hash(fmt.Sprintf("%s%d", c.Config.FTP.Access, documentID))

	url := fmt.Sprintf("%s/download?path=%d&token=%s", c.Config.FTP.Host, documentID, cypher)

	return ctx.Redirect(url, fiber.StatusPermanentRedirect)
}
