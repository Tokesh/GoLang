package entity

type Payment struct {
	PaymentID   int     `json:"payment_id"`
	Summ        float64 `json:"summ"`
	TotalPrice  float64 `json:"total_price"`
	PaymentType string  `json:"payment_type"`
}

//payment_id, summ, total_price, payment_type
