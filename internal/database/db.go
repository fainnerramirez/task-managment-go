package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func GetOrCreateConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the SQLite database successfully.")

	return db, nil
}

func InitializeDatabase() (*sql.DB, error) {

	db, err := GetOrCreateConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL
		);
	`

	_, errorExe := db.Exec(createTableQuery)
	if errorExe != nil {
		return nil, fmt.Errorf("failed to create tasks table: %v", errorExe)
	}

	fmt.Println("Tabla tasks creada o ya existe.")
	return db, nil
}
