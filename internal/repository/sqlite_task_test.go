package repository

import (
	"database/sql"
	"managment-tasks/internal/entities"
	"testing"
)

func TestSQLiteTaskRepositoryListsCreatedTasks(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	if _, err := db.Exec(`
		CREATE TABLE tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL,
			due TEXT NOT NULL DEFAULT ''
		)
	`); err != nil {
		t.Fatal(err)
	}

	repo := NewSQLiteTaskRepository()
	created := &entities.Task{
		Title:       "Test task",
		Description: "Created during test",
		Due:         "2026-11-03",
		Status:      "pending",
	}
	if err := repo.Create(db, created); err != nil {
		t.Fatal(err)
	}

	tasks, err := repo.List(db)
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Id != created.Id || tasks[0].Title != created.Title || tasks[0].Due != created.Due {
		t.Fatalf("listed task does not match created task: got %+v, want %+v", tasks[0], created)
	}
}
