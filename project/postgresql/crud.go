package postgresql

import (
	"context"
	"projectGoLang/entity"
)

type repository struct {
	client Client
	//logger *logging.Logger
}

func NewRepository(client Client) *repository {
	return &repository{
		client: client,
		//logger add TODO
	}
}

func (r *repository) FindOne(ctx context.Context, id string) (entity.Product, error) {
	q := `
		SELECT product_id, upc_code FROM public.products WHERE product_id = $1
	`

	var prod entity.Product
	err := r.client.QueryRow(ctx, q, id).Scan(&prod.ProductId, &prod.UpcCode)
	if err != nil {
		return entity.Product{}, err
	}

	return prod, nil
}
