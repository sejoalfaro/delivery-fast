package repository

import (
	"database/sql"
	"delivery/internal/domain"
)

// DepolySQLiteRepository is a struct that represents the deploy repository
type DepolySQLiteRepository struct {
	db *sql.DB
}

// NewDeploySQLiteRepository is a function that creates a new deploy repository
func NewDeploySQLiteRepository(db *sql.DB) DeployInterface {
	return &DepolySQLiteRepository{db: db}
}

// Add is a function that adds a new deploy
func (r *DepolySQLiteRepository) Add(deploy *domain.Deploy) error {
	var insertQuery = `
	INSERT INTO deploy (name, application, environment, domain, traefik_rule, version) 
	VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(deploy.Name, deploy.Application, deploy.Environment, deploy.Domain, deploy.TraefikRule, deploy.Version)
	return err
}

// RemoveByName is a function that removes a deploy by name
func (r *DepolySQLiteRepository) RemoveByName(name string) error {
	var deleteQuery = "DELETE FROM deploy WHERE name = ?"
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	return err
}

// FindAll is a function that returns all deploys
func (r *DepolySQLiteRepository) FindAll() ([]*domain.Deploy, error) {
	var findAllQuery = "SELECT name, application, environment, domain, traefik_rule, version FROM deploy"
	rows, err := r.db.Query(findAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	deployments := make([]*domain.Deploy, 0)
	for rows.Next() {
		deploy := &domain.Deploy{}
		if err := rows.Scan(&deploy.Name, &deploy.Application, &deploy.Environment, &deploy.Domain, &deploy.TraefikRule, &deploy.Version); err != nil {
			return nil, err
		}
		deployments = append(deployments, deploy)
	}
	return deployments, nil
}

// FindByName is a function that returns a deploy by name
func (r *DepolySQLiteRepository) FindByName(name string) (*domain.Deploy, error) {
	var findOneQuery = "SELECT name, application, environment, domain, traefik_rule, version FROM deploy WHERE name = ?"
	row := r.db.QueryRow(findOneQuery, name)
	deploy := &domain.Deploy{}
	if err := row.Scan(&deploy.Name, &deploy.Application, &deploy.Environment, &deploy.Domain, &deploy.TraefikRule, &deploy.Version); err != nil {
		return nil, err
	}
	return deploy, nil
}

// FindByApplication is a function that returns all deploys by application
func (r *DepolySQLiteRepository) FindByApplication(application string) ([]*domain.Deploy, error) {
	findQuery := "SELECT name, application, environment, domain, traefik_rule, version FROM deploy WHERE application = ?"
	rows, err := r.db.Query(findQuery, application)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	deployments := make([]*domain.Deploy, 0)
	for rows.Next() {
		deploy := &domain.Deploy{}
		if err := rows.Scan(&deploy.Name, &deploy.Application, &deploy.Environment, &deploy.Domain, &deploy.TraefikRule, &deploy.Version); err != nil {
			return nil, err
		}
		deployments = append(deployments, deploy)
	}
	return deployments, nil
}
