package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ptaas-tool/gateway/internal/http/response"
)

func (c Controller) GetTracksList(ctx *fiber.Ctx) error {
	tmp, _ := ctx.ParamsInt("project_id", 0)
	projectID := uint(tmp)

	id := uint(ctx.QueryInt("id", 0))

	tracks, err := c.Models.Tracks.Get(id, projectID)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(
			ctx,
			fmt.Errorf("[controller.Tracks.Get] record not found error=%w", err),
			MessageFailedEntityList,
		)
	}

	records := make([]*response.TrackResponse, 0)

	for _, track := range tracks {
		records = append(records, response.TrackResponse{}.DTO(track))
	}

	return ctx.Status(fiber.StatusOK).JSON(records)
}
