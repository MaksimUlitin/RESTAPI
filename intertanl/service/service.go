package service

import "github.com/MaksimUlitin/pkg/repository"

type Athorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Athorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
