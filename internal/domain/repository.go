package domain

import "fmt"

type Repository struct {
	ID     int
	URL    string
	Branch string
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
