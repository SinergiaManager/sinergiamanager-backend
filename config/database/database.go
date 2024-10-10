package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DB     *mongo.Database
	client *mongo.Client
)

func ConnectDb() error {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		return fmt.Errorf("failed to connect MongoDB: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	DB = client.Database(os.Getenv("MONGO_DB"))

	log.Println("Connected to MongoDB!")

	return nil
}

func DisconnectDb() {
	if client == nil {
		log.Println("No MongoDB client")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
		return
	}

	log.Println("Disconnected from MongoDB")
}
