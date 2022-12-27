package repositories

import (
	"context"
	"fmt"
	"projectGoLang/source/domain/entity"
)

// order_id, client_id, delivery_number_id, payment_id, store_id, datetime, cashair
func (r *Repository) AddOrderRepo(order entity.Order) error {
	q := `
		INSERT INTO orders(client_id, delivery_number_id, payment_id, store_id, date_time, cash_air) 
		VALUES($1, $2, $3, $4, $5, $6)
	`
	_, err := r.client.Query(context.TODO(), q, order.ClientId, order.DeliveryNumberId, order.PaymentId, order.StoreId,
		order.DateTime, order.CashAir)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteOrderRepo(orderId int) error {
	q := `
		delete from orders where order_id = $1
	`
	_, err := r.client.Query(context.TODO(), q, orderId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectOrderRepo(orderId int) (entity.Order, error) {
	q := `
		SELECT order_id, client_id, delivery_number_id, payment_id, store_id, date_time, cash_air
		FROM public.orders WHERE order_id = $1
	`
	var order entity.Order
	err := r.client.QueryRow(context.TODO(), q, orderId).Scan(&order.OrderId, &order.ClientId, &order.DeliveryNumberId,
		&order.PaymentId, &order.StoreId, &order.DateTime, &order.CashAir)
	if err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (r *Repository) UpdateOrderRepo(order entity.Order) error {
	q := `
		update orders set client_id = $1, delivery_number_id = $2, payment_id = $3, store_id = $4, 
		        date_time = $5, cash_air = $6 where order_id=$7
	`
	_, err := r.client.Query(context.TODO(), q, order.ClientId, order.DeliveryNumberId, order.PaymentId, order.StoreId,
		order.DateTime, order.CashAir, order.OrderId)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) FillOrder(cart []entity.Cart, orderId int) error {
	q := `
		INSERT INTO order_items(order_id, product_id, count) 
		VALUES 
	`
	orderIdStr := string(orderId)
	for i := 0; i < len(cart); i++ {
		if i != 0 {
			q += ", "
		}
		q += "(" + orderIdStr + "," + string(cart[i].ProductID) + "," + string(cart[i].Quantity) + ")"
	}
	fmt.Print("Everything is fine")
	return nil
}
