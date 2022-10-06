package service

import (
	"context"
	"fmt"
	"projectGoLang/entity"
)

func (r *Repository) AddToCart(cart *entity.Cart) *entity.Cart {
	q := `
		INSERT INTO shopping_cart(user_id, product_id, quantity) 
		VALUES($1, $2, $3)
	`
	err := r.client.QueryRow(context.TODO(), q, cart.UserID, cart.ProductID, cart.Quantity)
	if err != nil {
		fmt.Errorf("Something go wrong", err)
	}
	return cart
}

func (r *Repository) DeleteFromCartByUserID(cart *entity.Cart) *entity.Cart {
	q := `
		delete from shopping_cart where user_id = $1
	`
	err := r.client.QueryRow(context.TODO(), q, cart.UserID)
	if err != nil {
		return cart
	}
	return cart
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
