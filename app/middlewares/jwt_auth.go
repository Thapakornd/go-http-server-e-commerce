package middlewares

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thapakornd/fiber-go/app/utils"
)

func AuthorizeUser(c *fiber.Ctx) error {
	token := strings.Split(c.Get("Authorization"), " ")[1]

	if token == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "Can't find any token",
		})
	}

	_, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "token invalid",
		})
	}

	return c.Next()
}

func AuthorizeAdmin(c *fiber.Ctx) error {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
		return c.Status(fiber.ErrConflict.Code).JSON(err)
	}

	token := strings.Split(c.Get("Authorization"), " ")[1]

	if token == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "Can't find any token",
		})
	}

	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "token invalid",
		})
	}

	if []byte(strconv.FormatInt(claims["id"].(int64), 10))[0] != '1' {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
