package service

import (
	"errors"
	"time"

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
func (s *BookingService) Create(roomId int, dateStartStr, dateEndStr string) (int, error) {
	// TODO: check if roomId exists
	dateStart, err := time.Parse(model.DateFormat, dateStartStr)
	if err != nil {
		return 0, errors.New("bad date_start")
	}
	dateEnd, err := time.Parse(model.DateFormat, dateEndStr)
	if err != nil {
		return 0, errors.New("bad date_end")
	}
	// TODO: equal dates
	if !dateStart.Before(dateEnd) {
		return 0, errors.New("date_start after date_end")
	}
	booking := &model.Booking{
		RoomId:    roomId,
		DateStart: dateStart,
		DateEnd:   dateEnd,
	}
	return s.repo.Create(booking)
}
