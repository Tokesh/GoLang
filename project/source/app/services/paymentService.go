package services

import (
	"projectGoLang/source/domain/entity"
)

func (s Service) AddToPayment(payment entity.Payment) error {
	return s.AddToPaymentRepo(&payment)
}

func (s Service) SelectPaymentByPaymentID(payment_id int) (entity.Payment, error) {
	return s.SelectPaymentByPaymentIdRepo(payment_id)
}
func (s Service) SelectPaymentByUserID(user_id int) ([]entity.Payment, error) {
	return s.SelectPaymentByUserIdRepo(user_id)
}
func (s Service) DeletePayment(user_id int) error {
	return s.DeletePaymentRepo(user_id)
}

func (s Service) UpdatePayment(payment entity.Payment) error {
	return s.UpdatePaymentRepo(&payment)
}
