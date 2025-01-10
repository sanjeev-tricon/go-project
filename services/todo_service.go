package services

import (
	"todo-app/models"
	"todo-app/repositories"
)

func AddTodoService(todo *models.Todo) error {
	return repositories.CreateTodo(todo)
}

func UpdateTodoService(id string, updatedTodo *models.Todo) error {
	return repositories.UpdateTodo(id, updatedTodo)
}

func DeleteTodoService(id string) error {
	return repositories.DeleteTodo(id)
}

func GetAllTodosService() ([]models.Todo, error) {
	return repositories.GetAllTodos()
}

func GetTodoByIDService(id string) (*models.Todo, error) {
	return repositories.GetTodoByID(id)
}
