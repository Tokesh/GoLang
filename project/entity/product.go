package entity

type Product struct {
	ProductId     int    `json:"ProductId"`
	ProductTypeId int    `json:"ProductTypeId"`
	UpcCode       string `json:"UpcCode"`
	Title         string `json:"Title"`
}
