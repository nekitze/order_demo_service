package repository

import "orders_service/internal/models"

type OrderRepository interface {
	FindAll() ([]*models.Order, error)
	FindById(i string) (*models.Order, error)
	Save(order *models.Order) error
}
