package repository

import "delivery/internal/domain"

// RepoRepository es la interfaz que define las operaciones CRUD para los repositorios.
type Repository interface {
	Add(repository *domain.Repository) error
	RemoveByID(id string) error
	FindAll() ([]*domain.Repository, error)
	FindByURL(url string) (*domain.Repository, error)
	FindByID(url string) (*domain.Repository, error)
}
