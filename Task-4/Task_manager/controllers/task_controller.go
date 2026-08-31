package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"Task_manager/data"
	"Task_manager/models"
)

type TaskController struct {
	service *data.TaskService
}

func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

// GetTasks returns all tasks.
func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks := c.service.GetAllTasks()

	ctx.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

// GetTask returns a single task by ID.
func (c *TaskController) GetTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	task, err := c.service.GetTaskByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// CreateTask creates a new task.
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	createdTask := c.service.CreateTask(task)

	ctx.JSON(http.StatusCreated, createdTask)
}

// UpdateTask updates an existing task.
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	updatedTask, err := c.service.UpdateTask(id, task)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}

// DeleteTask deletes an existing task.
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid task ID",
		})
		return
	}

	err = c.service.DeleteTask(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}
