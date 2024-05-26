package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
)

func (h *Handler) GetAllProducts(c *fiber.Ctx) error {

	p := []models.APIProduct{}
	limit := 10
	offset := 0

	if err := h.productStore.GetAll(limit, offset, &p); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	fmt.Println(p)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   p,
	})
}

func (h *Handler) GetProduct(c *fiber.Ctx) error {

	return nil
}

func (h *Handler) AddNewProduct(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) RemoveProduct(c *fiber.Ctx) error {
	return nil
}
