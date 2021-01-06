package handler

import (
	"strconv"

	. "github.com/architectv/property-task/pkg/error"
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
		if err == ErrEmptyDescription || err == ErrNotPositivePrice {
			return sendError(ctx, fiber.StatusBadRequest, err)
		}
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{"room_id": id})
}

func (h *Handler) deleteRoom(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}

	err = h.services.Room.Delete(id)
	if err != nil {
		if err == ErrWrongRoomId {
			return sendError(ctx, fiber.StatusBadRequest, err)
		}
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{})
}

func (h *Handler) getAllRooms(ctx *fiber.Ctx) error {
	sortField := ctx.Query("sort")

	rooms, err := h.services.Room.GetAll(sortField)
	if err != nil {
		if err == ErrWrongSortField {
			return sendError(ctx, fiber.StatusBadRequest, err)
		}
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(rooms)
}
