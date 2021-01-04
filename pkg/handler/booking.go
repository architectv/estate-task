package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) createBooking(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}

func (h *Handler) deleteBooking(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}

func (h *Handler) getBookingsByRoomId(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}
