package services

import (
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/repositories"
)

type Service struct {
	repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Repository: *repo,
	}
}

func (s Service) SelectCart(user_id int) ([]entity.Cart, error) {
	return s.SelectCartByUserID(user_id)
}

func (s Service) DeleteCart(user_id int) err {
	return s.DeleteFromCartByUserID(user_id)
}
