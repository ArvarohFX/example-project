package usecase

import (
	"github.com/ArvarohFX/example-project/config"
	"go.uber.org/zap"
)

type Usecase interface {
	User() User
	Order() Order
}

type usecase struct {
	cfg    *config.Config
	logger *zap.SugaredLogger

	userUsecase  User
	orderUsecase Order
}

func New(cfg *config.Config, logger *zap.SugaredLogger) Usecase {
	return &usecase{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *usecase) User() User {
	if s.userUsecase != nil {
		return s.userUsecase
	}
	return &userUsecase{}
}

func (s *usecase) Order() Order {
	if s.orderUsecase != nil {
		return s.orderUsecase
	}
	return &orderService{}
}
