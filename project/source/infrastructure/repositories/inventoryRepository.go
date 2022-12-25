package repositories

import (
	"context"
	"projectGoLang/source/domain/entity"
)

// store_id, product_id, brand_id, price, quantity

func (r *Repository) AddStoreInventoryRepo(inventory entity.StoreInventory) error {
	q := `
		INSERT INTO store_inventory(store_id, product_id, brand_id, price, quantity) 
		VALUES($1, $2, $3, $4, $5)
	`
	_, err := r.client.Query(context.TODO(), q, inventory.StoreId, inventory.ProductId, inventory.BrandId, inventory.Price,
		inventory.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteStoryInventoryRepo(store_id int, product_id int) error {
	q := `
		delete from store_inventory where store_id = $1 and product_id = $2
	`
	_, err := r.client.Query(context.TODO(), q, store_id, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectStoreInventoryRepo(store_id int, product_id int) (entity.StoreInventory, error) {
	q := `
		SELECT store_id, product_id, brand_id, price, quantity 
		FROM public.store_inventory
		WHERE store_id = $1 and product_id = $2
	`
	var ProdInventory entity.StoreInventory
	err := r.client.QueryRow(context.TODO(), q, store_id, product_id).Scan(&ProdInventory.StoreId, &ProdInventory.ProductId,
		&ProdInventory.BrandId, &ProdInventory.Price, &ProdInventory.Quantity)
	if err != nil {
		return entity.StoreInventory{}, err
	}

	return ProdInventory, nil
}

func (r *Repository) UpdateStoreInventoryRepo(inventory entity.StoreInventory) error {
	q := `
		update store_inventory set brand_id = $1, price = $2, quantity = $3 where store_id=$4 and product_id = $5
	`
	_, err := r.client.Query(context.TODO(), q, inventory.BrandId, inventory.Price, inventory.Quantity, inventory.StoreId,
		inventory.ProductId)
	if err != nil {
		return err
	}
	return nil
}
