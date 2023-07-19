package cmd

import (
	"fmt"
	"log"

	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/http/controller"
	"github.com/automated-pen-testing/api/internal/http/middleware"
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/internal/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

// API command is used to start API server
type API struct {
	Cfg config.Config
}

func (a API) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "build apt api",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a API) main() {
	// create new jwt authenticator
	authenticator := jwt.New(a.Cfg.JWT)

	// create redis connection
	redisConnection, err := redis.New(a.Cfg.Redis)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to redis cluster: %w", err))
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
		return ctx.SendString(ctx.Locals("name").(string))
	})

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", 8080)); err != nil {
		log.Fatal(err)
	}
}
