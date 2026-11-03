package repository

import (
	"database/sql"
	"managment-tasks/internal/database"
	"managment-tasks/internal/entities"
)

type SQLiteTaskRepository struct{}

func NewSQLiteTaskRepository() *SQLiteTaskRepository {
	return &SQLiteTaskRepository{}
}

func (r *SQLiteTaskRepository) Create(db *sql.DB, task *entities.Task) error {
	return database.CreateTask(db, task)
}

func (r *SQLiteTaskRepository) GetByID(db *sql.DB, id int) (*entities.Task, error) {
	return database.GetById(db, id)
}

func (r *SQLiteTaskRepository) List(db *sql.DB) ([]*entities.Task, error) {
	return database.GetAllTasks(db)
}

func (r *SQLiteTaskRepository) Update(db *sql.DB, task *entities.Task) error {
	return database.UpdateTask(db, task)
}

func (r *SQLiteTaskRepository) Delete(db *sql.DB, id int) error {
	return database.DeleteTask(db, id)
}
