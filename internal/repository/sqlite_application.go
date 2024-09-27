package repository

import (
	"database/sql"
	"delivery/internal/domain"
	"errors"
	"fmt"
)

type SQLiteApplication struct {
	db *sql.DB
}

var applicationTableName = "application"

func NewApplicationSQLRepository() ApplicationInterface {
	return &SQLiteApplication{db: DB}
}

func (r *SQLiteApplication) Add(repo *domain.Application) error {
	stmt, err := r.db.Prepare(fmt.Sprintf("INSERT INTO %s (id, name, url, branch) VALUES (?, ?, ?, ?)", applicationTableName))
	if err != nil {
		return fmt.Errorf("error al preparar la inserción: %s", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(repo.ID, repo.Name, repo.URL, repo.Branch)
	if err != nil {
		return fmt.Errorf("error al ejecutar la inserción: %s", err)
	}
	return nil
}

func (r *SQLiteApplication) Remove(url string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE url = ?", applicationTableName)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(url)
	return err
}

func (r *SQLiteApplication) FindAll() ([]*domain.Application, error) {
	var findAllQuery = fmt.Sprintf("SELECT id, name, url, branch FROM %s", applicationTableName)
	rows, err := r.db.Query(findAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []*domain.Application
	for rows.Next() {
		var repo domain.Application
		if err := rows.Scan(&repo.ID, &repo.Name, &repo.URL, &repo.Branch); err != nil {
			return nil, err
		}
		repos = append(repos, &repo)
	}
	return repos, nil
}

func (r *SQLiteApplication) FindByURL(url string) (*domain.Application, error) {
	var repo domain.Application
	var FindByURLQuery = fmt.Sprintf("SELECT id, url, branch FROM %s WHERE url = ?", applicationTableName)
	err := r.db.QueryRow(FindByURLQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteApplication) FindByID(id string) (*domain.Application, error) {
	var repo domain.Application
	var FindByIDQuery = fmt.Sprintf("SELECT id, name, url, branch FROM %s WHERE id = ?", applicationTableName)
	err := r.db.QueryRow(FindByIDQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteApplication) FindByName(name string) (*domain.Application, error) {
	var app domain.Application
	var FindByNameQuery = fmt.Sprintf("SELECT id, name, url, branch FROM %s WHERE name = ?", applicationTableName)
	row := r.db.QueryRow(FindByNameQuery)
	err := row.Scan(&app.ID, &app.Name, &app.URL, &app.Branch)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *SQLiteApplication) RemoveByID(id string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE id = ?", applicationTableName)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func (r *SQLiteApplication) RemoveByName(name string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE name = ?", applicationTableName)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("repository not found")
	}
	return err
}
