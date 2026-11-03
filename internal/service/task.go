package service

import (
	"database/sql"
	"managment-tasks/internal/entities"
)

type TaskRepository interface {
	Create(db *sql.DB, task *entities.Task) error
	GetByID(db *sql.DB, id int) (*entities.Task, error)
	List(db *sql.DB) ([]*entities.Task, error)
	Update(db *sql.DB, task *entities.Task) error
	Delete(db *sql.DB, id int) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(db *sql.DB, task *entities.Task) error {
	return s.repo.Create(db, task)
}

func (s *TaskService) GetTaskByID(db *sql.DB, id int) (*entities.Task, error) {
	return s.repo.GetByID(db, id)
}

func (s *TaskService) ListTasks(db *sql.DB) ([]*entities.Task, error) {
	return s.repo.List(db)
}

func (s *TaskService) UpdateTask(db *sql.DB, task *entities.Task) error {
	return s.repo.Update(db, task)
}

func (s *TaskService) DeleteTask(db *sql.DB, id int) error {
	return s.repo.Delete(db, id)
}
