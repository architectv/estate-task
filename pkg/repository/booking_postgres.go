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

func (r *BookingPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", bookingsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *BookingPostgres) GetByRoomId(roomId int) ([]*model.Booking, error) {
	var bookings []*model.Booking

	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE room_id=$1 ORDER BY date_start`, bookingsTable)
	err := r.db.Select(&bookings, query, roomId)

	return bookings, err
}
