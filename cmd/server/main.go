package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"orders_service/config"
	"orders_service/internal/database"
	"orders_service/internal/handler/http"
	"orders_service/internal/handler/nats_streaming"
	"orders_service/internal/repository/postgres"
	"orders_service/internal/server"
	"orders_service/internal/service"
)

func main() {

	pgOptions, err := config.LoadDatabaseConfig("config/database/config.json")
	if err != nil {
		panic(err)
	}

	db := database.NewPostgresDatabase(pgOptions)
	ordersRepository := postgres.NewOrderRepository(db)
	orderService := service.NewOrderService(ordersRepository)

	httpHandler := http.NewHandler(orderService)
	streamHandler := nats_streaming.NewHandler(orderService)

	serverOptions, err := config.LoadServerConfig("config/server/config.json")
	if err != nil {
		panic(err)
	}

	serverOptions.Handler = httpHandler.InitRoutes()
	srv := server.NewServer(serverOptions)
	srv.SubscribeNatsStream(nats.DefaultURL, "orders.*", streamHandler)
	if err := srv.Run(); err != nil {
		log.Fatalf("error while runnning server: %s", err.Error())
	}
}
