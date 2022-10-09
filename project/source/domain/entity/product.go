package entity

type Product struct {
	ProductId     int    `json:"product_id"`
	ProductTypeId int    `json:"product_type_id"`
	UpcCode       string `json:"upc_code"`
	Title         string `json:"title"`
}
