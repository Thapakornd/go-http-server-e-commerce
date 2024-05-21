package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

func SignUp(c *fiber.Ctx, db *gorm.DB, v *Validator) error {
	newUser := models.User{}
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
