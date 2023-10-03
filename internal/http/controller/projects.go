package controller

import (
	"fmt"

	"github.com/ptaas-tool/gateway/internal/http/request"
	"github.com/ptaas-tool/gateway/internal/http/response"
	"github.com/ptaas-tool/gateway/internal/utils/crypto"

	"github.com/ptaas-tool/base-api/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// CreateProject into system
func (c Controller) CreateProject(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(*user.User)

	req := new(request.ProjectRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return c.ErrHandler.ErrBodyParser(
			ctx,
			fmt.Errorf("[controller.project.Create] failed to parse body error=%w", err),
			MessageRequestBody,
		)
	}

	if err := c.Models.Projects.Create(req.ToModel(ctx.Locals("namespace").(uint), u.Username)); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.project.Create] failed to create project error=%w", err),
			MessageFailedEntityCreate,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetProject by its id
func (c Controller) GetProject(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("id", 0)
	id := uint(tmp)

	project, err := c.Models.Projects.GetByID(id)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.project.Get] record not found error=%w", err),
			MessageFailedEntityList,
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.ProjectResponse{}.DTO(project))
}

// DeleteProject by its id
func (c Controller) DeleteProject(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("id", 0)
	id := uint(tmp)

	if err := c.Models.Projects.Delete(id); err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[controller.project.Delete] failed to delete project error=%w", err),
			MessageFailedEntityRemove,
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// ExecuteProject will send http request to core
func (c Controller) ExecuteProject(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("id", 0)
	projectID := uint(tmp)
	url := fmt.Sprintf("%s/%d", c.Config.HTTP.Core, projectID)

	c.Metrics.TotalExecutes++

	rsp, err := c.Client.Get(url, fmt.Sprintf("x-secure:%s", c.Config.HTTP.CoreSecret))
	if err != nil {
		c.Metrics.FailedRequests++

		return c.ErrHandler.ErrLogical(
			ctx,
			fmt.Errorf("[controller.project.Execute] failed to execute project error=%w", err),
			MessageFailedEntityList,
		)
	}

	if rsp.StatusCode != 200 {
		c.Metrics.FailedRequests++

		return c.ErrHandler.ErrLogical(
			ctx,
			fmt.Errorf("[controller.project.Execute] failed to execute project error=%w", err),
			MessageFailedExecute,
		)
	}

	c.Metrics.SuccessfulRequests++

	return ctx.SendStatus(fiber.StatusOK)
}

// DownloadProjectDocument will download the project document
func (c Controller) DownloadProjectDocument(ctx *fiber.Ctx) error {
	documentID, _ := ctx.ParamsInt("document_id", 0)

	cypher := crypto.GetMD5Hash(fmt.Sprintf("%s%d", c.Config.FTP.Access, documentID))

	url := fmt.Sprintf("%s/download?path=%d&token=%s", c.Config.FTP.Host, documentID, cypher)

	c.Metrics.TotalDownloads++

	return ctx.Redirect(url, fiber.StatusPermanentRedirect)
}
