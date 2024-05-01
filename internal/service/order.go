package service

import (
	"errors"
	"log"
	"orders_service/internal/database"
	"orders_service/internal/models"
)

type OrderService struct {
	db    *database.Database
	cache map[string]*models.Order
}

func NewOrderService(db *database.Database) *OrderService {
	os := &OrderService{db: db, cache: map[string]*models.Order{}}
	os.loadCache()
	return os
}

func (s OrderService) loadCache() {
	orders := make([]models.Order, 0)

	err := s.db.DB.Model(&orders).Select()
	if err != nil {
		log.Println(err)
	}

	for _, order := range orders {
		s.cache[order.OrderUid] = &order
	}

	log.Println("Cache loading finished")
}

func (s OrderService) Save(order *models.Order) error {
	_, err := s.db.DB.Model(order).Insert()
	if err != nil {
		return err
	}

	s.cache[order.OrderUid] = order
	return nil
}

func (s OrderService) FindById(id string) (*models.Order, error) {
	order, ok := s.cache[id]

	if !ok {
		return nil, errors.New("not found")
	}

	return order, nil
}
