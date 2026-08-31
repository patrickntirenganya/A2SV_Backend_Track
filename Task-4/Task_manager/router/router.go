package router

import (
	"github.com/gin-gonic/gin"

	"Task_manager/controllers"
	"Task_manager/data"
)

func SetupRouter(service *data.TaskService) *gin.Engine {
	r := gin.Default()

	controller := controllers.NewTaskController(service)

	tasks := r.Group("/tasks")
	{
		tasks.GET("", controller.GetTasks)
		tasks.GET("/:id", controller.GetTask)
		tasks.POST("", controller.CreateTask)
		tasks.PUT("/:id", controller.UpdateTask)
		tasks.DELETE("/:id", controller.DeleteTask)
	}

	return r
}
