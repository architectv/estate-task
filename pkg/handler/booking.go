package handler

import (
	"strconv"

	. "github.com/architectv/estate-task/pkg/error"
	"github.com/architectv/estate-task/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createBooking(ctx *fiber.Ctx) error {
	input := &model.Booking{}
	if err := ctx.BodyParser(input); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}

	id, err := h.services.Booking.Create(input)
	if err != nil {
		if err == ErrWrongRoomId || err == ErrWrongDates {
			return sendError(ctx, fiber.StatusBadRequest, err)
		}
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
		if err == ErrWrongBookingId {
			return sendError(ctx, fiber.StatusBadRequest, err)
		}
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON("OK")
}

func (h *Handler) getBookingsByRoomId(ctx *fiber.Ctx) error {
	roomId, err := strconv.Atoi(ctx.Query("room_id"))
	if err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err)
	}

	bookings, err := h.services.Booking.GetByRoomId(roomId)
	if err != nil {
		if err == ErrWrongRoomId {
			return sendError(ctx, fiber.StatusBadRequest, err)
		}
		return sendError(ctx, fiber.StatusInternalServerError, err)
	}

	return ctx.JSON(bookings)
}
