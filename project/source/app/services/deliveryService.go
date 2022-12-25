package services

import "projectGoLang/source/domain/entity"

// AddStatus, DeleteStatus, SelectDelivery, SelectLastDelivery, Update Delivery
func (s Service) AddDeliveryStatus(delivery entity.Delivery) error {
	return s.AddStatusToDeliveryRepo(&delivery)
}

func (s Service) DeleteDeliveryStatus(delivery entity.Delivery) error {
	return s.DeleteStatusFromDeliveryRepo(&delivery)
}

func (s Service) SelectDelivery(order_id int) ([]entity.Delivery, error) {
	return s.SelectDeliveryStatusRepo(order_id)
}
func (s Service) SelectLastDelivery(order_id int) (entity.Delivery, error) {
	return s.SelectLastDeliveryStatusRepo(order_id)
}
func (s Service) UpdateDelivery(delivery entity.Delivery) error {
	return s.UpdateDeliveryStatusRepo(&delivery)
}
