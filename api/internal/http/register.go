package http

import (
	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/http/controller"
	"github.com/automated-pen-testing/api/internal/http/middleware"
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/internal/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

type Register struct {
	Cfg config.Config
	Rdb redis.Connector
}

func (r Register) Create(app *fiber.App) {
	// create new jwt authenticator
	authenticator := jwt.New(r.Cfg.JWT)

	// create middleware and controller
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
	}
	ctl := controller.Controller{
		JWTAuthenticator: authenticator,
		RedisConnector:   r.Rdb,
	}

	app.Post("/login", ctl.Login)

	auth := app.Use(mid.Auth)

	users := auth.Group("/users")

	users.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(ctx.Locals("name").(string))
	})
}
