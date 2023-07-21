package http

import (
	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/http/controller"
	"github.com/automated-pen-testing/api/internal/http/controller/handler"
	"github.com/automated-pen-testing/api/internal/http/middleware"
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/internal/utils/jwt"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Register struct {
	Cfg config.Config
	Rdb redis.Connector
	Mdb *models.Interface
}

func (r Register) Create(app *fiber.App) {
	// create new jwt authenticator
	authenticator := jwt.New(r.Cfg.JWT)

	eh := handler.ErrorHandler{DevMode: true}

	// create middleware and controller
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
		Models:           r.Mdb,
		RedisConnector:   r.Rdb,
		ErrHandler:       eh,
	}
	ctl := controller.Controller{
		JWTAuthenticator: authenticator,
		Models:           r.Mdb,
		RedisConnector:   r.Rdb,
		ErrHandler:       eh,
	}

	// register endpoints
	app.Post("/login", ctl.UserLogin)

	auth := app.Use(mid.Auth)
	admin := auth.Use(mid.Admin)

	users := admin.Group("/users")

	users.Post("/register", ctl.UserRegister)
}
