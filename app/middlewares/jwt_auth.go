package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/utils"
)

type MiddlewareStruct struct {
	role string
}

func NewJWTAuth(role string) *MiddlewareStruct {
	return &MiddlewareStruct{
		role: role,
	}
}

func (r *MiddlewareStruct) AuthorizePermission(c *fiber.Ctx) error {

	if err := c.Get("Authorization"); err == "" {
		return c.Status(fiber.ErrForbidden.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "permission not allowed",
		})
	}

	token := strings.Split(c.Get("Authorization"), " ")[1]

	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "token invalid",
		})
	}

	permission := claims["role"].(string)

	if permission != r.role {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "Unauthorized",
			"message": "fail-auth",
		})
	}

	return c.Next()
}
