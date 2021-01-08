package service

import (
	. "github.com/architectv/estate-task/pkg/error"
	"github.com/architectv/estate-task/pkg/model"
	"github.com/architectv/estate-task/pkg/repository"
)

type BookingService struct {
	repo     repository.Booking
	roomRepo repository.Room
}

func NewBookingService(repo repository.Booking, roomRepo repository.Room) *BookingService {
	return &BookingService{repo: repo, roomRepo: roomRepo}
}

func (s *BookingService) Create(booking *model.Booking) (int, error) {
	_, err := s.roomRepo.GetById(booking.RoomId)
	if err != nil {
		return 0, ErrWrongRoomId
	}
	if !booking.DateStart.Before(booking.DateEnd) {
		return 0, ErrWrongDates
	}

	return s.repo.Create(booking)
}

func (s *BookingService) Delete(id int) error {
	_, err := s.repo.GetById(id)
	if err != nil {
		return ErrWrongBookingId
	}

	return s.repo.Delete(id)
}

func (s *BookingService) GetByRoomId(roomId int) ([]*model.Booking, error) {
	_, err := s.roomRepo.GetById(roomId)
	if err != nil {
		return nil, ErrWrongRoomId
	}

	return s.repo.GetByRoomId(roomId)
}
