package services

import "projectGoLang/source/domain/entity"

func (s Service) AddToPayment(payment entity.Payment) error {
	return s.AddToPaymentRepo(&payment)
}
