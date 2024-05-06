package main

import (
	"log"
	"orders_service/internal/database"
	"orders_service/internal/handlers/http"
	"orders_service/internal/handlers/nats_streaming"
	"orders_service/internal/repository/postgres"
	"orders_service/internal/server"
	"orders_service/internal/service"
)

func main() {
	db := database.NewPostgresDatabase()
	ordersRepository := postgres.NewOrderRepository(db)
	orderService := service.NewOrderService(ordersRepository)

	httpHandler := http.NewHandler(orderService)
	streamHandler := nats_streaming.NewHandler(orderService)

	srv := server.NewServer(httpHandler.InitRoutes())
	srv.SubscribeNatsStream(streamHandler)

	if err := srv.Run(); err != nil {
		log.Fatalf("error while runnning server: %s", err.Error())
	}
}
