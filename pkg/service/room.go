package service

import (
	"errors"

	"github.com/architectv/property-task/pkg/model"
	"github.com/architectv/property-task/pkg/repository"
)

type RoomService struct {
	repo repository.Room
}

func NewRoomService(repo repository.Room) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(room *model.Room) (int, error) {
	if room.Description == "" {
		return 0, errors.New("description should not be empty")
	}
	if room.Price <= 0 {
		return 0, errors.New("price should be positive number")
	}
	return s.repo.Create(room)
}
