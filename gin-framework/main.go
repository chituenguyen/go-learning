// main.go
package main

import (
	"gin-framework/database"
	"gin-framework/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize JSON database
	db := database.NewJSONDatabase("db.json")

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(db)

	// Setup Gin router with logger & recovery middleware
	router := gin.Default()

	// Welcome route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Task Management API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"tasks": gin.H{
					"GET /api/tasks":     "Get all tasks",
					"GET /api/tasks/:id": "Get task by ID",
					"POST /api/tasks":    "Create new task",
					"PUT /api/tasks/:id": "Update task",
					"DELETE /api/tasks/:id": "Delete task",
				},
			},
		})
	})

	// API routes group
	api := router.Group("/api")
	{
		// Task routes
		tasks := api.Group("/tasks")
		{
			tasks.GET("", taskHandler.GetAllTasks)
			tasks.GET("/:id", taskHandler.GetTaskByID)
			tasks.POST("", taskHandler.CreateTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

	// Start server
	router.Run(":8080") // Listen on port 8080
}
