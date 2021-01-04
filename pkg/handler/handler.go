package handler

import (
	"fmt"
	"runtime"

	"github.com/architectv/property-task/pkg/service"
	"github.com/gofiber/fiber/v2"
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
		rooms.Post("/create", h.createRoom)
		rooms.Delete("/:id", h.deleteRoom)
		rooms.Get("/list", h.getAllRooms)
	}
	bookings := router.Group("/bookings")
	{
		bookings.Post("/create", h.createBooking)
		bookings.Delete("/:id", h.deleteBooking)
		bookings.Get("/list", h.getBookingsByRoomId)
	}
}

func implementMe() string {
	pc, fn, line, _ := runtime.Caller(1)
	return fmt.Sprintf("Implement me in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
}
