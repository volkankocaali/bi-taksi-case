package usecases

import (
	"context"
	"github.com/volkankocaali/bi-taksi-case/internal/domain"
	"github.com/volkankocaali/bi-taksi-case/internal/repository"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/response"
	"github.com/volkankocaali/bi-taksi-case/pkg/haversine"
	"sort"
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

func (uc *DriverLocationUseCase) CreateOrUpdateDriverLocations(ctx context.Context, locations []domain.DriverLocation) error {
	if len(locations) == 1 {
		locations[0].Timestamp = time.Now()
		return uc.repo.UpsertLocation(ctx, locations[0])
	}

	workerCount := calculateWorkerCount(len(locations))

	locationChan := make(chan domain.DriverLocation, workerCount)
	errChan := make(chan error, len(locations))
	doneChan := make(chan struct{})

	for i := 0; i < workerCount; i++ {
		go func() {
			for loc := range locationChan {
				loc.Timestamp = time.Now()
				if err := uc.repo.UpsertLocation(ctx, loc); err != nil {
					errChan <- err
					return
				}
			}
			doneChan <- struct{}{}
		}()
	}

	go func() {
		for _, loc := range locations {
			locationChan <- loc
		}
		close(locationChan)
	}()

	for i := 0; i < workerCount; i++ {
		<-doneChan
	}
	close(errChan)
	close(doneChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func calculateWorkerCount(dataCount int) int {
	switch {
	case dataCount < 100:
		return 1
	case dataCount < 1000:
		return 5
	default:
		return 10
	}
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

	sort.SliceStable(driversWithinRadius, func(i, j int) bool {
		return driversWithinRadius[i].Distance < driversWithinRadius[j].Distance
	})

	return driversWithinRadius, nil
}
