package repository

import (
	"database/sql"
	"delivery/internal/domain"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) Repository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Add(repo *domain.Repository) error {
	stmt, err := r.db.Prepare("INSERT INTO repos (url, branch) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(repo.URL, repo.Branch)
	return err
}

func (r *SQLiteRepository) Remove(url string) error {
	stmt, err := r.db.Prepare("DELETE FROM repos WHERE url = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url)
	return err
}

func (r *SQLiteRepository) FindAll() ([]*domain.Repository, error) {
	rows, err := r.db.Query("SELECT id, url, branch FROM repos")
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
	err := r.db.QueryRow("SELECT id, url, branch FROM repos WHERE url = ?", url).Scan(&repo.ID, &repo.URL, &repo.Branch)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *SQLiteRepository) FindByID(id string) (*domain.Repository, error) {
	var repo domain.Repository
	err := r.db.QueryRow("SELECT id, url, branch FROM repos WHERE id = ?", id).Scan(&repo.ID, &repo.URL, &repo.Branch)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *SQLiteRepository) RemoveByID(id string) error {
	stmt, err := r.db.Prepare("DELETE FROM repos WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
