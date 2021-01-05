package service

import (
	"github.com/architectv/property-task/pkg/model"
	"github.com/architectv/property-task/pkg/repository"
)

type Room interface {
	Create(room *model.Room) (int, error)
	Delete(id int) error
}

type Booking interface {
}

type Service struct {
	Room
	Booking
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Room: NewRoomService(repos.Room),
	}
}
