package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/app/services"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/utils"
	"strconv"
)

type Controller struct {
	Service services.Service
}

func New(service services.Service) Controller {
	return Controller{
		Service: service,
	}
}

func (c *Controller) FindOneProduct(ctx *gin.Context) {
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	prod, err := c.Service.SelectProduct(ProductIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"product": prod,
	})
}

func (c *Controller) FindAllProducts(ctx *gin.Context) {
	prods, err := c.Service.SelectAllProducts()
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"product": prods,
	})
}
func (c *Controller) CreateOneProduct(ctx *gin.Context) {
	prod := entity.Product{}
	err := ctx.BindJSON(&prod)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.AddProduct(prod)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}

func (c *Controller) DeleteOneProduct(ctx *gin.Context) {
	ProductId := ctx.Param("id")
	ProductIdInt, err := strconv.Atoi(ProductId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.DeleteProduct(ProductIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful",
	})
}

func (c *Controller) UpdateOneProduct(ctx *gin.Context) {
	prod := entity.Product{}
	err := ctx.BindJSON(&prod)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.UpdateProduct(prod)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
