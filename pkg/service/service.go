package service

import "cine-base-app/pkg/repository"

type Authorization interface {
}

type MovieList interface {
}

type MovieItem interface {
}

type Service struct {
	Authorization
	MovieList
	MovieItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
