package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/http"
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// API command is used to start API server
type API struct {
	Cfg config.Config
	Db  *gorm.DB
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
	// create redis connection
	redisConnection, err := redis.New(a.Cfg.Redis)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to redis cluster: %w", err))
	}

	// create new models interface
	modelsInstance := models.New(a.Db)

	// creating a new fiber app
	app := fiber.New()

	app.Use(cors.New())

	// register http
	http.Register{
		Config:          a.Cfg,
		RedisConnector:  redisConnection,
		ModelsInterface: modelsInstance,
	}.Create(app)

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", a.Cfg.HTTP.Port)); err != nil {
		log.Fatal(err)
	}
}
