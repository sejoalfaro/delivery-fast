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

var DBFileName = "./delivery.db"
var dbTableBame = "repository"

func NewSQLiteApplication(db *sql.DB) Application {
	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id TEXT PRIMARY KEY NOT NULL,
		name TEXT NOT NULL UNIQUE,
        url TEXT NOT NULL,
        branch TEXT NOT NULL
    )`, dbTableBame)

	_, err := db.Exec(createTableQuery)
	if err != nil {
		fmt.Printf("Error al crear la tabla repos: %v\n", err)
	}
	return &SQLiteApplication{db: db}
}

func (r *SQLiteApplication) Add(repo *domain.Application) error {
	stmt, err := r.db.Prepare(fmt.Sprintf("INSERT INTO %s (id, name, url, branch) VALUES (?, ?, ?, ?)", dbTableBame))
	if err != nil {
		return fmt.Errorf("Error al preparar la inserción: %s", err)
	}
	_, err = stmt.Exec(repo.ID, repo.Name, repo.URL, repo.Branch)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la inserción: %s", err)
	}
	return nil
}

func (r *SQLiteApplication) Remove(url string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE url = ?", dbTableBame)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url)
	return err
}

func (r *SQLiteApplication) FindAll() ([]*domain.Application, error) {
	var findAllQuery = fmt.Sprintf("SELECT id, name, url, branch FROM %s", dbTableBame)
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
	var FindByURLQuery = fmt.Sprintf("SELECT id, url, branch FROM %s WHERE url = ?", dbTableBame)
	err := r.db.QueryRow(FindByURLQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteApplication) FindByID(id string) (*domain.Application, error) {
	var repo domain.Application
	var FindByIDQuery = fmt.Sprintf("SELECT id, name, url, branch FROM %s WHERE id = ?", dbTableBame)
	err := r.db.QueryRow(FindByIDQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteApplication) FindByName(name string) (*domain.Application, error) {
	var repo domain.Application
	var FindByNameQuery = fmt.Sprintf("SELECT id, name, url, branch FROM %s WHERE name = ?", dbTableBame)
	err := r.db.QueryRow(FindByNameQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteApplication) RemoveByID(id string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE id = ?", dbTableBame)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (r *SQLiteApplication) RemoveByName(name string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE name = ?", dbTableBame)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(name)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("repository not found")
	}
	return err
}
