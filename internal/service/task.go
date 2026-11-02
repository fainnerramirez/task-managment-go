package service

import (
	"managment-tasks/internal/entities"
)

type TaskRepository interface {
	Create(task *entities.Task) error
	GetByID(id int) (*entities.Task, error)
	List() ([]*entities.Task, error)
	Update(task *entities.Task) error
	Delete(id int) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *entities.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) GetTaskByID(id int) (*entities.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) ListTasks() ([]*entities.Task, error) {
	return s.repo.List()
}

func (s *TaskService) UpdateTask(task *entities.Task) error {
	return s.repo.Update(task)
}
