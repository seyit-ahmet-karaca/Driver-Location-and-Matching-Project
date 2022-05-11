package model

type Point struct {
	Title    string             `json:"title" bson:"title"`
	Location Location           `json:"location" bson:"location"`
}

type Location struct {
	Type      string `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
