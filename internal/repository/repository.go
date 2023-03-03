package repository

import (
	"github.com/ArvarohFX/example-project/config"
	"go.uber.org/zap"
)

type Store interface {
	User() UserRepository
	Order() OrderRepository
}

type store struct {
	//dbStorage *internal.DBStorage
	cfg    *config.Config
	logger *zap.SugaredLogger

	userRepository  UserRepository
	orderRepository OrderRepository
}

func (s *store) User() UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	return &userRepository{}
}

func (s *store) Order() OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}
	return &orderRepository{}
}
