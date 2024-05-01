package server

import (
	"context"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer       *http.Server
	natsSubscription *nats.Subscription
}

func (s Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down...")
	return s.httpServer.Shutdown(ctx)
}
