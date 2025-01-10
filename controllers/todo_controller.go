package controllers

import (
	"net/http"
	"todo-app/models"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

func GetAllTodosHandler(c *gin.Context) {
	todos, err := services.GetAllTodosService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")

	todo, err := services.GetTodoByIDService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func AddTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation: Check if Title is provided
	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	// Set default value for Completed if not provided
	if !todo.Completed {
		todo.Completed = false
	}

	err := services.AddTodoService(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodoHandler(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation: Check if Title is provided
	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	// Set default value for Completed if not provided
	if !todo.Completed {
		todo.Completed = false
	}

	err := services.UpdateTodoService(id, &todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}

func DeleteTodoHandler(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteTodoService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
