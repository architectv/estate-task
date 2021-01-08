package error

import "errors"

var (
	ErrEmptyDescription = errors.New("description should not be empty")
	ErrNotPositivePrice = errors.New("price should be positive number")
	ErrWrongSortField   = errors.New("wrong sort param")
	ErrWrongRoomId      = errors.New("wrong room_id")
	ErrWrongDates       = errors.New("date_start should be before date_end")
	ErrWrongBookingId   = errors.New("wrong booking_id")
	ErrInternalService  = errors.New("something went wrong")
)
