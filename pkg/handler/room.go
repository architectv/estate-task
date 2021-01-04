package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) createRoom(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}

func (h *Handler) deleteRoom(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}

func (h *Handler) getAllRooms(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}
