package repositories

import (
	"context"
	"fmt"
	"projectGoLang/source/domain/entity"
)

func (r *Repository) AddToPaymentRepo(payment *entity.Payment) error {
	q := `
		INSERT INTO payment(payment_id, summ, total_price, payment_type) 
		VALUES($1, $2, $3, $4)
	`
	_, err := r.client.Query(context.TODO(), q, payment.PaymentID, payment.Summ, payment.TotalPrice, payment.PaymentType)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) DeletePaymentRepo(userId int) error {
	q := `
		delete from payment where payment_id = $1
	`
	_, err := r.client.Query(context.TODO(), q, userId)
	if err != nil {
		return err
	}
	return nil
}

// payment_id, summ, total_price, payment_type
func (r *Repository) SelectPaymentByPaymentIdRepo(payment_id int) (entity.Payment, error) {
	q := `
		SELECT payment_id,summ, total_price, payment_type FROM payment WHERE payment_id = $1
	`
	var payment entity.Payment
	fmt.Println("AMP")
	err := r.client.QueryRow(context.TODO(), q, payment_id).Scan(&payment.PaymentID, &payment.Summ, &payment.TotalPrice, &payment.PaymentType)
	if err != nil {
		return entity.Payment{}, err
	}
	fmt.Println(payment)
	return payment, nil
}

func (r *Repository) SelectPaymentByUserIdRepo(user_id int) ([]entity.Payment, error) {
	q := `
		select payment.payment_id, summ, total_price, payment_type
		from payment left join orders o on payment.payment_id = o.payment_id 
		where client_id = $1
	`
	payments := make([]entity.Payment, 0)
	rows, err := r.client.Query(context.TODO(), q, user_id)
	if err != nil {
		return make([]entity.Payment, 0), nil
	}
	for rows.Next() {
		var paym entity.Payment
		err = rows.Scan(&paym.PaymentID, &paym.Summ, &paym.TotalPrice, &paym.PaymentType)
		if err != nil {
			return make([]entity.Payment, 0), nil
		}
		payments = append(payments, paym)
	}
	return payments, nil
}

func (r *Repository) UpdatePaymentRepo(payment *entity.Payment) error {
	q := `
		update payment set  summ = $1, total_price = $2,payment_type = $3  where payment_id=$4
	`
	_, err := r.client.Query(context.TODO(), q, &payment.Summ, &payment.TotalPrice, &payment.PaymentType, &payment.PaymentID)
	if err != nil {
		return err
	}
	return nil
}

//payment_id, summ, total_price, payment_type
