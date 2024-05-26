package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	controllers "github.com/thapakornd/fiber-go/app/handler"
	"github.com/thapakornd/fiber-go/app/middlewares"
	"github.com/thapakornd/fiber-go/app/routes"
	"github.com/thapakornd/fiber-go/app/store"
	"github.com/thapakornd/fiber-go/platform/config"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	app.Use(logger.New())

	d := config.New()
	config.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewAddressStore(d)
	cs := store.NewCartStore(d)
	oss := store.NewOrderStore(d)
	ps := store.NewPaymentStore(d)
	prs := store.NewProductStore(d)
	css := store.NewCategoryStore(d)

	h := controllers.NewHandler(
		us, as, cs, css, oss, ps, prs,
	)

	// Auth-jwt-middleware
	m_user := middlewares.NewJWTAuth(os.Getenv("USER_ROLE"))
	m_admin := middlewares.NewJWTAuth(os.Getenv("ADMIN_ROLE"))

	routes.RegisterUser(app, h, m_user)
	routes.RegisterAdmin(app, h, m_admin)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Listen(":5000")
}
