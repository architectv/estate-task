package repository

import (
	"fmt"

	"github.com/architectv/property-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{db: db}
}

func (r *RoomPostgres) Create(room *model.Room) (int, error) {
	var id int
	query := fmt.Sprintf(
		`INSERT INTO %s (description, price) VALUES ($1, $2) RETURNING id`,
		roomsTable)
	row := r.db.QueryRow(query, room.Description, room.Price)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
