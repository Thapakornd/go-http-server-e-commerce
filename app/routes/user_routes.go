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

	user := v1.Group("/user", m.AuthorizePermission)
	user.Get("/:id", h.CurrentUser)
	user.Put("/update", h.UpdateUser)
}
