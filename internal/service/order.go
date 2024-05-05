package service

import (
	"errors"
	"log"
	"orders_service/internal/models"
	"orders_service/internal/repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
	cache           map[string]*models.Order
}

func NewOrderService(r repository.OrderRepository) *OrderService {
	os := &OrderService{orderRepository: r, cache: map[string]*models.Order{}}
	os.loadCache()
	return os
}

func (s OrderService) loadCache() {
	orders, err := s.orderRepository.FindAll()
	if err != nil {
		log.Println("failed to load cache", err)
		return
	}

	for _, order := range orders {
		s.cache[order.OrderUid] = order
	}

	log.Println("cache loading finished")
}

func (s OrderService) Save(order *models.Order) error {
	err := s.orderRepository.Save(order)
	if err != nil {
		return err
	}

	s.cache[order.OrderUid] = order
	return nil
}

func (s OrderService) FindById(id string) (*models.Order, error) {
	order, ok := s.cache[id]

	if !ok {
		return nil, errors.New("order not found")
	}

	return order, nil
}
