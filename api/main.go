package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/http/controller"
	"github.com/automated-pen-testing/api/internal/http/middleware"
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/internal/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// creating application flags
	var (
		PortFlag = flag.Int("port", 8080, "http port of api")
	)

	// parse flags
	flag.Parse()

	// load configs
	cfg := config.Load()

	// create new jwt authenticator
	authenticator := jwt.New(cfg.JWT)

	// create redis connection
	redisConnection, err := redis.New(cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	// create middleware and controller
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
	}
	ctl := controller.Controller{
		JWTAuthenticator: authenticator,
		RedisConnector:   redisConnection,
	}

	// creating a new fiber app
	app := fiber.New()

	app.Post("/login", ctl.Login)

	auth := app.Use(mid.Auth)

	users := auth.Group("/users")

	users.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(ctx.Locals("email").(string))
	})

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", *PortFlag)); err != nil {
		log.Fatal(err)
	}
}
