package service

import (
	"context"
	"fmt"
	"projectGoLang/entity"
	"projectGoLang/postgresql"
)

type Repository struct {
	client postgresql.Client
	//logger *logging.Logger
}

func New(client postgresql.Client) Repository {
	return Repository{
		client: client,
		//logger add TODO
	}
}

func (r *Repository) FindOneProduct(id string) entity.Product {
	q := `
		SELECT product_id,product_type_id, upc_code, title FROM public.products WHERE product_id = $1
	`

	var prod entity.Product
	err := r.client.QueryRow(context.TODO(), q, id).Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
	if err != nil {
		return entity.Product{}
	}

	return prod
}

func (r *Repository) FindAllProducts() []entity.Product {
	q := `
		SELECT product_id,product_type_id, upc_code, title FROM public.products
	`
	products := make([]entity.Product, 0)
	rows, err := r.client.Query(context.TODO(), q)
	if err != nil {
		return products
	}
	for rows.Next() {
		var prod entity.Product
		err = rows.Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
		if err != nil {
			return products
		}
		products = append(products, prod)
	}
	return products
}

func (r *Repository) CreateProduct(product *entity.Product) entity.Product {
	q := `
		INSERT INTO products(product_type_id, upc_code, title) 
		VALUES($1, $2, $3)
	`
	r.client.Query(context.TODO(), q, product.ProductTypeId, product.UpcCode, product.Title)
	//err := r.client.QueryRow(context.TODO(), q, product.ProductTypeId, product.UpcCode, product.Title).Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)

	qSearchingAddedProduct := `
		SELECT product_id, product_type_id, upc_code, title from public.products where upc_code=$1;
	`
	err := r.client.QueryRow(context.TODO(), qSearchingAddedProduct, product.UpcCode).Scan(&product.ProductId, &product.ProductTypeId, &product.UpcCode, &product.Title)
	if err != nil {
		fmt.Errorf("Some problems with creating ", err)
	}
	fmt.Println(product)
	return *product
}

func (r *Repository) DeleteProductByID(id string) entity.Product {
	qSearchingBeforeDeletingProduct := `
		SELECT product_id, product_type_id, upc_code, title from public.products where product_id=$1;
	`
	prod := entity.Product{}
	err := r.client.QueryRow(context.TODO(), qSearchingBeforeDeletingProduct, id).Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
	if err != nil {
		fmt.Errorf("No existing element", err)
		return prod
	}

	q := `
		delete from products where product_id = $1
	`
	err = r.client.QueryRow(context.TODO(), q, id).Scan()
	if err != nil {
		fmt.Errorf("Cannot delete this element sorry!", err)
	}
	return prod
}

//type ProductService struct {
//	Products []entity.Product
//}
//
//func New() ProductService {
//	return ProductService{}
//}
//
//func (service *ProductService) AddProduct(product entity.Product) entity.Product {
//	service.Products = append(service.Products, product)
//	return product
//}
//func (service *ProductService) FindAllProducts() []entity.Product {
//	return service.Products
//}
