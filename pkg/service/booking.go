package service

import (
	"errors"

	"github.com/architectv/property-task/pkg/model"
	"github.com/architectv/property-task/pkg/repository"
)

type BookingService struct {
	repo repository.Booking
}

func NewBookingService(repo repository.Booking) *BookingService {
	return &BookingService{repo: repo}
}

// TODO: errors in separate file
func (s *BookingService) Create(booking *model.Booking) (int, error) {
	// TODO: check if roomId exists
	// dateStart, err := time.Parse(model.DateFormat, dateStartStr)
	// if err != nil {
	// 	return 0, errors.New("bad date_start")
	// }
	// dateEnd, err := time.Parse(model.DateFormat, dateEndStr)
	// if err != nil {
	// 	return 0, errors.New("bad date_end")
	// }
	if booking.RoomId <= 0 {
		return 0, errors.New("room_id should be positive")
	}
	if booking.DateStart.After(booking.DateEnd) {
		return 0, errors.New("date_start after date_end")
	}
	// booking := &model.Booking{
	// 	RoomId:    roomId,
	// 	DateStart: dateStart,
	// 	DateEnd:   dateEnd,
	// }
	return s.repo.Create(booking)
}

func (s *BookingService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id should be positive")
	}
	return s.repo.Delete(id)
}

func (s *BookingService) GetByRoomId(roomId int) ([]*model.Booking, error) {
	if roomId <= 0 {
		return nil, errors.New("room_id should be positive")
	}
	return s.repo.GetByRoomId(roomId)
}
