package service

import (
	"github.com/ArvarohFX/example-project/config"
	"github.com/ArvarohFX/example-project/internal/app/example/usecase"
	"go.uber.org/zap"
)

type Service interface {
	User() User
	Order() Order
}

type service struct {
	cfg     *config.Config
	logger  *zap.SugaredLogger
	usecase usecase.Usecase

	userService  User
	orderService Order
}

func New(cfg *config.Config, logger *zap.SugaredLogger, usecase usecase.Usecase) Service {
	return &service{
		cfg:     cfg,
		logger:  logger,
		usecase: usecase,
	}
}

func (s *service) User() User {
	if s.userService != nil {
		return s.userService
	}
	return &userService{
		usecase: s.usecase,
	}
}

func (s *service) Order() Order {
	if s.orderService != nil {
		return s.orderService
	}
	return &orderService{
		usecase: s.usecase,
	}
}
