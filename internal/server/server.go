package server

import (
	"context"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"orders_service/internal/handler/nats_streaming"
)

type Server struct {
	httpServer       *http.Server
	natsSubscription *nats.Subscription
}

func NewServer(httpServer *http.Server) *Server {
	return &Server{httpServer: httpServer}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) SubscribeNatsStream(natsUrl, subject string, streamHandler *nats_streaming.Handler) {
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Println(err)
		return
	}

	s.natsSubscription, err = nc.Subscribe(subject, streamHandler.IncomeOrder)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("subscribed to nats stream")
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("shutting down...")
	//s.natsSubscription.Drain()
	return s.httpServer.Shutdown(ctx)
}
