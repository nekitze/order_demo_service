package nats_streaming

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"orders_service/internal/models"
	"orders_service/internal/service"
)

type Handler struct {
	orderService *service.OrderService
}

func NewHandler(s *service.OrderService) *Handler {
	return &Handler{orderService: s}
}

func (h *Handler) IncomeOrder(msg *nats.Msg) {
	order := &models.Order{}
	if err := json.Unmarshal(msg.Data, order); err != nil {
		log.Printf("wrong json: %v\n", order)
		return
	}

	err := h.orderService.Save(order)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("order received")
}
