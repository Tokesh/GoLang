package controller

import (
	"github.com/gin-gonic/gin"
	"projectGoLang/entity"
	"projectGoLang/service"
)

type ProductController struct {
	Repository service.Repository
}

func New(repo service.Repository) ProductController {
	return ProductController{
		Repository: repo,
	}
}

func (c *ProductController) FindOneProduct(ctx *gin.Context) entity.Product {
	ProductID := ctx.Param("id")
	return c.Repository.FindOneProduct(ProductID)
}
func (c *ProductController) FindAllProducts(ctx *gin.Context) []entity.Product {
	return c.Repository.FindAllProducts()
}
func (c *ProductController) CreateOneProduct(ctx *gin.Context) entity.Product {
	prod := entity.Product{}
	ctx.BindJSON(&prod)
	return c.Repository.CreateProduct(&prod)
}
func (c *ProductController) DeleteOneProduct(ctx *gin.Context) entity.Product {
	ProductID := ctx.Param("id")
	return c.Repository.DeleteProductByID(ProductID)
}
