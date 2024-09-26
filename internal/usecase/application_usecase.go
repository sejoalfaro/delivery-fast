package usecase

import (
	"delivery/internal/domain"
	"delivery/internal/repository"
)

type ApplicationUseCase struct {
	application repository.Application
}

func NewApplicationUseCase(application repository.Application) *ApplicationUseCase {
	return &ApplicationUseCase{application: application}
}

func (uc *ApplicationUseCase) AddApplication(url, name, branch string) error {
	repo := domain.NewRepository(url, name, branch)
	if err := repo.Validar(); err != nil {
		return err
	}

	return uc.application.Add(repo)
}

func (uc *ApplicationUseCase) ListApplications() ([]*domain.Application, error) {
	return uc.application.FindAll()
}

func (uc *ApplicationUseCase) RemoveApplication(name string) error {
	return uc.application.RemoveByName(name)
}
