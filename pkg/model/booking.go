package model

import "time"

const DateFormat = "2006-01-02"

type Booking struct {
	Id        int       `json:"id" db:"id"`
	RoomId    int       `json:"room_id" db:"room_id"`
	DateStart time.Time `json:"date_start" db:"date_start"`
	DateEnd   time.Time `json:"date_end" db:"date_end"`
}
