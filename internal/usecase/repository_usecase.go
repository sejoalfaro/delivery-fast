package usecase

import (
	"delivery/internal/domain"
	"delivery/internal/repository"
)

type RepositoryUseCase struct {
	repository repository.Repository
}

func NewRepoUseCase(repository repository.Repository) *RepositoryUseCase {
	return &RepositoryUseCase{repository: repository}
}

func (uc *RepositoryUseCase) AddRepository(url, branch string) error {
	repo := &domain.Repository{URL: url, Branch: branch}

	if err := repo.Validar(); err != nil {
		return err
	}

	return uc.repository.Add(repo)
}

func (uc *RepositoryUseCase) ListRepositories() ([]*domain.Repository, error) {
	return uc.repository.FindAll()
}

func (uc *RepositoryUseCase) RemoveRepository(id string) error {
	return uc.repository.RemoveByID(id)
}
