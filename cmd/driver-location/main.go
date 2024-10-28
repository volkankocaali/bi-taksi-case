package main

import (
	"fmt"
	"github.com/volkankocaali/bi-taksi-case/bootstrap"
	"github.com/volkankocaali/bi-taksi-case/config"
	"github.com/volkankocaali/bi-taksi-case/internal/router"
)

const configName = "./config/driver-location.yaml"

func main() {
	if err := config.LoadConfig(configName); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	cfg := config.GetConfig()

	// initialize the router
	routes := bootstrap.RouterInit()

	// logger initialization
	bootstrap.LoggerInit(cfg.Server.Port)

	// mongodb initialization
	mongo := bootstrap.MongoInit(cfg)

	// register routes
	router.RegisterDriverLocationRoutes(routes, *cfg, mongo)

	// start the server
	bootstrap.StartServer(*cfg, routes)
}
