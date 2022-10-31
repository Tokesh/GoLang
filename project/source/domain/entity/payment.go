package entity

type Payment struct {
	PaymentID   int `json:"payment_id"`
	Summ        int `json:"summ"`
	TotalPrice  int `json:"total_price"`
	PaymentType int `json:"payment_type"`
}

//payment_id, summ, total_price, payment_type
