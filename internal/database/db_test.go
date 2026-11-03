package database

import (
	"database/sql"
	"managment-tasks/internal/entities"
	"testing"
)

func TestEnsureTasksTableMigratesDueColumn(t *testing.T) {
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
			status TEXT NOT NULL
		);
		INSERT INTO tasks (title, description, status) VALUES ('Existing task', 'Before migration', 'pending');
	`); err != nil {
		t.Fatal(err)
	}

	if err := ensureTasksTable(db); err != nil {
		t.Fatal(err)
	}

	tasks, err := GetAllTasks(db)
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected existing task to remain after migration, got %d tasks", len(tasks))
	}
	if tasks[0].Title != "Existing task" || tasks[0].Due != "" {
		t.Fatalf("unexpected migrated task: %+v", tasks[0])
	}

	created := &entities.Task{
		Title:       "New task",
		Description: "After migration",
		Due:         "2026-11-03",
		Status:      "pending",
	}
	if err := CreateTask(db, created); err != nil {
		t.Fatal(err)
	}
	tasks, err = GetAllTasks(db)
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 || tasks[1].Due != created.Due {
		t.Fatalf("new task was not returned correctly after migration: %+v", tasks)
	}
}
