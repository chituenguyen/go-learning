package handlers

import (
	"gin-framework/database"
	"gin-framework/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	db *database.JSONDatabase
}

func NewTaskHandler(db *database.JSONDatabase) *TaskHandler {
	return &TaskHandler{db: db}
}

// GetAllTasks retrieves all tasks
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	if err := h.db.ReadData(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  tasks,
		"count": len(tasks),
	})
}

// GetTaskByID retrieves a single task by ID
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var tasks []models.Task
	if err := h.db.ReadData(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read tasks"})
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": task})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// CreateTask creates a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input models.CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tasks []models.Task
	if err := h.db.ReadData(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read tasks"})
		return
	}

	// Generate new ID
	newID := 1
	if len(tasks) > 0 {
		// Find max ID
		for _, task := range tasks {
			if task.ID >= newID {
				newID = task.ID + 1
			}
		}
	}

	newTask := models.Task{
		ID:          newID,
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := h.db.WriteData(tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"data":    newTask,
	})
}

// UpdateTask updates an existing task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var input models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tasks []models.Task
	if err := h.db.ReadData(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read tasks"})
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			found = true

			// Update only provided fields
			if input.Title != nil {
				tasks[i].Title = *input.Title
			}
			if input.Description != nil {
				tasks[i].Description = *input.Description
			}
			if input.Completed != nil {
				tasks[i].Completed = *input.Completed
			}
			tasks[i].UpdatedAt = time.Now()

			if err := h.db.WriteData(tasks); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Task updated successfully",
				"data":    tasks[i],
			})
			return
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

// DeleteTask deletes a task by ID
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var tasks []models.Task
	if err := h.db.ReadData(&tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read tasks"})
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			found = true
			// Remove task from slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := h.db.WriteData(tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
