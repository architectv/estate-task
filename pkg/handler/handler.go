package handler

import (
	"fmt"
	"runtime"

	"github.com/architectv/property-task/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router fiber.Router) {
	rooms := router.Group("/rooms")
	{
		rooms.Post("/", h.createRoom)
		rooms.Delete("/:id", h.deleteRoom)
		rooms.Get("/", h.getAllRooms)
	}
	bookings := router.Group("/bookings")
	{
		bookings.Post("/", h.createBooking)
		bookings.Delete("/:id", h.deleteBooking)
		bookings.Get("/", h.getBookingsByRoomId)
	}
}

func sendError(ctx *fiber.Ctx, status int, err error) error {
	logrus.Error(err.Error())
	ctx.Status(status)
	return ctx.JSON(fiber.Map{"error": err.Error()})
}

func implementMe() string {
	pc, fn, line, _ := runtime.Caller(1)
	return fmt.Sprintf("Implement me in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
}
