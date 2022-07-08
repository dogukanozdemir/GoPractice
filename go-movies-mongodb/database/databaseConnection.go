package database

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}
