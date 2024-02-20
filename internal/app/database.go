package app

import (
	"context"
	"log"
	"time"

	"github.com/mystpen/weather-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func NewMongoDB(config *config.Config) *mongo.Client {
	// Set client options
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("There was a connection error", err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	client.Disconnect(ctx)
	log.Println("Connection to MongoDB closed.")
}

func NewMongoCollection(client *mongo.Client, config *config.Config, collectionName string) *mongo.Collection {
	return client.Database(config.DBName).Collection(collectionName)
}
