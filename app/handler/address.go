package controllers

import "github.com/gofiber/fiber/v2"

func (h *Handler) CurrentAddresses(c *fiber.Ctx) error {
	h.addressStore.GetByUserID(1)
	return nil
}

func (h *Handler) AddAddress(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) RemoveAddress(c *fiber.Ctx) error {
	return nil
}
