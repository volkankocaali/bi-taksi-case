package domain

import "time"

type DriverLocation struct {
	DriverID  string    `json:"driver_id" bson:"driver_id"`
	Location  GeoJSON   `json:"location" bson:"location"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type GeoJSON struct {
	Type        string     `json:"type" bson:"type"`
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}
