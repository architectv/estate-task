package model

import (
	"encoding/json"
	"errors"
	"time"
)

// year-month-day
const DateFormat = "2006-01-02"

type Booking struct {
	Id        int       `json:"booking_id" db:"id"`
	RoomId    int       `json:"-" db:"room_id"`
	DateStart time.Time `json:"date_start" db:"date_start"`
	DateEnd   time.Time `json:"date_end" db:"date_end"`
}

func (b *Booking) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Id        int    `json:"booking_id"`
		DateStart string `json:"date_start"`
		DateEnd   string `json:"date_end"`
	}{
		Id:        b.Id,
		DateStart: b.DateStart.Format(DateFormat),
		DateEnd:   b.DateEnd.Format(DateFormat),
	})
}

func (b *Booking) UnmarshalJSON(data []byte) error {
	var buffer struct {
		RoomId    int    `json:"room_id"`
		DateStart string `json:"date_start"`
		DateEnd   string `json:"date_end"`
	}
	if err := json.Unmarshal(data, &buffer); err != nil {
		return err
	}

	dateStart, err := time.Parse(DateFormat, buffer.DateStart)
	if err != nil {
		return errors.New("bad date_start")
	}
	dateEnd, err := time.Parse(DateFormat, buffer.DateEnd)
	if err != nil {
		return errors.New("bad date_end")
	}

	b.RoomId = buffer.RoomId
	b.DateStart = dateStart
	b.DateEnd = dateEnd

	return nil
}
