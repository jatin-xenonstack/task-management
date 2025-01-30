package main

import (
	"task-manager/database"
	"task-manager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks/:id", handlers.GetTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)
	r.GET("/tasks", handlers.ListTasks)

	r.Run(":8030") // Start server on port 8080
}
