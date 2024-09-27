package repository

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DBFileName string = "delivery.db"
var DBDriver string = "sqlite"
var DB *sql.DB = nil

func initDB() error {
	var err error
	DB, err = sql.Open(DBDriver, DBFileName)
	if err != nil {
		fmt.Println("could not open the database: " + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns((5))

	return nil
}

func createTableApplication() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS application(
	id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL,
	branch TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := DB.Exec(createTableQuery)
	if err != nil {
		fmt.Printf("Error al crear la tabla application: %v\n", err)
	}
	return nil
}

func createTableDeployment() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS deployment (
	id TEXT PRIMARY KEY NOT NULL,
	application_id TEXT NOT NULL,
	environment TEXT NULL,
	version TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(application_id) REFERENCES application(id)
	)
	`

	_, err := DB.Exec(createTableQuery)
	if err != nil {
		fmt.Printf("Error creating deploymnet table: %v\n", err)
	}
	return nil
}

func init() {
	initDB()
	createTableApplication()
	createTableDeployment()
}
