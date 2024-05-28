package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
)

func (h *Handler) GetAllProducts(c *fiber.Ctx) error {

	p := []models.Product{}
	page, err := strconv.ParseInt(c.Query("page"), 10, 0)
	limit := 10
	var offset int

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	offset = (int(page) - 1) * limit

	total, err := h.productStore.GetAll(limit, offset, &p)
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   p,
		"metadata": fiber.Map{
			"total_records": total,
			"current_page":  page,
			"total_pages":   int64(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func (h *Handler) GetProduct(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	p, err := h.productStore.GetByID(id)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   p,
	})
}

func (h *Handler) AddNewProduct(c *fiber.Ctx) error {

	req := models.APIProduct{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.productStore.Create(&req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "success created",
	})
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {

	req := models.APIProduct{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "success",
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
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.productStore.Update(&req, id); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success updated",
	})
}

func (h *Handler) RemoveProduct(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.productStore.Delete(id); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success deleted",
	})
}
