package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

type ProductQueries struct {
	*gorm.DB
}

func AllProducts(c *fiber.Ctx, db *gorm.DB) error {
	products := []models.Product{}
	db.Find(&products)
	return c.Status(200).JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(products),
		"products": products,
	})
}
