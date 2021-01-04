package repository

import "github.com/jmoiron/sqlx"

type Room interface {
}

type Booking interface {
}

type Repository struct {
	Room
	Booking
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
