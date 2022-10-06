package controller

import (
	"github.com/gin-gonic/gin"
	"projectGoLang/entity"
	"projectGoLang/service"
)

type Controller struct {
	Repository service.Repository
}

func New(repo service.Repository) Controller {
	return Controller{
		Repository: repo,
	}
}

func (c *Controller) FindOneProduct(ctx *gin.Context) entity.Product {
	ProductID := ctx.Param("id")
	return c.Repository.FindOneProduct(ProductID)
}
func (c *Controller) FindAllProducts(ctx *gin.Context) []entity.Product {
	return c.Repository.FindAllProducts()
}
func (c *Controller) CreateOneProduct(ctx *gin.Context) entity.Product {
	prod := entity.Product{}
	ctx.BindJSON(&prod)
	return c.Repository.CreateProduct(&prod)
}
func (c *Controller) DeleteOneProduct(ctx *gin.Context) entity.Product {
	ProductID := ctx.Param("id")
	return c.Repository.DeleteProductByID(ProductID)
}
