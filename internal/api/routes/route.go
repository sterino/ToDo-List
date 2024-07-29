package routes

import (
	"github.com/gin-gonic/gin"
	"todo-list/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, taskHandler *handler.TaskHandler) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", taskHandler.ListTasks)
		tasks.POST("/", taskHandler.CreateTask)
		tasks.GET("/:id", taskHandler.GetTask)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
		tasks.PUT("/:id/done", taskHandler.DoneTask)
	}
}
