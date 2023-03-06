package usecase

import "github.com/ArvarohFX/example-project/internal/repository"

type User interface {
	Create() error
}

type userUsecase struct {
	store repository.Store
}

func (s *userUsecase) Create() error {
	err := s.store.User().Get()
	if err != nil {
		return err
	}
	return nil
}
