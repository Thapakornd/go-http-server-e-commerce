package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/thapakornd/fiber-go/app/handler"
	"github.com/thapakornd/fiber-go/app/middlewares"
)

func Register(r *fiber.App, h *controllers.Handler) {
	v1 := r.Group("/api")

	guestUsers := v1.Group("/users")
	guestUsers.Post("/signUp", h.SignUp)
	guestUsers.Post("/signIn", h.SignIn)

	user := v1.Group("/user", middlewares.AuthorizeUser)
	user.Get("/signOut", h.SignOut)
	user.Post("/refresh", h.Refresh)
	user.Get("/:id", h.CurrentUser)
	user.Put("/update", h.UpdateUser)

	products := v1.Group("/products")
	// products.Get("")
	// products.Get("/:id")

	adminProduct := products.Group("/admin", middlewares.AuthorizeAdmin)
	adminProduct.Post("/add", h.AddNewProduct)
	adminProduct.Put("/update", h.UpdateProduct)
	adminProduct.Delete("/:id", h.RemoveProduct)

	// addresses := v1.Group("/addresses")
	// addresses.Get("")
	// addresses.Get("/:id")
}
