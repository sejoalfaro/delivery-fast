package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
	ID     string
	URL    string
	Branch string
}

func NewRepository(url, branch string) *Repository {
	return &Repository{
		ID:     uuid.NewString(),
		URL:    url,
		Branch: branch,
	}
}

func (r *Repository) Validar() error {
	if r.URL == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	if r.Branch == "" {
		return fmt.Errorf("Branch cannot be empty")
	}
	return nil
}
