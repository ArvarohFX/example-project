package usecase

import (
	"github.com/ArvarohFX/example-project/config"
	"github.com/ArvarohFX/example-project/internal/repository"
	"go.uber.org/zap"
)

type Usecase interface {
	User() User
	Order() Order
}

type usecase struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	store  repository.Store

	userUsecase  User
	orderUsecase Order
}

func New(cfg *config.Config, logger *zap.SugaredLogger, store repository.Store) Usecase {
	return &usecase{
		cfg:    cfg,
		logger: logger,
		store:  store,
	}
}

func (s *usecase) User() User {
	if s.userUsecase != nil {
		return s.userUsecase
	}
	return &userUsecase{
		store: s.store,
	}
}

func (s *usecase) Order() Order {
	if s.orderUsecase != nil {
		return s.orderUsecase
	}
	return &orderService{}
}
