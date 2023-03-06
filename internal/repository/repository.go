package repository

import (
	"github.com/ArvarohFX/example-project/config"
	db "github.com/ArvarohFX/example-project/internal"
	"go.uber.org/zap"
)

type Store interface {
	User() UserRepository
	Order() OrderRepository
}

type store struct {
	dbStorage *db.DBStorage
	cfg       *config.Config
	logger    *zap.SugaredLogger

	userRepository  UserRepository
	orderRepository OrderRepository
}

func New(cfg *config.Config, logger *zap.SugaredLogger, dbStorage *db.DBStorage) Store {
	return &store{
		dbStorage: dbStorage,
		cfg:       cfg,
		logger:    logger,
	}
}

func (s *store) User() UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	return &userRepository{
		DB: s.dbStorage.DB,
	}
}

func (s *store) Order() OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}
	return &orderRepository{
		DB: s.dbStorage.DB,
	}
}
