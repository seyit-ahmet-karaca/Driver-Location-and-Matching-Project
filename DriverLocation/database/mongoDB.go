package database

import (
	"context"
	"driverlocation/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	db *mongo.Database
)

func Connect() (*mongo.Client, *context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Get().MongoDbUrl))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}

	db = client.Database(config.Get().MongoDbName)
	return client, &ctx, err
}
