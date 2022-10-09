package controller

import (
	"projectGoLang/source/app/services"
)

type Controller struct {
	Service services.Service
}

func New(service services.Service) Controller {
	return Controller{
		Service: service,
	}
}

//func (c *Controller) FindOneProduct(ctx *gin.Context) entity.Product {
//	ProductID := ctx.Param("id")
//	return c.Repository.FindOneProduct(ProductID)
//}
//func (c *Controller) FindAllProducts(ctx *gin.Context) []entity.Product {
//	return c.Repository.FindAllProducts()
//}
//func (c *Controller) CreateOneProduct(ctx *gin.Context) entity.Product {
//	prod := entity.Product{}
//	ctx.BindJSON(&prod)
//	return c.Repository.CreateProduct(&prod)
//}
//func (c *Controller) DeleteOneProduct(ctx *gin.Context) entity.Product {
//	ProductID := ctx.Param("id")
//	return c.Repository.DeleteProductByID(ProductID)
//}
