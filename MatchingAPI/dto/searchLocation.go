package dto

type SearchLocation struct {
	Radius    int     `json:"radius" bson:"radius"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
