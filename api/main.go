package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// creating application flags
	var (
		PortFlag = flag.Int("port", 8080, "http port of api")
	)

	// parse flags
	flag.Parse()

	// creating a new fiber app
	app := fiber.New()

	// starting app on choosing port
	if err := app.Listen(fmt.Sprintf(":%d", *PortFlag)); err != nil {
		log.Fatal(err)
	}
}
