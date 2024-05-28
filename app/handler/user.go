package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
)

func (h *Handler) CurrentUser(c *fiber.Ctx) error {

	ids, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "something went wrong to server",
		})
	}

	u, err := h.userStore.GetByID(uint(ids))
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   u,
	})
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {

	req := models.APIUser{}
	ids, err := c.ParamsInt("id")
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

	if err := h.userStore.Update(&req, ids); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "successful updated",
	})
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {

	u := []models.APIUser{}
	page, err := strconv.ParseInt(c.Query("page"), 10, 0)
	var offset int
	limit := 10

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "bad request parameter",
		})
	}

	offset = (int(page) - 1) * limit

	total, err := h.userStore.GetAll(limit, offset, &u)
	if err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": "something went wrong to server",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   u,
		"metadata": fiber.Map{
			"total_records": total,
			"current_page":  page,
			"total_pages":   int64(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func (h *Handler) RemoveUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.userStore.Delete(id); err != nil {
		return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "success deleted",
	})
}
