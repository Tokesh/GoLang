package entity

type StoreInventory struct {
	StoreId   int `json:"store_id"`
	ProductId int `json:"product_id"`
	BrandId   int `json:"brand_id"`
	Price     int `json:"price"`
	Quantity  int `json:"quantity"`
}
