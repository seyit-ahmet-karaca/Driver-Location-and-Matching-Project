package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title" bson:"title"`
	Location Location           `json:"location" bson:"location"`
}

type Location struct {
	Type      string `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
