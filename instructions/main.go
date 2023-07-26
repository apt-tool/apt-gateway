package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/automated-pen-testing/instructions/internal/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// get env variables
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	private := os.Getenv("PRIVATE_KEY")
	access := os.Getenv("ACCESS_KEY")

	// create new fiber app
	app := fiber.New()

	app.Use(cors.New())

	// create new handler
	h := http.Handler{
		AccessKey:  access,
		PrivateKey: private,
	}

	app.Get("/", h.List)
	app.Get("/health", h.Health)
	app.Get("/download", h.AccessMiddleware, h.Download)
	app.Post("/upload", h.Upload)
	app.Post("/execute", h.AuthMiddleware, h.Execute)

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(fmt.Errorf("failed to start ftp server error=%w", err))
	}
}
