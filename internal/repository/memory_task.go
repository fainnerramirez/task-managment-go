package repository

import (
	"managment-tasks/internal/entities"
	"sync"
)

type MemoryTaskRepository struct {
	mu     sync.Mutex
	tasks  []*entities.Task
	nextID int
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks:  []*entities.Task{},
		nextID: 1,
	}
}

func (r *MemoryTaskRepository) Create(task *entities.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.Id = r.nextID
	r.tasks = append(r.tasks, task)
	r.nextID++
	return nil
}

func (r *MemoryTaskRepository) GetByID(id int) (*entities.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, task := range r.tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return nil, nil
}

func (r *MemoryTaskRepository) List() ([]*entities.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.tasks, nil
}

func (r *MemoryTaskRepository) Update(task *entities.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, t := range r.tasks {
		if t.Id == task.Id {
			r.tasks[i] = task
			return nil
		}
	}
	return nil
}

func (r *MemoryTaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, task := range r.tasks {
		if task.Id == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}
