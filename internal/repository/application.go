package repository

import "delivery/internal/domain"

// RepoRepository es la interfaz que define las operaciones CRUD para los repositorios.
type Application interface {
	Add(repository *domain.Application) error
	RemoveByID(id string) error
	RemoveByName(name string) error
	FindAll() ([]*domain.Application, error)
	FindByName(name string) (*domain.Application, error)
	FindByURL(url string) (*domain.Application, error)
	FindByID(url string) (*domain.Application, error)
}
