package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createBooking(ctx *fiber.Ctx) error {
	type InputBooking struct {
		RoomId    int    `json:"room_id" xml:"room_id" form:"room_id"`
		DateStart string `json:"date_start" xml:"date_start" form:"date_start"`
		DateEnd   string `json:"date_end" xml:"date_end" form:"date_end"`
	}
	input := &InputBooking{}
	if err := ctx.BodyParser(input); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	id, err := h.services.Booking.Create(input.RoomId, input.DateStart, input.DateEnd)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{"booking_id": id})
}

func (h *Handler) deleteBooking(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}

	err = h.services.Booking.Delete(id)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(fiber.Map{})
}

func (h *Handler) getBookingsByRoomId(ctx *fiber.Ctx) error {
	return ctx.SendString(implementMe())
}
