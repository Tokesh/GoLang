package repositories

import (
	"context"
	"projectGoLang/source/domain/entity"
)

func (r *Repository) AddBrandRepo(brand string) error {
	q := `
		INSERT INTO brands(brand_name) 
		VALUES($1)
	`
	_, err := r.client.Query(context.TODO(), q, brand)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteBrandByName(brand string) error {
	q := `
		delete from brands where brand_name = $1
	`
	_, err := r.client.Query(context.TODO(), q, brand)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectBrandRepo(brandId int) (entity.Brand, error) {
	q := `
		SELECT brand_id, brand_name FROM public.brands WHERE brand_id = $1
	`
	var brand entity.Brand
	err := r.client.QueryRow(context.TODO(), q, brandId).Scan(&brand.BrandID, &brand.BrandName)
	if err != nil {
		return entity.Brand{}, err
	}

	return brand, nil
}

func (r *Repository) UpdateBrandRepo(brand entity.Brand) error {
	q := `
		update brands set brand_name = $1 where brand_id=$2
	`
	_, err := r.client.Query(context.TODO(), q, brand.BrandName, brand.BrandID)
	if err != nil {
		return err
	}
	return nil
}
