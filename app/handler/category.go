package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
)

func (h *Handler) GetAllCategories(c *fiber.Ctx) error {

	css := []models.Category{}

	if err := h.categoryStore.GetAllCategories(&css); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   css,
	})
}

func (h *Handler) AddCategory(c *fiber.Ctx) error {

	req := models.Category{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.categoryStore.Create(&req); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "created successfully",
	})
}

func (h *Handler) RemoveCategory(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.categoryStore.Delete(id); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success deleted category",
	})
}

func (h *Handler) ChangeCategoryInfo(c *fiber.Ctx) error {

	req := models.Category{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail-body",
			"message": err.Error(),
		})
	}

	if err := h.categoryStore.Update(&req, id); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success updated category",
	})
}
