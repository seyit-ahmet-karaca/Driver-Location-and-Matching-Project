package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	point = "point"
)

func PointCollection() *mongo.Collection {
	return db.Collection(point)
}
