package main

import (
	"github.com/valverdethiago/volume-take-home-assignment/api"
	"github.com/valverdethiago/volume-take-home-assignment/config"
	"github.com/valverdethiago/volume-take-home-assignment/domain"
	"github.com/valverdethiago/volume-take-home-assignment/server"
	"log"
)

func main() {
	cfg, _ := config.LoadConfig("./env", "app")
	service := domain.NewFlightPathServiceImpl()
	controller := api.NewFlightMapApiHandlerImpl(service)
	configureServer(cfg, controller)
}

func configureServer(cfg config.AppConfig, controller server.Controller) *server.Server {
	srv := server.NewServer(cfg)
	srv.ConfigureController(controller)
	srv.ConfigureLogging()
	err := srv.Start()
	if err != nil {
		log.Fatal("Error running the server: ", err.Error())
	}
	return srv
}
