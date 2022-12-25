package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/utils"
	"strconv"
)

func (c *Controller) AddBrandController(ctx *gin.Context) {
	brand := ctx.Param("brand_name")
	err := c.Service.AddToBrands(brand)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}
func (c *Controller) DeleteBrandController(ctx *gin.Context) {
	brand := ctx.Param("brand_name")
	err := c.Service.DeleteBrand(brand)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (c *Controller) SelectBrandController(ctx *gin.Context) {
	brandId := ctx.Param("id")
	brandIdInt, err := strconv.Atoi(brandId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	brand, err := c.Service.SelectBrand(brandIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"delivery": brand,
	})
}
func (c *Controller) UpdateBrandController(ctx *gin.Context) {
	brand := entity.Brand{}
	err := ctx.BindJSON(&brand)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.UpdateBrand(brand)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
