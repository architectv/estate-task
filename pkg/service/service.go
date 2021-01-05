package service

import (
	"github.com/architectv/property-task/pkg/model"
	"github.com/architectv/property-task/pkg/repository"
)

type Room interface {
	Create(room *model.Room) (int, error)
	Delete(id int) error
	GetAll(sortField string, asc bool) ([]*model.Room, error)
}

type Booking interface {
	Create(booking *model.Booking) (int, error)
	Delete(id int) error
	GetByRoomId(roomId int) ([]*model.Booking, error)
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
