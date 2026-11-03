package database

import (
	"log"
	"managment-tasks/internal/entities"
)

func GetAllTasks() ([]*entities.Task, error) {
	db, err := GetOrCreateConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT * FROM tasks")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	tasks := []*entities.Task{}

	for rows.Next() {
		var t entities.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Due); err != nil {
			log.Fatal(err.Error())
		}

		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return tasks, nil
}
