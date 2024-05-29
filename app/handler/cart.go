package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thapakornd/fiber-go/app/models"
	"gorm.io/gorm"
)

func (h *Handler) CurrentCart(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) AddItem(c *fiber.Ctx) error {

	req := models.APIAddItem{}
	cart_session := models.CartSession{}

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

	// Find User
	user_exits, err := h.userStore.GetByID(uint(req.UserID))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Find product
	if _, err := h.productStore.GetByID(int(req.ProductID)); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := h.cartStore.GetByUserID(int(user_exits.ID), &cart_session); err == gorm.ErrRecordNotFound {

		cart_session.CreatedBy = user_exits.ID
		cart_session.Status = req.Status

		if err := h.cartStore.Create(&cart_session); err != nil {
			fmt.Printf("\n\nEnter this section\n\n")
			return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		new_item := models.CartItem{
			CartSessionID: uint64(cart_session.ID),
			ProductID:     req.ProductID,
			Quantity:      req.Quantity,
		}

		if err := h.cartItemStore.Create(&new_item); err != nil {
			return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

	} else if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	} else {

		cart_item, err := h.cartItemStore.GetByProductID(int(req.ProductID), int(cart_session.ID))
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		cart_item.Quantity = req.Quantity
		fmt.Print(cart_item)
		fmt.Println(req.ProductID)

		if err := h.cartItemStore.Update(cart_item, int(cart_item.ID)); err != nil {
			return c.Status(fiber.ErrConflict.Code).JSON(fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "created successfully",
	})
}

func (h *Handler) RemoveItem(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateItem(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetAllCarts(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) RemoveCart(c *fiber.Ctx) error {
	return nil
}
