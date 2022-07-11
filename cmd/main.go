package main

import (
	cine_base_app "cine-base-app"
	"cine-base-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(cine_base_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
