package handler

import (
	"github.com/architectv/property-task/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createRoom(ctx *fiber.Ctx) error {
	input := &model.Room{}
	if err := ctx.BodyParser(input); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}
	logrus.Println(string(ctx.Body()))

	id, err := h.services.Room.Create(input)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{"room_id": id})
}

func (h *Handler) deleteRoom(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}

func (h *Handler) getAllRooms(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}
