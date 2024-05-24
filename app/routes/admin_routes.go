package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/thapakornd/fiber-go/app/handler"
	"github.com/thapakornd/fiber-go/app/middlewares"
)

func RegisterAdmin(r *fiber.App, h *controllers.Handler, m *middlewares.MiddlewareStruct) {
	v1 := r.Group("/api")

	v1.Post("/signIn-admin", h.AdminSignIn)

	admin := v1.Group("/admin", m.AuthorizePermission)

	product := admin.Group("/products")
	product.Get("", h.GetAllProducts)

}
