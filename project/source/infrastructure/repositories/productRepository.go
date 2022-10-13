package repositories

import (
	"context"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/postgresql"
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

func (r *Repository) FindOneProduct(id int) (entity.Product, error) {
	q := `
		SELECT product_id,product_type_id, upc_code, title FROM public.products WHERE product_id = $1
	`
	var prod entity.Product
	err := r.client.QueryRow(context.TODO(), q, id).Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
	if err != nil {
		return entity.Product{}, err
	}

	return prod, nil
}

func (r *Repository) FindAllProducts() ([]entity.Product, error) {
	q := `
		SELECT product_id,product_type_id, upc_code, title FROM public.products
	`
	products := make([]entity.Product, 0)
	rows, err := r.client.Query(context.TODO(), q)
	if err != nil {
		return make([]entity.Product, 0), nil
	}
	for rows.Next() {
		var prod entity.Product
		err = rows.Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
		if err != nil {
			return make([]entity.Product, 0), nil
		}
		products = append(products, prod)
	}
	return products, nil
}

func (r *Repository) CreateProduct(product *entity.Product) error {
	q := `
		INSERT INTO products(product_type_id, upc_code, title) 
		VALUES($1, $2, $3)
	`
	_, err := r.client.Query(context.TODO(), q, product.ProductTypeId, product.UpcCode, product.Title)
	//err := r.client.QueryRow(context.TODO(), q, product.ProductTypeId, product.UpcCode, product.Title).Scan(&prod.ProductId, &prod.ProductTypeId, &prod.UpcCode, &prod.Title)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteProductByID(id int) error {
	q := `
		delete from products where product_id = $1
	`
	_, err := r.client.Query(context.TODO(), q, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateProductByID(prod entity.Product) error {
	q := `
		update products set product_type_id = $2, 
		                    upc_code = $3, 
		                    title = $4 
		                where product_id = $1
	`
	_, err := r.client.Query(context.TODO(), q, prod.ProductId, prod.ProductTypeId, prod.UpcCode, prod.Title)
	if err != nil {
		return err
	}
	return nil
}

//type ProductService struct {
//	Products []entity.Product
//}
//
//func New() ProductService {
//	return ProductService{}
//}
//
//func (repositories *ProductService) AddProduct(product entity.Product) entity.Product {
//	repositories.Products = append(repositories.Products, product)
//	return product
//}
//func (repositories *ProductService) FindAllProducts() []entity.Product {
//	return repositories.Products
//}
