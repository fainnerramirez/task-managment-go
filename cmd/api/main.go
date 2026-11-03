package main

import (
	"log"
	"managment-tasks/internal/database"
	"managment-tasks/internal/handler"
	"managment-tasks/internal/repository"
	"managment-tasks/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db, errDB := database.InitializeDatabase()
	if errDB != nil {
		log.Fatal(errDB.Error())
	}

	defer db.Close()

	r := gin.Default()

	repo := repository.NewMemoryTaskRepository()
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r.GET("/tasks", handler.GetAllTasks)
	r.POST("/tasks", handler.CreateTask)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())
	}
}
