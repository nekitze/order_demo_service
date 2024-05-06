package postgres

import (
	"errors"
	"orders_service/internal/database"
	"orders_service/internal/models"
)

type OrderRepositoryPostgres struct {
	db *database.PostgresDatabase
}

func NewOrderRepository(db *database.PostgresDatabase) *OrderRepositoryPostgres {
	return &OrderRepositoryPostgres{db: db}
}

func (r OrderRepositoryPostgres) Save(order *models.Order) error {
	_, err := r.db.DB.Model(order).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (r OrderRepositoryPostgres) FindById(id string) (*models.Order, error) {
	order := &models.Order{OrderUid: id}

	err := r.db.DB.Model(order).WherePK().Select()
	if err != nil {
		return nil, errors.New("not found")
	}

	return order, nil
}

func (r OrderRepositoryPostgres) FindAll() ([]*models.Order, error) {
	var orders []*models.Order
	err := r.db.DB.Model(&orders).Select()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
