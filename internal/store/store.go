package store

import "managment-tasks/internal/entities"

var tasks []*entities.Task = []*entities.Task{
	{Id: 1, Title: "Task 1", Description: "Description for Task 1", Status: "Planned", Due: "2026-10-09"},
	{Id: 2, Title: "Task 2", Description: "Description for Task 2", Status: "In Progress", Due: "2026-10-10"},
	{Id: 3, Title: "Task 3", Description: "Description for Task 3", Status: "Completed", Due: "2026-10-11"},
}

func GetAll() ([]*entities.Task, error) {
	return tasks, nil
}
