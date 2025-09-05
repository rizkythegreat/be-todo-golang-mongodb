package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	// Mongod connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Set username and password
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// setting auth credential
	clientOptions.SetAuth(options.Credential{
		Username:   username,
		Password:   password,
		AuthSource: "admin",
	})

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to MongoDB")
	log.Println("DB_USER:", username)
	log.Println("DB_PASSWORD:", password)
	return client, nil
}
