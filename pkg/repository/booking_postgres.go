package repository

import (
	"fmt"

	"github.com/architectv/property-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type BookingPostgres struct {
	db *sqlx.DB
}

func NewBookingPostgres(db *sqlx.DB) *BookingPostgres {
	return &BookingPostgres{db: db}
}

func (r *BookingPostgres) Create(booking *model.Booking) (int, error) {
	var id int
	query := fmt.Sprintf(
		`INSERT INTO %s (room_id, date_start, date_end) VALUES ($1, $2, $3) RETURNING id`,
		bookingsTable)
	row := r.db.QueryRow(query, booking.RoomId, booking.DateStart, booking.DateEnd)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
