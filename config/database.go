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

func ConnectDb() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = client.Ping(ctx, readpref.Primary())

}
