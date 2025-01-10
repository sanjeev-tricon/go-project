package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ObjectId (with BSON tag)
	Title     string             `bson:"title"`         // Title of the Todo item
	Completed bool               `bson:"completed"`     // Status of the Todo item
}