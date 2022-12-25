package services

import "projectGoLang/source/domain/entity"

func (s Service) AddToBrands(brandName string) error {
	return s.AddBrandRepo(brandName)
}

func (s Service) SelectBrand(brandId int) (entity.Brand, error) {
	return s.SelectBrandRepo(brandId)
}

func (s Service) DeleteBrand(brandName string) error {
	return s.DeleteBrandByName(brandName)
}
func (s Service) UpdateBrand(brand entity.Brand) error {
	return s.UpdateBrandRepo(brand)
}
