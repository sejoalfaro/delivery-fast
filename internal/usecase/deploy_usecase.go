package usecase

import (
	"delivery/internal/domain"
	"delivery/internal/repository"
)

// DeployUseCase is a implementation of the DeployInterface
type DeployUseCase struct {
	deploy repository.DeployInterface
}

// NewDeployUseCase is a function that creates a new DeployUseCase
func NewDeployUseCase(deploy repository.DeployInterface) *DeployUseCase {
	return &DeployUseCase{deploy: deploy}
}

// AddDeploy is a function that adds a new deploy
func (uc *DeployUseCase) AddDeploy(name, application, environment, domainName, traefikRule, version string) error {
	deploy := domain.NewDeploy(name, application, environment, domainName, traefikRule, version)

	return uc.deploy.Add(deploy)
}
