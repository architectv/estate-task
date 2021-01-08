package repository

import (
	"github.com/architectv/estate-task/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Room interface {
	Create(room *model.Room) (int, error)
	Delete(id int) error
	GetAll(sortField string, desc bool) ([]*model.Room, error)
	GetById(id int) (*model.Room, error)
}

type Booking interface {
	Create(booking *model.Booking) (int, error)
	Delete(id int) error
	GetByRoomId(roomId int) ([]*model.Booking, error)
	GetById(id int) (*model.Booking, error)
}

type Repository struct {
	Room
	Booking
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Room:    NewRoomPostgres(db),
		Booking: NewBookingPostgres(db),
	}
}
