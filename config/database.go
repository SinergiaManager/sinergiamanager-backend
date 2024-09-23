package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database
var client *mongo.Client

func ConnectDb() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(os.Getenv("MONGO_DB"))

	log.Println("Connected to MongoDB!")
}

func DisconnectDb() {
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB!")
}
