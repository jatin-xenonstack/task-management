package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

// Create Task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := "INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	id, _ := result.LastInsertId()
	task.ID = int(id)
	c.JSON(http.StatusCreated, task)
}

// Retrieve Task
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	query := "SELECT id, title, description, due_date, status FROM tasks WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// Update Task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(task)

	query := "UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?"
	_, err := database.DB.Exec(query, task.Title, task.Description, task.DueDate, task.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	converted_id, _ := strconv.Atoi(id)

	task.ID = converted_id
	c.JSON(http.StatusOK, task)
}

// Delete Task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM tasks WHERE id = ?"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	changedRows, _ := result.RowsAffected()
	if changedRows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// List All Tasks
func ListTasks(c *gin.Context) {
	query := "SELECT id, title, description, due_date, status FROM tasks"
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch"})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Tasks Not fetched Successfully"})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}
