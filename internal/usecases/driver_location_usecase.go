package usecases

import (
	"context"
	"github.com/volkankocaali/bi-taksi-case/internal/domain"
	"github.com/volkankocaali/bi-taksi-case/internal/repository"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/response"
	"github.com/volkankocaali/bi-taksi-case/pkg/haversine"
	"time"
)

type DriverLocationUseCase struct {
	repo repository.DriverLocationRepository
}

func NewDriverLocationUseCase(repo repository.DriverLocationRepository) *DriverLocationUseCase {
	return &DriverLocationUseCase{
		repo: repo,
	}
}

func (uc *DriverLocationUseCase) CreateOrUpdateDriverLocations(ctx context.Context, location []domain.DriverLocation) error {
	for _, loc := range location {
		loc.Timestamp = time.Now()
		if err := uc.repo.UpsertLocation(ctx, loc); err != nil {
			return err
		}
	}
	return nil
}

func (uc *DriverLocationUseCase) GetLatestDriverLocation(ctx context.Context, driverID string) (*domain.DriverLocation, error) {
	location, err := uc.repo.GetLatestLocation(ctx, driverID)

	if err != nil {
		return nil, err
	}
	return location, nil
}

func (uc *DriverLocationUseCase) FindDriversWithinRadius(ctx context.Context, lat, lon, radius float64) ([]response.Driver, error) {
	point := haversine.NewLocation(lat, lon)
	drivers, err := uc.repo.GetAllDrivers(ctx)
	if err != nil {
		return nil, err
	}

	var driversWithinRadius []response.Driver
	for _, driver := range drivers {
		driverLocation := haversine.NewLocation(driver.Location.Coordinates[0], driver.Location.Coordinates[1])
		distance := haversine.HaversineDistance(point, driverLocation)
		if distance <= radius {
			driversWithinRadius = append(driversWithinRadius, response.Driver{
				ID:        driver.DriverID,
				Latitude:  driver.Location.Coordinates[0],
				Longitude: driver.Location.Coordinates[1],
				Distance:  distance,
			})
		}
	}

	return driversWithinRadius, nil
}
