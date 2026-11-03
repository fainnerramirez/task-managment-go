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

	if err := ensureTasksTable(db); err != nil {
		return nil, err
	}

	fmt.Println("Tabla tasks creada o ya existe.")
	return db, nil
}

func ensureTasksTable(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL,
			due TEXT NOT NULL DEFAULT ''
		);
	`

	if _, err := db.Exec(createTableQuery); err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}

	rows, err := db.Query("PRAGMA table_info(tasks)")
	if err != nil {
		return fmt.Errorf("failed to inspect tasks table: %w", err)
	}
	hasDueColumn := false
	for rows.Next() {
		var cid, notNull, primaryKey int
		var name, columnType string
		var defaultValue sql.NullString
		if err := rows.Scan(&cid, &name, &columnType, &notNull, &defaultValue, &primaryKey); err != nil {
			rows.Close()
			return fmt.Errorf("failed to inspect tasks table columns: %w", err)
		}
		if name == "due" {
			hasDueColumn = true
		}
	}
	if err := rows.Err(); err != nil {
		rows.Close()
		return fmt.Errorf("failed to inspect tasks table columns: %w", err)
	}
	if err := rows.Close(); err != nil {
		return fmt.Errorf("failed to close tasks table inspection: %w", err)
	}
	if !hasDueColumn {
		if _, err := db.Exec("ALTER TABLE tasks ADD COLUMN due TEXT NOT NULL DEFAULT ''"); err != nil {
			return fmt.Errorf("failed to add due column to tasks table: %w", err)
		}
	}

	return nil
}
