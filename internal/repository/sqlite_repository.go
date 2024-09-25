package repository

import (
	"database/sql"
	"delivery/internal/domain"
	"errors"
	"fmt"
)

type SQLiteRepository struct {
	db *sql.DB
}

var DBFileName = "./delivery.db"
var dbTableBame = "repository"

func NewSQLiteRepo(db *sql.DB) Repository {
	createTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id TEXT PRIMARY KEY NOT NULL,
        url TEXT NOT NULL,
        branch TEXT NOT NULL
    )`, dbTableBame)

	_, err := db.Exec(createTableQuery)
	if err != nil {
		fmt.Printf("Error al crear la tabla repos: %v\n", err)
	}
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Add(repo *domain.Repository) error {
	stmt, err := r.db.Prepare(fmt.Sprintf("INSERT INTO %s (id, url, branch) VALUES (?, ?, ?)", dbTableBame))
	if err != nil {
		return fmt.Errorf("Error al preparar la inserción: %v", err)
	}
	_, err = stmt.Exec(repo.ID, repo.URL, repo.Branch)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la inserción: %v", err)
	}
	return nil
}

func (r *SQLiteRepository) Remove(url string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE url = ?", dbTableBame)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url)
	return err
}

func (r *SQLiteRepository) FindAll() ([]*domain.Repository, error) {
	var findAllQuery = fmt.Sprintf("SELECT id, url, branch FROM %s", dbTableBame)
	rows, err := r.db.Query(findAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repos []*domain.Repository
	for rows.Next() {
		var repo domain.Repository
		if err := rows.Scan(&repo.ID, &repo.URL, &repo.Branch); err != nil {
			return nil, err
		}
		repos = append(repos, &repo)
	}
	return repos, nil
}

func (r *SQLiteRepository) FindByURL(url string) (*domain.Repository, error) {
	var repo domain.Repository
	var FindByURLQuery = fmt.Sprintf("SELECT id, url, branch FROM %s WHERE url = ?", dbTableBame)
	err := r.db.QueryRow(FindByURLQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteRepository) FindByID(id string) (*domain.Repository, error) {
	var repo domain.Repository
	var FindByIDQuery = fmt.Sprintf("SELECT id, url, branch FROM %s WHERE id = ?", dbTableBame)
	err := r.db.QueryRow(FindByIDQuery)
	if err != nil {
		return nil, errors.New("repository not found")
	}
	return &repo, nil
}

func (r *SQLiteRepository) RemoveByID(id string) error {
	var deleteQuery = fmt.Sprintf("DELETE FROM %s WHERE id = ?", dbTableBame)
	stmt, err := r.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
