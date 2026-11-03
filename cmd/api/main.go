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

	repo := repository.NewSQLiteTaskRepository()
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r.GET("/tasks", func(c *gin.Context) { handler.GetAllTasks(db, c) })
	r.POST("/tasks", func(c *gin.Context) { handler.CreateTask(db, c) })
	r.PUT("/tasks/:id", func(c *gin.Context) { handler.UpdateTask(db, c) })
	r.DELETE("/tasks/:id", func(c *gin.Context) { handler.DeleteTask(db, c) })

	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())
	}
}
