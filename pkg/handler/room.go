package handler

import (
	"errors"
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
	sortField := "id"
	asc := true

	sort := ctx.Query("sort")
	if len(sort) > 2 {
		if sort[0] == '-' {
			asc = false
			sort = sort[1:]
		}
	}
	if sort == "price" || sort == "id" {
		sortField = sort
	} else if sort != "" {
		return sendError(ctx, fiber.StatusBadRequest, errors.New("wrong sort param"))
	}

	rooms, err := h.services.Room.GetAllRooms(sortField, asc)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{"rooms": rooms})
}
