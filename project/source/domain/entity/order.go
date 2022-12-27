package entity

import (
	"database/sql"
	"github.com/jackc/pgtype"
)

// order_id, client_id, delivery_number_id, payment_id, store_id, datetime, cashair

type Order struct {
	OrderId          int           `json:"order_id"`
	ClientId         int           `json:"client_id"`
	DeliveryNumberId sql.NullInt32 `json:"delivery_number_id"`
	PaymentId        int           `json:"payment_id"`
	StoreId          int           `json:"store_id"`
	DateTime         pgtype.Date   `json:"date_time"`
	CashAir          sql.NullInt32 `json:"cash_air"`
}
