package repository

import (
	"delivery/internal/domain"
)

// DeployInterface is an interface that defines the methods that a deploy repository should implement
type DeployInterface interface {
	Add(deploy *domain.Deploy) error
	RemoveByName(name string) error
	FindAll() ([]*domain.Deploy, error)
	FindByName(name string) (*domain.Deploy, error)
	FindByApplication(application string) ([]*domain.Deploy, error)
}
