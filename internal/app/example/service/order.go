package service

import "github.com/ArvarohFX/example-project/internal/app/example/usecase"

type Order interface {
	Get()
}

type orderService struct {
	usecase usecase.Usecase
}

func (s *orderService) Get() {

}
