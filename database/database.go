package database

import (
	"context"
	"log"
	"time"

	"github.com/thyagofr/go-rabbit-email-service/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase(config *configs.MongoDBConfig) *mongo.Client {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URL()))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
