package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// creating application flags
	var (
		PortFlag = flag.Int("port", 8080, "http port of runner service")
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
