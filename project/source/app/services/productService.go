package services

import "projectGoLang/source/domain/entity"

func (s Service) AddProduct(prod entity.Product) error {
	return s.CreateProduct(&prod)
}

func (s Service) SelectProduct(user_id int) (entity.Product, error) {
	return s.FindOneProduct(user_id)
}

func (s Service) SelectAllProducts() ([]entity.Product, error) {
	return s.FindAllProducts()
}

func (s Service) DeleteProduct(user_id int) error {
	return s.DeleteProductByID(user_id)
}

func (s Service) UpdateProduct(product entity.Product) error {
	return s.UpdateProductByID(product)
}
