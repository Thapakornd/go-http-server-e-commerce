package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/platform/config"
)

func main() {
	app := fiber.New()

	d := config.New()
	config.AutoMigrate(d)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Listen(":8080")
}
