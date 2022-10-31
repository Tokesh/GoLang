package repositories

import (
	"context"
	"projectGoLang/source/domain/entity"
)

func (r *Repository) AddToPaymentRepo(payment *entity.Payment) error {
	q := `
		INSERT INTO shopping_cart(payment_id, summ, total_price, payment_type) 
		VALUES($1, $2, $3, $4)
	`
	_, err := r.client.Query(context.TODO(), q, payment.PaymentID, payment.Summ, payment.TotalPrice, payment.PaymentType)
	if err != nil {
		return err
	}
	return nil
}

//payment_id, summ, total_price, payment_type
