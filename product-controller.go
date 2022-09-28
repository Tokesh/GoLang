package controller

import (
	"fmt"
	"projectGoLang/entity"
	"projectGoLang/service"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	AddProduct(ctx *gin.Context) entity.Product
	FindAllProducts() []entity.Product
}

type productController struct {
	service service.ProductService
}

func New(service service.ProductService) *productController {
	return &productController{
		service: service,
	}
}
func (c *productController) AddProduct(ctx *gin.Context) entity.Product {
	var product entity.Product
	ctx.BindJSON(&product)
	fmt.Println(product)
	c.service.AddProduct(product)
	return product
}
func (c *productController) FindAllProducts() []entity.Product {
	return c.service.FindAllProducts()
}
