package repository

import (
	"context"
	"driverlocation/database"
	"driverlocation/exception"
	"driverlocation/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var lock sync.Mutex
var singleInstance ILocationRepository

type ILocationRepository interface {
	Insert(*model.Point) error
	InsertBulk(items []interface{}) error
	FindDriverInDistanceWithRadius(*model.Location, int) ([]*model.Point, error)
}

type LocationRepository struct {
	collection *mongo.Collection
}

func NewLocationRepository() ILocationRepository {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &LocationRepository{
				collection: database.PointCollection(),
			}
		}
	}
	return singleInstance
}

func (l LocationRepository) Insert(point *model.Point) error {
	var ctx = context.Background()
	_, err := l.collection.InsertOne(ctx, point)
	if err != nil {
		return exception.GetMongoDbInsertException(err.Error())
	}
	return err
}

func (l LocationRepository) InsertBulk(items []interface{}) error {
	var ctx = context.Background()
	_, err := l.collection.InsertMany(ctx, items)
	if err != nil {
		return exception.GetMongoDbInsertException(err.Error())
	}
	return err
}

func (l LocationRepository) FindDriverInDistanceWithRadius(location *model.Location, radius int) ([]*model.Point, error) {
	var ctx = context.Background()
	var results []*model.Point

	cur, err := l.collection.Find(ctx, bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{location.Coordinates[0], location.Coordinates[1]},
				},
				"$maxDistance": radius,
			},
		},
	})

	if err != nil {
		return nil, exception.GetMongoDbReadException(err.Error())
	}

	for cur.Next(ctx) {
		var element model.Point
		err := cur.Decode(&element)
		if err != nil {
			return nil, exception.GetGenericException(err.Error())
		}

		results = append(results, &element)
	}
	return results, nil
}
