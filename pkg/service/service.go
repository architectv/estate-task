package service

import "github.com/architectv/property-task/pkg/repository"

type Room interface {
}

type Booking interface {
}

type Service struct {
	Room
	Booking
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
