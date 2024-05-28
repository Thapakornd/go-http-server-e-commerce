package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/thapakornd/fiber-go/app/handler"
	"github.com/thapakornd/fiber-go/app/middlewares"
)

func RegisterAdmin(r *fiber.App, h *controllers.Handler, m *middlewares.MiddlewareStruct) {

	// This will be routes some route that doesn't exist in user permissions.

	v1 := r.Group("/api")
	v1.Post("/signIn-admin", h.AdminSignIn)

	admin := v1.Group("/admin", m.AuthorizePermission)

	users := admin.Group("/users")
	users.Get("", h.GetAllUsers)
	users.Delete("/:id", h.RemoveUser)

	products := admin.Group("/products")
	products.Post("/add", h.AddNewProduct)
	products.Put("/:id", h.UpdateProduct)
	products.Delete("/:id", h.RemoveProduct)

	payments := admin.Group("/payments")
	payments.Get("/", h.GetAllPayments)

	orders := admin.Group("/orders")
	orders.Get("", h.GetAllOrders)
	orders.Delete("/:id", h.RemoveOrder)

	carts := admin.Group("/carts")
	carts.Get("", h.GetAllCarts)
	carts.Delete("/:id", h.RemoveCart)

	addresses := admin.Group("/addresses")
	addresses.Get("", h.GetAllAddresses)

	categories := admin.Group("/categories")
	categories.Get("", h.GetAllCategories)
	categories.Post("/add", h.AddCategory)
	categories.Put("/:id", h.ChangeCategoryInfo)
	categories.Delete("/:id", h.RemoveCategory)
}
