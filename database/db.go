package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func InitDB() {
	dbUrl := os.Getenv("MONGODB_URI")

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	DB, err = mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	err = DB.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB is unreachable: ", err)
	}

	log.Println("Connected to MongoDB")
}

func GetCollection(database, collection string) *mongo.Collection {
	return DB.Database(database).Collection(collection)
}
