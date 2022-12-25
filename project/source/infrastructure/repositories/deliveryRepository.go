package repositories

import (
	"context"
	"errors"
	"projectGoLang/source/domain/entity"
)

// delivery_status_id
// order_delivery_id
// date_time
// status
func (r *Repository) AddStatusToDeliveryRepo(delivery *entity.Delivery) error {
	//fmt.Println(delivery.StatusId, delivery.Status, delivery.OrderId, delivery.DateTime)
	if delivery.Status != "Unloaded" && delivery.Status != "Delivered" && delivery.Status != "Left the warehouse" {
		return errors.New("not correct status, check again")
	}
	q := `
		INSERT INTO delivery_status(order_delivery_id, date_time, status) 
		VALUES($1, $2, $3)
	`
	_, err := r.client.Query(context.TODO(), q, delivery.OrderId, delivery.DateTime, delivery.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteStatusFromDeliveryRepo(delivery *entity.Delivery) error {
	if delivery.Status != "Unloaded" && delivery.Status != "Delivered" && delivery.Status != "Left the warehouse" {
		return errors.New("not correct status, check again")
	}
	q := `
		delete from delivery_status where order_delivery_id = $1 and status = $2
	`
	_, err := r.client.Query(context.TODO(), q, delivery.OrderId, delivery.Status)
	if err != nil {
		return err
	}
	return nil
}

// delivery_status_id
// order_delivery_id
// date_time
// status

func (r *Repository) SelectDeliveryStatusRepo(order_id int) ([]entity.Delivery, error) {
	q := `
		SELECT delivery_status_id, order_delivery_id, date_time, status FROM public.delivery_status WHERE order_delivery_id = $1
	`
	deliveries := make([]entity.Delivery, 0)
	rows, err := r.client.Query(context.TODO(), q, order_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var delivery entity.Delivery
		err = rows.Scan(&delivery.StatusId, &delivery.OrderId, &delivery.DateTime, &delivery.Status)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, delivery)
	}
	return deliveries, nil
}
func (r *Repository) SelectLastDeliveryStatusRepo(order_delivery_id int) (entity.Delivery, error) {

	q := `
		select delivery_status_id, order_delivery_id, date_time, status
			from delivery_status 
		    where order_delivery_id = $1 
		    order by delivery_status_id desc 
			limit 1
	`
	var delivery entity.Delivery
	err := r.client.QueryRow(context.TODO(), q, order_delivery_id).Scan(&delivery.StatusId, &delivery.OrderId, &delivery.DateTime, &delivery.Status)
	if err != nil {
		return entity.Delivery{}, err
	}

	return delivery, nil
}

func (r *Repository) UpdateDeliveryStatusRepo(delivery *entity.Delivery) error {
	if delivery.Status != "Unloaded" && delivery.Status != "Delivered" && delivery.Status != "Left the warehouse" {
		return errors.New("not correct status, check again")
	}
	q := `
		update delivery_status set status = $1 where delivery_status_id=$2
	`
	_, err := r.client.Query(context.TODO(), q, &delivery.Status, &delivery.StatusId)
	if err != nil {
		return err
	}
	return nil
}
