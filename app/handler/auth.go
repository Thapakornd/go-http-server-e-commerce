package controllers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thapakornd/fiber-go/app/models"
	"github.com/thapakornd/fiber-go/app/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	newUser := models.User{}
	req := &models.SignUpUser{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	if h, err := req.HashPassword(req.Password); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	} else {
		newUser.Password = h
	}

	newUser.FirstName = req.FirstName
	newUser.LastName = req.LastName
	newUser.Email = req.Email
	newUser.Username = req.Username
	newUser.BirthOfDate = req.BirthOfDate
	newUser.Phone = req.Phone
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	if err := h.userStore.Create(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": newUser,
		},
	})
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	req := &models.SignInUser{}
	var err error
	var u *models.User

	if err := godotenv.Load(); err != nil {
		log.Fatal("environment error .env")
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "server error",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if req.Email != "" {
		u, err = h.userStore.GetByEmail(req.Email)
	} else if req.Username != "" {
		u, err = h.userStore.GetByUsername(req.Username)
	} else {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "Must have at least Username or Password",
		})
	}

	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		fmt.Printf("Unauthorized %v", req.Username)
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": gorm.ErrInvalidValue,
		})
	}

	genToken := &models.GenerateToken{
		IDS:      u.IDS,
		Username: u.Username,
	}

	accessToken := utils.GenerateJWT(genToken, 24*3, os.Getenv("USER_ROLE"))
	refreshToken := utils.GenerateJWT(genToken, 24*7, os.Getenv("USER_ROLE"))

	c.Cookie(&fiber.Cookie{
		Name:     "access-t",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * 3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh-t",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Login",
	})
}

func (h *Handler) SignOut(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name:  "access-t",
		Value: "",
	})

	c.Cookie(&fiber.Cookie{
		Name:  "refresh-t",
		Value: "",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful logout",
	})
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	req := &models.RefreshToken{}
	var newToken string

	if len(c.Body()) == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "Empty JSON body is not allowed",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	claims, err := utils.VerifyJWT(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": err.Error(),
		})
	}

	genToken := &models.GenerateToken{
		IDS:      int64(claims["id"].(float64)),
		Username: claims["username"].(string),
	}
	newToken = utils.GenerateJWT(genToken, 24*3, claims["role"].(string))

	c.Cookie(&fiber.Cookie{
		Name:     "access-t",
		Value:    newToken,
		Expires:  time.Now().Add(24 * 3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Get new access token",
	})
}

func (h *Handler) AdminSignIn(c *fiber.Ctx) error {

	if err := godotenv.Load(); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	req := &models.SignInUser{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "invalid json",
		})
	}

	if req.Email != os.Getenv("ADMIN_EMAIL") || req.Password != os.Getenv("ADMIN_PASSWORD") {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": "email or password incorrect",
		})
	}

	adminIds := os.Getenv("ADMIN_IDS")
	ids, err := strconv.ParseInt(adminIds, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail-server",
			"message": err.Error(),
		})
	}

	genToken := &models.GenerateToken{
		IDS:      ids,
		Username: os.Getenv("ADMIN_EMAIL"),
	}

	accessToken := utils.GenerateJWT(genToken, 24*3, os.Getenv("ADMIN_ROLE"))
	refreshToken := utils.GenerateJWT(genToken, 24*7, os.Getenv("ADMIN_ROLE"))

	c.Cookie(&fiber.Cookie{
		Name:     "access-t",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * 3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh-t",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful login",
	})
}
