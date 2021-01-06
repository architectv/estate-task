package service

import (
	. "github.com/architectv/property-task/pkg/error"
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
		return 0, ErrEmptyDescription
	}
	if room.Price <= 0 {
		return 0, ErrNotPositivePrice
	}

	return s.repo.Create(room)
}

func (s *RoomService) Delete(id int) error {
	_, err := s.repo.GetById(id)
	if err != nil {
		return ErrWrongRoomId
	}

	return s.repo.Delete(id)
}

func (s *RoomService) GetAll(sortField string) ([]*model.Room, error) {
	const (
		idField    = "id"
		priceField = "price"
	)
	desc := false

	if sortField == "" {
		sortField = idField
	} else if len(sortField) >= 2 {
		if sortField[0] == '-' {
			desc = true
			sortField = sortField[1:]
		}
		if sortField != idField && sortField != priceField {
			return nil, ErrWrongSortField
		}
	} else {
		return nil, ErrWrongSortField
	}

	return s.repo.GetAll(sortField, desc)
}
