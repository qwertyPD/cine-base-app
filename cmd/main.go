package main

import (
	cine_base_app "cine-base-app"
	"cine-base-app/pkg/handler"
	"cine-base-app/pkg/repository"
	"cine-base-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cine_base_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
