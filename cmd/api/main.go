package api

import (
	"log"
	"managment-tasks/internal/handler"
	"managment-tasks/internal/repository"
	"managment-tasks/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	repo := repository.NewMemoryTaskRepository()
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r.GET("/all", handler.GetAllTasks)

	err := r.Run(":8080")

	if err != nil {
		log.Fatal(err.Error())
	}
}
