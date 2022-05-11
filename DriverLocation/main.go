package main

import (
	"context"
	"driverlocation/database"
	"driverlocation/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var mongoClient *mongo.Client
var ctx *context.Context

func init() {
	var err error
	mongoClient, ctx, err = database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	mod := mongo.IndexModel{
		Keys: bson.M{
			"location": "2dsphere",
		},
	}

	database.PointCollection().Indexes().CreateOne(*ctx, mod)
}

// @title Driver Location API
// @version 1.0
// @description Driver Location API
// @host localhost:3000
func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}(mongoClient, *ctx)

	s := server.NewServer()
	s.StartServer(3000)
}
