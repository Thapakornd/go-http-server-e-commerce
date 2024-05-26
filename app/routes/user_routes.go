package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/thapakornd/fiber-go/app/handler"
	"github.com/thapakornd/fiber-go/app/middlewares"
)

func RegisterUser(r *fiber.App, h *controllers.Handler, m *middlewares.MiddlewareStruct) {
	v1 := r.Group("/api")

	guestUsers := v1.Group("/users")
	guestUsers.Post("/signUp", h.SignUp)
	guestUsers.Post("/signIn", h.SignIn)
	guestUsers.Post("/refresh", h.Refresh)
	guestUsers.Get("/signOut", h.SignOut)

	products := v1.Group("/products")
	products.Get("", h.GetAllProducts)
	products.Get("/:id", h.GetProduct)

	user := v1.Group("/user", m.AuthorizePermission)
	user.Get("/:id", h.CurrentUser)
	user.Put("/update/:id", h.UpdateUser)

	orders := user.Group("/orders")
	orders.Get("/:id", h.CurrentOrder)
	orders.Put("/:id", h.ChangeOrderInfo)
	orders.Get("/cancel/:id", h.CancelOrder)

	cart := user.Group("/cart")
	cart.Get("/:id", h.CurrentCart)
	cart.Post("/add/:id", h.AddItem)
	cart.Post("/del/:id", h.RemoveItem)
	cart.Put("/:id", h.UpdateItem)

	addresses := user.Group("/addresses")
	addresses.Get("/:id", h.CurrentAddresses)
	addresses.Post("/:id", h.AddAddress)
	addresses.Delete("/:id", h.RemoveAddress)
	addresses.Put("/:id", h.UpdateAddress)

	payments := user.Group("/payments")
	payments.Get("/:id", h.CurrentPayments)
	payments.Post("/:id", h.AddPayment)
	payments.Put("/:id", h.UpdatePayment)
	payments.Delete("/:id", h.RemovePayment)

}
