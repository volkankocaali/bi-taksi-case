package router

import (
	"github.com/gorilla/mux"
	"github.com/volkankocaali/bi-taksi-case/config"
	"github.com/volkankocaali/bi-taksi-case/internal/auth"
	"github.com/volkankocaali/bi-taksi-case/internal/handler"
	"github.com/volkankocaali/bi-taksi-case/internal/repository"
	"github.com/volkankocaali/bi-taksi-case/internal/usecases"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterDriverLocationRoutes(r *mux.Router, cfg config.Config, mongo *mongo.Client) {
	db := mongo.Database("bitaksi")

	driverLocationCollection := db.Collection("driver-locations")
	userCollection := db.Collection("users")

	// user login and register
	userRepo := repository.NewMongoUserRepository(userCollection)
	userUsecase := usecases.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// create the repository, usecase and interfaces
	driverLocationRepo := repository.NewMongoDriverLocationRepository(driverLocationCollection)
	driverLocationUsecase := usecases.NewDriverLocationUseCase(driverLocationRepo)
	driverLocationHandler := handler.NewDriverHandler(driverLocationUsecase)

	r.HandleFunc("/api/v1/login", userHandler.UserLogin).Methods("POST")
	r.HandleFunc("/api/v1/register", userHandler.UserRegister).Methods("POST")

	r.HandleFunc("/api/v1/driver-locations/{driver_id}", auth.JWTAuthMiddleware(cfg.JwtSecretKey, cfg.JwtIssuer, driverLocationHandler.GetLatestDriverLocation)).Methods("GET")
	r.HandleFunc("/api/v1/driver-locations", auth.JWTAuthMiddleware(cfg.JwtSecretKey, cfg.JwtIssuer, driverLocationHandler.DriverCreateOrUpdate)).Methods("POST", "PUT")
	r.HandleFunc("/api/v1/find-driver-within-radius", auth.JWTAuthMiddleware(cfg.JwtSecretKey, cfg.JwtIssuer, driverLocationHandler.FindDriverWithinRadius)).Methods("POST")
}

func RegisterMatchingApiRoutes(r *mux.Router, cfg config.Config, mongo *mongo.Client) {
	db := mongo.Database("bitaksi")

	driverCollection := db.Collection("driver-locations")
	userCollection := db.Collection("users")

	// create the repository, usecase and interfaces

	// user login and register
	userRepo := repository.NewMongoUserRepository(userCollection)
	userUsecase := usecases.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	driverLocationRepo := repository.NewMongoDriverLocationRepository(driverCollection)
	matchingUseCase := usecases.NewMatchingUseCase(driverLocationRepo)
	matchingHandler := handler.NewMatchingHandler(matchingUseCase)

	r.HandleFunc("/api/v1/login", userHandler.UserLogin).Methods("POST")
	r.HandleFunc("/api/v1/register", userHandler.UserRegister).Methods("POST")

	r.HandleFunc("/api/v1/match-driver", auth.JWTAuthMiddleware(cfg.JwtSecretKey, cfg.JwtIssuer, matchingHandler.MatchDriver)).Methods("POST")
}
