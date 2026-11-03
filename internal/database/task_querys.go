package database

import (
	"database/sql"
	"fmt"
	"managment-tasks/internal/entities"
)

func GetAllTasks(db *sql.DB) ([]*entities.Task, error) {
	rows, err := db.Query("SELECT id, title, description, status, due FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	defer rows.Close()

	tasks := []*entities.Task{}

	for rows.Next() {
		var t entities.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Due); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	return tasks, nil
}

func CreateTask(db *sql.DB, task *entities.Task) error {
	result, err := db.Exec("INSERT INTO tasks (title, description, status, due) VALUES (?, ?, ?, ?)", task.Title, task.Description, task.Status, task.Due)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get created task ID: %w", err)
	}
	task.Id = int(id)
	return nil
}

func GetById(db *sql.DB, id int) (*entities.Task, error) {
	row := db.QueryRow("SELECT id, title, description, status, due FROM tasks WHERE id = ?", id)
	var task entities.Task
	if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Due); err != nil {
		return nil, fmt.Errorf("error al obtener la tarea con ID %d: %w", id, err)
	}
	return &task, nil
}

func UpdateTask(db *sql.DB, task *entities.Task) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, due = ? WHERE id = ?", task.Title, task.Description, task.Status, task.Due, task.Id)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func DeleteTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
