package handler

import (
	"strconv"

	"github.com/architectv/property-task/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createRoom(ctx *fiber.Ctx) error {
	input := &model.Room{}
	if err := ctx.BodyParser(input); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}

	id, err := h.services.Room.Create(input)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{"room_id": id})
}

func (h *Handler) deleteRoom(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}

	err = h.services.Room.Delete(id)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{})
}

func (h *Handler) getAllRooms(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}
