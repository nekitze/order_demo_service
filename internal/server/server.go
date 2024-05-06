package server

import (
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"orders_service/internal/handlers/nats_streaming"
	"os"
)

type Server struct {
	httpServer       *http.Server
	natsSubscription *nats.Subscription
}

func NewServer(handler http.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: handler,
	}}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) SubscribeNatsStream(streamHandler *nats_streaming.Handler) {
	nc, err := nats.Connect(os.Getenv("NATS_STREAMING_URL"))
	if err != nil {
		log.Println(err)
		return
	}

	s.natsSubscription, err = nc.Subscribe(os.Getenv("NATS_STREAMING_SUBJECT"), streamHandler.IncomeOrder)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("subscribed to nats stream")
}
