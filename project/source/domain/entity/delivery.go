package entity

import "github.com/jackc/pgtype"

type Delivery struct {
	StatusId int         `json:"delivery_status_id"`
	OrderId  int         `json:"order_delivery_id"`
	DateTime pgtype.Date `json:"date_time"`
	Status   string      `json:"status"`
}

//delivery_status_id
//order_delivery_id
//date_time
//status
