package repository

import "delivery/internal/domain"

// ApplicationInterface is an interface to interact with the application repository
type ApplicationInterface interface {
	Add(repository *domain.Application) error
	RemoveByID(id string) error
	RemoveByName(name string) error
	FindAll() ([]*domain.Application, error)
	FindByName(name string) (*domain.Application, error)
	FindByURL(url string) (*domain.Application, error)
	FindByID(url string) (*domain.Application, error)
}
