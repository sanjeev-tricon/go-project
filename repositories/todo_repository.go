package repositories

import (
	"context"
	"time"
	"todo-app/database"
	"todo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	databaseName   = "todoapp"
	collectionName = "todos"
)

func GetAllTodos() ([]models.Todo, error) {
	collection := database.GetCollection(databaseName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todos []models.Todo
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodoByID(id string) (*models.Todo, error) {
	collection := database.GetCollection(databaseName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var todo models.Todo
	err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func CreateTodo(todo *models.Todo) error {
	collection := database.GetCollection(databaseName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, todo)
	return err
}

func UpdateTodo(id string, updatedTodo *models.Todo) error {
	collection := database.GetCollection(databaseName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": updatedTodo})
	return err
}

func DeleteTodo(id string) error {
	collection := database.GetCollection(databaseName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}
