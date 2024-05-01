package main

import (
	"log"
	"orders_service/internal/database"
	"orders_service/internal/handler"
	"orders_service/internal/server"
	"orders_service/internal/service"
)

func main() {
	db := database.NewDatabase()
	handlers := handler.NewHandler(service.NewOrderService(db))

	srv := server.Server{}
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while runnning server: %s", err.Error())
	}
}
