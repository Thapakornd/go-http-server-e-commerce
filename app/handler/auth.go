package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/controllers"
	"github.com/thapakornd/fiber-go/app/models"
	"github.com/thapakornd/fiber-go/app/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(c *fiber.Ctx) error {
	newUser := models.User{}
	var v *controllers.Validator
	req := &models.SignUpUser{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := v.Validate(req); err != nil {
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

	if err := db.Create(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": newUser,
		},
	})
}

func SignIn(c *fiber.Ctx, db *gorm.DB, v *Validator, u *models.User) error {
	req := &models.SignInUser{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if req.Email == "" && req.Username == "" {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "Must have email or username",
		})
	}

	if result := db.Where("email = ?", req.Email).Or("username = ?", req.Username).First(&u); result.RowsAffected == 0 {
		fmt.Println("No records found")
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": gorm.ErrRecordNotFound,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		fmt.Printf("Unauthorized %v", req.Username)
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"status":  "fail-auth",
			"message": gorm.ErrInvalidValue,
		})
	}

	accessToken := utils.GenerateJWT(u, 24*3)
	refreshToken := utils.GenerateJWT(u, 24*7)

	c.Cookie(&fiber.Cookie{
		Name:    "access-t",
		Value:   accessToken,
		Expires: time.Now().Add(24 * 3 * time.Hour),
	})

	c.Cookie(&fiber.Cookie{
		Name:    "refresh-t",
		Value:   refreshToken,
		Expires: time.Now().Add(24 * 7 * time.Hour),
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful Login",
	})
}

func SignOut(c *fiber.Ctx) error {

	c.ClearCookie("refresh-t", "access-t")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successful logout",
	})
}
