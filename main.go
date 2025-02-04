package main

import (
	"task-manager/database"
	"task-manager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	router := gin.Default()

	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks/:id", handlers.GetTask)
	router.PUT("/tasks/:id", handlers.UpdateTask)
	router.DELETE("/tasks/:id", handlers.DeleteTask)
	router.GET("/tasks", handlers.ListTasks)
	router.GET("/tasks/pending", handlers.PendingTasks)

	router.Run(":8030") // Start server on port 8080
}
