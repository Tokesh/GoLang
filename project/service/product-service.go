package service

import (
	"projectGoLang/entity"
)

type ProductService interface {
	AddProduct(entity.Product) entity.Product
	FindAllProducts() []entity.Product
	//SearchProduct() entity.Product
}

type productService struct {
	products []entity.Product
}

func New() *productService {
	return &productService{}
}

func (service *productService) AddProduct(product entity.Product) entity.Product {
	service.products = append(service.products, product)
	return product
}
func (service *productService) FindAllProducts() []entity.Product {
	return service.products
}
