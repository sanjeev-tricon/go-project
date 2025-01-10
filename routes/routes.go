package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/todos", controllers.GetAllTodosHandler)
		api.GET("/todos/:id", controllers.GetTodoByIDHandler)
		api.POST("/todos", controllers.AddTodoHandler)
		api.PUT("/todos/:id", controllers.UpdateTodoHandler)
		api.DELETE("/todos/:id", controllers.DeleteTodoHandler)
	}
}
