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

func (s Service) AddToCart(cart entity.Cart) error {
	return s.AddToCartRepo(&cart)
}

func (s Service) SelectCart(user_id int) ([]entity.Cart, error) {
	return s.SelectCartByUserID(user_id)
}

func (s Service) DeleteCart(user_id int) error {
	return s.DeleteFromCartByUserID(user_id)
}
func (s Service) UpdateCart(cart entity.Cart) error {
	return s.UpdateCartRepo(&cart)
}
