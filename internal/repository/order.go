package repository

import "gorm.io/gorm"

type OrderRepository interface {
	Create() (string, int, error)
}

type orderRepository struct {
	DB *gorm.DB
}

func (r *orderRepository) Create() (string, int, error) {
	return "param", 10, nil
}
