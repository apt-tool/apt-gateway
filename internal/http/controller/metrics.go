package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// MetricsHandler returns cluster metrics
func (c Controller) MetricsHandler(ctx *fiber.Ctx) error {
	users, err := c.Models.Users.GetAll()
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[metrics] failed to get users error=%w", err),
			MessageFailedEntityList,
		)
	}

	namespaces, err := c.Models.Namespaces.GetAll()
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[metrics] failed to get namespaces error=%w", err),
			MessageFailedEntityList,
		)
	}

	tmp := 0
	for _, item := range namespaces {
		namespace, _ := c.Models.Namespaces.GetByID(item.ID)

		tmp = tmp + len(namespace.Projects)
	}

	c.Metrics.TotalProjects = tmp

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"users":      len(users),
		"namespaces": len(namespaces),
		"core":       c.Config.HTTP.Core,
		"ftp":        c.Config.FTP.Host,
		"jwt":        c.Config.JWT.ExpireTime,
		"mysql":      fmt.Sprintf("%s:%d", c.Config.MySQL.Host, c.Config.MySQL.Port),
		"metrics":    c.Metrics,
	})
}
