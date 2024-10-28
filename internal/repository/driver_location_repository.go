package repository

import (
	"context"
	"errors"
	"github.com/volkankocaali/bi-taksi-case/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DriverLocationRepository interface {
	UpsertLocation(ctx context.Context, location domain.DriverLocation) error
	GetLatestLocation(ctx context.Context, driverID string) (*domain.DriverLocation, error)
	GetAllDrivers(ctx context.Context) ([]domain.DriverLocation, error)
}

type MongoDriverLocationRepository struct {
	collection *mongo.Collection
}

func NewMongoDriverLocationRepository(collection *mongo.Collection) *MongoDriverLocationRepository {
	return &MongoDriverLocationRepository{
		collection: collection,
	}
}

func (repo *MongoDriverLocationRepository) UpsertLocation(ctx context.Context, location domain.DriverLocation) error {
	filter := bson.D{{"driver_id", location.DriverID}}
	update := bson.M{
		"$set": bson.M{
			"location":  domain.GeoJSON{Type: "Point", Coordinates: location.Location.Coordinates},
			"timestamp": location.Timestamp,
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := repo.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (repo *MongoDriverLocationRepository) FindDriversByLocation(ctx context.Context, location domain.GeoJSON, radius float64) ([]domain.DriverLocation, error) {
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry":    location,
				"$maxDistance": radius,
			},
		},
	}

	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var locations []domain.DriverLocation
	for cursor.Next(ctx) {
		var location domain.DriverLocation
		if err := cursor.Decode(&location); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func (repo *MongoDriverLocationRepository) GetLatestLocation(ctx context.Context, driverID string) (*domain.DriverLocation, error) {
	filter := bson.D{{"driver_id", driverID}}
	opts := options.FindOne().SetSort(bson.D{{"timestamp", -1}})
	result := repo.collection.FindOne(ctx, filter, opts)

	var location domain.DriverLocation
	if err := result.Decode(&location); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("location not found")
		}
		return nil, err
	}

	return &location, nil
}

func (repo *MongoDriverLocationRepository) GetAllDrivers(ctx context.Context) ([]domain.DriverLocation, error) {
	var drivers []domain.DriverLocation
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var driver domain.DriverLocation
		if err := cursor.Decode(&driver); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}
