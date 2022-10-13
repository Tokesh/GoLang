package repositories

import (
	"context"
	"projectGoLang/source/domain/entity"
)

func (r *Repository) AddToCartRepo(cart *entity.Cart) error {
	q := `
		INSERT INTO shopping_cart(user_id, product_id, quantity) 
		VALUES($1, $2, $3)
	`
	_, err := r.client.Query(context.TODO(), q, cart.UserID, cart.ProductID, cart.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromCartByUserID(userId int) error {
	q := `
		delete from shopping_cart where user_id = $1
	`
	_, err := r.client.Query(context.TODO(), q, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SelectCartByUserID(userId int) ([]entity.Cart, error) {
	q := `
		SELECT user_id, product_id, quantity FROM public.shopping_cart WHERE user_id = $1
	`
	carts := make([]entity.Cart, 0)
	rows, err := r.client.Query(context.TODO(), q, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cart entity.Cart
		err = rows.Scan(&cart.UserID, &cart.ProductID, &cart.Quantity)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func (r *Repository) UpdateCartRepo(cart *entity.Cart) error {
	q := `
		update shopping_cart set quantity = $1 where user_id=$2 and product_id = $3
	`
	_, err := r.client.Query(context.TODO(), q, &cart.Quantity, &cart.UserID, &cart.ProductID)
	if err != nil {
		return err
	}
	return nil
}

//func (r *Repository) DeleteProductByID(id string) entity.Product {
//	qSearchingBeforeDeletingProduct := `
//		SELECT product_id, product_type_id, upc_code, title from public.products where product_id=$1;
//	`
//	prod := entity.Product{}
//	err := r.client.QueryRow(context.TODO(), qSearchingBeforeDeletingProduct, id).Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
//	if err != nil {
//		fmt.Errorf("No existing element", err)
//		return prod
//	}
//
//	q := `
//		delete from products where product_id = $1
//	`
//	err = r.client.QueryRow(context.TODO(), q, id).Scan()
//	if err != nil {
//		fmt.Errorf("Cannot delete this element sorry!", err)
//	}
//	return prod
//}
