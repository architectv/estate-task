package service

import (
	"github.com/architectv/property-task/pkg/model"
	"github.com/architectv/property-task/pkg/repository"
)

type Room interface {
	Create(room *model.Room) (int, error)
	Delete(id int) error
	GetAllRooms(sortField string, asc bool) ([]*model.Room, error)
}

type Booking interface {
	Create(roomId int, dateStart, dateEnd string) (int, error)
}

type Service struct {
	Room
	Booking
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Room:    NewRoomService(repos.Room),
		Booking: NewBookingService(repos.Booking),
	}
}
