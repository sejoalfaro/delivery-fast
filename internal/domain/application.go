package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Application struct {
	ID     string
	Name   string
	URL    string
	Branch string
}

func NewRepository(url, name, branch string) *Application {
	return &Application{
		ID:     uuid.NewString(),
		URL:    url,
		Name:   name,
		Branch: branch,
	}
}

func (r *Application) Validar() error {
	if r.URL == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	if r.Branch == "" {
		return fmt.Errorf("Branch cannot be empty")
	}
	return nil
}
