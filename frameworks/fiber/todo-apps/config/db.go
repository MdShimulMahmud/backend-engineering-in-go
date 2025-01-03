package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")
	
	if MONGODB_URI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}