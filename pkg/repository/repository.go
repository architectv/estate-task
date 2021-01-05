package repository

import (
	"github.com/architectv/property-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Room interface {
	Create(room *model.Room) (int, error)
	Delete(id int) error
	GetAllRooms(sortField string, asc bool) ([]*model.Room, error)
}

type Booking interface {
}

type Repository struct {
	Room
	Booking
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Room: NewRoomPostgres(db),
	}
}
