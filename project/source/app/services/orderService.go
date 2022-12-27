package services

import "projectGoLang/source/domain/entity"

func (s Service) AddToOrder(order entity.Order) error {
	return s.AddOrderRepo(order)
}

func (s Service) FillOrderService(cart []entity.Cart, orderId int) error {
	return s.FillOrder(cart, orderId)
}

func (s Service) SelectOrder(orderId int) (entity.Order, error) {
	return s.SelectOrderRepo(orderId)
}

func (s Service) DeleteOrder(orderId int) error {
	return s.DeleteOrderRepo(orderId)
}
func (s Service) UpdateOrder(order entity.Order) error {
	return s.UpdateOrderRepo(order)
}
