package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	controllers "github.com/thapakornd/fiber-go/app/handler"
	"github.com/thapakornd/fiber-go/app/routes"
	"github.com/thapakornd/fiber-go/app/store"
	"github.com/thapakornd/fiber-go/platform/config"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	d := config.New()
	config.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewAddressStore(d)
	cs := store.NewCartStore(d)
	os := store.NewOrderStore(d)
	ps := store.NewPaymentStore(d)
	prs := store.NewProductStore(d)

	h := controllers.NewHandler(
		us, as, cs, os, ps, prs,
	)

	routes.Register(app, h)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Listen(":5000")
}
